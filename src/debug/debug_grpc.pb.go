// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: debug/debug.proto

package debug

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Debug_Profile_FullMethodName               = "/debug_v2.Debug/Profile"
	Debug_Binary_FullMethodName                = "/debug_v2.Debug/Binary"
	Debug_Dump_FullMethodName                  = "/debug_v2.Debug/Dump"
	Debug_SetLogLevel_FullMethodName           = "/debug_v2.Debug/SetLogLevel"
	Debug_GetDumpV2Template_FullMethodName     = "/debug_v2.Debug/GetDumpV2Template"
	Debug_DumpV2_FullMethodName                = "/debug_v2.Debug/DumpV2"
	Debug_RunPFSLoadTest_FullMethodName        = "/debug_v2.Debug/RunPFSLoadTest"
	Debug_RunPFSLoadTestDefault_FullMethodName = "/debug_v2.Debug/RunPFSLoadTestDefault"
)

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DebugClient interface {
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error)
	Binary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (Debug_BinaryClient, error)
	Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Debug_DumpClient, error)
	SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error)
	GetDumpV2Template(ctx context.Context, in *GetDumpV2TemplateRequest, opts ...grpc.CallOption) (*GetDumpV2TemplateResponse, error)
	DumpV2(ctx context.Context, in *DumpV2Request, opts ...grpc.CallOption) (Debug_DumpV2Client, error)
	// RunLoadTest runs a load test.
	RunPFSLoadTest(ctx context.Context, in *RunPFSLoadTestRequest, opts ...grpc.CallOption) (*RunPFSLoadTestResponse, error)
	// RunLoadTestDefault runs the default load tests.
	RunPFSLoadTestDefault(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RunPFSLoadTestResponse, error)
}

type debugClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugClient(cc grpc.ClientConnInterface) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Debug_ServiceDesc.Streams[0], Debug_Profile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &debugProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_ProfileClient interface {
	Recv() (*wrapperspb.BytesValue, error)
	grpc.ClientStream
}

type debugProfileClient struct {
	grpc.ClientStream
}

