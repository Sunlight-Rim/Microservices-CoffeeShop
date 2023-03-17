package gw

import (
	orders_pb "coffeeshop/internal/orders/pb"

	// users_pb "coffeeshop/internal/users/pb"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ( // TODO: move to config
	grpcAddress = "localhost:50051"
	gwAddress   = "localhost:8080"
)

/// SERVER DEFINITION

type Gateway struct {
	server       *http.Server
	clientOrders orders_pb.OrdersServiceClient
}

func New() Gateway {
	// Connect to gRPC server
	conn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	client := orders_pb.NewOrdersServiceClient(conn)
	// REST router
	router := gin.Default()
	g := Gateway{
		server: &http.Server{
			Addr:    gwAddress,
			Handler: router,
		},
		clientOrders: client,
	}
	// Orders routing
	routerOrders := router.Group("/order")
	{
		routerOrders.POST("", g.createOrder)
		routerOrders.GET(":id", g.getOrder)
		routerOrders.GET("", g.listOrder)
		routerOrders.PUT(":id", g.updateOrder)
		routerOrders.DELETE(":id", g.deleteOrder)
	}
	return g
}

func (g Gateway) Start() error {
	return g.server.ListenAndServe()
}

/// API METHODS

func (g Gateway) createOrder(c *gin.Context) {
	var req orders_pb.CreateOrderRequest

	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating order request")
	}
	resp, err := g.clientOrders.Create(c.Request.Context(), &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating order")
	}
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "Error sending order response")
	}
}

func (g Gateway) getOrder(c *gin.Context) {
	log.Printf("id is: %v", c.Param("id"))
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (g Gateway) listOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (g Gateway) updateOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (g Gateway) deleteOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
