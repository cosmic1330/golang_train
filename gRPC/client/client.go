package main

import (
	"context"
	pb "grpc/protobuffer"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSaySomethingServiceClient(conn)
	TestRPC(c, 12)

}

func TestRPC(c pb.SaySomethingServiceClient, id int64) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.TestRPC(ctx, &pb.TestRequest{Id: id})
	if err != nil {
		log.Fatalf("could not mkdfk2020: %v", err)
	}
	log.Printf("gRPC response: %s", res)
}
