package users

import (
	db "coffeeshop/internal/users/database"
	pb "coffeeshop/internal/users/grpc/pb"
	"database/sql"

	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

/// TRANSPORT LAYER (gRPC)

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	db *sql.DB
}

func Start(host, port, dbPath string) {
	// Connect to DB
	db, err := db.Connect(dbPath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	UsersService := UsersServiceServer{db: db}
	pb.RegisterUsersServiceServer(grpcServer, &UsersService)
	lis, err := net.Listen("tcp", host+":"+port)
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