func (x *debugProfileClient) Recv() (*wrapperspb.BytesValue, error) {
	m := new(wrapperspb.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Binary(ctx context.Context, in *BinaryRequest, opts ...grpc.CallOption) (Debug_BinaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Debug_ServiceDesc.Streams[1], Debug_Binary_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &debugBinaryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_BinaryClient interface {
	Recv() (*wrapperspb.BytesValue, error)
	grpc.ClientStream
}

type debugBinaryClient struct {
	grpc.ClientStream
}

func (x *debugBinaryClient) Recv() (*wrapperspb.BytesValue, error) {
	m := new(wrapperspb.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Debug_DumpClient, error) {
	stream, err := c.cc.NewStream(ctx, &Debug_ServiceDesc.Streams[2], Debug_Dump_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &debugDumpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_DumpClient interface {
	Recv() (*wrapperspb.BytesValue, error)
	grpc.ClientStream
}

type debugDumpClient struct {
	grpc.ClientStream
}

func (x *debugDumpClient) Recv() (*wrapperspb.BytesValue, error) {
	m := new(wrapperspb.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error) {
	out := new(SetLogLevelResponse)
	err := c.cc.Invoke(ctx, Debug_SetLogLevel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) GetDumpV2Template(ctx context.Context, in *GetDumpV2TemplateRequest, opts ...grpc.CallOption) (*GetDumpV2TemplateResponse, error) {
	out := new(GetDumpV2TemplateResponse)
	err := c.cc.Invoke(ctx, Debug_GetDumpV2Template_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) DumpV2(ctx context.Context, in *DumpV2Request, opts ...grpc.CallOption) (Debug_DumpV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &Debug_ServiceDesc.Streams[3], Debug_DumpV2_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &debugDumpV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_DumpV2Client interface {
	Recv() (*DumpChunk, error)
	grpc.ClientStream
}

type debugDumpV2Client struct {
	grpc.ClientStream
}

func (x *debugDumpV2Client) Recv() (*DumpChunk, error) {
	m := new(DumpChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) RunPFSLoadTest(ctx context.Context, in *RunPFSLoadTestRequest, opts ...grpc.CallOption) (*RunPFSLoadTestResponse, error) {
	out := new(RunPFSLoadTestResponse)
	err := c.cc.Invoke(ctx, Debug_RunPFSLoadTest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) RunPFSLoadTestDefault(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RunPFSLoadTestResponse, error) {
	out := new(RunPFSLoadTestResponse)
	err := c.cc.Invoke(ctx, Debug_RunPFSLoadTestDefault_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServer is the server API for Debug service.
// All implementations must embed UnimplementedDebugServer
// for forward compatibility
type DebugServer interface {
	Profile(*ProfileRequest, Debug_ProfileServer) error
	Binary(*BinaryRequest, Debug_BinaryServer) error
	Dump(*DumpRequest, Debug_DumpServer) error
	SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error)
	GetDumpV2Template(context.Context, *GetDumpV2TemplateRequest) (*GetDumpV2TemplateResponse, error)
	DumpV2(*DumpV2Request, Debug_DumpV2Server) error
	// RunLoadTest runs a load test.
	RunPFSLoadTest(context.Context, *RunPFSLoadTestRequest) (*RunPFSLoadTestResponse, error)
	// RunLoadTestDefault runs the default load tests.
	RunPFSLoadTestDefault(context.Context, *emptypb.Empty) (*RunPFSLoadTestResponse, error)
	mustEmbedUnimplementedDebugServer()
}

// UnimplementedDebugServer must be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (UnimplementedDebugServer) Profile(*ProfileRequest, Debug_ProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedDebugServer) Binary(*BinaryRequest, Debug_BinaryServer) error {
	return status.Errorf(codes.Unimplemented, "method Binary not implemented")
}
func (UnimplementedDebugServer) Dump(*DumpRequest, Debug_DumpServer) error {
	return status.Errorf(codes.Unimplemented, "method Dump not implemented")
}
func (UnimplementedDebugServer) SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}
func (UnimplementedDebugServer) GetDumpV2Template(context.Context, *GetDumpV2TemplateRequest) (*GetDumpV2TemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDumpV2Template not implemented")
}
func (UnimplementedDebugServer) DumpV2(*DumpV2Request, Debug_DumpV2Server) error {
	return status.Errorf(codes.Unimplemented, "method DumpV2 not implemented")
}
func (UnimplementedDebugServer) RunPFSLoadTest(context.Context, *RunPFSLoadTestRequest) (*RunPFSLoadTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPFSLoadTest not implemented")
}
func (UnimplementedDebugServer) RunPFSLoadTestDefault(context.Context, *emptypb.Empty) (*RunPFSLoadTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPFSLoadTestDefault not implemented")
}
func (UnimplementedDebugServer) mustEmbedUnimplementedDebugServer() {}

// UnsafeDebugServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugServer will
// result in compilation errors.
type UnsafeDebugServer interface {
	mustEmbedUnimplementedDebugServer()
}

func RegisterDebugServer(s grpc.ServiceRegistrar, srv DebugServer) {
	s.RegisterService(&Debug_ServiceDesc, srv)
}

func _Debug_Profile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProfileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Profile(m, &debugProfileServer{stream})
}

type Debug_ProfileServer interface {
	Send(*wrapperspb.BytesValue) error
	grpc.ServerStream
}

type debugProfileServer struct {
	grpc.ServerStream
}

func (x *debugProfileServer) Send(m *wrapperspb.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Binary_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BinaryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Binary(m, &debugBinaryServer{stream})
}

type Debug_BinaryServer interface {
	Send(*wrapperspb.BytesValue) error
	grpc.ServerStream
}

type debugBinaryServer struct {
	grpc.ServerStream
}

func (x *debugBinaryServer) Send(m *wrapperspb.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Dump_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Dump(m, &debugDumpServer{stream})
}

type Debug_DumpServer interface {
	Send(*wrapperspb.BytesValue) error
	grpc.ServerStream
}

type debugDumpServer struct {
	grpc.ServerStream
}

func (x *debugDumpServer) Send(m *wrapperspb.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Debug_SetLogLevel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).SetLogLevel(ctx, req.(*SetLogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_GetDumpV2Template_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDumpV2TemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).GetDumpV2Template(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Debug_GetDumpV2Template_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).GetDumpV2Template(ctx, req.(*GetDumpV2TemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_DumpV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpV2Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).DumpV2(m, &debugDumpV2Server{stream})
}

type Debug_DumpV2Server interface {
	Send(*DumpChunk) error
	grpc.ServerStream
}

type debugDumpV2Server struct {
	grpc.ServerStream
}

func (x *debugDumpV2Server) Send(m *DumpChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_RunPFSLoadTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunPFSLoadTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).RunPFSLoadTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Debug_RunPFSLoadTest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).RunPFSLoadTest(ctx, req.(*RunPFSLoadTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_RunPFSLoadTestDefault_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).RunPFSLoadTestDefault(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Debug_RunPFSLoadTestDefault_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).RunPFSLoadTestDefault(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Debug_ServiceDesc is the grpc.ServiceDesc for Debug service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Debug_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "debug_v2.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLogLevel",
			Handler:    _Debug_SetLogLevel_Handler,
		},
		{
			MethodName: "GetDumpV2Template",
			Handler:    _Debug_GetDumpV2Template_Handler,
		},
		{
			MethodName: "RunPFSLoadTest",
			Handler:    _Debug_RunPFSLoadTest_Handler,
		},
		{
			MethodName: "RunPFSLoadTestDefault",
			Handler:    _Debug_RunPFSLoadTestDefault_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Profile",
			Handler:       _Debug_Profile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Binary",
			Handler:       _Debug_Binary_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Dump",
			Handler:       _Debug_Dump_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DumpV2",
			Handler:       _Debug_DumpV2_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "debug/debug.proto",
}
