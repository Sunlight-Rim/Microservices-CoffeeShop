package auth

import (
	configuration "coffeeshop/config"
	pb "coffeeshop/internal/auth/pb"
	pbUsers "coffeeshop/internal/users/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/// gRPC SERVER

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	users pbUsers.UsersServiceClient // client of Users service
	redis map[string]uint32
}

func Start(config *configuration.Config) {
	// TODO: Connect to Redis with config.Services["auth"].DB
	redis := make(map[string]uint32)

	// Connect to Users
	usersConn, err := grpc.Dial(config.Host+":"+config.Services["users"].Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to gRPC: %v", err)
	}

	// Start gRPC server
	grpcServer := grpc.NewServer()
	authService := AuthServiceServer{redis: redis, users: pbUsers.NewUsersServiceClient(usersConn)}
	pb.RegisterAuthServiceServer(grpcServer, &authService)
	lis, err := net.Listen("tcp", config.Host+":"+config.Services["auth"].Port)
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
