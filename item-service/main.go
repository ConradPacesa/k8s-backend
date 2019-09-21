package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ConradPacesa/k8s-backend/item-service/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetItem(ctx context.Context, in *pb.ItemRequest) (*pb.ItemReply, error) {
	log.Printf("Recieved request")
	return &pb.ItemReply{
		Id:        1,
		Name:      "hello this is your item to do",
		ExpiresAt: 123,
		Completed: false,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterItemServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
