// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package internal_handler

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EvhubInternalHandlerClient is the client API for EvhubInternalHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvhubInternalHandlerClient interface {
	Fork(ctx context.Context, in *ForkReq, opts ...grpc.CallOption) (*ForkRsp, error)
}

type evhubInternalHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewEvhubInternalHandlerClient(cc grpc.ClientConnInterface) EvhubInternalHandlerClient {
	return &evhubInternalHandlerClient{cc}
}

func (c *evhubInternalHandlerClient) Fork(ctx context.Context, in *ForkReq, opts ...grpc.CallOption) (*ForkRsp, error) {
	out := new(ForkRsp)
	err := c.cc.Invoke(ctx, "/evhub_internal_handler.evhubInternalHandler/Fork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvhubInternalHandlerServer is the server API for EvhubInternalHandler service.
// All implementations should embed UnimplementedEvhubInternalHandlerServer
// for forward compatibility
type EvhubInternalHandlerServer interface {
	Fork(context.Context, *ForkReq) (*ForkRsp, error)
}

// UnimplementedEvhubInternalHandlerServer should be embedded to have forward compatible implementations.
type UnimplementedEvhubInternalHandlerServer struct {
}

func (UnimplementedEvhubInternalHandlerServer) Fork(context.Context, *ForkReq) (*ForkRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fork not implemented")
}

// UnsafeEvhubInternalHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvhubInternalHandlerServer will
// result in compilation errors.
type UnsafeEvhubInternalHandlerServer interface {
	mustEmbedUnimplementedEvhubInternalHandlerServer()
}

func RegisterEvhubInternalHandlerServer(s *grpc.Server, srv EvhubInternalHandlerServer) {
	s.RegisterService(&_EvhubInternalHandler_serviceDesc, srv)
}

func _EvhubInternalHandler_Fork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvhubInternalHandlerServer).Fork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evhub_internal_handler.evhubInternalHandler/Fork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvhubInternalHandlerServer).Fork(ctx, req.(*ForkReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EvhubInternalHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evhub_internal_handler.evhubInternalHandler",
	HandlerType: (*EvhubInternalHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fork",
			Handler:    _EvhubInternalHandler_Fork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal_handler/internal_handler.proto",
}
