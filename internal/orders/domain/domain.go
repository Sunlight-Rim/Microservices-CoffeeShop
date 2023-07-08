package domain

import (
	"time"
)

/// DOMAIN LAYER

type Order struct {
	Id      uint32
	Userid  uint32
	Status  int32
	Coffee  string
	Topping string
	Sugar   uint32
	Total   float32
	Date    time.Time
}
