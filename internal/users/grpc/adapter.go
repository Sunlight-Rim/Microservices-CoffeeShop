package transport

import (
	"coffeeshop/internal/users/grpc/pb"

	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/// TRANSPORT LAYER (gRPC adapter)

func getUserID(ctx *context.Context) uint32 {
	// exception handling can be skipped here due to Auth() middleware
	md, _ := metadata.FromIncomingContext(*ctx)
	tokenPayload, _ := base64.StdEncoding.DecodeString(strings.Split(strings.Split(md["authorization"][0], " ")[1], ".")[1] + "==")
	var payload map[string]uint32
	json.Unmarshal(tokenPayload, &payload)
	return payload["id"]
}

func (s *UsersServiceServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.logic.Create(in.GetUsername(), in.GetPassword(), in.GetAddress())
	if err != nil { return nil, err }
	return &pb.CreateUserResponse{User: &pb.User{
		Id:       user.Id,
		Username: user.Username,
		Address:  user.Address,
		Regdate:  timestamppb.New(user.Regdate),
	}}, nil
}

func (s *UsersServiceServer) Login(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	if err := s.logic.Login(in.GetUsername(), in.GetPassword());
	err != nil { return nil, err }
	return &pb.LoginUserResponse{Access: true}, nil
}

func (s *UsersServiceServer) GetMe(ctx context.Context, empty *empty.Empty) (*pb.GetMeUserResponse, error) {
	user, err := s.logic.GetMe(getUserID(&ctx))
	if err != nil { return nil, err }
	return &pb.GetMeUserResponse{User: &pb.User{
		Id:       user.Id,
		Username: user.Username,
		Address:  user.Address,
		Regdate:  timestamppb.New(user.Regdate),
	}}, nil
}