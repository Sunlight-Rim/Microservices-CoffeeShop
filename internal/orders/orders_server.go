package orders

import (
	pb "coffeeshop/internal/orders/pb"

	"context"
	"log"
)

type OrdersServiceServer struct {
	pb.UnimplementedOrdersServiceServer
}

func (s *OrdersServiceServer) Create(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("Recieved: %v", in.GetCoffees())
	// Adding recieved data to DataBase, generating id
	return &pb.CreateOrderResponse{}, nil
}
