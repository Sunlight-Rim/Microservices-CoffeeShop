// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: internal/orders/proto/orders.proto

package orders_pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order_Status int32

const (
	Order_PENDING   Order_Status = 0
	Order_DELIVERED Order_Status = 1
)

// Enum value maps for Order_Status.
var (
	Order_Status_name = map[int32]string{
		0: "PENDING",
		1: "DELIVERED",
	}
	Order_Status_value = map[string]int32{
		"PENDING":   0,
		"DELIVERED": 1,
	}
)

func (x Order_Status) Enum() *Order_Status {
	p := new(Order_Status)
	*p = x
	return p
}

func (x Order_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_orders_proto_orders_proto_enumTypes[0].Descriptor()
}

func (Order_Status) Type() protoreflect.EnumType {
	return &file_internal_orders_proto_orders_proto_enumTypes[0]
}

func (x Order_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order_Status.Descriptor instead.
func (Order_Status) EnumDescriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{1, 0}
}

type Coffee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sugar int32  `protobuf:"varint,2,opt,name=sugar,proto3" json:"sugar,omitempty"`
}

func (x *Coffee) Reset() {
	*x = Coffee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coffee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coffee) ProtoMessage() {}

func (x *Coffee) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coffee.ProtoReflect.Descriptor instead.
func (*Coffee) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{0}
}

func (x *Coffee) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Coffee) GetSugar() int32 {
	if x != nil {
		return x.Sugar
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid  int64                `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Status  Order_Status         `protobuf:"varint,3,opt,name=status,proto3,enum=orders.Order_Status" json:"status,omitempty"`
	Date    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Coffees []*Coffee            `protobuf:"bytes,5,rep,name=coffees,proto3" json:"coffees,omitempty"`
	Total   float32              `protobuf:"fixed32,6,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *Order) GetStatus() Order_Status {
	if x != nil {
		return x.Status
	}
	return Order_PENDING
}

func (x *Order) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Order) GetCoffees() []*Coffee {
	if x != nil {
		return x.Coffees
	}
	return nil
}

func (x *Order) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid  int64     `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Coffees []*Coffee `protobuf:"bytes,2,rep,name=coffees,proto3" json:"coffees,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *CreateOrderRequest) GetCoffees() []*Coffee {
	if x != nil {
		return x.Coffees
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Id     int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *GetOrderRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type ListOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Shift  int32 `protobuf:"varint,2,opt,name=shift,proto3" json:"shift,omitempty"`
}

func (x *ListOrderRequest) Reset() {
	*x = ListOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderRequest) ProtoMessage() {}

func (x *ListOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderRequest.ProtoReflect.Descriptor instead.
func (*ListOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{6}
}

func (x *ListOrderRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *ListOrderRequest) GetShift() int32 {
	if x != nil {
		return x.Shift
	}
	return 0
}

type ListOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *ListOrderResponse) Reset() {
	*x = ListOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrderResponse) ProtoMessage() {}

