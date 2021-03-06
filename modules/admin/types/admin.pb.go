// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: admin/admin.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

// Role is a type alias that represents a proposal status as a byte
type Role int32

const (
	// ROOT_ADMIN defines the root admin role index.
	RoleRootAdmin Role = 0
	// PERM_ADMIN defines the permission admin role index.
	RolePermAdmin Role = 1
	// BLACKLIST_ADMIN defines the blacklist admin role index.
	RoleBlacklistAdmin Role = 2
	// NODE_ADMIN defines the validator node admin role index.
	RoleNodeAdmin Role = 3
	// PARAM_ADMIN defines the param admin role index.
	RoleParamAdmin Role = 4
	// POWER_USER defines the power user role index.
	RolePowerUser Role = 5
	// RELAYER_USER defines the relayer role index.
	RoleRelayerUser Role = 6
	// IDAdmin defines the identity role index.
	RoleIDAdmin Role = 7
)

var Role_name = map[int32]string{
	0: "ROOT_ADMIN",
	1: "PERM_ADMIN",
	2: "BLACKLIST_ADMIN",
	3: "NODE_ADMIN",
	4: "PARAM_ADMIN",
	5: "POWER_USER",
	6: "RELAYER_USER",
	7: "ID_ADMIN",
}

var Role_value = map[string]int32{
	"ROOT_ADMIN":      0,
	"PERM_ADMIN":      1,
	"BLACKLIST_ADMIN": 2,
	"NODE_ADMIN":      3,
	"PARAM_ADMIN":     4,
	"POWER_USER":      5,
	"RELAYER_USER":    6,
	"ID_ADMIN":        7,
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8595c8dce2486799, []int{0}
}

func init() {
	proto.RegisterEnum("iritamod.admin.Role", Role_name, Role_value)
}

func init() { proto.RegisterFile("admin/admin.proto", fileDescriptor_8595c8dce2486799) }

var fileDescriptor_8595c8dce2486799 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4c, 0xc9, 0xcd,
	0xcc, 0xd3, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x7c, 0x99, 0x45, 0x99, 0x25,
	0x89, 0xb9, 0xf9, 0x29, 0x7a, 0x60, 0x51, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x94, 0x3e,
	0x88, 0x05, 0x51, 0xa5, 0xb5, 0x83, 0x89, 0x8b, 0x25, 0x28, 0x3f, 0x27, 0x55, 0x48, 0x91, 0x8b,
	0x2b, 0xc8, 0xdf, 0x3f, 0x24, 0xde, 0xd1, 0xc5, 0xd7, 0xd3, 0x4f, 0x80, 0x41, 0x4a, 0xb0, 0x6b,
	0xae, 0x02, 0x2f, 0x48, 0x26, 0x28, 0x3f, 0xbf, 0xc4, 0x11, 0x64, 0x02, 0x48, 0x49, 0x80, 0x6b,
	0x90, 0x2f, 0x54, 0x09, 0x23, 0x42, 0x49, 0x40, 0x6a, 0x51, 0x2e, 0x44, 0x89, 0x36, 0x17, 0xbf,
	0x93, 0x8f, 0xa3, 0xb3, 0xb7, 0x8f, 0x67, 0x30, 0xcc, 0x28, 0x26, 0x29, 0xb1, 0xae, 0xb9, 0x0a,
	0x42, 0x20, 0x75, 0x4e, 0x39, 0x89, 0xc9, 0xd9, 0x39, 0x99, 0xc5, 0x08, 0xf3, 0xfc, 0xfc, 0x5d,
	0x5c, 0xa1, 0xea, 0x98, 0x11, 0xe6, 0xf9, 0xe5, 0xa7, 0xa4, 0x42, 0x94, 0x28, 0x73, 0x71, 0x07,
	0x38, 0x06, 0x39, 0xc2, 0xec, 0x64, 0x91, 0x12, 0xea, 0x9a, 0xab, 0xc0, 0x07, 0xb6, 0x33, 0xb1,
	0x28, 0x31, 0x17, 0xe1, 0x2e, 0xff, 0x70, 0xd7, 0xa0, 0xf8, 0xd0, 0x60, 0xd7, 0x20, 0x01, 0x56,
	0x24, 0x77, 0xe5, 0x97, 0xa7, 0x16, 0x85, 0x16, 0xa7, 0x16, 0x09, 0xa9, 0x72, 0xf1, 0x04, 0xb9,
	0xfa, 0x38, 0x46, 0xc2, 0x14, 0xb1, 0x49, 0x09, 0x77, 0xcd, 0x55, 0xe0, 0x07, 0xfb, 0x2f, 0x35,
	0x27, 0xb1, 0x12, 0xaa, 0x4c, 0x96, 0x8b, 0xc3, 0xd3, 0x05, 0x6a, 0x17, 0xbb, 0x14, 0x7f, 0xd7,
	0x5c, 0x05, 0x6e, 0x90, 0x12, 0x4f, 0x17, 0xb0, 0x45, 0x52, 0x3c, 0x1d, 0x8b, 0xe5, 0x18, 0x56,
	0x2c, 0x91, 0x63, 0xd8, 0xb0, 0x44, 0x8e, 0xd1, 0xc9, 0xff, 0xc4, 0x43, 0x39, 0x86, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xca, 0x4c, 0xcc, 0xcb, 0xca, 0x4c, 0x4d, 0xcc, 0xd4, 0x87, 0xc5,
	0x89, 0x7e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e, 0x6a, 0x31, 0x24, 0xc6, 0xf4, 0x4b, 0x2a, 0x0b, 0x52,
	0x8b, 0x93, 0xd8, 0xc0, 0x51, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x63, 0x2d, 0xec, 0x0d,
	0xcd, 0x01, 0x00, 0x00,
}

func (x Role) String() string {
	s, ok := Role_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
