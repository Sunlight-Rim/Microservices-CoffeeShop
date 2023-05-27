package main

import (
	gw "coffeeshop/internal/api-gw"
	"coffeeshop/internal/orders"
	"coffeeshop/internal/users"
)

///	ENTRY POINT

func main() {
	orders.Start()
	users.Start()
	gw.Start()
}
