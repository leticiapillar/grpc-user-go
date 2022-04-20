package main

import (
	"context"
	"fmt"
	"log"

	"github.com/leticiapillar/grpc-user-go/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)
}

// test request to service
func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "321",
		Name:  "Leticia",
		Email: "leticia@test.com",
	}
	
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}
