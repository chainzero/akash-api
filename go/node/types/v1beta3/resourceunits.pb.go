// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/base/v1beta3/resourceunits.proto

package v1beta3

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// ResourceUnits describes all available resources types for deployment/node etc
// if field is nil resource is not present in the given data-structure
type ResourceUnits struct {
	CPU       *CPU      `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory    *Memory   `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty" yaml:"memory,omitempty"`
	Storage   Volumes   `protobuf:"bytes,3,rep,name=storage,proto3,castrepeated=Volumes" json:"storage,omitempty" yaml:"storage,omitempty"`
	GPU       *GPU      `protobuf:"bytes,4,opt,name=gpu,proto3" json:"gpu,omitempty" yaml:"gpu,omitempty"`
	Endpoints Endpoints `protobuf:"bytes,5,rep,name=endpoints,proto3,castrepeated=Endpoints" json:"endpoints" yaml:"endpoints"`
}

func (m *ResourceUnits) Reset()         { *m = ResourceUnits{} }
func (m *ResourceUnits) String() string { return proto.CompactTextString(m) }
func (*ResourceUnits) ProtoMessage()    {}
func (*ResourceUnits) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4047260f8e6b24, []int{0}
}
func (m *ResourceUnits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceUnits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceUnits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceUnits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceUnits.Merge(m, src)
}
func (m *ResourceUnits) XXX_Size() int {
	return m.Size()
}
func (m *ResourceUnits) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceUnits.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceUnits proto.InternalMessageInfo

func (m *ResourceUnits) GetCPU() *CPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *ResourceUnits) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ResourceUnits) GetStorage() Volumes {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *ResourceUnits) GetGPU() *GPU {
	if m != nil {
		return m.GPU
	}
	return nil
}

func (m *ResourceUnits) GetEndpoints() Endpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceUnits)(nil), "akash.base.v1beta3.ResourceUnits")
}

func init() {
	proto.RegisterFile("akash/base/v1beta3/resourceunits.proto", fileDescriptor_8e4047260f8e6b24)
}

var fileDescriptor_8e4047260f8e6b24 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xb1, 0x6e, 0xd4, 0x30,
	0x1c, 0xc6, 0x2f, 0xe4, 0x68, 0xd5, 0x54, 0x95, 0x4a, 0x54, 0xa9, 0xd1, 0x51, 0xc5, 0x47, 0x06,
	0xd4, 0x01, 0x62, 0x71, 0xc7, 0x80, 0x3a, 0xa1, 0x54, 0xa8, 0x13, 0x52, 0x15, 0x74, 0x0c, 0x2c,
	0x28, 0x49, 0x2d, 0x37, 0x6a, 0x13, 0x5b, 0xb1, 0x4d, 0x95, 0xb7, 0xe0, 0x11, 0x58, 0xe1, 0x49,
	0x6e, 0xec, 0xc8, 0x64, 0x50, 0x6e, 0x41, 0x37, 0xf6, 0x09, 0x50, 0x62, 0xa7, 0xe5, 0x7a, 0xa6,
	0x5b, 0xe2, 0xef, 0xf7, 0xff, 0x7f, 0xf9, 0xbe, 0xd8, 0x79, 0x9e, 0x5c, 0x24, 0xec, 0x1c, 0xa6,
	0x09, 0x43, 0xf0, 0xcb, 0xab, 0x14, 0xf1, 0x64, 0x0a, 0x2b, 0xc4, 0x88, 0xa8, 0x32, 0x24, 0xca,
	0x9c, 0xb3, 0x90, 0x56, 0x84, 0x13, 0xd7, 0xed, 0xb8, 0xb0, 0xe5, 0x42, 0xcd, 0x8d, 0xf6, 0x30,
	0xc1, 0xa4, 0x93, 0x61, 0xfb, 0xa4, 0xc8, 0xd1, 0x81, 0x61, 0x63, 0x46, 0xc5, 0x03, 0x2a, 0xbe,
	0x55, 0x81, 0x41, 0x2d, 0x50, 0x41, 0xaa, 0x5a, 0x03, 0x63, 0x03, 0xc0, 0x38, 0xa9, 0x12, 0x8c,
	0x34, 0xf1, 0xcc, 0x40, 0xa0, 0xf2, 0x8c, 0x92, 0xbc, 0xe4, 0x0a, 0x09, 0xbe, 0x0f, 0x9d, 0x9d,
	0x58, 0x67, 0x9c, 0xb5, 0x19, 0xdd, 0xcf, 0x8e, 0x9d, 0x51, 0xe1, 0x59, 0x63, 0xeb, 0x70, 0x7b,
	0xb2, 0x1f, 0xae, 0x67, 0x0d, 0x8f, 0x4f, 0x67, 0xd1, 0xeb, 0xb9, 0x04, 0x56, 0x23, 0x81, 0x7d,
	0x7c, 0x3a, 0x5b, 0x4a, 0xb0, 0x93, 0x51, 0xf1, 0x82, 0x14, 0x39, 0x47, 0x05, 0xe5, 0xf5, 0x8d,
	0x04, 0x7b, 0x75, 0x52, 0x5c, 0x1e, 0x05, 0x2b, 0xc7, 0x41, 0xdc, 0x6e, 0x76, 0xb1, 0xb3, 0xa1,
	0x72, 0x78, 0x8f, 0x3a, 0x8f, 0x91, 0xc9, 0xe3, 0x7d, 0x47, 0x44, 0xd3, 0xd6, 0x66, 0x29, 0xc1,
	0xae, 0x9a, 0x58, 0xb1, 0xd8, 0x57, 0x16, 0xf7, 0x95, 0x20, 0xd6, 0xeb, 0xdd, 0x2b, 0x67, 0x53,
	0xf7, 0xe1, 0xd9, 0x63, 0xfb, 0x70, 0x7b, 0xf2, 0xd4, 0xe4, 0xf4, 0x41, 0x21, 0xd1, 0xdb, 0xb9,
	0x04, 0x83, 0xa5, 0x04, 0x4f, 0xf4, 0xcc, 0x8a, 0x97, 0xa7, 0xbc, 0xd6, 0xa4, 0xe0, 0xc7, 0x2f,
	0xb0, 0xf9, 0x91, 0x5c, 0x8a, 0x02, 0xb1, 0xb8, 0x77, 0x6b, 0x2b, 0xc4, 0x54, 0x78, 0xc3, 0xff,
	0x57, 0x78, 0xf2, 0x6f, 0x85, 0x27, 0xaa, 0x42, 0x6c, 0xae, 0x10, 0xdf, 0xab, 0x10, 0x53, 0xe1,
	0x96, 0xce, 0x56, 0xff, 0x1f, 0x99, 0xf7, 0xb8, 0xcb, 0x76, 0x60, 0xb2, 0x79, 0xa7, 0xa1, 0x68,
	0xa2, 0xc3, 0xdd, 0x8d, 0xdd, 0x48, 0xb0, 0xab, 0x0c, 0x6e, 0x8f, 0xda, 0x30, 0x5b, 0xfd, 0x08,
	0x8b, 0xef, 0xd8, 0xa3, 0xe1, 0x9f, 0x6f, 0xc0, 0x8a, 0xe2, 0x79, 0xe3, 0x5b, 0xd7, 0x8d, 0x6f,
	0xfd, 0x6e, 0x7c, 0xeb, 0xeb, 0xc2, 0x1f, 0x5c, 0x2f, 0xfc, 0xc1, 0xcf, 0x85, 0x3f, 0xf8, 0xf4,
	0x06, 0xe7, 0xfc, 0x5c, 0xa4, 0x61, 0x46, 0x0a, 0xd8, 0x7d, 0xc6, 0xcb, 0x12, 0xf1, 0x2b, 0x52,
	0x5d, 0xe8, 0xb7, 0x84, 0xe6, 0x10, 0x13, 0x58, 0x92, 0x33, 0x04, 0x79, 0x4d, 0x11, 0xeb, 0xaf,
	0x63, 0xba, 0xd1, 0x5d, 0xc3, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0xbb, 0x74, 0xd7,
	0x7c, 0x03, 0x00, 0x00,
}

