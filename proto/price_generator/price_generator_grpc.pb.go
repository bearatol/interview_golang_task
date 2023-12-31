// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/price_generator/price_generator.proto

package price_generator

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PriceGeneratorClient is the client API for PriceGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceGeneratorClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Pong, error)
	Set(ctx context.Context, in *PriceFileSetRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *PriceFileRequest, opts ...grpc.CallOption) (*PriceFileResponse, error)
	Delete(ctx context.Context, in *PriceFileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type priceGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceGeneratorClient(cc grpc.ClientConnInterface) PriceGeneratorClient {
	return &priceGeneratorClient{cc}
}

func (c *priceGeneratorClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/price_generator.PriceGenerator/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceGeneratorClient) Set(ctx context.Context, in *PriceFileSetRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/price_generator.PriceGenerator/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceGeneratorClient) Get(ctx context.Context, in *PriceFileRequest, opts ...grpc.CallOption) (*PriceFileResponse, error) {
	out := new(PriceFileResponse)
	err := c.cc.Invoke(ctx, "/price_generator.PriceGenerator/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceGeneratorClient) Delete(ctx context.Context, in *PriceFileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/price_generator.PriceGenerator/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceGeneratorServer is the server API for PriceGenerator service.
// All implementations must embed UnimplementedPriceGeneratorServer
// for forward compatibility
type PriceGeneratorServer interface {
	Ping(context.Context, *empty.Empty) (*Pong, error)
	Set(context.Context, *PriceFileSetRequest) (*empty.Empty, error)
	Get(context.Context, *PriceFileRequest) (*PriceFileResponse, error)
	Delete(context.Context, *PriceFileRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPriceGeneratorServer()
}

// UnimplementedPriceGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedPriceGeneratorServer struct {
}

func (UnimplementedPriceGeneratorServer) Ping(context.Context, *empty.Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPriceGeneratorServer) Set(context.Context, *PriceFileSetRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedPriceGeneratorServer) Get(context.Context, *PriceFileRequest) (*PriceFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPriceGeneratorServer) Delete(context.Context, *PriceFileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPriceGeneratorServer) mustEmbedUnimplementedPriceGeneratorServer() {}

// UnsafePriceGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceGeneratorServer will
// result in compilation errors.
type UnsafePriceGeneratorServer interface {
	mustEmbedUnimplementedPriceGeneratorServer()
}

func RegisterPriceGeneratorServer(s grpc.ServiceRegistrar, srv PriceGeneratorServer) {
	s.RegisterService(&PriceGenerator_ServiceDesc, srv)
}

func _PriceGenerator_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceGeneratorServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price_generator.PriceGenerator/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceGeneratorServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceGenerator_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceFileSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceGeneratorServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price_generator.PriceGenerator/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceGeneratorServer).Set(ctx, req.(*PriceFileSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceGenerator_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceGeneratorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price_generator.PriceGenerator/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceGeneratorServer).Get(ctx, req.(*PriceFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceGenerator_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceGeneratorServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/price_generator.PriceGenerator/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceGeneratorServer).Delete(ctx, req.(*PriceFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceGenerator_ServiceDesc is the grpc.ServiceDesc for PriceGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "price_generator.PriceGenerator",
	HandlerType: (*PriceGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PriceGenerator_Ping_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _PriceGenerator_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PriceGenerator_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PriceGenerator_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/price_generator/price_generator.proto",
}
