#!/bin/bash

curl --header "Content-Type: application/json" \
    --request POST \
    --data '{ "key" : "var1", "value" : "testing" }' \
    http://localhost:8000/store-variable