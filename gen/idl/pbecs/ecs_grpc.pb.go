// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbecs

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

// ECSServiceClient is the client API for ECSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ECSServiceClient interface {
	ECSStatistic(ctx context.Context, in *ECSStatisticReq, opts ...grpc.CallOption) (*ECSStatisticResp, error)
}

type eCSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewECSServiceClient(cc grpc.ClientConnInterface) ECSServiceClient {
	return &eCSServiceClient{cc}
}

func (c *eCSServiceClient) ECSStatistic(ctx context.Context, in *ECSStatisticReq, opts ...grpc.CallOption) (*ECSStatisticResp, error) {
	out := new(ECSStatisticResp)
	err := c.cc.Invoke(ctx, "/pbecs.ECSService/ECSStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ECSServiceServer is the server API for ECSService service.
// All implementations must embed UnimplementedECSServiceServer
// for forward compatibility
type ECSServiceServer interface {
	ECSStatistic(context.Context, *ECSStatisticReq) (*ECSStatisticResp, error)
	mustEmbedUnimplementedECSServiceServer()
}

// UnimplementedECSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedECSServiceServer struct {
}

func (UnimplementedECSServiceServer) ECSStatistic(context.Context, *ECSStatisticReq) (*ECSStatisticResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ECSStatistic not implemented")
}
func (UnimplementedECSServiceServer) mustEmbedUnimplementedECSServiceServer() {}

// UnsafeECSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ECSServiceServer will
// result in compilation errors.
type UnsafeECSServiceServer interface {
	mustEmbedUnimplementedECSServiceServer()
}

func RegisterECSServiceServer(s grpc.ServiceRegistrar, srv ECSServiceServer) {
	s.RegisterService(&ECSService_ServiceDesc, srv)
}

func _ECSService_ECSStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ECSStatisticReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ECSServiceServer).ECSStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbecs.ECSService/ECSStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ECSServiceServer).ECSStatistic(ctx, req.(*ECSStatisticReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ECSService_ServiceDesc is the grpc.ServiceDesc for ECSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ECSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbecs.ECSService",
	HandlerType: (*ECSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ECSStatistic",
			Handler:    _ECSService_ECSStatistic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/pbecs/ecs.proto",
}
