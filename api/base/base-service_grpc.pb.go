// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.14.0
// source: base-service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BaseService_LoginWeChat_FullMethodName             = "/base_service.v1.BaseService/LoginWeChat"
	BaseService_FastRegisterWeChat_FullMethodName      = "/base_service.v1.BaseService/FastRegisterWeChat"
	BaseService_LocationCommonSearch_FullMethodName    = "/base_service.v1.BaseService/LocationCommonSearch"
	BaseService_MediaPutURLBatchGet_FullMethodName     = "/base_service.v1.BaseService/MediaPutURLBatchGet"
	BaseService_ChannelTypeList_FullMethodName         = "/base_service.v1.BaseService/ChannelTypeList"
	BaseService_ChannelCreate_FullMethodName           = "/base_service.v1.BaseService/ChannelCreate"
	BaseService_ChannelUpdate_FullMethodName           = "/base_service.v1.BaseService/ChannelUpdate"
	BaseService_ChannelDelete_FullMethodName           = "/base_service.v1.BaseService/ChannelDelete"
	BaseService_ChannelBaseQueryByBound_FullMethodName = "/base_service.v1.BaseService/ChannelBaseQueryByBound"
	BaseService_ChannelFullQueryById_FullMethodName    = "/base_service.v1.BaseService/ChannelFullQueryById"
	BaseService_ChannelInx_FullMethodName              = "/base_service.v1.BaseService/ChannelInx"
	BaseService_ChannelComment_FullMethodName          = "/base_service.v1.BaseService/ChannelComment"
)

// BaseServiceClient is the client API for BaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type BaseServiceClient interface {
	// 微信小程序登录接口
	// @security BearerAuth
	LoginWeChat(ctx context.Context, in *LoginWeChatReq, opts ...grpc.CallOption) (*LoginWeChatRes, error)
	// 微信小程序快速登录注册接口定义
	FastRegisterWeChat(ctx context.Context, in *FastRegisterWeChatReq, opts ...grpc.CallOption) (*FastRegisterWeChatRes, error)
	// 通用地点搜索
	LocationCommonSearch(ctx context.Context, in *LocationCommonSearchReq, opts ...grpc.CallOption) (*LocationCommonSearchRes, error)
	// 批量获取对象上传预签名URL
	MediaPutURLBatchGet(ctx context.Context, in *MediaPutURLBatchGetReq, opts ...grpc.CallOption) (*MediaPutURLBatchGetRes, error)
	// 列表查询足迹频道类型
	ChannelTypeList(ctx context.Context, in *ChannelTypeListReq, opts ...grpc.CallOption) (*ChannelTypeListRes, error)
	// 创建足迹频道
	ChannelCreate(ctx context.Context, in *ChannelCreateReq, opts ...grpc.CallOption) (*ChannelCreateRes, error)
	// 更新足迹频道
	ChannelUpdate(ctx context.Context, in *ChannelUpdateReq, opts ...grpc.CallOption) (*ChannelUpdateRes, error)
	// 删除足迹频道
	ChannelDelete(ctx context.Context, in *ChannelDeleteReq, opts ...grpc.CallOption) (*ChannelDeleteRes, error)
	// 按照范围查询足迹基础信息
	ChannelBaseQueryByBound(ctx context.Context, in *ChannelBaseQueryByBoundReq, opts ...grpc.CallOption) (*ChannelBaseQueryByBoundRes, error)
	// 按照id查询足迹频道动态信息
	ChannelFullQueryById(ctx context.Context, in *ChannelFullQueryByIdReq, opts ...grpc.CallOption) (*ChannelFullQueryByIdRes, error)
	// 足迹频道互动
	ChannelInx(ctx context.Context, in *ChannelInxReq, opts ...grpc.CallOption) (*ChannelInxRes, error)
	// 足迹频道评论
	ChannelComment(ctx context.Context, in *ChannelCommentReq, opts ...grpc.CallOption) (*ChannelCommentRes, error)
}

type baseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseServiceClient(cc grpc.ClientConnInterface) BaseServiceClient {
	return &baseServiceClient{cc}
}

