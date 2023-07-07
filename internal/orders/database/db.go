package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/// REPOSITORY LAYER

// Connect to DB
func Connect(dbPath string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}
