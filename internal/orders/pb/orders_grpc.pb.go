// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: proto/orders.proto

package orders_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrdersServiceClient interface {
	Create(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	Get(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	Update(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	Delete(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) Create(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) Get(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) Update(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) Delete(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, "/orders.OrdersService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
// All implementations must embed UnimplementedOrdersServiceServer
// for forward compatibility
type OrdersServiceServer interface {
	Create(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	Get(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	Update(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	Delete(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error)
	mustEmbedUnimplementedOrdersServiceServer()
}

// UnimplementedOrdersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrdersServiceServer struct {
}

func (UnimplementedOrdersServiceServer) Create(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrdersServiceServer) Get(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOrdersServiceServer) Update(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrdersServiceServer) Delete(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOrdersServiceServer) mustEmbedUnimplementedOrdersServiceServer() {}

// UnsafeOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServiceServer will
// result in compilation errors.
type UnsafeOrdersServiceServer interface {
	mustEmbedUnimplementedOrdersServiceServer()
}

func RegisterOrdersServiceServer(s grpc.ServiceRegistrar, srv OrdersServiceServer) {
	s.RegisterService(&OrdersService_ServiceDesc, srv)
}

func _OrdersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Create(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Get(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Update(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrdersService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).Delete(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersService_ServiceDesc is the grpc.ServiceDesc for OrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrdersService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _OrdersService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrdersService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrdersService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/orders.proto",
}