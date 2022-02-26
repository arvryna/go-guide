#! /bin/bash

echo "Package app in docker container ..."
docker build -t goguide/greeter .

# To run the container directly..
# echo "Run the container ..."
# docker run --rm goguide/greeter

echo "Lauch kubernetes"
kubectl apply -f ./greeter.yaml
