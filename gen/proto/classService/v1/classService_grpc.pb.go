// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/classService/v1/classService.proto

package classServicev1

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
	ClassService_SearchClass_FullMethodName        = "/classService.v1.ClassService/SearchClass"
	ClassService_AddClass_FullMethodName           = "/classService.v1.ClassService/AddClass"
	ClassService_QueryFreeClassroom_FullMethodName = "/classService.v1.ClassService/QueryFreeClassroom"
)

// ClassServiceClient is the client API for ClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The greeting service definition.
type ClassServiceClient interface {
	// 数据源是所有使用匣子的用户的课表，从其中搜索相应的课程
	SearchClass(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
	//添加课程
	AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*AddClassReply, error)
	QueryFreeClassroom(ctx context.Context, in *QueryFreeClassroomReq, opts ...grpc.CallOption) (*QueryFreeClassroomResp, error)
}

type classServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClassServiceClient(cc grpc.ClientConnInterface) ClassServiceClient {
	return &classServiceClient{cc}
}

func (c *classServiceClient) SearchClass(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, ClassService_SearchClass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) AddClass(ctx context.Context, in *AddClassRequest, opts ...grpc.CallOption) (*AddClassReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddClassReply)
	err := c.cc.Invoke(ctx, ClassService_AddClass_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *classServiceClient) QueryFreeClassroom(ctx context.Context, in *QueryFreeClassroomReq, opts ...grpc.CallOption) (*QueryFreeClassroomResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryFreeClassroomResp)
	err := c.cc.Invoke(ctx, ClassService_QueryFreeClassroom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassServiceServer is the server API for ClassService service.
// All implementations must embed UnimplementedClassServiceServer
// for forward compatibility.
//
// The greeting service definition.
type ClassServiceServer interface {
	// 数据源是所有使用匣子的用户的课表，从其中搜索相应的课程
	SearchClass(context.Context, *SearchRequest) (*SearchReply, error)
	//添加课程
	AddClass(context.Context, *AddClassRequest) (*AddClassReply, error)
	QueryFreeClassroom(context.Context, *QueryFreeClassroomReq) (*QueryFreeClassroomResp, error)
	mustEmbedUnimplementedClassServiceServer()
}

// UnimplementedClassServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClassServiceServer struct{}

func (UnimplementedClassServiceServer) SearchClass(context.Context, *SearchRequest) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchClass not implemented")
}
func (UnimplementedClassServiceServer) AddClass(context.Context, *AddClassRequest) (*AddClassReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClass not implemented")
}
func (UnimplementedClassServiceServer) QueryFreeClassroom(context.Context, *QueryFreeClassroomReq) (*QueryFreeClassroomResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFreeClassroom not implemented")
}
func (UnimplementedClassServiceServer) mustEmbedUnimplementedClassServiceServer() {}
func (UnimplementedClassServiceServer) testEmbeddedByValue()                      {}

// UnsafeClassServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassServiceServer will
// result in compilation errors.
type UnsafeClassServiceServer interface {
	mustEmbedUnimplementedClassServiceServer()
}

func RegisterClassServiceServer(s grpc.ServiceRegistrar, srv ClassServiceServer) {
	// If the following call pancis, it indicates UnimplementedClassServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClassService_ServiceDesc, srv)
}

func _ClassService_SearchClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).SearchClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_SearchClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).SearchClass(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_AddClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).AddClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_AddClass_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).AddClass(ctx, req.(*AddClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClassService_QueryFreeClassroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFreeClassroomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassServiceServer).QueryFreeClassroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClassService_QueryFreeClassroom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassServiceServer).QueryFreeClassroom(ctx, req.(*QueryFreeClassroomReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ClassService_ServiceDesc is the grpc.ServiceDesc for ClassService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClassService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "classService.v1.ClassService",
	HandlerType: (*ClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchClass",
			Handler:    _ClassService_SearchClass_Handler,
		},
		{
			MethodName: "AddClass",
			Handler:    _ClassService_AddClass_Handler,
		},
		{
			MethodName: "QueryFreeClassroom",
			Handler:    _ClassService_QueryFreeClassroom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/classService/v1/classService.proto",
}
