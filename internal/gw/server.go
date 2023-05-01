package gw

import (
	ordersPB "coffeeshop/internal/orders/pb"
	usersPB "coffeeshop/internal/users/pb"
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
	restPort   = "8080"
	ordersPort = "50051"
	usersPort  = "50052"
)

/// SERVER DEFINITION

type GatewayServer struct {
	restServer   *http.Server
	ordersClient ordersPB.OrdersServiceClient
	usersClient  usersPB.UsersServiceClient
}

func Start() error {
	log.Print("API gateway (REST->gRPC) server listening at http://localhost:" + restPort)
	// Connect to gRPC servers
	ordersConn, err := grpc.Dial("localhost:"+ordersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to gRPC: %v", err)
	}
	ordersClient := ordersPB.NewOrdersServiceClient(ordersConn)
	usersConn, err := grpc.Dial("localhost:"+usersPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to gRPC: %v", err)
	}
	usersClient := usersPB.NewUsersServiceClient(usersConn)
	// REST router
	router := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
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
		routerOrders.POST("", marshalMW(&ordersPB.CreateOrderRequest{}), g.createOrder)
		// routerOrders.GET(":id", marshalMW(&ordersPB.GetOrderRequest{}), g.getOrder)
		// routerOrders.GET("", g.listOrders)
		// routerOrders.PUT(":id", g.updateOrder)
		// routerOrders.DELETE(":id", g.deleteOrder)
	}
	routerUsers := router.Group("/user")
	{
		routerUsers.POST("/signup", marshalMW(&usersPB.CreateUserRequest{}), g.createUser)
		routerUsers.POST("/login", marshalMW(&usersPB.LoginUserRequest{}), g.loginUser)
		routerUsers.GET("", authToken, marshalMW(&usersPB.GetUserRequest{}), g.getUser)
		routerUsers.PATCH("", authToken, marshalMW(&usersPB.UpdateUserRequest{}), g.updateUser)
		// routerUsers.DELETE(":id", g.deleteUser)
	}
	// Start server
	return g.restServer.ListenAndServe()
}

/// MIDDLEWARE

func authToken(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized"})
		return
	}
	c.Set("token", token)
}

func marshalMW(req proto.Message) gin.HandlerFunc {
	u := &jsonpb.Unmarshaler{}
	m := &jsonpb.Marshaler{}
	return func(c *gin.Context) {
		req.Reset()
		// Unmarshal request
		if err := u.Unmarshal(c.Request.Body, req); err != nil {
			// Check token
			if _, ok := c.Get("token"); !ok || err.Error() != "EOF" {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error in your request"})
				return
			}
		}
		c.Set("req", req)
		// Perform requested method
		c.Next()
		if err, _ := c.Get("err"); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.(error).Error()})
			return
		}
		// Send a response
		if err := m.Marshal(c.Writer, c.MustGet("resp").(proto.Message)); err != nil {
			log.Print(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error sending response"})
		}
	}
}

/// API METHODS (REST)

func (g GatewayServer) createUser(c *gin.Context) {
	resp, err := g.usersClient.Create(c.Request.Context(), c.MustGet("req").(*usersPB.CreateUserRequest))
	c.Set("err", err)
	c.Set("resp", resp)
}

func (g GatewayServer) loginUser(c *gin.Context) {
	resp, err := g.usersClient.Login(c.Request.Context(), c.MustGet("req").(*usersPB.LoginUserRequest))
	c.Set("err", err)
	c.Set("resp", resp)
}

func (g GatewayServer) getUser(c *gin.Context) {
	req := c.MustGet("req").(*usersPB.GetUserRequest)
	req.Token = c.MustGet("token").(string)
	resp, err := g.usersClient.Get(c.Request.Context(), req)
	c.Set("err", err)
	c.Set("resp", resp)
}

func (g GatewayServer) updateUser(c *gin.Context) {
	req := c.MustGet("req").(*usersPB.UpdateUserRequest)
	req.Token = c.MustGet("token").(string)
	resp, err := g.usersClient.Update(c.Request.Context(), req)
	c.Set("err", err)
	c.Set("resp", resp)
}

func (g GatewayServer) createOrder(c *gin.Context) {
	req, _ := c.MustGet("req").(*ordersPB.CreateOrderRequest)
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
	req, _ := c.MustGet("req").(*ordersPB.GetOrderRequest)
	req.Ids = []int64{id}
	// Get order using Orders service
	resp, err := g.ordersClient.Get(c.Request.Context(), req)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error creating order")
	}
	c.Set("resp", resp)
}

func (g GatewayServer) listOrders(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (g GatewayServer) updateOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (g GatewayServer) deleteOrder(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
