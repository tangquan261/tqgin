// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userinfo.proto

package login

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserInfo struct {
	PlayerID             int64    `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	PlayerName           string   `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
	Sex                  int32    `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	PlayerIcon           string   `protobuf:"bytes,4,opt,name=playerIcon,proto3" json:"playerIcon,omitempty"`
	TitleID              int32    `protobuf:"varint,5,opt,name=titleID,proto3" json:"titleID,omitempty"`
	TitleName            string   `protobuf:"bytes,6,opt,name=titleName,proto3" json:"titleName,omitempty"`
	Age                  int32    `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`
	Auth                 int32    `protobuf:"varint,8,opt,name=Auth,proto3" json:"Auth,omitempty"`
	Sign                 string   `protobuf:"bytes,9,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetPlayerID() int64 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *UserInfo) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *UserInfo) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *UserInfo) GetPlayerIcon() string {
	if m != nil {
		return m.PlayerIcon
	}
	return ""
}

func (m *UserInfo) GetTitleID() int32 {
	if m != nil {
		return m.TitleID
	}
	return 0
}

func (m *UserInfo) GetTitleName() string {
	if m != nil {
		return m.TitleName
	}
	return ""
}

func (m *UserInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfo) GetAuth() int32 {
	if m != nil {
		return m.Auth
	}
	return 0
}

func (m *UserInfo) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

type RoomUserInfo struct {
	UserInfo             *UserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoomUserInfo) Reset()         { *m = RoomUserInfo{} }
func (m *RoomUserInfo) String() string { return proto.CompactTextString(m) }
func (*RoomUserInfo) ProtoMessage()    {}
func (*RoomUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{1}
}

func (m *RoomUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomUserInfo.Unmarshal(m, b)
}
func (m *RoomUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomUserInfo.Marshal(b, m, deterministic)
}
func (m *RoomUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomUserInfo.Merge(m, src)
}
func (m *RoomUserInfo) XXX_Size() int {
	return xxx_messageInfo_RoomUserInfo.Size(m)
}
func (m *RoomUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomUserInfo proto.InternalMessageInfo

func (m *RoomUserInfo) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type OneUserinfo struct {
	UserInfo             *UserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OneUserinfo) Reset()         { *m = OneUserinfo{} }
func (m *OneUserinfo) String() string { return proto.CompactTextString(m) }
func (*OneUserinfo) ProtoMessage()    {}
func (*OneUserinfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{2}
}

func (m *OneUserinfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneUserinfo.Unmarshal(m, b)
}
func (m *OneUserinfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneUserinfo.Marshal(b, m, deterministic)
}
func (m *OneUserinfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneUserinfo.Merge(m, src)
}
func (m *OneUserinfo) XXX_Size() int {
	return xxx_messageInfo_OneUserinfo.Size(m)
}
func (m *OneUserinfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OneUserinfo.DiscardUnknown(m)
}

var xxx_messageInfo_OneUserinfo proto.InternalMessageInfo

func (m *OneUserinfo) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "login.userInfo")
	proto.RegisterType((*RoomUserInfo)(nil), "login.RoomUserInfo")
	proto.RegisterType((*OneUserinfo)(nil), "login.OneUserinfo")
}

func init() { proto.RegisterFile("userinfo.proto", fileDescriptor_785a78c34699a93d) }

var fileDescriptor_785a78c34699a93d = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0xd3, 0xb4, 0xc9, 0x54, 0x54, 0xe6, 0x34, 0x88, 0x48, 0xc8, 0x29, 0x20, 0xe4,
	0xa0, 0x37, 0x3d, 0x09, 0xbd, 0xe4, 0xa2, 0xb0, 0xd0, 0x07, 0x88, 0x32, 0x8d, 0x81, 0x74, 0xa7,
	0x24, 0x5b, 0xd0, 0x57, 0xf6, 0x29, 0x64, 0xa7, 0xa6, 0xe9, 0xb5, 0xb7, 0x7f, 0xbe, 0xdd, 0x8f,
	0x61, 0x7e, 0xb8, 0xda, 0x0f, 0xdc, 0xb7, 0x6e, 0x23, 0xe5, 0xae, 0x17, 0x2f, 0x18, 0x77, 0xd2,
	0xb4, 0x2e, 0xff, 0x35, 0x90, 0x84, 0x97, 0xca, 0x6d, 0x04, 0x6f, 0x21, 0xd9, 0x75, 0xf5, 0x0f,
	0xf7, 0xd5, 0x8a, 0x4c, 0x66, 0x8a, 0xc8, 0x1e, 0x67, 0xbc, 0x07, 0x38, 0xe4, 0xb7, 0x7a, 0xcb,
	0x74, 0x91, 0x99, 0x22, 0xb5, 0x27, 0x04, 0x6f, 0x20, 0x1a, 0xf8, 0x9b, 0xa2, 0xcc, 0x14, 0xb1,
	0x0d, 0x71, 0x32, 0xaa, 0x4f, 0x71, 0x34, 0x3b, 0x35, 0x02, 0x41, 0x82, 0x85, 0x6f, 0x7d, 0xc7,
	0xd5, 0x8a, 0x62, 0xb5, 0xc6, 0x11, 0xef, 0x20, 0xd5, 0xa8, 0xab, 0xe6, 0x2a, 0x4e, 0x20, 0x6c,
	0xaa, 0x1b, 0xa6, 0xc5, 0x61, 0x53, 0xdd, 0x30, 0x22, 0xcc, 0x5e, 0xf7, 0xfe, 0x8b, 0x12, 0x45,
	0x9a, 0x03, 0x1b, 0xda, 0xc6, 0x51, 0xaa, 0xba, 0xe6, 0xfc, 0x05, 0x2e, 0xad, 0xc8, 0x76, 0x3d,
	0xde, 0xfb, 0x30, 0xdd, 0xae, 0xf7, 0x2e, 0x1f, 0xaf, 0x4b, 0xad, 0xa5, 0x1c, 0xb1, 0x3d, 0x7e,
	0xc8, 0x9f, 0x61, 0xf9, 0xee, 0x78, 0xfd, 0xdf, 0xe2, 0x59, 0xee, 0xc7, 0x5c, 0x3b, 0x7f, 0xfa,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x22, 0xdb, 0x5d, 0x85, 0x01, 0x00, 0x00,
}
