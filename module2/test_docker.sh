#!/bin/bash

name=1k2222/homework-of-cloud-native-http-server-in-docker
tag=0.0.3

# 构建本地镜像
sudo docker ps -a | grep "module2" | awk '{print $1}' | xargs -r sudo docker rm -f
sudo docker images | grep $name | awk '{print $3}' | xargs -r sudo docker rmi -f
sudo docker build \
    --network host \
    -t $name:$tag .

# 将镜像推送至 docker 官方镜像仓库
#sudo docker push $name:$tag

image_id=$(sudo docker images | grep $name | awk '{print $3}')
echo "image id: $image_id"

# 通过 docker 命令本地启动 httpserver
container_id=$(sudo docker run -d -p 10080:8080 "$image_id")
echo "container id: $container_id"
pid=$(sudo docker inspect "$container_id" | grep '"Pid"' | sed 's/[^0-9]//g')
echo "pid: $pid"

# 通过 nsenter 进入容器查看 IP 配置
docker_ip=$(sudo nsenter -t "$pid" -n ip a | grep "eth0" | grep "inet" | awk '{print $2}' | awk -F'/' '{print $1}')
echo "docker ip: $docker_ip"

printf "GET http://127.0.0.1:10080/greet\n"
curl "http://127.0.0.1:10080/greet"

printf "\nGET http://127.0.0.1:10080/healthz\n"
curl "http://127.0.0.1:10080/healthz"
