// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: category.proto

package pb

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
	CategoryService_FindCategories_FullMethodName = "/category.CategoryService/FindCategories"
	CategoryService_FindCategory_FullMethodName   = "/category.CategoryService/FindCategory"
	CategoryService_SaveCategory_FullMethodName   = "/category.CategoryService/SaveCategory"
	CategoryService_DeleteCategory_FullMethodName = "/category.CategoryService/DeleteCategory"
)

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	FindCategories(ctx context.Context, in *FindCategoriesRequest, opts ...grpc.CallOption) (*FindCategoriesResponse, error)
	FindCategory(ctx context.Context, in *FindCategoryRequest, opts ...grpc.CallOption) (*FindCategoryResponse, error)
	SaveCategory(ctx context.Context, in *SaveCategoryRequest, opts ...grpc.CallOption) (*SaveRecordResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) FindCategories(ctx context.Context, in *FindCategoriesRequest, opts ...grpc.CallOption) (*FindCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCategoriesResponse)
	err := c.cc.Invoke(ctx, CategoryService_FindCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) FindCategory(ctx context.Context, in *FindCategoryRequest, opts ...grpc.CallOption) (*FindCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_FindCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) SaveCategory(ctx context.Context, in *SaveCategoryRequest, opts ...grpc.CallOption) (*SaveRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveRecordResponse)
	err := c.cc.Invoke(ctx, CategoryService_SaveCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteCategory(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRecordResponse)
	err := c.cc.Invoke(ctx, CategoryService_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility.
type CategoryServiceServer interface {
	FindCategories(context.Context, *FindCategoriesRequest) (*FindCategoriesResponse, error)
	FindCategory(context.Context, *FindCategoryRequest) (*FindCategoryResponse, error)
	SaveCategory(context.Context, *SaveCategoryRequest) (*SaveRecordResponse, error)
	DeleteCategory(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCategoryServiceServer struct{}

func (UnimplementedCategoryServiceServer) FindCategories(context.Context, *FindCategoriesRequest) (*FindCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCategories not implemented")
}
func (UnimplementedCategoryServiceServer) FindCategory(context.Context, *FindCategoryRequest) (*FindCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCategory not implemented")
}
func (UnimplementedCategoryServiceServer) SaveCategory(context.Context, *SaveCategoryRequest) (*SaveRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCategory not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteCategory(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}
func (UnimplementedCategoryServiceServer) testEmbeddedByValue()                         {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedCategoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_FindCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).FindCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_FindCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).FindCategories(ctx, req.(*FindCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_FindCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).FindCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_FindCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).FindCategory(ctx, req.(*FindCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_SaveCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).SaveCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_SaveCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).SaveCategory(ctx, req.(*SaveCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "category.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindCategories",
			Handler:    _CategoryService_FindCategories_Handler,
		},
		{
			MethodName: "FindCategory",
			Handler:    _CategoryService_FindCategory_Handler,
		},
		{
			MethodName: "SaveCategory",
			Handler:    _CategoryService_SaveCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CategoryService_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}
