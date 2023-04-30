package gw

import (
	orders_pb "coffeeshop/internal/orders/pb"
	users_pb "coffeeshop/internal/users/pb"
	"strconv"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ( // TODO: move to config
	ordersPort = "50051"
	usersPort  = "50052"
	restPort   = "8080"
)

/// SERVER DEFINITION

type GatewayServer struct {
	restServer   *http.Server
	ordersClient orders_pb.OrdersServiceClient
	usersClient  users_pb.UsersServiceClient
}

func Start() error {
	log.Print("API gateway (REST->gRPC) server listening at http://localhost:" + restPort)
	// Connect to gRPC server
	ordersConn, err := grpc.Dial("localhost:"+ordersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	ordersClient := orders_pb.NewOrdersServiceClient(ordersConn)
	usersConn, err := grpc.Dial("localhost:"+usersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect: %v", err)
	}
	usersClient := users_pb.NewUsersServiceClient(usersConn)
	// REST router
	router := gin.Default()
	g := GatewayServer{
		restServer: &http.Server{
			Addr:    "localhost:" + restPort,
			Handler: router,
		},
		ordersClient: ordersClient,
		usersClient:  usersClient,
	}
	routerOrders := router.Group("/order")
	{
		routerOrders.POST("", marshalMW(&orders_pb.CreateOrderRequest{}), g.createOrder)
		// routerOrders.GET(":id", marshalMW(&orders_pb.GetOrderRequest{}), g.getOrder)
		// routerOrders.GET("", g.listOrders)
		// routerOrders.PUT(":id", g.updateOrder)
		// routerOrders.DELETE(":id", g.deleteOrder)
	}
	routerUsers := router.Group("/user")
	{
		routerUsers.POST("/signup", marshalMW(&users_pb.CreateUserRequest{}), g.createUser)
		routerUsers.POST("/login", marshalMW(&users_pb.LoginUserRequest{}), g.loginUser)
		routerUsers.GET("", marshalMW(&users_pb.GetUserRequest{}), g.getUser)
		// routerUsers.GET("", g.listUsers)
		// routerUsers.PUT(":id", g.updateUser)
		// routerUsers.DELETE(":id", g.deleteUser)
	}
	// Start server
	return g.restServer.ListenAndServe()
}

/// MIDDLEWARE

func marshalMW(req proto.Message) gin.HandlerFunc {
	u := &jsonpb.Unmarshaler{}
	m := &jsonpb.Marshaler{}
	return func(c *gin.Context) {
		req.Reset()
		// Unmarshal request
		if err := u.Unmarshal(c.Request.Body, req); err != nil && err.Error() != "EOF" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in your request"})
			c.Abort()
			return
		}
		c.Set("req", req)
		// Perform requested method
		c.Next()
		// Send response
		resp, ok := c.Get("resp")
		if !ok {
			return
		}
		if err := m.Marshal(c.Writer, resp.(proto.Message)); err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending response"})
		}
	}
}

/// API METHODS (REST)

func (g GatewayServer) createUser(c *gin.Context) {
	resp, err := g.usersClient.Create(c.Request.Context(), c.MustGet("req").(*users_pb.CreateUserRequest))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("resp", resp)
}

func (g GatewayServer) loginUser(c *gin.Context) {
	resp, err := g.usersClient.Login(c.Request.Context(), c.MustGet("req").(*users_pb.LoginUserRequest))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("resp", resp)
}

func (g GatewayServer) getUser(c *gin.Context) {
	req := c.MustGet("req").(*users_pb.GetUserRequest)
	req.Token = c.Request.Header.Get("token")
	if req.Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized"})
		return
	}
	resp, err := g.usersClient.Get(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("resp", resp)
}

func (g GatewayServer) createOrder(c *gin.Context) {
	req, _ := c.MustGet("req").(*orders_pb.CreateOrderRequest)
	// Create order using Orders service
	resp, err := g.ordersClient.Create(c.Request.Context(), req)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating order")
	}
	c.Set("resp", resp)
}

func (g GatewayServer) getOrder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error in url request parameter")
	}
	req, _ := c.MustGet("req").(*orders_pb.GetOrderRequest)
	req.Ids = []int64{id}
	// Get order using Orders service
	resp, err := g.ordersClient.Get(c.Request.Context(), req)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating order")
	}
	c.Set("resp", resp)
}

func (g GatewayServer) listOrders(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yetx")
}

func (g GatewayServer) updateOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (g GatewayServer) deleteOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
