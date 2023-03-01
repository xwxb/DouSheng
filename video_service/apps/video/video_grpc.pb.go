// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: apps/video/pb/video.proto

package video

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// 视频 feed 流
	FeedVideos(ctx context.Context, in *FeedVideosRequest, opts ...grpc.CallOption) (*FeedSetResponse, error)
	// 视频发布
	PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error)
	// 用户发布视频的列表
	PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error)
	// 根据视频ID获取视频信息
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*Video, error)
	// 根据用户ID获取 用户发布视频数目
	PublishListCount(ctx context.Context, in *PublishListCountRequest, opts ...grpc.CallOption) (*PublishListCountResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) FeedVideos(ctx context.Context, in *FeedVideosRequest, opts ...grpc.CallOption) (*FeedSetResponse, error) {
	out := new(FeedSetResponse)
	err := c.cc.Invoke(ctx, "/dousheng.video.Service/FeedVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error) {
	out := new(PublishVideoResponse)
	err := c.cc.Invoke(ctx, "/dousheng.video.Service/PublishVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error) {
	out := new(PublishListResponse)
	err := c.cc.Invoke(ctx, "/dousheng.video.Service/PublishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*Video, error) {
	out := new(Video)
	err := c.cc.Invoke(ctx, "/dousheng.video.Service/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PublishListCount(ctx context.Context, in *PublishListCountRequest, opts ...grpc.CallOption) (*PublishListCountResponse, error) {
	out := new(PublishListCountResponse)
	err := c.cc.Invoke(ctx, "/dousheng.video.Service/PublishListCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// 视频 feed 流
	FeedVideos(context.Context, *FeedVideosRequest) (*FeedSetResponse, error)
	// 视频发布
	PublishVideo(context.Context, *PublishVideoRequest) (*PublishVideoResponse, error)
	// 用户发布视频的列表
	PublishList(context.Context, *PublishListRequest) (*PublishListResponse, error)
	// 根据视频ID获取视频信息
	GetVideo(context.Context, *GetVideoRequest) (*Video, error)
	// 根据用户ID获取 用户发布视频数目
	PublishListCount(context.Context, *PublishListCountRequest) (*PublishListCountResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) FeedVideos(context.Context, *FeedVideosRequest) (*FeedSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVideos not implemented")
}
func (UnimplementedServiceServer) PublishVideo(context.Context, *PublishVideoRequest) (*PublishVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (UnimplementedServiceServer) PublishList(context.Context, *PublishListRequest) (*PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedServiceServer) GetVideo(context.Context, *GetVideoRequest) (*Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedServiceServer) PublishListCount(context.Context, *PublishListCountRequest) (*PublishListCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishListCount not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_FeedVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedVideosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FeedVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.video.Service/FeedVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FeedVideos(ctx, req.(*FeedVideosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PublishVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PublishVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.video.Service/PublishVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PublishVideo(ctx, req.(*PublishVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.video.Service/PublishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PublishList(ctx, req.(*PublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.video.Service/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PublishListCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PublishListCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dousheng.video.Service/PublishListCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PublishListCount(ctx, req.(*PublishListCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dousheng.video.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FeedVideos",
			Handler:    _Service_FeedVideos_Handler,
		},
		{
			MethodName: "PublishVideo",
			Handler:    _Service_PublishVideo_Handler,
		},
		{
			MethodName: "PublishList",
			Handler:    _Service_PublishList_Handler,
		},
		{
			MethodName: "GetVideo",
			Handler:    _Service_GetVideo_Handler,
		},
		{
			MethodName: "PublishListCount",
			Handler:    _Service_PublishListCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/video/pb/video.proto",
}
