// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upgrade/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgUpgradeSoftware - struct for upgrade software
type MsgUpgradeSoftware struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The time after which the upgrade must be performed.
	// Leave set to its zero value to use a pre-defined Height instead.
	Time time.Time `protobuf:"bytes,2,opt,name=time,proto3,stdtime" json:"time"`
	// The height at which the upgrade must be performed.
	// Only used if Time is not set.
	Height int64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info     string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Operator string `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgUpgradeSoftware) Reset()         { *m = MsgUpgradeSoftware{} }
func (m *MsgUpgradeSoftware) String() string { return proto.CompactTextString(m) }
func (*MsgUpgradeSoftware) ProtoMessage()    {}
func (*MsgUpgradeSoftware) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d40a02da488be15, []int{0}
}
func (m *MsgUpgradeSoftware) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpgradeSoftware) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpgradeSoftware.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpgradeSoftware) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpgradeSoftware.Merge(m, src)
}
func (m *MsgUpgradeSoftware) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpgradeSoftware) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpgradeSoftware.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpgradeSoftware proto.InternalMessageInfo

func (m *MsgUpgradeSoftware) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgUpgradeSoftware) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *MsgUpgradeSoftware) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MsgUpgradeSoftware) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *MsgUpgradeSoftware) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

// MsgUpgradeSoftwareResponse defines the Msg/UpgradeSoftware response type.
type MsgUpgradeSoftwareResponse struct {
}

func (m *MsgUpgradeSoftwareResponse) Reset()         { *m = MsgUpgradeSoftwareResponse{} }
func (m *MsgUpgradeSoftwareResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpgradeSoftwareResponse) ProtoMessage()    {}
func (*MsgUpgradeSoftwareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d40a02da488be15, []int{1}
}
func (m *MsgUpgradeSoftwareResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpgradeSoftwareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpgradeSoftwareResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpgradeSoftwareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpgradeSoftwareResponse.Merge(m, src)
}
func (m *MsgUpgradeSoftwareResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpgradeSoftwareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpgradeSoftwareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpgradeSoftwareResponse proto.InternalMessageInfo

// MsgCancelUpgrade - struct for cancel software upgrade
type MsgCancelUpgrade struct {
	Operator string `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (m *MsgCancelUpgrade) Reset()         { *m = MsgCancelUpgrade{} }
func (m *MsgCancelUpgrade) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgrade) ProtoMessage()    {}
func (*MsgCancelUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d40a02da488be15, []int{2}
}
func (m *MsgCancelUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgrade.Merge(m, src)
}
func (m *MsgCancelUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgrade proto.InternalMessageInfo

func (m *MsgCancelUpgrade) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

// MsgCancelUpgradeResponse defines the Msg/CancelUpgrade response type.
type MsgCancelUpgradeResponse struct {
}

func (m *MsgCancelUpgradeResponse) Reset()         { *m = MsgCancelUpgradeResponse{} }
func (m *MsgCancelUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgradeResponse) ProtoMessage()    {}
func (*MsgCancelUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d40a02da488be15, []int{3}
}
func (m *MsgCancelUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgradeResponse.Merge(m, src)
}
func (m *MsgCancelUpgradeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgradeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpgradeSoftware)(nil), "iritamod.upgrade.MsgUpgradeSoftware")
	proto.RegisterType((*MsgUpgradeSoftwareResponse)(nil), "iritamod.upgrade.MsgUpgradeSoftwareResponse")
	proto.RegisterType((*MsgCancelUpgrade)(nil), "iritamod.upgrade.MsgCancelUpgrade")
	proto.RegisterType((*MsgCancelUpgradeResponse)(nil), "iritamod.upgrade.MsgCancelUpgradeResponse")
}

func init() { proto.RegisterFile("upgrade/tx.proto", fileDescriptor_9d40a02da488be15) }

