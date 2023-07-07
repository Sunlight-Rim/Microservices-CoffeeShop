package orders

import "google.golang.org/protobuf/types/known/timestamppb"

/// DOMAIN LAYER

type Order struct {
	Id      uint32
	UserId  uint32
	Status  int32
	Coffee  string
	Topping string
	Sugar   uint32
	Total   float64
	Date    *timestamppb.Timestamp
}
