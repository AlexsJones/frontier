package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/paulbellamy/ratecounter"
)

func main() {

	client := &http.Client{}
	counter := ratecounter.NewRateCounter(1 * time.Second)
	for i := 0; i < 50000; {
		var jsonStr = []byte(fmt.Sprintf(`{ "PrimeCandidate": %d}`, i))
		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/processor", bytes.NewBuffer(jsonStr))

		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		i++
		counter.Incr(1)
	}
	fmt.Printf("Requests per second: %d\n", counter.Rate())
}
