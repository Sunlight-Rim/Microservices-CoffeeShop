package auth

import (
	pb "coffeeshop/internal/auth/pb"
	users_pb "coffeeshop/internal/users/pb"
	"context"
)

func (s *AuthServiceServer) Signup(ctx context.Context, in *pb.SignupAuthRequest) (*pb.SignupAuthResponse, error) {
	// Use the Users service to create new user
	creationResp, err := s.users.Create(ctx, &users_pb.CreateUserRequest{
		Username: in.GetUsername(),
		Address:  in.GetAddress(),
		Password: in.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.SignupAuthResponse{User: &pb.User{
		Id:       creationResp.User.GetId(),
		Username: creationResp.User.GetUsername(),
		Address:  creationResp.User.GetAddress(),
		Regdate:  creationResp.User.GetRegdate(),
	}}, nil
}
