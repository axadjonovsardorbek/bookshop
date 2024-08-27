// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: publishers.proto

package books

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

// PublishersServiceClient is the client API for PublishersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishersServiceClient interface {
	Create(ctx context.Context, in *PublishersRes, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*PublishersGetByIdRes, error)
	GetAll(ctx context.Context, in *PublishersGetAllReq, opts ...grpc.CallOption) (*PublishersGetAllRes, error)
	Update(ctx context.Context, in *PublishersUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type publishersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishersServiceClient(cc grpc.ClientConnInterface) PublishersServiceClient {
	return &publishersServiceClient{cc}
}

func (c *publishersServiceClient) Create(ctx context.Context, in *PublishersRes, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/books.PublishersService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishersServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*PublishersGetByIdRes, error) {
	out := new(PublishersGetByIdRes)
	err := c.cc.Invoke(ctx, "/books.PublishersService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishersServiceClient) GetAll(ctx context.Context, in *PublishersGetAllReq, opts ...grpc.CallOption) (*PublishersGetAllRes, error) {
	out := new(PublishersGetAllRes)
	err := c.cc.Invoke(ctx, "/books.PublishersService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishersServiceClient) Update(ctx context.Context, in *PublishersUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/books.PublishersService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishersServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/books.PublishersService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishersServiceServer is the server API for PublishersService service.
// All implementations must embed UnimplementedPublishersServiceServer
// for forward compatibility
type PublishersServiceServer interface {
	Create(context.Context, *PublishersRes) (*Void, error)
	GetById(context.Context, *ById) (*PublishersGetByIdRes, error)
	GetAll(context.Context, *PublishersGetAllReq) (*PublishersGetAllRes, error)
	Update(context.Context, *PublishersUpdateReq) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedPublishersServiceServer()
}

// UnimplementedPublishersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublishersServiceServer struct {
}

func (UnimplementedPublishersServiceServer) Create(context.Context, *PublishersRes) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPublishersServiceServer) GetById(context.Context, *ById) (*PublishersGetByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedPublishersServiceServer) GetAll(context.Context, *PublishersGetAllReq) (*PublishersGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPublishersServiceServer) Update(context.Context, *PublishersUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPublishersServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPublishersServiceServer) mustEmbedUnimplementedPublishersServiceServer() {}

// UnsafePublishersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishersServiceServer will
// result in compilation errors.
type UnsafePublishersServiceServer interface {
	mustEmbedUnimplementedPublishersServiceServer()
}

func RegisterPublishersServiceServer(s grpc.ServiceRegistrar, srv PublishersServiceServer) {
	s.RegisterService(&PublishersService_ServiceDesc, srv)
}

func _PublishersService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishersRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishersServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.PublishersService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishersServiceServer).Create(ctx, req.(*PublishersRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishersService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishersServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.PublishersService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishersServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishersService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishersGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishersServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.PublishersService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishersServiceServer).GetAll(ctx, req.(*PublishersGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishersService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishersUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishersServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.PublishersService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishersServiceServer).Update(ctx, req.(*PublishersUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishersService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishersServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/books.PublishersService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishersServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// PublishersService_ServiceDesc is the grpc.ServiceDesc for PublishersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublishersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.PublishersService",
	HandlerType: (*PublishersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PublishersService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _PublishersService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PublishersService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PublishersService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PublishersService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publishers.proto",
}