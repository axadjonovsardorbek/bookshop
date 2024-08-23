// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: books.proto

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

const (
	BooksService_Create_FullMethodName  = "/books.BooksService/Create"
	BooksService_GetById_FullMethodName = "/books.BooksService/GetById"
	BooksService_GetAll_FullMethodName  = "/books.BooksService/GetAll"
	BooksService_Update_FullMethodName  = "/books.BooksService/Update"
	BooksService_Delete_FullMethodName  = "/books.BooksService/Delete"
)

// BooksServiceClient is the client API for BooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksServiceClient interface {
	Create(ctx context.Context, in *BooksCreateReq, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BooksRes, error)
	GetAll(ctx context.Context, in *BooksGetAllReq, opts ...grpc.CallOption) (*BooksGetAllRes, error)
	Update(ctx context.Context, in *BooksUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type booksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksServiceClient(cc grpc.ClientConnInterface) BooksServiceClient {
	return &booksServiceClient{cc}
}

func (c *booksServiceClient) Create(ctx context.Context, in *BooksCreateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, BooksService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BooksRes, error) {
	out := new(BooksRes)
	err := c.cc.Invoke(ctx, BooksService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) GetAll(ctx context.Context, in *BooksGetAllReq, opts ...grpc.CallOption) (*BooksGetAllRes, error) {
	out := new(BooksGetAllRes)
	err := c.cc.Invoke(ctx, BooksService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) Update(ctx context.Context, in *BooksUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, BooksService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *booksServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, BooksService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServiceServer is the server API for BooksService service.
// All implementations must embed UnimplementedBooksServiceServer
// for forward compatibility
type BooksServiceServer interface {
	Create(context.Context, *BooksCreateReq) (*Void, error)
	GetById(context.Context, *ById) (*BooksRes, error)
	GetAll(context.Context, *BooksGetAllReq) (*BooksGetAllRes, error)
	Update(context.Context, *BooksUpdateReq) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedBooksServiceServer()
}

// UnimplementedBooksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBooksServiceServer struct {
}

func (UnimplementedBooksServiceServer) Create(context.Context, *BooksCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBooksServiceServer) GetById(context.Context, *ById) (*BooksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBooksServiceServer) GetAll(context.Context, *BooksGetAllReq) (*BooksGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBooksServiceServer) Update(context.Context, *BooksUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBooksServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBooksServiceServer) mustEmbedUnimplementedBooksServiceServer() {}

// UnsafeBooksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServiceServer will
// result in compilation errors.
type UnsafeBooksServiceServer interface {
	mustEmbedUnimplementedBooksServiceServer()
}

func RegisterBooksServiceServer(s grpc.ServiceRegistrar, srv BooksServiceServer) {
	s.RegisterService(&BooksService_ServiceDesc, srv)
}

func _BooksService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BooksCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).Create(ctx, req.(*BooksCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BooksGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).GetAll(ctx, req.(*BooksGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BooksUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).Update(ctx, req.(*BooksUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BooksService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BooksService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// BooksService_ServiceDesc is the grpc.ServiceDesc for BooksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BooksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.BooksService",
	HandlerType: (*BooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BooksService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BooksService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BooksService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BooksService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BooksService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books.proto",
}
