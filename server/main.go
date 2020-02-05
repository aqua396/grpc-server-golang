package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "../pb"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreetServer
}

func (s *server) Say(ctx context.Context, in *pb.Question) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.Reply{Message: "Hello"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
