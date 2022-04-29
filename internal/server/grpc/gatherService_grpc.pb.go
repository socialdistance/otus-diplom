// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.11.4
// source: gatherService.proto

package grpc

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

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	ListGather(ctx context.Context, in *GatherRequest, opts ...grpc.CallOption) (StreamService_ListGatherClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) ListGather(ctx context.Context, in *GatherRequest, opts ...grpc.CallOption) (StreamService_ListGatherClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], "/api.StreamService/ListGather", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListGatherClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListGatherClient interface {
	Recv() (*GatherResponse, error)
	grpc.ClientStream
}

type streamServiceListGatherClient struct {
	grpc.ClientStream
}

func (x *streamServiceListGatherClient) Recv() (*GatherResponse, error) {
	m := new(GatherResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	ListGather(*GatherRequest, StreamService_ListGatherServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) ListGather(*GatherRequest, StreamService_ListGatherServer) error {
	return status.Errorf(codes.Unimplemented, "method ListGather not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_ListGather_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GatherRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).ListGather(m, &streamServiceListGatherServer{stream})
}

type StreamService_ListGatherServer interface {
	Send(*GatherResponse) error
	grpc.ServerStream
}

type streamServiceListGatherServer struct {
	grpc.ServerStream
}

func (x *streamServiceListGatherServer) Send(m *GatherResponse) error {
	return x.ServerStream.SendMsg(m)
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListGather",
			Handler:       _StreamService_ListGather_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gatherService.proto",
}
