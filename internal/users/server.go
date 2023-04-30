package users

import (
	pb "coffeeshop/internal/users/pb"
	"errors"
	"fmt"
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
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	// Check if name is already exists
	var id int64
	s.db.QueryRow("select Id from USERS where Username == $1", in.GetUsername()).Scan(&id)
	if id != 0 {
		return nil, errors.New("the username is already taken")
	}
	// Fill a new row
	t := time.Now()
	res, err := s.db.Exec("insert into USERS (Username, Password, Address, RegDate, OrderIds, Token) values ($1, $2, $3, $4, '', '')",
		in.GetUsername(), in.GetPassword(), in.GetAddress(), t.Format(time.RFC3339))
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	id, _ = res.LastInsertId()
	return &pb.CreateUserResponse{User: &pb.User{
		Id:       id,
		Username: in.GetUsername(),
		Address:  in.GetAddress(),
		Regdate:  timestamppb.New(t),
		OrderIds: []int64{},
	}}, nil
}

func (s *UsersServiceServer) Login(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// Check for username and password is correct
	var id int64
	if s.db.QueryRow("select Id from USERS where Username == $1 and Password == $2",
		in.GetUsername(), in.GetPassword()).Scan(&id); id == 0 {
		return nil, errors.New("username or password is incorrect")
	}
	// Generate token
	t := time.Now()
	hash := md5.Sum([]byte(in.GetUsername() + in.GetPassword() + t.Format(time.RFC3339)))
	token := hex.EncodeToString(hash[:])
	if _, err := s.db.Exec("update USERS set Token = $1 where Id = $2", token, id); err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	return &pb.LoginUserResponse{Id: id, Token: token}, nil
}

func (s *UsersServiceServer) Get(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var (
		users  = []*pb.User{}
		t      time.Time
		orders string
	)
	// Find user account by token if Ids not requested
	if in.GetIds() == nil {
		user := pb.User{}
		if s.db.QueryRow("select Id, Username, Address, RegDate, OrderIds from USERS where Token == $1",
			in.GetToken()).Scan(&user.Id, &user.Username, &user.Address, &t, &orders); user.Id == 0 {
			return nil, errors.New("token is incorrect")
		}
		user.Regdate = timestamppb.New(t)
		json.Unmarshal([]byte(orders), &user.OrderIds)
		users = append(users, &user)
	} else {
		// Check token
		var id int64
		if s.db.QueryRow("select Id from USERS where Token == $1", in.GetToken()).Scan(&id); id == 0 {
			return nil, errors.New("token is incorrect")
		}
		// Find users by requested Ids
		rows, err := s.db.Query("select Id, Username, Address, RegDate, OrderIds from USERS where Id in (" +
			strings.Trim(strings.Join(strings.Fields(fmt.Sprint(in.GetIds())), ","), "[]") + ")")
		if err != nil {
			log.Printf("DB request error: %v", err)
			return nil, errors.New("there is some problem with DB")
		}
		defer rows.Close()
		for i := 0; rows.Next(); i++ {
			user := pb.User{}
			rows.Scan(&user.Id, &user.Username, &user.Address, &t, &orders)
			json.Unmarshal([]byte(orders), &user.OrderIds)
			user.Regdate = timestamppb.New(t)
			users = append(users, &user)
		}
	}
	return &pb.GetUserResponse{Users: users}, nil
}
