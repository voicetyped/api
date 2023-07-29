// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: transcriber.proto

package transcriberv1

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
	TranscriberService_Transcribe_FullMethodName = "/apis.TranscriberService/Transcribe"
	TranscriberService_Prompt_FullMethodName     = "/apis.TranscriberService/Prompt"
)

// TranscriberServiceClient is the client API for TranscriberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranscriberServiceClient interface {
	// Performs bidirectional streaming speech recognition: receive results while
	// sending audio.
	Transcribe(ctx context.Context, opts ...grpc.CallOption) (TranscriberService_TranscribeClient, error)
	Prompt(ctx context.Context, opts ...grpc.CallOption) (TranscriberService_PromptClient, error)
}

type transcriberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranscriberServiceClient(cc grpc.ClientConnInterface) TranscriberServiceClient {
	return &transcriberServiceClient{cc}
}

func (c *transcriberServiceClient) Transcribe(ctx context.Context, opts ...grpc.CallOption) (TranscriberService_TranscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &TranscriberService_ServiceDesc.Streams[0], TranscriberService_Transcribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transcriberServiceTranscribeClient{stream}
	return x, nil
}

type TranscriberService_TranscribeClient interface {
	Send(*TranscribeRequest) error
	Recv() (*TranscribeResponse, error)
	grpc.ClientStream
}

type transcriberServiceTranscribeClient struct {
	grpc.ClientStream
}

func (x *transcriberServiceTranscribeClient) Send(m *TranscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transcriberServiceTranscribeClient) Recv() (*TranscribeResponse, error) {
	m := new(TranscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transcriberServiceClient) Prompt(ctx context.Context, opts ...grpc.CallOption) (TranscriberService_PromptClient, error) {
	stream, err := c.cc.NewStream(ctx, &TranscriberService_ServiceDesc.Streams[1], TranscriberService_Prompt_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transcriberServicePromptClient{stream}
	return x, nil
}

type TranscriberService_PromptClient interface {
	Send(*QRequest) error
	CloseAndRecv() (*QResponse, error)
	grpc.ClientStream
}

type transcriberServicePromptClient struct {
	grpc.ClientStream
}

func (x *transcriberServicePromptClient) Send(m *QRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transcriberServicePromptClient) CloseAndRecv() (*QResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(QResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TranscriberServiceServer is the server API for TranscriberService service.
// All implementations must embed UnimplementedTranscriberServiceServer
// for forward compatibility
type TranscriberServiceServer interface {
	// Performs bidirectional streaming speech recognition: receive results while
	// sending audio.
	Transcribe(TranscriberService_TranscribeServer) error
	Prompt(TranscriberService_PromptServer) error
	mustEmbedUnimplementedTranscriberServiceServer()
}

// UnimplementedTranscriberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTranscriberServiceServer struct {
}

func (UnimplementedTranscriberServiceServer) Transcribe(TranscriberService_TranscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Transcribe not implemented")
}
func (UnimplementedTranscriberServiceServer) Prompt(TranscriberService_PromptServer) error {
	return status.Errorf(codes.Unimplemented, "method Prompt not implemented")
}
func (UnimplementedTranscriberServiceServer) mustEmbedUnimplementedTranscriberServiceServer() {}

// UnsafeTranscriberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranscriberServiceServer will
// result in compilation errors.
type UnsafeTranscriberServiceServer interface {
	mustEmbedUnimplementedTranscriberServiceServer()
}

func RegisterTranscriberServiceServer(s grpc.ServiceRegistrar, srv TranscriberServiceServer) {
	s.RegisterService(&TranscriberService_ServiceDesc, srv)
}

func _TranscriberService_Transcribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TranscriberServiceServer).Transcribe(&transcriberServiceTranscribeServer{stream})
}

type TranscriberService_TranscribeServer interface {
	Send(*TranscribeResponse) error
	Recv() (*TranscribeRequest, error)
	grpc.ServerStream
}

type transcriberServiceTranscribeServer struct {
	grpc.ServerStream
}

func (x *transcriberServiceTranscribeServer) Send(m *TranscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transcriberServiceTranscribeServer) Recv() (*TranscribeRequest, error) {
	m := new(TranscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TranscriberService_Prompt_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TranscriberServiceServer).Prompt(&transcriberServicePromptServer{stream})
}

type TranscriberService_PromptServer interface {
	SendAndClose(*QResponse) error
	Recv() (*QRequest, error)
	grpc.ServerStream
}

type transcriberServicePromptServer struct {
	grpc.ServerStream
}

func (x *transcriberServicePromptServer) SendAndClose(m *QResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transcriberServicePromptServer) Recv() (*QRequest, error) {
	m := new(QRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TranscriberService_ServiceDesc is the grpc.ServiceDesc for TranscriberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranscriberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apis.TranscriberService",
	HandlerType: (*TranscriberServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transcribe",
			Handler:       _TranscriberService_Transcribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Prompt",
			Handler:       _TranscriberService_Prompt_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "transcriber.proto",
}
