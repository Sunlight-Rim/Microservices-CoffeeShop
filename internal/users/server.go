package users

import (
	pb "coffeeshop/internal/users/pb"
	"database/sql"

	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

/// gRPC SERVER

const grpcPort = "50052" // TODO: move to config

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	db *sql.DB
}

func Start() {
	// Connect to DB
	db, err := sql.Open("sqlite3", "internal/users/database/users.db")
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	userService := UsersServiceServer{db: db}
	pb.RegisterUsersServiceServer(grpcServer, &userService)
	lis, err := net.Listen("tcp", "localhost:"+grpcPort)
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
