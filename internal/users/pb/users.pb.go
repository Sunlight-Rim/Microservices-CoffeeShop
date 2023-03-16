// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/users/proto/users.proto

package users_pb

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

type User struct {
	UserId               int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Date                 *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	OrderIds             []int64                `protobuf:"varint,5,rep,packed,name=order_ids,json=orderIds,proto3" json:"order_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetDate() *timestamppb.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *User) GetOrderIds() []int64 {
	if m != nil {
		return m.OrderIds
	}
	return nil
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pass                 string   `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

func (m *CreateUserRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	UserIds              []int64  `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{3}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserIds() []int64 {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type GetUserResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{4}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UpdateUserRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{5}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UpdateUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{6}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

func (m *UpdateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{7}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeleteUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type DeleteUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2602c7761890406d, []int{8}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterType((*CreateUserRequest)(nil), "users.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "users.CreateUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "users.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "users.GetUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "users.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "users.UpdateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "users.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "users.DeleteUserResponse")
}

func init() {
	proto.RegisterFile("internal/users/proto/users.proto", fileDescriptor_2602c7761890406d)
}

var fileDescriptor_2602c7761890406d = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0xaf, 0x9a, 0x40,
	0x10, 0x0f, 0x82, 0xf8, 0xde, 0xf8, 0xd2, 0xc6, 0x4d, 0xff, 0xec, 0xa3, 0x87, 0x47, 0x39, 0x91,
	0x34, 0xc1, 0xc4, 0x3e, 0x4f, 0x8d, 0x17, 0xdb, 0xc4, 0x78, 0xa5, 0xf5, 0xd2, 0x8b, 0x41, 0x77,
	0x6a, 0x48, 0x15, 0xe8, 0xce, 0xda, 0x7b, 0xbf, 0x46, 0x3f, 0x6d, 0xc3, 0x2e, 0x10, 0x14, 0x4d,
	0x4c, 0x7a, 0xdb, 0xd9, 0xf9, 0xcd, 0xfe, 0xfe, 0x0c, 0x80, 0x9f, 0x66, 0x0a, 0x65, 0x96, 0xec,
	0xc7, 0x47, 0x42, 0x49, 0xe3, 0x42, 0xe6, 0x2a, 0x37, 0xe7, 0x48, 0x9f, 0x59, 0x5f, 0x17, 0xde,
	0xd3, 0x2e, 0xcf, 0x77, 0x7b, 0x34, 0x80, 0xcd, 0xf1, 0xc7, 0x58, 0xa5, 0x07, 0x24, 0x95, 0x1c,
	0x0a, 0x83, 0x0b, 0xfe, 0x5a, 0xe0, 0xac, 0x08, 0x25, 0x7b, 0x0b, 0x83, 0x72, 0x64, 0x9d, 0x0a,
	0x6e, 0xf9, 0x56, 0x68, 0xc7, 0x6e, 0x59, 0x2e, 0x05, 0x63, 0xe0, 0x64, 0xc9, 0x01, 0x79, 0xcf,
	0xb7, 0xc2, 0xfb, 0x58, 0x9f, 0x19, 0x87, 0x41, 0x22, 0x84, 0x44, 0x22, 0x6e, 0xeb, 0xeb, 0xba,
	0x64, 0x11, 0x38, 0x22, 0x51, 0xc8, 0x1d, 0xdf, 0x0a, 0x87, 0x13, 0x2f, 0x32, 0xfc, 0x51, 0xcd,
	0x1f, 0x7d, 0xab, 0xf9, 0x63, 0x8d, 0x63, 0xef, 0xe0, 0x3e, 0x97, 0x42, 0xf3, 0x12, 0xef, 0xfb,
	0x76, 0x68, 0xc7, 0x77, 0xfa, 0x62, 0x29, 0x28, 0x58, 0xc1, 0xe8, 0xb3, 0xc4, 0x44, 0x61, 0xa9,
	0x30, 0xc6, 0x5f, 0x47, 0x24, 0xd5, 0xe8, 0xb1, 0x5a, 0x7a, 0x18, 0x38, 0x45, 0x42, 0x54, 0x6b,
	0x2c, 0xcf, 0xd7, 0x35, 0x06, 0x53, 0x60, 0xed, 0x67, 0xa9, 0xc8, 0x33, 0x42, 0xf6, 0x04, 0x4e,
	0xe9, 0x58, 0xbf, 0x3b, 0x9c, 0x0c, 0x23, 0x93, 0xa6, 0x86, 0xe8, 0x46, 0xf0, 0x01, 0x5e, 0x2c,
	0x50, 0xb5, 0xa5, 0x3c, 0xc2, 0x5d, 0x95, 0x19, 0x71, 0x4b, 0x6b, 0x1f, 0x98, 0xd0, 0x28, 0x78,
	0x86, 0x97, 0x0d, 0xb8, 0x22, 0x78, 0x0f, 0x66, 0x29, 0x1a, 0x7a, 0xc6, 0x60, 0x3a, 0xc1, 0x16,
	0x46, 0xab, 0x42, 0x9c, 0x19, 0x7e, 0x05, 0x7d, 0x95, 0xff, 0xc4, 0xac, 0x72, 0x6c, 0x8a, 0xf6,
	0xbe, 0x7a, 0x27, 0xfb, 0xaa, 0x7d, 0xd8, 0xd7, 0x7c, 0x4c, 0x81, 0xb5, 0x49, 0x6e, 0xb5, 0x3f,
	0x87, 0xd1, 0x17, 0xdc, 0xe3, 0xff, 0x68, 0x2b, 0xa9, 0xdb, 0x6f, 0xdc, 0x48, 0x3d, 0xf9, 0xd3,
	0x83, 0x87, 0xb2, 0xa4, 0xaf, 0x28, 0x7f, 0xa7, 0x5b, 0x64, 0x33, 0x70, 0xcd, 0x06, 0x19, 0xaf,
	0xd0, 0x9d, 0xef, 0xc4, 0x7b, 0xbc, 0xd0, 0xa9, 0x08, 0x9f, 0xc1, 0x5e, 0xa0, 0x62, 0xaf, 0x2b,
	0xc4, 0xe9, 0x56, 0xbd, 0x37, 0xe7, 0xd7, 0xd5, 0xd4, 0x0c, 0x5c, 0x93, 0x5b, 0x43, 0xda, 0xd9,
	0x55, 0x43, 0x7a, 0x21, 0xe0, 0x19, 0xb8, 0xc6, 0x7b, 0x33, 0xde, 0x89, 0xb3, 0x19, 0xef, 0x86,
	0x34, 0x7f, 0xf8, 0x0e, 0xd1, 0x27, 0xdd, 0x5d, 0x17, 0x9b, 0x8d, 0xab, 0x7f, 0xa8, 0x8f, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xdb, 0x27, 0xb5, 0x09, 0x04, 0x00, 0x00,
}
