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
	QubicLiveService_GetBalance_FullMethodName           = "/qubic.http.qubic.pb.QubicLiveService/GetBalance"
	QubicLiveService_BroadcastTransaction_FullMethodName = "/qubic.http.qubic.pb.QubicLiveService/BroadcastTransaction"
	QubicLiveService_GetTickInfo_FullMethodName          = "/qubic.http.qubic.pb.QubicLiveService/GetTickInfo"
	QubicLiveService_GetBlockHeight_FullMethodName       = "/qubic.http.qubic.pb.QubicLiveService/GetBlockHeight"
)

// QubicLiveServiceClient is the client API for QubicLiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QubicLiveServiceClient interface {
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*BroadcastTransactionResponse, error)
	GetTickInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTickInfoResponse, error)
	GetBlockHeight(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetBlockHeightResponse, error)
}

type qubicLiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQubicLiveServiceClient(cc grpc.ClientConnInterface) QubicLiveServiceClient {
	return &qubicLiveServiceClient{cc}
}

func (c *qubicLiveServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, QubicLiveService_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qubicLiveServiceClient) BroadcastTransaction(ctx context.Context, in *BroadcastTransactionRequest, opts ...grpc.CallOption) (*BroadcastTransactionResponse, error) {
	out := new(BroadcastTransactionResponse)
	err := c.cc.Invoke(ctx, QubicLiveService_BroadcastTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qubicLiveServiceClient) GetTickInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTickInfoResponse, error) {
	out := new(GetTickInfoResponse)
	err := c.cc.Invoke(ctx, QubicLiveService_GetTickInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qubicLiveServiceClient) GetBlockHeight(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetBlockHeightResponse, error) {
	out := new(GetBlockHeightResponse)
	err := c.cc.Invoke(ctx, QubicLiveService_GetBlockHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QubicLiveServiceServer is the server API for QubicLiveService service.
// All implementations must embed UnimplementedQubicLiveServiceServer
// for forward compatibility
type QubicLiveServiceServer interface {
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*BroadcastTransactionResponse, error)
	GetTickInfo(context.Context, *emptypb.Empty) (*GetTickInfoResponse, error)
	GetBlockHeight(context.Context, *emptypb.Empty) (*GetBlockHeightResponse, error)
	mustEmbedUnimplementedQubicLiveServiceServer()
}

// UnimplementedQubicLiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQubicLiveServiceServer struct {
}

func (UnimplementedQubicLiveServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedQubicLiveServiceServer) BroadcastTransaction(context.Context, *BroadcastTransactionRequest) (*BroadcastTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTransaction not implemented")
}
func (UnimplementedQubicLiveServiceServer) GetTickInfo(context.Context, *emptypb.Empty) (*GetTickInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickInfo not implemented")
}
func (UnimplementedQubicLiveServiceServer) GetBlockHeight(context.Context, *emptypb.Empty) (*GetBlockHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeight not implemented")
}
func (UnimplementedQubicLiveServiceServer) mustEmbedUnimplementedQubicLiveServiceServer() {}

// UnsafeQubicLiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QubicLiveServiceServer will
// result in compilation errors.
type UnsafeQubicLiveServiceServer interface {
	mustEmbedUnimplementedQubicLiveServiceServer()
}

func RegisterQubicLiveServiceServer(s grpc.ServiceRegistrar, srv QubicLiveServiceServer) {
	s.RegisterService(&QubicLiveService_ServiceDesc, srv)
}

func _QubicLiveService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QubicLiveServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QubicLiveService_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QubicLiveServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QubicLiveService_BroadcastTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QubicLiveServiceServer).BroadcastTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QubicLiveService_BroadcastTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QubicLiveServiceServer).BroadcastTransaction(ctx, req.(*BroadcastTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QubicLiveService_GetTickInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QubicLiveServiceServer).GetTickInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QubicLiveService_GetTickInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QubicLiveServiceServer).GetTickInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QubicLiveService_GetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QubicLiveServiceServer).GetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QubicLiveService_GetBlockHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QubicLiveServiceServer).GetBlockHeight(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// QubicLiveService_ServiceDesc is the grpc.ServiceDesc for QubicLiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QubicLiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qubic.http.qubic.pb.QubicLiveService",
	HandlerType: (*QubicLiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _QubicLiveService_GetBalance_Handler,
		},
		{
			MethodName: "BroadcastTransaction",
			Handler:    _QubicLiveService_BroadcastTransaction_Handler,
		},
		{
			MethodName: "GetTickInfo",
			Handler:    _QubicLiveService_GetTickInfo_Handler,
		},
		{
			MethodName: "GetBlockHeight",
			Handler:    _QubicLiveService_GetBlockHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qubic.proto",
}
