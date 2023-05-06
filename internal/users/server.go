package users

import (
	pb "coffeeshop/internal/users/pb"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"context"
	"log"
	"net"

	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const ( // TODO: move to config
	grpcPort = "50052"
)

/// SERVER DEFINITION

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	db *sql.DB
}

func Start() {
	// Connect to DB
	db, err := sql.Open("sqlite3", "internal/users/users.db")
	if err != nil {
		log.Fatalf("%v", err)
	}
	userService := UsersServiceServer{db: db}
	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, &userService)
	lis, err := net.Listen("tcp", "localhost:"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Users server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
		defer userService.db.Close()
	}()
}

/// API METHODS (gRPC)

func (s *UsersServiceServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Validate input
	if in.GetUsername() == "" {
		return nil, errors.New("enter correct username")
	}
	if in.GetPassword() == "" || len(in.GetPassword()) < 6 {
		return nil, errors.New("password must exceed 6 chars")
	}
	// Check if name is already exist
	var id int64
	if s.db.QueryRow("SELECT id FROM users WHERE username == $1", in.GetUsername()).Scan(&id); id != 0 {
		return nil, errors.New("this username is already taken")
	}
	// Fill a new row
	regdate := time.Now()
	user := pb.User{
		Username: in.GetUsername(),
		Address:  in.GetAddress(),
		Regdate:  timestamppb.New(regdate),
		OrderIds: nil,
	}
	res, err := s.db.Exec("INSERT INTO users (username, password, address, reg_date) VALUES ($1, $2, $3, $4)",
						  &user.Username, in.GetPassword(), &user.Address, regdate)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	user.Id, _ = res.LastInsertId()
	return &pb.CreateUserResponse{User: &user}, nil
}

func (s *UsersServiceServer) Login(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// Check username and password
	var id int64
	if s.db.QueryRow("SELECT id FROM users WHERE username == $1 AND password == $2",
					 in.GetUsername(), in.GetPassword()).Scan(&id); id == 0 {
		return nil, errors.New("username or password is incorrect")
	}
	// Generate token
	var token string
	for i := 0; i < 10; i++ {
		hash := md5.Sum([]byte(in.GetUsername() + in.GetPassword() + time.Now().Format(time.RFC3339)))
		token = hex.EncodeToString(hash[:])
		var checkId int64
		if s.db.QueryRow("SELECT id FROM users WHERE token == $1", token).Scan(&checkId); checkId == 0 { break }
	}
	if _, err := s.db.Exec("UPDATE users SET token = $1 WHERE id == $2", token, id); err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	return &pb.LoginUserResponse{Id: id, Token: token}, nil
}

func (s *UsersServiceServer) Get(ctx context.Context, in *pb.GetUserRequest) (*pb.ListUserResponse, error) {
	var ( user    pb.User
		  orders  string
		  regdate time.Time )
	// Check token & get data
	if s.db.QueryRow("SELECT id, username, address, reg_date, order_ids FROM users WHERE token == $1",
					 in.GetToken()).Scan(&user.Id, &user.Username, &user.Address, &regdate, &orders);
					 user.Id == 0 {
		return nil, errors.New("token is incorrect")
	}
	if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &user.OrderIds) }
	user.Regdate = timestamppb.New(regdate)
	return &pb.ListUserResponse{Users: []*pb.User{ &user }}, nil
}

func (s *UsersServiceServer) List(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	// Check token
	var id int64
	if s.db.QueryRow("SELECT id FROM users WHERE token == $1", in.GetToken()).Scan(&id); id == 0 {
		return nil, errors.New("token is incorrect")
	}
	// Find users by requested Ids
	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(in.GetIds())), ","), "[]")
	rows, err := s.db.Query("SELECT id, username, address, reg_date, order_ids FROM users WHERE id IN ("+ids+")")
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	defer rows.Close()
	var ( users   []*pb.User
		  orders  string
		  regdate time.Time )
	// Fill response users data
	for rows.Next() {
		user := pb.User{}
		rows.Scan(&user.Id, &user.Username, &user.Address, &regdate, &orders)
		if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &user.OrderIds) }
		user.Regdate = timestamppb.New(regdate)
		users = append(users, &user)
	}
	return &pb.ListUserResponse{Users: users}, nil
}

func (s *UsersServiceServer) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var ( user    pb.User
		  orders  string
		  regdate time.Time )
	// Check token & get data
	if s.db.QueryRow("SELECT id, username, address, reg_date, order_ids FROM users " +
					 "WHERE token == $1", in.GetToken()).Scan(&user.Id, &user.Username,
					  &user.Address, &regdate, &orders); user.Id == 0 {
		return nil, errors.New("token is incorrect")
	}
	// Check if name is already exist
	if username := in.GetUser().Username; username != "" {
		var id int64
		if s.db.QueryRow("SELECT id FROM users WHERE username == $1", username).Scan(&id); id != 0 {
			return nil, errors.New("this username is already taken")
		}
		user.Username = username
	}
	if address := in.GetUser().Address; address != "" {
		user.Address = address
	}
	// Update info
	if _, err := s.db.Exec("UPDATE users SET username = $1, address = $2 WHERE token == $3",
						   &user.Username, &user.Address, in.GetToken()); err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &user.OrderIds) }
	user.Regdate = timestamppb.New(regdate)
	return &pb.UpdateUserResponse{User: &user}, nil
}

func (s *UsersServiceServer) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	var ( user    pb.User
		  orders  string
		  regdate time.Time )
	// Check token & get data
	if s.db.QueryRow("SELECT id, username, address, reg_date, order_ids FROM users " +
					 "WHERE token == $1", in.GetToken()).Scan(&user.Id, &user.Username,
					 &user.Address, &regdate, &orders); user.Id == 0 {
		return nil, errors.New("token is incorrect")
	}
	// Delete account
	if _, err := s.db.Exec("DELETE FROM users WHERE id == $1", user.Id); err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &user.OrderIds) }
	user.Regdate = timestamppb.New(regdate)
	return &pb.DeleteUserResponse{User: &user}, nil
}

func (s *UsersServiceServer) AuthUser(ctx context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	var id int64
	if s.db.QueryRow("SELECT id FROM users WHERE token == $1", in.GetToken()).Scan(&id); id == 0 {
		return nil, errors.New("token is incorrect")
	}
	return &pb.AuthUserResponse{Id: id}, nil
}

func (s *UsersServiceServer) CreateUserOrder(ctx context.Context, in *pb.CreateUserOrderRequest) (*emptypb.Empty, error) {
	if _, err := s.db.Exec("UPDATE users SET order_ids = order_ids || $1 WHERE id == $2",
						   strconv.FormatInt(in.GetOrderId(), 10)+",", in.GetId()); err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	return &emptypb.Empty{}, nil
}

func (s *UsersServiceServer) GetUserOrders(ctx context.Context, in *pb.GetUserOrdersRequest) (*pb.GetUserOrdersResponse, error) {
	var ( orders    string
		  ordersIds []int64 )
	s.db.QueryRow("SELECT order_ids FROM users WHERE id == $1", in.GetId()).Scan(&orders)
	if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &ordersIds) }
	return &pb.GetUserOrdersResponse{OrderIds: ordersIds}, nil
}
