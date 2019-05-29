// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/loading_todo_response.proto

package message // import "github.com/guilhermesteves/go-todo-api/internal/pkg/core/message"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/guilhermesteves/go-todo-api/internal/pkg/core/model"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoadingToDoResponse struct {
	Todo                 *model.ToDo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LoadingToDoResponse) Reset()         { *m = LoadingToDoResponse{} }
func (m *LoadingToDoResponse) String() string { return proto.CompactTextString(m) }
func (*LoadingToDoResponse) ProtoMessage()    {}
func (*LoadingToDoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_loading_todo_response_64527161496ad0d2, []int{0}
}
func (m *LoadingToDoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadingToDoResponse.Unmarshal(m, b)
}
func (m *LoadingToDoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadingToDoResponse.Marshal(b, m, deterministic)
}
func (dst *LoadingToDoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadingToDoResponse.Merge(dst, src)
}
func (m *LoadingToDoResponse) XXX_Size() int {
	return xxx_messageInfo_LoadingToDoResponse.Size(m)
}
func (m *LoadingToDoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadingToDoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadingToDoResponse proto.InternalMessageInfo

func (m *LoadingToDoResponse) GetTodo() *model.ToDo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadingToDoResponse)(nil), "LoadingToDoResponse")
}

func init() {
	proto.RegisterFile("message/loading_todo_response.proto", fileDescriptor_loading_todo_response_64527161496ad0d2)
}

var fileDescriptor_loading_todo_response_64527161496ad0d2 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0xbd, 0x0a, 0xc2, 0x30,
	0x14, 0x46, 0x29, 0xa8, 0x43, 0x5d, 0x44, 0x17, 0x75, 0x12, 0x5d, 0x5c, 0xda, 0x2b, 0xfa, 0x02,
	0x52, 0x1c, 0x9d, 0xc4, 0xc9, 0xa5, 0xa4, 0xed, 0x25, 0x0d, 0x26, 0xf9, 0x42, 0x92, 0xfa, 0xfc,
	0xd2, 0x9f, 0xf5, 0x9c, 0x73, 0xb9, 0x5f, 0x7a, 0x32, 0x1c, 0x82, 0x90, 0x4c, 0x1a, 0xa2, 0x51,
	0x56, 0x96, 0x11, 0x0d, 0x4a, 0xcf, 0xc1, 0xc1, 0x06, 0xce, 0x9d, 0x47, 0xc4, 0x7e, 0x65, 0xd0,
	0xb0, 0xa6, 0x5e, 0x8d, 0xe4, 0x78, 0x49, 0x37, 0xcf, 0xf1, 0xe0, 0x8d, 0x07, 0x5e, 0x53, 0xbe,
	0xde, 0xa5, 0xb3, 0x3e, 0xda, 0x26, 0x87, 0xe4, 0xbc, 0xbc, 0xce, 0xf3, 0x41, 0x0e, 0xa8, 0x28,
	0x3e, 0x77, 0xa9, 0x62, 0xdb, 0x55, 0x79, 0x0d, 0x43, 0xb2, 0x53, 0xba, 0x65, 0x6f, 0x38, 0x44,
	0xfe, 0x71, 0x20, 0x89, 0xac, 0xaf, 0x32, 0xe1, 0x14, 0x29, 0x1b, 0xd9, 0x5b, 0xa1, 0xc9, 0x7d,
	0x25, 0xd5, 0xf0, 0x4c, 0xd3, 0xc6, 0x6a, 0x31, 0x3c, 0xbf, 0xfd, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x08, 0xd3, 0x33, 0xfb, 0xb5, 0x00, 0x00, 0x00,
}
