// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/delivery.proto

package __

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

// CourierServiceClient is the client API for CourierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourierServiceClient interface {
	CreateCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	GetCourier(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*CourierResponse, error)
	GetCouriers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CourierListResponse, error)
	UpdateCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	DeleteCourier(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*Message, error)
}

type courierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierServiceClient(cc grpc.ClientConnInterface) CourierServiceClient {
	return &courierServiceClient{cc}
}

func (c *courierServiceClient) CreateCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, "/proto.CourierService/CreateCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) GetCourier(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*CourierResponse, error) {
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, "/proto.CourierService/GetCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) GetCouriers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CourierListResponse, error) {
	out := new(CourierListResponse)
	err := c.cc.Invoke(ctx, "/proto.CourierService/GetCouriers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) UpdateCourier(ctx context.Context, in *CourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, "/proto.CourierService/UpdateCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) DeleteCourier(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/proto.CourierService/DeleteCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierServiceServer is the server API for CourierService service.
// All implementations must embed UnimplementedCourierServiceServer
// for forward compatibility
type CourierServiceServer interface {
	CreateCourier(context.Context, *CourierRequest) (*CourierResponse, error)
	GetCourier(context.Context, *ItemID) (*CourierResponse, error)
	GetCouriers(context.Context, *EmptyRequest) (*CourierListResponse, error)
	UpdateCourier(context.Context, *CourierRequest) (*CourierResponse, error)
	DeleteCourier(context.Context, *ItemID) (*Message, error)
	mustEmbedUnimplementedCourierServiceServer()
}

// UnimplementedCourierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourierServiceServer struct {
}

func (UnimplementedCourierServiceServer) CreateCourier(context.Context, *CourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourier not implemented")
}
func (UnimplementedCourierServiceServer) GetCourier(context.Context, *ItemID) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourier not implemented")
}
func (UnimplementedCourierServiceServer) GetCouriers(context.Context, *EmptyRequest) (*CourierListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouriers not implemented")
}
func (UnimplementedCourierServiceServer) UpdateCourier(context.Context, *CourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourier not implemented")
}
func (UnimplementedCourierServiceServer) DeleteCourier(context.Context, *ItemID) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourier not implemented")
}
func (UnimplementedCourierServiceServer) mustEmbedUnimplementedCourierServiceServer() {}

// UnsafeCourierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierServiceServer will
// result in compilation errors.
type UnsafeCourierServiceServer interface {
	mustEmbedUnimplementedCourierServiceServer()
}

func RegisterCourierServiceServer(s grpc.ServiceRegistrar, srv CourierServiceServer) {
	s.RegisterService(&CourierService_ServiceDesc, srv)
}

func _CourierService_CreateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).CreateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CourierService/CreateCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).CreateCourier(ctx, req.(*CourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_GetCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).GetCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CourierService/GetCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).GetCourier(ctx, req.(*ItemID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_GetCouriers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).GetCouriers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CourierService/GetCouriers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).GetCouriers(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_UpdateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).UpdateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CourierService/UpdateCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).UpdateCourier(ctx, req.(*CourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_DeleteCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).DeleteCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CourierService/DeleteCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).DeleteCourier(ctx, req.(*ItemID))
	}
	return interceptor(ctx, in, info, handler)
}

