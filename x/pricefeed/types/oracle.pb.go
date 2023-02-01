// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consumer/pricefeed/oracle.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_x_gov_types_v1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
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

type RequestInterval struct {
	Symbol         string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	OracleScriptId uint64 `protobuf:"varint,2,opt,name=oracle_script_id,json=oracleScriptId,proto3" json:"oracle_script_id,omitempty"`
	BlockInterval  uint64 `protobuf:"varint,3,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

func (m *RequestInterval) Reset()         { *m = RequestInterval{} }
func (m *RequestInterval) String() string { return proto.CompactTextString(m) }
func (*RequestInterval) ProtoMessage()    {}
func (*RequestInterval) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe39f7b31fac89b, []int{0}
}
func (m *RequestInterval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestInterval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestInterval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestInterval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInterval.Merge(m, src)
}
func (m *RequestInterval) XXX_Size() int {
	return m.Size()
}
func (m *RequestInterval) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInterval.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInterval proto.InternalMessageInfo

func (m *RequestInterval) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *RequestInterval) GetOracleScriptId() uint64 {
	if m != nil {
		return m.OracleScriptId
	}
	return 0
}

func (m *RequestInterval) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

type PriceFeed struct {
	Symbol      string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price       uint64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	ResolveTime int64  `protobuf:"varint,3,opt,name=resolve_time,json=resolveTime,proto3" json:"resolve_time,omitempty"`
}

func (m *PriceFeed) Reset()         { *m = PriceFeed{} }
func (m *PriceFeed) String() string { return proto.CompactTextString(m) }
func (*PriceFeed) ProtoMessage()    {}
func (*PriceFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe39f7b31fac89b, []int{1}
}
func (m *PriceFeed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PriceFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PriceFeed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PriceFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceFeed.Merge(m, src)
}
func (m *PriceFeed) XXX_Size() int {
	return m.Size()
}
func (m *PriceFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceFeed.DiscardUnknown(m)
}

var xxx_messageInfo_PriceFeed proto.InternalMessageInfo

func (m *PriceFeed) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *PriceFeed) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *PriceFeed) GetResolveTime() int64 {
	if m != nil {
		return m.ResolveTime
	}
	return 0
}

type UpdateSymbolRequestProposal struct {
	// Types that are valid to be assigned to SymbolRequestProposal:
	//	*UpdateSymbolRequestProposal_Title
	//	*UpdateSymbolRequestProposal_Description
	//	*UpdateSymbolRequestProposal_Symbols
	SymbolRequestProposal isUpdateSymbolRequestProposal_SymbolRequestProposal `protobuf_oneof:"SymbolRequestProposal"`
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

type isUpdateSymbolRequestProposal_SymbolRequestProposal interface {
	isUpdateSymbolRequestProposal_SymbolRequestProposal()
	MarshalTo([]byte) (int, error)
	Size() int
}

type UpdateSymbolRequestProposal_Title struct {
	Title string `protobuf:"bytes,1,opt,name=title,proto3,oneof" json:"title,omitempty"`
}
type UpdateSymbolRequestProposal_Description struct {
	Description string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
}
type UpdateSymbolRequestProposal_Symbols struct {
	Symbols *SymbolRequests `protobuf:"bytes,3,opt,name=Symbols,proto3,oneof" json:"Symbols,omitempty"`
}

func (*UpdateSymbolRequestProposal_Title) isUpdateSymbolRequestProposal_SymbolRequestProposal() {}
func (*UpdateSymbolRequestProposal_Description) isUpdateSymbolRequestProposal_SymbolRequestProposal() {
}
func (*UpdateSymbolRequestProposal_Symbols) isUpdateSymbolRequestProposal_SymbolRequestProposal() {}

func (m *UpdateSymbolRequestProposal) GetSymbolRequestProposal() isUpdateSymbolRequestProposal_SymbolRequestProposal {
	if m != nil {
		return m.SymbolRequestProposal
	}
	return nil
}

func (m *UpdateSymbolRequestProposal) GetTitle() string {
	if x, ok := m.GetSymbolRequestProposal().(*UpdateSymbolRequestProposal_Title); ok {
		return x.Title
	}
	return ""
}

func (m *UpdateSymbolRequestProposal) GetDescription() string {
	if x, ok := m.GetSymbolRequestProposal().(*UpdateSymbolRequestProposal_Description); ok {
		return x.Description
	}
	return ""
}

func (m *UpdateSymbolRequestProposal) GetSymbols() *SymbolRequests {
	if x, ok := m.GetSymbolRequestProposal().(*UpdateSymbolRequestProposal_Symbols); ok {
		return x.Symbols
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UpdateSymbolRequestProposal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UpdateSymbolRequestProposal_Title)(nil),
		(*UpdateSymbolRequestProposal_Description)(nil),
		(*UpdateSymbolRequestProposal_Symbols)(nil),
	}
}

type SymbolRequests struct {
	Symbols []*SymbolRequest `protobuf:"bytes,1,rep,name=Symbols,proto3" json:"Symbols,omitempty"`
}

func (m *SymbolRequests) Reset()         { *m = SymbolRequests{} }
func (m *SymbolRequests) String() string { return proto.CompactTextString(m) }
func (*SymbolRequests) ProtoMessage()    {}
func (*SymbolRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_afe39f7b31fac89b, []int{3}
}
func (m *SymbolRequests) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SymbolRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SymbolRequests.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SymbolRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymbolRequests.Merge(m, src)
}
func (m *SymbolRequests) XXX_Size() int {
	return m.Size()
}
func (m *SymbolRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_SymbolRequests.DiscardUnknown(m)
}

var xxx_messageInfo_SymbolRequests proto.InternalMessageInfo

func (m *SymbolRequests) GetSymbols() []*SymbolRequest {
	if m != nil {
		return m.Symbols
	}
	return nil
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
	return fileDescriptor_afe39f7b31fac89b, []int{4}
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

func init() {
	proto.RegisterType((*RequestInterval)(nil), "consumer.pricefeed.RequestInterval")
	proto.RegisterType((*PriceFeed)(nil), "consumer.pricefeed.PriceFeed")
	proto.RegisterType((*UpdateSymbolRequestProposal)(nil), "consumer.pricefeed.UpdateSymbolRequestProposal")
	proto.RegisterType((*SymbolRequests)(nil), "consumer.pricefeed.SymbolRequests")
	proto.RegisterType((*SymbolRequest)(nil), "consumer.pricefeed.SymbolRequest")
}

func init() { proto.RegisterFile("consumer/pricefeed/oracle.proto", fileDescriptor_afe39f7b31fac89b) }

var fileDescriptor_afe39f7b31fac89b = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0xd3, 0x8e, 0xbb, 0xb2, 0x1d, 0x77, 0x94, 0x66, 0xd5, 0x71, 0x84, 0xec, 0x6c, 0x40,
	0xc8, 0x65, 0x13, 0x76, 0xf5, 0x20, 0x0a, 0x22, 0x2b, 0xc8, 0xee, 0x41, 0x58, 0xb2, 0x7a, 0x11,
	0x21, 0xe4, 0x4f, 0x19, 0x9b, 0x49, 0x52, 0xb1, 0xbb, 0x27, 0xcc, 0xf8, 0x04, 0x1e, 0x7d, 0x04,
	0x1f, 0xc2, 0x87, 0x10, 0x4f, 0x73, 0xf4, 0x28, 0x33, 0xaf, 0xe0, 0x03, 0xc8, 0x74, 0x67, 0xc6,
	0x19, 0x54, 0xbc, 0x79, 0x4a, 0xaa, 0xbe, 0x4a, 0xfd, 0xea, 0x4b, 0x77, 0xd1, 0xfd, 0x14, 0x2b,
	0x39, 0x2a, 0x41, 0x04, 0xb5, 0xe0, 0x29, 0xbc, 0x01, 0xc8, 0x02, 0x14, 0x71, 0x5a, 0x80, 0x5f,
	0x0b, 0x54, 0xc8, 0xd8, 0xb2, 0xc0, 0x5f, 0x15, 0xf4, 0xf7, 0x72, 0xcc, 0x51, 0xcb, 0xc1, 0xe2,
	0xcd, 0x54, 0xf6, 0x6f, 0xa7, 0x28, 0x4b, 0x94, 0x91, 0x11, 0x4c, 0x60, 0x24, 0xf7, 0x3d, 0xbd,
	0x16, 0xc2, 0xbb, 0x11, 0x48, 0x75, 0x56, 0x29, 0x10, 0x4d, 0x5c, 0xb0, 0x9b, 0x74, 0x5b, 0x4e,
	0xca, 0x04, 0x8b, 0x1e, 0x19, 0x10, 0x6f, 0x27, 0x6c, 0x23, 0xe6, 0xd1, 0xeb, 0x86, 0x1f, 0xc9,
	0x54, 0xf0, 0x5a, 0x45, 0x3c, 0xeb, 0x5d, 0x1a, 0x10, 0xef, 0x72, 0xd8, 0x35, 0xf9, 0x0b, 0x9d,
	0x3e, 0xcb, 0xd8, 0x5d, 0xda, 0x4d, 0x0a, 0x4c, 0x87, 0x11, 0x6f, 0x7b, 0xf6, 0x3a, 0xba, 0x6e,
	0x57, 0x67, 0x97, 0x20, 0xf7, 0x35, 0xdd, 0x39, 0x5f, 0x4c, 0xfe, 0x0c, 0x20, 0xfb, 0x2b, 0x75,
	0x8f, 0x6e, 0x69, 0x7b, 0x2d, 0xca, 0x04, 0xec, 0x80, 0x5e, 0x15, 0x20, 0xb1, 0x68, 0x20, 0x52,
	0xbc, 0x04, 0xdd, 0xbf, 0x13, 0xda, 0x6d, 0xee, 0x05, 0x2f, 0xc1, 0xfd, 0x41, 0xe8, 0x9d, 0x97,
	0x75, 0x16, 0x2b, 0xb8, 0xd0, 0x9d, 0x5a, 0x9b, 0xe7, 0x02, 0x6b, 0x94, 0xda, 0xe6, 0x96, 0xe2,
	0xaa, 0x00, 0xc3, 0x3b, 0xb5, 0x42, 0x13, 0x32, 0x97, 0xda, 0x19, 0x18, 0x87, 0x1c, 0x2b, 0x8d,
	0x5d, 0xa8, 0xeb, 0x49, 0xf6, 0x98, 0x5e, 0x31, 0x4d, 0xa5, 0x26, 0xdb, 0xc7, 0xae, 0xff, 0xfb,
	0x61, 0xf8, 0x1b, 0x5c, 0x79, 0x6a, 0x85, 0xcb, 0x8f, 0x1e, 0x3e, 0xf9, 0xf0, 0x69, 0xdf, 0xfa,
	0xfa, 0xf9, 0xf0, 0x41, 0xce, 0xd5, 0xdb, 0x51, 0xe2, 0xa7, 0x58, 0xb6, 0xc7, 0xd2, 0x3e, 0x0e,
	0x65, 0x36, 0x0c, 0xc6, 0x41, 0x8e, 0x4d, 0xa0, 0x26, 0x35, 0xc8, 0xa0, 0x39, 0x4a, 0x40, 0xc5,
	0x47, 0xfe, 0x53, 0xac, 0x14, 0x54, 0xea, 0xe4, 0x16, 0xbd, 0xf1, 0x47, 0x5b, 0xee, 0x73, 0xda,
	0xdd, 0xe4, 0xb2, 0x47, 0xbf, 0x86, 0x25, 0x83, 0x8e, 0x67, 0x1f, 0x1f, 0xfc, 0x73, 0xd8, 0xd5,
	0xa4, 0xee, 0x98, 0xee, 0x6e, 0x28, 0xff, 0xed, 0x76, 0x9c, 0xdc, 0xff, 0x32, 0x73, 0xc8, 0x74,
	0xe6, 0x90, 0xef, 0x33, 0x87, 0x7c, 0x9c, 0x3b, 0xd6, 0x74, 0xee, 0x58, 0xdf, 0xe6, 0x8e, 0xf5,
	0xaa, 0xbf, 0xda, 0x8c, 0xf1, 0xda, 0x6e, 0xe8, 0xff, 0x94, 0x6c, 0xeb, 0x6b, 0x7d, 0xef, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xb8, 0xd3, 0xbd, 0x3e, 0x03, 0x00, 0x00,
}

func (this *UpdateSymbolRequestProposal) GetContent() github_com_cosmos_cosmos_sdk_x_gov_types_v1beta1.Content {
	if x := this.GetTitle(); x != nil {
		return x
	}
	if x := this.GetDescription(); x != nil {
		return x
	}
	if x := this.GetSymbols(); x != nil {
		return x
	}
	return nil
}

func (this *UpdateSymbolRequestProposal) SetContent(value github_com_cosmos_cosmos_sdk_x_gov_types_v1beta1.Content) error {
	if value == nil {
		this.SymbolRequestProposal = nil
		return nil
	}
	switch vt := value.(type) {
	case string:
		this.SymbolRequestProposal = &UpdateSymbolRequestProposal_Title{vt}
		return nil
	case string:
		this.SymbolRequestProposal = &UpdateSymbolRequestProposal_Description{vt}
		return nil
	case *SymbolRequests:
		this.SymbolRequestProposal = &UpdateSymbolRequestProposal_Symbols{vt}
		return nil
	case SymbolRequests:
		this.SymbolRequestProposal = &UpdateSymbolRequestProposal_Symbols{&vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message UpdateSymbolRequestProposal", value)
}

func (m *RequestInterval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestInterval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestInterval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *PriceFeed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PriceFeed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PriceFeed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if m.SymbolRequestProposal != nil {
		{
			size := m.SymbolRequestProposal.Size()
			i -= size
			if _, err := m.SymbolRequestProposal.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *UpdateSymbolRequestProposal_Title) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSymbolRequestProposal_Title) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Title)
	copy(dAtA[i:], m.Title)
	i = encodeVarintOracle(dAtA, i, uint64(len(m.Title)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *UpdateSymbolRequestProposal_Description) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSymbolRequestProposal_Description) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Description)
	copy(dAtA[i:], m.Description)
	i = encodeVarintOracle(dAtA, i, uint64(len(m.Description)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *UpdateSymbolRequestProposal_Symbols) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateSymbolRequestProposal_Symbols) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Symbols != nil {
		{
			size, err := m.Symbols.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOracle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *SymbolRequests) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SymbolRequests) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SymbolRequests) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Symbols[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *RequestInterval) Size() (n int) {
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

func (m *PriceFeed) Size() (n int) {
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

func (m *UpdateSymbolRequestProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SymbolRequestProposal != nil {
		n += m.SymbolRequestProposal.Size()
	}
	return n
}

func (m *UpdateSymbolRequestProposal_Title) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	n += 1 + l + sovOracle(uint64(l))
	return n
}
func (m *UpdateSymbolRequestProposal_Description) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	n += 1 + l + sovOracle(uint64(l))
	return n
}
func (m *UpdateSymbolRequestProposal_Symbols) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Symbols != nil {
		l = m.Symbols.Size()
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}
func (m *SymbolRequests) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, e := range m.Symbols {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
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

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestInterval) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RequestInterval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestInterval: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *PriceFeed) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PriceFeed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceFeed: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.SymbolRequestProposal = &UpdateSymbolRequestProposal_Title{string(dAtA[iNdEx:postIndex])}
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
			m.SymbolRequestProposal = &UpdateSymbolRequestProposal_Description{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
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
			v := &SymbolRequests{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.SymbolRequestProposal = &UpdateSymbolRequestProposal_Symbols{v}
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
func (m *SymbolRequests) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SymbolRequests: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SymbolRequests: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
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
			m.Symbols = append(m.Symbols, &SymbolRequest{})
			if err := m.Symbols[len(m.Symbols)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
