package kafka

import (
	"fmt"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

/*
Example local kafka configuration
./zookeeper-server-start.sh config/zookeeper.properties
./kafka-server-start.sh config/server.properties
./kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test
*/

//ServiceBus holds the type information and hinged method references for the kafka producer
type ServiceBus struct {
	Producer *kafka.Producer
}

var (
	instance        *ServiceBus
	once            sync.Once
	kafkaserverlist = os.Getenv("KAFKA_SERVER_LIST")
	kafkatopic      = os.Getenv("KAFKA_PRODUCER_TOPIC")
)

//GetServiceBus for kafka
func GetServiceBus() *ServiceBus {
	once.Do(func() {

		instance = &ServiceBus{}

		if kafkaserverlist == "" || kafkatopic == "" {
			fmt.Println("Kafka requires both KAFKA_SERVER_LIST and KAFKA_PRODUCER_TOPIC envs set.")
			os.Exit(1)
		}

		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
		if err != nil {
			panic(err)
		}
		// Delivery report handler for produced messages
		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()
		instance.Producer = p
	})
	return instance
}

//Produce will send messages to the kafka topic defined in the env, or can be overriden
func (s *ServiceBus) Produce(data []byte, optTopic ...string) {

	var currentTopic = kafkatopic
	if len(optTopic) > 0 {
		if optTopic[0] != "" {
			//Use the optional override if it exists
			currentTopic = optTopic[0]
		}
	}
	s.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &currentTopic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	// Wait for message deliveries
	//p.Flush(15 * 1000)
}