func (this *ResourceUnits) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResourceUnits)
	if !ok {
		that2, ok := that.(ResourceUnits)
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
	if !this.CPU.Equal(that1.CPU) {
		return false
	}
	if !this.Memory.Equal(that1.Memory) {
		return false
	}
	if len(this.Storage) != len(that1.Storage) {
		return false
	}
	for i := range this.Storage {
		if !this.Storage[i].Equal(&that1.Storage[i]) {
			return false
		}
	}
	if !this.GPU.Equal(that1.GPU) {
		return false
	}
	if len(this.Endpoints) != len(that1.Endpoints) {
		return false
	}
	for i := range this.Endpoints {
		if !this.Endpoints[i].Equal(&that1.Endpoints[i]) {
			return false
		}
	}
	return true
}
func (m *ResourceUnits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceUnits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceUnits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResourceunits(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.GPU != nil {
		{
			size, err := m.GPU.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResourceunits(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Storage) > 0 {
		for iNdEx := len(m.Storage) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Storage[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResourceunits(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Memory != nil {
		{
			size, err := m.Memory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResourceunits(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CPU != nil {
		{
			size, err := m.CPU.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResourceunits(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResourceunits(dAtA []byte, offset int, v uint64) int {
	offset -= sovResourceunits(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourceUnits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CPU != nil {
		l = m.CPU.Size()
		n += 1 + l + sovResourceunits(uint64(l))
	}
	if m.Memory != nil {
		l = m.Memory.Size()
		n += 1 + l + sovResourceunits(uint64(l))
	}
	if len(m.Storage) > 0 {
		for _, e := range m.Storage {
			l = e.Size()
			n += 1 + l + sovResourceunits(uint64(l))
		}
	}
	if m.GPU != nil {
		l = m.GPU.Size()
		n += 1 + l + sovResourceunits(uint64(l))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovResourceunits(uint64(l))
		}
	}
	return n
}

func sovResourceunits(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResourceunits(x uint64) (n int) {
	return sovResourceunits(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceUnits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceunits
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
			return fmt.Errorf("proto: ResourceUnits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceUnits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPU", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CPU == nil {
				m.CPU = &CPU{}
			}
			if err := m.CPU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Memory == nil {
				m.Memory = &Memory{}
			}
			if err := m.Memory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = append(m.Storage, Storage{})
			if err := m.Storage[len(m.Storage)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GPU", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GPU == nil {
				m.GPU = &GPU{}
			}
			if err := m.GPU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceunits(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResourceunits
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
func skipResourceunits(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceunits
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
					return 0, ErrIntOverflowResourceunits
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
					return 0, ErrIntOverflowResourceunits
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
				return 0, ErrInvalidLengthResourceunits
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResourceunits
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResourceunits
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResourceunits        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceunits          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResourceunits = fmt.Errorf("proto: unexpected end of group")
)
