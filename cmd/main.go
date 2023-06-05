package main

import (
	gw "coffeeshop/internal/api-gw"
	"coffeeshop/internal/auth"
	"coffeeshop/internal/orders"
	"coffeeshop/internal/users"
)

///	ENTRY POINT

func main() {
	auth.Start()
	users.Start()
	orders.Start()
	gw.Start()
}