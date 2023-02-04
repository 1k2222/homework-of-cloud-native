#!/bin/bash

echo "1. build docker"
if [ $(docker images | grep "1k2222/homework-of-cloud-native-http-server-in-docker" | grep "0.0.3" | wc -l) = "0" ]; then
    cd ../module2
   chmod u+x docker_build.sh
   ./docker_build.sh
   cd -
else
   echo "已经build过，跳过"
fi

echo "2. 部署k8s"
kubectl create configmap configmap-of-httpserver --from-file=config.properties
kubectl create -f pod.yaml
kubectl create -f nodeport.yaml
