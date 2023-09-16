// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: types.proto

package apis

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
	RankService_Register_FullMethodName        = "/apis.RankService/Register"
	RankService_Update_FullMethodName          = "/apis.RankService/Update"
	RankService_GetFatRate_FullMethodName      = "/apis.RankService/GetFatRate"
	RankService_GetTop_FullMethodName          = "/apis.RankService/GetTop"
	RankService_RegisterPersons_FullMethodName = "/apis.RankService/RegisterPersons"
	RankService_WatchPersons_FullMethodName    = "/apis.RankService/WatchPersons"
)

// RankServiceClient is the client API for RankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RankServiceClient interface {
	Register(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalInformation, error)
	Update(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalInformationFatRate, error)
	GetFatRate(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalRank, error)
	GetTop(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Ranks, error)
	RegisterPersons(ctx context.Context, opts ...grpc.CallOption) (RankService_RegisterPersonsClient, error)
	WatchPersons(ctx context.Context, in *Null, opts ...grpc.CallOption) (RankService_WatchPersonsClient, error)
}

type rankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRankServiceClient(cc grpc.ClientConnInterface) RankServiceClient {
	return &rankServiceClient{cc}
}

func (c *rankServiceClient) Register(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalInformation, error) {
	out := new(PersonalInformation)
	err := c.cc.Invoke(ctx, RankService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankServiceClient) Update(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalInformationFatRate, error) {
	out := new(PersonalInformationFatRate)
	err := c.cc.Invoke(ctx, RankService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankServiceClient) GetFatRate(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalRank, error) {
	out := new(PersonalRank)
	err := c.cc.Invoke(ctx, RankService_GetFatRate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankServiceClient) GetTop(ctx context.Context, in *Null, opts ...grpc.CallOption) (*Ranks, error) {
	out := new(Ranks)
	err := c.cc.Invoke(ctx, RankService_GetTop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankServiceClient) RegisterPersons(ctx context.Context, opts ...grpc.CallOption) (RankService_RegisterPersonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RankService_ServiceDesc.Streams[0], RankService_RegisterPersons_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rankServiceRegisterPersonsClient{stream}
	return x, nil
}

type RankService_RegisterPersonsClient interface {
	Send(*PersonalInformation) error
	CloseAndRecv() (*PersonalInformationList, error)
	grpc.ClientStream
}

type rankServiceRegisterPersonsClient struct {
	grpc.ClientStream
}

func (x *rankServiceRegisterPersonsClient) Send(m *PersonalInformation) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rankServiceRegisterPersonsClient) CloseAndRecv() (*PersonalInformationList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PersonalInformationList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankServiceClient) WatchPersons(ctx context.Context, in *Null, opts ...grpc.CallOption) (RankService_WatchPersonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &RankService_ServiceDesc.Streams[1], RankService_WatchPersons_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rankServiceWatchPersonsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RankService_WatchPersonsClient interface {
	Recv() (*PersonalInformation, error)
	grpc.ClientStream
}

type rankServiceWatchPersonsClient struct {
	grpc.ClientStream
}

func (x *rankServiceWatchPersonsClient) Recv() (*PersonalInformation, error) {
	m := new(PersonalInformation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RankServiceServer is the server API for RankService service.
// All implementations must embed UnimplementedRankServiceServer
// for forward compatibility
type RankServiceServer interface {
	Register(context.Context, *PersonalInformation) (*PersonalInformation, error)
	Update(context.Context, *PersonalInformation) (*PersonalInformationFatRate, error)
	GetFatRate(context.Context, *PersonalInformation) (*PersonalRank, error)
	GetTop(context.Context, *Null) (*Ranks, error)
	RegisterPersons(RankService_RegisterPersonsServer) error
	WatchPersons(*Null, RankService_WatchPersonsServer) error
}

// UnimplementedRankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRankServiceServer struct {
}

func (UnimplementedRankServiceServer) Register(context.Context, *PersonalInformation) (*PersonalInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRankServiceServer) Update(context.Context, *PersonalInformation) (*PersonalInformationFatRate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRankServiceServer) GetFatRate(context.Context, *PersonalInformation) (*PersonalRank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFatRate not implemented")
}
func (UnimplementedRankServiceServer) GetTop(context.Context, *Null) (*Ranks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTop not implemented")
}
func (UnimplementedRankServiceServer) RegisterPersons(RankService_RegisterPersonsServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterPersons not implemented")
}
func (UnimplementedRankServiceServer) WatchPersons(*Null, RankService_WatchPersonsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPersons not implemented")
}
func (UnimplementedRankServiceServer) mustEmbedUnimplementedRankServiceServer() {}

// UnsafeRankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankServiceServer will
// result in compilation errors.
type UnsafeRankServiceServer interface {
	mustEmbedUnimplementedRankServiceServer()
}

func RegisterRankServiceServer(s grpc.ServiceRegistrar, srv RankServiceServer) {
	s.RegisterService(&RankService_ServiceDesc, srv)
}

func _RankService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonalInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RankService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).Register(ctx, req.(*PersonalInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonalInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RankService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).Update(ctx, req.(*PersonalInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankService_GetFatRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonalInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).GetFatRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RankService_GetFatRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).GetFatRate(ctx, req.(*PersonalInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankService_GetTop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Null)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).GetTop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RankService_GetTop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).GetTop(ctx, req.(*Null))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankService_RegisterPersons_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RankServiceServer).RegisterPersons(&rankServiceRegisterPersonsServer{stream})
}

type RankService_RegisterPersonsServer interface {
	SendAndClose(*PersonalInformationList) error
	Recv() (*PersonalInformation, error)
	grpc.ServerStream
}

type rankServiceRegisterPersonsServer struct {
	grpc.ServerStream
}

func (x *rankServiceRegisterPersonsServer) SendAndClose(m *PersonalInformationList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rankServiceRegisterPersonsServer) Recv() (*PersonalInformation, error) {
	m := new(PersonalInformation)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RankService_WatchPersons_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Null)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServiceServer).WatchPersons(m, &rankServiceWatchPersonsServer{stream})
}

type RankService_WatchPersonsServer interface {
	Send(*PersonalInformation) error
	grpc.ServerStream
}

type rankServiceWatchPersonsServer struct {
	grpc.ServerStream
}

func (x *rankServiceWatchPersonsServer) Send(m *PersonalInformation) error {
	return x.ServerStream.SendMsg(m)
}

// RankService_ServiceDesc is the grpc.ServiceDesc for RankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.RankService",
	HandlerType: (*RankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RankService_Register_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RankService_Update_Handler,
		},
		{
			MethodName: "GetFatRate",
			Handler:    _RankService_GetFatRate_Handler,
		},
		{
			MethodName: "GetTop",
			Handler:    _RankService_GetTop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterPersons",
			Handler:       _RankService_RegisterPersons_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "WatchPersons",
			Handler:       _RankService_WatchPersons_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "types.proto",
}