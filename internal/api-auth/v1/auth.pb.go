// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
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

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
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

func (m *CreateUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SigninRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigninRequest) Reset()         { *m = SigninRequest{} }
func (m *SigninRequest) String() string { return proto.CompactTextString(m) }
func (*SigninRequest) ProtoMessage()    {}
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *SigninRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigninRequest.Unmarshal(m, b)
}
func (m *SigninRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigninRequest.Marshal(b, m, deterministic)
}
func (m *SigninRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigninRequest.Merge(m, src)
}
func (m *SigninRequest) XXX_Size() int {
	return xxx_messageInfo_SigninRequest.Size(m)
}
func (m *SigninRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SigninRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SigninRequest proto.InternalMessageInfo

func (m *SigninRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *SigninRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SigninResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigninResponse) Reset()         { *m = SigninResponse{} }
func (m *SigninResponse) String() string { return proto.CompactTextString(m) }
func (*SigninResponse) ProtoMessage()    {}
func (*SigninResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *SigninResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigninResponse.Unmarshal(m, b)
}
func (m *SigninResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigninResponse.Marshal(b, m, deterministic)
}
func (m *SigninResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigninResponse.Merge(m, src)
}
func (m *SigninResponse) XXX_Size() int {
	return xxx_messageInfo_SigninResponse.Size(m)
}
func (m *SigninResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SigninResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SigninResponse proto.InternalMessageInfo

func (m *SigninResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ListUsersRequest struct {
	PageSize             int32    `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUsersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListUsersResponse struct {
	// See https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/monitoring/apiv3/metric_client.go#L137
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalPageSize        int32    `protobuf:"varint,3,opt,name=total_page_size,json=totalPageSize,proto3" json:"total_page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListUsersResponse) GetTotalPageSize() int32 {
	if m != nil {
		return m.TotalPageSize
	}
	return 0
}

type GetUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
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

func (m *GetUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
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

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	// or here https://github.com/gogo/grpc-example/commit/6c217371b67a89609c632f047477fa5a1123ac93
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{9}
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

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{10}
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

func (m *DeleteUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mail                 string   `protobuf:"bytes,2,opt,name=mail,proto3" json:"mail,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Login                string   `protobuf:"bytes,4,opt,name=login,proto3" json:"login,omitempty"`
	Firstname            string   `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Password             string   `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Role                 int32    `protobuf:"varint,8,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{11}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *User) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "flatsharing.auth.v1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "flatsharing.auth.v1.CreateUserResponse")
	proto.RegisterType((*SigninRequest)(nil), "flatsharing.auth.v1.SigninRequest")
	proto.RegisterType((*SigninResponse)(nil), "flatsharing.auth.v1.SigninResponse")
	proto.RegisterType((*LogoutRequest)(nil), "flatsharing.auth.v1.LogoutRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "flatsharing.auth.v1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "flatsharing.auth.v1.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "flatsharing.auth.v1.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "flatsharing.auth.v1.GetUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "flatsharing.auth.v1.UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "flatsharing.auth.v1.DeleteUserRequest")
	proto.RegisterType((*User)(nil), "flatsharing.auth.v1.User")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x4e, 0x13, 0x41,
	0x18, 0x4d, 0x4b, 0x5b, 0xe8, 0x87, 0xe5, 0x67, 0xa8, 0x58, 0x17, 0x48, 0xc8, 0x0a, 0xd5, 0x90,
	0xb0, 0x1b, 0xf0, 0x4e, 0x2f, 0x14, 0xfc, 0xbb, 0x41, 0x43, 0x8a, 0xdc, 0xe8, 0x45, 0x33, 0xd0,
	0xe9, 0x76, 0xd2, 0xed, 0xce, 0xba, 0x33, 0xad, 0x8a, 0x31, 0x1a, 0x5f, 0x41, 0xdf, 0xcc, 0x57,
	0xd0, 0xf7, 0x30, 0xf3, 0xcd, 0xf4, 0x8f, 0x6e, 0x21, 0x86, 0x9b, 0x66, 0x67, 0xe6, 0xcc, 0x39,
	0xdf, 0xcc, 0x77, 0xce, 0x14, 0x80, 0x76, 0x55, 0xcb, 0x8b, 0x13, 0xa1, 0x04, 0x59, 0x69, 0x86,
	0x54, 0xc9, 0x16, 0x4d, 0x78, 0x14, 0x78, 0x38, 0xdf, 0xdb, 0x73, 0xd6, 0x02, 0x21, 0x82, 0x90,
	0xf9, 0x08, 0x39, 0xeb, 0x36, 0x7d, 0xd6, 0x89, 0xd5, 0x67, 0xb3, 0xc3, 0xd9, 0xbc, 0xbc, 0xd8,
	0xe4, 0x2c, 0x6c, 0xd4, 0x3b, 0x54, 0xb6, 0x2d, 0x62, 0xdd, 0x22, 0x68, 0xcc, 0x7d, 0x1a, 0x45,
	0x42, 0x51, 0xc5, 0x45, 0x24, 0xcd, 0xaa, 0x7b, 0x08, 0xcb, 0xcf, 0x12, 0x46, 0x15, 0x3b, 0x95,
	0x2c, 0xa9, 0xb1, 0x0f, 0x5d, 0x26, 0x15, 0xd9, 0x85, 0x5c, 0x57, 0xb2, 0xa4, 0x92, 0xd9, 0xcc,
	0x3c, 0x98, 0xdf, 0xbf, 0xeb, 0xa5, 0x54, 0xe5, 0x21, 0x1e, 0x61, 0xee, 0x16, 0x90, 0x51, 0x0e,
	0x19, 0x8b, 0x48, 0x32, 0xb2, 0x00, 0x59, 0xde, 0x40, 0x8a, 0x62, 0x2d, 0xcb, 0x1b, 0xee, 0x01,
	0x94, 0x4e, 0x78, 0x10, 0xf1, 0xa8, 0xaf, 0x52, 0x86, 0x7c, 0x28, 0x02, 0x1e, 0x59, 0x8c, 0x19,
	0x10, 0x07, 0xe6, 0x62, 0x2a, 0xe5, 0x47, 0x91, 0x34, 0x2a, 0x59, 0x5c, 0x18, 0x8c, 0xdd, 0x2a,
	0x2c, 0xf4, 0x29, 0xac, 0x48, 0x19, 0xf2, 0x4a, 0xb4, 0xd9, 0x80, 0x03, 0x07, 0xee, 0x36, 0x94,
	0x8e, 0x44, 0x20, 0xba, 0x6a, 0x44, 0x2a, 0x05, 0xf6, 0x06, 0x96, 0x8e, 0xb8, 0x54, 0xba, 0x6a,
	0xd9, 0x47, 0xae, 0x41, 0x31, 0xa6, 0x01, 0xab, 0x4b, 0x7e, 0xc1, 0x10, 0x9d, 0xd7, 0xfa, 0x01,
	0x3b, 0xe1, 0x17, 0x8c, 0x6c, 0x00, 0xe0, 0xa2, 0xe1, 0x32, 0xd5, 0x21, 0xfc, 0x2d, 0xf2, 0xfd,
	0xca, 0xc0, 0xf2, 0x08, 0xa1, 0x2d, 0xd1, 0x87, 0xbc, 0xbe, 0x25, 0x59, 0xc9, 0x6c, 0xce, 0x5c,
	0x7d, 0x9b, 0x06, 0x47, 0xaa, 0xb0, 0x18, 0xb1, 0x4f, 0xaa, 0x3e, 0x21, 0x55, 0xd2, 0xd3, 0xc7,
	0x7d, 0x39, 0x8d, 0x53, 0x42, 0xd1, 0xb0, 0x3e, 0x2c, 0x78, 0x06, 0x0b, 0x2e, 0xe1, 0xf4, 0xb1,
	0xad, 0xda, 0x7d, 0x02, 0x0b, 0xaf, 0x98, 0xba, 0x41, 0x7f, 0x9f, 0xc2, 0xe2, 0x80, 0xc0, 0x1e,
	0xea, 0x3f, 0x19, 0xbe, 0xc1, 0xf2, 0x69, 0xdc, 0xb8, 0x91, 0xcb, 0xc8, 0x63, 0x98, 0xef, 0x22,
	0x07, 0x9a, 0x1b, 0xaf, 0x64, 0x7e, 0xdf, 0xf1, 0x8c, 0xbb, 0xbd, 0xbe, 0xff, 0xbd, 0x97, 0xda,
	0xff, 0xaf, 0xa9, 0x6c, 0xd7, 0xc0, 0xc0, 0xf5, 0xb7, 0xb6, 0xf9, 0x73, 0x16, 0xb2, 0x1b, 0xd9,
	0xfc, 0x6f, 0x06, 0x72, 0x7a, 0x78, 0xd9, 0xd9, 0x84, 0x40, 0xae, 0x43, 0x79, 0x68, 0xbb, 0x84,
	0xdf, 0xda, 0xc6, 0x7a, 0x53, 0x44, 0x3b, 0xa6, 0x2b, 0xc5, 0xda, 0x60, 0x3c, 0x34, 0x7e, 0x6e,
	0xd4, 0xf8, 0xeb, 0x50, 0x6c, 0xf2, 0x44, 0x2a, 0xdc, 0x92, 0x37, 0xde, 0x1a, 0x4c, 0x68, 0xbe,
	0x90, 0xda, 0xc5, 0x82, 0xe1, 0xeb, 0x8f, 0xc7, 0x22, 0x33, 0x3b, 0x1e, 0x19, 0x5d, 0x5b, 0x22,
	0x42, 0x56, 0x99, 0x43, 0x67, 0xe0, 0xb7, 0xb6, 0xf1, 0x39, 0xe6, 0xb5, 0x51, 0xa7, 0xaa, 0x52,
	0x34, 0x52, 0x76, 0xe6, 0x40, 0xed, 0x7f, 0x2f, 0xc0, 0xad, 0x83, 0xae, 0x6a, 0x9d, 0xb0, 0xa4,
	0xc7, 0xcf, 0x99, 0x24, 0x2d, 0x28, 0x98, 0xd8, 0x11, 0x37, 0xf5, 0x8e, 0xc6, 0x62, 0xed, 0xdc,
	0xbb, 0x12, 0x63, 0xfc, 0xe3, 0xde, 0xfe, 0xf1, 0xfb, 0xcf, 0xcf, 0xec, 0xa2, 0x0b, 0x7e, 0x6f,
	0xcf, 0x97, 0xb8, 0xf6, 0x28, 0xb3, 0x43, 0xde, 0x43, 0xc1, 0x04, 0x77, 0x8a, 0xd2, 0x58, 0xaa,
	0x9d, 0xd5, 0x89, 0xe6, 0xbf, 0xd0, 0x2f, 0xe3, 0x38, 0x79, 0x88, 0x5b, 0x34, 0x79, 0x02, 0x30,
	0x7c, 0xa6, 0x48, 0x35, 0x55, 0x60, 0xe2, 0x2d, 0x74, 0xee, 0x5f, 0x8b, 0xb3, 0x47, 0x5a, 0x41,
	0xd5, 0x92, 0x3b, 0xa7, 0x55, 0x75, 0xaf, 0xb5, 0x66, 0x08, 0xc5, 0xc1, 0x8b, 0x40, 0xb6, 0xd3,
	0xcf, 0x74, 0xe9, 0x09, 0x72, 0xaa, 0xd7, 0xc1, 0xac, 0xe0, 0x12, 0x0a, 0x02, 0x19, 0x08, 0x92,
	0x18, 0x66, 0x6d, 0x50, 0x49, 0x7a, 0x17, 0xc6, 0xdf, 0x01, 0x67, 0xeb, 0x6a, 0x90, 0xd5, 0x71,
	0x50, 0xa7, 0x4c, 0x48, 0x5f, 0xc7, 0xff, 0xa2, 0x7f, 0x3d, 0xde, 0xf8, 0x4a, 0xda, 0x00, 0xc3,
	0x60, 0x4f, 0xb9, 0xd3, 0x89, 0xe4, 0x4f, 0x6d, 0xdc, 0x06, 0x2a, 0xdd, 0x71, 0x52, 0x94, 0xf4,
	0x65, 0xb6, 0x00, 0x86, 0x21, 0x9e, 0x22, 0x36, 0x91, 0xf2, 0xa9, 0x62, 0xf6, 0x58, 0x3b, 0x29,
	0x62, 0x87, 0xab, 0xef, 0xca, 0x3c, 0x52, 0x3a, 0xae, 0xa1, 0xfe, 0xdf, 0xdc, 0xd5, 0xec, 0x7e,
	0x6f, 0xef, 0xac, 0x80, 0x1c, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x85, 0x6d, 0x7a, 0xda,
	0xb4, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServicesClient is the client API for AuthServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServicesClient interface {
	Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServicesClient struct {
	cc *grpc.ClientConn
}

func NewAuthServicesClient(cc *grpc.ClientConn) AuthServicesClient {
	return &authServicesClient{cc}
}

func (c *authServicesClient) Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninResponse, error) {
	out := new(SigninResponse)
	err := c.cc.Invoke(ctx, "/flatsharing.auth.v1.AuthServices/Signin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicesClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/flatsharing.auth.v1.AuthServices/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicesClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/flatsharing.auth.v1.AuthServices/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicesClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/flatsharing.auth.v1.AuthServices/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicesClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/flatsharing.auth.v1.AuthServices/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicesClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/flatsharing.auth.v1.AuthServices/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServicesClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/flatsharing.auth.v1.AuthServices/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServicesServer is the server API for AuthServices service.
type AuthServicesServer interface {
	Signin(context.Context, *SigninRequest) (*SigninResponse, error)
	Logout(context.Context, *LogoutRequest) (*empty.Empty, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
}

func RegisterAuthServicesServer(s *grpc.Server, srv AuthServicesServer) {
	s.RegisterService(&_AuthServices_serviceDesc, srv)
}

func _AuthServices_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigninRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicesServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flatsharing.auth.v1.AuthServices/Signin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicesServer).Signin(ctx, req.(*SigninRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServices_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicesServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flatsharing.auth.v1.AuthServices/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicesServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServices_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicesServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flatsharing.auth.v1.AuthServices/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicesServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServices_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicesServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flatsharing.auth.v1.AuthServices/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicesServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServices_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicesServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flatsharing.auth.v1.AuthServices/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicesServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServices_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicesServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flatsharing.auth.v1.AuthServices/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicesServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServices_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServicesServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flatsharing.auth.v1.AuthServices/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServicesServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flatsharing.auth.v1.AuthServices",
	HandlerType: (*AuthServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signin",
			Handler:    _AuthServices_Signin_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthServices_Logout_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AuthServices_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _AuthServices_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthServices_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthServices_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthServices_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}