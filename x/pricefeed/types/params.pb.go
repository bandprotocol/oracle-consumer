// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pricefeed/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	// AskCount is the number of validators that are requested to respond to this
	// oracle request. Higher value means more security, at a higher gas cost.
	AskCount uint64 `protobuf:"varint,1,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty"`
	// MinCount is the minimum number of validators necessary for the request to
	// proceed to the execution phase. Higher value means more security, at the
	// cost of liveness.
	MinCount uint64 `protobuf:"varint,2,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	// MinDsCount is the minimum number of responses required from data sources
	// (oracles) to consider the request as successful.
	MinDsCount uint64 `protobuf:"varint,3,opt,name=min_ds_count,json=minDsCount,proto3" json:"min_ds_count,omitempty"`
	// PrepareGasBase is the amount of base gas needed to prepare an oracle
	// request.
	PrepareGasBase uint64 `protobuf:"varint,4,opt,name=prepare_gas_base,json=prepareGasBase,proto3" json:"prepare_gas_base,omitempty"`
	// PrepareGasEach is the amount of additional gas needed per symbol in an
	// oracle request.
	PrepareGasEach uint64 `protobuf:"varint,5,opt,name=prepare_gas_each,json=prepareGasEach,proto3" json:"prepare_gas_each,omitempty"`
	// ExecuteGasBase is the amount of base gas needed to execute an oracle
	// request.
	ExecuteGasBase uint64 `protobuf:"varint,6,opt,name=execute_gas_base,json=executeGasBase,proto3" json:"execute_gas_base,omitempty"`
	// ExecuteGasEach is the amount of additional gas needed per symbol in
	// executing an oracle request.
	ExecuteGasEach uint64 `protobuf:"varint,7,opt,name=execute_gas_each,json=executeGasEach,proto3" json:"execute_gas_each,omitempty"`
	// SourceChannel is the source channel to use when sending the oracle request
	// to BandChain via IBC
	SourceChannel string `protobuf:"bytes,8,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	// FeeLimit is the maximum tokens that will be paid to all data source
	// providers.
	FeeLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,9,rep,name=fee_limit,json=feeLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_limit"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_76ce94fb65bc3a91, []int{0}
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

func (m *Params) GetAskCount() uint64 {
	if m != nil {
		return m.AskCount
	}
	return 0
}

func (m *Params) GetMinCount() uint64 {
	if m != nil {
		return m.MinCount
	}
	return 0
}

func (m *Params) GetMinDsCount() uint64 {
	if m != nil {
		return m.MinDsCount
	}
	return 0
}

func (m *Params) GetPrepareGasBase() uint64 {
	if m != nil {
		return m.PrepareGasBase
	}
	return 0
}

func (m *Params) GetPrepareGasEach() uint64 {
	if m != nil {
		return m.PrepareGasEach
	}
	return 0
}

func (m *Params) GetExecuteGasBase() uint64 {
	if m != nil {
		return m.ExecuteGasBase
	}
	return 0
}

func (m *Params) GetExecuteGasEach() uint64 {
	if m != nil {
		return m.ExecuteGasEach
	}
	return 0
}

func (m *Params) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *Params) GetFeeLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "pricefeed.Params")
}

func init() { proto.RegisterFile("pricefeed/params.proto", fileDescriptor_76ce94fb65bc3a91) }

