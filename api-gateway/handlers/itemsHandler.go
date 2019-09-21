package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/ConradPacesa/k8s-backend/api-gateway/grpc"
	"google.golang.org/grpc"
)

const (
	address = "item-service:50051"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewItemClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := c.GetItem(ctx, &pb.ItemRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Item: %s\nExpires at: %d\n Completed: %t", req.Name, req.ExpiresAt, req.Completed)

	resp := fmt.Sprintf("Item: %s\nExpires at: %d\n Completed: %t", req.Name, req.ExpiresAt, req.Completed)

	io.WriteString(w, resp)
}

func NewItem(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, from the items Handler! this is the post method, to be implemented\n")
}
