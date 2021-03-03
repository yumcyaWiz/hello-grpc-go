package main

import (
	"context"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	pb "hello-grpc-go/proto"
)

const (
	address = "localhost:50051"
	defaultName = "world"
)

func main() {
	// setup a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// input user name
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("cound not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}