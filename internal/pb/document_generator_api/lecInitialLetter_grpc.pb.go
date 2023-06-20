// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/document_generator_api/lecInitialLetter.proto

package document_generator_api

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

// GenerateLecInitialLetterClient is the client API for GenerateLecInitialLetter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenerateLecInitialLetterClient interface {
	Generate(ctx context.Context, in *GenerateLetterRequest, opts ...grpc.CallOption) (*GenerateLetterResponse, error)
}

type generateLecInitialLetterClient struct {
	cc grpc.ClientConnInterface
}

func NewGenerateLecInitialLetterClient(cc grpc.ClientConnInterface) GenerateLecInitialLetterClient {
	return &generateLecInitialLetterClient{cc}
}

func (c *generateLecInitialLetterClient) Generate(ctx context.Context, in *GenerateLetterRequest, opts ...grpc.CallOption) (*GenerateLetterResponse, error) {
	out := new(GenerateLetterResponse)
	err := c.cc.Invoke(ctx, "/lecInitialLetter.GenerateLecInitialLetter/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateLecInitialLetterServer is the server API for GenerateLecInitialLetter service.
// All implementations must embed UnimplementedGenerateLecInitialLetterServer
// for forward compatibility
type GenerateLecInitialLetterServer interface {
	Generate(context.Context, *GenerateLetterRequest) (*GenerateLetterResponse, error)
	mustEmbedUnimplementedGenerateLecInitialLetterServer()
}

// UnimplementedGenerateLecInitialLetterServer must be embedded to have forward compatible implementations.
type UnimplementedGenerateLecInitialLetterServer struct {
}

func (UnimplementedGenerateLecInitialLetterServer) Generate(context.Context, *GenerateLetterRequest) (*GenerateLetterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedGenerateLecInitialLetterServer) mustEmbedUnimplementedGenerateLecInitialLetterServer() {
}

// UnsafeGenerateLecInitialLetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenerateLecInitialLetterServer will
// result in compilation errors.
type UnsafeGenerateLecInitialLetterServer interface {
	mustEmbedUnimplementedGenerateLecInitialLetterServer()
}

func RegisterGenerateLecInitialLetterServer(s grpc.ServiceRegistrar, srv GenerateLecInitialLetterServer) {
	s.RegisterService(&GenerateLecInitialLetter_ServiceDesc, srv)
}

func _GenerateLecInitialLetter_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateLetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateLecInitialLetterServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecInitialLetter.GenerateLecInitialLetter/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateLecInitialLetterServer).Generate(ctx, req.(*GenerateLetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenerateLecInitialLetter_ServiceDesc is the grpc.ServiceDesc for GenerateLecInitialLetter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenerateLecInitialLetter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lecInitialLetter.GenerateLecInitialLetter",
	HandlerType: (*GenerateLecInitialLetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _GenerateLecInitialLetter_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/document_generator_api/lecInitialLetter.proto",
}
