package users

import (
	pb "coffeeshop/internal/users/pb"

	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = "50052"
)

/// SERVER DEFINITION

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
}

func Start() {
	grpcServer := grpc.NewServer()
	userService := UsersServiceServer{}
	pb.RegisterUsersServiceServer(grpcServer, &userService)

	lis, err := net.Listen("tcp", "localhost:"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Users server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
}

/// API METHODS (gRPC)

func (s *UsersServiceServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Recieved: %v, %v, %v\n", in.GetUsername(), in.GetAddress(), in.GetPassword())
	// Check username in DB, add password to DB, generate token, generate id, add and response timestamp
	return &pb.CreateUserResponse{User: &pb.User{
		Id:       1,
		Username: in.GetUsername(),
		Address:  in.GetAddress(),
		OrderIds: []int64{1, 2, 3},
	}}, nil
}
