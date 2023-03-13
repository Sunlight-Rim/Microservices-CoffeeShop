package main

import (
	"coffeeshop/config"
	"coffeeshop/internal/orders"
	"coffeeshop/internal/users"

	"fmt"
)

///	ENTRY POINT

func main() {
	fmt.Println(config.GetSocket())
	var (
		user_server  users.UsersServiceServer
		order_server orders.OrdersServiceServer
	)
	user_server.Create(nil, nil)
	order_server.Create(nil, nil)
}
