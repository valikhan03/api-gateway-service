// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/auctions-service/protobuf/auctions_management_service.proto

package auctions_management_service

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

// AuctionsManagementServiceClient is the client API for AuctionsManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionsManagementServiceClient interface {
	StartAuction(ctx context.Context, in *StartAuctionRequest, opts ...grpc.CallOption) (*StartAuctionResponse, error)
	CancelAuction(ctx context.Context, in *CancelAuctionRequest, opts ...grpc.CallOption) (*CancelAuctionResponse, error)
	SetStartTime(ctx context.Context, in *SetStartTimeRequest, opts ...grpc.CallOption) (*SetStartTimeResponse, error)
}

type auctionsManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionsManagementServiceClient(cc grpc.ClientConnInterface) AuctionsManagementServiceClient {
	return &auctionsManagementServiceClient{cc}
}

func (c *auctionsManagementServiceClient) StartAuction(ctx context.Context, in *StartAuctionRequest, opts ...grpc.CallOption) (*StartAuctionResponse, error) {
	out := new(StartAuctionResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AuctionsManagementService/StartAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionsManagementServiceClient) CancelAuction(ctx context.Context, in *CancelAuctionRequest, opts ...grpc.CallOption) (*CancelAuctionResponse, error) {
	out := new(CancelAuctionResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AuctionsManagementService/CancelAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionsManagementServiceClient) SetStartTime(ctx context.Context, in *SetStartTimeRequest, opts ...grpc.CallOption) (*SetStartTimeResponse, error) {
	out := new(SetStartTimeResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AuctionsManagementService/SetStartTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionsManagementServiceServer is the server API for AuctionsManagementService service.
// All implementations must embed UnimplementedAuctionsManagementServiceServer
// for forward compatibility
type AuctionsManagementServiceServer interface {
	StartAuction(context.Context, *StartAuctionRequest) (*StartAuctionResponse, error)
	CancelAuction(context.Context, *CancelAuctionRequest) (*CancelAuctionResponse, error)
	SetStartTime(context.Context, *SetStartTimeRequest) (*SetStartTimeResponse, error)
	mustEmbedUnimplementedAuctionsManagementServiceServer()
}

// UnimplementedAuctionsManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionsManagementServiceServer struct {
}

func (UnimplementedAuctionsManagementServiceServer) StartAuction(context.Context, *StartAuctionRequest) (*StartAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAuction not implemented")
}
func (UnimplementedAuctionsManagementServiceServer) CancelAuction(context.Context, *CancelAuctionRequest) (*CancelAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAuction not implemented")
}
func (UnimplementedAuctionsManagementServiceServer) SetStartTime(context.Context, *SetStartTimeRequest) (*SetStartTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStartTime not implemented")
}
func (UnimplementedAuctionsManagementServiceServer) mustEmbedUnimplementedAuctionsManagementServiceServer() {
}

// UnsafeAuctionsManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionsManagementServiceServer will
// result in compilation errors.
type UnsafeAuctionsManagementServiceServer interface {
	mustEmbedUnimplementedAuctionsManagementServiceServer()
}

func RegisterAuctionsManagementServiceServer(s grpc.ServiceRegistrar, srv AuctionsManagementServiceServer) {
	s.RegisterService(&AuctionsManagementService_ServiceDesc, srv)
}

func _AuctionsManagementService_StartAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartAuctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionsManagementServiceServer).StartAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AuctionsManagementService/StartAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionsManagementServiceServer).StartAuction(ctx, req.(*StartAuctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionsManagementService_CancelAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelAuctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionsManagementServiceServer).CancelAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AuctionsManagementService/CancelAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionsManagementServiceServer).CancelAuction(ctx, req.(*CancelAuctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionsManagementService_SetStartTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStartTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionsManagementServiceServer).SetStartTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AuctionsManagementService/SetStartTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionsManagementServiceServer).SetStartTime(ctx, req.(*SetStartTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionsManagementService_ServiceDesc is the grpc.ServiceDesc for AuctionsManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionsManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.AuctionsManagementService",
	HandlerType: (*AuctionsManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartAuction",
			Handler:    _AuctionsManagementService_StartAuction_Handler,
		},
		{
			MethodName: "CancelAuction",
			Handler:    _AuctionsManagementService_CancelAuction_Handler,
		},
		{
			MethodName: "SetStartTime",
			Handler:    _AuctionsManagementService_SetStartTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/auctions-service/protobuf/auctions_management_service.proto",
}
