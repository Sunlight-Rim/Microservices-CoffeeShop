package transport

import (
	"coffeeshop/internal/orders/grpc/pb"

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
	order, err := s.logic.Create(getUserID(&ctx), in.GetSugar(), in.GetCoffee(), in.GetTopping())
	if err != nil { return nil, err }
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
	order, err := s.logic.Get(getUserID(&ctx), in.GetId())
	if err != nil { return nil, err }
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
	orders, err := s.logic.GetSome(getUserID(&ctx), in.GetShift())
	if err != nil { return nil, err }
	var ordersPb []*pb.Order
	for _, order := range orders {
		ordersPb = append(ordersPb, &pb.Order{
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

func (s *OrdersServiceServer) Cancel(ctx context.Context, in *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	order, err := s.logic.Cancel(getUserID(&ctx), in.GetId())
	if err != nil { return nil, err }
	return &pb.CancelOrderResponse{Order: &pb.Order{
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

func (s *OrdersServiceServer) Delete(ctx context.Context, in *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	order, err := s.logic.Delete(getUserID(&ctx), in.GetId())
	if err != nil { return nil, err }
	return &pb.DeleteOrderResponse{Order: &pb.Order{
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