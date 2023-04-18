package orders

import (
	pb "coffeeshop/internal/orders/pb"

	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	grpcPort = "50051"
)

/// SERVER DEFINITION

type OrdersServiceServer struct {
	pb.UnimplementedOrdersServiceServer
}

func Start() {
	grpcServer := grpc.NewServer()
	orderService := OrdersServiceServer{}
	pb.RegisterOrdersServiceServer(grpcServer, &orderService)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Orders server listening at %v", lis.Addr())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
}

/// API METHODS (gRPC)

func (s *OrdersServiceServer) Create(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("Recieved: %v", in.GetCoffees())
	// Adding recieved data to DataBase, generating id
	return &pb.CreateOrderResponse{Order: &pb.Order{Id: 1, Total: 1642}}, nil
}
