package main

import (
	"context"
	"fmt"
	"io"
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
	// AddUser(client)
	AddUserStream(client)
}

// test request to service gRCP - API unary
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

// test request to service gRCP - API server streaming
func AddUserStream(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "456",
		Name:  "Maria",
		Email: "maria@test.com",
	}

	res, err := client.AddUserStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}
		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}
