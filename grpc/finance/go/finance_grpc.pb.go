// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package financePb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FinanceClient is the client API for Finance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinanceClient interface {
	SetUpStripeUser(ctx context.Context, opts ...grpc.CallOption) (Finance_SetUpStripeUserClient, error)
	GetClientSecret(ctx context.Context, in *GetClientSecretRequest, opts ...grpc.CallOption) (*GetClientSecretResponse, error)
	GetCustomerPaymentMethods(ctx context.Context, in *GetCustomerPaymentMethodsRequest, opts ...grpc.CallOption) (*GetCustomerPaymentMethodsResponse, error)
	DeleteCustomerPaymentMethod(ctx context.Context, in *DeleteCustomerPaymentMethodRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	SetDefaultPaymentMethod(ctx context.Context, in *SetDefaultPaymentMethodRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCourierPayable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCourierPayableResponse, error)
	GetRequestTransactions(ctx context.Context, in *GetRequestTransactionsRequest, opts ...grpc.CallOption) (*GetRequestTransactionsResponse, error)
	GetTransactionsPaidByCustomer(ctx context.Context, in *GetTransactionsPaidByCustomerRequest, opts ...grpc.CallOption) (*GetTransactionsPaidByCustomerResponse, error)
	GetTransactionsPaidToCourier(ctx context.Context, in *GetTransactionsPaidToCourierRequest, opts ...grpc.CallOption) (*GetTransactionsPaidToCourierResponse, error)
	GetAmountPaidToCourier(ctx context.Context, in *GetAmountPaidToCourierRequest, opts ...grpc.CallOption) (*GetAmountPaidToCourierResponse, error)
}

type financeClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceClient(cc grpc.ClientConnInterface) FinanceClient {
	return &financeClient{cc}
}

func (c *financeClient) SetUpStripeUser(ctx context.Context, opts ...grpc.CallOption) (Finance_SetUpStripeUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &Finance_ServiceDesc.Streams[0], "/artin.couriermanagement.finance.Finance/SetUpStripeUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &financeSetUpStripeUserClient{stream}
	return x, nil
}

type Finance_SetUpStripeUserClient interface {
	Send(*SetUpStripeUserRequest) error
	Recv() (*SetUpStripeUserResponse, error)
	grpc.ClientStream
}

type financeSetUpStripeUserClient struct {
	grpc.ClientStream
}

