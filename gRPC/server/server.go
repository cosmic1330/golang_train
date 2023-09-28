package main

import (
	"context"
	pb "grpc/protobuffer"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}


func main() {
	// Create gRPC Server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("gRPC server is running.")
	pb.RegisterSaySomethingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) TestRPC(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	log.Println("i receive:", in.Id)
	return &pb.TestResponse{Name: "Jay", Weight: 75, Heigh: 170}, nil
}
