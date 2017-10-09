// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pyvcloudprovider.proto

/*
Package pyvcloudprovider is a generated protocol buffer package.

It is generated from these files:
	pyvcloudprovider.proto

It has these top-level messages:
	TenantCredentials
	LoginResult
	Catalog
	IsPresentCatalogResult
*/
package pyvcloudprovider

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Tenant VCD crendentials
type TenantCredentials struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Org      string `protobuf:"bytes,3,opt,name=org" json:"org,omitempty"`
	Ip       string `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
}

func (m *TenantCredentials) Reset()                    { *m = TenantCredentials{} }
func (m *TenantCredentials) String() string            { return proto.CompactTextString(m) }
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
func (m *LoginResult) String() string            { return proto.CompactTextString(m) }
func (*LoginResult) ProtoMessage()               {}
func (*LoginResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResult) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Catalog struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Catalog) Reset()                    { *m = Catalog{} }
func (m *Catalog) String() string            { return proto.CompactTextString(m) }
func (*Catalog) ProtoMessage()               {}
func (*Catalog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Catalog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IsPresentCatalogResult struct {
	IsPresent bool `protobuf:"varint,1,opt,name=isPresent" json:"isPresent,omitempty"`
}

func (m *IsPresentCatalogResult) Reset()                    { *m = IsPresentCatalogResult{} }
func (m *IsPresentCatalogResult) String() string            { return proto.CompactTextString(m) }
func (*IsPresentCatalogResult) ProtoMessage()               {}
func (*IsPresentCatalogResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsPresentCatalogResult) GetIsPresent() bool {
	if m != nil {
		return m.IsPresent
	}
	return false
}

func init() {
	proto.RegisterType((*TenantCredentials)(nil), "pyvcloudprovider.TenantCredentials")
	proto.RegisterType((*LoginResult)(nil), "pyvcloudprovider.LoginResult")
	proto.RegisterType((*Catalog)(nil), "pyvcloudprovider.Catalog")
	proto.RegisterType((*IsPresentCatalogResult)(nil), "pyvcloudprovider.IsPresentCatalogResult")
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
}

type pyVcloudProviderClient struct {
	cc *grpc.ClientConn
}

func NewPyVcloudProviderClient(cc *grpc.ClientConn) PyVcloudProviderClient {
	return &pyVcloudProviderClient{cc}
}

func (c *pyVcloudProviderClient) Login(ctx context.Context, in *TenantCredentials, opts ...grpc.CallOption) (*LoginResult, error) {
	out := new(LoginResult)
	err := grpc.Invoke(ctx, "/pyvcloudprovider.PyVcloudProvider/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pyVcloudProviderClient) IsPresentCatalog(ctx context.Context, in *Catalog, opts ...grpc.CallOption) (*IsPresentCatalogResult, error) {
	out := new(IsPresentCatalogResult)
	err := grpc.Invoke(ctx, "/pyvcloudprovider.PyVcloudProvider/isPresentCatalog", in, out, c.cc, opts...)
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
		FullMethod: "/pyvcloudprovider.PyVcloudProvider/Login",
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
		FullMethod: "/pyvcloudprovider.PyVcloudProvider/IsPresentCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyVcloudProviderServer).IsPresentCatalog(ctx, req.(*Catalog))
	}
	return interceptor(ctx, in, info, handler)
}

var _PyVcloudProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pyvcloudprovider.PyVcloudProvider",
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pyvcloudprovider.proto",
}

func init() { proto.RegisterFile("pyvcloudprovider.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x6d, 0xfa, 0xe7, 0xf7, 0x6b, 0x47, 0x90, 0xb8, 0x68, 0x89, 0xa1, 0x05, 0xd9, 0x5e, 0x7a,
	0xca, 0x41, 0x41, 0xef, 0xed, 0x49, 0x50, 0x08, 0x41, 0x04, 0x8f, 0x6b, 0x33, 0x84, 0xc5, 0x64,
	0x77, 0xd9, 0xdd, 0xa4, 0xf4, 0xbb, 0xf9, 0xe1, 0x24, 0xdb, 0xa4, 0x96, 0xc4, 0xdb, 0xcc, 0x7b,
	0xf3, 0x76, 0xdf, 0x9b, 0x81, 0xb9, 0x3a, 0x54, 0xbb, 0x5c, 0x96, 0xa9, 0xd2, 0xb2, 0xe2, 0x29,
	0xea, 0x48, 0x69, 0x69, 0x25, 0xf1, 0xbb, 0x38, 0x2d, 0xe0, 0xea, 0x0d, 0x05, 0x13, 0x76, 0xab,
	0x31, 0x45, 0x61, 0x39, 0xcb, 0x0d, 0x09, 0x61, 0x5a, 0x1a, 0xd4, 0x82, 0x15, 0x18, 0x78, 0x77,
	0xde, 0x7a, 0x96, 0x9c, 0xfa, 0x9a, 0x53, 0xcc, 0x98, 0xbd, 0xd4, 0x69, 0x30, 0x3c, 0x72, 0x6d,
	0x4f, 0x7c, 0x18, 0x49, 0x9d, 0x05, 0x23, 0x07, 0xd7, 0x25, 0xb9, 0x84, 0x21, 0x57, 0xc1, 0xd8,
	0x01, 0x43, 0xae, 0xe8, 0x0a, 0x2e, 0x5e, 0x64, 0xc6, 0x45, 0x82, 0xa6, 0xcc, 0x2d, 0xb9, 0x86,
	0x89, 0x95, 0x5f, 0x28, 0x9a, 0x5f, 0x8e, 0x0d, 0x5d, 0xc2, 0xff, 0x2d, 0xb3, 0x2c, 0x97, 0x19,
	0x21, 0x30, 0x3e, 0x73, 0xe1, 0x6a, 0xfa, 0x08, 0xf3, 0x67, 0x13, 0x6b, 0x34, 0x28, 0x6c, 0x33,
	0xd7, 0x3c, 0xb7, 0x80, 0x19, 0x6f, 0x19, 0x27, 0x99, 0x26, 0xbf, 0xc0, 0xfd, 0xb7, 0x07, 0x7e,
	0x7c, 0x78, 0x77, 0xf9, 0xe3, 0x26, 0x3f, 0x79, 0x85, 0x89, 0x33, 0x44, 0x56, 0x51, 0x6f, 0x67,
	0xbd, 0xc5, 0x84, 0xcb, 0xfe, 0xd0, 0x59, 0x1c, 0x3a, 0x20, 0x1f, 0xe0, 0xf3, 0x8e, 0x37, 0x72,
	0xdb, 0x17, 0x35, 0x54, 0xb8, 0xee, 0x53, 0x7f, 0x47, 0xa3, 0x83, 0xcd, 0x13, 0x2c, 0x76, 0xb2,
	0x88, 0xaa, 0x62, 0xcf, 0x34, 0x9e, 0x74, 0x51, 0x2b, 0xdc, 0xdc, 0x74, 0xb3, 0xc5, 0xf5, 0xc9,
	0x63, 0xef, 0xf3, 0x9f, 0xbb, 0xfd, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xa8, 0x08,
	0xed, 0x15, 0x02, 0x00, 0x00,
}
