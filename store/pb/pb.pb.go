// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package pb

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

type TopicMap struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                *Topic   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicMap) Reset()         { *m = TopicMap{} }
func (m *TopicMap) String() string { return proto.CompactTextString(m) }
func (*TopicMap) ProtoMessage()    {}
func (*TopicMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{0}
}
func (m *TopicMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMap.Unmarshal(m, b)
}
func (m *TopicMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMap.Marshal(b, m, deterministic)
}
func (dst *TopicMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMap.Merge(dst, src)
}
func (m *TopicMap) XXX_Size() int {
	return xxx_messageInfo_TopicMap.Size(m)
}
func (m *TopicMap) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMap.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMap proto.InternalMessageInfo

func (m *TopicMap) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TopicMap) GetTopic() *Topic {
	if m != nil {
		return m.Topic
	}
	return nil
}

type Topic struct {
	Id                   uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TopicMapId           uint64        `protobuf:"varint,2,opt,name=topic_map_id,json=topicMapId,proto3" json:"topic_map_id,omitempty"`
	Names                []*Name       `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
	Occurrences          []*Occurrence `protobuf:"bytes,4,rep,name=occurrences,proto3" json:"occurrences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{1}
}
func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (dst *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(dst, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetTopicMapId() uint64 {
	if m != nil {
		return m.TopicMapId
	}
	return 0
}

func (m *Topic) GetNames() []*Name {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Topic) GetOccurrences() []*Occurrence {
	if m != nil {
		return m.Occurrences
	}
	return nil
}

type Name struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             uint64   `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{2}
}
func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (dst *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(dst, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Name) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Name) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Occurrence struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             uint64   `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Occurrence) Reset()         { *m = Occurrence{} }
func (m *Occurrence) String() string { return proto.CompactTextString(m) }
func (*Occurrence) ProtoMessage()    {}
func (*Occurrence) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{3}
}
func (m *Occurrence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Occurrence.Unmarshal(m, b)
}
func (m *Occurrence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Occurrence.Marshal(b, m, deterministic)
}
func (dst *Occurrence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Occurrence.Merge(dst, src)
}
func (m *Occurrence) XXX_Size() int {
	return xxx_messageInfo_Occurrence.Size(m)
}
func (m *Occurrence) XXX_DiscardUnknown() {
	xxx_messageInfo_Occurrence.DiscardUnknown(m)
}

var xxx_messageInfo_Occurrence proto.InternalMessageInfo

