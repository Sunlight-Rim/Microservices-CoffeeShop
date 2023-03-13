// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users.proto

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
	Id                   int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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
	return fileDescriptor_b1c161a4c7514913, []int{0}
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

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
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
	return fileDescriptor_b1c161a4c7514913, []int{1}
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
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{2}
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

func (m *CreateUserResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetUserRequest struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{3}
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

func (m *GetUserRequest) GetIds() []int64 {
	if m != nil {
		return m.Ids
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
	return fileDescriptor_b1c161a4c7514913, []int{4}
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
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{5}
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

func (m *UpdateUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
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
	return fileDescriptor_b1c161a4c7514913, []int{6}
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
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1c161a4c7514913, []int{7}
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

func (m *DeleteUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
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
	return fileDescriptor_b1c161a4c7514913, []int{8}
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
	proto.RegisterFile("proto/users.proto", fileDescriptor_b1c161a4c7514913)
}

var fileDescriptor_b1c161a4c7514913 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0xaf, 0xd2, 0x40,
	0x14, 0xcd, 0xb4, 0xa5, 0xca, 0x85, 0xa0, 0x9d, 0xa8, 0xa9, 0x75, 0x41, 0x9d, 0x55, 0xe3, 0xa2,
	0x24, 0x88, 0x0b, 0x63, 0xd8, 0xa8, 0x09, 0x71, 0x5b, 0x65, 0xc3, 0x86, 0x14, 0xe7, 0x4a, 0x1a,
	0x81, 0xd6, 0x99, 0xc1, 0xbd, 0x3f, 0xc1, 0x7f, 0xfc, 0xd2, 0x99, 0x69, 0xd3, 0x47, 0x21, 0x79,
	0x79, 0xbb, 0xfb, 0x71, 0xee, 0x3d, 0x67, 0xee, 0x69, 0x21, 0xa8, 0x44, 0xa9, 0xca, 0xd9, 0x59,
	0xa2, 0x90, 0xa9, 0x8e, 0xe9, 0x40, 0x27, 0xd1, 0x74, 0x5f, 0x96, 0xfb, 0x03, 0xce, 0x74, 0x71,
	0x77, 0xfe, 0x35, 0x53, 0xc5, 0x11, 0xa5, 0xca, 0x8f, 0x95, 0xc1, 0xb1, 0xff, 0x04, 0xbc, 0xb5,
	0x44, 0x41, 0x27, 0xe0, 0x14, 0x3c, 0x24, 0x31, 0x49, 0xdc, 0xcc, 0x29, 0x38, 0xa5, 0xe0, 0x9d,
	0xf2, 0x23, 0x86, 0x4e, 0x4c, 0x92, 0x61, 0xa6, 0x63, 0x1a, 0xc2, 0x93, 0x9c, 0x73, 0x81, 0x52,
	0x86, 0xae, 0x2e, 0x37, 0x29, 0x4d, 0xc1, 0xe3, 0xb9, 0xc2, 0xd0, 0x8b, 0x49, 0x32, 0x9a, 0x47,
	0xa9, 0xa1, 0x4d, 0x1b, 0xda, 0xf4, 0x47, 0x43, 0x9b, 0x69, 0x1c, 0x7d, 0x03, 0xc3, 0x52, 0x70,
	0x14, 0xdb, 0x82, 0xcb, 0x70, 0x10, 0xbb, 0x89, 0x9b, 0x3d, 0xd5, 0x85, 0x6f, 0x5c, 0xb2, 0x35,
	0x04, 0x5f, 0x04, 0xe6, 0x0a, 0x6b, 0x61, 0x19, 0xfe, 0x39, 0xa3, 0x54, 0xad, 0x1e, 0xd2, 0xd1,
	0x43, 0xc1, 0xab, 0x72, 0x29, 0x1b, 0x8d, 0x75, 0x7c, 0x5b, 0x23, 0x7b, 0x07, 0xb4, 0xbb, 0x56,
	0x56, 0xe5, 0x49, 0x22, 0x7d, 0x01, 0x03, 0x55, 0xfe, 0xc6, 0x93, 0x5d, 0x6c, 0x12, 0xc6, 0x60,
	0xb2, 0x42, 0xd5, 0xe5, 0x7f, 0x0e, 0x6e, 0xad, 0x95, 0x68, 0xad, 0x75, 0xc8, 0x16, 0xf0, 0xac,
	0xc5, 0xd8, 0x65, 0x6f, 0xc1, 0xdc, 0x5d, 0xc3, 0x46, 0xf3, 0x51, 0x6a, 0x2c, 0xd1, 0x18, 0xd3,
	0x61, 0x1b, 0x08, 0xd6, 0x15, 0xbf, 0x78, 0xdc, 0x55, 0x11, 0xd6, 0x12, 0xa7, 0xb5, 0x64, 0x0a,
	0x5e, 0xbd, 0x43, 0xbf, 0xeb, 0x62, 0xb9, 0x6e, 0xb0, 0x0f, 0x40, 0xbb, 0xbb, 0xad, 0xa8, 0x66,
	0x8c, 0xdc, 0x1a, 0xfb, 0x08, 0xc1, 0x57, 0x3c, 0xe0, 0x23, 0x24, 0xd5, 0x8c, 0xdd, 0xd1, 0x07,
	0x32, 0xce, 0xff, 0x39, 0x30, 0xae, 0x53, 0xf9, 0x1d, 0xc5, 0xdf, 0xe2, 0x27, 0xd2, 0x25, 0xf8,
	0xc6, 0x1b, 0x1a, 0x5a, 0x74, 0xef, 0x0b, 0x88, 0x5e, 0x5f, 0xe9, 0x58, 0xc2, 0x05, 0xb8, 0x2b,
	0x54, 0xf4, 0xa5, 0x45, 0xdc, 0xb7, 0x2e, 0x7a, 0x75, 0x59, 0xb6, 0x53, 0x4b, 0xf0, 0xcd, 0xb9,
	0x5a, 0xd2, 0x9e, 0x33, 0x2d, 0xe9, 0x95, 0xbb, 0x2e, 0xc1, 0x37, 0x6f, 0x6f, 0xc7, 0x7b, 0x57,
	0x6c, 0xc7, 0xfb, 0x47, 0xfa, 0x3c, 0xde, 0x40, 0xfa, 0x49, 0x77, 0xb7, 0xd5, 0x6e, 0xe7, 0xeb,
	0x5f, 0xe5, 0xfd, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0xd7, 0xbb, 0xce, 0xcb, 0x03, 0x00,
	0x00,
}