var fileDescriptor_9d40a02da488be15 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x6a, 0xe3, 0x30,
	0x18, 0xc7, 0xad, 0x8b, 0x2f, 0xe4, 0x74, 0x1c, 0x17, 0xc4, 0x71, 0x18, 0x71, 0xd8, 0xc1, 0xdc,
	0x10, 0x8e, 0x43, 0x82, 0x64, 0x29, 0x1d, 0xd3, 0xd9, 0x8b, 0xdb, 0x2e, 0x5d, 0x8a, 0x9c, 0x28,
	0x8a, 0x4a, 0x6c, 0x19, 0x4b, 0xa6, 0xed, 0x5b, 0xe4, 0x11, 0xfa, 0x08, 0x7d, 0x8c, 0x2c, 0x85,
	0x8c, 0x9d, 0xda, 0x92, 0x2c, 0x7d, 0x8c, 0x62, 0x3b, 0x0e, 0x38, 0x29, 0xb4, 0xdb, 0xf7, 0x59,
	0xbf, 0xef, 0xf7, 0xff, 0xb0, 0x04, 0xbb, 0x79, 0x2a, 0x32, 0x36, 0xe1, 0xd4, 0xdc, 0x90, 0x34,
	0x53, 0x46, 0xa1, 0xae, 0xcc, 0xa4, 0x61, 0xb1, 0x9a, 0x90, 0xed, 0x11, 0xfe, 0x25, 0x94, 0x50,
	0xe5, 0x21, 0x2d, 0xaa, 0x8a, 0xc3, 0x9e, 0x50, 0x4a, 0xcc, 0x39, 0x2d, 0xbb, 0x28, 0x9f, 0x52,
	0x23, 0x63, 0xae, 0x0d, 0x8b, 0xd3, 0x0a, 0xf0, 0xef, 0x01, 0x44, 0x81, 0x16, 0xe7, 0x95, 0xe5,
	0x54, 0x4d, 0xcd, 0x35, 0xcb, 0x38, 0x42, 0xd0, 0x4e, 0x58, 0xcc, 0x1d, 0xd0, 0x03, 0xfd, 0x6f,
	0x61, 0x59, 0xa3, 0x23, 0x68, 0x17, 0xd3, 0xce, 0x97, 0x1e, 0xe8, 0x7f, 0x1f, 0x60, 0x52, 0xa9,
	0x49, 0xad, 0x26, 0x67, 0xb5, 0x7a, 0xd4, 0x59, 0x3e, 0x79, 0xd6, 0xe2, 0xd9, 0x03, 0x61, 0x39,
	0x81, 0x7e, 0xc3, 0xf6, 0x8c, 0x4b, 0x31, 0x33, 0x4e, 0xab, 0x07, 0xfa, 0xad, 0x70, 0xdb, 0x15,
	0x29, 0x32, 0x99, 0x2a, 0xc7, 0xae, 0x52, 0x8a, 0x1a, 0x61, 0xd8, 0x51, 0x29, 0xcf, 0x98, 0x51,
	0x99, 0xf3, 0xb5, 0xfc, 0xbe, 0xeb, 0x8f, 0xed, 0xd7, 0x3b, 0x0f, 0xf8, 0x7f, 0x20, 0x3e, 0xdc,
	0x38, 0xe4, 0x3a, 0x55, 0x89, 0xe6, 0x3e, 0x81, 0xdd, 0x40, 0x8b, 0x13, 0x96, 0x8c, 0xf9, 0x7c,
	0xcb, 0x34, 0x9c, 0xa0, 0xe9, 0xf4, 0x31, 0x74, 0xf6, 0xf9, 0xda, 0x35, 0x78, 0x00, 0xb0, 0x15,
	0x68, 0x81, 0x38, 0xfc, 0xb9, 0xff, 0x83, 0xfe, 0x92, 0xfd, 0x1b, 0x20, 0x87, 0x4b, 0xe1, 0xff,
	0x9f, 0xa1, 0xea, 0x38, 0x74, 0x09, 0x7f, 0x34, 0xf7, 0xf6, 0xdf, 0x1d, 0x6f, 0x30, 0xf8, 0xdf,
	0xc7, 0x4c, 0x1d, 0x30, 0x0a, 0x96, 0x6b, 0x17, 0xac, 0xd6, 0x2e, 0x78, 0x59, 0xbb, 0x60, 0xb1,
	0x71, 0xad, 0xd5, 0xc6, 0xb5, 0x1e, 0x37, 0xae, 0x75, 0x31, 0x14, 0xd2, 0xcc, 0xf2, 0x88, 0x8c,
	0x55, 0x4c, 0x23, 0xc9, 0x92, 0x2b, 0xc9, 0x99, 0xa4, 0xb5, 0x99, 0xc6, 0x6a, 0x92, 0xcf, 0xb9,
	0xa6, 0xbb, 0x77, 0x78, 0x9b, 0x72, 0x1d, 0xb5, 0xcb, 0xab, 0x1f, 0xbe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x95, 0x1b, 0x1b, 0x7b, 0x9f, 0x02, 0x00, 0x00,
}

