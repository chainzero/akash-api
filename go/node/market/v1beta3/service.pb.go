// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1beta3/service.proto

package v1beta3

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("akash/market/v1beta3/service.proto", fileDescriptor_0637a2f85fdb6b87)
}

var fileDescriptor_0637a2f85fdb6b87 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x01, 0xab, 0xd1, 0x83, 0xa8, 0xd1, 0x83, 0xaa, 0x91, 0x92, 0xc3, 0xaa, 0x33, 0x29, 0x33, 0x05,
	0xa2, 0x4b, 0x4a, 0x01, 0xab, 0x7c, 0x4e, 0x6a, 0x62, 0x31, 0xd4, 0x5c, 0xa3, 0x17, 0xcc, 0x5c,
	0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0xd1, 0x5c, 0x9c, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x4e, 0x99,
	0x29, 0x42, 0x4a, 0x7a, 0xd8, 0x6c, 0xd3, 0xf3, 0x2d, 0x4e, 0x87, 0xab, 0x91, 0xd2, 0x22, 0xac,
	0x26, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x28, 0x82, 0x8b, 0xc3, 0x39, 0x27, 0xbf,
	0x18, 0x6c, 0xb6, 0x22, 0x6e, 0x7d, 0x50, 0x25, 0x52, 0x9a, 0x04, 0x95, 0xc0, 0x4d, 0x4e, 0xe7,
	0xe2, 0x0d, 0xcf, 0x2c, 0xc9, 0x48, 0x29, 0x4a, 0x2c, 0xf7, 0x01, 0xf9, 0x4a, 0x48, 0x0d, 0xa7,
	0x5e, 0x14, 0x75, 0x52, 0x7a, 0xc4, 0xa9, 0x83, 0x5b, 0x94, 0xc8, 0xc5, 0x0d, 0xf1, 0x17, 0xc4,
	0x1a, 0x15, 0x02, 0xbe, 0x87, 0x58, 0xa2, 0x43, 0x8c, 0x2a, 0xb8, 0x15, 0x71, 0x5c, 0x5c, 0x60,
	0xff, 0x41, 0x6c, 0x50, 0xc6, 0x1f, 0x08, 0x10, 0x0b, 0xb4, 0x89, 0x50, 0x04, 0x33, 0xdf, 0x29,
	0xf8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x2c, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xc1, 0x06, 0xea, 0xe6, 0xa5, 0x96, 0x94, 0xe7, 0x17,
	0x65, 0x43, 0x79, 0x89, 0x05, 0x99, 0xfa, 0xe9, 0xf9, 0xfa, 0x79, 0xf9, 0x29, 0xa9, 0x68, 0x69,
	0x29, 0x89, 0x0d, 0x9c, 0x8c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x0e, 0xeb, 0xba,
	0xc4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateBid defines a method to create a bid given proper inputs.
	CreateBid(ctx context.Context, in *MsgCreateBid, opts ...grpc.CallOption) (*MsgCreateBidResponse, error)
	// CloseBid defines a method to close a bid given proper inputs.
	CloseBid(ctx context.Context, in *MsgCloseBid, opts ...grpc.CallOption) (*MsgCloseBidResponse, error)
	// WithdrawLease withdraws accrued funds from the lease payment
	WithdrawLease(ctx context.Context, in *MsgWithdrawLease, opts ...grpc.CallOption) (*MsgWithdrawLeaseResponse, error)
	// CreateLease creates a new lease
	CreateLease(ctx context.Context, in *MsgCreateLease, opts ...grpc.CallOption) (*MsgCreateLeaseResponse, error)
	// CloseLease defines a method to close an order given proper inputs.
	CloseLease(ctx context.Context, in *MsgCloseLease, opts ...grpc.CallOption) (*MsgCloseLeaseResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBid(ctx context.Context, in *MsgCreateBid, opts ...grpc.CallOption) (*MsgCreateBidResponse, error) {
	out := new(MsgCreateBidResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta3.Msg/CreateBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseBid(ctx context.Context, in *MsgCloseBid, opts ...grpc.CallOption) (*MsgCloseBidResponse, error) {
	out := new(MsgCloseBidResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta3.Msg/CloseBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawLease(ctx context.Context, in *MsgWithdrawLease, opts ...grpc.CallOption) (*MsgWithdrawLeaseResponse, error) {
	out := new(MsgWithdrawLeaseResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta3.Msg/WithdrawLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateLease(ctx context.Context, in *MsgCreateLease, opts ...grpc.CallOption) (*MsgCreateLeaseResponse, error) {
	out := new(MsgCreateLeaseResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta3.Msg/CreateLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CloseLease(ctx context.Context, in *MsgCloseLease, opts ...grpc.CallOption) (*MsgCloseLeaseResponse, error) {
	out := new(MsgCloseLeaseResponse)
	err := c.cc.Invoke(ctx, "/akash.market.v1beta3.Msg/CloseLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateBid defines a method to create a bid given proper inputs.
	CreateBid(context.Context, *MsgCreateBid) (*MsgCreateBidResponse, error)
	// CloseBid defines a method to close a bid given proper inputs.
	CloseBid(context.Context, *MsgCloseBid) (*MsgCloseBidResponse, error)
	// WithdrawLease withdraws accrued funds from the lease payment
	WithdrawLease(context.Context, *MsgWithdrawLease) (*MsgWithdrawLeaseResponse, error)
	// CreateLease creates a new lease
	CreateLease(context.Context, *MsgCreateLease) (*MsgCreateLeaseResponse, error)
	// CloseLease defines a method to close an order given proper inputs.
	CloseLease(context.Context, *MsgCloseLease) (*MsgCloseLeaseResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBid(ctx context.Context, req *MsgCreateBid) (*MsgCreateBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBid not implemented")
}
func (*UnimplementedMsgServer) CloseBid(ctx context.Context, req *MsgCloseBid) (*MsgCloseBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseBid not implemented")
}
func (*UnimplementedMsgServer) WithdrawLease(ctx context.Context, req *MsgWithdrawLease) (*MsgWithdrawLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawLease not implemented")
}
func (*UnimplementedMsgServer) CreateLease(ctx context.Context, req *MsgCreateLease) (*MsgCreateLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLease not implemented")
}
func (*UnimplementedMsgServer) CloseLease(ctx context.Context, req *MsgCloseLease) (*MsgCloseLeaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseLease not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta3.Msg/CreateBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBid(ctx, req.(*MsgCreateBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta3.Msg/CloseBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseBid(ctx, req.(*MsgCloseBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta3.Msg/WithdrawLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawLease(ctx, req.(*MsgWithdrawLease))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta3.Msg/CreateLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateLease(ctx, req.(*MsgCreateLease))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CloseLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCloseLease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CloseLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.market.v1beta3.Msg/CloseLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CloseLease(ctx, req.(*MsgCloseLease))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "akash.market.v1beta3.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBid",
			Handler:    _Msg_CreateBid_Handler,
		},
		{
			MethodName: "CloseBid",
			Handler:    _Msg_CloseBid_Handler,
		},
		{
			MethodName: "WithdrawLease",
			Handler:    _Msg_WithdrawLease_Handler,
		},
		{
			MethodName: "CreateLease",
			Handler:    _Msg_CreateLease_Handler,
		},
		{
			MethodName: "CloseLease",
			Handler:    _Msg_CloseLease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akash/market/v1beta3/service.proto",
}
