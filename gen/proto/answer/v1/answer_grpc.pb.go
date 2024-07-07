// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: proto/answer/v1/answer.proto

package answerv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AnswerService_Publish_FullMethodName          = "/answer.v1.AnswerService/Publish"
	AnswerService_Detail_FullMethodName           = "/answer.v1.AnswerService/Detail"
	AnswerService_ListForQuestion_FullMethodName  = "/answer.v1.AnswerService/ListForQuestion"
	AnswerService_ListForUser_FullMethodName      = "/answer.v1.AnswerService/ListForUser"
	AnswerService_CountForQuestion_FullMethodName = "/answer.v1.AnswerService/CountForQuestion"
	AnswerService_DelAnswerById_FullMethodName    = "/answer.v1.AnswerService/DelAnswerById"
)

// AnswerServiceClient is the client API for AnswerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnswerServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
	ListForQuestion(ctx context.Context, in *ListForQuestionRequest, opts ...grpc.CallOption) (*ListForQuestionResponse, error)
	ListForUser(ctx context.Context, in *ListForUserRequest, opts ...grpc.CallOption) (*ListForUserResponse, error)
	CountForQuestion(ctx context.Context, in *CountForQuestionRequest, opts ...grpc.CallOption) (*CountForQuestionResponse, error)
	DelAnswerById(ctx context.Context, in *DelAnswerByIdRequest, opts ...grpc.CallOption) (*DelAnswerByIdResponse, error)
}

type answerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnswerServiceClient(cc grpc.ClientConnInterface) AnswerServiceClient {
	return &answerServiceClient{cc}
}

func (c *answerServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, AnswerService_Publish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerServiceClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, AnswerService_Detail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerServiceClient) ListForQuestion(ctx context.Context, in *ListForQuestionRequest, opts ...grpc.CallOption) (*ListForQuestionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListForQuestionResponse)
	err := c.cc.Invoke(ctx, AnswerService_ListForQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerServiceClient) ListForUser(ctx context.Context, in *ListForUserRequest, opts ...grpc.CallOption) (*ListForUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListForUserResponse)
	err := c.cc.Invoke(ctx, AnswerService_ListForUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerServiceClient) CountForQuestion(ctx context.Context, in *CountForQuestionRequest, opts ...grpc.CallOption) (*CountForQuestionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountForQuestionResponse)
	err := c.cc.Invoke(ctx, AnswerService_CountForQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerServiceClient) DelAnswerById(ctx context.Context, in *DelAnswerByIdRequest, opts ...grpc.CallOption) (*DelAnswerByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelAnswerByIdResponse)
	err := c.cc.Invoke(ctx, AnswerService_DelAnswerById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnswerServiceServer is the server API for AnswerService service.
// All implementations must embed UnimplementedAnswerServiceServer
// for forward compatibility
type AnswerServiceServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
	ListForQuestion(context.Context, *ListForQuestionRequest) (*ListForQuestionResponse, error)
	ListForUser(context.Context, *ListForUserRequest) (*ListForUserResponse, error)
	CountForQuestion(context.Context, *CountForQuestionRequest) (*CountForQuestionResponse, error)
	DelAnswerById(context.Context, *DelAnswerByIdRequest) (*DelAnswerByIdResponse, error)
	mustEmbedUnimplementedAnswerServiceServer()
}

// UnimplementedAnswerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnswerServiceServer struct {
}

func (UnimplementedAnswerServiceServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedAnswerServiceServer) Detail(context.Context, *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedAnswerServiceServer) ListForQuestion(context.Context, *ListForQuestionRequest) (*ListForQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForQuestion not implemented")
}
func (UnimplementedAnswerServiceServer) ListForUser(context.Context, *ListForUserRequest) (*ListForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForUser not implemented")
}
func (UnimplementedAnswerServiceServer) CountForQuestion(context.Context, *CountForQuestionRequest) (*CountForQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountForQuestion not implemented")
}
func (UnimplementedAnswerServiceServer) DelAnswerById(context.Context, *DelAnswerByIdRequest) (*DelAnswerByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAnswerById not implemented")
}
func (UnimplementedAnswerServiceServer) mustEmbedUnimplementedAnswerServiceServer() {}

// UnsafeAnswerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnswerServiceServer will
// result in compilation errors.
type UnsafeAnswerServiceServer interface {
	mustEmbedUnimplementedAnswerServiceServer()
}

func RegisterAnswerServiceServer(s grpc.ServiceRegistrar, srv AnswerServiceServer) {
	s.RegisterService(&AnswerService_ServiceDesc, srv)
}

func _AnswerService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnswerService_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerService_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnswerService_Detail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerService_ListForQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListForQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).ListForQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnswerService_ListForQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).ListForQuestion(ctx, req.(*ListForQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerService_ListForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).ListForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnswerService_ListForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).ListForUser(ctx, req.(*ListForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerService_CountForQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountForQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).CountForQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnswerService_CountForQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).CountForQuestion(ctx, req.(*CountForQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerService_DelAnswerById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelAnswerByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerServiceServer).DelAnswerById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnswerService_DelAnswerById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerServiceServer).DelAnswerById(ctx, req.(*DelAnswerByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnswerService_ServiceDesc is the grpc.ServiceDesc for AnswerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnswerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "answer.v1.AnswerService",
	HandlerType: (*AnswerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _AnswerService_Publish_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _AnswerService_Detail_Handler,
		},
		{
			MethodName: "ListForQuestion",
			Handler:    _AnswerService_ListForQuestion_Handler,
		},
		{
			MethodName: "ListForUser",
			Handler:    _AnswerService_ListForUser_Handler,
		},
		{
			MethodName: "CountForQuestion",
			Handler:    _AnswerService_CountForQuestion_Handler,
		},
		{
			MethodName: "DelAnswerById",
			Handler:    _AnswerService_DelAnswerById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/answer/v1/answer.proto",
}
