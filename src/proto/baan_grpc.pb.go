// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: baan.proto

package proto

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

// BaanServiceClient is the client API for BaanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaanServiceClient interface {
	FindAllBaan(ctx context.Context, in *FindAllBaanRequest, opts ...grpc.CallOption) (*FindAllBaanResponse, error)
	FindOneBaan(ctx context.Context, in *FindOneBaanRequest, opts ...grpc.CallOption) (*FindOneBaanResponse, error)
}

type baanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaanServiceClient(cc grpc.ClientConnInterface) BaanServiceClient {
	return &baanServiceClient{cc}
}

func (c *baanServiceClient) FindAllBaan(ctx context.Context, in *FindAllBaanRequest, opts ...grpc.CallOption) (*FindAllBaanResponse, error) {
	out := new(FindAllBaanResponse)
	err := c.cc.Invoke(ctx, "/baan.BaanService/FindAllBaan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baanServiceClient) FindOneBaan(ctx context.Context, in *FindOneBaanRequest, opts ...grpc.CallOption) (*FindOneBaanResponse, error) {
	out := new(FindOneBaanResponse)
	err := c.cc.Invoke(ctx, "/baan.BaanService/FindOneBaan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaanServiceServer is the server API for BaanService service.
// All implementations should embed UnimplementedBaanServiceServer
// for forward compatibility
type BaanServiceServer interface {
	FindAllBaan(context.Context, *FindAllBaanRequest) (*FindAllBaanResponse, error)
	FindOneBaan(context.Context, *FindOneBaanRequest) (*FindOneBaanResponse, error)
}

// UnimplementedBaanServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBaanServiceServer struct {
}

func (UnimplementedBaanServiceServer) FindAllBaan(context.Context, *FindAllBaanRequest) (*FindAllBaanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllBaan not implemented")
}
func (UnimplementedBaanServiceServer) FindOneBaan(context.Context, *FindOneBaanRequest) (*FindOneBaanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneBaan not implemented")
}

// UnsafeBaanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaanServiceServer will
// result in compilation errors.
type UnsafeBaanServiceServer interface {
	mustEmbedUnimplementedBaanServiceServer()
}

func RegisterBaanServiceServer(s grpc.ServiceRegistrar, srv BaanServiceServer) {
	s.RegisterService(&BaanService_ServiceDesc, srv)
}

func _BaanService_FindAllBaan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllBaanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaanServiceServer).FindAllBaan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baan.BaanService/FindAllBaan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaanServiceServer).FindAllBaan(ctx, req.(*FindAllBaanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaanService_FindOneBaan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneBaanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaanServiceServer).FindOneBaan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baan.BaanService/FindOneBaan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaanServiceServer).FindOneBaan(ctx, req.(*FindOneBaanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BaanService_ServiceDesc is the grpc.ServiceDesc for BaanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "baan.BaanService",
	HandlerType: (*BaanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllBaan",
			Handler:    _BaanService_FindAllBaan_Handler,
		},
		{
			MethodName: "FindOneBaan",
			Handler:    _BaanService_FindOneBaan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "baan.proto",
}
