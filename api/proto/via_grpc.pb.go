// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: api/proto/via.proto

package viapb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Via_GetDangerProbability_FullMethodName = "/viapb.Via/GetDangerProbability"
	Via_GetIncident_FullMethodName          = "/viapb.Via/GetIncident"
)

// ViaClient is the client API for Via service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViaClient interface {
	GetDangerProbability(ctx context.Context, in *GetDangerProbabilityRequest, opts ...grpc.CallOption) (*GetDangerProbabilityResponse, error)
	GetIncident(ctx context.Context, in *GetIncidentRequest, opts ...grpc.CallOption) (*GetIncidentResponse, error)
}

type viaClient struct {
	cc grpc.ClientConnInterface
}

func NewViaClient(cc grpc.ClientConnInterface) ViaClient {
	return &viaClient{cc}
}

func (c *viaClient) GetDangerProbability(ctx context.Context, in *GetDangerProbabilityRequest, opts ...grpc.CallOption) (*GetDangerProbabilityResponse, error) {
	out := new(GetDangerProbabilityResponse)
	err := c.cc.Invoke(ctx, Via_GetDangerProbability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viaClient) GetIncident(ctx context.Context, in *GetIncidentRequest, opts ...grpc.CallOption) (*GetIncidentResponse, error) {
	out := new(GetIncidentResponse)
	err := c.cc.Invoke(ctx, Via_GetIncident_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViaServer is the server API for Via service.
// All implementations must embed UnimplementedViaServer
// for forward compatibility
type ViaServer interface {
	GetDangerProbability(context.Context, *GetDangerProbabilityRequest) (*GetDangerProbabilityResponse, error)
	GetIncident(context.Context, *GetIncidentRequest) (*GetIncidentResponse, error)
	mustEmbedUnimplementedViaServer()
}

// UnimplementedViaServer must be embedded to have forward compatible implementations.
type UnimplementedViaServer struct {
}

func (UnimplementedViaServer) GetDangerProbability(context.Context, *GetDangerProbabilityRequest) (*GetDangerProbabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDangerProbability not implemented")
}
func (UnimplementedViaServer) GetIncident(context.Context, *GetIncidentRequest) (*GetIncidentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncident not implemented")
}
func (UnimplementedViaServer) mustEmbedUnimplementedViaServer() {}

// UnsafeViaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViaServer will
// result in compilation errors.
type UnsafeViaServer interface {
	mustEmbedUnimplementedViaServer()
}

func RegisterViaServer(s grpc.ServiceRegistrar, srv ViaServer) {
	s.RegisterService(&Via_ServiceDesc, srv)
}

func _Via_GetDangerProbability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDangerProbabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViaServer).GetDangerProbability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Via_GetDangerProbability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViaServer).GetDangerProbability(ctx, req.(*GetDangerProbabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Via_GetIncident_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncidentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViaServer).GetIncident(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Via_GetIncident_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViaServer).GetIncident(ctx, req.(*GetIncidentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Via_ServiceDesc is the grpc.ServiceDesc for Via service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Via_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viapb.Via",
	HandlerType: (*ViaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDangerProbability",
			Handler:    _Via_GetDangerProbability_Handler,
		},
		{
			MethodName: "GetIncident",
			Handler:    _Via_GetIncident_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/via.proto",
}
