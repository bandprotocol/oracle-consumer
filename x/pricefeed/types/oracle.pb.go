// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consumer/pricefeed/oracle.proto

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

type Price struct {
	Symbol      string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price       uint64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	ResolveTime int64  `protobuf:"varint,3,opt,name=resolve_time,json=resolveTime,proto3" json:"resolve_time,omitempty"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe39f7b31fac89b, []int{0}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Price.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return m.Size()
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Price) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Price) GetResolveTime() int64 {
	if m != nil {
		return m.ResolveTime
	}
	return 0
}

type SymbolRequest struct {
	Symbol         string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	OracleScriptId uint64 `protobuf:"varint,2,opt,name=oracle_script_id,json=oracleScriptId,proto3" json:"oracle_script_id,omitempty"`
	BlockInterval  uint64 `protobuf:"varint,3,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

func (m *SymbolRequest) Reset()         { *m = SymbolRequest{} }
func (m *SymbolRequest) String() string { return proto.CompactTextString(m) }
func (*SymbolRequest) ProtoMessage()    {}
func (*SymbolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe39f7b31fac89b, []int{1}
}
func (m *SymbolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SymbolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SymbolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SymbolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymbolRequest.Merge(m, src)
}
func (m *SymbolRequest) XXX_Size() int {
	return m.Size()
}
func (m *SymbolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SymbolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SymbolRequest proto.InternalMessageInfo

func (m *SymbolRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SymbolRequest) GetOracleScriptId() uint64 {
	if m != nil {
		return m.OracleScriptId
	}
	return 0
}

func (m *SymbolRequest) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

type UpdateSymbolRequestProposal struct {
	Title          string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description    string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	SymbolRequests []SymbolRequest `protobuf:"bytes,3,rep,name=symbol_requests,json=symbolRequests,proto3" json:"symbol_requests"`
}

func (m *UpdateSymbolRequestProposal) Reset()         { *m = UpdateSymbolRequestProposal{} }
func (m *UpdateSymbolRequestProposal) String() string { return proto.CompactTextString(m) }
func (*UpdateSymbolRequestProposal) ProtoMessage()    {}
func (*UpdateSymbolRequestProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe39f7b31fac89b, []int{2}
}
func (m *UpdateSymbolRequestProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateSymbolRequestProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateSymbolRequestProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateSymbolRequestProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSymbolRequestProposal.Merge(m, src)
}
func (m *UpdateSymbolRequestProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateSymbolRequestProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSymbolRequestProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSymbolRequestProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Price)(nil), "consumer.pricefeed.Price")
	proto.RegisterType((*SymbolRequest)(nil), "consumer.pricefeed.SymbolRequest")
	proto.RegisterType((*UpdateSymbolRequestProposal)(nil), "consumer.pricefeed.UpdateSymbolRequestProposal")
}

func init() { proto.RegisterFile("consumer/pricefeed/oracle.proto", fileDescriptor_afe39f7b31fac89b) }

