package domain

import (
	"time"
)

/// DOMAIN LAYER

type User struct {
	Id       uint32
	Username string
	Address  string
	Regdate  time.Time
}
