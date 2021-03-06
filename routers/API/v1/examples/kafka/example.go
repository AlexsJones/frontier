package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/AlexsJones/frontier/components/kafka"
	"github.com/AlexsJones/frontier/processing"
)

//DTO ...
type DTO struct {
	PrimeCandidate int
}

//IsPrime is a pretty simple function for our example
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

//Post is for illustrative purposes of how to use the processing API
func Post(arg1 http.ResponseWriter, arg2 *http.Request) {

	decoder := json.NewDecoder(arg2.Body)

	var e DTO
	err := decoder.Decode(&e)
	if err != nil {
		log.Println(err.Error())
		arg1.WriteHeader(http.StatusBadRequest)
		return
	}
	var j processing.Job
	//Cast the ExampleSpecificDTO into a generic DTO interface
	j.DTO = e
	//Setting up the post-processing function hook
	j.Process = func(j processing.Job) {

		var castback = j.DTO.(DTO)

		if IsPrime(castback.PrimeCandidate) {
			//If this is prime then lets send to kafka
			kafka.GetServiceBus().Produce([]byte(fmt.Sprintf("%d", castback.PrimeCandidate)))
		}
	}

	processing.GetDispatcher().JobQueue <- j

	arg1.WriteHeader(http.StatusOK)
}
