package main

import (
	configuration "coffeeshop/config"
	gw "coffeeshop/internal/api-gw"
	"coffeeshop/internal/auth"
	"coffeeshop/internal/orders"
	"coffeeshop/internal/users"
)

///	ENTRY POINT

func main() {
	config := configuration.New()

	auth.Start(config)
	users.Start(config)
	orders.Start(config)
	gw.Start(config)
}