func (this *MsgUpgradeSoftware) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUpgradeSoftware)
	if !ok {
		that2, ok := that.(MsgUpgradeSoftware)
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
	if this.Name != that1.Name {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	if this.Height != that1.Height {
		return false
	}
	if this.Info != that1.Info {
		return false
	}
	if this.Operator != that1.Operator {
		return false
	}
	return true
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
	// CreateClient defines a method for creating a light client.
	UpgradeSoftware(ctx context.Context, in *MsgUpgradeSoftware, opts ...grpc.CallOption) (*MsgUpgradeSoftwareResponse, error)
	// CancelUpgrade defines a method for updating a light client state.
	CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpgradeSoftware(ctx context.Context, in *MsgUpgradeSoftware, opts ...grpc.CallOption) (*MsgUpgradeSoftwareResponse, error) {
	out := new(MsgUpgradeSoftwareResponse)
	err := c.cc.Invoke(ctx, "/iritamod.upgrade.Msg/UpgradeSoftware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error) {
	out := new(MsgCancelUpgradeResponse)
	err := c.cc.Invoke(ctx, "/iritamod.upgrade.Msg/CancelUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateClient defines a method for creating a light client.
	UpgradeSoftware(context.Context, *MsgUpgradeSoftware) (*MsgUpgradeSoftwareResponse, error)
	// CancelUpgrade defines a method for updating a light client state.
	CancelUpgrade(context.Context, *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpgradeSoftware(ctx context.Context, req *MsgUpgradeSoftware) (*MsgUpgradeSoftwareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeSoftware not implemented")
}
func (*UnimplementedMsgServer) CancelUpgrade(ctx context.Context, req *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUpgrade not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpgradeSoftware_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpgradeSoftware)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpgradeSoftware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iritamod.upgrade.Msg/UpgradeSoftware",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpgradeSoftware(ctx, req.(*MsgUpgradeSoftware))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iritamod.upgrade.Msg/CancelUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelUpgrade(ctx, req.(*MsgCancelUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iritamod.upgrade.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpgradeSoftware",
			Handler:    _Msg_UpgradeSoftware_Handler,
		},
		{
			MethodName: "CancelUpgrade",
			Handler:    _Msg_CancelUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upgrade/tx.proto",
}

func (m *MsgUpgradeSoftware) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpgradeSoftware) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpgradeSoftware) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpgradeSoftwareResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpgradeSoftwareResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpgradeSoftwareResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Operator) > 0 {
		i -= len(m.Operator)
		copy(dAtA[i:], m.Operator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Operator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpgradeSoftware) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTx(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTx(uint64(m.Height))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpgradeSoftwareResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCancelUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Operator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCancelUpgradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpgradeSoftware) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpgradeSoftware: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpgradeSoftware: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpgradeSoftwareResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpgradeSoftwareResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpgradeSoftwareResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCancelUpgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCancelUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCancelUpgradeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