func (m *Occurrence) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Occurrence) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Occurrence) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetTopicMapsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicMapsRequest) Reset()         { *m = GetTopicMapsRequest{} }
func (m *GetTopicMapsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicMapsRequest) ProtoMessage()    {}
func (*GetTopicMapsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{4}
}
func (m *GetTopicMapsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicMapsRequest.Unmarshal(m, b)
}
func (m *GetTopicMapsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicMapsRequest.Marshal(b, m, deterministic)
}
func (dst *GetTopicMapsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicMapsRequest.Merge(dst, src)
}
func (m *GetTopicMapsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicMapsRequest.Size(m)
}
func (m *GetTopicMapsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicMapsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicMapsRequest proto.InternalMessageInfo

type GetTopicMapsResponse struct {
	TopicMaps            []*TopicMap `protobuf:"bytes,1,rep,name=topic_maps,json=topicMaps,proto3" json:"topic_maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetTopicMapsResponse) Reset()         { *m = GetTopicMapsResponse{} }
func (m *GetTopicMapsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopicMapsResponse) ProtoMessage()    {}
func (*GetTopicMapsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{5}
}
func (m *GetTopicMapsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicMapsResponse.Unmarshal(m, b)
}
func (m *GetTopicMapsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicMapsResponse.Marshal(b, m, deterministic)
}
func (dst *GetTopicMapsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicMapsResponse.Merge(dst, src)
}
func (m *GetTopicMapsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopicMapsResponse.Size(m)
}
func (m *GetTopicMapsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicMapsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicMapsResponse proto.InternalMessageInfo

func (m *GetTopicMapsResponse) GetTopicMaps() []*TopicMap {
	if m != nil {
		return m.TopicMaps
	}
	return nil
}

type CreateTopicMapRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicMapRequest) Reset()         { *m = CreateTopicMapRequest{} }
func (m *CreateTopicMapRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTopicMapRequest) ProtoMessage()    {}
func (*CreateTopicMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{6}
}
func (m *CreateTopicMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicMapRequest.Unmarshal(m, b)
}
func (m *CreateTopicMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicMapRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTopicMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicMapRequest.Merge(dst, src)
}
func (m *CreateTopicMapRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTopicMapRequest.Size(m)
}
func (m *CreateTopicMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicMapRequest proto.InternalMessageInfo

type CreateTopicMapResponse struct {
	TopicMap             *TopicMap `protobuf:"bytes,1,opt,name=topic_map,json=topicMap,proto3" json:"topic_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateTopicMapResponse) Reset()         { *m = CreateTopicMapResponse{} }
func (m *CreateTopicMapResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTopicMapResponse) ProtoMessage()    {}
func (*CreateTopicMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_6653ae2df246e174, []int{7}
}
func (m *CreateTopicMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicMapResponse.Unmarshal(m, b)
}
func (m *CreateTopicMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicMapResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTopicMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicMapResponse.Merge(dst, src)
}
func (m *CreateTopicMapResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTopicMapResponse.Size(m)
}
func (m *CreateTopicMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicMapResponse proto.InternalMessageInfo

func (m *CreateTopicMapResponse) GetTopicMap() *TopicMap {
	if m != nil {
		return m.TopicMap
	}
	return nil
}

func init() {
	proto.RegisterType((*TopicMap)(nil), "TopicMap")
	proto.RegisterType((*Topic)(nil), "Topic")
	proto.RegisterType((*Name)(nil), "Name")
	proto.RegisterType((*Occurrence)(nil), "Occurrence")
	proto.RegisterType((*GetTopicMapsRequest)(nil), "GetTopicMapsRequest")
	proto.RegisterType((*GetTopicMapsResponse)(nil), "GetTopicMapsResponse")
	proto.RegisterType((*CreateTopicMapRequest)(nil), "CreateTopicMapRequest")
	proto.RegisterType((*CreateTopicMapResponse)(nil), "CreateTopicMapResponse")
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_pb_6653ae2df246e174) }

var fileDescriptor_pb_6653ae2df246e174 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0x26, 0xa9, 0x92, 0x0b, 0x62, 0x30, 0x2d, 0x58, 0x2a, 0x43, 0xe4, 0x01, 0x65,
	0x21, 0x43, 0x59, 0x18, 0x91, 0x18, 0x50, 0x06, 0xa8, 0x64, 0xb1, 0x57, 0x6e, 0x72, 0x43, 0x24,
	0x92, 0x18, 0xdb, 0xe1, 0x1b, 0xf0, 0xbd, 0x51, 0x9d, 0xc6, 0xa5, 0x74, 0x65, 0xf4, 0x7b, 0xef,
	0x7e, 0xf7, 0x47, 0x86, 0x58, 0xed, 0x0a, 0xa5, 0x7b, 0xdb, 0xf3, 0x47, 0x88, 0xdf, 0x7b, 0xd5,
	0x54, 0xaf, 0x52, 0xd1, 0x4b, 0x98, 0x35, 0x35, 0x23, 0x19, 0xc9, 0x43, 0x31, 0x6b, 0x6a, 0x7a,
	0x0b, 0x91, 0xdd, 0x7b, 0x6c, 0x96, 0x91, 0x3c, 0x5d, 0xcf, 0x0b, 0x97, 0x14, 0xa3, 0xc8, 0xbf,
	0x09, 0x44, 0x4e, 0x38, 0xab, 0xcb, 0xe0, 0xc2, 0x45, 0xb6, 0xad, 0x54, 0xdb, 0xa6, 0x76, 0xe5,
	0xa1, 0x00, 0x7b, 0xe8, 0x53, 0xd6, 0x74, 0x05, 0x51, 0x27, 0x5b, 0x34, 0x2c, 0xc8, 0x82, 0x3c,
	0x5d, 0x47, 0xc5, 0x9b, 0x6c, 0x51, 0x8c, 0x1a, 0xbd, 0x87, 0xb4, 0xaf, 0xaa, 0x41, 0x6b, 0xec,
	0x2a, 0x34, 0x2c, 0x74, 0x91, 0xb4, 0xd8, 0x78, 0x4d, 0xfc, 0xf6, 0x79, 0x09, 0xe1, 0xbe, 0xfa,
	0x6c, 0x8a, 0x15, 0x24, 0x4a, 0x6a, 0xec, 0xec, 0x71, 0x84, 0x78, 0x14, 0xca, 0x9a, 0x2e, 0x20,
	0xfa, 0x92, 0x1f, 0x03, 0xb2, 0x20, 0x23, 0x79, 0x22, 0xc6, 0x07, 0xdf, 0x00, 0x1c, 0xbb, 0xfc,
	0x07, 0x70, 0x09, 0x57, 0x2f, 0x68, 0xa7, 0x03, 0x1b, 0x81, 0x9f, 0x03, 0x1a, 0xcb, 0x9f, 0x60,
	0x71, 0x2a, 0x1b, 0xd5, 0x77, 0x06, 0x69, 0x0e, 0xe0, 0x0f, 0x67, 0x18, 0x71, 0x8b, 0x27, 0xc5,
	0x94, 0x13, 0xc9, 0x74, 0x41, 0xc3, 0x6f, 0x60, 0xf9, 0xac, 0x51, 0x5a, 0xf4, 0xa6, 0x47, 0x5f,
	0xff, 0x35, 0x0e, 0xf0, 0x3b, 0x48, 0x3c, 0xdc, 0x6d, 0x75, 0xc2, 0x8e, 0x27, 0xf6, 0x6e, 0xee,
	0x3e, 0xc6, 0xc3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0xa8, 0xd8, 0x80, 0x24, 0x02, 0x00,
	0x00,
}
