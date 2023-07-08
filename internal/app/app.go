package app

import (
	"coffeeshop/config"
	gw "coffeeshop/internal/api-gw"
	auth "coffeeshop/internal/auth/grpc"
	orders "coffeeshop/internal/orders/grpc"
	users "coffeeshop/internal/users/grpc"
)

/// APP COMMON LOGIC

func Start(configPath string) {
	cfg := config.New(configPath)

	auth.Start(cfg.Host, cfg.Services["auth"].Port, cfg.Services["auth"].DB, cfg.Services["users"].Port)
	users.Start(cfg.Host, cfg.Services["users"].Port, cfg.Services["users"].DB)
	orders.Start(cfg.Host, cfg.Services["orders"].Port, cfg.Services["orders"].DB)
	gw.Start(cfg.Host, cfg.Port, cfg.JWTKey,
		cfg.Services["auth"].URL, cfg.Services["auth"].Port,
		cfg.Services["users"].URL, cfg.Services["users"].Port,
		cfg.Services["orders"].URL, cfg.Services["orders"].Port)
	// TODO: graceful shutdown
}
