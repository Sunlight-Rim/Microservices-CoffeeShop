package orders

import (
	pb "coffeeshop/internal/orders/pb"
	usersPB "coffeeshop/internal/users/pb"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"context"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const ( // TODO: move to config
	grpcPort  = "50051"
	usersPort = "50052"
)

/// SERVER DEFINITION

type OrdersServiceServer struct {
	pb.UnimplementedOrdersServiceServer
	db    *sql.DB
	users usersPB.UsersServiceClient
}

func Start() {
	// Connect to DB
	db, err := sql.Open("sqlite3", "internal/orders/orders.db")
	if err != nil {
		log.Fatalf("%v", err)
	}
	usersConn, err := grpc.Dial("localhost:"+usersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to gRPC: %v", err)
	}
	orderService := OrdersServiceServer{db: db, users: usersPB.NewUsersServiceClient(usersConn)}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterOrdersServiceServer(grpcServer, &orderService)
	lis, err := net.Listen("tcp", "localhost:"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Orders server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
		defer orderService.db.Close()
	}()
}

var coffeePrices = map[string]float32{
	"Espresso":  2,
	"Americano": 2.5,
}

/// API METHODS (gRPC)

func (s *OrdersServiceServer) Create(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Validate input
	if len(in.GetCoffees()) == 0 {
		return nil, errors.New("you didn't order any coffee")
	}
	// Check token
	authUser, err := s.users.AuthUser(ctx, &usersPB.AuthUserRequest{Token: in.GetToken()})
	if err != nil {
		return nil, err
	}
	// Create order
	date := time.Now()
	order := pb.Order{
		Coffees: in.GetCoffees(),
		Date:    timestamppb.New(date),
		Status:  pb.Order_Status(0),
	}
	for _, c := range in.GetCoffees() {
		price, ok := coffeePrices[c.Type]
		if !ok {
			return nil, errors.New("unknown coffee type")
		}
		order.Total += price
	}
	coffees, _ := json.Marshal(in.GetCoffees())
	res, err := s.db.Exec("INSERT INTO orders (coffees, total, date, status) VALUES ($1, $2, $3, $4)",
						  string(coffees), &order.Total, date.Format(time.RFC3339), 0)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	order.Id, _ = res.LastInsertId()
	// Add order to users table in DB
	if _, err := s.users.CreateUserOrder(ctx, &usersPB.CreateUserOrderRequest{
		Id: authUser.Id, OrderId: order.Id}); err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{Order: &order}, nil
}

func (s *OrdersServiceServer) Get(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// Check token
	authUser, err := s.users.AuthUser(ctx, &usersPB.AuthUserRequest{Token: in.GetToken()})
	if err != nil {
		return nil, err
	}
	userOrders, err := s.users.GetUserOrders(ctx, &usersPB.GetUserOrdersRequest{Id: authUser.Id})
	if err != nil {
		return nil, err
	}
	// Fill response order data
	var ( order   pb.Order
		  coffees string
		  date    time.Time )
	for _, id := range userOrders.OrderIds {
		if id == in.GetId() {
			s.db.QueryRow("SELECT id, coffees, date, total, status FROM orders WHERE id == $1", id).Scan(
						  &order.Id, &coffees, &date, &order.Total, &order.Status)
			break
		}
	}
	order.Date = timestamppb.New(date)
	json.Unmarshal([]byte(coffees), &order.Coffees)
	return &pb.GetOrderResponse{Order: &order}, nil
}

func (s *OrdersServiceServer) List(ctx context.Context, in *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
	// Check token
	authUser, err := s.users.AuthUser(ctx, &usersPB.AuthUserRequest{Token: in.GetToken()})
	if err != nil {
		return nil, err
	}
	userOrders, err := s.users.GetUserOrders(ctx, &usersPB.GetUserOrdersRequest{Id: authUser.Id})
	if err != nil {
		return nil, err
	}
	// Get all orders from user
	var ( orders  []*pb.Order
		  coffees string
		  date    time.Time )
	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(userOrders.OrderIds)), ","), "[]")
	rows, err := s.db.Query("SELECT id, coffees, date, total, status FROM orders WHERE id IN ("+ids+")")
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	defer rows.Close()
	for rows.Next() {
		order := pb.Order{}
		rows.Scan(&order.Id, &coffees, &date, &order.Total, &order.Status)
		order.Date = timestamppb.New(date)
		json.Unmarshal([]byte(coffees), &order.Coffees)
		orders = append(orders, &order)
	}
	// Recieve data from DataBase
	return &pb.ListOrderResponse{Orders: orders}, nil
}
