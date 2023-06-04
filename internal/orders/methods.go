package orders

import (
	pb "coffeeshop/internal/orders/pb"
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

/// API METHODS IMPLEMENTATION

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
	order := &pb.Order{
		Userid:  getUserID(&ctx),
		Coffee:  in.GetCoffee(),
		Topping: in.GetTopping(),
		Sugar:   in.GetSugar(),
		Date:    timestamppb.New(time.Now()),
		Status:  pb.Order_Status(0),
	}
	var coffeeID, toppingID uint32
	// Sum coffee price & get id
	if err := s.db.QueryRow(`SELECT coffeeID, price FROM coffee WHERE name == $1`,
			  order.Coffee).Scan(&coffeeID, &order.Total);
	err != nil {
		return nil, errors.New("specified coffee type was wrong")
	}
	// Sum topping price & get id
	if order.Topping != "" {
		var price float32
		if err := s.db.QueryRow(`SELECT toppingID, price FROM topping WHERE name == $1`,
				  order.Topping).Scan(&toppingID, &price);
		err != nil {
			return nil, errors.New("specified topping type was wrong")
		}
		order.Total += price
	}
	// Append to DB
	res, err := s.db.Exec(`INSERT INTO order_ (userID, coffeeID, toppingID, sugar) VALUES ($1, $2, $3, $4)`,
				order.Userid, coffeeID, toppingID, order.Sugar)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	id, _ := res.LastInsertId()
	order.Id = uint32(id)
	return &pb.CreateOrderResponse{Order: order}, nil
}

func (s *OrdersServiceServer) Get(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Validate input
	if in.GetId() < 1 {
		return nil, errors.New("order ID is wrong")
	}
	order := &pb.Order{
		Id: in.GetId(),
		Userid: getUserID(&ctx),
	}
	var toppingPrice float32
	var date time.Time
	if err := s.db.QueryRow(
		`SELECT coffee.name, coffee.price, topping.name, topping.price, sugar, status, date
		FROM order_ INNER JOIN
			 coffee ON order_.coffeeID == coffee.coffeeID INNER JOIN
			 topping ON order_.toppingID == topping.toppingID
		WHERE order_.orderID == $1 AND order_.userID == $2;`,
		order.Id, order.Userid).Scan(&order.Coffee, &order.Total, &order.Topping,
									 &toppingPrice, &order.Sugar, &order.Status, &date);
	err != nil {
		return nil, err
	}
	order.Date = timestamppb.New(date)
	order.Total += toppingPrice
	return &pb.GetOrderResponse{Order: order}, nil
}

// func (s *OrdersServiceServer) List(ctx context.Context, in *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
// 	// Validate input
// 	if in.GetShift() < 0 {
// 		return nil, errors.New("shift can be only positive")
// 	}
// 	rows, err := s.db.Query("SELECT orderID, userID, coffees, date, total, status FROM orders WHERE id > $1 AND userID == $2",
// 							in.GetShift(), getUserID(&ctx))
// 	if err != nil {
// 		log.Printf("DB request error: %v", err)
// 		return nil, errors.New("there is some problem with DB")
// 	}
// 	defer rows.Close()
// 	var (
// 		orders  []*pb.Order
// 		coffees string
// 		date    time.Time
// 	)
// 	for rows.Next() {
// 		order := pb.Order{}
// 		rows.Scan(&order.Id, &coffees, &date, &order.Total, &order.Status)
// 		order.Date = timestamppb.New(date)
// 		json.Unmarshal([]byte(coffees), &order.Coffees)
// 		orders = append(orders, &order)
// 	}
// 	return &pb.ListOrderResponse{Orders: orders}, nil
// }