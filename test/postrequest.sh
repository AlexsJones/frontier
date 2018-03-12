#!/bin/bash

# Hits the /api/v1/processor route
curl -H "Content-Type: application/json" -X POST -d '{"payload":"yada yada"}' http://localhost:8080/api/v1/processor
