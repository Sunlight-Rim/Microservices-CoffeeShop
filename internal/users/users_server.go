package users

import (
	pb "coffeeshop/internal/users/pb"

	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = "50051"
)

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
}

func (s *UsersServiceServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Recieved: %v, %v, %v\n", in.GetUsername(), in.GetAddress(), in.GetPass())
	// Adding recieved data to DataBase, generating hash token, generating id
	return &pb.CreateUserResponse{}, nil
}

func Start(userService UsersServiceServer) {
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, &userService)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Users server listening at %v", lis.Addr())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()
}
