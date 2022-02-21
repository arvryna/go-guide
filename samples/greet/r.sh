#! /bin/bash

echo "Package app in docker container ..."
docker build -t greet .

echo "Run the container ..."
docker run --rm greet
