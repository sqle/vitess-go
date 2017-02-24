// Code generated by protoc-gen-go.
// source: automationservice.proto
// DO NOT EDIT!

/*
Package automationservice is a generated protocol buffer package.

It is generated from these files:
	automationservice.proto

It has these top-level messages:
*/
package automationservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import automation "gopkg.in/sqle/vitess-go.v1/vt/proto/automation"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Automation service

type AutomationClient interface {
	// Start a cluster operation.
	EnqueueClusterOperation(ctx context.Context, in *automation.EnqueueClusterOperationRequest, opts ...grpc.CallOption) (*automation.EnqueueClusterOperationResponse, error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(ctx context.Context, in *automation.GetClusterOperationDetailsRequest, opts ...grpc.CallOption) (*automation.GetClusterOperationDetailsResponse, error)
}

type automationClient struct {
	cc *grpc.ClientConn
}

func NewAutomationClient(cc *grpc.ClientConn) AutomationClient {
	return &automationClient{cc}
}

func (c *automationClient) EnqueueClusterOperation(ctx context.Context, in *automation.EnqueueClusterOperationRequest, opts ...grpc.CallOption) (*automation.EnqueueClusterOperationResponse, error) {
	out := new(automation.EnqueueClusterOperationResponse)
	err := grpc.Invoke(ctx, "/automationservice.Automation/EnqueueClusterOperation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationClient) GetClusterOperationDetails(ctx context.Context, in *automation.GetClusterOperationDetailsRequest, opts ...grpc.CallOption) (*automation.GetClusterOperationDetailsResponse, error) {
	out := new(automation.GetClusterOperationDetailsResponse)
	err := grpc.Invoke(ctx, "/automationservice.Automation/GetClusterOperationDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Automation service

type AutomationServer interface {
	// Start a cluster operation.
	EnqueueClusterOperation(context.Context, *automation.EnqueueClusterOperationRequest) (*automation.EnqueueClusterOperationResponse, error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(context.Context, *automation.GetClusterOperationDetailsRequest) (*automation.GetClusterOperationDetailsResponse, error)
}

func RegisterAutomationServer(s *grpc.Server, srv AutomationServer) {
	s.RegisterService(&_Automation_serviceDesc, srv)
}

func _Automation_EnqueueClusterOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(automation.EnqueueClusterOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServer).EnqueueClusterOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automationservice.Automation/EnqueueClusterOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServer).EnqueueClusterOperation(ctx, req.(*automation.EnqueueClusterOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Automation_GetClusterOperationDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(automation.GetClusterOperationDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServer).GetClusterOperationDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automationservice.Automation/GetClusterOperationDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServer).GetClusterOperationDetails(ctx, req.(*automation.GetClusterOperationDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Automation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "automationservice.Automation",
	HandlerType: (*AutomationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnqueueClusterOperation",
			Handler:    _Automation_EnqueueClusterOperation_Handler,
		},
		{
			MethodName: "GetClusterOperationDetails",
			Handler:    _Automation_GetClusterOperationDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automationservice.proto",
}

func init() { proto.RegisterFile("automationservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2c, 0x2d, 0xc9,
	0xcf, 0x4d, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc4, 0x90, 0x90, 0x12, 0x40, 0x08, 0x41, 0x14, 0x19, 0x35, 0x32,
	0x71, 0x71, 0x39, 0xc2, 0x05, 0x85, 0x4a, 0xb8, 0xc4, 0x5d, 0xf3, 0x0a, 0x4b, 0x53, 0x4b, 0x53,
	0x9d, 0x73, 0x4a, 0x8b, 0x4b, 0x52, 0x8b, 0xfc, 0x0b, 0x52, 0x8b, 0x20, 0x52, 0x5a, 0x7a, 0x48,
	0x9a, 0x71, 0x28, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0xd2, 0x26, 0x4a, 0x6d, 0x71,
	0x01, 0xc8, 0x65, 0x4a, 0x0c, 0x42, 0xb5, 0x5c, 0x52, 0xee, 0xa9, 0x25, 0xe8, 0x0a, 0x5c, 0x52,
	0x4b, 0x12, 0x33, 0x73, 0x8a, 0x85, 0x74, 0x91, 0x0d, 0xc3, 0xad, 0x0e, 0x66, 0xb7, 0x1e, 0xb1,
	0xca, 0x61, 0xd6, 0x27, 0xb1, 0x81, 0x83, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x3e,
	0xd2, 0x25, 0x4a, 0x01, 0x00, 0x00,
}
