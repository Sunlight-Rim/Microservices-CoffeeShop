package main

import (
	gw "coffeeshop/internal/api-gw"
	"coffeeshop/internal/orders"
)

///	ENTRY POINT

func main() {
	orders.Start()

	gwServer := gw.New()
	gwServer.Start()
}
