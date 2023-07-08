package database

import (
	"coffeeshop/internal/users/domain"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/// REPOSITORY LAYER (sqlite adapter)

type Repo struct {
	*sql.DB
}

// Connect to DB
func Connect(dbPath string) (Repo, error) {
	db, err := sql.Open("sqlite3", dbPath)
	return Repo{db}, err
}

func (repo *Repo) DoesUsernameExists(username string) (existence bool) {
	var id int32
	repo.QueryRow(
		`SELECT userID FROM user WHERE username = $1`,
		username).Scan(&id)
	existence = id != 0
	return
}

func (repo *Repo) CreateUser(username, passwordHash, address string, date time.Time) (uint32, error) {
	res, err := repo.Exec(
		`INSERT INTO user (username, passwordHash, address, date) VALUES ($1, $2, $3, $4)`,
		username, passwordHash, address, date)
	id, _ := res.LastInsertId()
	return uint32(id), err
}

func (repo *Repo) GetPasswordHashByName(username string) (passwordHash string, err error) {
	err = repo.QueryRow(
		`SELECT passwordHash FROM user WHERE username = $1`,
		username).Scan(&passwordHash)
	return
}

func (repo *Repo) GetUserById(userID uint32) (user *domain.User, err error) {
	err = repo.QueryRow(
		`SELECT username, address, date FROM user WHERE userID = $1`,
		user.Id).Scan(&user.Username, &user.Address, &user.Regdate)
	return
}