var fileDescriptor_76ce94fb65bc3a91 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0x13, 0x5a, 0x4a, 0x63, 0xe0, 0x84, 0x2a, 0x84, 0xc2, 0x21, 0xa5, 0x15, 0x12, 0x52,
	0x96, 0x8b, 0x39, 0x98, 0x60, 0x6c, 0x41, 0x2c, 0x0c, 0xa8, 0x6c, 0x2c, 0x91, 0xe3, 0xbe, 0x26,
	0x56, 0x63, 0x3b, 0xf2, 0x4b, 0xd0, 0xf1, 0x2d, 0x18, 0x19, 0x91, 0xd8, 0xf8, 0x24, 0x37, 0x76,
	0x64, 0x02, 0xd4, 0x7e, 0x11, 0x64, 0x3b, 0xe2, 0xa2, 0x4e, 0x49, 0xfe, 0xbf, 0x5f, 0xfe, 0xb1,
	0xde, 0x0b, 0x79, 0xd4, 0x18, 0xc1, 0x61, 0x0b, 0xb0, 0xa1, 0x0d, 0x33, 0x4c, 0x62, 0xd6, 0x18,
	0xdd, 0xea, 0x59, 0xf4, 0x3f, 0x3f, 0x7f, 0x58, 0xea, 0x52, 0xbb, 0x94, 0xda, 0x3b, 0x2f, 0x9c,
	0x27, 0x5c, 0xa3, 0xd4, 0x48, 0x0b, 0x86, 0x40, 0x3f, 0x5f, 0x16, 0xd0, 0xb2, 0x4b, 0xca, 0xb5,
	0x50, 0x9e, 0x3f, 0xfd, 0x31, 0x22, 0x93, 0x0f, 0xae, 0x71, 0xf6, 0x84, 0x44, 0x0c, 0x77, 0x39,
	0xd7, 0x9d, 0x6a, 0xe3, 0x70, 0x11, 0xa6, 0xe3, 0xf5, 0x94, 0xe1, 0x6e, 0x65, 0x9f, 0x2d, 0x94,
	0x42, 0xf5, 0xf0, 0x96, 0x87, 0x52, 0x28, 0x0f, 0x17, 0xe4, 0x9e, 0x85, 0x1b, 0xec, 0xf9, 0xc8,
	0x71, 0x22, 0x85, 0x7a, 0x83, 0xde, 0x48, 0xc9, 0x83, 0xc6, 0x40, 0xc3, 0x0c, 0xe4, 0x25, 0xc3,
	0xdc, 0x9e, 0x26, 0x1e, 0x3b, 0xeb, 0xac, 0xcf, 0xdf, 0x31, 0x5c, 0x32, 0x84, 0x53, 0x13, 0x18,
	0xaf, 0xe2, 0xdb, 0xa7, 0xe6, 0x5b, 0xc6, 0x2b, 0x6b, 0xc2, 0x15, 0xf0, 0xae, 0x1d, 0x74, 0x4e,
	0xbc, 0xd9, 0xe7, 0x83, 0xce, 0xa1, 0xe9, 0x3a, 0xef, 0x9c, 0x9a, 0xae, 0xf3, 0x19, 0x39, 0x43,
	0xdd, 0x19, 0x0e, 0x39, 0xaf, 0x98, 0x52, 0x50, 0xc7, 0xd3, 0x45, 0x98, 0x46, 0xeb, 0xfb, 0x3e,
	0x5d, 0xf9, 0x70, 0x56, 0x91, 0x68, 0x0b, 0x90, 0xd7, 0x42, 0x8a, 0x36, 0x8e, 0x16, 0xa3, 0xf4,
	0xee, 0x8b, 0xc7, 0x99, 0x9f, 0x74, 0x66, 0xcf, 0x91, 0xf5, 0x93, 0xce, 0x56, 0x5a, 0xa8, 0xe5,
	0xf3, 0xeb, 0xdf, 0xf3, 0xe0, 0xe7, 0x9f, 0x79, 0x5a, 0x8a, 0xb6, 0xea, 0x8a, 0x8c, 0x6b, 0x49,
	0xfb, 0xb5, 0xf8, 0xcb, 0x05, 0x6e, 0x76, 0xb4, 0xfd, 0xd2, 0x00, 0xba, 0x17, 0x70, 0x3d, 0xdd,
	0x02, 0xbc, 0xb7, 0xe5, 0xaf, 0xc7, 0xdf, 0xbe, 0xcf, 0x83, 0xe5, 0xc7, 0xeb, 0x43, 0x12, 0xee,
	0x0f, 0x49, 0xf8, 0xf7, 0x90, 0x84, 0x5f, 0x8f, 0x49, 0xb0, 0x3f, 0x26, 0xc1, 0xaf, 0x63, 0x12,
	0x7c, 0x7a, 0x35, 0xe8, 0x2c, 0x98, 0xda, 0xb8, 0xad, 0x72, 0x5d, 0x53, 0x6d, 0x18, 0xaf, 0xe1,
	0x82, 0x6b, 0x85, 0x9d, 0x04, 0x43, 0xaf, 0xe8, 0xcd, 0x2f, 0xe4, 0x3e, 0x55, 0x4c, 0x9c, 0xfb,
	0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x98, 0x23, 0xca, 0x5c, 0x02, 0x00, 0x00,
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
	if len(m.FeeLimit) > 0 {
		for iNdEx := len(m.FeeLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x42
	}
	if m.ExecuteGasEach != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ExecuteGasEach))
		i--
		dAtA[i] = 0x38
	}
	if m.ExecuteGasBase != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ExecuteGasBase))
		i--
		dAtA[i] = 0x30
	}
	if m.PrepareGasEach != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PrepareGasEach))
		i--
		dAtA[i] = 0x28
	}
	if m.PrepareGasBase != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PrepareGasBase))
		i--
		dAtA[i] = 0x20
	}
	if m.MinDsCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinDsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.MinCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinCount))
		i--
		dAtA[i] = 0x10
	}
	if m.AskCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AskCount))
		i--
		dAtA[i] = 0x8
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
	if m.AskCount != 0 {
		n += 1 + sovParams(uint64(m.AskCount))
	}
	if m.MinCount != 0 {
		n += 1 + sovParams(uint64(m.MinCount))
	}
	if m.MinDsCount != 0 {
		n += 1 + sovParams(uint64(m.MinDsCount))
	}
	if m.PrepareGasBase != 0 {
		n += 1 + sovParams(uint64(m.PrepareGasBase))
	}
	if m.PrepareGasEach != 0 {
		n += 1 + sovParams(uint64(m.PrepareGasEach))
	}
	if m.ExecuteGasBase != 0 {
		n += 1 + sovParams(uint64(m.ExecuteGasBase))
	}
	if m.ExecuteGasEach != 0 {
		n += 1 + sovParams(uint64(m.ExecuteGasEach))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.FeeLimit) > 0 {
		for _, e := range m.FeeLimit {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
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
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskCount", wireType)
			}
			m.AskCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AskCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCount", wireType)
			}
			m.MinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDsCount", wireType)
			}
			m.MinDsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGasBase", wireType)
			}
			m.PrepareGasBase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGasBase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGasEach", wireType)
			}
			m.PrepareGasEach = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGasEach |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGasBase", wireType)
			}
			m.ExecuteGasBase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGasBase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGasEach", wireType)
			}
			m.ExecuteGasEach = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGasEach |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
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
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeLimit = append(m.FeeLimit, types.Coin{})
			if err := m.FeeLimit[len(m.FeeLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
