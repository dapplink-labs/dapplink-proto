// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: dapplink/market.proto

package market

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
	PriceService_GetExchanges_FullMethodName       = "/dapplink.market.PriceService/getExchanges"
	PriceService_GetAssets_FullMethodName          = "/dapplink.market.PriceService/getAssets"
	PriceService_GetSymbols_FullMethodName         = "/dapplink.market.PriceService/getSymbols"
	PriceService_GetSymbolPrices_FullMethodName    = "/dapplink.market.PriceService/getSymbolPrices"
	PriceService_GetStableCoins_FullMethodName     = "/dapplink.market.PriceService/getStableCoins"
	PriceService_GetStableCoinPrice_FullMethodName = "/dapplink.market.PriceService/getStableCoinPrice"
)

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceServiceClient interface {
	GetExchanges(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error)
	GetAssets(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetResponse, error)
	GetSymbols(ctx context.Context, in *SymbolRequest, opts ...grpc.CallOption) (*SymbolResponse, error)
	GetSymbolPrices(ctx context.Context, in *SymbolPriceRequest, opts ...grpc.CallOption) (*SymbolPriceResponse, error)
	GetStableCoins(ctx context.Context, in *StableCoinRequest, opts ...grpc.CallOption) (*StableCoinResponse, error)
	GetStableCoinPrice(ctx context.Context, in *StableCoinPriceRequest, opts ...grpc.CallOption) (*StableCoinPriceResponse, error)
}

type priceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceServiceClient(cc grpc.ClientConnInterface) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) GetExchanges(ctx context.Context, in *ExchangeRequest, opts ...grpc.CallOption) (*ExchangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeResponse)
	err := c.cc.Invoke(ctx, PriceService_GetExchanges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) GetAssets(ctx context.Context, in *AssetRequest, opts ...grpc.CallOption) (*AssetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssetResponse)
	err := c.cc.Invoke(ctx, PriceService_GetAssets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) GetSymbols(ctx context.Context, in *SymbolRequest, opts ...grpc.CallOption) (*SymbolResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymbolResponse)
	err := c.cc.Invoke(ctx, PriceService_GetSymbols_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) GetSymbolPrices(ctx context.Context, in *SymbolPriceRequest, opts ...grpc.CallOption) (*SymbolPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymbolPriceResponse)
	err := c.cc.Invoke(ctx, PriceService_GetSymbolPrices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) GetStableCoins(ctx context.Context, in *StableCoinRequest, opts ...grpc.CallOption) (*StableCoinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StableCoinResponse)
	err := c.cc.Invoke(ctx, PriceService_GetStableCoins_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) GetStableCoinPrice(ctx context.Context, in *StableCoinPriceRequest, opts ...grpc.CallOption) (*StableCoinPriceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StableCoinPriceResponse)
	err := c.cc.Invoke(ctx, PriceService_GetStableCoinPrice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceServiceServer is the server API for PriceService service.
// All implementations should embed UnimplementedPriceServiceServer
// for forward compatibility.
type PriceServiceServer interface {
	GetExchanges(context.Context, *ExchangeRequest) (*ExchangeResponse, error)
	GetAssets(context.Context, *AssetRequest) (*AssetResponse, error)
	GetSymbols(context.Context, *SymbolRequest) (*SymbolResponse, error)
	GetSymbolPrices(context.Context, *SymbolPriceRequest) (*SymbolPriceResponse, error)
	GetStableCoins(context.Context, *StableCoinRequest) (*StableCoinResponse, error)
	GetStableCoinPrice(context.Context, *StableCoinPriceRequest) (*StableCoinPriceResponse, error)
}

// UnimplementedPriceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPriceServiceServer struct{}

func (UnimplementedPriceServiceServer) GetExchanges(context.Context, *ExchangeRequest) (*ExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchanges not implemented")
}
func (UnimplementedPriceServiceServer) GetAssets(context.Context, *AssetRequest) (*AssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssets not implemented")
}
func (UnimplementedPriceServiceServer) GetSymbols(context.Context, *SymbolRequest) (*SymbolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSymbols not implemented")
}
func (UnimplementedPriceServiceServer) GetSymbolPrices(context.Context, *SymbolPriceRequest) (*SymbolPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSymbolPrices not implemented")
}
func (UnimplementedPriceServiceServer) GetStableCoins(context.Context, *StableCoinRequest) (*StableCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStableCoins not implemented")
}
func (UnimplementedPriceServiceServer) GetStableCoinPrice(context.Context, *StableCoinPriceRequest) (*StableCoinPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStableCoinPrice not implemented")
}
func (UnimplementedPriceServiceServer) testEmbeddedByValue() {}

// UnsafePriceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceServiceServer will
// result in compilation errors.
type UnsafePriceServiceServer interface {
	mustEmbedUnimplementedPriceServiceServer()
}

func RegisterPriceServiceServer(s grpc.ServiceRegistrar, srv PriceServiceServer) {
	// If the following call pancis, it indicates UnimplementedPriceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PriceService_ServiceDesc, srv)
}

func _PriceService_GetExchanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetExchanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_GetExchanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetExchanges(ctx, req.(*ExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_GetAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_GetAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetAssets(ctx, req.(*AssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_GetSymbols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetSymbols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_GetSymbols_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetSymbols(ctx, req.(*SymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_GetSymbolPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymbolPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetSymbolPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_GetSymbolPrices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetSymbolPrices(ctx, req.(*SymbolPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_GetStableCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StableCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetStableCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_GetStableCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetStableCoins(ctx, req.(*StableCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_GetStableCoinPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StableCoinPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetStableCoinPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_GetStableCoinPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetStableCoinPrice(ctx, req.(*StableCoinPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceService_ServiceDesc is the grpc.ServiceDesc for PriceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapplink.market.PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getExchanges",
			Handler:    _PriceService_GetExchanges_Handler,
		},
		{
			MethodName: "getAssets",
			Handler:    _PriceService_GetAssets_Handler,
		},
		{
			MethodName: "getSymbols",
			Handler:    _PriceService_GetSymbols_Handler,
		},
		{
			MethodName: "getSymbolPrices",
			Handler:    _PriceService_GetSymbolPrices_Handler,
		},
		{
			MethodName: "getStableCoins",
			Handler:    _PriceService_GetStableCoins_Handler,
		},
		{
			MethodName: "getStableCoinPrice",
			Handler:    _PriceService_GetStableCoinPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapplink/market.proto",
}
