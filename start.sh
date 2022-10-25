#!/bin/bash

# Check if we're root and re-execute if we're not.
rootcheck() {
    if [ $(id -u) != "0" ]; then
        sudo "$0" "$@" # Modified as suggested below.
        exit $?
    fi
}

rootcheck "${@}"

$APIPORT = 8888                                                                                                         # The API Port the api will hear
sudo docker network create problem-network                                                                              # Create a docker network to make possible two or more containers running in the same host communicate
sudo docker run -d -e POSTGRES_PASSWORD=postgres --name=problem-database --network=problem-network postgres:14.5-alpine # Run a POSTGRE container and connect it to docker network
IP=$(sudo docker network inspect problem-network | grep -o -P "(?<=\"IPv4Address\": \").*(?=/)")                        # Discovery the postgree IP in the docker network
sudo docker build --build-arg port=$APIPORT -t pc/test-project .                                                        # Build the image of backend application
sudo docker run -e APIPORT=$APIPORT -e DATABASEIP=$IP --network=problem-network pc/test-project                         # Run a container from the image created before
