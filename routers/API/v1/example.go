package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AlexsJones/frontier/processing"
)

//ExamplePost is for illustrative purposes of how to use the processing API
func ExamplePost(arg1 http.ResponseWriter, arg2 *http.Request) {

	decoder := json.NewDecoder(arg2.Body)

	//Here we are directly calling the processing API DTO
	//In reality you'll probably have some mapping between the incoming JSON object and the Job DTO
	var t processing.Job
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err.Error())
		arg1.WriteHeader(http.StatusBadRequest)
	}

	//Because our example uses our Job DTO we can process it directly into the JobQueue
	log.Println("Writing job")
	processing.GetDispatcher().JobQueue <- t
	log.Println("Wrote job")
	arg1.WriteHeader(http.StatusOK)
}
