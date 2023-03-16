// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/orders/proto/orders.proto

package orders_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Order_Status int32

const (
	Order_PENDING   Order_Status = 0
	Order_DELIVERED Order_Status = 1
)

var Order_Status_name = map[int32]string{
	0: "PENDING",
	1: "DELIVERED",
}

var Order_Status_value = map[string]int32{
	"PENDING":   0,
	"DELIVERED": 1,
}

func (x Order_Status) String() string {
	return proto.EnumName(Order_Status_name, int32(x))
}

func (Order_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{1, 0}
}

type Coffee struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coffee) Reset()         { *m = Coffee{} }
func (m *Coffee) String() string { return proto.CompactTextString(m) }
func (*Coffee) ProtoMessage()    {}
func (*Coffee) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{0}
}

func (m *Coffee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coffee.Unmarshal(m, b)
}
func (m *Coffee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coffee.Marshal(b, m, deterministic)
}
func (m *Coffee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coffee.Merge(m, src)
}
func (m *Coffee) XXX_Size() int {
	return xxx_messageInfo_Coffee.Size(m)
}
func (m *Coffee) XXX_DiscardUnknown() {
	xxx_messageInfo_Coffee.DiscardUnknown(m)
}

var xxx_messageInfo_Coffee proto.InternalMessageInfo

func (m *Coffee) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Coffee) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Order struct {
	OrderId              int64                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Coffees              []*Coffee              `protobuf:"bytes,2,rep,name=coffees,proto3" json:"coffees,omitempty"`
	Total                float32                `protobuf:"fixed32,3,opt,name=total,proto3" json:"total,omitempty"`
	Date                 *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Status               Order_Status           `protobuf:"varint,5,opt,name=status,proto3,enum=orders.Order_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *Order) GetCoffees() []*Coffee {
	if m != nil {
		return m.Coffees
	}
	return nil
}

func (m *Order) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Order) GetDate() *timestamppb.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Order) GetStatus() Order_Status {
	if m != nil {
		return m.Status
	}
	return Order_PENDING
}

type CreateOrderRequest struct {
	Token                string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               int64     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Coffees              []*Coffee `protobuf:"bytes,3,rep,name=coffees,proto3" json:"coffees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{2}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreateOrderRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateOrderRequest) GetCoffees() []*Coffee {
	if m != nil {
		return m.Coffees
	}
	return nil
}

type CreateOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{3}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type GetOrderRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Ids                  []int64  `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderRequest) Reset()         { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()    {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{4}
}

func (m *GetOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderRequest.Unmarshal(m, b)
}
func (m *GetOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderRequest.Merge(m, src)
}
func (m *GetOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderRequest.Size(m)
}
func (m *GetOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderRequest proto.InternalMessageInfo

func (m *GetOrderRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GetOrderRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetOrderRequest) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetOrderResponse struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderResponse) Reset()         { *m = GetOrderResponse{} }
func (m *GetOrderResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderResponse) ProtoMessage()    {}
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{5}
}

func (m *GetOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderResponse.Unmarshal(m, b)
}
func (m *GetOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderResponse.Merge(m, src)
}
func (m *GetOrderResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderResponse.Size(m)
}
func (m *GetOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderResponse proto.InternalMessageInfo

func (m *GetOrderResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type UpdateOrderRequest struct {
	Token                string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               int64     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId              int64     `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Coffees              []*Coffee `protobuf:"bytes,4,rep,name=coffees,proto3" json:"coffees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateOrderRequest) Reset()         { *m = UpdateOrderRequest{} }
func (m *UpdateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateOrderRequest) ProtoMessage()    {}
func (*UpdateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{6}
}

func (m *UpdateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOrderRequest.Unmarshal(m, b)
}
func (m *UpdateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOrderRequest.Marshal(b, m, deterministic)
}
func (m *UpdateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOrderRequest.Merge(m, src)
}
func (m *UpdateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateOrderRequest.Size(m)
}
func (m *UpdateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOrderRequest proto.InternalMessageInfo

func (m *UpdateOrderRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateOrderRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateOrderRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *UpdateOrderRequest) GetCoffees() []*Coffee {
	if m != nil {
		return m.Coffees
	}
	return nil
}

type UpdateOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateOrderResponse) Reset()         { *m = UpdateOrderResponse{} }
func (m *UpdateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateOrderResponse) ProtoMessage()    {}
func (*UpdateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{7}
}

func (m *UpdateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOrderResponse.Unmarshal(m, b)
}
func (m *UpdateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOrderResponse.Marshal(b, m, deterministic)
}
func (m *UpdateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOrderResponse.Merge(m, src)
}
func (m *UpdateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateOrderResponse.Size(m)
}
func (m *UpdateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOrderResponse proto.InternalMessageInfo

func (m *UpdateOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type DeleteOrderRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId              int64    `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOrderRequest) Reset()         { *m = DeleteOrderRequest{} }
func (m *DeleteOrderRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteOrderRequest) ProtoMessage()    {}
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{8}
}

