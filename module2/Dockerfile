FROM ubuntu:22.04

WORKDIR ~/module2

COPY . .

RUN apt update -y && apt install ca-certificates curl golang -y && go mod tidy && go build .

ENTRYPOINT ["./module2"]
