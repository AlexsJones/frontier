# frontier

A very simple golang gorilla/negroni http server example with a few useful paradigms

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