func (m *DeleteOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOrderRequest.Unmarshal(m, b)
}
func (m *DeleteOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOrderRequest.Marshal(b, m, deterministic)
}
func (m *DeleteOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOrderRequest.Merge(m, src)
}
func (m *DeleteOrderRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteOrderRequest.Size(m)
}
func (m *DeleteOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOrderRequest proto.InternalMessageInfo

func (m *DeleteOrderRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeleteOrderRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DeleteOrderRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type DeleteOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOrderResponse) Reset()         { *m = DeleteOrderResponse{} }
func (m *DeleteOrderResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteOrderResponse) ProtoMessage()    {}
func (*DeleteOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b42e94a9bc8688, []int{9}
}

func (m *DeleteOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOrderResponse.Unmarshal(m, b)
}
func (m *DeleteOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOrderResponse.Marshal(b, m, deterministic)
}
func (m *DeleteOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOrderResponse.Merge(m, src)
}
func (m *DeleteOrderResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteOrderResponse.Size(m)
}
func (m *DeleteOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOrderResponse proto.InternalMessageInfo

func (m *DeleteOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func init() {
	proto.RegisterEnum("orders.Order_Status", Order_Status_name, Order_Status_value)
	proto.RegisterType((*Coffee)(nil), "orders.Coffee")
	proto.RegisterType((*Order)(nil), "orders.Order")
	proto.RegisterType((*CreateOrderRequest)(nil), "orders.CreateOrderRequest")
	proto.RegisterType((*CreateOrderResponse)(nil), "orders.CreateOrderResponse")
	proto.RegisterType((*GetOrderRequest)(nil), "orders.GetOrderRequest")
	proto.RegisterType((*GetOrderResponse)(nil), "orders.GetOrderResponse")
	proto.RegisterType((*UpdateOrderRequest)(nil), "orders.UpdateOrderRequest")
	proto.RegisterType((*UpdateOrderResponse)(nil), "orders.UpdateOrderResponse")
	proto.RegisterType((*DeleteOrderRequest)(nil), "orders.DeleteOrderRequest")
	proto.RegisterType((*DeleteOrderResponse)(nil), "orders.DeleteOrderResponse")
}

func init() {
	proto.RegisterFile("internal/orders/proto/orders.proto", fileDescriptor_15b42e94a9bc8688)
}

var fileDescriptor_15b42e94a9bc8688 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x76, 0xe2, 0xd0, 0xb1, 0x52, 0xa2, 0xa1, 0x52, 0x8d, 0x39, 0x60, 0x19, 0x90, 0x7c,
	0x40, 0x8e, 0x64, 0x2e, 0x7c, 0x9c, 0x42, 0x13, 0x45, 0x91, 0x50, 0x41, 0x5b, 0xe0, 0xc0, 0x81,
	0xca, 0x89, 0x27, 0x95, 0x45, 0x12, 0x1b, 0xef, 0x06, 0x89, 0x7f, 0xc0, 0x99, 0x1f, 0xc9, 0xef,
	0x40, 0xde, 0x5d, 0x63, 0xdc, 0x04, 0x50, 0x55, 0xf5, 0xb6, 0x6f, 0x67, 0x76, 0xde, 0x9b, 0xe7,
	0x27, 0x43, 0x90, 0x6d, 0x04, 0x95, 0x9b, 0x64, 0x35, 0xcc, 0xcb, 0x94, 0x4a, 0x3e, 0x2c, 0xca,
	0x5c, 0xe4, 0x1a, 0x44, 0x12, 0xa0, 0xad, 0x90, 0xf7, 0xe0, 0x22, 0xcf, 0x2f, 0x56, 0xa4, 0x5a,
	0xe6, 0xdb, 0xe5, 0x50, 0x64, 0x6b, 0xe2, 0x22, 0x59, 0x17, 0xaa, 0x31, 0x88, 0xc1, 0x3e, 0xc9,
	0x97, 0x4b, 0x22, 0x44, 0xe8, 0x88, 0x6f, 0x05, 0xb9, 0x86, 0x6f, 0x84, 0x07, 0x4c, 0x9e, 0xf1,
	0x08, 0xba, 0x45, 0x99, 0x2d, 0xc8, 0x35, 0x7d, 0x23, 0x34, 0x99, 0x02, 0xc1, 0x4f, 0x03, 0xba,
	0x6f, 0xaa, 0xf9, 0x78, 0x0f, 0x6e, 0x4b, 0xa2, 0xf3, 0x2c, 0x95, 0xef, 0x2c, 0xd6, 0x93, 0x78,
	0x96, 0x62, 0x08, 0xbd, 0x85, 0x1c, 0xcc, 0x5d, 0xd3, 0xb7, 0x42, 0x27, 0x3e, 0x8c, 0xb4, 0x42,
	0xc5, 0xc7, 0xea, 0x72, 0x45, 0x22, 0x72, 0x91, 0xac, 0x5c, 0x4b, 0x91, 0x48, 0x80, 0x11, 0x74,
	0xd2, 0x44, 0x90, 0xdb, 0xf1, 0x8d, 0xd0, 0x89, 0xbd, 0x48, 0x2d, 0x12, 0xd5, 0x8b, 0x44, 0xef,
	0xea, 0x45, 0x98, 0xec, 0xc3, 0x27, 0x60, 0x73, 0x91, 0x88, 0x2d, 0x77, 0xbb, 0xbe, 0x11, 0x1e,
	0xc6, 0x47, 0x35, 0x9d, 0x54, 0x1a, 0x9d, 0xc9, 0x1a, 0xd3, 0x3d, 0xc1, 0x23, 0xb0, 0xd5, 0x0d,
	0x3a, 0xd0, 0x7b, 0x3b, 0x39, 0x1d, 0xcf, 0x4e, 0xa7, 0x83, 0x5b, 0xd8, 0x87, 0x83, 0xf1, 0xe4,
	0xf5, 0xec, 0xc3, 0x84, 0x4d, 0xc6, 0x03, 0x23, 0x58, 0x03, 0x9e, 0x94, 0x94, 0x08, 0x92, 0x33,
	0x18, 0x7d, 0xd9, 0x12, 0x17, 0x4a, 0xef, 0x67, 0xda, 0x68, 0xa7, 0x14, 0xc0, 0x63, 0xe8, 0x6d,
	0xb9, 0x72, 0xc2, 0x94, 0x4e, 0xd8, 0x15, 0x6c, 0x1b, 0x61, 0xfd, 0xd3, 0x88, 0xe0, 0x05, 0xdc,
	0x6d, 0xd1, 0xf1, 0x22, 0xdf, 0x70, 0xc2, 0x87, 0xd0, 0x95, 0x0f, 0x24, 0x9f, 0x13, 0xf7, 0x5b,
	0x8b, 0x31, 0x55, 0x0b, 0x18, 0xdc, 0x99, 0x92, 0xb8, 0x8e, 0xce, 0x01, 0x58, 0x59, 0xaa, 0x34,
	0x5a, 0xac, 0x3a, 0x06, 0xcf, 0x61, 0xd0, 0xcc, 0xd4, 0x62, 0x1e, 0x83, 0x8e, 0x96, 0x6b, 0xc8,
	0x65, 0x2e, 0xa9, 0xd1, 0xc5, 0xe0, 0xbb, 0x01, 0xf8, 0xbe, 0x48, 0xaf, 0x69, 0xdd, 0x9f, 0xf1,
	0xb2, 0xfe, 0x1a, 0xaf, 0xce, 0x7f, 0x5d, 0x6d, 0x29, 0xb9, 0x8a, 0xab, 0x9f, 0x00, 0xc7, 0xb4,
	0xa2, 0x9b, 0xda, 0xa2, 0xd2, 0xd6, 0x9a, 0x7f, 0x05, 0x6d, 0xf1, 0x0f, 0x13, 0xfa, 0xf2, 0x82,
	0x9f, 0x51, 0xf9, 0x35, 0x5b, 0x10, 0x8e, 0xc0, 0x56, 0xf9, 0x41, 0xef, 0xb7, 0x19, 0x3b, 0xf1,
	0xf5, 0xee, 0xef, 0xad, 0x69, 0xe6, 0x67, 0x60, 0x4d, 0x49, 0xe0, 0x71, 0xdd, 0x73, 0x29, 0x53,
	0x9e, 0xbb, 0x5b, 0xd0, 0x2f, 0x47, 0x60, 0x2b, 0x9b, 0x1b, 0xf2, 0xdd, 0x00, 0x34, 0xe4, 0xfb,
	0x3e, 0xc9, 0x08, 0x6c, 0xe5, 0x46, 0x33, 0x62, 0xd7, 0xfd, 0x66, 0xc4, 0x1e, 0xe7, 0x5e, 0xf5,
	0x3f, 0x3a, 0xd1, 0x4b, 0x55, 0x3f, 0x2f, 0xe6, 0x73, 0x5b, 0xfe, 0x2e, 0x9e, 0xfe, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x55, 0xbd, 0x51, 0x7c, 0x33, 0x05, 0x00, 0x00,
}
