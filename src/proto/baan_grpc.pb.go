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
	GetAllBaan(ctx context.Context, in *GetAllBaanRequest, opts ...grpc.CallOption) (*GetAllBaanResponse, error)
	GetBaan(ctx context.Context, in *GetBaanRequest, opts ...grpc.CallOption) (*GetBaanResponse, error)
}

type baanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaanServiceClient(cc grpc.ClientConnInterface) BaanServiceClient {
	return &baanServiceClient{cc}
}

func (c *baanServiceClient) GetAllBaan(ctx context.Context, in *GetAllBaanRequest, opts ...grpc.CallOption) (*GetAllBaanResponse, error) {
	out := new(GetAllBaanResponse)
	err := c.cc.Invoke(ctx, "/baan.BaanService/GetAllBaan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baanServiceClient) GetBaan(ctx context.Context, in *GetBaanRequest, opts ...grpc.CallOption) (*GetBaanResponse, error) {
	out := new(GetBaanResponse)
	err := c.cc.Invoke(ctx, "/baan.BaanService/GetBaan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaanServiceServer is the server API for BaanService service.
// All implementations should embed UnimplementedBaanServiceServer
// for forward compatibility
type BaanServiceServer interface {
	GetAllBaan(context.Context, *GetAllBaanRequest) (*GetAllBaanResponse, error)
	GetBaan(context.Context, *GetBaanRequest) (*GetBaanResponse, error)
}

// UnimplementedBaanServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBaanServiceServer struct {
}

func (UnimplementedBaanServiceServer) GetAllBaan(context.Context, *GetAllBaanRequest) (*GetAllBaanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBaan not implemented")
}
func (UnimplementedBaanServiceServer) GetBaan(context.Context, *GetBaanRequest) (*GetBaanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBaan not implemented")
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

func _BaanService_GetAllBaan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBaanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaanServiceServer).GetAllBaan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baan.BaanService/GetAllBaan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaanServiceServer).GetAllBaan(ctx, req.(*GetAllBaanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaanService_GetBaan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBaanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaanServiceServer).GetBaan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baan.BaanService/GetBaan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaanServiceServer).GetBaan(ctx, req.(*GetBaanRequest))
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
			MethodName: "GetAllBaan",
			Handler:    _BaanService_GetAllBaan_Handler,
		},
		{
			MethodName: "GetBaan",
			Handler:    _BaanService_GetBaan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "baan.proto",
}