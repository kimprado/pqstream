// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pqstream.proto

/*
Package pqs is a generated protocol buffer package.

It is generated from these files:
	pqstream.proto

It has these top-level messages:
	ListenRequest
	Event
*/
package pqs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

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

// An operation in the database.
type Operation int32

const (
	Operation_UNKNOWN  Operation = 0
	Operation_INSERT   Operation = 1
	Operation_UPDATE   Operation = 2
	Operation_DELETE   Operation = 3
	Operation_TRUNCATE Operation = 4
)

var Operation_name = map[int32]string{
	0: "UNKNOWN",
	1: "INSERT",
	2: "UPDATE",
	3: "DELETE",
	4: "TRUNCATE",
}
var Operation_value = map[string]int32{
	"UNKNOWN":  0,
	"INSERT":   1,
	"UPDATE":   2,
	"DELETE":   3,
	"TRUNCATE": 4,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A request to listen to database event streams.
type ListenRequest struct {
	// if provided, this string will be used to match table names to track.
	TableRegexp string `protobuf:"bytes,1,opt,name=table_regexp,json=tableRegexp" json:"table_regexp,omitempty"`
}

func (m *ListenRequest) Reset()                    { *m = ListenRequest{} }
func (m *ListenRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()               {}
func (*ListenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListenRequest) GetTableRegexp() string {
	if m != nil {
		return m.TableRegexp
	}
	return ""
}

// A database event.
type Event struct {
	Schema string    `protobuf:"bytes,1,opt,name=schema" json:"schema,omitempty"`
	Table  string    `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
	Op     Operation `protobuf:"varint,3,opt,name=op,enum=pqs.Operation" json:"op,omitempty"`
	// if the id column exists, this will populate it
	Id string `protobuf:"bytes,4,opt,name=id" json:"id,omitempty"`
	// payload is a json encoded representation of the changed object.
	Payload *google_protobuf.Struct `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Event) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *Event) GetOp() Operation {
	if m != nil {
		return m.Op
	}
	return Operation_UNKNOWN
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetPayload() *google_protobuf.Struct {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*ListenRequest)(nil), "pqs.ListenRequest")
	proto.RegisterType((*Event)(nil), "pqs.Event")
	proto.RegisterEnum("pqs.Operation", Operation_name, Operation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PQStream service

type PQStreamClient interface {
	// Listen responds with a stream of database operations.
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (PQStream_ListenClient, error)
}

type pQStreamClient struct {
	cc *grpc.ClientConn
}

func NewPQStreamClient(cc *grpc.ClientConn) PQStreamClient {
	return &pQStreamClient{cc}
}

func (c *pQStreamClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (PQStream_ListenClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PQStream_serviceDesc.Streams[0], c.cc, "/pqs.PQStream/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &pQStreamListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PQStream_ListenClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type pQStreamListenClient struct {
	grpc.ClientStream
}

func (x *pQStreamListenClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PQStream service

type PQStreamServer interface {
	// Listen responds with a stream of database operations.
	Listen(*ListenRequest, PQStream_ListenServer) error
}

func RegisterPQStreamServer(s *grpc.Server, srv PQStreamServer) {
	s.RegisterService(&_PQStream_serviceDesc, srv)
}

func _PQStream_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PQStreamServer).Listen(m, &pQStreamListenServer{stream})
}

type PQStream_ListenServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type pQStreamListenServer struct {
	grpc.ServerStream
}

func (x *pQStreamListenServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _PQStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pqs.PQStream",
	HandlerType: (*PQStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _PQStream_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pqstream.proto",
}

func init() { proto.RegisterFile("pqstream.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4f, 0xc2, 0x30,
	0x14, 0xc6, 0xd9, 0x80, 0x01, 0x0f, 0x5c, 0x96, 0xc6, 0xe8, 0xc2, 0xc1, 0x20, 0x27, 0x62, 0x4c,
	0xa7, 0x33, 0x26, 0x5e, 0x8d, 0xec, 0xa0, 0x92, 0x81, 0x65, 0xc4, 0xa3, 0xe9, 0xa0, 0x8e, 0x25,
	0x63, 0xed, 0xd6, 0xce, 0xc8, 0x7f, 0xe2, 0x9f, 0x6b, 0x28, 0x60, 0xe2, 0xa9, 0xef, 0x7d, 0xef,
	0x7d, 0xcd, 0xfb, 0x7e, 0x60, 0x8b, 0x42, 0xaa, 0x92, 0xd1, 0x0d, 0x16, 0x25, 0x57, 0x1c, 0xd5,
	0x45, 0x21, 0xfb, 0xf7, 0x49, 0xaa, 0xd6, 0x55, 0x8c, 0x97, 0x7c, 0xe3, 0x25, 0x3c, 0xa3, 0x79,
	0xe2, 0xe9, 0x69, 0x5c, 0x7d, 0x7a, 0x42, 0x6d, 0x05, 0x93, 0x9e, 0x54, 0x65, 0xb5, 0x54, 0x87,
	0x67, 0xef, 0x1d, 0xfa, 0x70, 0x32, 0x49, 0xa5, 0x62, 0x39, 0x61, 0x45, 0xc5, 0xa4, 0x42, 0x97,
	0xd0, 0x53, 0x34, 0xce, 0xd8, 0x47, 0xc9, 0x12, 0xf6, 0x2d, 0x5c, 0x63, 0x60, 0x8c, 0x3a, 0xa4,
	0xab, 0x35, 0xa2, 0xa5, 0xe1, 0x8f, 0x01, 0xcd, 0xe0, 0x8b, 0xe5, 0x0a, 0x9d, 0x81, 0x25, 0x97,
	0x6b, 0xb6, 0xa1, 0x87, 0xb5, 0x43, 0x87, 0x4e, 0xa1, 0xa9, 0x0d, 0xae, 0xa9, 0xe5, 0x7d, 0x83,
	0x2e, 0xc0, 0xe4, 0xc2, 0xad, 0x0f, 0x8c, 0x91, 0xed, 0xdb, 0x58, 0x14, 0x12, 0x4f, 0x05, 0x2b,
	0xa9, 0x4a, 0x79, 0x4e, 0x4c, 0x2e, 0x90, 0x0d, 0x66, 0xba, 0x72, 0x1b, 0xda, 0x62, 0xa6, 0x2b,
	0x74, 0x0b, 0x2d, 0x41, 0xb7, 0x19, 0xa7, 0x2b, 0xb7, 0x39, 0x30, 0x46, 0x5d, 0xff, 0x1c, 0x27,
	0x9c, 0x27, 0x19, 0xc3, 0xc7, 0x64, 0x78, 0xae, 0xb3, 0x90, 0xe3, 0xde, 0xd5, 0x0b, 0x74, 0xfe,
	0xfe, 0x44, 0x5d, 0x68, 0x2d, 0xc2, 0xd7, 0x70, 0xfa, 0x1e, 0x3a, 0x35, 0x04, 0x60, 0x3d, 0x87,
	0xf3, 0x80, 0x44, 0x8e, 0xb1, 0xab, 0x17, 0xb3, 0xf1, 0x63, 0x14, 0x38, 0xe6, 0xae, 0x1e, 0x07,
	0x93, 0x20, 0x0a, 0x9c, 0x3a, 0xea, 0x41, 0x3b, 0x22, 0x8b, 0xf0, 0x69, 0x37, 0x69, 0xf8, 0x0f,
	0xd0, 0x9e, 0xbd, 0xcd, 0x35, 0x68, 0x74, 0x0d, 0xd6, 0x1e, 0x13, 0x42, 0xfa, 0xf0, 0x7f, 0xcc,
	0xfa, 0xa0, 0x35, 0x8d, 0x64, 0x58, 0xbb, 0x31, 0x62, 0x4b, 0xdf, 0x77, 0xf7, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x97, 0xf9, 0xd4, 0xa9, 0x01, 0x00, 0x00,
}
