# go-grpc-hello
Hello world using Golang and gRPC

## There are two parts for this POC
- Server
- Client

Both components use gRPC for communication

## Run server unit test(s)
Note that running unit test(s) is incorporated when building docker image
```
go test server/*.go
```

## Networking
In the docker-compose definition both server and client were setup use `host` (localhost as seen by container)

## Dependencies
- Golang
- gRPC (along with Protobuf)
- Docker (and docker-compose)

## Setup and run
Run below commands
```bash
git clone https://github.com/basicsbeauty/go-grpc-hello
cd go-grpc-hello
docker-compose up
```

# Diagram
![Alt text](diagram.png?raw=true "Go gRPC Hello World")
