// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

/*
Package game is a generated protocol buffer package.

It is generated from these files:
	base.proto

It has these top-level messages:
	CSErrorMsg
	CSLoginReq
	CSLoginResp
	C2ServerMsg
	S2ClientMsg
*/
package game

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CS_CMD_ID int32

const (
	CS_CMD_ID_CS_CMD_LOGIN              CS_CMD_ID = 500
	CS_CMD_ID_CS_CMD_HEART_BEAT         CS_CMD_ID = 501
	CS_CMD_ID_CS_CMD_ENTER_SERVER       CS_CMD_ID = 502
	CS_CMD_ID_CS_CMD_LOGOUT             CS_CMD_ID = 503
	CS_CMD_ID_CS_CMD_DIRECT_ENTER       CS_CMD_ID = 504
	CS_CMD_ID_CS_CMD_CANENTER           CS_CMD_ID = 505
	CS_CMD_ID_CS_CMD_QUEUE              CS_CMD_ID = 506
	CS_CMD_ID_CS_CMD_CHANGE_NAME        CS_CMD_ID = 507
	CS_CMD_ID_CS_CMD_GM_COMMAND         CS_CMD_ID = 508
	CS_CMD_ID_CS_CMD_USE_GIFT_CODE      CS_CMD_ID = 510
	CS_CMD_ID_CS_CMD_QUERY_PLAYERINFO   CS_CMD_ID = 511
	CS_CMD_ID_CS_CMD_CHANGE_PLAYER_ICON CS_CMD_ID = 512
	CS_CMD_ID_CS_CMD_SET_NEWBIE         CS_CMD_ID = 513
	CS_CMD_ID_CS_CMD_GET_NEWBIE_LIST    CS_CMD_ID = 514
	CS_CMD_ID_CS_CMD_SHOP_GUIDE         CS_CMD_ID = 515
	CS_CMD_ID_CS_CMD_SET_TITLE          CS_CMD_ID = 517
	CS_CMD_ID_CS_CMD_SHOP_BUY           CS_CMD_ID = 518
	CS_CMD_ID_CS_CMD_KICK_PLAYER        CS_CMD_ID = 519
	CS_CMD_ID_CS_CMD_STOP_KICK_PLAYLER  CS_CMD_ID = 520
)

var CS_CMD_ID_name = map[int32]string{
	500: "CS_CMD_LOGIN",
	501: "CS_CMD_HEART_BEAT",
	502: "CS_CMD_ENTER_SERVER",
	503: "CS_CMD_LOGOUT",
	504: "CS_CMD_DIRECT_ENTER",
	505: "CS_CMD_CANENTER",
	506: "CS_CMD_QUEUE",
	507: "CS_CMD_CHANGE_NAME",
	508: "CS_CMD_GM_COMMAND",
	510: "CS_CMD_USE_GIFT_CODE",
	511: "CS_CMD_QUERY_PLAYERINFO",
	512: "CS_CMD_CHANGE_PLAYER_ICON",
	513: "CS_CMD_SET_NEWBIE",
	514: "CS_CMD_GET_NEWBIE_LIST",
	515: "CS_CMD_SHOP_GUIDE",
	517: "CS_CMD_SET_TITLE",
	518: "CS_CMD_SHOP_BUY",
	519: "CS_CMD_KICK_PLAYER",
	520: "CS_CMD_STOP_KICK_PLAYLER",
}
var CS_CMD_ID_value = map[string]int32{
	"CS_CMD_LOGIN":              500,
	"CS_CMD_HEART_BEAT":         501,
	"CS_CMD_ENTER_SERVER":       502,
	"CS_CMD_LOGOUT":             503,
	"CS_CMD_DIRECT_ENTER":       504,
	"CS_CMD_CANENTER":           505,
	"CS_CMD_QUEUE":              506,
	"CS_CMD_CHANGE_NAME":        507,
	"CS_CMD_GM_COMMAND":         508,
	"CS_CMD_USE_GIFT_CODE":      510,
	"CS_CMD_QUERY_PLAYERINFO":   511,
	"CS_CMD_CHANGE_PLAYER_ICON": 512,
	"CS_CMD_SET_NEWBIE":         513,
	"CS_CMD_GET_NEWBIE_LIST":    514,
	"CS_CMD_SHOP_GUIDE":         515,
	"CS_CMD_SET_TITLE":          517,
	"CS_CMD_SHOP_BUY":           518,
	"CS_CMD_KICK_PLAYER":        519,
	"CS_CMD_STOP_KICK_PLAYLER":  520,
}

