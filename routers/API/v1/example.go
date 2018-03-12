package v1

import (
	"encoding/json"
	"log"
	"net/http"
)

//ExamplePostRequest is a DTO
type ExamplePostRequest struct {
	Number string
}

//ExamplePost is for illustrative purposes of how to use the processing API
func ExamplePost(arg1 http.ResponseWriter, arg2 *http.Request) {

	decoder := json.NewDecoder(arg2.Body)

	var t ExamplePostRequest
	err := decoder.Decode(&t)
	if err != nil {
		log.Println(err.Error())
		arg1.WriteHeader(http.StatusBadRequest)
	}

	arg1.WriteHeader(http.StatusOK)
}
