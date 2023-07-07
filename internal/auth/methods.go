package auth

import (
	pb "coffeeshop/internal/auth/grpc/pb"
)

/// BUSINESS LOGIC LAYER

func Signup(username, password, address string, createUser func(username, password, address string) (UserGetter, error)) (*pb.User, error) {
	createdUser, err := createUser(username, password, address)
	return &pb.User{
		Id:       createdUser.GetId(),
		Username: createdUser.GetUsername(),
		Address:  createdUser.GetAddress(),
		Regdate:  createdUser.GetRegdate(),
	}, err
}
