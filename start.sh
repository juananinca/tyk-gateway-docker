#!/bin/bash
docker-compose up -d --build
sleep 15
mkdir -p tyk_apis

API_ANIMAL_ID=$(curl \
    -X POST "localhost:8080/tyk/apis/oas/import?listenPath=/animal&validateRequest=true" \
    -d @apidocs/animal-api.json \
    -H 'x-tyk-authorization: foo' \
    -H 'content-type: text/plain' | jq -r '.key')
sleep 5
API_VEHICLE_ID=$(curl \
    -X POST "localhost:8080/tyk/apis/oas/import?listenPath=/vehicle&validateRequest=true" \
    -d @apidocs/vehicle-api.json \
    -H 'x-tyk-authorization: foo' \
    -H 'content-type: text/plain' | jq -r '.key')
curl localhost:8080/tyk/reload/group -H 'x-tyk-authorization: foo'
sleep 5
echo "API_ANIMAL_ID: $API_ANIMAL_ID"
curl \
    -X GET "localhost:8080/tyk/apis/oas/$API_ANIMAL_ID" \
    -H 'x-tyk-authorization: foo' \
    -o tyk_apis/animal-api.json
sleep 5
echo "API_VEHICLE_ID: $API_VEHICLE_ID"
curl \
    -X GET "localhost:8080/tyk/apis/oas/$API_VEHICLE_ID" \
    -H 'x-tyk-authorization: foo' \
    -o tyk_apis/vehicle-api.json
sleep 5
