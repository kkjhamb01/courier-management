// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	// a client wants to rate a courier
	// can be sent by a client
	// for each ride_id there is only one courier-rate
	CreateCourierRating(ctx context.Context, in *CreateCourierRatingRequest, opts ...grpc.CallOption) (*CreateCourierRatingResponse, error)
	// a courier wants to rate a client
	// can be sent by a courier
	// for each ride_id there is only one client-rate
	CreateClientRating(ctx context.Context, in *CreateClientRatingRequest, opts ...grpc.CallOption) (*CreateClientRatingResponse, error)
	// get rate details of a courier. Which clients with what scores rate a courier
	// can be sent by either courier or client
	GetCourierRating(ctx context.Context, in *GetCourierRatingRequest, opts ...grpc.CallOption) (*GetCourierRatingResponse, error)
	// get rate details of a client. Which couriers with what scores rate a client
	// can be sent by either courier or client
	GetClientRating(ctx context.Context, in *GetClientRatingRequest, opts ...grpc.CallOption) (*GetClientRatingResponse, error)
	// get rate summary for a courier
	// can be sent by anyone, even without token
	GetCourierRatingStat(ctx context.Context, in *GetCourierRatingStatRequest, opts ...grpc.CallOption) (*GetCourierRatingStatResponse, error)
	// get rate summary for a client
	// can be sent by anyone, even without token
	GetClientRatingStat(ctx context.Context, in *GetClientRatingStatRequest, opts ...grpc.CallOption) (*GetClientRatingStatResponse, error)
	// get rate summary for a courier by its token
	// can be sent by anyone, even without token
	GetCourierRatingStatByToken(ctx context.Context, in *GetCourierRatingStatByTokenRequest, opts ...grpc.CallOption) (*GetCourierRatingStatByTokenResponse, error)
	// get rate summary for a client by its token
	// can be sent by anyone, even without token
	GetClientRatingStatByToken(ctx context.Context, in *GetClientRatingStatByTokenRequest, opts ...grpc.CallOption) (*GetClientRatingStatByTokenResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) CreateCourierRating(ctx context.Context, in *CreateCourierRatingRequest, opts ...grpc.CallOption) (*CreateCourierRatingResponse, error) {
	out := new(CreateCourierRatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/CreateCourierRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateClientRating(ctx context.Context, in *CreateClientRatingRequest, opts ...grpc.CallOption) (*CreateClientRatingResponse, error) {
	out := new(CreateClientRatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/CreateClientRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetCourierRating(ctx context.Context, in *GetCourierRatingRequest, opts ...grpc.CallOption) (*GetCourierRatingResponse, error) {
	out := new(GetCourierRatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetCourierRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetClientRating(ctx context.Context, in *GetClientRatingRequest, opts ...grpc.CallOption) (*GetClientRatingResponse, error) {
	out := new(GetClientRatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetClientRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetCourierRatingStat(ctx context.Context, in *GetCourierRatingStatRequest, opts ...grpc.CallOption) (*GetCourierRatingStatResponse, error) {
	out := new(GetCourierRatingStatResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetCourierRatingStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetClientRatingStat(ctx context.Context, in *GetClientRatingStatRequest, opts ...grpc.CallOption) (*GetClientRatingStatResponse, error) {
	out := new(GetClientRatingStatResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetClientRatingStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetCourierRatingStatByToken(ctx context.Context, in *GetCourierRatingStatByTokenRequest, opts ...grpc.CallOption) (*GetCourierRatingStatByTokenResponse, error) {
	out := new(GetCourierRatingStatByTokenResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetCourierRatingStatByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetClientRatingStatByToken(ctx context.Context, in *GetClientRatingStatByTokenRequest, opts ...grpc.CallOption) (*GetClientRatingStatByTokenResponse, error) {
	out := new(GetClientRatingStatByTokenResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetClientRatingStatByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	// a client wants to rate a courier
	// can be sent by a client
	// for each ride_id there is only one courier-rate
	CreateCourierRating(context.Context, *CreateCourierRatingRequest) (*CreateCourierRatingResponse, error)
	// a courier wants to rate a client
	// can be sent by a courier
	// for each ride_id there is only one client-rate
	CreateClientRating(context.Context, *CreateClientRatingRequest) (*CreateClientRatingResponse, error)
	// get rate details of a courier. Which clients with what scores rate a courier
	// can be sent by either courier or client
	GetCourierRating(context.Context, *GetCourierRatingRequest) (*GetCourierRatingResponse, error)
	// get rate details of a client. Which couriers with what scores rate a client
	// can be sent by either courier or client
	GetClientRating(context.Context, *GetClientRatingRequest) (*GetClientRatingResponse, error)
	// get rate summary for a courier
	// can be sent by anyone, even without token
	GetCourierRatingStat(context.Context, *GetCourierRatingStatRequest) (*GetCourierRatingStatResponse, error)
	// get rate summary for a client
	// can be sent by anyone, even without token
	GetClientRatingStat(context.Context, *GetClientRatingStatRequest) (*GetClientRatingStatResponse, error)
	// get rate summary for a courier by its token
	// can be sent by anyone, even without token
	GetCourierRatingStatByToken(context.Context, *GetCourierRatingStatByTokenRequest) (*GetCourierRatingStatByTokenResponse, error)
	// get rate summary for a client by its token
	// can be sent by anyone, even without token
	GetClientRatingStatByToken(context.Context, *GetClientRatingStatByTokenRequest) (*GetClientRatingStatByTokenResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) CreateCourierRating(context.Context, *CreateCourierRatingRequest) (*CreateCourierRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourierRating not implemented")
}
func (UnimplementedRatingServiceServer) CreateClientRating(context.Context, *CreateClientRatingRequest) (*CreateClientRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientRating not implemented")
}
func (UnimplementedRatingServiceServer) GetCourierRating(context.Context, *GetCourierRatingRequest) (*GetCourierRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierRating not implemented")
}
func (UnimplementedRatingServiceServer) GetClientRating(context.Context, *GetClientRatingRequest) (*GetClientRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientRating not implemented")
}
func (UnimplementedRatingServiceServer) GetCourierRatingStat(context.Context, *GetCourierRatingStatRequest) (*GetCourierRatingStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierRatingStat not implemented")
}
func (UnimplementedRatingServiceServer) GetClientRatingStat(context.Context, *GetClientRatingStatRequest) (*GetClientRatingStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientRatingStat not implemented")
}
func (UnimplementedRatingServiceServer) GetCourierRatingStatByToken(context.Context, *GetCourierRatingStatByTokenRequest) (*GetCourierRatingStatByTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierRatingStatByToken not implemented")
}
func (UnimplementedRatingServiceServer) GetClientRatingStatByToken(context.Context, *GetClientRatingStatByTokenRequest) (*GetClientRatingStatByTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientRatingStatByToken not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_CreateCourierRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourierRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateCourierRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/CreateCourierRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateCourierRating(ctx, req.(*CreateCourierRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateClientRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateClientRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/CreateClientRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateClientRating(ctx, req.(*CreateClientRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetCourierRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourierRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetCourierRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetCourierRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetCourierRating(ctx, req.(*GetCourierRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetClientRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetClientRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetClientRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetClientRating(ctx, req.(*GetClientRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetCourierRatingStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourierRatingStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetCourierRatingStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetCourierRatingStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetCourierRatingStat(ctx, req.(*GetCourierRatingStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetClientRatingStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRatingStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetClientRatingStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetClientRatingStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetClientRatingStat(ctx, req.(*GetClientRatingStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetCourierRatingStatByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourierRatingStatByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetCourierRatingStatByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetCourierRatingStatByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetCourierRatingStatByToken(ctx, req.(*GetCourierRatingStatByTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetClientRatingStatByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRatingStatByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetClientRatingStatByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetClientRatingStatByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetClientRatingStatByToken(ctx, req.(*GetClientRatingStatByTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourierRating",
			Handler:    _RatingService_CreateCourierRating_Handler,
		},
		{
			MethodName: "CreateClientRating",
			Handler:    _RatingService_CreateClientRating_Handler,
		},
		{
			MethodName: "GetCourierRating",
			Handler:    _RatingService_GetCourierRating_Handler,
		},
		{
			MethodName: "GetClientRating",
			Handler:    _RatingService_GetClientRating_Handler,
		},
		{
			MethodName: "GetCourierRatingStat",
			Handler:    _RatingService_GetCourierRatingStat_Handler,
		},
		{
			MethodName: "GetClientRatingStat",
			Handler:    _RatingService_GetClientRatingStat_Handler,
		},
		{
			MethodName: "GetCourierRatingStatByToken",
			Handler:    _RatingService_GetCourierRatingStatByToken_Handler,
		},
		{
			MethodName: "GetClientRatingStatByToken",
			Handler:    _RatingService_GetClientRatingStatByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating.proto",
}