// CourierService_ServiceDesc is the grpc.ServiceDesc for CourierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CourierService",
	HandlerType: (*CourierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourier",
			Handler:    _CourierService_CreateCourier_Handler,
		},
		{
			MethodName: "GetCourier",
			Handler:    _CourierService_GetCourier_Handler,
		},
		{
			MethodName: "GetCouriers",
			Handler:    _CourierService_GetCouriers_Handler,
		},
		{
			MethodName: "UpdateCourier",
			Handler:    _CourierService_UpdateCourier_Handler,
		},
		{
			MethodName: "DeleteCourier",
			Handler:    _CourierService_DeleteCourier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/delivery.proto",
}

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
	GetCustomer(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*CustomerResponse, error)
	GetCustomers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CustomerListResponse, error)
	UpdateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error)
	DeleteCustomer(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*Message, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomers(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CustomerListResponse, error) {
	out := new(CustomerListResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/GetCustomers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*CustomerResponse, error) {
	out := new(CustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) DeleteCustomer(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
	GetCustomer(context.Context, *ItemID) (*CustomerResponse, error)
	GetCustomers(context.Context, *EmptyRequest) (*CustomerListResponse, error)
	UpdateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error)
	DeleteCustomer(context.Context, *ItemID) (*Message, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) CreateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomer(context.Context, *ItemID) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) GetCustomers(context.Context, *EmptyRequest) (*CustomerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomers not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomer(context.Context, *CustomerRequest) (*CustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) DeleteCustomer(context.Context, *ItemID) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
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

func _CustomerService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*ItemID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/GetCustomers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomers(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, req.(*ItemID))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerService_CreateCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
		{
			MethodName: "GetCustomers",
			Handler:    _CustomerService_GetCustomers_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _CustomerService_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _CustomerService_DeleteCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/delivery.proto",
}

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	CreateRestaurant(ctx context.Context, in *RestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error)
	GetRestaurant(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*RestaurantResponse, error)
	GetRestaurants(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RestaurantListResponse, error)
	UpdateRestaurant(ctx context.Context, in *RestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error)
	DeleteRestaurant(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*Message, error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) CreateRestaurant(ctx context.Context, in *RestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error) {
	out := new(RestaurantResponse)
	err := c.cc.Invoke(ctx, "/proto.RestaurantService/CreateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetRestaurant(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*RestaurantResponse, error) {
	out := new(RestaurantResponse)
	err := c.cc.Invoke(ctx, "/proto.RestaurantService/GetRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetRestaurants(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RestaurantListResponse, error) {
	out := new(RestaurantListResponse)
	err := c.cc.Invoke(ctx, "/proto.RestaurantService/GetRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) UpdateRestaurant(ctx context.Context, in *RestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error) {
	out := new(RestaurantResponse)
	err := c.cc.Invoke(ctx, "/proto.RestaurantService/UpdateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) DeleteRestaurant(ctx context.Context, in *ItemID, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/proto.RestaurantService/DeleteRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
// All implementations must embed UnimplementedRestaurantServiceServer
// for forward compatibility
type RestaurantServiceServer interface {
	CreateRestaurant(context.Context, *RestaurantRequest) (*RestaurantResponse, error)
	GetRestaurant(context.Context, *ItemID) (*RestaurantResponse, error)
	GetRestaurants(context.Context, *EmptyRequest) (*RestaurantListResponse, error)
	UpdateRestaurant(context.Context, *RestaurantRequest) (*RestaurantResponse, error)
	DeleteRestaurant(context.Context, *ItemID) (*Message, error)
	mustEmbedUnimplementedRestaurantServiceServer()
}

// UnimplementedRestaurantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRestaurantServiceServer struct {
}

func (UnimplementedRestaurantServiceServer) CreateRestaurant(context.Context, *RestaurantRequest) (*RestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) GetRestaurant(context.Context, *ItemID) (*RestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) GetRestaurants(context.Context, *EmptyRequest) (*RestaurantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurants not implemented")
}
func (UnimplementedRestaurantServiceServer) UpdateRestaurant(context.Context, *RestaurantRequest) (*RestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) DeleteRestaurant(context.Context, *ItemID) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) mustEmbedUnimplementedRestaurantServiceServer() {}

// UnsafeRestaurantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServiceServer will
// result in compilation errors.
type UnsafeRestaurantServiceServer interface {
	mustEmbedUnimplementedRestaurantServiceServer()
}

func RegisterRestaurantServiceServer(s grpc.ServiceRegistrar, srv RestaurantServiceServer) {
	s.RegisterService(&RestaurantService_ServiceDesc, srv)
}

func _RestaurantService_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RestaurantService/CreateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, req.(*RestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RestaurantService/GetRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetRestaurant(ctx, req.(*ItemID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RestaurantService/GetRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetRestaurants(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_UpdateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).UpdateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RestaurantService/UpdateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).UpdateRestaurant(ctx, req.(*RestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_DeleteRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).DeleteRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RestaurantService/DeleteRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).DeleteRestaurant(ctx, req.(*ItemID))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantService_ServiceDesc is the grpc.ServiceDesc for RestaurantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRestaurant",
			Handler:    _RestaurantService_CreateRestaurant_Handler,
		},
		{
			MethodName: "GetRestaurant",
			Handler:    _RestaurantService_GetRestaurant_Handler,
		},
		{
			MethodName: "GetRestaurants",
			Handler:    _RestaurantService_GetRestaurants_Handler,
		},
		{
			MethodName: "UpdateRestaurant",
			Handler:    _RestaurantService_UpdateRestaurant_Handler,
		},
		{
			MethodName: "DeleteRestaurant",
			Handler:    _RestaurantService_DeleteRestaurant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/delivery.proto",
}
