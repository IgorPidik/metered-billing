// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: customer_service.proto

package customer_service_proto

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

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	CheckCustomerDetailsValidity(ctx context.Context, in *CustomerDetails, opts ...grpc.CallOption) (*CustomerDetailsValid, error)
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerReply, error)
	CreateCustomerService(ctx context.Context, in *CreateCustomerServiceRequest, opts ...grpc.CallOption) (*CreateCustomerServiceReply, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) CheckCustomerDetailsValidity(ctx context.Context, in *CustomerDetails, opts ...grpc.CallOption) (*CustomerDetailsValid, error) {
	out := new(CustomerDetailsValid)
	err := c.cc.Invoke(ctx, "/customer_service_proto.CustomerService/CheckCustomerDetailsValidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerReply, error) {
	out := new(CreateCustomerReply)
	err := c.cc.Invoke(ctx, "/customer_service_proto.CustomerService/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomerService(ctx context.Context, in *CreateCustomerServiceRequest, opts ...grpc.CallOption) (*CreateCustomerServiceReply, error) {
	out := new(CreateCustomerServiceReply)
	err := c.cc.Invoke(ctx, "/customer_service_proto.CustomerService/CreateCustomerService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	CheckCustomerDetailsValidity(context.Context, *CustomerDetails) (*CustomerDetailsValid, error)
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerReply, error)
	CreateCustomerService(context.Context, *CreateCustomerServiceRequest) (*CreateCustomerServiceReply, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) CheckCustomerDetailsValidity(context.Context, *CustomerDetails) (*CustomerDetailsValid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCustomerDetailsValidity not implemented")
}
func (UnimplementedCustomerServiceServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) CreateCustomerService(context.Context, *CreateCustomerServiceRequest) (*CreateCustomerServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomerService not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_CheckCustomerDetailsValidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CheckCustomerDetailsValidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_service_proto.CustomerService/CheckCustomerDetailsValidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CheckCustomerDetailsValidity(ctx, req.(*CustomerDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_service_proto.CustomerService/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomerService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomerService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_service_proto.CustomerService/CreateCustomerService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomerService(ctx, req.(*CreateCustomerServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer_service_proto.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckCustomerDetailsValidity",
			Handler:    _CustomerService_CheckCustomerDetailsValidity_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerService_CreateCustomer_Handler,
		},
		{
			MethodName: "CreateCustomerService",
			Handler:    _CustomerService_CreateCustomerService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer_service.proto",
}