func (c *baseServiceClient) LoginWeChat(ctx context.Context, in *LoginWeChatReq, opts ...grpc.CallOption) (*LoginWeChatRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginWeChatRes)
	err := c.cc.Invoke(ctx, BaseService_LoginWeChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) FastRegisterWeChat(ctx context.Context, in *FastRegisterWeChatReq, opts ...grpc.CallOption) (*FastRegisterWeChatRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FastRegisterWeChatRes)
	err := c.cc.Invoke(ctx, BaseService_FastRegisterWeChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) LocationCommonSearch(ctx context.Context, in *LocationCommonSearchReq, opts ...grpc.CallOption) (*LocationCommonSearchRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LocationCommonSearchRes)
	err := c.cc.Invoke(ctx, BaseService_LocationCommonSearch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) MediaPutURLBatchGet(ctx context.Context, in *MediaPutURLBatchGetReq, opts ...grpc.CallOption) (*MediaPutURLBatchGetRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MediaPutURLBatchGetRes)
	err := c.cc.Invoke(ctx, BaseService_MediaPutURLBatchGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelTypeList(ctx context.Context, in *ChannelTypeListReq, opts ...grpc.CallOption) (*ChannelTypeListRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelTypeListRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelTypeList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelCreate(ctx context.Context, in *ChannelCreateReq, opts ...grpc.CallOption) (*ChannelCreateRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelCreateRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelCreate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelUpdate(ctx context.Context, in *ChannelUpdateReq, opts ...grpc.CallOption) (*ChannelUpdateRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelUpdateRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelDelete(ctx context.Context, in *ChannelDeleteReq, opts ...grpc.CallOption) (*ChannelDeleteRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelDeleteRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelBaseQueryByBound(ctx context.Context, in *ChannelBaseQueryByBoundReq, opts ...grpc.CallOption) (*ChannelBaseQueryByBoundRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelBaseQueryByBoundRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelBaseQueryByBound_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelFullQueryById(ctx context.Context, in *ChannelFullQueryByIdReq, opts ...grpc.CallOption) (*ChannelFullQueryByIdRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelFullQueryByIdRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelFullQueryById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelInx(ctx context.Context, in *ChannelInxReq, opts ...grpc.CallOption) (*ChannelInxRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelInxRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelInx_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseServiceClient) ChannelComment(ctx context.Context, in *ChannelCommentReq, opts ...grpc.CallOption) (*ChannelCommentRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChannelCommentRes)
	err := c.cc.Invoke(ctx, BaseService_ChannelComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServiceServer is the server API for BaseService service.
// All implementations must embed UnimplementedBaseServiceServer
// for forward compatibility.
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type BaseServiceServer interface {
	// 微信小程序登录接口
	// @security BearerAuth
	LoginWeChat(context.Context, *LoginWeChatReq) (*LoginWeChatRes, error)
	// 微信小程序快速登录注册接口定义
	FastRegisterWeChat(context.Context, *FastRegisterWeChatReq) (*FastRegisterWeChatRes, error)
	// 通用地点搜索
	LocationCommonSearch(context.Context, *LocationCommonSearchReq) (*LocationCommonSearchRes, error)
	// 批量获取对象上传预签名URL
	MediaPutURLBatchGet(context.Context, *MediaPutURLBatchGetReq) (*MediaPutURLBatchGetRes, error)
	// 列表查询足迹频道类型
	ChannelTypeList(context.Context, *ChannelTypeListReq) (*ChannelTypeListRes, error)
	// 创建足迹频道
	ChannelCreate(context.Context, *ChannelCreateReq) (*ChannelCreateRes, error)
	// 更新足迹频道
	ChannelUpdate(context.Context, *ChannelUpdateReq) (*ChannelUpdateRes, error)
	// 删除足迹频道
	ChannelDelete(context.Context, *ChannelDeleteReq) (*ChannelDeleteRes, error)
	// 按照范围查询足迹基础信息
	ChannelBaseQueryByBound(context.Context, *ChannelBaseQueryByBoundReq) (*ChannelBaseQueryByBoundRes, error)
	// 按照id查询足迹频道动态信息
	ChannelFullQueryById(context.Context, *ChannelFullQueryByIdReq) (*ChannelFullQueryByIdRes, error)
	// 足迹频道互动
	ChannelInx(context.Context, *ChannelInxReq) (*ChannelInxRes, error)
	// 足迹频道评论
	ChannelComment(context.Context, *ChannelCommentReq) (*ChannelCommentRes, error)
	mustEmbedUnimplementedBaseServiceServer()
}

// UnimplementedBaseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBaseServiceServer struct{}

func (UnimplementedBaseServiceServer) LoginWeChat(context.Context, *LoginWeChatReq) (*LoginWeChatRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWeChat not implemented")
}
func (UnimplementedBaseServiceServer) FastRegisterWeChat(context.Context, *FastRegisterWeChatReq) (*FastRegisterWeChatRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FastRegisterWeChat not implemented")
}
func (UnimplementedBaseServiceServer) LocationCommonSearch(context.Context, *LocationCommonSearchReq) (*LocationCommonSearchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationCommonSearch not implemented")
}
func (UnimplementedBaseServiceServer) MediaPutURLBatchGet(context.Context, *MediaPutURLBatchGetReq) (*MediaPutURLBatchGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MediaPutURLBatchGet not implemented")
}
func (UnimplementedBaseServiceServer) ChannelTypeList(context.Context, *ChannelTypeListReq) (*ChannelTypeListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelTypeList not implemented")
}
func (UnimplementedBaseServiceServer) ChannelCreate(context.Context, *ChannelCreateReq) (*ChannelCreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelCreate not implemented")
}
func (UnimplementedBaseServiceServer) ChannelUpdate(context.Context, *ChannelUpdateReq) (*ChannelUpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelUpdate not implemented")
}
func (UnimplementedBaseServiceServer) ChannelDelete(context.Context, *ChannelDeleteReq) (*ChannelDeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelDelete not implemented")
}
func (UnimplementedBaseServiceServer) ChannelBaseQueryByBound(context.Context, *ChannelBaseQueryByBoundReq) (*ChannelBaseQueryByBoundRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelBaseQueryByBound not implemented")
}
func (UnimplementedBaseServiceServer) ChannelFullQueryById(context.Context, *ChannelFullQueryByIdReq) (*ChannelFullQueryByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelFullQueryById not implemented")
}
func (UnimplementedBaseServiceServer) ChannelInx(context.Context, *ChannelInxReq) (*ChannelInxRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelInx not implemented")
}
func (UnimplementedBaseServiceServer) ChannelComment(context.Context, *ChannelCommentReq) (*ChannelCommentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelComment not implemented")
}
func (UnimplementedBaseServiceServer) mustEmbedUnimplementedBaseServiceServer() {}
func (UnimplementedBaseServiceServer) testEmbeddedByValue()                     {}

// UnsafeBaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServiceServer will
// result in compilation errors.
type UnsafeBaseServiceServer interface {
	mustEmbedUnimplementedBaseServiceServer()
}

func RegisterBaseServiceServer(s grpc.ServiceRegistrar, srv BaseServiceServer) {
	// If the following call pancis, it indicates UnimplementedBaseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BaseService_ServiceDesc, srv)
}

func _BaseService_LoginWeChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWeChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).LoginWeChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_LoginWeChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).LoginWeChat(ctx, req.(*LoginWeChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_FastRegisterWeChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastRegisterWeChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).FastRegisterWeChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_FastRegisterWeChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).FastRegisterWeChat(ctx, req.(*FastRegisterWeChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_LocationCommonSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationCommonSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).LocationCommonSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_LocationCommonSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).LocationCommonSearch(ctx, req.(*LocationCommonSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_MediaPutURLBatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaPutURLBatchGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).MediaPutURLBatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_MediaPutURLBatchGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).MediaPutURLBatchGet(ctx, req.(*MediaPutURLBatchGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelTypeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelTypeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelTypeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelTypeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelTypeList(ctx, req.(*ChannelTypeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelCreate(ctx, req.(*ChannelCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelUpdate(ctx, req.(*ChannelUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelDelete(ctx, req.(*ChannelDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelBaseQueryByBound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelBaseQueryByBoundReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelBaseQueryByBound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelBaseQueryByBound_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelBaseQueryByBound(ctx, req.(*ChannelBaseQueryByBoundReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelFullQueryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelFullQueryByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelFullQueryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelFullQueryById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelFullQueryById(ctx, req.(*ChannelFullQueryByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelInx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelInxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelInx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelInx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelInx(ctx, req.(*ChannelInxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseService_ChannelComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServiceServer).ChannelComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BaseService_ChannelComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServiceServer).ChannelComment(ctx, req.(*ChannelCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BaseService_ServiceDesc is the grpc.ServiceDesc for BaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "base_service.v1.BaseService",
	HandlerType: (*BaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginWeChat",
			Handler:    _BaseService_LoginWeChat_Handler,
		},
		{
			MethodName: "FastRegisterWeChat",
			Handler:    _BaseService_FastRegisterWeChat_Handler,
		},
		{
			MethodName: "LocationCommonSearch",
			Handler:    _BaseService_LocationCommonSearch_Handler,
		},
		{
			MethodName: "MediaPutURLBatchGet",
			Handler:    _BaseService_MediaPutURLBatchGet_Handler,
		},
		{
			MethodName: "ChannelTypeList",
			Handler:    _BaseService_ChannelTypeList_Handler,
		},
		{
			MethodName: "ChannelCreate",
			Handler:    _BaseService_ChannelCreate_Handler,
		},
		{
			MethodName: "ChannelUpdate",
			Handler:    _BaseService_ChannelUpdate_Handler,
		},
		{
			MethodName: "ChannelDelete",
			Handler:    _BaseService_ChannelDelete_Handler,
		},
		{
			MethodName: "ChannelBaseQueryByBound",
			Handler:    _BaseService_ChannelBaseQueryByBound_Handler,
		},
		{
			MethodName: "ChannelFullQueryById",
			Handler:    _BaseService_ChannelFullQueryById_Handler,
		},
		{
			MethodName: "ChannelInx",
			Handler:    _BaseService_ChannelInx_Handler,
		},
		{
			MethodName: "ChannelComment",
			Handler:    _BaseService_ChannelComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base-service.proto",
}
