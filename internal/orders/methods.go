package orders

import (
	pb "coffeeshop/internal/orders/grpc/pb"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/// BUSINESS LOGIC LAYER

func getUserID(ctx *context.Context) uint32 {
	// exception handling can be skipped here due to Auth() middleware
	md, _ := metadata.FromIncomingContext(*ctx)
	tokenPayload, _ := base64.StdEncoding.DecodeString(strings.Split(strings.Split(md["authorization"][0], " ")[1], ".")[1] + "==")
	var payload map[string]uint32
	json.Unmarshal(tokenPayload, &payload)
	return payload["id"]
}

func (s *OrdersServiceServer) Create(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Validate input
	if len(in.GetCoffee()) == 0 {
		return nil, errors.New("you didn't specify any coffee")
	}
	// Create order
	var (
		coffeeID  uint32
		toppingID uint32 = 1
		order            = &pb.Order{
			Userid:  getUserID(&ctx),
			Coffee:  in.GetCoffee(),
			Topping: in.GetTopping(),
			Sugar:   in.GetSugar(),
			Date:    timestamppb.New(time.Now()),
			Status:  pb.Order_Status(0),
		}
	)
	// Sum coffee price & get id
	if err := s.db.QueryRow(
		`SELECT coffeeID, price FROM coffee WHERE name = $1;`,
		order.Coffee).Scan(&coffeeID, &order.Total); err != nil {
		return nil, errors.New("specified coffee type was wrong")
	}
	// Sum topping price & get id
	if order.Topping != "" {
		var price float32
		if err := s.db.QueryRow(
			`SELECT toppingID, price FROM topping WHERE name = $1;`,
			order.Topping).Scan(&toppingID, &price); err != nil {
			return nil, errors.New("specified topping type was wrong")
		}
		order.Total += price
	}
	// Append to DB
	res, err := s.db.Exec(
		`INSERT INTO order_ (userID, coffeeID, toppingID, sugar) VALUES ($1, $2, $3, $4);`,
		order.Userid, coffeeID, toppingID, order.Sugar)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	id, _ := res.LastInsertId()
	order.Id = uint32(id)
	return &pb.CreateOrderResponse{Order: order}, nil
}

func (s *OrdersServiceServer) Get(ctx context.Context, in *pb.GetOneOrderRequest) (*pb.GetOneOrderResponse, error) {
	// Validate input
	if in.GetId() == 0 {
		return nil, errors.New("order ID is wrong")
	}
	var (
		toppingPrice float32
		date         time.Time
		order        = &pb.Order{
			Id:     in.GetId(),
			Userid: getUserID(&ctx),
		}
	)
	if err := s.db.QueryRow(
		`SELECT coffee.name, coffee.price, topping.name, topping.price, sugar, status, date
		FROM order_ INNER JOIN
			 coffee ON order_.coffeeID = coffee.coffeeID INNER JOIN
			 topping ON order_.toppingID = topping.toppingID
		WHERE order_.orderID = $1 AND order_.userID = $2;`,
		order.Id, order.Userid).Scan(&order.Coffee, &order.Total, &order.Topping,
		&toppingPrice, &order.Sugar, &order.Status, &date); err != nil {
		return nil, err
	}
	order.Date = timestamppb.New(date)
	order.Total += toppingPrice
	return &pb.GetOneOrderResponse{Order: order}, nil
}

func (s *OrdersServiceServer) GetSome(ctx context.Context, in *pb.GetSomeOrderRequest) (*pb.GetSomeOrderResponse, error) {
	userID := getUserID(&ctx)
	rows, err := s.db.Query(
		`SELECT orderID, coffee.name, coffee.price, topping.name, topping.price, sugar, status, date
		FROM order_ INNER JOIN
			 coffee ON order_.coffeeID = coffee.coffeeID INNER JOIN
			 topping ON order_.toppingID = topping.toppingID
		WHERE order_.userID = $1
		LIMIT 5 OFFSET $2;`,
		userID, in.GetShift())
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	defer rows.Close()
	var (
		date         time.Time
		orders       []*pb.Order
		toppingPrice float32
	)
	for rows.Next() {
		order := pb.Order{Userid: userID}
		rows.Scan(&order.Id, &order.Coffee, &order.Total, &order.Topping,
			&toppingPrice, &order.Sugar, &order.Status, &date)
		order.Date = timestamppb.New(date)
		orders = append(orders, &order)
	}
	return &pb.GetSomeOrderResponse{Orders: orders}, nil
}

func (s *OrdersServiceServer) Cancel(ctx context.Context, in *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	// Get changing order
	getResponce, err := s.Get(ctx, &pb.GetOneOrderRequest{
		Id: in.GetId(),
	})
	if err != nil {
		return nil, err
	}
	// Validate order status
	switch getResponce.Order.Status {
	case 1:
		return nil, errors.New("order was already DELIVERED")
	case 2:
		return nil, errors.New("order was already CANCELLED")
	}
	// Change order status to CANCELLED
	if _, err := s.db.Exec(
		`UPDATE order_ SET status = 2 WHERE orderID = $1 AND userID = $2;`,
		getResponce.Order.Id, getResponce.Order.Userid); err != nil {
		return nil, err
	}
	getResponce.Order.Status = 2
	return &pb.CancelOrderResponse{Order: getResponce.Order}, nil
}

func (s *OrdersServiceServer) Delete(ctx context.Context, in *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	// Validate input
	if in.GetId() == 0 {
		return nil, errors.New("order ID is wrong")
	}
	// Get deleting order
	getResponce, err := s.Get(ctx, &pb.GetOneOrderRequest{
		Id: in.GetId(),
	})
	if err != nil {
		return nil, err
	}
	// Delete
	if _, err := s.db.Exec(
		`DELETE FROM order_ WHERE orderID = $1 AND userID = $2;`,
		getResponce.Order.Id, getResponce.Order.Userid); err != nil {
		return nil, err
	}
	return &pb.DeleteOrderResponse{Order: getResponce.Order}, nil
}
