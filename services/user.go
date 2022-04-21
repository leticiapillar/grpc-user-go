package services

import (
	"context"
	"fmt"
	"time"

	"github.com/leticiapillar/grpc-user-go/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// Service gRPC API Unary
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println(req.GetName())

	return &pb.User{
		Id: "123",
		Name: req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

// Service gRPC API Server Streaming
func (*UserService) AddUserStream(req *pb.User, stream pb.UserService_AddUserStreamServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "init",
		User: &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "inserting",
		User: &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)
	
	stream.Send(&pb.UserResultStream{
		Status: "completed",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)
	return nil
}