func (x *financeSetUpStripeUserClient) Send(m *SetUpStripeUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *financeSetUpStripeUserClient) Recv() (*SetUpStripeUserResponse, error) {
	m := new(SetUpStripeUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *financeClient) GetClientSecret(ctx context.Context, in *GetClientSecretRequest, opts ...grpc.CallOption) (*GetClientSecretResponse, error) {
	out := new(GetClientSecretResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/GetClientSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) GetCustomerPaymentMethods(ctx context.Context, in *GetCustomerPaymentMethodsRequest, opts ...grpc.CallOption) (*GetCustomerPaymentMethodsResponse, error) {
	out := new(GetCustomerPaymentMethodsResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/GetCustomerPaymentMethods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) DeleteCustomerPaymentMethod(ctx context.Context, in *DeleteCustomerPaymentMethodRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/DeleteCustomerPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) SetDefaultPaymentMethod(ctx context.Context, in *SetDefaultPaymentMethodRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/SetDefaultPaymentMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) GetCourierPayable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCourierPayableResponse, error) {
	out := new(GetCourierPayableResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/GetCourierPayable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) GetRequestTransactions(ctx context.Context, in *GetRequestTransactionsRequest, opts ...grpc.CallOption) (*GetRequestTransactionsResponse, error) {
	out := new(GetRequestTransactionsResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/GetRequestTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) GetTransactionsPaidByCustomer(ctx context.Context, in *GetTransactionsPaidByCustomerRequest, opts ...grpc.CallOption) (*GetTransactionsPaidByCustomerResponse, error) {
	out := new(GetTransactionsPaidByCustomerResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/GetTransactionsPaidByCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) GetTransactionsPaidToCourier(ctx context.Context, in *GetTransactionsPaidToCourierRequest, opts ...grpc.CallOption) (*GetTransactionsPaidToCourierResponse, error) {
	out := new(GetTransactionsPaidToCourierResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/GetTransactionsPaidToCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *financeClient) GetAmountPaidToCourier(ctx context.Context, in *GetAmountPaidToCourierRequest, opts ...grpc.CallOption) (*GetAmountPaidToCourierResponse, error) {
	out := new(GetAmountPaidToCourierResponse)
	err := c.cc.Invoke(ctx, "/artin.couriermanagement.finance.Finance/GetAmountPaidToCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceServer is the server API for Finance service.
// All implementations must embed UnimplementedFinanceServer
// for forward compatibility
type FinanceServer interface {
	SetUpStripeUser(Finance_SetUpStripeUserServer) error
	GetClientSecret(context.Context, *GetClientSecretRequest) (*GetClientSecretResponse, error)
	GetCustomerPaymentMethods(context.Context, *GetCustomerPaymentMethodsRequest) (*GetCustomerPaymentMethodsResponse, error)
	DeleteCustomerPaymentMethod(context.Context, *DeleteCustomerPaymentMethodRequest) (*emptypb.Empty, error)
	CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	SetDefaultPaymentMethod(context.Context, *SetDefaultPaymentMethodRequest) (*emptypb.Empty, error)
	GetCourierPayable(context.Context, *emptypb.Empty) (*GetCourierPayableResponse, error)
	GetRequestTransactions(context.Context, *GetRequestTransactionsRequest) (*GetRequestTransactionsResponse, error)
	GetTransactionsPaidByCustomer(context.Context, *GetTransactionsPaidByCustomerRequest) (*GetTransactionsPaidByCustomerResponse, error)
	GetTransactionsPaidToCourier(context.Context, *GetTransactionsPaidToCourierRequest) (*GetTransactionsPaidToCourierResponse, error)
	GetAmountPaidToCourier(context.Context, *GetAmountPaidToCourierRequest) (*GetAmountPaidToCourierResponse, error)
	mustEmbedUnimplementedFinanceServer()
}

// UnimplementedFinanceServer must be embedded to have forward compatible implementations.
type UnimplementedFinanceServer struct {
}

func (UnimplementedFinanceServer) SetUpStripeUser(Finance_SetUpStripeUserServer) error {
	return status.Errorf(codes.Unimplemented, "method SetUpStripeUser not implemented")
}
func (UnimplementedFinanceServer) GetClientSecret(context.Context, *GetClientSecretRequest) (*GetClientSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientSecret not implemented")
}
func (UnimplementedFinanceServer) GetCustomerPaymentMethods(context.Context, *GetCustomerPaymentMethodsRequest) (*GetCustomerPaymentMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerPaymentMethods not implemented")
}
func (UnimplementedFinanceServer) DeleteCustomerPaymentMethod(context.Context, *DeleteCustomerPaymentMethodRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomerPaymentMethod not implemented")
}
func (UnimplementedFinanceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedFinanceServer) SetDefaultPaymentMethod(context.Context, *SetDefaultPaymentMethodRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultPaymentMethod not implemented")
}
func (UnimplementedFinanceServer) GetCourierPayable(context.Context, *emptypb.Empty) (*GetCourierPayableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierPayable not implemented")
}
func (UnimplementedFinanceServer) GetRequestTransactions(context.Context, *GetRequestTransactionsRequest) (*GetRequestTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestTransactions not implemented")
}
func (UnimplementedFinanceServer) GetTransactionsPaidByCustomer(context.Context, *GetTransactionsPaidByCustomerRequest) (*GetTransactionsPaidByCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsPaidByCustomer not implemented")
}
func (UnimplementedFinanceServer) GetTransactionsPaidToCourier(context.Context, *GetTransactionsPaidToCourierRequest) (*GetTransactionsPaidToCourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsPaidToCourier not implemented")
}
func (UnimplementedFinanceServer) GetAmountPaidToCourier(context.Context, *GetAmountPaidToCourierRequest) (*GetAmountPaidToCourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAmountPaidToCourier not implemented")
}
func (UnimplementedFinanceServer) mustEmbedUnimplementedFinanceServer() {}

// UnsafeFinanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinanceServer will
// result in compilation errors.
type UnsafeFinanceServer interface {
	mustEmbedUnimplementedFinanceServer()
}

func RegisterFinanceServer(s grpc.ServiceRegistrar, srv FinanceServer) {
	s.RegisterService(&Finance_ServiceDesc, srv)
}

func _Finance_SetUpStripeUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FinanceServer).SetUpStripeUser(&financeSetUpStripeUserServer{stream})
}

type Finance_SetUpStripeUserServer interface {
	Send(*SetUpStripeUserResponse) error
	Recv() (*SetUpStripeUserRequest, error)
	grpc.ServerStream
}

type financeSetUpStripeUserServer struct {
	grpc.ServerStream
}

func (x *financeSetUpStripeUserServer) Send(m *SetUpStripeUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *financeSetUpStripeUserServer) Recv() (*SetUpStripeUserRequest, error) {
	m := new(SetUpStripeUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Finance_GetClientSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).GetClientSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/GetClientSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).GetClientSecret(ctx, req.(*GetClientSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_GetCustomerPaymentMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerPaymentMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).GetCustomerPaymentMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/GetCustomerPaymentMethods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).GetCustomerPaymentMethods(ctx, req.(*GetCustomerPaymentMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_DeleteCustomerPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).DeleteCustomerPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/DeleteCustomerPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).DeleteCustomerPaymentMethod(ctx, req.(*DeleteCustomerPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_SetDefaultPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDefaultPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).SetDefaultPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/SetDefaultPaymentMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).SetDefaultPaymentMethod(ctx, req.(*SetDefaultPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_GetCourierPayable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).GetCourierPayable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/GetCourierPayable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).GetCourierPayable(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_GetRequestTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).GetRequestTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/GetRequestTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).GetRequestTransactions(ctx, req.(*GetRequestTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_GetTransactionsPaidByCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsPaidByCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).GetTransactionsPaidByCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/GetTransactionsPaidByCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).GetTransactionsPaidByCustomer(ctx, req.(*GetTransactionsPaidByCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_GetTransactionsPaidToCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsPaidToCourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).GetTransactionsPaidToCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/GetTransactionsPaidToCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).GetTransactionsPaidToCourier(ctx, req.(*GetTransactionsPaidToCourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finance_GetAmountPaidToCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAmountPaidToCourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceServer).GetAmountPaidToCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/artin.couriermanagement.finance.Finance/GetAmountPaidToCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceServer).GetAmountPaidToCourier(ctx, req.(*GetAmountPaidToCourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Finance_ServiceDesc is the grpc.ServiceDesc for Finance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Finance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "artin.couriermanagement.finance.Finance",
	HandlerType: (*FinanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientSecret",
			Handler:    _Finance_GetClientSecret_Handler,
		},
		{
			MethodName: "GetCustomerPaymentMethods",
			Handler:    _Finance_GetCustomerPaymentMethods_Handler,
		},
		{
			MethodName: "DeleteCustomerPaymentMethod",
			Handler:    _Finance_DeleteCustomerPaymentMethod_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _Finance_CreatePayment_Handler,
		},
		{
			MethodName: "SetDefaultPaymentMethod",
			Handler:    _Finance_SetDefaultPaymentMethod_Handler,
		},
		{
			MethodName: "GetCourierPayable",
			Handler:    _Finance_GetCourierPayable_Handler,
		},
		{
			MethodName: "GetRequestTransactions",
			Handler:    _Finance_GetRequestTransactions_Handler,
		},
		{
			MethodName: "GetTransactionsPaidByCustomer",
			Handler:    _Finance_GetTransactionsPaidByCustomer_Handler,
		},
		{
			MethodName: "GetTransactionsPaidToCourier",
			Handler:    _Finance_GetTransactionsPaidToCourier_Handler,
		},
		{
			MethodName: "GetAmountPaidToCourier",
			Handler:    _Finance_GetAmountPaidToCourier_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetUpStripeUser",
			Handler:       _Finance_SetUpStripeUser_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "finance.proto",
}
