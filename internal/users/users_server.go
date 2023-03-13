package users

import (
	pb "coffeeshop/internal/users/pb"

	"context"
	"log"
)

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
}

func (s *UsersServiceServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Recieved: %v, %v, %v\n", in.GetName(), in.GetAddress(), in.GetPass())
	// Adding recieved data to DataBase, generating hash token, generating id
	return &pb.CreateUserResponse{Token: in.GetPass()}, nil
}
