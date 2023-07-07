package main

import "coffeeshop/internal/app"

///	ENTRY POINT

const configPath = "config/config.yaml"

func main() {
	app.Start(configPath)
}
