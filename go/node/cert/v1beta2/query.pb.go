// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/cert/v1beta2/query.proto

package v1beta2

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// CertificateResponse contains a single X509 certificate and its serial number
type CertificateResponse struct {
	Certificate Certificate `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate" yaml:"certificate"`
	Serial      string      `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial" yaml:"serial"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56dee19acb66f387, []int{0}
}
func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificate() Certificate {
	if m != nil {
		return m.Certificate
	}
	return Certificate{}
}

func (m *CertificateResponse) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

// QueryDeploymentsRequest is request type for the Query/Deployments RPC method
type QueryCertificatesRequest struct {
	Filter     CertificateFilter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCertificatesRequest) Reset()         { *m = QueryCertificatesRequest{} }
func (m *QueryCertificatesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCertificatesRequest) ProtoMessage()    {}
func (*QueryCertificatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56dee19acb66f387, []int{1}
}
func (m *QueryCertificatesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificatesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificatesRequest.Merge(m, src)
}
func (m *QueryCertificatesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificatesRequest proto.InternalMessageInfo

func (m *QueryCertificatesRequest) GetFilter() CertificateFilter {
	if m != nil {
		return m.Filter
	}
	return CertificateFilter{}
}

func (m *QueryCertificatesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCertificatesResponse is response type for the Query/Certificates RPC method
type QueryCertificatesResponse struct {
	Certificates CertificatesResponse `protobuf:"bytes,1,rep,name=certificates,proto3,castrepeated=CertificatesResponse" json:"certificates"`
	Pagination   *query.PageResponse  `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCertificatesResponse) Reset()         { *m = QueryCertificatesResponse{} }
func (m *QueryCertificatesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCertificatesResponse) ProtoMessage()    {}
func (*QueryCertificatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56dee19acb66f387, []int{2}
}
func (m *QueryCertificatesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificatesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificatesResponse.Merge(m, src)
}
func (m *QueryCertificatesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificatesResponse proto.InternalMessageInfo

func (m *QueryCertificatesResponse) GetCertificates() CertificatesResponse {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *QueryCertificatesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateResponse)(nil), "akash.cert.v1beta2.CertificateResponse")
	proto.RegisterType((*QueryCertificatesRequest)(nil), "akash.cert.v1beta2.QueryCertificatesRequest")
	proto.RegisterType((*QueryCertificatesResponse)(nil), "akash.cert.v1beta2.QueryCertificatesResponse")
}

func init() { proto.RegisterFile("akash/cert/v1beta2/query.proto", fileDescriptor_56dee19acb66f387) }

