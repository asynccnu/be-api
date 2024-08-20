// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/banner/v1/banner.proto

package bannerv1

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
	BannerService_GetBanners_FullMethodName = "/banner.v1.BannerService/GetBanners"
	BannerService_SaveBanner_FullMethodName = "/banner.v1.BannerService/SaveBanner"
	BannerService_AddBanner_FullMethodName  = "/banner.v1.BannerService/AddBanner"
	BannerService_DelBanner_FullMethodName  = "/banner.v1.BannerService/DelBanner"
)

// BannerServiceClient is the client API for BannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BannerServiceClient interface {
	GetBanners(ctx context.Context, in *GetBannersRequest, opts ...grpc.CallOption) (*GetBannersResponse, error)
	SaveBanner(ctx context.Context, in *SaveBannerRequest, opts ...grpc.CallOption) (*SaveBannerResponse, error)
	AddBanner(ctx context.Context, in *AddBannerRequest, opts ...grpc.CallOption) (*AddBannerResponse, error)
	DelBanner(ctx context.Context, in *DelBannerRequest, opts ...grpc.CallOption) (*DelBannerResponse, error)
}

type bannerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBannerServiceClient(cc grpc.ClientConnInterface) BannerServiceClient {
	return &bannerServiceClient{cc}
}

func (c *bannerServiceClient) GetBanners(ctx context.Context, in *GetBannersRequest, opts ...grpc.CallOption) (*GetBannersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBannersResponse)
	err := c.cc.Invoke(ctx, BannerService_GetBanners_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) SaveBanner(ctx context.Context, in *SaveBannerRequest, opts ...grpc.CallOption) (*SaveBannerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveBannerResponse)
	err := c.cc.Invoke(ctx, BannerService_SaveBanner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) AddBanner(ctx context.Context, in *AddBannerRequest, opts ...grpc.CallOption) (*AddBannerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBannerResponse)
	err := c.cc.Invoke(ctx, BannerService_AddBanner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) DelBanner(ctx context.Context, in *DelBannerRequest, opts ...grpc.CallOption) (*DelBannerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelBannerResponse)
	err := c.cc.Invoke(ctx, BannerService_DelBanner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BannerServiceServer is the server API for BannerService service.
// All implementations must embed UnimplementedBannerServiceServer
// for forward compatibility.
type BannerServiceServer interface {
	GetBanners(context.Context, *GetBannersRequest) (*GetBannersResponse, error)
	SaveBanner(context.Context, *SaveBannerRequest) (*SaveBannerResponse, error)
	AddBanner(context.Context, *AddBannerRequest) (*AddBannerResponse, error)
	DelBanner(context.Context, *DelBannerRequest) (*DelBannerResponse, error)
	mustEmbedUnimplementedBannerServiceServer()
}

// UnimplementedBannerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBannerServiceServer struct{}

func (UnimplementedBannerServiceServer) GetBanners(context.Context, *GetBannersRequest) (*GetBannersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanners not implemented")
}
func (UnimplementedBannerServiceServer) SaveBanner(context.Context, *SaveBannerRequest) (*SaveBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBanner not implemented")
}
func (UnimplementedBannerServiceServer) AddBanner(context.Context, *AddBannerRequest) (*AddBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBanner not implemented")
}
func (UnimplementedBannerServiceServer) DelBanner(context.Context, *DelBannerRequest) (*DelBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBanner not implemented")
}
func (UnimplementedBannerServiceServer) mustEmbedUnimplementedBannerServiceServer() {}
func (UnimplementedBannerServiceServer) testEmbeddedByValue()                       {}

// UnsafeBannerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BannerServiceServer will
// result in compilation errors.
type UnsafeBannerServiceServer interface {
	mustEmbedUnimplementedBannerServiceServer()
}

func RegisterBannerServiceServer(s grpc.ServiceRegistrar, srv BannerServiceServer) {
	// If the following call pancis, it indicates UnimplementedBannerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BannerService_ServiceDesc, srv)
}

func _BannerService_GetBanners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBannersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).GetBanners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_GetBanners_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).GetBanners(ctx, req.(*GetBannersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_SaveBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).SaveBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_SaveBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).SaveBanner(ctx, req.(*SaveBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_AddBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).AddBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_AddBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).AddBanner(ctx, req.(*AddBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_DelBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).DelBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BannerService_DelBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).DelBanner(ctx, req.(*DelBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BannerService_ServiceDesc is the grpc.ServiceDesc for BannerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BannerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "banner.v1.BannerService",
	HandlerType: (*BannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBanners",
			Handler:    _BannerService_GetBanners_Handler,
		},
		{
			MethodName: "SaveBanner",
			Handler:    _BannerService_SaveBanner_Handler,
		},
		{
			MethodName: "AddBanner",
			Handler:    _BannerService_AddBanner_Handler,
		},
		{
			MethodName: "DelBanner",
			Handler:    _BannerService_DelBanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/banner/v1/banner.proto",
}