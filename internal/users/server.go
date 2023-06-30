package users

import (
	configuration "coffeeshop/config"
	pb "coffeeshop/internal/users/pb"
	"database/sql"

	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

/// gRPC SERVER

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	db *sql.DB
}

func Start(config *configuration.Config) {
	// Connect to DB
	db, err := sql.Open("sqlite3", config.Services["users"].DB)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	UsersService := UsersServiceServer{db: db}
	pb.RegisterUsersServiceServer(grpcServer, &UsersService)
	lis, err := net.Listen("tcp", config.Host+":"+config.Services["users"].Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Users server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
		defer db.Close()
	}()
}
