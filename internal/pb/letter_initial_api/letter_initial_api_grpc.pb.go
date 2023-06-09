// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/letter_initial_api/letter_initial_api.proto

package letter_initial_api

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

// LetterInitialServiceClient is the client API for LetterInitialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LetterInitialServiceClient interface {
	AddLetter(ctx context.Context, in *AddLetterRequest, opts ...grpc.CallOption) (*AddLetterResponse, error)
	GetLetter(ctx context.Context, in *GetLetterRequest, opts ...grpc.CallOption) (*GetLetterResponse, error)
	GetFullLetter(ctx context.Context, in *GetFullLetterRequest, opts ...grpc.CallOption) (*GetFullLetterResponse, error)
	RenameLetter(ctx context.Context, in *RenameLetterRequest, opts ...grpc.CallOption) (*RenameLetterResponse, error)
	DeleteLetter(ctx context.Context, in *DeleteLetterRequest, opts ...grpc.CallOption) (*DeleteLetterResponse, error)
	AddObjective(ctx context.Context, in *AddObjectiveRequest, opts ...grpc.CallOption) (*AddObjectiveResponse, error)
	GetObjective(ctx context.Context, in *GetObjectiveRequest, opts ...grpc.CallOption) (*GetObjectiveResponse, error)
	GetObjectiveList(ctx context.Context, in *GetObjectiveListRequest, opts ...grpc.CallOption) (*GetObjectiveListResponse, error)
	RenameObjective(ctx context.Context, in *RenameObjectiveRequest, opts ...grpc.CallOption) (*RenameObjectiveResponse, error)
	ChangeObjectiveOrderNumber(ctx context.Context, in *ChangeObjectiveOrderNumberRequest, opts ...grpc.CallOption) (*ChangeObjectiveOrderNumberResponse, error)
	DeleteObjective(ctx context.Context, in *DeleteObjectiveRequest, opts ...grpc.CallOption) (*DeleteObjectiveResponse, error)
	AddPurpose(ctx context.Context, in *AddPurposeRequest, opts ...grpc.CallOption) (*AddPurposeResponse, error)
	GetPurpose(ctx context.Context, in *GetPurposeRequest, opts ...grpc.CallOption) (*GetPurposeResponse, error)
	GetPurposeList(ctx context.Context, in *GetPurposeListRequest, opts ...grpc.CallOption) (*GetPurposeListResponse, error)
	RenamePurpose(ctx context.Context, in *RenamePurposeRequest, opts ...grpc.CallOption) (*RenamePurposeResponse, error)
	ChangePurposeOrderNumber(ctx context.Context, in *ChangePurposeOrderNumberRequest, opts ...grpc.CallOption) (*ChangePurposeOrderNumberResponse, error)
	DeletePurpose(ctx context.Context, in *DeletePurposeRequest, opts ...grpc.CallOption) (*DeletePurposeResponse, error)
	GenerateLetter(ctx context.Context, in *GenerateLetterRequest, opts ...grpc.CallOption) (*GenerateLetterResponse, error)
}

type letterInitialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLetterInitialServiceClient(cc grpc.ClientConnInterface) LetterInitialServiceClient {
	return &letterInitialServiceClient{cc}
}

