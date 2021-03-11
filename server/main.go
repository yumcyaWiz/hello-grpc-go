package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hello-grpc-go/proto"
	"log"
	"net"
)

const (
	port = ":50051"
)

// Server is used to implement GreeterServer
type Server struct {
	proto.UnimplementedGreeterServer
}

// implementation of SayHello
func (s *Server) SayHello(ctx context.Context, in *proto.Request) (*proto.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &proto.Reply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// listen web server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// setup gRPC server
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
