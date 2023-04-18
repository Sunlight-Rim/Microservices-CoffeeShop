package gw

import (
	orders_pb "coffeeshop/internal/orders/pb"
	"strconv"

	// users_pb "coffeeshop/internal/users/pb"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ( // TODO: move to config
	grpcAddress = "localhost:50051"
	gwAddress   = "localhost:8080"
)

/// SERVER DEFINITION

type GatewayServer struct {
	server       *http.Server
	clientOrders orders_pb.OrdersServiceClient
}

func New() GatewayServer {
	// Connect to gRPC server
	conn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	client := orders_pb.NewOrdersServiceClient(conn)
	// REST router
	router := gin.Default()
	g := GatewayServer{
		server: &http.Server{
			Addr:    gwAddress,
			Handler: router,
		},
		clientOrders: client,
	}
	// Orders routing
	routerOrders := router.Group("/order")
	{
		routerOrders.POST("", marshalMiddleware(&orders_pb.CreateOrderRequest{}), g.createOrder)
		routerOrders.GET(":id", marshalMiddleware(&orders_pb.GetOrderRequest{}), g.getOrder)
		routerOrders.GET("", g.listOrder)
		routerOrders.PUT(":id", g.updateOrder)
		routerOrders.DELETE(":id", g.deleteOrder)
	}
	return g
}

func (g GatewayServer) Start() error {
	return g.server.ListenAndServe()
}

/// MIDDLEWARE

func marshalMiddleware(req proto.Message) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Request unmarshal
		err := jsonpb.Unmarshal(c.Request.Body, req)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error in your order request")
		}
		c.Set("req", req)
		// Perform requested method
		c.Next()
		// Send response
		resp, _ := c.MustGet("resp").(proto.Message)
		m := &jsonpb.Marshaler{}
		if err := m.Marshal(c.Writer, resp); err != nil {
			log.Print(err)
			c.String(http.StatusInternalServerError, "Error sending order response")
		}
	}
}

/// API METHODS (REST)

func (g GatewayServer) createOrder(c *gin.Context) {
	req, _ := c.MustGet("req").(*orders_pb.CreateOrderRequest)
	// Create order using Orders service
	resp, err := g.clientOrders.Create(c.Request.Context(), req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating order")
	}
	c.Set("resp", resp)
}

func (g GatewayServer) getOrder(c *gin.Context) {
	req, _ := c.MustGet("req").(*orders_pb.GetOrderRequest)
	println(c.Param("id"))
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error in order request parameter")
	}
	req.Ids = []int64{id}
	// Get order using Orders service
	resp, err := g.clientOrders.Get(c.Request.Context(), req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating order")
	}
	c.Set("resp", resp)
}

func (g GatewayServer) listOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yetx")
}

func (g GatewayServer) updateOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (g GatewayServer) deleteOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
