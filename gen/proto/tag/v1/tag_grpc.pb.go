// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/tag/v1/tag.proto

package tagv1

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
	TagService_AttachAssessmentTags_FullMethodName = "/tag.v1.TagService/AttachAssessmentTags"
	TagService_AttachFeatureTags_FullMethodName    = "/tag.v1.TagService/AttachFeatureTags"
)

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	AttachAssessmentTags(ctx context.Context, in *AttachAssessmentTagsRequest, opts ...grpc.CallOption) (*AttachAssessmentTagsResponse, error)
	AttachFeatureTags(ctx context.Context, in *AttachFeatureTagsRequest, opts ...grpc.CallOption) (*AttachFeatureTagsResponse, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) AttachAssessmentTags(ctx context.Context, in *AttachAssessmentTagsRequest, opts ...grpc.CallOption) (*AttachAssessmentTagsResponse, error) {
	out := new(AttachAssessmentTagsResponse)
	err := c.cc.Invoke(ctx, TagService_AttachAssessmentTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) AttachFeatureTags(ctx context.Context, in *AttachFeatureTagsRequest, opts ...grpc.CallOption) (*AttachFeatureTagsResponse, error) {
	out := new(AttachFeatureTagsResponse)
	err := c.cc.Invoke(ctx, TagService_AttachFeatureTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	AttachAssessmentTags(context.Context, *AttachAssessmentTagsRequest) (*AttachAssessmentTagsResponse, error)
	AttachFeatureTags(context.Context, *AttachFeatureTagsRequest) (*AttachFeatureTagsResponse, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) AttachAssessmentTags(context.Context, *AttachAssessmentTagsRequest) (*AttachAssessmentTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachAssessmentTags not implemented")
}
func (UnimplementedTagServiceServer) AttachFeatureTags(context.Context, *AttachFeatureTagsRequest) (*AttachFeatureTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachFeatureTags not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_AttachAssessmentTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachAssessmentTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).AttachAssessmentTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_AttachAssessmentTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).AttachAssessmentTags(ctx, req.(*AttachAssessmentTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_AttachFeatureTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachFeatureTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).AttachFeatureTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_AttachFeatureTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).AttachFeatureTags(ctx, req.(*AttachFeatureTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tag.v1.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttachAssessmentTags",
			Handler:    _TagService_AttachAssessmentTags_Handler,
		},
		{
			MethodName: "AttachFeatureTags",
			Handler:    _TagService_AttachFeatureTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tag/v1/tag.proto",
}