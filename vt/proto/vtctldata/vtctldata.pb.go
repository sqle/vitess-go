// Code generated by protoc-gen-go.
// source: vtctldata.proto
// DO NOT EDIT!

/*
Package vtctldata is a generated protocol buffer package.

It is generated from these files:
	vtctldata.proto

It has these top-level messages:
	ExecuteVtctlCommandRequest
	ExecuteVtctlCommandResponse
*/
package vtctldata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import logutil "gopkg.in/sqle/vitess-go.v2/vt/proto/logutil"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ExecuteVtctlCommandRequest is the payload for ExecuteVtctlCommand.
// timeouts are in nanoseconds.
type ExecuteVtctlCommandRequest struct {
	Args          []string `protobuf:"bytes,1,rep,name=args" json:"args,omitempty"`
	ActionTimeout int64    `protobuf:"varint,2,opt,name=action_timeout,json=actionTimeout" json:"action_timeout,omitempty"`
}

func (m *ExecuteVtctlCommandRequest) Reset()                    { *m = ExecuteVtctlCommandRequest{} }
func (m *ExecuteVtctlCommandRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteVtctlCommandRequest) ProtoMessage()               {}
func (*ExecuteVtctlCommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExecuteVtctlCommandRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ExecuteVtctlCommandRequest) GetActionTimeout() int64 {
	if m != nil {
		return m.ActionTimeout
	}
	return 0
}

// ExecuteVtctlCommandResponse is streamed back by ExecuteVtctlCommand.
type ExecuteVtctlCommandResponse struct {
	Event *logutil.Event `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
}

func (m *ExecuteVtctlCommandResponse) Reset()                    { *m = ExecuteVtctlCommandResponse{} }
func (m *ExecuteVtctlCommandResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecuteVtctlCommandResponse) ProtoMessage()               {}
func (*ExecuteVtctlCommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecuteVtctlCommandResponse) GetEvent() *logutil.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteVtctlCommandRequest)(nil), "vtctldata.ExecuteVtctlCommandRequest")
	proto.RegisterType((*ExecuteVtctlCommandResponse)(nil), "vtctldata.ExecuteVtctlCommandResponse")
}

func init() { proto.RegisterFile("vtctldata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2b, 0x49, 0x2e,
	0xc9, 0x49, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x48,
	0xf1, 0xe6, 0xe4, 0xa7, 0x97, 0x96, 0x64, 0xe6, 0x40, 0x64, 0x94, 0xc2, 0xb9, 0xa4, 0x5c, 0x2b,
	0x52, 0x93, 0x4b, 0x4b, 0x52, 0xc3, 0x40, 0x4a, 0x9c, 0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x82,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x12, 0x8b, 0xd2, 0x8b, 0x25, 0x18,
	0x15, 0x98, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x55, 0x2e, 0xbe, 0xc4, 0xe4, 0x92, 0xcc, 0xfc,
	0xbc, 0xf8, 0x92, 0xcc, 0xdc, 0xd4, 0xfc, 0xd2, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20,
	0x5e, 0x88, 0x68, 0x08, 0x44, 0x50, 0xc9, 0x99, 0x4b, 0x1a, 0xab, 0xc1, 0xc5, 0x05, 0xf9, 0x79,
	0xc5, 0xa9, 0x42, 0x2a, 0x5c, 0xac, 0xa9, 0x65, 0xa9, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x7c, 0x7a, 0x30, 0x67, 0xb9, 0x82, 0x44, 0x83, 0x20, 0x92, 0x49, 0x6c, 0x60, 0x47,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x9e, 0x0a, 0x31, 0xd1, 0x00, 0x00, 0x00,
}
