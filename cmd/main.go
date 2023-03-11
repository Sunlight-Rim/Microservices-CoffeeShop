package main

import (
	"coffeeshop/config"
	"fmt"
)

///	ENTRY POINT

func main() {
	fmt.Print(config.GetSocket())
}
