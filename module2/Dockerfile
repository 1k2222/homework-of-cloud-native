FROM ubuntu:22.04

WORKDIR ~/module2

COPY . .

ENV HTTP_PROXY 192.168.31.107:2022
ENV HTTPS_PROXY 192.168.31.107:2022

RUN apt update -y && apt install ca-certificates curl golang -y && go mod tidy && go build .

ENTRYPOINT ["./module2"]