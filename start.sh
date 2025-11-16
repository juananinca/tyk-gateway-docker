#!/bin/bash
docker-compose up -d --build
sleep 15
curl -X POST "localhost:8080/tyk/apis/oas/import?listenPath=/animal&validateRequest=true" -d @apidocs/animal-api.json -H 'x-tyk-authorization: foo' -H 'content-type: text/plain' 
curl -X POST "localhost:8080/tyk/apis/oas/import?listenPath=/vehicle&validateRequest=true" -d @apidocs/vehicle-api.json -H 'x-tyk-authorization: foo' -H 'content-type: text/plain' 
sleep 5
curl localhost:8080/tyk/reload/group -H 'x-tyk-authorization: foo'
