# frontier

A very simple golang gorilla/negroni http server example with a few useful paradigms

## Routing

This example uses sub routers to break down paths as part of versioning API.
The sub routers use an irouter interface that adheres them to a few common methods.
Having this pattern is useful for loading routes via batch.

## Data processing

Data processing follows a worker/job double buffered channel to push/pop without delaying requests.
