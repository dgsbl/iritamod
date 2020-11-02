// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slashing/slashing.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUnjailValidator - struct for unjailing jailed validator
type MsgUnjailValidator struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Operator string `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgUnjailValidator) Reset()         { *m = MsgUnjailValidator{} }
func (m *MsgUnjailValidator) String() string { return proto.CompactTextString(m) }
func (*MsgUnjailValidator) ProtoMessage()    {}
func (*MsgUnjailValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_b24ff443e5dfee94, []int{0}
}
func (m *MsgUnjailValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnjailValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnjailValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnjailValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnjailValidator.Merge(m, src)
}
func (m *MsgUnjailValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnjailValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnjailValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnjailValidator proto.InternalMessageInfo

// ValidatorSigningInfo defines a validator's signing info for monitoring their liveness activity.
type ValidatorSigningInfo struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// height at which validator was first a candidate OR was unjailed
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty" yaml:"start_height"`
	// index offset into signed block bit array
	IndexOffset int64 `protobuf:"varint,3,opt,name=index_offset,json=indexOffset,proto3" json:"index_offset,omitempty" yaml:"index_offset"`
	// timestamp validator cannot be unjailed until
	JailedUntil time.Time `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until" yaml:"jailed_until"`
	// whether or not a validator has been tombstoned (killed out of validator set)
	Tombstoned bool `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	// missed blocks counter (to avoid scanning the array every time)
	MissedBlocksCounter int64 `protobuf:"varint,6,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty" yaml:"missed_blocks_counter"`
}

func (m *ValidatorSigningInfo) Reset()         { *m = ValidatorSigningInfo{} }
func (m *ValidatorSigningInfo) String() string { return proto.CompactTextString(m) }
func (*ValidatorSigningInfo) ProtoMessage()    {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b24ff443e5dfee94, []int{1}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

// Params represents the parameters used for by the slashing module.
type Params struct {
	SignedBlocksWindow      int64                                  `protobuf:"varint,1,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty" yaml:"signed_blocks_window"`
	MinSignedPerWindow      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=min_signed_per_window,json=minSignedPerWindow,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_signed_per_window" yaml:"min_signed_per_window"`
	DowntimeJailDuration    time.Duration                          `protobuf:"bytes,3,opt,name=downtime_jail_duration,json=downtimeJailDuration,proto3,stdduration" json:"downtime_jail_duration" yaml:"downtime_jail_duration"`
	SlashFractionDoubleSign github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slash_fraction_double_sign,json=slashFractionDoubleSign,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_double_sign" yaml:"slash_fraction_double_sign"`
	SlashFractionDowntime   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=slash_fraction_downtime,json=slashFractionDowntime,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_downtime" yaml:"slash_fraction_downtime"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b24ff443e5dfee94, []int{2}
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

func init() {
	proto.RegisterType((*MsgUnjailValidator)(nil), "iritamod.slashing.MsgUnjailValidator")
	proto.RegisterType((*ValidatorSigningInfo)(nil), "iritamod.slashing.ValidatorSigningInfo")
	proto.RegisterType((*Params)(nil), "iritamod.slashing.Params")
}

func init() { proto.RegisterFile("slashing/slashing.proto", fileDescriptor_b24ff443e5dfee94) }

