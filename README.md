# frontier

A very simple golang gorilla/negroni http server example with a few useful paradigms.
I hope this goes some small way to demonstrate a few useful ideas around create a high performance API.

<img src="https://i.imgur.com/HpKOfUt.png" width="580"/>


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

To run the example please set the following google credentials in your shell
They are required to upload the test strings to GCS (you can replace this with some other example)
```
projectID  = os.Getenv("GOOGLE_PROJECT")
bucketName = os.Getenv("GOOGLE_BUCKET")
```
To remove the example please remove `routers/API/v1/example` folder
Also remove `API/v1/router.go` with reference to the example here:
```
	//Example route to demonstrate processing components
	d.BaseRouter.Router.HandleFunc("/processor", example.ExamplePost).Methods("POST")
```

### testing

Change into the test directory whilst running the API in another tab

This load test requires `vegeta`

```
./load_test.sh
```

or you can run the go tests `go run test.go`
