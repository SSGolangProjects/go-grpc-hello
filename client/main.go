package main

import (
	"context"
	helloworld "github.com/basicsbeauty/go-grpc-hello/api"
	"google.golang.org/grpc"
	"log"
	"time"
)

const serverAddress = "localhost:8081"

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("HelloWorld client: Failed to connect: %v", err)
	}
	defer conn.Close()

	c := helloworld.NewHelloWorldServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.HandleHelloWorld(ctx, &helloworld.HelloWorldRequest{Input: "Test"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("HelloWorld Client: HelloWork Response: %v", r)
}
