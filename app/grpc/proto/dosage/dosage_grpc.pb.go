// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: app/grpc/proto/dosage/dosage.proto

package dosage

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

// DosageServiceClient is the client API for DosageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DosageServiceClient interface {
	GetDosageById(ctx context.Context, in *DosageRequest, opts ...grpc.CallOption) (*DosageResponse, error)
}

type dosageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDosageServiceClient(cc grpc.ClientConnInterface) DosageServiceClient {
	return &dosageServiceClient{cc}
}

func (c *dosageServiceClient) GetDosageById(ctx context.Context, in *DosageRequest, opts ...grpc.CallOption) (*DosageResponse, error) {
	out := new(DosageResponse)
	err := c.cc.Invoke(ctx, "/grpc.DosageService/GetDosageById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DosageServiceServer is the server API for DosageService service.
// All implementations must embed UnimplementedDosageServiceServer
// for forward compatibility
type DosageServiceServer interface {
	GetDosageById(context.Context, *DosageRequest) (*DosageResponse, error)
	mustEmbedUnimplementedDosageServiceServer()
}

// UnimplementedDosageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDosageServiceServer struct {
}

func (UnimplementedDosageServiceServer) GetDosageById(context.Context, *DosageRequest) (*DosageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDosageById not implemented")
}
func (UnimplementedDosageServiceServer) mustEmbedUnimplementedDosageServiceServer() {}

// UnsafeDosageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DosageServiceServer will
// result in compilation errors.
type UnsafeDosageServiceServer interface {
	mustEmbedUnimplementedDosageServiceServer()
}

func RegisterDosageServiceServer(s grpc.ServiceRegistrar, srv DosageServiceServer) {
	s.RegisterService(&DosageService_ServiceDesc, srv)
}

func _DosageService_GetDosageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DosageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DosageServiceServer).GetDosageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DosageService/GetDosageById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DosageServiceServer).GetDosageById(ctx, req.(*DosageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DosageService_ServiceDesc is the grpc.ServiceDesc for DosageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DosageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DosageService",
	HandlerType: (*DosageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDosageById",
			Handler:    _DosageService_GetDosageById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/grpc/proto/dosage/dosage.proto",
}