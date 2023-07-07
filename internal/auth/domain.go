package auth

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/// DOMAIN LAYER

type User struct {
	Id       uint32
	Username string
	Address  string
	RegDate  *timestamppb.Timestamp
}

type UserGetter interface {
	GetId() uint32
	GetUsername() string
	GetAddress() string
	GetRegdate() *timestamp.Timestamp
}
