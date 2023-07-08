package transport

import (
	"coffeeshop/internal/auth/grpc/pb"
	pbUsers "coffeeshop/internal/users/grpc/pb"

	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

/// TRANSPORT LAYER (gRPC adapters)

// Signup new user
func (s *AuthServiceServer) Signup(ctx context.Context, in *pb.SignupAuthRequest) (*pb.SignupAuthResponse, error) {
	// Use the Users service to create a new user in Users service DB
	createUser := func(username, password, address string) (uint32, *time.Time, error) {
		resp, err := s.users.Create(ctx, &pbUsers.CreateUserRequest{
			Username: username,
			Password: password,
			Address:  address,
		})
		if err != nil {
			return 0, nil, err
		}
		date := resp.User.GetRegdate().AsTime()
		return resp.User.GetId(), &date, nil
	}
	user, err := s.logic.Signup(in.GetUsername(), in.GetPassword(), in.GetAddress(), createUser)
	if err != nil {
		return nil, err
	}
	return &pb.SignupAuthResponse{User: &pb.User{
		Id:       user.Id,
		Username: user.Username,
		Address:  user.Address,
		Regdate:  timestamppb.New(*user.Regdate),
	}}, nil
}
