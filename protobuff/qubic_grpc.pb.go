// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: qubic.proto

package protobuff

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QubicService_GetBalance_FullMethodName           = "/qubic.http.qubic.pb.QubicService/GetBalance"
	QubicService_BroadcastTransaction_FullMethodName = "/qubic.http.qubic.pb.QubicService/BroadcastTransaction"
	QubicService_GetTickInfo_FullMethodName          = "/qubic.http.qubic.pb.QubicService/GetTickInfo"
)

// QubicServiceClient is the client API for QubicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QubicServiceClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*BroadcastTransactionResponse, error)
	GetTickInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTickInfoResponse, error)
}

type qubicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQubicServiceClient(cc grpc.ClientConnInterface) QubicServiceClient {
	return &qubicServiceClient{cc}
}

func (c *qubicServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, QubicService_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qubicServiceClient) BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*BroadcastTransactionResponse, error) {
	out := new(BroadcastTransactionResponse)
	err := c.cc.Invoke(ctx, QubicService_BroadcastTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qubicServiceClient) GetTickInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTickInfoResponse, error) {
	out := new(GetTickInfoResponse)
	err := c.cc.Invoke(ctx, QubicService_GetTickInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QubicServiceServer is the server API for QubicService service.
// All implementations must embed UnimplementedQubicServiceServer
// for forward compatibility
type QubicServiceServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*BroadcastTransactionResponse, error)
	GetTickInfo(context.Context, *emptypb.Empty) (*GetTickInfoResponse, error)
	mustEmbedUnimplementedQubicServiceServer()
}

// UnimplementedQubicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQubicServiceServer struct {
}

func (UnimplementedQubicServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedQubicServiceServer) BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*BroadcastTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTransaction not implemented")
}
func (UnimplementedQubicServiceServer) GetTickInfo(context.Context, *emptypb.Empty) (*GetTickInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickInfo not implemented")
}
func (UnimplementedQubicServiceServer) mustEmbedUnimplementedQubicServiceServer() {}

// UnsafeQubicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QubicServiceServer will
// result in compilation errors.
type UnsafeQubicServiceServer interface {
	mustEmbedUnimplementedQubicServiceServer()
}

func RegisterQubicServiceServer(s grpc.ServiceRegistrar, srv QubicServiceServer) {
	s.RegisterService(&QubicService_ServiceDesc, srv)
}

func _QubicService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QubicServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QubicService_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QubicServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QubicService_BroadcastTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QubicServiceServer).BroadcastTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QubicService_BroadcastTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QubicServiceServer).BroadcastTransaction(ctx, req.(*BroadcastTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QubicService_GetTickInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QubicServiceServer).GetTickInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QubicService_GetTickInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QubicServiceServer).GetTickInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// QubicService_ServiceDesc is the grpc.ServiceDesc for QubicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QubicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qubic.http.qubic.pb.QubicService",
	HandlerType: (*QubicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _QubicService_GetBalance_Handler,
		},
		{
			MethodName: "BroadcastTransaction",
			Handler:    _QubicService_BroadcastTransaction_Handler,
		},
		{
			MethodName: "GetTickInfo",
			Handler:    _QubicService_GetTickInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qubic.proto",
}
