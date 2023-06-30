package gw

import (
	configuration "coffeeshop/config"
	pbAuth "coffeeshop/internal/auth/pb"
	pbOrders "coffeeshop/internal/orders/pb"
	pbUsers "coffeeshop/internal/users/pb"
	"context"
	"crypto/hmac"
	"crypto/sha256"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/// REST SERVER

func Start(config *configuration.Config) {
	log.Print("API gateway (REST->gRPC) server listening at http://"+config.Host+":"+config.Port)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Dial with started gRPC servers
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authMux := runtime.NewServeMux()
	if err := pbAuth.RegisterAuthServiceHandlerFromEndpoint(ctx, authMux, config.Host+":"+config.Services["auth"].Port, opts); err != nil {
		log.Fatalf("Failed to dial with Auth endpoint: %v", err)
	}
	usersMux := runtime.NewServeMux()
	if err := pbUsers.RegisterUsersServiceHandlerFromEndpoint(ctx, usersMux, config.Host+":"+config.Services["users"].Port, opts); err != nil {
		log.Fatalf("Failed to dial with Users endpoint: %v", err)
	}
	ordersMux := runtime.NewServeMux()
	if err := pbOrders.RegisterOrdersServiceHandlerFromEndpoint(ctx, ordersMux, config.Host+":"+config.Services["orders"].Port, opts); err != nil {
		log.Fatalf("Failed to dial with Orders endpoint: %v", err)
	}
	// Handlers
	router := gin.Default()
	registerMux := func(url string, handlers ...gin.HandlerFunc) {
		routerGroup := router.Group(url, handlers...)
		routerGroup.Any("")
		routerGroup.Any("*any")
	}
	tokenHash := hmac.New(sha256.New, []byte(config.JWTKey))
	registerMux(config.Services["auth"].URL, gin.WrapF(authMux.ServeHTTP))
	registerMux(config.Services["users"].URL, Auth(tokenHash), gin.WrapF(usersMux.ServeHTTP))
	registerMux(config.Services["orders"].URL, Auth(tokenHash), gin.WrapF(ordersMux.ServeHTTP))
	// Start server
	if err := router.Run(":"+config.Port); err != nil {
		log.Fatal(err)
	}
}