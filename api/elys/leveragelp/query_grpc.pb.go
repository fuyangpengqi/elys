// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package leveragelp

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
	// Queries a list of GetPositions items.
	QueryPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error)
	// Queries a list of GetPositionsByPool items.
	QueryPositionsByPool(ctx context.Context, in *PositionsByPoolRequest, opts ...grpc.CallOption) (*PositionsByPoolResponse, error)
	// Queries a list of GetStatus items.
	GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// Queries a list of GetPositionsForAddress items.
	QueryPositionsForAddress(ctx context.Context, in *PositionsForAddressRequest, opts ...grpc.CallOption) (*PositionsForAddressResponse, error)
	// Queries a list of GetWhitelist items.
	GetWhitelist(ctx context.Context, in *WhitelistRequest, opts ...grpc.CallOption) (*WhitelistResponse, error)
	// Queries a list of IsWhitelisted items.
	IsWhitelisted(ctx context.Context, in *IsWhitelistedRequest, opts ...grpc.CallOption) (*IsWhitelistedResponse, error)
	// Queries a list of Pool items.
	Pool(ctx context.Context, in *QueryGetPoolRequest, opts ...grpc.CallOption) (*QueryGetPoolResponse, error)
	Pools(ctx context.Context, in *QueryAllPoolRequest, opts ...grpc.CallOption) (*QueryAllPoolResponse, error)
	// Queries a list of Position items.
	Position(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionResponse, error)
	// Query a liquidation price for a position
	LiquidationPrice(ctx context.Context, in *QueryLiquidationPriceRequest, opts ...grpc.CallOption) (*QueryLiquidationPriceResponse, error)
	// Get estimated amount of return value opening a position
	OpenEst(ctx context.Context, in *QueryOpenEstRequest, opts ...grpc.CallOption) (*QueryOpenEstResponse, error)
	// Get estimated amount of return value closing a position
	CloseEst(ctx context.Context, in *QueryCloseEstRequest, opts ...grpc.CallOption) (*QueryCloseEstResponse, error)
	// Queries rewards on leveragelp
	Rewards(ctx context.Context, in *QueryRewardsRequest, opts ...grpc.CallOption) (*QueryRewardsResponse, error)
	// Queries a list of CommittedTokensLocked items.
	CommittedTokensLocked(ctx context.Context, in *QueryCommittedTokensLockedRequest, opts ...grpc.CallOption) (*QueryCommittedTokensLockedResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error) {
	out := new(PositionsResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/QueryPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPositionsByPool(ctx context.Context, in *PositionsByPoolRequest, opts ...grpc.CallOption) (*PositionsByPoolResponse, error) {
	out := new(PositionsByPoolResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/QueryPositionsByPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryPositionsForAddress(ctx context.Context, in *PositionsForAddressRequest, opts ...grpc.CallOption) (*PositionsForAddressResponse, error) {
	out := new(PositionsForAddressResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/QueryPositionsForAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetWhitelist(ctx context.Context, in *WhitelistRequest, opts ...grpc.CallOption) (*WhitelistResponse, error) {
	out := new(WhitelistResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/GetWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IsWhitelisted(ctx context.Context, in *IsWhitelistedRequest, opts ...grpc.CallOption) (*IsWhitelistedResponse, error) {
	out := new(IsWhitelistedResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/IsWhitelisted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pool(ctx context.Context, in *QueryGetPoolRequest, opts ...grpc.CallOption) (*QueryGetPoolResponse, error) {
	out := new(QueryGetPoolResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/Pool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pools(ctx context.Context, in *QueryAllPoolRequest, opts ...grpc.CallOption) (*QueryAllPoolResponse, error) {
	out := new(QueryAllPoolResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/Pools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Position(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionResponse, error) {
	out := new(PositionResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/Position", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidationPrice(ctx context.Context, in *QueryLiquidationPriceRequest, opts ...grpc.CallOption) (*QueryLiquidationPriceResponse, error) {
	out := new(QueryLiquidationPriceResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/LiquidationPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) OpenEst(ctx context.Context, in *QueryOpenEstRequest, opts ...grpc.CallOption) (*QueryOpenEstResponse, error) {
	out := new(QueryOpenEstResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/OpenEst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CloseEst(ctx context.Context, in *QueryCloseEstRequest, opts ...grpc.CallOption) (*QueryCloseEstResponse, error) {
	out := new(QueryCloseEstResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/CloseEst", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Rewards(ctx context.Context, in *QueryRewardsRequest, opts ...grpc.CallOption) (*QueryRewardsResponse, error) {
	out := new(QueryRewardsResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/Rewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CommittedTokensLocked(ctx context.Context, in *QueryCommittedTokensLockedRequest, opts ...grpc.CallOption) (*QueryCommittedTokensLockedResponse, error) {
	out := new(QueryCommittedTokensLockedResponse)
	err := c.cc.Invoke(ctx, "/elys.leveragelp.Query/CommittedTokensLocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
	// Queries a list of GetPositions items.
	QueryPositions(context.Context, *PositionsRequest) (*PositionsResponse, error)
	// Queries a list of GetPositionsByPool items.
	QueryPositionsByPool(context.Context, *PositionsByPoolRequest) (*PositionsByPoolResponse, error)
	// Queries a list of GetStatus items.
	GetStatus(context.Context, *StatusRequest) (*StatusResponse, error)
	// Queries a list of GetPositionsForAddress items.
	QueryPositionsForAddress(context.Context, *PositionsForAddressRequest) (*PositionsForAddressResponse, error)
	// Queries a list of GetWhitelist items.
	GetWhitelist(context.Context, *WhitelistRequest) (*WhitelistResponse, error)
	// Queries a list of IsWhitelisted items.
	IsWhitelisted(context.Context, *IsWhitelistedRequest) (*IsWhitelistedResponse, error)
	// Queries a list of Pool items.
	Pool(context.Context, *QueryGetPoolRequest) (*QueryGetPoolResponse, error)
	Pools(context.Context, *QueryAllPoolRequest) (*QueryAllPoolResponse, error)
	// Queries a list of Position items.
	Position(context.Context, *PositionRequest) (*PositionResponse, error)
	// Query a liquidation price for a position
	LiquidationPrice(context.Context, *QueryLiquidationPriceRequest) (*QueryLiquidationPriceResponse, error)
	// Get estimated amount of return value opening a position
	OpenEst(context.Context, *QueryOpenEstRequest) (*QueryOpenEstResponse, error)
	// Get estimated amount of return value closing a position
	CloseEst(context.Context, *QueryCloseEstRequest) (*QueryCloseEstResponse, error)
	// Queries rewards on leveragelp
	Rewards(context.Context, *QueryRewardsRequest) (*QueryRewardsResponse, error)
	// Queries a list of CommittedTokensLocked items.
	CommittedTokensLocked(context.Context, *QueryCommittedTokensLockedRequest) (*QueryCommittedTokensLockedResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) QueryPositions(context.Context, *PositionsRequest) (*PositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPositions not implemented")
}
func (UnimplementedQueryServer) QueryPositionsByPool(context.Context, *PositionsByPoolRequest) (*PositionsByPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPositionsByPool not implemented")
}
func (UnimplementedQueryServer) GetStatus(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedQueryServer) QueryPositionsForAddress(context.Context, *PositionsForAddressRequest) (*PositionsForAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPositionsForAddress not implemented")
}
func (UnimplementedQueryServer) GetWhitelist(context.Context, *WhitelistRequest) (*WhitelistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhitelist not implemented")
}
func (UnimplementedQueryServer) IsWhitelisted(context.Context, *IsWhitelistedRequest) (*IsWhitelistedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsWhitelisted not implemented")
}
func (UnimplementedQueryServer) Pool(context.Context, *QueryGetPoolRequest) (*QueryGetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (UnimplementedQueryServer) Pools(context.Context, *QueryAllPoolRequest) (*QueryAllPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pools not implemented")
}
func (UnimplementedQueryServer) Position(context.Context, *PositionRequest) (*PositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}
func (UnimplementedQueryServer) LiquidationPrice(context.Context, *QueryLiquidationPriceRequest) (*QueryLiquidationPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidationPrice not implemented")
}
func (UnimplementedQueryServer) OpenEst(context.Context, *QueryOpenEstRequest) (*QueryOpenEstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenEst not implemented")
}
func (UnimplementedQueryServer) CloseEst(context.Context, *QueryCloseEstRequest) (*QueryCloseEstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseEst not implemented")
}
func (UnimplementedQueryServer) Rewards(context.Context, *QueryRewardsRequest) (*QueryRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rewards not implemented")
}
func (UnimplementedQueryServer) CommittedTokensLocked(context.Context, *QueryCommittedTokensLockedRequest) (*QueryCommittedTokensLockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommittedTokensLocked not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/QueryPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPositions(ctx, req.(*PositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPositionsByPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsByPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPositionsByPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/QueryPositionsByPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPositionsByPool(ctx, req.(*PositionsByPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryPositionsForAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsForAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPositionsForAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/QueryPositionsForAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPositionsForAddress(ctx, req.(*PositionsForAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhitelistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/GetWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetWhitelist(ctx, req.(*WhitelistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IsWhitelisted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsWhitelistedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IsWhitelisted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/IsWhitelisted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IsWhitelisted(ctx, req.(*IsWhitelistedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/Pool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryGetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/Pools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pools(ctx, req.(*QueryAllPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/Position",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Position(ctx, req.(*PositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidationPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidationPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidationPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/LiquidationPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidationPrice(ctx, req.(*QueryLiquidationPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_OpenEst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOpenEstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).OpenEst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/OpenEst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).OpenEst(ctx, req.(*QueryOpenEstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CloseEst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCloseEstRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CloseEst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/CloseEst",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CloseEst(ctx, req.(*QueryCloseEstRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Rewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Rewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/Rewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Rewards(ctx, req.(*QueryRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CommittedTokensLocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommittedTokensLockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommittedTokensLocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elys.leveragelp.Query/CommittedTokensLocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommittedTokensLocked(ctx, req.(*QueryCommittedTokensLockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "elys.leveragelp.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QueryPositions",
			Handler:    _Query_QueryPositions_Handler,
		},
		{
			MethodName: "QueryPositionsByPool",
			Handler:    _Query_QueryPositionsByPool_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Query_GetStatus_Handler,
		},
		{
			MethodName: "QueryPositionsForAddress",
			Handler:    _Query_QueryPositionsForAddress_Handler,
		},
		{
			MethodName: "GetWhitelist",
			Handler:    _Query_GetWhitelist_Handler,
		},
		{
			MethodName: "IsWhitelisted",
			Handler:    _Query_IsWhitelisted_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "Pools",
			Handler:    _Query_Pools_Handler,
		},
		{
			MethodName: "Position",
			Handler:    _Query_Position_Handler,
		},
		{
			MethodName: "LiquidationPrice",
			Handler:    _Query_LiquidationPrice_Handler,
		},
		{
			MethodName: "OpenEst",
			Handler:    _Query_OpenEst_Handler,
		},
		{
			MethodName: "CloseEst",
			Handler:    _Query_CloseEst_Handler,
		},
		{
			MethodName: "Rewards",
			Handler:    _Query_Rewards_Handler,
		},
		{
			MethodName: "CommittedTokensLocked",
			Handler:    _Query_CommittedTokensLocked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elys/leveragelp/query.proto",
}
