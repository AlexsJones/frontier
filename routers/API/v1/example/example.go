package example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/AlexsJones/frontier/processing"
)

//ExampleSpecificDTO ...
type ExampleSpecificDTO struct {
	PrimeCandidate int
}

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

	var e ExampleSpecificDTO
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

		var castback = j.DTO.(ExampleSpecificDTO)

		dest := os.Getenv("REQUESTBIN_URL")

		if dest == "" {
			fmt.Printf("Please set REQUESTBIN_URL to run this example\n")
			os.Exit(1)
		}

		if IsPrime(castback.PrimeCandidate) {

			//Further processing
			client := &http.Client{}
			var jsonStr = []byte(fmt.Sprintf(`{ "PrimeFound": %d }`, castback.PrimeCandidate))
			req, err := http.NewRequest("POST", dest, bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
		}
	}

	processing.GetDispatcher().JobQueue <- j

	arg1.WriteHeader(http.StatusOK)
}
