# frontier


[![Go Report Card](https://goreportcard.com/badge/github.com/AlexsJones/frontier)](https://goreportcard.com/report/github.com/AlexsJones/frontier)
[![Maintainability](https://api.codeclimate.com/v1/badges/3eceeb49838553cb94fb/maintainability)](https://codeclimate.com/github/AlexsJones/frontier/maintainability)

<img src="https://i.imgur.com/HpKOfUt.png" width="580"/>


A very simple golang gorilla http server example with a few useful paradigms.
I hope this goes some small way to demonstrate a few useful ideas around create a high performance API.

## Routing

This example uses sub routers to break down paths as part of versioning API.
The sub routers use an irouter interface that adheres them to a few common methods.
Having this pattern is useful for loading routes via batch.

The current structure is the following:
```
Base router ---> /
API router --------> /api
V1 router --------------> /v1
```

This means using the V1 router within `routers/API/v1/router.go` is ideal for grouping API to this version.

## Data processing

Data processing follows a worker/job via [channel over channel](https://www.goin5minutes.com/blog/channel_over_channel/) to push/pop without delaying requests.

## The example

To run the example please grab a free requestbin URL to forward our requests too
This is required to upload the test JSON once primes have been processed
e.g...
```
export REQUESTBIN_URL="https://requestb.in/poe29ypo"
```

Change into the test directory whilst running the API in another tab

This load test requires `vegeta` a golang package

```
cd test
go run test.go #Runs the incremental prime tests
./load_test.sh #Runs a static POST request scan
```

### Example code throughput

The example is pretty arbitrary but hopefully demonstrates it is perfectly possibly to process
many thousands of requests per second.

Bench marking on a laptop at around 10k RP/S it is easy to imagine that load balancing several frontier API 
as containers could easily achieve over 100,000 RP/S

### Removing example code

To remove the example please remove `routers/API/v1/example` folder
Also remove `API/v1/router.go` with reference to the example here:
```
	//Example route to demonstrate processing components
	d.BaseRouter.Router.HandleFunc("/processor", example.ExamplePost).Methods("POST")
```
