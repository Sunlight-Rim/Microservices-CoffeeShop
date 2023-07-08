package business

import (
	db "coffeeshop/internal/auth/database"
	"coffeeshop/internal/auth/domain"
	"time"
)

/// BUSINESS LOGIC LAYER

type Business struct {
	repo *db.Repo
}

func New(repo *db.Repo) Business {
	return Business{repo: repo}
}

func (b *Business) Signup(username, password, address string,
						  createUser func(username, password, address string) (uint32, time.Time, error),
						  ) (*domain.User, error) {
	userID, userRegdate, err := createUser(username, password, address)
	return &domain.User{
		Id:       userID,
		Username: username,
		Address:  address,
		Regdate:  &userRegdate,
	}, err
}
