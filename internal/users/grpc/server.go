package transport

import (
	"coffeeshop/internal/users/business"
	db "coffeeshop/internal/users/database"
	"coffeeshop/internal/users/grpc/pb"

	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

/// TRANSPORT LAYER (gRPC server)

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	logic business.Logic
}

func Start(host, port, dbPath string) {
	// Connect to DB
	repo, err := db.Connect(dbPath)
	if err != nil { log.Fatalf("%v", err) }
	// Init gRPC server
	grpcServer := grpc.NewServer()
	usersService := UsersServiceServer{
		logic: business.New(&repo),
	}
	pb.RegisterUsersServiceServer(grpcServer, &usersService)
	// Start gRPC server
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil { log.Fatalf("Failed to listen: %v", err) }
	log.Printf("Users server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
		defer repo.Close()
	}()
}
