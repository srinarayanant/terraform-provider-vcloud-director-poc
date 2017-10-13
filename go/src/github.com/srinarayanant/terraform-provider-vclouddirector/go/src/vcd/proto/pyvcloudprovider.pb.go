// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pyvcloudprovider.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	pyvcloudprovider.proto

It has these top-level messages:
	TenantCredentials
	LoginResult
	Catalog
	IsPresentCatalogResult
	CreateCatalogResult
	DeleteCatalogResult
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Tenant VCD crendentials
type TenantCredentials struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Org      string `protobuf:"bytes,3,opt,name=org" json:"org,omitempty"`
	Ip       string `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
}

func (m *TenantCredentials) Reset()                    { *m = TenantCredentials{} }
func (m *TenantCredentials) String() string            { return proto1.CompactTextString(m) }
func (*TenantCredentials) ProtoMessage()               {}
func (*TenantCredentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TenantCredentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *TenantCredentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *TenantCredentials) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

func (m *TenantCredentials) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type LoginResult struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResult) Reset()                    { *m = LoginResult{} }
func (m *LoginResult) String() string            { return proto1.CompactTextString(m) }
func (*LoginResult) ProtoMessage()               {}
func (*LoginResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResult) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Catalog struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Shared      bool   `protobuf:"varint,3,opt,name=shared" json:"shared,omitempty"`
}

func (m *Catalog) Reset()                    { *m = Catalog{} }
func (m *Catalog) String() string            { return proto1.CompactTextString(m) }
func (*Catalog) ProtoMessage()               {}
func (*Catalog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Catalog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Catalog) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Catalog) GetShared() bool {
	if m != nil {
		return m.Shared
	}
	return false
}

type IsPresentCatalogResult struct {
	Present bool `protobuf:"varint,1,opt,name=present" json:"present,omitempty"`
}

func (m *IsPresentCatalogResult) Reset()                    { *m = IsPresentCatalogResult{} }
func (m *IsPresentCatalogResult) String() string            { return proto1.CompactTextString(m) }
func (*IsPresentCatalogResult) ProtoMessage()               {}
func (*IsPresentCatalogResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsPresentCatalogResult) GetPresent() bool {
	if m != nil {
		return m.Present
	}
	return false
}

type CreateCatalogResult struct {
	Created bool `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
}

func (m *CreateCatalogResult) Reset()                    { *m = CreateCatalogResult{} }
func (m *CreateCatalogResult) String() string            { return proto1.CompactTextString(m) }
func (*CreateCatalogResult) ProtoMessage()               {}
func (*CreateCatalogResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateCatalogResult) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

