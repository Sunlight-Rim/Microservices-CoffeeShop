package business

import (
	db "coffeeshop/internal/auth/database"
	"coffeeshop/internal/auth/domain"
	"time"
)

/// BUSINESS LOGIC LAYER

type Logic struct {
	repo *db.Repo
}

func New(repo *db.Repo) Logic {
	return Logic{repo: repo}
}

func (l *Logic) Signup(username, password, address string,
						  createUser func(username, password, address string) (uint32, *time.Time, error),
						  ) (*domain.User, error) {
	userID, userRegdate, err := createUser(username, password, address)
	return &domain.User{
		Id:       userID,
		Username: username,
		Address:  address,
		Regdate:  userRegdate,
	}, err
}
