// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/ccnu/ccnu.proto

package ccnuv1

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

const (
	CCNUService_Login_FullMethodName      = "/ccnu.v1.CCNUService/Login"
	CCNUService_CourseList_FullMethodName = "/ccnu.v1.CCNUService/CourseList"
)

// CCNUServiceClient is the client API for CCNUService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CCNUServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CourseList(ctx context.Context, in *CourseListRequest, opts ...grpc.CallOption) (*CourseListResponse, error)
}

type cCNUServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCCNUServiceClient(cc grpc.ClientConnInterface) CCNUServiceClient {
	return &cCNUServiceClient{cc}
}

func (c *cCNUServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, CCNUService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cCNUServiceClient) CourseList(ctx context.Context, in *CourseListRequest, opts ...grpc.CallOption) (*CourseListResponse, error) {
	out := new(CourseListResponse)
	err := c.cc.Invoke(ctx, CCNUService_CourseList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CCNUServiceServer is the server API for CCNUService service.
// All implementations must embed UnimplementedCCNUServiceServer
// for forward compatibility
type CCNUServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CourseList(context.Context, *CourseListRequest) (*CourseListResponse, error)
	mustEmbedUnimplementedCCNUServiceServer()
}

// UnimplementedCCNUServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCCNUServiceServer struct {
}

func (UnimplementedCCNUServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCCNUServiceServer) CourseList(context.Context, *CourseListRequest) (*CourseListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CourseList not implemented")
}
func (UnimplementedCCNUServiceServer) mustEmbedUnimplementedCCNUServiceServer() {}

// UnsafeCCNUServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CCNUServiceServer will
// result in compilation errors.
type UnsafeCCNUServiceServer interface {
	mustEmbedUnimplementedCCNUServiceServer()
}

func RegisterCCNUServiceServer(s grpc.ServiceRegistrar, srv CCNUServiceServer) {
	s.RegisterService(&CCNUService_ServiceDesc, srv)
}

func _CCNUService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCNUServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CCNUService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCNUServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CCNUService_CourseList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CCNUServiceServer).CourseList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CCNUService_CourseList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CCNUServiceServer).CourseList(ctx, req.(*CourseListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CCNUService_ServiceDesc is the grpc.ServiceDesc for CCNUService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CCNUService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ccnu.v1.CCNUService",
	HandlerType: (*CCNUServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _CCNUService_Login_Handler,
		},
		{
			MethodName: "CourseList",
			Handler:    _CCNUService_CourseList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ccnu/ccnu.proto",
}