type DeleteCatalogResult struct {
	Deleted bool `protobuf:"varint,1,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteCatalogResult) Reset()                    { *m = DeleteCatalogResult{} }
func (m *DeleteCatalogResult) String() string            { return proto1.CompactTextString(m) }
func (*DeleteCatalogResult) ProtoMessage()               {}
func (*DeleteCatalogResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteCatalogResult) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func init() {
	proto1.RegisterType((*TenantCredentials)(nil), "proto.TenantCredentials")
	proto1.RegisterType((*LoginResult)(nil), "proto.LoginResult")
	proto1.RegisterType((*Catalog)(nil), "proto.Catalog")
	proto1.RegisterType((*IsPresentCatalogResult)(nil), "proto.IsPresentCatalogResult")
	proto1.RegisterType((*CreateCatalogResult)(nil), "proto.CreateCatalogResult")
	proto1.RegisterType((*DeleteCatalogResult)(nil), "proto.DeleteCatalogResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PyVcloudProvider service

type PyVcloudProviderClient interface {
	// Tenant Loging to VCD
	Login(ctx context.Context, in *TenantCredentials, opts ...grpc.CallOption) (*LoginResult, error)
	IsPresentCatalog(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*IsPresentCatalogResult, error)
	CreateCatalog(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*CreateCatalogResult, error)
	DeleteCatalog(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*DeleteCatalogResult, error)
}

type pyVcloudProviderClient struct {
	cc *grpc.ClientConn
}

func NewPyVcloudProviderClient(cc *grpc.ClientConn) PyVcloudProviderClient {
	return &pyVcloudProviderClient{cc}
}

func (c *pyVcloudProviderClient) Login(ctx context.Context, in *TenantCredentials, opts ...grpc.CallOption) (*LoginResult, error) {
	out := new(LoginResult)
	err := grpc.Invoke(ctx, "/proto.PyVcloudProvider/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pyVcloudProviderClient) IsPresentCatalog(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*IsPresentCatalogResult, error) {
	out := new(IsPresentCatalogResult)
	err := grpc.Invoke(ctx, "/proto.PyVcloudProvider/isPresentCatalog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pyVcloudProviderClient) CreateCatalog(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*CreateCatalogResult, error) {
	out := new(CreateCatalogResult)
	err := grpc.Invoke(ctx, "/proto.PyVcloudProvider/CreateCatalog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pyVcloudProviderClient) DeleteCatalog(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*DeleteCatalogResult, error) {
	out := new(DeleteCatalogResult)
	err := grpc.Invoke(ctx, "/proto.PyVcloudProvider/DeleteCatalog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PyVcloudProvider service

type PyVcloudProviderServer interface {
	// Tenant Loging to VCD
	Login(context.Context, *TenantCredentials) (*LoginResult, error)
	IsPresentCatalog(context.Context, *Catalog) (*IsPresentCatalogResult, error)
	CreateCatalog(context.Context, *Catalog) (*CreateCatalogResult, error)
	DeleteCatalog(context.Context, *Catalog) (*DeleteCatalogResult, error)
}

func RegisterPyVcloudProviderServer(s *grpc.Server, srv PyVcloudProviderServer) {
	s.RegisterService(&_PyVcloudProvider_serviceDesc, srv)
}

func _PyVcloudProvider_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyVcloudProviderServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PyVcloudProvider/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyVcloudProviderServer).Login(ctx, req.(*TenantCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _PyVcloudProvider_IsPresentCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyVcloudProviderServer).IsPresentCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PyVcloudProvider/IsPresentCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyVcloudProviderServer).IsPresentCatalog(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _PyVcloudProvider_CreateCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyVcloudProviderServer).CreateCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PyVcloudProvider/CreateCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyVcloudProviderServer).CreateCatalog(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

func _PyVcloudProvider_DeleteCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Catalog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyVcloudProviderServer).DeleteCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PyVcloudProvider/DeleteCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyVcloudProviderServer).DeleteCatalog(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

var _PyVcloudProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PyVcloudProvider",
	HandlerType: (*PyVcloudProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _PyVcloudProvider_Login_Handler,
		},
		{
			MethodName: "isPresentCatalog",
			Handler:    _PyVcloudProvider_IsPresentCatalog_Handler,
		},
		{
			MethodName: "CreateCatalog",
			Handler:    _PyVcloudProvider_CreateCatalog_Handler,
		},
		{
			MethodName: "DeleteCatalog",
			Handler:    _PyVcloudProvider_DeleteCatalog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pyvcloudprovider.proto",
}

func init() { proto1.RegisterFile("pyvcloudprovider.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0xcf, 0xd3, 0x30,
	0x10, 0xc6, 0xdb, 0xbc, 0xfd, 0xc7, 0x55, 0x54, 0xc5, 0x40, 0x15, 0x45, 0x20, 0x55, 0x66, 0x61,
	0x0a, 0x52, 0x19, 0x3a, 0xb1, 0x34, 0x2c, 0x48, 0x0c, 0x51, 0x84, 0x60, 0x36, 0xf1, 0xa9, 0x58,
	0x24, 0xb6, 0x65, 0xbb, 0xad, 0xfa, 0x25, 0xf8, 0xcc, 0x28, 0x8e, 0x53, 0xb5, 0x4d, 0xdf, 0x29,
	0xfe, 0x3d, 0xce, 0xe3, 0x3b, 0x3f, 0x67, 0x58, 0xe9, 0xf3, 0xb1, 0xac, 0xd4, 0x81, 0x6b, 0xa3,
	0x8e, 0x82, 0xa3, 0x49, 0xb5, 0x51, 0x4e, 0x91, 0xb1, 0xff, 0xd0, 0x1a, 0x5e, 0xfd, 0x40, 0xc9,
	0xa4, 0xcb, 0x0c, 0x72, 0x94, 0x4e, 0xb0, 0xca, 0x92, 0x04, 0x66, 0x07, 0x8b, 0x46, 0xb2, 0x1a,
	0xe3, 0xe1, 0x7a, 0xf8, 0xf1, 0x45, 0x71, 0xe1, 0x66, 0x4f, 0x33, 0x6b, 0x4f, 0xca, 0xf0, 0x38,
	0x6a, 0xf7, 0x3a, 0x26, 0x4b, 0x78, 0x52, 0x66, 0x1f, 0x3f, 0x79, 0xb9, 0x59, 0x92, 0x05, 0x44,
	0x42, 0xc7, 0x23, 0x2f, 0x44, 0x42, 0xd3, 0x0f, 0x30, 0xff, 0xae, 0xf6, 0x42, 0x16, 0x68, 0x0f,
	0x95, 0x23, 0x6f, 0x60, 0xec, 0xd4, 0x5f, 0x94, 0xa1, 0x4a, 0x0b, 0xf4, 0x17, 0x4c, 0x33, 0xe6,
	0x58, 0xa5, 0xf6, 0x84, 0xc0, 0xe8, 0xaa, 0x0b, 0xbf, 0x26, 0x6b, 0x98, 0x73, 0xb4, 0xa5, 0x11,
	0xda, 0x09, 0x25, 0x43, 0x13, 0xd7, 0x12, 0x59, 0xc1, 0xc4, 0xfe, 0x61, 0x06, 0xb9, 0x6f, 0x65,
	0x56, 0x04, 0xa2, 0x1b, 0x58, 0x7d, 0xb3, 0xb9, 0x41, 0x8b, 0xd2, 0x85, 0x0a, 0xa1, 0x91, 0x18,
	0xa6, 0xba, 0xd5, 0x7d, 0xa9, 0x59, 0xd1, 0x21, 0xfd, 0x04, 0xaf, 0x33, 0x83, 0xcc, 0x61, 0xcf,
	0x50, 0x7a, 0x99, 0x77, 0x86, 0x80, 0x8d, 0xe1, 0x2b, 0x56, 0xf8, 0xc0, 0xc0, 0xbd, 0x7c, 0x31,
	0x04, 0xdc, 0xfc, 0x8b, 0x60, 0x99, 0x9f, 0x7f, 0xfa, 0x21, 0xe5, 0x61, 0x48, 0x64, 0x0b, 0x63,
	0x1f, 0x14, 0x89, 0xdb, 0x79, 0xa5, 0xbd, 0x29, 0x25, 0x24, 0xec, 0x5c, 0x05, 0x4a, 0x07, 0x24,
	0x83, 0xa5, 0xb8, 0xbb, 0x23, 0x59, 0x84, 0x3f, 0x03, 0x27, 0xef, 0x03, 0x3f, 0x0e, 0x83, 0x0e,
	0xc8, 0x17, 0x78, 0x79, 0x73, 0xe9, 0xde, 0x09, 0x49, 0xc7, 0xfd, 0x68, 0x5a, 0xfb, 0x4d, 0x04,
	0xcf, 0xda, 0x1f, 0x04, 0x45, 0x07, 0xbb, 0x2d, 0xbc, 0x2b, 0x55, 0x9d, 0x1e, 0xeb, 0x13, 0x33,
	0x98, 0x76, 0xef, 0x37, 0xed, 0x1e, 0xf0, 0xee, 0xed, 0x7d, 0x5a, 0x79, 0x73, 0x58, 0x3e, 0xfc,
	0x3d, 0xf1, 0xa7, 0x7e, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x36, 0x56, 0xf9, 0xc5, 0xf4, 0x02,
	0x00, 0x00,
}