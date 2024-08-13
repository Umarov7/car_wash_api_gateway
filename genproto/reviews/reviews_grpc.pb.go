// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: reviews.proto

package reviews

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

// ReviewsClient is the client API for Reviews service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReviewsClient interface {
	CreateReview(ctx context.Context, in *NewReview, opts ...grpc.CallOption) (*CreateResp, error)
	GetReview(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Review, error)
	UpdateReview(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*UpdateResp, error)
	DeleteReview(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error)
	ListReviews(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*ReviewsList, error)
}

type reviewsClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewsClient(cc grpc.ClientConnInterface) ReviewsClient {
	return &reviewsClient{cc}
}

func (c *reviewsClient) CreateReview(ctx context.Context, in *NewReview, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewsClient) GetReview(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/GetReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewsClient) UpdateReview(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewsClient) DeleteReview(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/DeleteReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewsClient) ListReviews(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*ReviewsList, error) {
	out := new(ReviewsList)
	err := c.cc.Invoke(ctx, "/reviews.Reviews/ListReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewsServer is the server API for Reviews service.
// All implementations must embed UnimplementedReviewsServer
// for forward compatibility
type ReviewsServer interface {
	CreateReview(context.Context, *NewReview) (*CreateResp, error)
	GetReview(context.Context, *ID) (*Review, error)
	UpdateReview(context.Context, *NewData) (*UpdateResp, error)
	DeleteReview(context.Context, *ID) (*Void, error)
	ListReviews(context.Context, *Pagination) (*ReviewsList, error)
	mustEmbedUnimplementedReviewsServer()
}

// UnimplementedReviewsServer must be embedded to have forward compatible implementations.
type UnimplementedReviewsServer struct {
}

func (UnimplementedReviewsServer) CreateReview(context.Context, *NewReview) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedReviewsServer) GetReview(context.Context, *ID) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReview not implemented")
}
func (UnimplementedReviewsServer) UpdateReview(context.Context, *NewData) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedReviewsServer) DeleteReview(context.Context, *ID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReview not implemented")
}
func (UnimplementedReviewsServer) ListReviews(context.Context, *Pagination) (*ReviewsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReviews not implemented")
}
func (UnimplementedReviewsServer) mustEmbedUnimplementedReviewsServer() {}

// UnsafeReviewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReviewsServer will
// result in compilation errors.
type UnsafeReviewsServer interface {
	mustEmbedUnimplementedReviewsServer()
}

func RegisterReviewsServer(s grpc.ServiceRegistrar, srv ReviewsServer) {
	s.RegisterService(&Reviews_ServiceDesc, srv)
}

func _Reviews_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).CreateReview(ctx, req.(*NewReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviews_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/GetReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).GetReview(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviews_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).UpdateReview(ctx, req.(*NewData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviews_DeleteReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).DeleteReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/DeleteReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).DeleteReview(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reviews_ListReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewsServer).ListReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reviews.Reviews/ListReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewsServer).ListReviews(ctx, req.(*Pagination))
	}
	return interceptor(ctx, in, info, handler)
}

// Reviews_ServiceDesc is the grpc.ServiceDesc for Reviews service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reviews_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reviews.Reviews",
	HandlerType: (*ReviewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReview",
			Handler:    _Reviews_CreateReview_Handler,
		},
		{
			MethodName: "GetReview",
			Handler:    _Reviews_GetReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _Reviews_UpdateReview_Handler,
		},
		{
			MethodName: "DeleteReview",
			Handler:    _Reviews_DeleteReview_Handler,
		},
		{
			MethodName: "ListReviews",
			Handler:    _Reviews_ListReviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reviews.proto",
}
