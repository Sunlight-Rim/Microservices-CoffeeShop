package users

import (
	pb "coffeeshop/internal/users/pb"
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"

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

func (s *UsersServiceServer) GetUser(ctx context.Context, empty *empty.Empty) (*pb.GetUserResponse, error) {
	user := &pb.User{
		Id: getUserID(&ctx),
	}
	var date time.Time
	s.db.QueryRow("SELECT username, address, date FROM user WHERE userID == $1",
					 user.Id).Scan(&user.Username, &user.Address, &date)
	user.Regdate = timestamppb.New(date)
	return &pb.GetUserResponse{User: user}, nil
}

// func (s *UsersServiceServer) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
// 	// Validate input
// 	if in.GetUsername() == "" {
// 		return nil, errors.New("enter correct username")
// 	}
// 	if in.GetPassword() == "" || len(in.GetPassword()) < 6 {
// 		return nil, errors.New("password must exceed 6 chars")
// 	}
// 	// Check if name is already exist
// 	var id int64
// 	if s.db.QueryRow("SELECT id FROM users WHERE username == $1", in.GetUsername()).Scan(&id); id != 0 {
// 		return nil, errors.New("this username is already taken")
// 	}
// 	// Fill a new row
// 	regdate := time.Now()
// 	user := pb.User{
// 		Username: in.GetUsername(),
// 		Address:  in.GetAddress(),
// 		Regdate:  timestamppb.New(regdate),
// 		OrderIds: nil,
// 	}
// 	res, err := s.db.Exec("INSERT INTO users (username, password, address, reg_date, order_ids) VALUES " +
// 						  "($1, $2, $3, $4, '')", &user.Username, in.GetPassword(), &user.Address, regdate)
// 	if err != nil {
// 		log.Printf("DB request error: %v", err)
// 		return nil, errors.New("there is some problem with DB")
// 	}
// 	user.Id, _ = res.LastInsertId()
// 	return &pb.CreateUserResponse{User: &user}, nil
// }

// func (s *UsersServiceServer) List(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
// 	// Check token
// 	var id int64
// 	if s.db.QueryRow("SELECT id FROM users WHERE token == $1", in.GetToken()).Scan(&id); id == 0 {
// 		return nil, errors.New("token is incorrect")
// 	}
// 	// Find users by requested Ids
// 	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(in.GetIds())), ","), "[]")
// 	rows, err := s.db.Query("SELECT id, username, address, reg_date, order_ids FROM users WHERE id IN ("+ids+")")
// 	if err != nil {
// 		log.Printf("DB request error: %v", err)
// 		return nil, errors.New("there is some problem with DB")
// 	}
// 	defer rows.Close()
// 	var ( users   []*pb.User
// 		  orders  string
// 		  regdate time.Time )
// 	// Fill response users data
// 	for rows.Next() {
// 		user := pb.User{}
// 		rows.Scan(&user.Id, &user.Username, &user.Address, &regdate, &orders)
// 		if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &user.OrderIds) }
// 		user.Regdate = timestamppb.New(regdate)
// 		users = append(users, &user)
// 	}
// 	return &pb.ListUserResponse{Users: users}, nil
// }

// func (s *UsersServiceServer) Update(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
// 	var ( user    pb.User
// 		  orders  string
// 		  regdate time.Time )
// 	// Check token & get data
// 	if s.db.QueryRow("SELECT id, username, address, reg_date, order_ids FROM users " +
// 					 "WHERE token == $1", in.GetToken()).Scan(&user.Id, &user.Username,
// 					  &user.Address, &regdate, &orders); user.Id == 0 {
// 		return nil, errors.New("token is incorrect")
// 	}
// 	// Check if name is already exist
// 	if username := in.GetUser().Username; username != "" {
// 		var id int64
// 		if s.db.QueryRow("SELECT id FROM users WHERE username == $1", username).Scan(&id); id != 0 {
// 			return nil, errors.New("this username is already taken")
// 		}
// 		user.Username = username
// 	}
// 	if address := in.GetUser().Address; address != "" {
// 		user.Address = address
// 	}
// 	// Update info
// 	if _, err := s.db.Exec("UPDATE users SET username = $1, address = $2 WHERE token == $3",
// 						   &user.Username, &user.Address, in.GetToken()); err != nil {
// 		log.Printf("DB request error: %v", err)
// 		return nil, errors.New("there is some problem with DB")
// 	}
// 	if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &user.OrderIds) }
// 	user.Regdate = timestamppb.New(regdate)
// 	return &pb.UpdateUserResponse{User: &user}, nil
// }

// func (s *UsersServiceServer) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
// 	var ( user    pb.User
// 		  orders  string
// 		  regdate time.Time )
// 	// Check token & get data
// 	if s.db.QueryRow("SELECT id, username, address, reg_date, order_ids FROM users " +
// 					 "WHERE token == $1", in.GetToken()).Scan(&user.Id, &user.Username,
// 					 &user.Address, &regdate, &orders); user.Id == 0 {
// 		return nil, errors.New("token is incorrect")
// 	}
// 	// Delete account
// 	if _, err := s.db.Exec("DELETE FROM users WHERE id == $1", user.Id); err != nil {
// 		log.Printf("DB request error: %v", err)
// 		return nil, errors.New("there is some problem with DB")
// 	}
// 	if orders != "" { json.Unmarshal([]byte("["+orders[:len(orders)-1]+"]"), &user.OrderIds) }
// 	user.Regdate = timestamppb.New(regdate)
// 	return &pb.DeleteUserResponse{User: &user}, nil
// }
