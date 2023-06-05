package gw

import (
	pbAuth "coffeeshop/internal/auth/pb"
	pbOrders "coffeeshop/internal/orders/pb"
	pbUsers "coffeeshop/internal/users/pb"
	"context"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/// REST SERVER

const ( // TODO: move to config
	restPort   = "8080"
	authPort   = "50050"
	usersPort  = "50052"
	ordersPort = "50051"
)

func Start() {
	log.Print("API gateway (REST->gRPC) server listening at http://localhost:"+restPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Dial with gRPC servers
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authMux := runtime.NewServeMux()
	if err := pbAuth.RegisterAuthServiceHandlerFromEndpoint(ctx, authMux, "localhost:"+authPort, opts); err != nil {
		log.Fatalf("Failed to dial with Auth endpoint: %v", err)
	}
	usersMux := runtime.NewServeMux()
	if err := pbUsers.RegisterUsersServiceHandlerFromEndpoint(ctx, usersMux, "localhost:"+usersPort, opts); err != nil {
		log.Fatalf("Failed to dial with Users endpoint: %v", err)
	}
	ordersMux := runtime.NewServeMux()
	if err := pbOrders.RegisterOrdersServiceHandlerFromEndpoint(ctx, ordersMux, "localhost:"+ordersPort, opts); err != nil {
		log.Fatalf("Failed to dial with Orders endpoint: %v", err)
	}
	// Handlers
	router := gin.Default()
	registerMux := func(url string, mux *runtime.ServeMux) {
		routerAuth := router.Group(url, gin.WrapF(mux.ServeHTTP))
		routerAuth.Any("")
		routerAuth.Any("*any")
	}
	registerMux("/auth", authMux)
	registerMux("/user", usersMux)
	registerMux("/order", ordersMux)
	// Start server
	if err := router.Run(":"+restPort); err != nil {
        log.Fatal(err)
    }
}