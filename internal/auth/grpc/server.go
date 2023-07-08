package transport

import (
	"coffeeshop/internal/auth/business"
	db "coffeeshop/internal/auth/database"
	pb "coffeeshop/internal/auth/grpc/pb"
	pbUsers "coffeeshop/internal/users/grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/// TRANSPORT LAYER (gRPC server)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	users  pbUsers.UsersServiceClient // client of Users service
	business business.Business
}

func Start(host, port, dbPort, usersPort string) {
	// TODO: Connect to Redis with dbPort
	repo, err := db.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Connect to Users
	usersConn, err := grpc.Dial(host+":"+usersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to gRPC: %v", err)
	}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	authService := AuthServiceServer{
		users:  pbUsers.NewUsersServiceClient(usersConn),
		business: business.New(&repo),
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
