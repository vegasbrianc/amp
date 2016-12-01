// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/api/rpc/project/project.proto
// DO NOT EDIT!

/*
Package project is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/api/rpc/project/project.proto

It has these top-level messages:
	CreateRequest
	CreateReply
*/
package project

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type CreateRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *CreateReply) Reset()                    { *m = CreateReply{} }
func (m *CreateReply) String() string            { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()               {}
func (*CreateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "project.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "project.CreateReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Project service

type ProjectClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
}

type projectClient struct {
	cc *grpc.ClientConn
}

func NewProjectClient(cc *grpc.ClientConn) ProjectClient {
	return &projectClient{cc}
}

func (c *projectClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := grpc.Invoke(ctx, "/project.Project/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Project service

type ProjectServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
}

func RegisterProjectServer(s *grpc.Server, srv ProjectServer) {
	s.RegisterService(&_Project_serviceDesc, srv)
}

func _Project_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.Project/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Project_serviceDesc = grpc.ServiceDesc{
	ServiceName: "project.Project",
	HandlerType: (*ProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Project_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/appcelerator/amp/api/rpc/project/project.proto",
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/api/rpc/project/project.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x8e, 0xb1, 0x8a, 0x83, 0x40,
	0x10, 0x86, 0x4f, 0x39, 0x94, 0x9b, 0xe3, 0xae, 0x58, 0x42, 0x90, 0x54, 0xc1, 0x26, 0xa9, 0x5c,
	0x88, 0x4d, 0x52, 0xfb, 0x02, 0xc1, 0x37, 0x18, 0xd7, 0xc1, 0x6c, 0x70, 0xb3, 0x93, 0x75, 0x2d,
	0x7c, 0xfb, 0x80, 0xba, 0x45, 0x48, 0x35, 0xf3, 0xff, 0xf0, 0xf1, 0x7f, 0x70, 0xe9, 0xb4, 0xbf,
	0x8d, 0x4d, 0xa1, 0xac, 0x91, 0xc8, 0xac, 0xa8, 0x27, 0x87, 0xde, 0x3a, 0x89, 0x86, 0x25, 0xb2,
	0x96, 0x8e, 0x95, 0x64, 0x67, 0xef, 0xa4, 0x7c, 0xb8, 0x05, 0x3b, 0xeb, 0xad, 0x48, 0xd7, 0x98,
	0x97, 0xf0, 0x57, 0x39, 0x42, 0x4f, 0x35, 0x3d, 0x47, 0x1a, 0xbc, 0xf8, 0x87, 0x58, 0xb7, 0x59,
	0xb4, 0x8f, 0x8e, 0x3f, 0x75, 0xac, 0x5b, 0x21, 0xe0, 0xfb, 0x81, 0x86, 0xb2, 0x78, 0x6e, 0xe6,
	0x3f, 0x3f, 0xc0, 0x6f, 0x80, 0xb8, 0x9f, 0x44, 0x06, 0xa9, 0xa1, 0x61, 0xc0, 0x8e, 0x56, 0x2e,
	0xc4, 0x53, 0x05, 0xe9, 0x75, 0x19, 0x12, 0x67, 0x48, 0x16, 0x46, 0x6c, 0x8b, 0xe0, 0xf2, 0xb6,
	0xbc, 0xdb, 0x7c, 0xf4, 0xdc, 0x4f, 0xf9, 0x57, 0x93, 0xcc, 0xca, 0xe5, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x55, 0x55, 0x88, 0xef, 0x00, 0x00, 0x00,
}