var fileDescriptor_afe39f7b31fac89b = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x6a, 0xea, 0x40,
	0x18, 0x85, 0x33, 0x37, 0x51, 0x70, 0xbc, 0x7a, 0x2f, 0x83, 0x94, 0x60, 0x21, 0x46, 0xa1, 0x90,
	0x55, 0x84, 0x76, 0x53, 0xba, 0x74, 0xe7, 0x4e, 0x62, 0x0b, 0xa5, 0x9b, 0x10, 0x93, 0xbf, 0x32,
	0x34, 0x71, 0xd2, 0x99, 0x51, 0xf4, 0x0d, 0xba, 0xec, 0x23, 0x74, 0xd9, 0x4d, 0xdf, 0xc3, 0xa5,
	0xcb, 0xae, 0x4a, 0x89, 0x2f, 0x52, 0x9c, 0x49, 0x45, 0x29, 0xdd, 0xcd, 0xff, 0xcd, 0xe1, 0x3f,
	0x67, 0xe6, 0xe0, 0x4e, 0xcc, 0x66, 0x62, 0x9e, 0x01, 0xef, 0xe7, 0x9c, 0xc6, 0x70, 0x0f, 0x90,
	0xf4, 0x19, 0x8f, 0xe2, 0x14, 0xfc, 0x9c, 0x33, 0xc9, 0x08, 0xf9, 0x16, 0xf8, 0x7b, 0x41, 0xbb,
	0x35, 0x65, 0x53, 0xa6, 0xae, 0xfb, 0xbb, 0x93, 0x56, 0xf6, 0x6e, 0x71, 0x65, 0xb4, 0x93, 0x90,
	0x13, 0x5c, 0x15, 0xab, 0x6c, 0xc2, 0x52, 0x1b, 0xb9, 0xc8, 0xab, 0x05, 0xe5, 0x44, 0x5a, 0xb8,
	0xa2, 0x76, 0xd8, 0x7f, 0x5c, 0xe4, 0x59, 0x81, 0x1e, 0x48, 0x17, 0xff, 0xe5, 0x20, 0x58, 0xba,
	0x80, 0x50, 0xd2, 0x0c, 0x6c, 0xd3, 0x45, 0x9e, 0x19, 0xd4, 0x4b, 0x76, 0x4d, 0x33, 0xe8, 0x2d,
	0x71, 0x63, 0xac, 0x56, 0x04, 0xf0, 0x38, 0x07, 0x21, 0x7f, 0x75, 0xf0, 0xf0, 0x7f, 0x1d, 0x3e,
	0x14, 0x31, 0xa7, 0xb9, 0x0c, 0x69, 0x52, 0x9a, 0x35, 0x35, 0x1f, 0x2b, 0x3c, 0x4c, 0xc8, 0x19,
	0x6e, 0x4e, 0x52, 0x16, 0x3f, 0x84, 0x74, 0x26, 0x81, 0x2f, 0xa2, 0x54, 0xf9, 0x5a, 0x41, 0x43,
	0xd1, 0x61, 0x09, 0x7b, 0x6f, 0x08, 0x9f, 0xde, 0xe4, 0x49, 0x24, 0xe1, 0x28, 0xc0, 0x88, 0xb3,
	0x9c, 0x89, 0x48, 0x3d, 0x49, 0x52, 0x99, 0x42, 0x99, 0x43, 0x0f, 0xc4, 0xc5, 0xf5, 0x04, 0x74,
	0x02, 0xca, 0x66, 0x2a, 0x41, 0x2d, 0x38, 0x44, 0x64, 0x84, 0xff, 0xe9, 0xc8, 0x21, 0xd7, 0x1b,
	0x85, 0x6d, 0xba, 0xa6, 0x57, 0x3f, 0xef, 0xfa, 0x3f, 0xff, 0xdb, 0x3f, 0xf2, 0x1e, 0x58, 0xeb,
	0x8f, 0x8e, 0x11, 0x34, 0xc5, 0x21, 0x14, 0x57, 0xd6, 0xd3, 0x4b, 0xc7, 0x18, 0x5c, 0xbe, 0x16,
	0x0e, 0x5a, 0x17, 0x0e, 0xda, 0x14, 0x0e, 0xfa, 0x2c, 0x1c, 0xf4, 0xbc, 0x75, 0x8c, 0xcd, 0xd6,
	0x31, 0xde, 0xb7, 0x8e, 0x71, 0xd7, 0xde, 0x97, 0xbd, 0x3c, 0xa8, 0x5b, 0xae, 0x72, 0x10, 0x93,
	0xaa, 0x2a, 0xf1, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x23, 0x51, 0x13, 0x11, 0x02, 0x00,
	0x00,
}

func (this *Price) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Price)
	if !ok {
		that2, ok := that.(Price)
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
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Price != that1.Price {
		return false
	}
	if this.ResolveTime != that1.ResolveTime {
		return false
	}
	return true
}
func (this *SymbolRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SymbolRequest)
	if !ok {
		that2, ok := that.(SymbolRequest)
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
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.OracleScriptId != that1.OracleScriptId {
		return false
	}
	if this.BlockInterval != that1.BlockInterval {
		return false
	}
	return true
}
func (this *UpdateSymbolRequestProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateSymbolRequestProposal)
	if !ok {
		that2, ok := that.(UpdateSymbolRequestProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.SymbolRequests) != len(that1.SymbolRequests) {
		return false
	}
	for i := range this.SymbolRequests {
		if !this.SymbolRequests[i].Equal(&that1.SymbolRequests[i]) {
			return false
		}
	}
	return true
}
func (m *Price) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Price) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Price) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResolveTime != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ResolveTime))
		i--
		dAtA[i] = 0x18
	}
	if m.Price != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SymbolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SymbolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SymbolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockInterval != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.BlockInterval))
		i--
		dAtA[i] = 0x18
	}
	if m.OracleScriptId != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.OracleScriptId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateSymbolRequestProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateSymbolRequestProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSymbolRequestProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SymbolRequests) > 0 {
		for iNdEx := len(m.SymbolRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SymbolRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Price) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovOracle(uint64(m.Price))
	}
	if m.ResolveTime != 0 {
		n += 1 + sovOracle(uint64(m.ResolveTime))
	}
	return n
}

func (m *SymbolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.OracleScriptId != 0 {
		n += 1 + sovOracle(uint64(m.OracleScriptId))
	}
	if m.BlockInterval != 0 {
		n += 1 + sovOracle(uint64(m.BlockInterval))
	}
	return n
}

func (m *UpdateSymbolRequestProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if len(m.SymbolRequests) > 0 {
		for _, e := range m.SymbolRequests {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Price) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Price: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Price: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveTime", wireType)
			}
			m.ResolveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolveTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *SymbolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: SymbolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SymbolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleScriptId", wireType)
			}
			m.OracleScriptId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OracleScriptId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInterval", wireType)
			}
			m.BlockInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *UpdateSymbolRequestProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: UpdateSymbolRequestProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateSymbolRequestProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SymbolRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SymbolRequests = append(m.SymbolRequests, SymbolRequest{})
			if err := m.SymbolRequests[len(m.SymbolRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
