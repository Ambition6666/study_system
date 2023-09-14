// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: private.proto

package problemrpc

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
	PrivateService_GetProblem_FullMethodName = "/private.PrivateService/GetProblem"
)

// PrivateServiceClient is the client API for PrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivateServiceClient interface {
	GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemResponse, error)
}

type privateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateServiceClient(cc grpc.ClientConnInterface) PrivateServiceClient {
	return &privateServiceClient{cc}
}

func (c *privateServiceClient) GetProblem(ctx context.Context, in *GetProblemRequest, opts ...grpc.CallOption) (*GetProblemResponse, error) {
	out := new(GetProblemResponse)
	err := c.cc.Invoke(ctx, PrivateService_GetProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateServiceServer is the server API for PrivateService service.
// All implementations must embed UnimplementedPrivateServiceServer
// for forward compatibility
type PrivateServiceServer interface {
	GetProblem(context.Context, *GetProblemRequest) (*GetProblemResponse, error)
	mustEmbedUnimplementedPrivateServiceServer()
}

// UnimplementedPrivateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPrivateServiceServer struct {
}

func (UnimplementedPrivateServiceServer) GetProblem(context.Context, *GetProblemRequest) (*GetProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblem not implemented")
}
func (UnimplementedPrivateServiceServer) mustEmbedUnimplementedPrivateServiceServer() {}

// UnsafePrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateServiceServer will
// result in compilation errors.
type UnsafePrivateServiceServer interface {
	mustEmbedUnimplementedPrivateServiceServer()
}

func RegisterPrivateServiceServer(s grpc.ServiceRegistrar, srv PrivateServiceServer) {
	s.RegisterService(&PrivateService_ServiceDesc, srv)
}

func _PrivateService_GetProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PrivateService_GetProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetProblem(ctx, req.(*GetProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivateService_ServiceDesc is the grpc.ServiceDesc for PrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "private.PrivateService",
	HandlerType: (*PrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProblem",
			Handler:    _PrivateService_GetProblem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "private.proto",
}