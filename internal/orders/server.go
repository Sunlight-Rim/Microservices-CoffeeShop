package orders

import (
	pb "coffeeshop/internal/orders/pb"
	usersPB "coffeeshop/internal/users/pb"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
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
	// Check token
	authUser, err := s.users.AuthUser(ctx, &usersPB.AuthUserRequest{Token: in.GetToken()})
	if err != nil {
		return nil, err
	}
	// Create Order
	var (
		date    time.Time = time.Now()
		total   float32   = 0
		coffees string
	)
	for _, c := range in.GetCoffees() {
		total += coffeePrices[c.Type]
		coffees += "{" + c.GetType() + ", " + strconv.Itoa(int(c.GetSugar())) + "}"
	}
	res, err := s.db.Exec("insert into ORDERS (Coffees, Total, Date, Status) values ($1, $2, $3, $4)",
		coffees, total, date.Format(time.RFC3339), 0)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	id, _ := res.LastInsertId()
	if _, err := s.users.CreateUserOrder(ctx, &usersPB.CreateUserOrderRequest{
		Id: authUser.Id, OrderId: id}); err != nil {
		return nil, err
	}
	// Adding recieved data to DataBase
	return &pb.CreateOrderResponse{Order: &pb.Order{
		Id:      id,
		Coffees: in.GetCoffees(),
		Date:    timestamppb.New(date),
		Total:   total,
		Status:  pb.Order_Status(0),
	}}, nil
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
	var (
		orders  []*pb.Order
		coffees string
		date    time.Time
		status  int8
	)
	// Get all orders from user
	if in.GetId() == 0 {
		rows, err := s.db.Query("select Id, Coffees, Date, Total, Status from ORDERS where Id in (" +
			strings.Trim(strings.Join(strings.Fields(fmt.Sprint(userOrders.OrderIds)), ","), "[]") + ")")
		if err != nil {
			log.Printf("DB request error: %v", err)
			return nil, errors.New("there is some problem with DB")
		}
		defer rows.Close()
		for rows.Next() {
			order := pb.Order{}
			rows.Scan(&order.Id, &coffees, &date, &order.Total, &status)
			order.Date = timestamppb.New(date)
			order.Status = pb.Order_Status(status)
			for _, c := range strings.Split(coffees[1:len(coffees)-1], "}{") {
				typeSugar := strings.Split(c, ", ")
				sugar, _ := strconv.Atoi(typeSugar[1])
				order.Coffees = append(order.Coffees, &pb.Coffee{Type: typeSugar[0], Sugar: int32(sugar)})
			}
			orders = append(orders, &order)
		}
		// Get specifed order from user
	} else {

	}
	// Recieve data from DataBase
	return &pb.GetOrderResponse{Orders: orders}, nil
}
