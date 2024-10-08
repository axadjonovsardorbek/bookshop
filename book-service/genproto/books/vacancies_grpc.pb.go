// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: vacancies.proto

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
	VacanciesService_Create_FullMethodName  = "/books.VacanciesService/Create"
	VacanciesService_GetById_FullMethodName = "/books.VacanciesService/GetById"
	VacanciesService_GetAll_FullMethodName  = "/books.VacanciesService/GetAll"
	VacanciesService_Update_FullMethodName  = "/books.VacanciesService/Update"
	VacanciesService_Delete_FullMethodName  = "/books.VacanciesService/Delete"
)

// VacanciesServiceClient is the client API for VacanciesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacanciesServiceClient interface {
	Create(ctx context.Context, in *VacanciesCreateReq, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*VacanciesRes, error)
	GetAll(ctx context.Context, in *VacanciesGetAllReq, opts ...grpc.CallOption) (*VacanciesGetAllRes, error)
	Update(ctx context.Context, in *VacanciesUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type vacanciesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVacanciesServiceClient(cc grpc.ClientConnInterface) VacanciesServiceClient {
	return &vacanciesServiceClient{cc}
}

func (c *vacanciesServiceClient) Create(ctx context.Context, in *VacanciesCreateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, VacanciesService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*VacanciesRes, error) {
	out := new(VacanciesRes)
	err := c.cc.Invoke(ctx, VacanciesService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) GetAll(ctx context.Context, in *VacanciesGetAllReq, opts ...grpc.CallOption) (*VacanciesGetAllRes, error) {
	out := new(VacanciesGetAllRes)
	err := c.cc.Invoke(ctx, VacanciesService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) Update(ctx context.Context, in *VacanciesUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, VacanciesService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, VacanciesService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacanciesServiceServer is the server API for VacanciesService service.
// All implementations must embed UnimplementedVacanciesServiceServer
// for forward compatibility
type VacanciesServiceServer interface {
	Create(context.Context, *VacanciesCreateReq) (*Void, error)
	GetById(context.Context, *ById) (*VacanciesRes, error)
	GetAll(context.Context, *VacanciesGetAllReq) (*VacanciesGetAllRes, error)
	Update(context.Context, *VacanciesUpdateReq) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedVacanciesServiceServer()
}

// UnimplementedVacanciesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVacanciesServiceServer struct {
}

func (UnimplementedVacanciesServiceServer) Create(context.Context, *VacanciesCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVacanciesServiceServer) GetById(context.Context, *ById) (*VacanciesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedVacanciesServiceServer) GetAll(context.Context, *VacanciesGetAllReq) (*VacanciesGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedVacanciesServiceServer) Update(context.Context, *VacanciesUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVacanciesServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedVacanciesServiceServer) mustEmbedUnimplementedVacanciesServiceServer() {}

// UnsafeVacanciesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacanciesServiceServer will
// result in compilation errors.
type UnsafeVacanciesServiceServer interface {
	mustEmbedUnimplementedVacanciesServiceServer()
}

func RegisterVacanciesServiceServer(s grpc.ServiceRegistrar, srv VacanciesServiceServer) {
	s.RegisterService(&VacanciesService_ServiceDesc, srv)
}

func _VacanciesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacanciesCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).Create(ctx, req.(*VacanciesCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacanciesGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).GetAll(ctx, req.(*VacanciesGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacanciesUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).Update(ctx, req.(*VacanciesUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// VacanciesService_ServiceDesc is the grpc.ServiceDesc for VacanciesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VacanciesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.VacanciesService",
	HandlerType: (*VacanciesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VacanciesService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _VacanciesService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _VacanciesService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VacanciesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VacanciesService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vacancies.proto",
}