var fileDescriptor_b24ff443e5dfee94 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xbe, 0xf4, 0x4a, 0x29, 0xbe, 0x13, 0x12, 0xee, 0x95, 0x1e, 0x07, 0x24, 0x47, 0x06, 0x54,
	0x86, 0x26, 0x52, 0xd9, 0xca, 0x82, 0x42, 0x85, 0x00, 0x09, 0x28, 0x69, 0x0b, 0x82, 0x81, 0xc8,
	0x39, 0xfb, 0x72, 0x6e, 0x13, 0xfb, 0x14, 0x3b, 0x2a, 0x65, 0x43, 0x62, 0x60, 0xec, 0xc8, 0xc8,
	0xcf, 0xe9, 0xd8, 0x11, 0x31, 0x1c, 0xd0, 0x2e, 0xcc, 0xdd, 0xd8, 0x90, 0xed, 0xa4, 0x3d, 0x95,
	0xeb, 0xd0, 0x29, 0x7e, 0xdf, 0xf7, 0xde, 0xe7, 0xe7, 0xf7, 0x5e, 0x1e, 0x58, 0x10, 0x29, 0x12,
	0x03, 0xca, 0x12, 0xbf, 0x3a, 0x78, 0xc3, 0x9c, 0x4b, 0x0e, 0xaf, 0xd1, 0x9c, 0x4a, 0x94, 0x71,
	0xec, 0x55, 0x44, 0xa7, 0x95, 0xf0, 0x84, 0x6b, 0xd6, 0x57, 0x27, 0xe3, 0xd8, 0xb1, 0x13, 0xce,
	0x93, 0x94, 0xf8, 0xda, 0x8a, 0x8b, 0xbe, 0x8f, 0x8b, 0x1c, 0x49, 0xca, 0x59, 0xc9, 0x3b, 0x67,
	0x79, 0x49, 0x33, 0x22, 0x24, 0xca, 0x86, 0xc6, 0xc1, 0x7d, 0x08, 0xe0, 0x73, 0x91, 0x6c, 0xb2,
	0x2d, 0x44, 0xd3, 0xd7, 0x28, 0xa5, 0x18, 0x49, 0x9e, 0xc3, 0xab, 0x60, 0x8a, 0xe2, 0xb6, 0xd5,
	0xb5, 0x16, 0xaf, 0x84, 0x53, 0x14, 0xc3, 0x0e, 0x98, 0xe5, 0x43, 0x92, 0x2b, 0xae, 0x3d, 0xa5,
	0xd1, 0x13, 0xdb, 0xfd, 0x5c, 0x07, 0xad, 0x93, 0xc8, 0x75, 0x9a, 0x30, 0xca, 0x92, 0xa7, 0xac,
	0xcf, 0x61, 0x1b, 0x5c, 0x46, 0x18, 0xe7, 0x44, 0x88, 0x52, 0xa9, 0x32, 0xe1, 0x0a, 0x68, 0x0a,
	0x89, 0x72, 0x19, 0x0d, 0x08, 0x4d, 0x06, 0x52, 0x4b, 0xd6, 0x83, 0x85, 0xe3, 0x91, 0x33, 0xb7,
	0x8b, 0xb2, 0x74, 0xc5, 0x1d, 0x67, 0xdd, 0xb0, 0xa1, 0xcd, 0x27, 0xda, 0x52, 0xb1, 0x94, 0x61,
	0xf2, 0x21, 0xe2, 0xfd, 0xbe, 0x20, 0xb2, 0x5d, 0x3f, 0x1b, 0x3b, 0xce, 0xba, 0x61, 0x43, 0x9b,
	0x2f, 0xb5, 0x05, 0xdf, 0x83, 0xa6, 0x7a, 0x27, 0xc1, 0x51, 0xc1, 0x24, 0x4d, 0xdb, 0xd3, 0x5d,
	0x6b, 0xb1, 0xb1, 0xdc, 0xf1, 0x4c, 0x91, 0xbc, 0xaa, 0x48, 0xde, 0x46, 0x55, 0xa4, 0xc0, 0xd9,
	0x1f, 0x39, 0xb5, 0x53, 0xed, 0xf1, 0x68, 0x77, 0xef, 0xa7, 0x63, 0x85, 0x0d, 0x03, 0x6d, 0x2a,
	0x04, 0xda, 0x00, 0x48, 0x9e, 0xc5, 0x42, 0x72, 0x46, 0x70, 0xfb, 0x52, 0xd7, 0x5a, 0x9c, 0x0d,
	0xc7, 0x10, 0xb8, 0x01, 0xe6, 0x33, 0x2a, 0x04, 0xc1, 0x51, 0x9c, 0xf2, 0xde, 0xb6, 0x88, 0x7a,
	0xbc, 0x60, 0x92, 0xe4, 0xed, 0x19, 0xfd, 0x88, 0xee, 0xf1, 0xc8, 0xb9, 0x65, 0x2e, 0x9a, 0xe8,
	0xe6, 0x86, 0x73, 0x06, 0x0f, 0x34, 0xfc, 0xc8, 0xa0, 0x2b, 0xd3, 0x7f, 0xbe, 0x39, 0x96, 0xfb,
	0x77, 0x1a, 0xcc, 0xac, 0xa1, 0x1c, 0x65, 0x02, 0xbe, 0x02, 0x2d, 0x41, 0x13, 0x76, 0x1a, 0xbf,
	0x43, 0x19, 0xe6, 0x3b, 0xba, 0x0b, 0xf5, 0xc0, 0x39, 0x1e, 0x39, 0x37, 0xcb, 0x32, 0x4f, 0xf0,
	0x72, 0x43, 0x68, 0x60, 0x73, 0xc9, 0x1b, 0x0d, 0xc2, 0x4f, 0x96, 0x4a, 0x9d, 0x45, 0x65, 0xc4,
	0x90, 0xe4, 0x95, 0xa8, 0xea, 0x5d, 0x33, 0x78, 0xa1, 0xea, 0xf4, 0x63, 0xe4, 0xdc, 0x4d, 0xa8,
	0x1c, 0x14, 0xb1, 0xd7, 0xe3, 0x99, 0xdf, 0xe3, 0x22, 0xe3, 0xa2, 0xfc, 0x2c, 0x09, 0xbc, 0xed,
	0xcb, 0xdd, 0x21, 0x11, 0xde, 0x2a, 0xe9, 0x8d, 0x3f, 0x74, 0x82, 0xa8, 0x1b, 0xc2, 0x8c, 0xb2,
	0x75, 0x0d, 0xaf, 0x91, 0xbc, 0xcc, 0xe1, 0x23, 0xb8, 0x8e, 0xf9, 0x0e, 0x53, 0x13, 0x1c, 0xa9,
	0xaa, 0x47, 0xd5, 0xac, 0xeb, 0x19, 0x68, 0x2c, 0xdf, 0xf8, 0xaf, 0x8f, 0xab, 0xa5, 0x43, 0x70,
	0xaf, 0x6c, 0xe3, 0x6d, 0x73, 0xe9, 0x64, 0x19, 0xf7, 0xab, 0x6a, 0x68, 0xab, 0x22, 0x9f, 0x21,
	0x9a, 0x56, 0x02, 0x70, 0xcf, 0x02, 0x1d, 0xfd, 0x2b, 0x46, 0xfd, 0x1c, 0xf5, 0x14, 0x14, 0x61,
	0x5e, 0xc4, 0x29, 0xd1, 0xc9, 0xeb, 0x41, 0x6a, 0x06, 0xeb, 0x17, 0x2e, 0xc2, 0x9d, 0xb2, 0x0f,
	0xe7, 0x2a, 0xbb, 0xa1, 0xd9, 0x11, 0x8f, 0x4b, 0x6e, 0x55, 0x53, 0xaa, 0x32, 0xf0, 0x8b, 0x55,
	0xee, 0x8f, 0xf1, 0x40, 0x93, 0xba, 0x1e, 0xbd, 0x66, 0xb0, 0x76, 0xe1, 0x7c, 0xec, 0x73, 0xf2,
	0x31, 0xb2, 0x6e, 0x38, 0x7f, 0x26, 0x19, 0x83, 0x07, 0x6f, 0xf7, 0x7f, 0xdb, 0xb5, 0xfd, 0x43,
	0xdb, 0x3a, 0x38, 0xb4, 0xad, 0x5f, 0x87, 0xb6, 0xb5, 0x77, 0x64, 0xd7, 0x0e, 0x8e, 0xec, 0xda,
	0xf7, 0x23, 0xbb, 0xf6, 0xee, 0x41, 0x42, 0x65, 0x8a, 0x62, 0x2f, 0xa6, 0x88, 0x6d, 0x51, 0xe2,
	0x21, 0xea, 0xeb, 0xf5, 0xb6, 0x34, 0xcc, 0xb9, 0x5f, 0x2d, 0x3a, 0x3f, 0xe3, 0xb8, 0x48, 0x89,
	0x38, 0xd9, 0x84, 0x26, 0xaf, 0x78, 0x46, 0x37, 0xf3, 0xfe, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x35, 0x8b, 0x98, 0x2b, 0x05, 0x00, 0x00,
}

