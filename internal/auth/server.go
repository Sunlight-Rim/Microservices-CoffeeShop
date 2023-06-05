package auth

import (
	pb "coffeeshop/internal/auth/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

/// gRPC SERVER

const grpcPort = "50050" // TODO: move to config
const usersPort = "50051" // TODO: move to config

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func Start() {
	// Connect to DB

	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &AuthServiceServer{})
	lis, err := net.Listen("tcp", "localhost:"+grpcPort)
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