func (x CS_CMD_ID) Enum() *CS_CMD_ID {
	p := new(CS_CMD_ID)
	*p = x
	return p
}
func (x CS_CMD_ID) String() string {
	return proto.EnumName(CS_CMD_ID_name, int32(x))
}
func (x *CS_CMD_ID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CS_CMD_ID_value, data, "CS_CMD_ID")
	if err != nil {
		return err
	}
	*x = CS_CMD_ID(value)
	return nil
}
func (CS_CMD_ID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CSErrorMsg struct {
	ErrCode          *uint32 `protobuf:"varint,1,req,name=ErrCode" json:"ErrCode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSErrorMsg) Reset()                    { *m = CSErrorMsg{} }
func (m *CSErrorMsg) String() string            { return proto.CompactTextString(m) }
func (*CSErrorMsg) ProtoMessage()               {}
func (*CSErrorMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CSErrorMsg) GetErrCode() uint32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

type CSLoginReq struct {
	OpenId           *string `protobuf:"bytes,1,req,name=OpenId" json:"OpenId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSLoginReq) Reset()                    { *m = CSLoginReq{} }
func (m *CSLoginReq) String() string            { return proto.CompactTextString(m) }
func (*CSLoginReq) ProtoMessage()               {}
func (*CSLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CSLoginReq) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

type CSLoginResp struct {
	Uin              *uint64 `protobuf:"varint,1,req,name=Uin" json:"Uin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSLoginResp) Reset()                    { *m = CSLoginResp{} }
func (m *CSLoginResp) String() string            { return proto.CompactTextString(m) }
func (*CSLoginResp) ProtoMessage()               {}
func (*CSLoginResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CSLoginResp) GetUin() uint64 {
	if m != nil && m.Uin != nil {
		return *m.Uin
	}
	return 0
}

type C2ServerMsg struct {
	Login_Req        *CSLoginReq `protobuf:"bytes,2,opt,name=Login_Req" json:"Login_Req,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *C2ServerMsg) Reset()                    { *m = C2ServerMsg{} }
func (m *C2ServerMsg) String() string            { return proto.CompactTextString(m) }
func (*C2ServerMsg) ProtoMessage()               {}
func (*C2ServerMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *C2ServerMsg) GetLogin_Req() *CSLoginReq {
	if m != nil {
		return m.Login_Req
	}
	return nil
}

type S2ClientMsg struct {
	Error_Msg        *CSErrorMsg  `protobuf:"bytes,1,opt,name=Error_Msg" json:"Error_Msg,omitempty"`
	Login_Resp       *CSLoginResp `protobuf:"bytes,2,opt,name=Login_Resp" json:"Login_Resp,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *S2ClientMsg) Reset()                    { *m = S2ClientMsg{} }
func (m *S2ClientMsg) String() string            { return proto.CompactTextString(m) }
func (*S2ClientMsg) ProtoMessage()               {}
func (*S2ClientMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *S2ClientMsg) GetError_Msg() *CSErrorMsg {
	if m != nil {
		return m.Error_Msg
	}
	return nil
}

func (m *S2ClientMsg) GetLogin_Resp() *CSLoginResp {
	if m != nil {
		return m.Login_Resp
	}
	return nil
}

func init() {
	proto.RegisterType((*CSErrorMsg)(nil), "game.CSErrorMsg")
	proto.RegisterType((*CSLoginReq)(nil), "game.CSLoginReq")
	proto.RegisterType((*CSLoginResp)(nil), "game.CSLoginResp")
	proto.RegisterType((*C2ServerMsg)(nil), "game.C2ServerMsg")
	proto.RegisterType((*S2ClientMsg)(nil), "game.S2ClientMsg")
	proto.RegisterEnum("game.CS_CMD_ID", CS_CMD_ID_name, CS_CMD_ID_value)
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xd5, 0x8e, 0x05, 0xca, 0x09, 0xa5, 0x93, 0x43, 0x69, 0x53, 0x68, 0x51, 0x15, 0x84,
	0x54, 0xb1, 0xc8, 0x22, 0x6f, 0xe0, 0x8c, 0x4f, 0x9d, 0x51, 0x7d, 0x29, 0xf6, 0x18, 0x94, 0xd5,
	0xa8, 0xa8, 0x56, 0x14, 0x09, 0x62, 0x63, 0x57, 0xac, 0xb9, 0x88, 0xcb, 0x7b, 0xf2, 0x0a, 0xdc,
	0x41, 0xc0, 0xd4, 0x71, 0x12, 0x87, 0xdd, 0xcc, 0xf9, 0xfe, 0xf3, 0xcd, 0xaf, 0x01, 0x78, 0x7a,
	0x5e, 0xa6, 0xfd, 0xbc, 0xc8, 0x2e, 0x33, 0xb4, 0x26, 0xe7, 0xcf, 0xd3, 0xde, 0x21, 0x80, 0x88,
	0xa9, 0x28, 0xb2, 0xc2, 0x2f, 0x27, 0xb8, 0x0d, 0xd7, 0xcd, 0x59, 0x64, 0x17, 0x69, 0x77, 0xe3,
	0x68, 0xf3, 0x78, 0xab, 0x77, 0x70, 0x85, 0xbd, 0x6c, 0x32, 0x9d, 0x45, 0xe9, 0x0b, 0xbc, 0x09,
	0xd7, 0xc2, 0x3c, 0x9d, 0xc9, 0x8b, 0x8a, 0xb6, 0x7a, 0x77, 0xa0, 0xbd, 0xa4, 0x65, 0x8e, 0x6d,
	0x60, 0xc9, 0x74, 0x56, 0x31, 0xab, 0x37, 0x30, 0x6c, 0x10, 0xa7, 0xc5, 0xcb, 0xb4, 0x32, 0xdf,
	0x87, 0x56, 0x15, 0xd4, 0xc6, 0xd3, 0xdd, 0x3c, 0xda, 0x38, 0x6e, 0x0f, 0x78, 0xff, 0xaa, 0x41,
	0x7f, 0xe5, 0xef, 0x8d, 0xa1, 0x1d, 0x0f, 0xc4, 0xb3, 0x69, 0x3a, 0xbb, 0xac, 0x77, 0xaa, 0x66,
	0xda, 0x5c, 0x8c, 0x75, 0x6d, 0x67, 0x59, 0xf9, 0x01, 0xc0, 0x42, 0x5c, 0xe6, 0xb5, 0xb9, 0xf3,
	0x9f, 0xb9, 0xcc, 0x1f, 0x7e, 0x62, 0xd0, 0x12, 0xb1, 0x16, 0xbe, 0xa3, 0xa5, 0x83, 0x1d, 0xb8,
	0x51, 0x5f, 0xbc, 0xd0, 0x95, 0x01, 0xff, 0xcc, 0x70, 0x17, 0x3a, 0xf5, 0x68, 0x44, 0x76, 0xa4,
	0xf4, 0x90, 0x6c, 0xc5, 0xbf, 0x30, 0xec, 0xc2, 0xad, 0x7a, 0x4e, 0x81, 0xa2, 0x48, 0xc7, 0x14,
	0x3d, 0xa6, 0x88, 0x7f, 0x65, 0x88, 0xb0, 0xb5, 0x92, 0x84, 0x89, 0xe2, 0xdf, 0x9a, 0x69, 0x47,
	0x46, 0x24, 0xd4, 0x7c, 0x89, 0x7f, 0x67, 0xb8, 0x03, 0xdb, 0x35, 0x11, 0x76, 0x30, 0x9f, 0xfe,
	0x60, 0x8d, 0x22, 0x8f, 0x12, 0x4a, 0x88, 0xff, 0x64, 0xb8, 0x07, 0xb8, 0x08, 0x8e, 0xec, 0xc0,
	0x25, 0x1d, 0xd8, 0x3e, 0xf1, 0x5f, 0xcd, 0x86, 0xae, 0xaf, 0x45, 0xe8, 0xfb, 0x76, 0xe0, 0xf0,
	0xdf, 0x0c, 0xf7, 0x61, 0xa7, 0x9e, 0x27, 0x31, 0x69, 0x57, 0x9e, 0x28, 0x43, 0x1d, 0xe2, 0x7f,
	0x18, 0x1e, 0xc0, 0xde, 0x4a, 0x1f, 0x8d, 0xf5, 0x99, 0x67, 0x8f, 0x29, 0x92, 0xc1, 0x49, 0xc8,
	0xff, 0x32, 0xbc, 0x07, 0xfb, 0xeb, 0x2f, 0xcd, 0xb1, 0x96, 0x22, 0x0c, 0xf8, 0x2b, 0xab, 0xf1,
	0x60, 0x4c, 0x4a, 0x07, 0xf4, 0x64, 0x28, 0x89, 0xbf, 0xb6, 0xf0, 0x2e, 0xec, 0x2e, 0x8a, 0x2c,
	0xe7, 0xda, 0x93, 0xb1, 0xe2, 0x6f, 0xd6, 0x96, 0x46, 0xe1, 0x99, 0x76, 0x13, 0x69, 0xaa, 0xbc,
	0xb5, 0xf0, 0x36, 0xf0, 0x86, 0x4c, 0x49, 0xe5, 0x11, 0x7f, 0x67, 0x35, 0xbe, 0xa5, 0x8a, 0x0f,
	0x93, 0x31, 0x7f, 0x6f, 0x35, 0xfe, 0xe0, 0x54, 0x8a, 0xd3, 0xba, 0x17, 0xff, 0x60, 0xe1, 0x21,
	0x74, 0x17, 0x71, 0x65, 0xe2, 0x4b, 0xea, 0x19, 0xfc, 0xd1, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0xbe, 0x24, 0x75, 0xe0, 0x02, 0x00, 0x00,
}
