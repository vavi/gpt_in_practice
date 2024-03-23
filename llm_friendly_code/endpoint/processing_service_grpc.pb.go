// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: grpc/processing_service.proto

package endpoint

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

// ProcessingServiceClient is the client API for ProcessingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessingServiceClient interface {
	// Sends a greeting
	Process(ctx context.Context, in *ProcessingRequest, opts ...grpc.CallOption) (*ProcessingReply, error)
	Retrieve(ctx context.Context, in *RetrievingRequest, opts ...grpc.CallOption) (*RetrievingReply, error)
}

type processingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessingServiceClient(cc grpc.ClientConnInterface) ProcessingServiceClient {
	return &processingServiceClient{cc}
}

func (c *processingServiceClient) Process(ctx context.Context, in *ProcessingRequest, opts ...grpc.CallOption) (*ProcessingReply, error) {
	out := new(ProcessingReply)
	err := c.cc.Invoke(ctx, "/ProcessingService/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processingServiceClient) Retrieve(ctx context.Context, in *RetrievingRequest, opts ...grpc.CallOption) (*RetrievingReply, error) {
	out := new(RetrievingReply)
	err := c.cc.Invoke(ctx, "/ProcessingService/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessingServiceServer is the server API for ProcessingService service.
// All implementations must embed UnimplementedProcessingServiceServer
// for forward compatibility
type ProcessingServiceServer interface {
	// Sends a greeting
	Process(context.Context, *ProcessingRequest) (*ProcessingReply, error)
	Retrieve(context.Context, *RetrievingRequest) (*RetrievingReply, error)
	mustEmbedUnimplementedProcessingServiceServer()
}

// UnimplementedProcessingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessingServiceServer struct {
}

func (UnimplementedProcessingServiceServer) Process(context.Context, *ProcessingRequest) (*ProcessingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (UnimplementedProcessingServiceServer) Retrieve(context.Context, *RetrievingRequest) (*RetrievingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedProcessingServiceServer) mustEmbedUnimplementedProcessingServiceServer() {}

// UnsafeProcessingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessingServiceServer will
// result in compilation errors.
type UnsafeProcessingServiceServer interface {
	mustEmbedUnimplementedProcessingServiceServer()
}

func RegisterProcessingServiceServer(s grpc.ServiceRegistrar, srv ProcessingServiceServer) {
	s.RegisterService(&ProcessingService_ServiceDesc, srv)
}

func _ProcessingService_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessingServiceServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessingService/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessingServiceServer).Process(ctx, req.(*ProcessingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessingService_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrievingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessingServiceServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcessingService/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessingServiceServer).Retrieve(ctx, req.(*RetrievingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessingService_ServiceDesc is the grpc.ServiceDesc for ProcessingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProcessingService",
	HandlerType: (*ProcessingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _ProcessingService_Process_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _ProcessingService_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/processing_service.proto",
}
