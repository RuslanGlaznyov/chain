// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/registry/v1beta1/params.proto

package types

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

// Params defines the set of params for the registry module.
type Params struct {
	// vote_slash ...
	VoteSlash string `protobuf:"bytes,3,opt,name=vote_slash,json=voteSlash,proto3" json:"vote_slash,omitempty"`
	// upload_slash ...
	UploadSlash string `protobuf:"bytes,4,opt,name=upload_slash,json=uploadSlash,proto3" json:"upload_slash,omitempty"`
	// timeout_slash ...
	TimeoutSlash string `protobuf:"bytes,5,opt,name=timeout_slash,json=timeoutSlash,proto3" json:"timeout_slash,omitempty"`
	// upload_timeout ...
	UploadTimeout uint64 `protobuf:"varint,6,opt,name=upload_timeout,json=uploadTimeout,proto3" json:"upload_timeout,omitempty"`
	// storage_cost ...
	StorageCost uint64 `protobuf:"varint,7,opt,name=storage_cost,json=storageCost,proto3" json:"storage_cost,omitempty"`
	// network_fee ...
	NetworkFee string `protobuf:"bytes,8,opt,name=network_fee,json=networkFee,proto3" json:"network_fee,omitempty"`
	// max_points ...
	MaxPoints uint64 `protobuf:"varint,9,opt,name=max_points,json=maxPoints,proto3" json:"max_points,omitempty"`
	// unbonding_staking_time ...
	UnbondingStakingTime uint64 `protobuf:"varint,10,opt,name=unbonding_staking_time,json=unbondingStakingTime,proto3" json:"unbonding_staking_time,omitempty"`
	// unbonding_delegation_time ...
	UnbondingDelegationTime uint64 `protobuf:"varint,11,opt,name=unbonding_delegation_time,json=unbondingDelegationTime,proto3" json:"unbonding_delegation_time,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca08e39f277f4aef, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetVoteSlash() string {
	if m != nil {
		return m.VoteSlash
	}
	return ""
}

func (m *Params) GetUploadSlash() string {
	if m != nil {
		return m.UploadSlash
	}
	return ""
}

func (m *Params) GetTimeoutSlash() string {
	if m != nil {
		return m.TimeoutSlash
	}
	return ""
}

func (m *Params) GetUploadTimeout() uint64 {
	if m != nil {
		return m.UploadTimeout
	}
	return 0
}

func (m *Params) GetStorageCost() uint64 {
	if m != nil {
		return m.StorageCost
	}
	return 0
}

func (m *Params) GetNetworkFee() string {
	if m != nil {
		return m.NetworkFee
	}
	return ""
}

func (m *Params) GetMaxPoints() uint64 {
	if m != nil {
		return m.MaxPoints
	}
	return 0
}

func (m *Params) GetUnbondingStakingTime() uint64 {
	if m != nil {
		return m.UnbondingStakingTime
	}
	return 0
}

func (m *Params) GetUnbondingDelegationTime() uint64 {
	if m != nil {
		return m.UnbondingDelegationTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "kyve.registry.v1beta1.Params")
}

func init() {
	proto.RegisterFile("kyve/registry/v1beta1/params.proto", fileDescriptor_ca08e39f277f4aef)
}

var fileDescriptor_ca08e39f277f4aef = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0xc7, 0xed, 0x96, 0xd2, 0x7a, 0x0d, 0x3d, 0x58, 0xb4, 0x75, 0x2b, 0xd5, 0x50, 0xaa, 0x48,
	0x5c, 0x62, 0x0b, 0x25, 0x27, 0x8e, 0xf9, 0x3c, 0x44, 0x8a, 0x10, 0x44, 0x91, 0x92, 0x8b, 0xb5,
	0x86, 0x8d, 0xb1, 0xc0, 0x1e, 0xcb, 0x3b, 0x26, 0xf0, 0x16, 0x39, 0xe6, 0x98, 0xc7, 0xc9, 0x91,
	0x63, 0x8e, 0x11, 0x3c, 0x41, 0xde, 0x20, 0xda, 0x8f, 0xc0, 0xc9, 0xd6, 0x6f, 0x7e, 0xff, 0x9d,
	0x19, 0x0d, 0x69, 0x4f, 0x97, 0x73, 0x16, 0x14, 0x2c, 0x4e, 0x38, 0x16, 0xcb, 0x60, 0xde, 0x8d,
	0x18, 0xd2, 0x6e, 0x90, 0xd3, 0x82, 0xa6, 0xdc, 0xcf, 0x0b, 0x40, 0x70, 0x7e, 0x08, 0xc7, 0xff,
	0x70, 0x7c, 0xed, 0xfc, 0x69, 0xc4, 0x10, 0x83, 0x34, 0x02, 0xf1, 0xa7, 0xe4, 0xf6, 0xdb, 0x27,
	0x52, 0xed, 0xcb, 0xb4, 0xf3, 0x97, 0x90, 0x39, 0x20, 0x0b, 0xf9, 0x8c, 0xf2, 0x89, 0xfb, 0xb9,
	0x65, 0x76, 0xac, 0x81, 0x25, 0xc8, 0x50, 0x00, 0xe7, 0x1f, 0xa9, 0x95, 0xf9, 0x0c, 0xe8, 0x58,
	0x0b, 0x15, 0x29, 0xd8, 0x8a, 0x29, 0xe5, 0x3f, 0xa9, 0x63, 0x92, 0x32, 0x28, 0x51, 0x3b, 0x5f,
	0xa4, 0x53, 0xd3, 0x50, 0x49, 0x7b, 0xe4, 0xbb, 0x7e, 0x47, 0x63, 0xb7, 0xda, 0x32, 0x3b, 0x95,
	0x41, 0x5d, 0xd1, 0x2b, 0x05, 0x45, 0x3b, 0x8e, 0x50, 0xd0, 0x98, 0x85, 0x23, 0xe0, 0xe8, 0x7e,
	0x95, 0x92, 0xad, 0xd9, 0x31, 0x70, 0x74, 0x9a, 0xc4, 0xce, 0x18, 0xde, 0x43, 0x31, 0x0d, 0xef,
	0x18, 0x73, 0xbf, 0xc9, 0x66, 0x44, 0xa3, 0x33, 0xc6, 0xc4, 0x46, 0x29, 0x5d, 0x84, 0x39, 0x24,
	0x19, 0x72, 0xd7, 0x92, 0x2f, 0x58, 0x29, 0x5d, 0xf4, 0x25, 0x70, 0x0e, 0xc9, 0xcf, 0x32, 0x8b,
	0x20, 0x1b, 0x27, 0x59, 0x1c, 0x72, 0xa4, 0x53, 0xf1, 0x15, 0x43, 0xb9, 0x44, 0xaa, 0x8d, 0x6d,
	0x75, 0xa8, 0x8a, 0x62, 0x36, 0xa7, 0x47, 0x7e, 0xef, 0x52, 0x63, 0x36, 0x63, 0x31, 0xc5, 0x04,
	0x32, 0x15, 0xb4, 0x65, 0xf0, 0xd7, 0x56, 0x38, 0xd9, 0xd6, 0x45, 0xb6, 0x57, 0x79, 0x7c, 0x6a,
	0x1a, 0x47, 0xe7, 0xcf, 0x6b, 0xcf, 0x5c, 0xad, 0x3d, 0xf3, 0x75, 0xed, 0x99, 0x0f, 0x1b, 0xcf,
	0x58, 0x6d, 0x3c, 0xe3, 0x65, 0xe3, 0x19, 0xb7, 0xfb, 0x71, 0x82, 0x93, 0x32, 0xf2, 0x47, 0x90,
	0x06, 0x17, 0x37, 0xd7, 0xa7, 0x97, 0x6a, 0x97, 0x60, 0x34, 0xa1, 0x49, 0x16, 0x2c, 0x76, 0x87,
	0xc7, 0x65, 0xce, 0x78, 0x54, 0x95, 0x37, 0x3c, 0x78, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xce, 0x28,
	0x32, 0xe5, 0x16, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnbondingDelegationTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnbondingDelegationTime))
		i--
		dAtA[i] = 0x58
	}
	if m.UnbondingStakingTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnbondingStakingTime))
		i--
		dAtA[i] = 0x50
	}
	if m.MaxPoints != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxPoints))
		i--
		dAtA[i] = 0x48
	}
	if len(m.NetworkFee) > 0 {
		i -= len(m.NetworkFee)
		copy(dAtA[i:], m.NetworkFee)
		i = encodeVarintParams(dAtA, i, uint64(len(m.NetworkFee)))
		i--
		dAtA[i] = 0x42
	}
	if m.StorageCost != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.StorageCost))
		i--
		dAtA[i] = 0x38
	}
	if m.UploadTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UploadTimeout))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TimeoutSlash) > 0 {
		i -= len(m.TimeoutSlash)
		copy(dAtA[i:], m.TimeoutSlash)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TimeoutSlash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.UploadSlash) > 0 {
		i -= len(m.UploadSlash)
		copy(dAtA[i:], m.UploadSlash)
		i = encodeVarintParams(dAtA, i, uint64(len(m.UploadSlash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VoteSlash) > 0 {
		i -= len(m.VoteSlash)
		copy(dAtA[i:], m.VoteSlash)
		i = encodeVarintParams(dAtA, i, uint64(len(m.VoteSlash)))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VoteSlash)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.UploadSlash)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TimeoutSlash)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.UploadTimeout != 0 {
		n += 1 + sovParams(uint64(m.UploadTimeout))
	}
	if m.StorageCost != 0 {
		n += 1 + sovParams(uint64(m.StorageCost))
	}
	l = len(m.NetworkFee)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxPoints != 0 {
		n += 1 + sovParams(uint64(m.MaxPoints))
	}
	if m.UnbondingStakingTime != 0 {
		n += 1 + sovParams(uint64(m.UnbondingStakingTime))
	}
	if m.UnbondingDelegationTime != 0 {
		n += 1 + sovParams(uint64(m.UnbondingDelegationTime))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteSlash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoteSlash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadSlash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UploadSlash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutSlash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeoutSlash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UploadTimeout", wireType)
			}
			m.UploadTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UploadTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageCost", wireType)
			}
			m.StorageCost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StorageCost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetworkFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPoints", wireType)
			}
			m.MaxPoints = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPoints |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingStakingTime", wireType)
			}
			m.UnbondingStakingTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingStakingTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDelegationTime", wireType)
			}
			m.UnbondingDelegationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingDelegationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
