// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rook/matchmaker/genesis.proto

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

// GenesisState defines the matchmaker module's genesis state.
type GenesisState struct {
	Params       Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	InitialModes []Mode `protobuf:"bytes,2,rep,name=initial_modes,json=initialModes,proto3" json:"initial_modes"`
	NextModeId   uint32 `protobuf:"varint,3,opt,name=next_mode_id,json=nextModeId,proto3" json:"next_mode_id,omitempty"`
	NextRoomId   uint64 `protobuf:"varint,4,opt,name=next_room_id,json=nextRoomId,proto3" json:"next_room_id,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_aba54f3fbb842714, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetInitialModes() []Mode {
	if m != nil {
		return m.InitialModes
	}
	return nil
}

func (m *GenesisState) GetNextModeId() uint32 {
	if m != nil {
		return m.NextModeId
	}
	return 0
}

func (m *GenesisState) GetNextRoomId() uint64 {
	if m != nil {
		return m.NextRoomId
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "rook.matchmaker.GenesisState")
}

func init() { proto.RegisterFile("rook/matchmaker/genesis.proto", fileDescriptor_aba54f3fbb842714) }

var fileDescriptor_aba54f3fbb842714 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x1b, 0x37, 0x76, 0xc8, 0x3a, 0x84, 0xa2, 0x58, 0x06, 0xc6, 0xe0, 0xa9, 0x17, 0x5b,
	0x9c, 0x78, 0x97, 0x5d, 0x64, 0x88, 0x20, 0xf5, 0xe6, 0x65, 0x64, 0x6d, 0xe8, 0xc2, 0x4c, 0x5f,
	0x49, 0x22, 0x6c, 0xdf, 0xc2, 0x8f, 0x35, 0xf0, 0xb2, 0xa3, 0x27, 0x91, 0xf6, 0x8b, 0x48, 0xd2,
	0xe2, 0x64, 0xde, 0x1e, 0xef, 0xff, 0xfb, 0xbd, 0x07, 0x7f, 0x7c, 0xae, 0x00, 0x56, 0x89, 0x64,
	0x26, 0x5b, 0x4a, 0xb6, 0xe2, 0x2a, 0x29, 0x78, 0xc9, 0xb5, 0xd0, 0x71, 0xa5, 0xc0, 0x40, 0x70,
	0x6c, 0xe3, 0x78, 0x1f, 0x8f, 0x4f, 0x0a, 0x28, 0xc0, 0x65, 0x89, 0x9d, 0x5a, 0x6c, 0x4c, 0x0f,
	0xaf, 0xec, 0xc7, 0x96, 0xb8, 0xfc, 0x40, 0xd8, 0xbf, 0x6f, 0x4f, 0x3f, 0x1b, 0x66, 0x78, 0x70,
	0x8b, 0x07, 0x15, 0x53, 0x4c, 0xea, 0x10, 0x51, 0x14, 0x0d, 0x27, 0x67, 0xf1, 0xc1, 0xab, 0xf8,
	0xc9, 0xc5, 0xd3, 0xfe, 0xf6, 0xeb, 0xc2, 0x4b, 0x3b, 0x38, 0xb8, 0xc3, 0x23, 0x51, 0x0a, 0x23,
	0xd8, 0xeb, 0x5c, 0x42, 0xce, 0x75, 0x78, 0x44, 0x7b, 0xd1, 0x70, 0x72, 0xfa, 0xcf, 0x7e, 0x84,
	0x9c, 0x77, 0xae, 0xdf, 0x19, 0x76, 0xa5, 0x03, 0x8a, 0xfd, 0x92, 0xaf, 0x8d, 0xd3, 0xe7, 0x22,
	0x0f, 0x7b, 0x14, 0x45, 0xa3, 0x14, 0xdb, 0x9d, 0x05, 0x66, 0xf9, 0x2f, 0xa1, 0x00, 0xa4, 0x25,
	0xfa, 0x14, 0x45, 0xfd, 0x96, 0x48, 0x01, 0xe4, 0x2c, 0x9f, 0x3e, 0x6c, 0x6b, 0x82, 0x76, 0x35,
	0x41, 0xdf, 0x35, 0x41, 0xef, 0x0d, 0xf1, 0x76, 0x0d, 0xf1, 0x3e, 0x1b, 0xe2, 0xbd, 0x5c, 0x17,
	0xc2, 0x2c, 0xdf, 0x16, 0x71, 0x06, 0x32, 0x61, 0x2a, 0x63, 0x25, 0xbf, 0xd2, 0x1b, 0x6d, 0xb8,
	0xd4, 0x89, 0xeb, 0x68, 0xfd, 0xb7, 0x25, 0xb3, 0xa9, 0xb8, 0x5e, 0x0c, 0x5c, 0x43, 0x37, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xb5, 0x19, 0x64, 0x8b, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextRoomId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextRoomId))
		i--
		dAtA[i] = 0x20
	}
	if m.NextModeId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextModeId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.InitialModes) > 0 {
		for iNdEx := len(m.InitialModes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialModes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.InitialModes) > 0 {
		for _, e := range m.InitialModes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextModeId != 0 {
		n += 1 + sovGenesis(uint64(m.NextModeId))
	}
	if m.NextRoomId != 0 {
		n += 1 + sovGenesis(uint64(m.NextRoomId))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialModes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialModes = append(m.InitialModes, Mode{})
			if err := m.InitialModes[len(m.InitialModes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextModeId", wireType)
			}
			m.NextModeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextModeId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRoomId", wireType)
			}
			m.NextRoomId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextRoomId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
