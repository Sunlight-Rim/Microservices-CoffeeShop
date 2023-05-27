package gw

import (
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
	ordersPort = "50051"
	usersPort  = "50052"
)

func Start() {
	log.Print("API gateway (REST->gRPC) server listening at http://localhost:"+restPort)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Dial with gRPC servers
	ordersMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pbOrders.RegisterOrdersServiceHandlerFromEndpoint(ctx, ordersMux, "localhost:"+ordersPort, opts); err != nil {
		log.Fatalf("Failed to dial with Orders endpoint: %v", err)
	}
	usersMux := runtime.NewServeMux()
	if err := pbUsers.RegisterUsersServiceHandlerFromEndpoint(ctx, usersMux, "localhost:"+usersPort, opts); err != nil {
		log.Fatalf("Failed to dial with Users endpoint: %v", err)
	}
	// Handlers
	router := gin.Default()
	router.Any("/*any",
		gin.WrapF(ordersMux.ServeHTTP),
		gin.WrapF(usersMux.ServeHTTP),
	)
	// Start server
	if err := router.Run(":"+restPort); err != nil {
        log.Fatal(err)
    }
}