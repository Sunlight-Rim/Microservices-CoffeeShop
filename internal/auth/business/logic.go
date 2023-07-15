package business

import (
	"coffeeshop/internal/auth/domain"
	"time"
)

/// BUSINESS LOGIC LAYER

type Repository interface{
	// DB methods
}

type Logic struct {
	repo Repository
}

func New(repo Repository) *Logic {
	return &Logic{repo: repo}
}

// Register new user
func (l *Logic) Signup(username, password, address string,
					   createUser func(username, password, address string) (uint32, *time.Time, error),
					   ) (*domain.User, error) {
	// createUser uses the Users service to create a new user in Users service DB
	userID, userRegdate, err := createUser(username, password, address)
	return &domain.User{
		Id:       userID,
		Username: username,
		Address:  address,
		Regdate:  userRegdate,
	}, err
}
