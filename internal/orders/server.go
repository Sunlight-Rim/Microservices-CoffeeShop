package orders

import (
	configuration "coffeeshop/config"
	pb "coffeeshop/internal/orders/pb"
	"database/sql"

	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

/// gRPC SERVER

type OrdersServiceServer struct {
	pb.UnimplementedOrdersServiceServer
	db *sql.DB
}

func Start(config *configuration.Config) {
	// Connect to DB
	db, err := sql.Open("sqlite3", config.Services["orders"].DB)
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	ordersService := OrdersServiceServer{db: db}
	pb.RegisterOrdersServiceServer(grpcServer, &ordersService)
	lis, err := net.Listen("tcp", config.Host+":"+config.Services["orders"].Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Orders server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
		defer db.Close()
	}()
}
