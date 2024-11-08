// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: main.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Register_Register_FullMethodName = "/Register/Register"
	Register_Login_FullMethodName    = "/Register/Login"
	Register_Auth_FullMethodName     = "/Register/Auth"
)

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterClient interface {
	Register(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error)
	Login(ctx context.Context, in *UserLogReq, opts ...grpc.CallOption) (*UserResp, error)
	Auth(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*User, error)
}

type registerClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterClient(cc grpc.ClientConnInterface) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) Register(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResp)
	err := c.cc.Invoke(ctx, Register_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) Login(ctx context.Context, in *UserLogReq, opts ...grpc.CallOption) (*UserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResp)
	err := c.cc.Invoke(ctx, Register_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) Auth(ctx context.Context, in *UserAuthReq, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, Register_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServer is the server API for Register service.
// All implementations must embed UnimplementedRegisterServer
// for forward compatibility.
type RegisterServer interface {
	Register(context.Context, *UserReq) (*UserResp, error)
	Login(context.Context, *UserLogReq) (*UserResp, error)
	Auth(context.Context, *UserAuthReq) (*User, error)
	mustEmbedUnimplementedRegisterServer()
}

// UnimplementedRegisterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRegisterServer struct{}

func (UnimplementedRegisterServer) Register(context.Context, *UserReq) (*UserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegisterServer) Login(context.Context, *UserLogReq) (*UserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRegisterServer) Auth(context.Context, *UserAuthReq) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedRegisterServer) mustEmbedUnimplementedRegisterServer() {}
func (UnimplementedRegisterServer) testEmbeddedByValue()                  {}

// UnsafeRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServer will
// result in compilation errors.
type UnsafeRegisterServer interface {
	mustEmbedUnimplementedRegisterServer()
}

func RegisterRegisterServer(s grpc.ServiceRegistrar, srv RegisterServer) {
	// If the following call pancis, it indicates UnimplementedRegisterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Register_ServiceDesc, srv)
}

func _Register_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Register(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Login(ctx, req.(*UserLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Auth(ctx, req.(*UserAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Register_ServiceDesc is the grpc.ServiceDesc for Register service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Register_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Register_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Register_Login_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Register_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
