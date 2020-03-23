package main

import (
	"context"
	"fmt"
	"github.com/basicsbeauty/go-grpc-hello/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const serverPort = ":8081"
const serverProtocol = "tcp"

const StatusCodeSuccess = 200

// Utility function for getting current year
func GetYear() int {
	return time.Now().Year()
}

// server for implementing HelloWorld
type server struct {
	helloworld.UnimplementedHelloWorldServiceServer
}

// HandleHelloWorld
func (s *server) HandleHelloWorld(ctx context.Context, req *helloworld.HelloWorldRequest) (*helloworld.HelloWorldResponse, error) {
	log.Println("HelloWorld Service: HelloWorld call: ", req.Input)
	message := fmt.Sprintf("Hello World from year %d", GetYear())
	return &helloworld.HelloWorldResponse{Message: message, Rcode: StatusCodeSuccess}, nil
}

func main() {
	socket, err := net.Listen(serverProtocol, serverPort)
	if err != nil {
		log.Fatalf("HelloWorld Service: Failed to bind the server: error: %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterHelloWorldServiceServer(s, &server{})
	err = s.Serve(socket)
	if err != nil {
		log.Fatalf("HelloWorld Service: Failed to start the server error:%v", err)
	}
}
