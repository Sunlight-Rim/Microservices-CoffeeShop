package business

import (
	"coffeeshop/internal/users/domain"

	"crypto/sha1"
	"encoding/hex"
	"errors"
	"log"
	"time"
	"unicode/utf8"
)

/// APPLICATION BUSINESS LOGIC LAYER

type Repository interface {
	DoesUsernameExists(string) (bool)
	CreateUser(username string, passwordHash string, address string, date time.Time) (uint32, error)
	GetPasswordHashByName(username string) (passwordHash string, err error)
	GetUserById(userID uint32) (user *domain.User, err error)
}

type Logic struct {
	repo Repository
}

func New(repo Repository) *Logic {
	return &Logic{repo: repo}
}

// Create user
func (l *Logic) Create(username, password, address string) (*domain.User, error) {
	// Validation
	if username == "" {
		return nil, errors.New("username is empty")
	}
	if address == "" {
		return nil, errors.New("address is empty")
	}
	if utf8.RuneCountInString(password) < 6 {
		return nil, errors.New("password must exceed 6 chars")
	}
	// Check if name is already exist
	existence := l.repo.DoesUsernameExists(username)
	if existence {
		return nil, errors.New("this username is already exists")
	}
	// Hash the password
	hasher := sha1.New()
	hasher.Write([]byte(password))
	// Append to DB
	date := time.Now()
	userID, err := l.repo.CreateUser(username, hex.EncodeToString(hasher.Sum(nil)), address, date)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	return &domain.User{
		Id: userID,
		Username: username,
		Address:  address,
		Regdate:  date,
	}, nil
}

// Check if username and password is valid
func (l *Logic) Login(username, password string) (error) {
	// Validation
	if username == "" {
		return errors.New("username is empty")
	}
	// Password hash by username
	passwordHash, err := l.repo.GetPasswordHashByName(username)
	if err != nil {
		return errors.New("username is wrong")
	}
	// Hashes verification
	hasher := sha1.New()
	hasher.Write([]byte(password))
	if passwordHash != hex.EncodeToString(hasher.Sum(nil)) {
		return errors.New("password is wrong")
	}
	return nil
}

// Get info about user himself
func (l *Logic) GetMe(userID uint32) (*domain.User, error) {
	user, err := l.repo.GetUserById(userID)
	if err != nil {
		return nil, errors.New("there is some problem with DB")
	}
	return user, nil
}

// func (s *UsersServiceServer) GetOther(context.Context, *pb.GetOtherUserRequest) (*pb.GetOtherUserResponse, error) {
// 	return &pb.GetOtherUserResponse{}, nil
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
