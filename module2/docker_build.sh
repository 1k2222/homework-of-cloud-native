#!/bin/bash

name=1k2222/homework-of-cloud-native-http-server-in-docker
tag=0.0.4

sudo docker build \
    --network host \
    -t $name:$tag .
