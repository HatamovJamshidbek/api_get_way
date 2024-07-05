// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: collaboration.proto

package genproto

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

// CollaborationServiceClient is the client API for CollaborationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollaborationServiceClient interface {
	CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*Void, error)
	UpdateInvite(ctx context.Context, in *UpdateInviteRequest, opts ...grpc.CallOption) (*Void, error)
	CreateCollaborators(ctx context.Context, in *CreateCollaborationRequest, opts ...grpc.CallOption) (*Void, error)
	GetCollaborators(ctx context.Context, in *GetCollaboratorsRequest, opts ...grpc.CallOption) (*CollaborationsResponse, error)
	UpdateCollaborators(ctx context.Context, in *UpdateCollaborationRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteCollaborators(ctx context.Context, in *DeleteCollaborationRequest, opts ...grpc.CallOption) (*Void, error)
	CreateComment(ctx context.Context, in *CreateCommitRequest, opts ...grpc.CallOption) (*Void, error)
	GetComment(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*CommitsResponse, error)
}

type collaborationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollaborationServiceClient(cc grpc.ClientConnInterface) CollaborationServiceClient {
	return &collaborationServiceClient{cc}
}

func (c *collaborationServiceClient) CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/CreateInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationServiceClient) UpdateInvite(ctx context.Context, in *UpdateInviteRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/UpdateInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationServiceClient) CreateCollaborators(ctx context.Context, in *CreateCollaborationRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/CreateCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationServiceClient) GetCollaborators(ctx context.Context, in *GetCollaboratorsRequest, opts ...grpc.CallOption) (*CollaborationsResponse, error) {
	out := new(CollaborationsResponse)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/GetCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationServiceClient) UpdateCollaborators(ctx context.Context, in *UpdateCollaborationRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/UpdateCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationServiceClient) DeleteCollaborators(ctx context.Context, in *DeleteCollaborationRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/DeleteCollaborators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationServiceClient) CreateComment(ctx context.Context, in *CreateCommitRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationServiceClient) GetComment(ctx context.Context, in *GetCommitRequest, opts ...grpc.CallOption) (*CommitsResponse, error) {
	out := new(CommitsResponse)
	err := c.cc.Invoke(ctx, "/protos.CollaborationService/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollaborationServiceServer is the server API for CollaborationService service.
// All implementations must embed UnimplementedCollaborationServiceServer
// for forward compatibility
type CollaborationServiceServer interface {
	CreateInvite(context.Context, *CreateInviteRequest) (*Void, error)
	UpdateInvite(context.Context, *UpdateInviteRequest) (*Void, error)
	CreateCollaborators(context.Context, *CreateCollaborationRequest) (*Void, error)
	GetCollaborators(context.Context, *GetCollaboratorsRequest) (*CollaborationsResponse, error)
	UpdateCollaborators(context.Context, *UpdateCollaborationRequest) (*Void, error)
	DeleteCollaborators(context.Context, *DeleteCollaborationRequest) (*Void, error)
	CreateComment(context.Context, *CreateCommitRequest) (*Void, error)
	GetComment(context.Context, *GetCommitRequest) (*CommitsResponse, error)
	mustEmbedUnimplementedCollaborationServiceServer()
}

// UnimplementedCollaborationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollaborationServiceServer struct {
}

func (UnimplementedCollaborationServiceServer) CreateInvite(context.Context, *CreateInviteRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvite not implemented")
}
func (UnimplementedCollaborationServiceServer) UpdateInvite(context.Context, *UpdateInviteRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInvite not implemented")
}
func (UnimplementedCollaborationServiceServer) CreateCollaborators(context.Context, *CreateCollaborationRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollaborators not implemented")
}
func (UnimplementedCollaborationServiceServer) GetCollaborators(context.Context, *GetCollaboratorsRequest) (*CollaborationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollaborators not implemented")
}
func (UnimplementedCollaborationServiceServer) UpdateCollaborators(context.Context, *UpdateCollaborationRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollaborators not implemented")
}
func (UnimplementedCollaborationServiceServer) DeleteCollaborators(context.Context, *DeleteCollaborationRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollaborators not implemented")
}
func (UnimplementedCollaborationServiceServer) CreateComment(context.Context, *CreateCommitRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCollaborationServiceServer) GetComment(context.Context, *GetCommitRequest) (*CommitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedCollaborationServiceServer) mustEmbedUnimplementedCollaborationServiceServer() {}

// UnsafeCollaborationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollaborationServiceServer will
// result in compilation errors.
type UnsafeCollaborationServiceServer interface {
	mustEmbedUnimplementedCollaborationServiceServer()
}

func RegisterCollaborationServiceServer(s grpc.ServiceRegistrar, srv CollaborationServiceServer) {
	s.RegisterService(&CollaborationService_ServiceDesc, srv)
}

func _CollaborationService_CreateInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).CreateInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/CreateInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).CreateInvite(ctx, req.(*CreateInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationService_UpdateInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).UpdateInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/UpdateInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).UpdateInvite(ctx, req.(*UpdateInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationService_CreateCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollaborationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).CreateCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/CreateCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).CreateCollaborators(ctx, req.(*CreateCollaborationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationService_GetCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollaboratorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).GetCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/GetCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).GetCollaborators(ctx, req.(*GetCollaboratorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationService_UpdateCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollaborationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).UpdateCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/UpdateCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).UpdateCollaborators(ctx, req.(*UpdateCollaborationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationService_DeleteCollaborators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCollaborationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).DeleteCollaborators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/DeleteCollaborators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).DeleteCollaborators(ctx, req.(*DeleteCollaborationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).CreateComment(ctx, req.(*CreateCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollaborationService_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollaborationServiceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.CollaborationService/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollaborationServiceServer).GetComment(ctx, req.(*GetCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollaborationService_ServiceDesc is the grpc.ServiceDesc for CollaborationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollaborationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CollaborationService",
	HandlerType: (*CollaborationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvite",
			Handler:    _CollaborationService_CreateInvite_Handler,
		},
		{
			MethodName: "UpdateInvite",
			Handler:    _CollaborationService_UpdateInvite_Handler,
		},
		{
			MethodName: "CreateCollaborators",
			Handler:    _CollaborationService_CreateCollaborators_Handler,
		},
		{
			MethodName: "GetCollaborators",
			Handler:    _CollaborationService_GetCollaborators_Handler,
		},
		{
			MethodName: "UpdateCollaborators",
			Handler:    _CollaborationService_UpdateCollaborators_Handler,
		},
		{
			MethodName: "DeleteCollaborators",
			Handler:    _CollaborationService_DeleteCollaborators_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _CollaborationService_CreateComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _CollaborationService_GetComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collaboration.proto",
}
