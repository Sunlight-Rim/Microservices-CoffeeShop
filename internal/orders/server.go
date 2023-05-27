package orders

import (
	pb "coffeeshop/internal/orders/pb"
	usersPB "coffeeshop/internal/users/pb"
	"database/sql"
	"encoding/json"
	"errors"
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

/// gRPC SERVER

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
	defer db.Close()
	// Connect to Users service
	usersConn, err := grpc.Dial("localhost:"+usersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to gRPC: %v", err)
	}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	orderService := OrdersServiceServer{db: db, users: usersPB.NewUsersServiceClient(usersConn)}
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
		return nil, errors.New("you didn't specify any coffee")
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
	// Append to DB
	res, err := s.db.Exec("INSERT INTO orders (userid, coffees, total, date, status) VALUES ($1, $2, $3, $4, $5)",
						  in.GetUserid(), string(coffees), &order.Total, date.Format(time.RFC3339), 0)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	order.Id, _ = res.LastInsertId()
	// Update order count in Users DB
	if _, err := s.users.IncUserOrder(ctx, &usersPB.IncUserOrderRequest{Id: in.GetUserid()}); err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{Order: &order}, nil
}

func (s *OrdersServiceServer) Get(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	log.Printf("yow! %v", in.GetId())
	var (
		order   pb.Order
		coffees string
		date    time.Time
	)
	if err := s.db.QueryRow("SELECT id, userid, coffees, date, total, status FROM orders WHERE id == $1 AND userid == $2",
							in.GetId(), in.GetUserid()).Scan(&order.Id, &order.Userid, &coffees, &date, &order.Total, &order.Status);
	err != nil {
		return nil, err
	}
	if order.Id == 0 {
		return nil, errors.New("order not found")
	}
	order.Date = timestamppb.New(date)
	json.Unmarshal([]byte(coffees), &order.Coffees)
	return &pb.GetOrderResponse{Order: &order}, nil
}

func (s *OrdersServiceServer) List(ctx context.Context, in *pb.ListOrderRequest) (*pb.ListOrderResponse, error) {
	rows, err := s.db.Query("SELECT id, userid, coffees, date, total, status FROM orders WHERE id > $1 AND userid == $2",
							in.GetShift(), in.GetUserid())
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	defer rows.Close()
	var (
		orders  []*pb.Order
		coffees string
		date    time.Time
	)
	for rows.Next() {
		order := pb.Order{}
		rows.Scan(&order.Id, &coffees, &date, &order.Total, &order.Status)
		order.Date = timestamppb.New(date)
		json.Unmarshal([]byte(coffees), &order.Coffees)
		orders = append(orders, &order)
	}
	return &pb.ListOrderResponse{Orders: orders}, nil
}
