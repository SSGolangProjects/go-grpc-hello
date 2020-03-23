package main

import (
	"context"
	helloworld "github.com/basicsbeauty/go-grpc-hello/api"
	"google.golang.org/grpc"
	"log"
	"strconv"
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

	counter := 10
	for i := 0; i < counter; i++ {
		r, err := c.HandleHelloWorld(ctx, &helloworld.HelloWorldRequest{Input: strconv.Itoa(i)})
		if err != nil {
			log.Fatalf("HelloWorld Client: HelloWork call: Failed: %v", err)
		}
		log.Printf("HelloWorld Client: HelloWork Response: [%d:%d] %v", i, counter, r)
	}
}
