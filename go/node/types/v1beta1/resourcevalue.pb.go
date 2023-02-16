// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/base/v1beta1/resourcevalue.proto

package v1beta1

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Unit stores cpu, memory and storage metrics
type ResourceValue struct {
	Val github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=val,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"val"`
}

func (m *ResourceValue) Reset()         { *m = ResourceValue{} }
func (m *ResourceValue) String() string { return proto.CompactTextString(m) }
func (*ResourceValue) ProtoMessage()    {}
func (*ResourceValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_85efdc01289cdacd, []int{0}
}
func (m *ResourceValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceValue.Merge(m, src)
}
func (m *ResourceValue) XXX_Size() int {
	return m.Size()
}
func (m *ResourceValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceValue.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceValue proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ResourceValue)(nil), "akash.base.v1beta1.ResourceValue")
}

func init() {
	proto.RegisterFile("akash/base/v1beta1/resourcevalue.proto", fileDescriptor_85efdc01289cdacd)
}

var fileDescriptor_85efdc01289cdacd = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xcc, 0x4e, 0x2c,
	0xce, 0xd0, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f,
	0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x02, 0xab, 0xd3, 0x03, 0xa9, 0xd3, 0x83, 0xaa, 0x93, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x4a, 0xe1, 0x5c, 0xbc, 0x41, 0x50, 0x03,
	0xc2, 0x40, 0x06, 0x08, 0x39, 0x70, 0x31, 0x97, 0x25, 0xe6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0,
	0x38, 0xe9, 0x9d, 0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x5a, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x31, 0x94, 0xd2, 0x2d, 0x4e,
	0xc9, 0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0xf3, 0xcc, 0x2b, 0x09, 0x02, 0x69, 0xb5, 0x62,
	0x79, 0xb1, 0x40, 0x9e, 0xd1, 0x29, 0xe8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x2c, 0x90, 0x0c, 0x03, 0xbb, 0x53, 0x37, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x1b, 0xca,
	0x4b, 0x2c, 0xc8, 0xd4, 0x4f, 0xcf, 0xd7, 0xcf, 0xcb, 0x4f, 0x49, 0x85, 0x18, 0x0d, 0xf3, 0x6a,
	0x12, 0x1b, 0xd8, 0xcd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xf5, 0x4a, 0x29, 0x07,
	0x01, 0x00, 0x00,
}

func (this *ResourceValue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResourceValue)
	if !ok {
		that2, ok := that.(ResourceValue)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Val.Equal(that1.Val) {
		return false
	}
	return true
}
func (m *ResourceValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Val.Size()
		i -= size
		if _, err := m.Val.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintResourcevalue(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintResourcevalue(dAtA []byte, offset int, v uint64) int {
	offset -= sovResourcevalue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourceValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Val.Size()
	n += 1 + l + sovResourcevalue(uint64(l))
	return n
}

func sovResourcevalue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResourcevalue(x uint64) (n int) {
	return sovResourcevalue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourcevalue
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
			return fmt.Errorf("proto: ResourceValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourcevalue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResourcevalue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResourcevalue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Val.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourcevalue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResourcevalue
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
func skipResourcevalue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourcevalue
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
					return 0, ErrIntOverflowResourcevalue
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
					return 0, ErrIntOverflowResourcevalue
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
				return 0, ErrInvalidLengthResourcevalue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResourcevalue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResourcevalue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResourcevalue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourcevalue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResourcevalue = fmt.Errorf("proto: unexpected end of group")
)
