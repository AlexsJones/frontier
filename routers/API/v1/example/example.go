package example

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AlexsJones/frontier/processing"
)

//ExamplePost is for illustrative purposes of how to use the processing API
func ExamplePost(arg1 http.ResponseWriter, arg2 *http.Request) {

	// decoder := json.NewDecoder(arg2.Body)
	//
	// //Here we are directly calling the processing API DTO
	// //In reality you'll probably have some mapping between the incoming JSON object and the Job DTO
	// //We do the decoding here as otherwise we're wasting more cycles queuing it
	// var t processing.Job
	// err := decoder.Decode(&t)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	arg1.WriteHeader(http.StatusBadRequest)
	// }
	bodyBytes, err := ioutil.ReadAll(arg2.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Because our example uses our Job DTO we can process it directly into the JobQueue
	processing.GetDispatcher().JobQueue <- processing.Job{Payload: bodyBytes}

	arg1.WriteHeader(http.StatusOK)
}
