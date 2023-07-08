package gw

import (
	pbAuth "coffeeshop/internal/auth/grpc/pb"
	pbOrders "coffeeshop/internal/orders/grpc/pb"
	pbUsers "coffeeshop/internal/users/grpc/pb"
	"context"
	"crypto/hmac"
	"crypto/sha256"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/// TRANSPORT LAYER (REST)

func Start(host, port, jwtKey, authUrl, authPort, usersUrl, userPort, ordersUrl, ordersPort string) {
	log.Print("API gateway (REST->gRPC) server listening at http://" + host + ":" + port)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Dial with started gRPC servers
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authMux := runtime.NewServeMux()
	if err := pbAuth.RegisterAuthServiceHandlerFromEndpoint(ctx, authMux, host+":"+authPort, opts); err != nil {
		log.Fatalf("Failed to dial with Auth endpoint: %v", err)
	}
	usersMux := runtime.NewServeMux()
	if err := pbUsers.RegisterUsersServiceHandlerFromEndpoint(ctx, usersMux, host+":"+userPort, opts); err != nil {
		log.Fatalf("Failed to dial with Users endpoint: %v", err)
	}
	ordersMux := runtime.NewServeMux()
	if err := pbOrders.RegisterOrdersServiceHandlerFromEndpoint(ctx, ordersMux, host+":"+ordersPort, opts); err != nil {
		log.Fatalf("Failed to dial with Orders endpoint: %v", err)
	}
	// Handlers
	router := gin.Default()
	registerMux := func(url string, handlers ...gin.HandlerFunc) {
		routerGroup := router.Group(url, handlers...)
		routerGroup.Any("")
		routerGroup.Any("*any")
	}
	tokenHash := hmac.New(sha256.New, []byte(jwtKey))
	registerMux(authUrl, gin.WrapF(authMux.ServeHTTP))
	registerMux(usersUrl, AuthMW(tokenHash), gin.WrapF(usersMux.ServeHTTP))
	registerMux(ordersUrl, AuthMW(tokenHash), gin.WrapF(ordersMux.ServeHTTP))
	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
