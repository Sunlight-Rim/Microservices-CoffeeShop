package transport

import (
	"coffeeshop/internal/auth"
	db "coffeeshop/internal/auth/database"
	pb "coffeeshop/internal/auth/grpc/pb"
	pbUsers "coffeeshop/internal/users/grpc/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/// TRANSPORT LAYER (gRPC)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	users pbUsers.UsersServiceClient // client of Users service
	db    map[string]uint32
}

func Start(host, port, dbPort, usersPort string) {
	// TODO: Connect to Redis with dbPort
	redis := db.Connect()

	// Connect to Users
	usersConn, err := grpc.Dial(host+":"+usersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to gRPC: %v", err)
	}

	// Start gRPC server
	grpcServer := grpc.NewServer()
	authService := AuthServiceServer{
		db:    redis,
		users: pbUsers.NewUsersServiceClient(usersConn),
	}
	pb.RegisterAuthServiceServer(grpcServer, &authService)
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Auth server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
}

func (s *AuthServiceServer) Signup(ctx context.Context, in *pb.SignupAuthRequest) (*pb.SignupAuthResponse, error) {
	// Use the Users service to create a new user
	createUser := func(username, password, address string) (auth.UserGetter, error) {
		resp, err := s.users.Create(ctx, &pbUsers.CreateUserRequest{
			Username: username,
			Password: password,
			Address:  address,
		})
		return resp.User, err
	}
	user, err := auth.Signup(in.GetUsername(), in.GetPassword(), in.GetAddress(), createUser)
	return &pb.SignupAuthResponse{User: user}, err
}