func (x *ListOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrderResponse.ProtoReflect.Descriptor instead.
func (*ListOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{7}
}

func (x *ListOrderResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type UpdateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid  int64     `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Id      int64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Coffees []*Coffee `protobuf:"bytes,3,rep,name=coffees,proto3" json:"coffees,omitempty"`
}

func (x *UpdateOrderRequest) Reset() {
	*x = UpdateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRequest) ProtoMessage() {}

func (x *UpdateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOrderRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *UpdateOrderRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateOrderRequest) GetCoffees() []*Coffee {
	if x != nil {
		return x.Coffees
	}
	return nil
}

type UpdateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpdateOrderResponse) Reset() {
	*x = UpdateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderResponse) ProtoMessage() {}

func (x *UpdateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type DeleteOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid int64 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Id     int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteOrderRequest) Reset() {
	*x = DeleteOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRequest) ProtoMessage() {}

func (x *DeleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderRequest.ProtoReflect.Descriptor instead.
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteOrderRequest) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *DeleteOrderRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *DeleteOrderResponse) Reset() {
	*x = DeleteOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_orders_proto_orders_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderResponse) ProtoMessage() {}

func (x *DeleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_orders_proto_orders_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderResponse.ProtoReflect.Descriptor instead.
func (*DeleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_orders_proto_orders_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_internal_orders_proto_orders_proto protoreflect.FileDescriptor

var file_internal_orders_proto_orders_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x06, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x75, 0x67, 0x61, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x75, 0x67, 0x61, 0x72, 0x22, 0xf3, 0x01, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f,
	0x66, 0x66, 0x65, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x24, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01,
	0x22, 0x56, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x28,
	0x0a, 0x07, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x52,
	0x07, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x40, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x69, 0x66, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x68, 0x69, 0x66, 0x74, 0x22, 0x3a, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x43,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x22, 0x3a,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x32, 0xbf, 0x03, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x07, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x22, 0x07, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x4c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x5a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x32,
	0x0c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x57, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_orders_proto_orders_proto_rawDescOnce sync.Once
	file_internal_orders_proto_orders_proto_rawDescData = file_internal_orders_proto_orders_proto_rawDesc
)

func file_internal_orders_proto_orders_proto_rawDescGZIP() []byte {
	file_internal_orders_proto_orders_proto_rawDescOnce.Do(func() {
		file_internal_orders_proto_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_orders_proto_orders_proto_rawDescData)
	})
	return file_internal_orders_proto_orders_proto_rawDescData
}

var file_internal_orders_proto_orders_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_orders_proto_orders_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_orders_proto_orders_proto_goTypes = []interface{}{
	(Order_Status)(0),           // 0: orders.Order.Status
	(*Coffee)(nil),              // 1: orders.Coffee
	(*Order)(nil),               // 2: orders.Order
	(*CreateOrderRequest)(nil),  // 3: orders.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 4: orders.CreateOrderResponse
	(*GetOrderRequest)(nil),     // 5: orders.GetOrderRequest
	(*GetOrderResponse)(nil),    // 6: orders.GetOrderResponse
	(*ListOrderRequest)(nil),    // 7: orders.ListOrderRequest
	(*ListOrderResponse)(nil),   // 8: orders.ListOrderResponse
	(*UpdateOrderRequest)(nil),  // 9: orders.UpdateOrderRequest
	(*UpdateOrderResponse)(nil), // 10: orders.UpdateOrderResponse
	(*DeleteOrderRequest)(nil),  // 11: orders.DeleteOrderRequest
	(*DeleteOrderResponse)(nil), // 12: orders.DeleteOrderResponse
	(*timestamp.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_internal_orders_proto_orders_proto_depIdxs = []int32{
	0,  // 0: orders.Order.status:type_name -> orders.Order.Status
	13, // 1: orders.Order.date:type_name -> google.protobuf.Timestamp
	1,  // 2: orders.Order.coffees:type_name -> orders.Coffee
	1,  // 3: orders.CreateOrderRequest.coffees:type_name -> orders.Coffee
	2,  // 4: orders.CreateOrderResponse.order:type_name -> orders.Order
	2,  // 5: orders.GetOrderResponse.order:type_name -> orders.Order
	2,  // 6: orders.ListOrderResponse.orders:type_name -> orders.Order
	1,  // 7: orders.UpdateOrderRequest.coffees:type_name -> orders.Coffee
	2,  // 8: orders.UpdateOrderResponse.order:type_name -> orders.Order
	2,  // 9: orders.DeleteOrderResponse.order:type_name -> orders.Order
	3,  // 10: orders.OrdersService.Create:input_type -> orders.CreateOrderRequest
	5,  // 11: orders.OrdersService.Get:input_type -> orders.GetOrderRequest
	7,  // 12: orders.OrdersService.List:input_type -> orders.ListOrderRequest
	9,  // 13: orders.OrdersService.Update:input_type -> orders.UpdateOrderRequest
	11, // 14: orders.OrdersService.Delete:input_type -> orders.DeleteOrderRequest
	4,  // 15: orders.OrdersService.Create:output_type -> orders.CreateOrderResponse
	6,  // 16: orders.OrdersService.Get:output_type -> orders.GetOrderResponse
	8,  // 17: orders.OrdersService.List:output_type -> orders.ListOrderResponse
	10, // 18: orders.OrdersService.Update:output_type -> orders.UpdateOrderResponse
	12, // 19: orders.OrdersService.Delete:output_type -> orders.DeleteOrderResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_internal_orders_proto_orders_proto_init() }
func file_internal_orders_proto_orders_proto_init() {
	if File_internal_orders_proto_orders_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_orders_proto_orders_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coffee); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_orders_proto_orders_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_orders_proto_orders_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_orders_proto_orders_proto_goTypes,
		DependencyIndexes: file_internal_orders_proto_orders_proto_depIdxs,
		EnumInfos:         file_internal_orders_proto_orders_proto_enumTypes,
		MessageInfos:      file_internal_orders_proto_orders_proto_msgTypes,
	}.Build()
	File_internal_orders_proto_orders_proto = out.File
	file_internal_orders_proto_orders_proto_rawDesc = nil
	file_internal_orders_proto_orders_proto_goTypes = nil
	file_internal_orders_proto_orders_proto_depIdxs = nil
}
