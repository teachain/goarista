// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: gnmireverse.proto

// Package gNMIReverse defines a service specification that reverses the
// direction of the dial for gNMI Subscribe and gNMI Get.
//
// gNMI is a "dial-in" service, where the telemetry collector must make a
// connection to the gNMI target. With gNMIReverse, a client can be run
// alongside the gNMI target and then "dial-out" to a gNMIReverse server
// to send streaming data.

package gnmireverse

import (
	context "context"
	gnmi "github.com/openconfig/gnmi/proto/gnmi"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GNMIReverse_Publish_FullMethodName    = "/gnmireverse.gNMIReverse/Publish"
	GNMIReverse_PublishGet_FullMethodName = "/gnmireverse.gNMIReverse/PublishGet"
)

// GNMIReverseClient is the client API for GNMIReverse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GNMIReverseClient interface {
	// Publish allows the client to publish gNMI SubscribeResponses to the
	// collector server. The client is typically run alongside the gNMI target
	// and forwards SubscribeResponses from the target to the collector server.
	// The request is specified by the client.
	Publish(ctx context.Context, opts ...grpc.CallOption) (GNMIReverse_PublishClient, error)
	// PublishGet allows the client to publish gNMI GetResponses to the
	// collector server. The client is typically run alongside the gNMI target
	// and forwards GetResponses from the target to the collector server.
	// The request and sample interval are specified by the client.
	PublishGet(ctx context.Context, opts ...grpc.CallOption) (GNMIReverse_PublishGetClient, error)
}

type gNMIReverseClient struct {
	cc grpc.ClientConnInterface
}

func NewGNMIReverseClient(cc grpc.ClientConnInterface) GNMIReverseClient {
	return &gNMIReverseClient{cc}
}

func (c *gNMIReverseClient) Publish(ctx context.Context, opts ...grpc.CallOption) (GNMIReverse_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &GNMIReverse_ServiceDesc.Streams[0], GNMIReverse_Publish_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gNMIReversePublishClient{stream}
	return x, nil
}

type GNMIReverse_PublishClient interface {
	Send(*gnmi.SubscribeResponse) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type gNMIReversePublishClient struct {
	grpc.ClientStream
}

func (x *gNMIReversePublishClient) Send(m *gnmi.SubscribeResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gNMIReversePublishClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gNMIReverseClient) PublishGet(ctx context.Context, opts ...grpc.CallOption) (GNMIReverse_PublishGetClient, error) {
	stream, err := c.cc.NewStream(ctx, &GNMIReverse_ServiceDesc.Streams[1], GNMIReverse_PublishGet_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gNMIReversePublishGetClient{stream}
	return x, nil
}

type GNMIReverse_PublishGetClient interface {
	Send(*gnmi.GetResponse) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type gNMIReversePublishGetClient struct {
	grpc.ClientStream
}

func (x *gNMIReversePublishGetClient) Send(m *gnmi.GetResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gNMIReversePublishGetClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GNMIReverseServer is the server API for GNMIReverse service.
// All implementations must embed UnimplementedGNMIReverseServer
// for forward compatibility
type GNMIReverseServer interface {
	// Publish allows the client to publish gNMI SubscribeResponses to the
	// collector server. The client is typically run alongside the gNMI target
	// and forwards SubscribeResponses from the target to the collector server.
	// The request is specified by the client.
	Publish(GNMIReverse_PublishServer) error
	// PublishGet allows the client to publish gNMI GetResponses to the
	// collector server. The client is typically run alongside the gNMI target
	// and forwards GetResponses from the target to the collector server.
	// The request and sample interval are specified by the client.
	PublishGet(GNMIReverse_PublishGetServer) error
	mustEmbedUnimplementedGNMIReverseServer()
}

// UnimplementedGNMIReverseServer must be embedded to have forward compatible implementations.
type UnimplementedGNMIReverseServer struct {
}

func (UnimplementedGNMIReverseServer) Publish(GNMIReverse_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedGNMIReverseServer) PublishGet(GNMIReverse_PublishGetServer) error {
	return status.Errorf(codes.Unimplemented, "method PublishGet not implemented")
}
func (UnimplementedGNMIReverseServer) mustEmbedUnimplementedGNMIReverseServer() {}

// UnsafeGNMIReverseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GNMIReverseServer will
// result in compilation errors.
type UnsafeGNMIReverseServer interface {
	mustEmbedUnimplementedGNMIReverseServer()
}

func RegisterGNMIReverseServer(s grpc.ServiceRegistrar, srv GNMIReverseServer) {
	s.RegisterService(&GNMIReverse_ServiceDesc, srv)
}

func _GNMIReverse_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GNMIReverseServer).Publish(&gNMIReversePublishServer{stream})
}

type GNMIReverse_PublishServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*gnmi.SubscribeResponse, error)
	grpc.ServerStream
}

type gNMIReversePublishServer struct {
	grpc.ServerStream
}

func (x *gNMIReversePublishServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gNMIReversePublishServer) Recv() (*gnmi.SubscribeResponse, error) {
	m := new(gnmi.SubscribeResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GNMIReverse_PublishGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GNMIReverseServer).PublishGet(&gNMIReversePublishGetServer{stream})
}

type GNMIReverse_PublishGetServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*gnmi.GetResponse, error)
	grpc.ServerStream
}

type gNMIReversePublishGetServer struct {
	grpc.ServerStream
}

func (x *gNMIReversePublishGetServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gNMIReversePublishGetServer) Recv() (*gnmi.GetResponse, error) {
	m := new(gnmi.GetResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GNMIReverse_ServiceDesc is the grpc.ServiceDesc for GNMIReverse service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GNMIReverse_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnmireverse.gNMIReverse",
	HandlerType: (*GNMIReverseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _GNMIReverse_Publish_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PublishGet",
			Handler:       _GNMIReverse_PublishGet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "gnmireverse.proto",
}
