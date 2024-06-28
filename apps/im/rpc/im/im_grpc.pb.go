// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: apps/im/rpc/im.proto

package im

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
	Im_GetChatLog_FullMethodName              = "/im.Im/GetChatLog"
	Im_SetUpUserConversation_FullMethodName   = "/im.Im/SetUpUserConversation"
	Im_GetConversations_FullMethodName        = "/im.Im/GetConversations"
	Im_PutConversations_FullMethodName        = "/im.Im/PutConversations"
	Im_CreateGroupConversation_FullMethodName = "/im.Im/CreateGroupConversation"
)

// ImClient is the client API for Im service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImClient interface {
	// 获取会话记录
	GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error)
	// 建立会话：群聊、私聊
	SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error)
	// 获取会话
	GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error)
	// 更新会话
	PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error)
	CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error)
}

type imClient struct {
	cc grpc.ClientConnInterface
}

func NewImClient(cc grpc.ClientConnInterface) ImClient {
	return &imClient{cc}
}

func (c *imClient) GetChatLog(ctx context.Context, in *GetChatLogReq, opts ...grpc.CallOption) (*GetChatLogResp, error) {
	out := new(GetChatLogResp)
	err := c.cc.Invoke(ctx, Im_GetChatLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imClient) SetUpUserConversation(ctx context.Context, in *SetUpUserConversationReq, opts ...grpc.CallOption) (*SetUpUserConversationResp, error) {
	out := new(SetUpUserConversationResp)
	err := c.cc.Invoke(ctx, Im_SetUpUserConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imClient) GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error) {
	out := new(GetConversationsResp)
	err := c.cc.Invoke(ctx, Im_GetConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imClient) PutConversations(ctx context.Context, in *PutConversationsReq, opts ...grpc.CallOption) (*PutConversationsResp, error) {
	out := new(PutConversationsResp)
	err := c.cc.Invoke(ctx, Im_PutConversations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imClient) CreateGroupConversation(ctx context.Context, in *CreateGroupConversationReq, opts ...grpc.CallOption) (*CreateGroupConversationResp, error) {
	out := new(CreateGroupConversationResp)
	err := c.cc.Invoke(ctx, Im_CreateGroupConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImServer is the server API for Im service.
// All implementations must embed UnimplementedImServer
// for forward compatibility
type ImServer interface {
	// 获取会话记录
	GetChatLog(context.Context, *GetChatLogReq) (*GetChatLogResp, error)
	// 建立会话：群聊、私聊
	SetUpUserConversation(context.Context, *SetUpUserConversationReq) (*SetUpUserConversationResp, error)
	// 获取会话
	GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error)
	// 更新会话
	PutConversations(context.Context, *PutConversationsReq) (*PutConversationsResp, error)
	CreateGroupConversation(context.Context, *CreateGroupConversationReq) (*CreateGroupConversationResp, error)
	mustEmbedUnimplementedImServer()
}

// UnimplementedImServer must be embedded to have forward compatible implementations.
type UnimplementedImServer struct {
}

func (UnimplementedImServer) GetChatLog(context.Context, *GetChatLogReq) (*GetChatLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatLog not implemented")
}
func (UnimplementedImServer) SetUpUserConversation(context.Context, *SetUpUserConversationReq) (*SetUpUserConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUpUserConversation not implemented")
}
func (UnimplementedImServer) GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversations not implemented")
}
func (UnimplementedImServer) PutConversations(context.Context, *PutConversationsReq) (*PutConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutConversations not implemented")
}
func (UnimplementedImServer) CreateGroupConversation(context.Context, *CreateGroupConversationReq) (*CreateGroupConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupConversation not implemented")
}
func (UnimplementedImServer) mustEmbedUnimplementedImServer() {}

// UnsafeImServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImServer will
// result in compilation errors.
type UnsafeImServer interface {
	mustEmbedUnimplementedImServer()
}

func RegisterImServer(s grpc.ServiceRegistrar, srv ImServer) {
	s.RegisterService(&Im_ServiceDesc, srv)
}

func _Im_GetChatLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServer).GetChatLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Im_GetChatLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServer).GetChatLog(ctx, req.(*GetChatLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Im_SetUpUserConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUpUserConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServer).SetUpUserConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Im_SetUpUserConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServer).SetUpUserConversation(ctx, req.(*SetUpUserConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Im_GetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServer).GetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Im_GetConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServer).GetConversations(ctx, req.(*GetConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Im_PutConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServer).PutConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Im_PutConversations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServer).PutConversations(ctx, req.(*PutConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Im_CreateGroupConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImServer).CreateGroupConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Im_CreateGroupConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImServer).CreateGroupConversation(ctx, req.(*CreateGroupConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Im_ServiceDesc is the grpc.ServiceDesc for Im service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Im_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "im.Im",
	HandlerType: (*ImServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChatLog",
			Handler:    _Im_GetChatLog_Handler,
		},
		{
			MethodName: "SetUpUserConversation",
			Handler:    _Im_SetUpUserConversation_Handler,
		},
		{
			MethodName: "GetConversations",
			Handler:    _Im_GetConversations_Handler,
		},
		{
			MethodName: "PutConversations",
			Handler:    _Im_PutConversations_Handler,
		},
		{
			MethodName: "CreateGroupConversation",
			Handler:    _Im_CreateGroupConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/im/rpc/im.proto",
}