var fileDescriptor_56dee19acb66f387 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x01, 0x95, 0x70, 0xc7, 0xc5, 0xec, 0x50, 0xca, 0x48, 0xaa, 0x48, 0xa3, 0x05,
	0x51, 0x5b, 0x4b, 0x0f, 0x48, 0x1c, 0x33, 0x69, 0x5c, 0x59, 0x8e, 0xdc, 0xdc, 0xf0, 0x2d, 0xb3,
	0x96, 0xc6, 0x5d, 0xec, 0x82, 0x76, 0xe5, 0x09, 0x90, 0xb8, 0x71, 0x46, 0x42, 0xe2, 0x09, 0x78,
	0x84, 0xdd, 0x98, 0xc4, 0x85, 0x53, 0x40, 0x2d, 0xa7, 0x1d, 0xf7, 0x04, 0x28, 0x8e, 0xd1, 0x3c,
	0x35, 0xa8, 0xbb, 0xc5, 0xfe, 0x7f, 0xff, 0xcf, 0xbf, 0xff, 0xe7, 0x18, 0x7b, 0xfc, 0x98, 0xab,
	0x23, 0x96, 0x40, 0xa1, 0xd9, 0xdb, 0xdd, 0x09, 0x68, 0x1e, 0xb2, 0x93, 0x39, 0x14, 0xa7, 0x74,
	0x56, 0x48, 0x2d, 0x09, 0x31, 0x3a, 0xad, 0x74, 0x6a, 0xf5, 0xde, 0x56, 0x2a, 0x53, 0x69, 0x64,
	0x56, 0x7d, 0xd5, 0x95, 0xbd, 0xed, 0x54, 0xca, 0x34, 0x03, 0xc6, 0x67, 0x82, 0xf1, 0x3c, 0x97,
	0x9a, 0x6b, 0x21, 0x73, 0x65, 0xd5, 0xa7, 0x89, 0x54, 0x53, 0xa9, 0xd8, 0x84, 0x2b, 0xa8, 0x0f,
	0xb0, 0xc7, 0xed, 0xb2, 0x19, 0x4f, 0x45, 0x6e, 0x8a, 0x6d, 0xed, 0xa3, 0x06, 0x26, 0x03, 0x60,
	0xe4, 0xe0, 0x1b, 0xc2, 0xf7, 0xf7, 0xa0, 0xd0, 0xe2, 0x50, 0x24, 0x5c, 0x43, 0x0c, 0x6a, 0x26,
	0x73, 0x05, 0x24, 0xc3, 0x9d, 0xe4, 0x6a, 0xbb, 0x8b, 0xfa, 0x68, 0xd8, 0x09, 0x7d, 0xba, 0x1a,
	0x80, 0x3a, 0xee, 0xe8, 0xc9, 0x59, 0xe9, 0xb7, 0x2e, 0x4a, 0xdf, 0xf5, 0x5e, 0x96, 0x3e, 0x39,
	0xe5, 0xd3, 0xec, 0x45, 0xe0, 0x6c, 0x06, 0xb1, 0x5b, 0x42, 0xc6, 0xb8, 0xad, 0xa0, 0x10, 0x3c,
	0xeb, 0x6e, 0xf4, 0xd1, 0xf0, 0x6e, 0xf4, 0xf0, 0xa2, 0xf4, 0xed, 0xce, 0x65, 0xe9, 0xdf, 0xab,
	0xed, 0xf5, 0x3a, 0x88, 0xad, 0x10, 0x7c, 0x41, 0xb8, 0x7b, 0x50, 0x85, 0x77, 0x08, 0x54, 0x0c,
	0x27, 0x73, 0x50, 0x9a, 0xec, 0xe1, 0xf6, 0xa1, 0xc8, 0x34, 0x14, 0x16, 0x7d, 0x67, 0x0d, 0xfa,
	0xbe, 0x29, 0x8e, 0x6e, 0x57, 0x01, 0x62, 0x6b, 0x25, 0xfb, 0x18, 0x5f, 0xcd, 0xd3, 0xa0, 0x75,
	0xc2, 0xc7, 0xb4, 0x1e, 0x3e, 0xad, 0x86, 0x4f, 0xeb, 0xdb, 0xb5, 0xc3, 0xa7, 0xaf, 0x78, 0x0a,
	0x16, 0x20, 0x76, 0x9c, 0xc1, 0x77, 0x84, 0x1f, 0x34, 0x90, 0xda, 0x51, 0x0b, 0xbc, 0xe9, 0xcc,
	0x42, 0x75, 0x51, 0xff, 0xd6, 0xb0, 0x13, 0x0e, 0xd6, 0x00, 0xff, 0xb3, 0x47, 0xdb, 0x15, 0xf2,
	0xd7, 0x5f, 0xfe, 0x56, 0x53, 0xf3, 0xf8, 0x5a, 0x6b, 0xf2, 0xb2, 0x21, 0xd0, 0x60, 0x6d, 0x20,
	0xdb, 0xca, 0xb1, 0x86, 0x9f, 0x11, 0xbe, 0x63, 0x12, 0x91, 0x4f, 0x08, 0x6f, 0xba, 0x27, 0x93,
	0x67, 0x4d, 0xe0, 0xff, 0xbb, 0xa7, 0xde, 0xe8, 0x86, 0xd5, 0x35, 0x43, 0x30, 0x7a, 0xff, 0xe3,
	0xcf, 0xc7, 0x8d, 0x01, 0xd9, 0x61, 0x2b, 0xbf, 0xf5, 0x98, 0xb9, 0x51, 0x59, 0x26, 0x94, 0x8e,
	0x0e, 0xce, 0x16, 0x1e, 0x3a, 0x5f, 0x78, 0xe8, 0xf7, 0xc2, 0x43, 0x1f, 0x96, 0x5e, 0xeb, 0x7c,
	0xe9, 0xb5, 0x7e, 0x2e, 0xbd, 0xd6, 0xeb, 0xe7, 0xa9, 0xd0, 0x47, 0xf3, 0x09, 0x4d, 0xe4, 0xb4,
	0x6e, 0x35, 0xca, 0x41, 0xbf, 0x93, 0xc5, 0xb1, 0x5d, 0x55, 0x0f, 0x2f, 0x95, 0x2c, 0x97, 0x6f,
	0xe0, 0xda, 0xdb, 0x99, 0xb4, 0xcd, 0xbb, 0x19, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xf4,
	0xa0, 0xc2, 0xec, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Certificates queries certificates
	Certificates(ctx context.Context, in *QueryCertificatesRequest, opts ...grpc.CallOption) (*QueryCertificatesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Certificates(ctx context.Context, in *QueryCertificatesRequest, opts ...grpc.CallOption) (*QueryCertificatesResponse, error) {
	out := new(QueryCertificatesResponse)
	err := c.cc.Invoke(ctx, "/akash.cert.v1beta2.Query/Certificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Certificates queries certificates
	Certificates(context.Context, *QueryCertificatesRequest) (*QueryCertificatesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Certificates(ctx context.Context, req *QueryCertificatesRequest) (*QueryCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Certificates not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Certificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Certificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.cert.v1beta2.Query/Certificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Certificates(ctx, req.(*QueryCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "akash.cert.v1beta2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certificates",
			Handler:    _Query_Certificates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akash/cert/v1beta2/query.proto",
}

func (m *CertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Serial) > 0 {
		i -= len(m.Serial)
		copy(dAtA[i:], m.Serial)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Serial)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Certificate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCertificatesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificatesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificatesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Filter.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCertificatesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificatesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificatesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Certificates) > 0 {
		for iNdEx := len(m.Certificates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certificates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Certificate.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = len(m.Serial)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCertificatesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Filter.Size()
	n += 1 + l + sovQuery(uint64(l))
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCertificatesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Certificate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Serial", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Serial = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCertificatesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertificatesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificatesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Filter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCertificatesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertificatesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificatesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificates = append(m.Certificates, CertificateResponse{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
