package transport

import (
	"coffeeshop/internal/orders/business"
	db "coffeeshop/internal/orders/database"
	"coffeeshop/internal/orders/grpc/pb"

	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

/// TRANSPORT LAYER (gRPC server)

type OrdersServiceServer struct {
	pb.UnimplementedOrdersServiceServer
	logic business.Logic
}

func Start(host, port, dbPath string) {
	// Connect to DB
	repo, err := db.Connect(dbPath)
	if err != nil { log.Fatalf("%v", err) }
	// Init logic
	logic := *business.New(&repo)
	// Init gRPC server
	grpcServer := grpc.NewServer()
	ordersService := OrdersServiceServer{
		logic: logic,
	}
	pb.RegisterOrdersServiceServer(grpcServer, &ordersService)
	// Start gRPC server
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil { log.Fatalf("Failed to listen: %v", err) }
	log.Printf("Orders server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
		defer repo.Close()
	}()
}