func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
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
	if this.Address != that1.Address {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.IndexOffset != that1.IndexOffset {
		return false
	}
	if !this.JailedUntil.Equal(that1.JailedUntil) {
		return false
	}
	if this.Tombstoned != that1.Tombstoned {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	return true
}
func (m *MsgUnjailValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnjailValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnjailValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintSlashing(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSlashing(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSlashing(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.IndexOffset != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.IndexOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSlashing(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	{
		size := m.SlashFractionDowntime.Size()
		i -= size
		if _, err := m.SlashFractionDowntime.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSlashing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SlashFractionDoubleSign.Size()
		i -= size
		if _, err := m.SlashFractionDoubleSign.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSlashing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DowntimeJailDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeJailDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSlashing(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinSignedPerWindow.Size()
		i -= size
		if _, err := m.MinSignedPerWindow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSlashing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SignedBlocksWindow != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.SignedBlocksWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSlashing(dAtA []byte, offset int, v uint64) int {
	offset -= sovSlashing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUnjailValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSlashing(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovSlashing(uint64(l))
	}
	return n
}

func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSlashing(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovSlashing(uint64(m.StartHeight))
	}
	if m.IndexOffset != 0 {
		n += 1 + sovSlashing(uint64(m.IndexOffset))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovSlashing(uint64(l))
	if m.Tombstoned {
		n += 2
	}
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovSlashing(uint64(m.MissedBlocksCounter))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedBlocksWindow != 0 {
		n += 1 + sovSlashing(uint64(m.SignedBlocksWindow))
	}
	l = m.MinSignedPerWindow.Size()
	n += 1 + l + sovSlashing(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeJailDuration)
	n += 1 + l + sovSlashing(uint64(l))
	l = m.SlashFractionDoubleSign.Size()
	n += 1 + l + sovSlashing(uint64(l))
	l = m.SlashFractionDowntime.Size()
	n += 1 + l + sovSlashing(uint64(l))
	return n
}

func sovSlashing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSlashing(x uint64) (n int) {
	return sovSlashing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUnjailValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: MsgUnjailValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnjailValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOffset", wireType)
			}
			m.IndexOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexOffset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
			}
			m.SignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBlocksWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSignedPerWindow", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSignedPerWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DowntimeJailDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DowntimeJailDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionDoubleSign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionDoubleSign.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionDowntime", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionDowntime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSlashing
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
func skipSlashing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
				return 0, ErrInvalidLengthSlashing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSlashing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSlashing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSlashing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlashing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSlashing = fmt.Errorf("proto: unexpected end of group")
)
