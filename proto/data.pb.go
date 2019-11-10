// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Message struct {
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	From                 string   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Message) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Void)(nil), "proto.Void")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0x31, 0x8f, 0xc2, 0x30,
	0x0c, 0x85, 0x9b, 0x5e, 0xae, 0xbd, 0xf3, 0x49, 0x95, 0xce, 0x53, 0xd4, 0x09, 0x65, 0xaa, 0x18,
	0x2a, 0x04, 0xfc, 0x03, 0x06, 0x26, 0x16, 0x24, 0xd8, 0x43, 0xeb, 0x52, 0x86, 0x26, 0x28, 0x31,
	0xff, 0x1f, 0x35, 0xb4, 0x0c, 0x4c, 0xf6, 0xf7, 0x6c, 0xbd, 0xf7, 0x00, 0x5a, 0xc3, 0xa6, 0xbe,
	0x7b, 0xc7, 0x0e, 0xbf, 0xe3, 0xd0, 0x7b, 0xc8, 0x0f, 0x14, 0x82, 0xb9, 0x12, 0x2a, 0xc8, 0x1b,
	0x67, 0x99, 0x2c, 0xab, 0x74, 0x21, 0xaa, 0xdf, 0xe3, 0x8c, 0x88, 0x20, 0x3b, 0xef, 0x06, 0xf5,
	0x15, 0xe5, 0xb8, 0x63, 0x01, 0x29, 0x3b, 0x25, 0xa3, 0x92, 0xb2, 0xd3, 0x1a, 0xe4, 0x29, 0x90,
	0xc7, 0x12, 0x7e, 0x1e, 0x81, 0xbc, 0x35, 0x03, 0x29, 0x11, 0xaf, 0x6f, 0xd6, 0x19, 0xc8, 0xb3,
	0xbb, 0xb5, 0xeb, 0x1e, 0xe4, 0xae, 0x37, 0x8c, 0xcb, 0x98, 0x68, 0xa9, 0x61, 0xfc, 0x7b, 0xd5,
	0xaa, 0x47, 0x8f, 0xb2, 0x98, 0x60, 0x6a, 0xa6, 0x93, 0x95, 0xc0, 0x2d, 0xfc, 0xb3, 0x37, 0x36,
	0x74, 0xe4, 0x79, 0xd2, 0x03, 0x7e, 0x3c, 0x96, 0xb3, 0xcb, 0x98, 0xa2, 0x93, 0x4a, 0x5c, 0xb2,
	0xc8, 0x9b, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xe3, 0x30, 0x0b, 0xfa, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (Chat_ConnectClient, error)
	TransfertMessages(ctx context.Context, opts ...grpc.CallOption) (Chat_TransfertMessagesClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (Chat_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/proto.Chat/connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ConnectClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatConnectClient struct {
	grpc.ClientStream
}

func (x *chatConnectClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) TransfertMessages(ctx context.Context, opts ...grpc.CallOption) (Chat_TransfertMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[1], "/proto.Chat/transfertMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatTransfertMessagesClient{stream}
	return x, nil
}

type Chat_TransfertMessagesClient interface {
	Send(*Message) error
	CloseAndRecv() (*Void, error)
	grpc.ClientStream
}

type chatTransfertMessagesClient struct {
	grpc.ClientStream
}

func (x *chatTransfertMessagesClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatTransfertMessagesClient) CloseAndRecv() (*Void, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Void)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Connect(*User, Chat_ConnectServer) error
	TransfertMessages(Chat_TransfertMessagesServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) Connect(req *User, srv Chat_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedChatServer) TransfertMessages(srv Chat_TransfertMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method TransfertMessages not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).Connect(m, &chatConnectServer{stream})
}

type Chat_ConnectServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatConnectServer struct {
	grpc.ServerStream
}

func (x *chatConnectServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_TransfertMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).TransfertMessages(&chatTransfertMessagesServer{stream})
}

type Chat_TransfertMessagesServer interface {
	SendAndClose(*Void) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatTransfertMessagesServer struct {
	grpc.ServerStream
}

func (x *chatTransfertMessagesServer) SendAndClose(m *Void) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatTransfertMessagesServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "connect",
			Handler:       _Chat_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "transfertMessages",
			Handler:       _Chat_TransfertMessages_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "data.proto",
}
