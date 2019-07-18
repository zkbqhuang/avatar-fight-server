// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/smsg/smsg.proto

/*
Package smsg is a generated protocol buffer package.

It is generated from these files:
	msg/smsg/smsg.proto

It has these top-level messages:
	Server2AllSession
	GaCeReqLogin
	GaCeRespLogin
	CeGaBindGameServer
	CeGamReqJoinGame
	CeGamRespJoinGame
	GamCeNoticeGameStart
	GamCeNoticeGameEnd
	CeGameUserDisconnect
	CeGameUserReconnect
	GaCeUserDisconnect
	CeGaCloseSession
*/
package smsg

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

type GaCeRespLogin_Error int32

const (
	GaCeRespLogin_Invalid GaCeRespLogin_Error = 0
)

var GaCeRespLogin_Error_name = map[int32]string{
	0: "Invalid",
}
var GaCeRespLogin_Error_value = map[string]int32{
	"Invalid": 0,
}

func (x GaCeRespLogin_Error) String() string {
	return proto.EnumName(GaCeRespLogin_Error_name, int32(x))
}
func (GaCeRespLogin_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type CeGamRespJoinGame_Error int32

const (
	CeGamRespJoinGame_Invalid      CeGamRespJoinGame_Error = 0
	CeGamRespJoinGame_GameNotExist CeGamRespJoinGame_Error = 1
)

var CeGamRespJoinGame_Error_name = map[int32]string{
	0: "Invalid",
	1: "GameNotExist",
}
var CeGamRespJoinGame_Error_value = map[string]int32{
	"Invalid":      0,
	"GameNotExist": 1,
}

func (x CeGamRespJoinGame_Error) String() string {
	return proto.EnumName(CeGamRespJoinGame_Error_name, int32(x))
}
func (CeGamRespJoinGame_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Server2AllSession struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Server2AllSession) Reset()                    { *m = Server2AllSession{} }
func (m *Server2AllSession) String() string            { return proto.CompactTextString(m) }
func (*Server2AllSession) ProtoMessage()               {}
func (*Server2AllSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Server2AllSession) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GaCeReqLogin struct {
	Sesid int32  `protobuf:"varint,1,opt,name=sesid" json:"sesid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *GaCeReqLogin) Reset()                    { *m = GaCeReqLogin{} }
func (m *GaCeReqLogin) String() string            { return proto.CompactTextString(m) }
func (*GaCeReqLogin) ProtoMessage()               {}
func (*GaCeReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GaCeReqLogin) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *GaCeReqLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GaCeRespLogin struct {
	UserID uint64              `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Err    GaCeRespLogin_Error `protobuf:"varint,2,opt,name=err,enum=smsg.GaCeRespLogin_Error" json:"err,omitempty"`
	Token  string              `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	InGame bool                `protobuf:"varint,4,opt,name=inGame" json:"inGame,omitempty"`
}

func (m *GaCeRespLogin) Reset()                    { *m = GaCeRespLogin{} }
func (m *GaCeRespLogin) String() string            { return proto.CompactTextString(m) }
func (*GaCeRespLogin) ProtoMessage()               {}
func (*GaCeRespLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GaCeRespLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GaCeRespLogin) GetErr() GaCeRespLogin_Error {
	if m != nil {
		return m.Err
	}
	return GaCeRespLogin_Invalid
}

func (m *GaCeRespLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GaCeRespLogin) GetInGame() bool {
	if m != nil {
		return m.InGame
	}
	return false
}

type CeGaBindGameServer struct {
	Sesid        int32 `protobuf:"varint,1,opt,name=sesid" json:"sesid,omitempty"`
	Gameserverid int32 `protobuf:"varint,2,opt,name=gameserverid" json:"gameserverid,omitempty"`
}

func (m *CeGaBindGameServer) Reset()                    { *m = CeGaBindGameServer{} }
func (m *CeGaBindGameServer) String() string            { return proto.CompactTextString(m) }
func (*CeGaBindGameServer) ProtoMessage()               {}
func (*CeGaBindGameServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CeGaBindGameServer) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *CeGaBindGameServer) GetGameserverid() int32 {
	if m != nil {
		return m.Gameserverid
	}
	return 0
}

type CeGamReqJoinGame struct {
	Userid       uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Sesid        int32  `protobuf:"varint,2,opt,name=sesid" json:"sesid,omitempty"`
	Nickname     string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	GateServerid int32  `protobuf:"varint,4,opt,name=gateServerid" json:"gateServerid,omitempty"`
}

func (m *CeGamReqJoinGame) Reset()                    { *m = CeGamReqJoinGame{} }
func (m *CeGamReqJoinGame) String() string            { return proto.CompactTextString(m) }
func (*CeGamReqJoinGame) ProtoMessage()               {}
func (*CeGamReqJoinGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CeGamReqJoinGame) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *CeGamReqJoinGame) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *CeGamReqJoinGame) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *CeGamReqJoinGame) GetGateServerid() int32 {
	if m != nil {
		return m.GateServerid
	}
	return 0
}

type CeGamRespJoinGame struct {
	Err    CeGamRespJoinGame_Error `protobuf:"varint,1,opt,name=err,enum=smsg.CeGamRespJoinGame_Error" json:"err,omitempty"`
	Gameid int64                   `protobuf:"varint,2,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *CeGamRespJoinGame) Reset()                    { *m = CeGamRespJoinGame{} }
func (m *CeGamRespJoinGame) String() string            { return proto.CompactTextString(m) }
func (*CeGamRespJoinGame) ProtoMessage()               {}
func (*CeGamRespJoinGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CeGamRespJoinGame) GetErr() CeGamRespJoinGame_Error {
	if m != nil {
		return m.Err
	}
	return CeGamRespJoinGame_Invalid
}

func (m *CeGamRespJoinGame) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type GamCeNoticeGameStart struct {
	Gameid int64 `protobuf:"varint,1,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *GamCeNoticeGameStart) Reset()                    { *m = GamCeNoticeGameStart{} }
func (m *GamCeNoticeGameStart) String() string            { return proto.CompactTextString(m) }
func (*GamCeNoticeGameStart) ProtoMessage()               {}
func (*GamCeNoticeGameStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GamCeNoticeGameStart) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type GamCeNoticeGameEnd struct {
	Gameid int64 `protobuf:"varint,1,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *GamCeNoticeGameEnd) Reset()                    { *m = GamCeNoticeGameEnd{} }
func (m *GamCeNoticeGameEnd) String() string            { return proto.CompactTextString(m) }
func (*GamCeNoticeGameEnd) ProtoMessage()               {}
func (*GamCeNoticeGameEnd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GamCeNoticeGameEnd) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type CeGameUserDisconnect struct {
	Userid uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
}

func (m *CeGameUserDisconnect) Reset()                    { *m = CeGameUserDisconnect{} }
func (m *CeGameUserDisconnect) String() string            { return proto.CompactTextString(m) }
func (*CeGameUserDisconnect) ProtoMessage()               {}
func (*CeGameUserDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CeGameUserDisconnect) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type CeGameUserReconnect struct {
	Userid    uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	GateID    int32  `protobuf:"varint,2,opt,name=gateID" json:"gateID,omitempty"`
	SessionID int32  `protobuf:"varint,3,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *CeGameUserReconnect) Reset()                    { *m = CeGameUserReconnect{} }
func (m *CeGameUserReconnect) String() string            { return proto.CompactTextString(m) }
func (*CeGameUserReconnect) ProtoMessage()               {}
func (*CeGameUserReconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CeGameUserReconnect) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *CeGameUserReconnect) GetGateID() int32 {
	if m != nil {
		return m.GateID
	}
	return 0
}

func (m *CeGameUserReconnect) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

type GaCeUserDisconnect struct {
	SessionID int32 `protobuf:"varint,1,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *GaCeUserDisconnect) Reset()                    { *m = GaCeUserDisconnect{} }
func (m *GaCeUserDisconnect) String() string            { return proto.CompactTextString(m) }
func (*GaCeUserDisconnect) ProtoMessage()               {}
func (*GaCeUserDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GaCeUserDisconnect) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

type CeGaCloseSession struct {
	SessionID int32 `protobuf:"varint,1,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *CeGaCloseSession) Reset()                    { *m = CeGaCloseSession{} }
func (m *CeGaCloseSession) String() string            { return proto.CompactTextString(m) }
func (*CeGaCloseSession) ProtoMessage()               {}
func (*CeGaCloseSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CeGaCloseSession) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

func init() {
	proto.RegisterType((*Server2AllSession)(nil), "smsg.Server2AllSession")
	proto.RegisterType((*GaCeReqLogin)(nil), "smsg.GaCeReqLogin")
	proto.RegisterType((*GaCeRespLogin)(nil), "smsg.GaCeRespLogin")
	proto.RegisterType((*CeGaBindGameServer)(nil), "smsg.CeGaBindGameServer")
	proto.RegisterType((*CeGamReqJoinGame)(nil), "smsg.CeGamReqJoinGame")
	proto.RegisterType((*CeGamRespJoinGame)(nil), "smsg.CeGamRespJoinGame")
	proto.RegisterType((*GamCeNoticeGameStart)(nil), "smsg.GamCeNoticeGameStart")
	proto.RegisterType((*GamCeNoticeGameEnd)(nil), "smsg.GamCeNoticeGameEnd")
	proto.RegisterType((*CeGameUserDisconnect)(nil), "smsg.CeGameUserDisconnect")
	proto.RegisterType((*CeGameUserReconnect)(nil), "smsg.CeGameUserReconnect")
	proto.RegisterType((*GaCeUserDisconnect)(nil), "smsg.GaCeUserDisconnect")
	proto.RegisterType((*CeGaCloseSession)(nil), "smsg.CeGaCloseSession")
	proto.RegisterEnum("smsg.GaCeRespLogin_Error", GaCeRespLogin_Error_name, GaCeRespLogin_Error_value)
	proto.RegisterEnum("smsg.CeGamRespJoinGame_Error", CeGamRespJoinGame_Error_name, CeGamRespJoinGame_Error_value)
}

func init() { proto.RegisterFile("msg/smsg/smsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0x4d, 0x36, 0xb4, 0x43, 0x40, 0xa9, 0x1b, 0x55, 0x01, 0x81, 0x14, 0xf9, 0x00, 0x91,
	0x40, 0x5b, 0x14, 0x6e, 0xdc, 0x60, 0x13, 0x45, 0x41, 0xa8, 0x07, 0x57, 0xfc, 0x00, 0xb3, 0x3b,
	0x8a, 0xac, 0x66, 0xed, 0xd4, 0x36, 0x15, 0x47, 0x0e, 0xfc, 0x08, 0x7e, 0x2e, 0xf2, 0x47, 0x92,
	0xdd, 0x40, 0xe0, 0xb2, 0xda, 0x19, 0xbd, 0x37, 0xf3, 0xde, 0xb3, 0x0d, 0x17, 0xb5, 0x5d, 0x5d,
	0xd9, 0xed, 0x27, 0xdf, 0x18, 0xed, 0x34, 0xed, 0xfa, 0x7f, 0xf6, 0x0a, 0xce, 0x6f, 0xd0, 0xdc,
	0xa3, 0x99, 0x7e, 0x58, 0xaf, 0x6f, 0xd0, 0x5a, 0xa9, 0x15, 0xa5, 0xd0, 0xad, 0x84, 0x13, 0x23,
	0x32, 0x26, 0x93, 0x3e, 0x0f, 0xff, 0xec, 0x3d, 0xf4, 0x17, 0xa2, 0x40, 0x8e, 0x77, 0x9f, 0xf5,
	0x4a, 0x2a, 0x3a, 0x84, 0xcc, 0xa2, 0x95, 0x55, 0x00, 0x65, 0x3c, 0x16, 0xbe, 0xeb, 0xf4, 0x2d,
	0xaa, 0xd1, 0xc9, 0x98, 0x4c, 0xce, 0x78, 0x2c, 0xd8, 0x2f, 0x02, 0x8f, 0x23, 0xd9, 0x6e, 0x22,
	0xfb, 0x12, 0x7a, 0xdf, 0x2c, 0x9a, 0xe5, 0x2c, 0xd0, 0xbb, 0x3c, 0x55, 0xf4, 0x35, 0x74, 0xd0,
	0x98, 0xc0, 0x7e, 0x32, 0x7d, 0x9a, 0x07, 0xb9, 0x2d, 0x66, 0x3e, 0x37, 0x46, 0x1b, 0xee, 0x51,
	0xfb, 0x65, 0x9d, 0xc6, 0x32, 0x3f, 0x5a, 0xaa, 0x85, 0xa8, 0x71, 0xd4, 0x1d, 0x93, 0xc9, 0x29,
	0x4f, 0x15, 0x1b, 0x42, 0x16, 0xb8, 0xf4, 0x11, 0x3c, 0x5c, 0xaa, 0x7b, 0xb1, 0x96, 0xd5, 0xe0,
	0x01, 0xbb, 0x06, 0x5a, 0xe0, 0x42, 0x7c, 0x94, 0xaa, 0xf2, 0xa8, 0x98, 0xc5, 0x11, 0x73, 0x0c,
	0xfa, 0x2b, 0x51, 0xa3, 0x0d, 0x18, 0x59, 0x05, 0x95, 0x19, 0x6f, 0xf5, 0xd8, 0x0f, 0x02, 0x03,
	0x3f, 0xb0, 0xe6, 0x78, 0xf7, 0x49, 0xc7, 0xd5, 0x5b, 0xb7, 0x69, 0x5e, 0x72, 0x1b, 0xd3, 0x8a,
	0x6b, 0x4e, 0x9a, 0x6b, 0x9e, 0xc1, 0xa9, 0x92, 0xe5, 0xad, 0xf2, 0x16, 0xa2, 0xb3, 0x5d, 0x1d,
	0x25, 0xb8, 0x24, 0x53, 0x56, 0xc1, 0x62, 0x90, 0xb0, 0xef, 0xb1, 0x9f, 0x04, 0xce, 0x93, 0x04,
	0xbb, 0xd9, 0x69, 0xb8, 0x8a, 0xc9, 0x92, 0x90, 0xec, 0x8b, 0x98, 0xec, 0x1f, 0xa8, 0x66, 0xba,
	0x97, 0xd0, 0xf3, 0xce, 0x92, 0xba, 0x0e, 0x4f, 0x15, 0x7b, 0xf9, 0xb7, 0x1c, 0xe9, 0xc0, 0x5f,
	0x8f, 0x1a, 0xaf, 0xb5, 0x9b, 0x7f, 0x97, 0xd6, 0x0d, 0x08, 0xcb, 0x61, 0xb8, 0x10, 0x75, 0xe1,
	0x5b, 0xb2, 0xc4, 0x10, 0xae, 0x13, 0xc6, 0x35, 0xe6, 0x92, 0xd6, 0xdc, 0x37, 0x40, 0x0f, 0xf0,
	0x73, 0x55, 0x1d, 0x45, 0xe7, 0x30, 0x0c, 0xea, 0xf1, 0x8b, 0x45, 0x33, 0x93, 0xb6, 0xd4, 0x4a,
	0x61, 0xe9, 0x8e, 0x45, 0xcd, 0x4a, 0xb8, 0xd8, 0xe3, 0x39, 0xfe, 0x07, 0x1e, 0xd7, 0x3a, 0x5c,
	0xce, 0xd2, 0xd1, 0xa4, 0x8a, 0x3e, 0x87, 0x33, 0x1b, 0x1f, 0xc9, 0x72, 0x16, 0x0e, 0x27, 0xe3,
	0xfb, 0x06, 0x9b, 0x7a, 0x0b, 0xc5, 0xa1, 0xa4, 0x16, 0x87, 0x1c, 0x72, 0xde, 0xc6, 0xfb, 0x52,
	0xac, 0xb5, 0xc5, 0xed, 0xfb, 0xfb, 0x27, 0xe3, 0x6b, 0x2f, 0xbc, 0xdf, 0x77, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xe8, 0xa1, 0x46, 0xff, 0xd6, 0x03, 0x00, 0x00,
}
