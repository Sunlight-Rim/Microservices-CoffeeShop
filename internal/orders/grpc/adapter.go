package transport

import (
	orders_pb "coffeeshop/internal/orders/grpc/pb"
	pb "coffeeshop/internal/orders/grpc/pb"
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/// TRANSPORT LAYER (gRPC adapter)

func getUserID(ctx *context.Context) uint32 {
	// exception handling can be skipped here due to Auth() middleware
	md, _ := metadata.FromIncomingContext(*ctx)
	tokenPayload, _ := base64.StdEncoding.DecodeString(strings.Split(strings.Split(md["authorization"][0], " ")[1], ".")[1] + "==")
	var payload map[string]uint32
	json.Unmarshal(tokenPayload, &payload)
	return payload["id"]
}

func (s *OrdersServiceServer) Create(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order, err := s.business.Create(getUserID(&ctx), in.GetSugar(), in.GetCoffee(), in.GetTopping())
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{Order: &pb.Order{
		Id:      order.Id,
		Userid:  order.Userid,
		Status:  pb.Order_Status(order.Status),
		Coffee:  order.Coffee,
		Topping: order.Topping,
		Sugar:   order.Sugar,
		Total:   order.Total,
		Date:    timestamppb.New(order.Date),
	}}, nil
}

func (s *OrdersServiceServer) Get(ctx context.Context, in *pb.GetOneOrderRequest) (*pb.GetOneOrderResponse, error) {
	order, err := s.business.Get(getUserID(&ctx), in.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetOneOrderResponse{Order: &pb.Order{
		Id:      order.Id,
		Userid:  order.Userid,
		Status:  pb.Order_Status(order.Status),
		Coffee:  order.Coffee,
		Topping: order.Topping,
		Sugar:   order.Sugar,
		Total:   order.Total,
		Date:    timestamppb.New(order.Date),
	}}, nil
}

func (s *OrdersServiceServer) GetSome(ctx context.Context, in *pb.GetSomeOrderRequest) (*pb.GetSomeOrderResponse, error) {
	orders, err := s.business.GetSome(getUserID(&ctx), in.GetShift())
	if err != nil {
		return nil, err
	}
	var ordersPb []*orders_pb.Order
	for _, order := range orders {
		ordersPb = append(ordersPb, &orders_pb.Order{
			Id:      order.Id,
			Userid:  order.Userid,
			Status:  pb.Order_Status(order.Status),
			Coffee:  order.Coffee,
			Topping: order.Topping,
			Sugar:   order.Sugar,
			Total:   order.Total,
			Date:    timestamppb.New(order.Date),
		})
	}
	return &pb.GetSomeOrderResponse{Orders: ordersPb}, nil
}