func (c *letterInitialServiceClient) AddLetter(ctx context.Context, in *AddLetterRequest, opts ...grpc.CallOption) (*AddLetterResponse, error) {
	out := new(AddLetterResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/AddLetter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) GetLetter(ctx context.Context, in *GetLetterRequest, opts ...grpc.CallOption) (*GetLetterResponse, error) {
	out := new(GetLetterResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/GetLetter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) GetFullLetter(ctx context.Context, in *GetFullLetterRequest, opts ...grpc.CallOption) (*GetFullLetterResponse, error) {
	out := new(GetFullLetterResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/GetFullLetter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) RenameLetter(ctx context.Context, in *RenameLetterRequest, opts ...grpc.CallOption) (*RenameLetterResponse, error) {
	out := new(RenameLetterResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/RenameLetter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) DeleteLetter(ctx context.Context, in *DeleteLetterRequest, opts ...grpc.CallOption) (*DeleteLetterResponse, error) {
	out := new(DeleteLetterResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/DeleteLetter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) AddObjective(ctx context.Context, in *AddObjectiveRequest, opts ...grpc.CallOption) (*AddObjectiveResponse, error) {
	out := new(AddObjectiveResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/AddObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) GetObjective(ctx context.Context, in *GetObjectiveRequest, opts ...grpc.CallOption) (*GetObjectiveResponse, error) {
	out := new(GetObjectiveResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/GetObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) GetObjectiveList(ctx context.Context, in *GetObjectiveListRequest, opts ...grpc.CallOption) (*GetObjectiveListResponse, error) {
	out := new(GetObjectiveListResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/GetObjectiveList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) RenameObjective(ctx context.Context, in *RenameObjectiveRequest, opts ...grpc.CallOption) (*RenameObjectiveResponse, error) {
	out := new(RenameObjectiveResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/RenameObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) ChangeObjectiveOrderNumber(ctx context.Context, in *ChangeObjectiveOrderNumberRequest, opts ...grpc.CallOption) (*ChangeObjectiveOrderNumberResponse, error) {
	out := new(ChangeObjectiveOrderNumberResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/ChangeObjectiveOrderNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) DeleteObjective(ctx context.Context, in *DeleteObjectiveRequest, opts ...grpc.CallOption) (*DeleteObjectiveResponse, error) {
	out := new(DeleteObjectiveResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/DeleteObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) AddPurpose(ctx context.Context, in *AddPurposeRequest, opts ...grpc.CallOption) (*AddPurposeResponse, error) {
	out := new(AddPurposeResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/AddPurpose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) GetPurpose(ctx context.Context, in *GetPurposeRequest, opts ...grpc.CallOption) (*GetPurposeResponse, error) {
	out := new(GetPurposeResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/GetPurpose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) GetPurposeList(ctx context.Context, in *GetPurposeListRequest, opts ...grpc.CallOption) (*GetPurposeListResponse, error) {
	out := new(GetPurposeListResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/GetPurposeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) RenamePurpose(ctx context.Context, in *RenamePurposeRequest, opts ...grpc.CallOption) (*RenamePurposeResponse, error) {
	out := new(RenamePurposeResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/RenamePurpose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) ChangePurposeOrderNumber(ctx context.Context, in *ChangePurposeOrderNumberRequest, opts ...grpc.CallOption) (*ChangePurposeOrderNumberResponse, error) {
	out := new(ChangePurposeOrderNumberResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/ChangePurposeOrderNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) DeletePurpose(ctx context.Context, in *DeletePurposeRequest, opts ...grpc.CallOption) (*DeletePurposeResponse, error) {
	out := new(DeletePurposeResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/DeletePurpose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *letterInitialServiceClient) GenerateLetter(ctx context.Context, in *GenerateLetterRequest, opts ...grpc.CallOption) (*GenerateLetterResponse, error) {
	out := new(GenerateLetterResponse)
	err := c.cc.Invoke(ctx, "/letter_initial_api.LetterInitialService/GenerateLetter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LetterInitialServiceServer is the server API for LetterInitialService service.
// All implementations must embed UnimplementedLetterInitialServiceServer
// for forward compatibility
type LetterInitialServiceServer interface {
	AddLetter(context.Context, *AddLetterRequest) (*AddLetterResponse, error)
	GetLetter(context.Context, *GetLetterRequest) (*GetLetterResponse, error)
	GetFullLetter(context.Context, *GetFullLetterRequest) (*GetFullLetterResponse, error)
	RenameLetter(context.Context, *RenameLetterRequest) (*RenameLetterResponse, error)
	DeleteLetter(context.Context, *DeleteLetterRequest) (*DeleteLetterResponse, error)
	AddObjective(context.Context, *AddObjectiveRequest) (*AddObjectiveResponse, error)
	GetObjective(context.Context, *GetObjectiveRequest) (*GetObjectiveResponse, error)
	GetObjectiveList(context.Context, *GetObjectiveListRequest) (*GetObjectiveListResponse, error)
	RenameObjective(context.Context, *RenameObjectiveRequest) (*RenameObjectiveResponse, error)
	ChangeObjectiveOrderNumber(context.Context, *ChangeObjectiveOrderNumberRequest) (*ChangeObjectiveOrderNumberResponse, error)
	DeleteObjective(context.Context, *DeleteObjectiveRequest) (*DeleteObjectiveResponse, error)
	AddPurpose(context.Context, *AddPurposeRequest) (*AddPurposeResponse, error)
	GetPurpose(context.Context, *GetPurposeRequest) (*GetPurposeResponse, error)
	GetPurposeList(context.Context, *GetPurposeListRequest) (*GetPurposeListResponse, error)
	RenamePurpose(context.Context, *RenamePurposeRequest) (*RenamePurposeResponse, error)
	ChangePurposeOrderNumber(context.Context, *ChangePurposeOrderNumberRequest) (*ChangePurposeOrderNumberResponse, error)
	DeletePurpose(context.Context, *DeletePurposeRequest) (*DeletePurposeResponse, error)
	GenerateLetter(context.Context, *GenerateLetterRequest) (*GenerateLetterResponse, error)
	mustEmbedUnimplementedLetterInitialServiceServer()
}

// UnimplementedLetterInitialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLetterInitialServiceServer struct {
}

func (UnimplementedLetterInitialServiceServer) AddLetter(context.Context, *AddLetterRequest) (*AddLetterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLetter not implemented")
}
func (UnimplementedLetterInitialServiceServer) GetLetter(context.Context, *GetLetterRequest) (*GetLetterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLetter not implemented")
}
func (UnimplementedLetterInitialServiceServer) GetFullLetter(context.Context, *GetFullLetterRequest) (*GetFullLetterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullLetter not implemented")
}
func (UnimplementedLetterInitialServiceServer) RenameLetter(context.Context, *RenameLetterRequest) (*RenameLetterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameLetter not implemented")
}
func (UnimplementedLetterInitialServiceServer) DeleteLetter(context.Context, *DeleteLetterRequest) (*DeleteLetterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLetter not implemented")
}
func (UnimplementedLetterInitialServiceServer) AddObjective(context.Context, *AddObjectiveRequest) (*AddObjectiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddObjective not implemented")
}
func (UnimplementedLetterInitialServiceServer) GetObjective(context.Context, *GetObjectiveRequest) (*GetObjectiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjective not implemented")
}
func (UnimplementedLetterInitialServiceServer) GetObjectiveList(context.Context, *GetObjectiveListRequest) (*GetObjectiveListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectiveList not implemented")
}
func (UnimplementedLetterInitialServiceServer) RenameObjective(context.Context, *RenameObjectiveRequest) (*RenameObjectiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameObjective not implemented")
}
func (UnimplementedLetterInitialServiceServer) ChangeObjectiveOrderNumber(context.Context, *ChangeObjectiveOrderNumberRequest) (*ChangeObjectiveOrderNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeObjectiveOrderNumber not implemented")
}
func (UnimplementedLetterInitialServiceServer) DeleteObjective(context.Context, *DeleteObjectiveRequest) (*DeleteObjectiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjective not implemented")
}
func (UnimplementedLetterInitialServiceServer) AddPurpose(context.Context, *AddPurposeRequest) (*AddPurposeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPurpose not implemented")
}
func (UnimplementedLetterInitialServiceServer) GetPurpose(context.Context, *GetPurposeRequest) (*GetPurposeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPurpose not implemented")
}
func (UnimplementedLetterInitialServiceServer) GetPurposeList(context.Context, *GetPurposeListRequest) (*GetPurposeListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPurposeList not implemented")
}
func (UnimplementedLetterInitialServiceServer) RenamePurpose(context.Context, *RenamePurposeRequest) (*RenamePurposeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenamePurpose not implemented")
}
func (UnimplementedLetterInitialServiceServer) ChangePurposeOrderNumber(context.Context, *ChangePurposeOrderNumberRequest) (*ChangePurposeOrderNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePurposeOrderNumber not implemented")
}
func (UnimplementedLetterInitialServiceServer) DeletePurpose(context.Context, *DeletePurposeRequest) (*DeletePurposeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePurpose not implemented")
}
func (UnimplementedLetterInitialServiceServer) GenerateLetter(context.Context, *GenerateLetterRequest) (*GenerateLetterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateLetter not implemented")
}
func (UnimplementedLetterInitialServiceServer) mustEmbedUnimplementedLetterInitialServiceServer() {}

// UnsafeLetterInitialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LetterInitialServiceServer will
// result in compilation errors.
type UnsafeLetterInitialServiceServer interface {
	mustEmbedUnimplementedLetterInitialServiceServer()
}

func RegisterLetterInitialServiceServer(s grpc.ServiceRegistrar, srv LetterInitialServiceServer) {
	s.RegisterService(&LetterInitialService_ServiceDesc, srv)
}

func _LetterInitialService_AddLetter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).AddLetter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/AddLetter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).AddLetter(ctx, req.(*AddLetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_GetLetter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).GetLetter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/GetLetter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).GetLetter(ctx, req.(*GetLetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_GetFullLetter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullLetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).GetFullLetter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/GetFullLetter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).GetFullLetter(ctx, req.(*GetFullLetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_RenameLetter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameLetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).RenameLetter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/RenameLetter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).RenameLetter(ctx, req.(*RenameLetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_DeleteLetter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).DeleteLetter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/DeleteLetter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).DeleteLetter(ctx, req.(*DeleteLetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_AddObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).AddObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/AddObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).AddObjective(ctx, req.(*AddObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_GetObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).GetObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/GetObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).GetObjective(ctx, req.(*GetObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_GetObjectiveList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectiveListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).GetObjectiveList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/GetObjectiveList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).GetObjectiveList(ctx, req.(*GetObjectiveListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_RenameObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).RenameObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/RenameObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).RenameObjective(ctx, req.(*RenameObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_ChangeObjectiveOrderNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeObjectiveOrderNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).ChangeObjectiveOrderNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/ChangeObjectiveOrderNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).ChangeObjectiveOrderNumber(ctx, req.(*ChangeObjectiveOrderNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_DeleteObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).DeleteObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/DeleteObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).DeleteObjective(ctx, req.(*DeleteObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_AddPurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPurposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).AddPurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/AddPurpose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).AddPurpose(ctx, req.(*AddPurposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_GetPurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPurposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).GetPurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/GetPurpose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).GetPurpose(ctx, req.(*GetPurposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_GetPurposeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPurposeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).GetPurposeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/GetPurposeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).GetPurposeList(ctx, req.(*GetPurposeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_RenamePurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenamePurposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).RenamePurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/RenamePurpose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).RenamePurpose(ctx, req.(*RenamePurposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_ChangePurposeOrderNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePurposeOrderNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).ChangePurposeOrderNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/ChangePurposeOrderNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).ChangePurposeOrderNumber(ctx, req.(*ChangePurposeOrderNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_DeletePurpose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePurposeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).DeletePurpose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/DeletePurpose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).DeletePurpose(ctx, req.(*DeletePurposeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LetterInitialService_GenerateLetter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateLetterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LetterInitialServiceServer).GenerateLetter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/letter_initial_api.LetterInitialService/GenerateLetter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LetterInitialServiceServer).GenerateLetter(ctx, req.(*GenerateLetterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LetterInitialService_ServiceDesc is the grpc.ServiceDesc for LetterInitialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LetterInitialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "letter_initial_api.LetterInitialService",
	HandlerType: (*LetterInitialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLetter",
			Handler:    _LetterInitialService_AddLetter_Handler,
		},
		{
			MethodName: "GetLetter",
			Handler:    _LetterInitialService_GetLetter_Handler,
		},
		{
			MethodName: "GetFullLetter",
			Handler:    _LetterInitialService_GetFullLetter_Handler,
		},
		{
			MethodName: "RenameLetter",
			Handler:    _LetterInitialService_RenameLetter_Handler,
		},
		{
			MethodName: "DeleteLetter",
			Handler:    _LetterInitialService_DeleteLetter_Handler,
		},
		{
			MethodName: "AddObjective",
			Handler:    _LetterInitialService_AddObjective_Handler,
		},
		{
			MethodName: "GetObjective",
			Handler:    _LetterInitialService_GetObjective_Handler,
		},
		{
			MethodName: "GetObjectiveList",
			Handler:    _LetterInitialService_GetObjectiveList_Handler,
		},
		{
			MethodName: "RenameObjective",
			Handler:    _LetterInitialService_RenameObjective_Handler,
		},
		{
			MethodName: "ChangeObjectiveOrderNumber",
			Handler:    _LetterInitialService_ChangeObjectiveOrderNumber_Handler,
		},
		{
			MethodName: "DeleteObjective",
			Handler:    _LetterInitialService_DeleteObjective_Handler,
		},
		{
			MethodName: "AddPurpose",
			Handler:    _LetterInitialService_AddPurpose_Handler,
		},
		{
			MethodName: "GetPurpose",
			Handler:    _LetterInitialService_GetPurpose_Handler,
		},
		{
			MethodName: "GetPurposeList",
			Handler:    _LetterInitialService_GetPurposeList_Handler,
		},
		{
			MethodName: "RenamePurpose",
			Handler:    _LetterInitialService_RenamePurpose_Handler,
		},
		{
			MethodName: "ChangePurposeOrderNumber",
			Handler:    _LetterInitialService_ChangePurposeOrderNumber_Handler,
		},
		{
			MethodName: "DeletePurpose",
			Handler:    _LetterInitialService_DeletePurpose_Handler,
		},
		{
			MethodName: "GenerateLetter",
			Handler:    _LetterInitialService_GenerateLetter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/letter_initial_api/letter_initial_api.proto",
}
