package users

import (
	pb "coffeeshop/internal/users/pb"
	"errors"
	"time"

	"context"
	"log"
	"net"

	"crypto/md5"
	"database/sql"
	"encoding/hex"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	grpcPort = "50052"
)

/// SERVER DEFINITION

type UsersServiceServer struct {
	pb.UnimplementedUsersServiceServer
	db *sql.DB
}

func Start() {
	userService := UsersServiceServer{}
	// Connect to DB
	var err error
	if userService.db, err = sql.Open("sqlite3", "internal/users/users.db"); err != nil {
		log.Fatalf("%v", err)
	}

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
	var checkName int
	s.db.QueryRow("select Id from USERS where Username == $1", in.GetUsername()).Scan(&checkName)
	if checkName != 0 {
		return nil, errors.New("the username is already taken")
	}
	// Fill the new row
	t := time.Now()
	res, err := s.db.Exec("insert into USERS (Username, Password, Address, RegDate) values ($1, $2, $3, $4)",
		in.GetUsername(), in.GetPassword(), in.GetAddress(), t.Format(time.RFC3339))
	if err != nil {
		log.Printf("Cannot request to DB: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	id, _ := res.LastInsertId()
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
	s.db.QueryRow("select id from USERS where Username == $1 and Password == $2",
		in.GetUsername(), in.GetPassword()).Scan(&id)
	if id == 0 {
		return nil, errors.New("username or password is incorrect")
	}
	// Generate token
	t := time.Now()
	hash := md5.Sum([]byte(in.GetUsername() + in.GetPassword() + t.Format(time.RFC3339)))
	token := hex.EncodeToString(hash[:])
	_, err := s.db.Exec("update USERS set Token = $1 where Id = $2", token, id)
	if err != nil {
		log.Printf("Cannot request to DB: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	return &pb.LoginUserResponse{Id: id, Token: token}, nil
}
