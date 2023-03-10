#!/bin/bash

docker ps -a | grep "httpserver-in-k8s" | xargs docker rm -f
docker rmi -f $(docker images | grep "1k2222/homework-of-cloud-native-http-server-in-docker" | awk {'print $3'})
kubectl delete configmap configmap-of-httpserver
kubectl delete pod httpserver-in-k8s
kubectl delete service service-for-httpserver-in-k8s
kubectl delete secret secret-of-httpserver-in-k8s
kubectl delete deployment deployment-of-httpserver-in-k8s

rm tls.crt tls.key
