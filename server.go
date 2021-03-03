package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "hello-grpc-go/proto"
)

const (
	port = ":50051"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Reply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}