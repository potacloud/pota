// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ImagesClient is the client API for Images service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImagesClient interface {
	CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error)
}

type imagesClient struct {
	cc grpc.ClientConnInterface
}

func NewImagesClient(cc grpc.ClientConnInterface) ImagesClient {
	return &imagesClient{cc}
}

func (c *imagesClient) CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error) {
	out := new(CreateImageResponse)
	err := c.cc.Invoke(ctx, "/images.Images/CreateImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImagesServer is the server API for Images service.
// All implementations must embed UnimplementedImagesServer
// for forward compatibility
type ImagesServer interface {
	CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error)
	mustEmbedUnimplementedImagesServer()
}

// UnimplementedImagesServer must be embedded to have forward compatible implementations.
type UnimplementedImagesServer struct {
}

func (UnimplementedImagesServer) CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (UnimplementedImagesServer) mustEmbedUnimplementedImagesServer() {}

// UnsafeImagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImagesServer will
// result in compilation errors.
type UnsafeImagesServer interface {
	mustEmbedUnimplementedImagesServer()
}

func RegisterImagesServer(s grpc.ServiceRegistrar, srv ImagesServer) {
	s.RegisterService(&Images_ServiceDesc, srv)
}

func _Images_CreateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).CreateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.Images/CreateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).CreateImage(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Images_ServiceDesc is the grpc.ServiceDesc for Images service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Images_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "images.Images",
	HandlerType: (*ImagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateImage",
			Handler:    _Images_CreateImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "images.proto",
}
