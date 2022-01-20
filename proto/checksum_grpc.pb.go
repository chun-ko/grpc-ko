// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: proto/checksum.proto

package redbelly_grpcrace

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

// ChecksumServiceClient is the client API for ChecksumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChecksumServiceClient interface {
	Checksum(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type checksumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChecksumServiceClient(cc grpc.ClientConnInterface) ChecksumServiceClient {
	return &checksumServiceClient{cc}
}

func (c *checksumServiceClient) Checksum(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/proto.ChecksumService/Checksum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChecksumServiceServer is the server API for ChecksumService service.
// All implementations must embed UnimplementedChecksumServiceServer
// for forward compatibility
type ChecksumServiceServer interface {
	Checksum(context.Context, *CheckRequest) (*CheckResponse, error)
	mustEmbedUnimplementedChecksumServiceServer()
}

// UnimplementedChecksumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChecksumServiceServer struct {
}

func (UnimplementedChecksumServiceServer) Checksum(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checksum not implemented")
}
func (UnimplementedChecksumServiceServer) mustEmbedUnimplementedChecksumServiceServer() {}

// UnsafeChecksumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChecksumServiceServer will
// result in compilation errors.
type UnsafeChecksumServiceServer interface {
	mustEmbedUnimplementedChecksumServiceServer()
}

func RegisterChecksumServiceServer(s grpc.ServiceRegistrar, srv ChecksumServiceServer) {
	s.RegisterService(&ChecksumService_ServiceDesc, srv)
}

func _ChecksumService_Checksum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChecksumServiceServer).Checksum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChecksumService/Checksum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChecksumServiceServer).Checksum(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChecksumService_ServiceDesc is the grpc.ServiceDesc for ChecksumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChecksumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChecksumService",
	HandlerType: (*ChecksumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Checksum",
			Handler:    _ChecksumService_Checksum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/checksum.proto",
}
