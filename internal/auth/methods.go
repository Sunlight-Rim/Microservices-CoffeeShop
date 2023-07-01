package auth

import (
	pb "coffeeshop/internal/auth/pb"
	"context"
)

func (s *AuthServiceServer) Signup(ctx context.Context, in *pb.SignupAuthRequest) (*pb.SignupAuthResponse, error) {
	return &pb.SignupAuthResponse{}, nil
}
