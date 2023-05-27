package orders

import (
	pb "coffeeshop/internal/orders/pb"
	"database/sql"

	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

/// gRPC SERVER

const ( // TODO: move to config
	grpcPort  = "50051"
	usersPort = "50052"
)

type OrdersServiceServer struct {
	pb.UnimplementedOrdersServiceServer
	db    *sql.DB
}

func Start() {
	// Connect to DB
	db, err := sql.Open("sqlite3", "internal/orders/database/orders.db")
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	ordersService := OrdersServiceServer{db: db}
	pb.RegisterOrdersServiceServer(grpcServer, &ordersService)
	lis, err := net.Listen("tcp", "localhost:"+grpcPort)
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