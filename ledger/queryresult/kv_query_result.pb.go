// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/queryresult/kv_query_result.proto

package queryresult

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// KV -- QueryResult for range/execute query. Holds a key and corresponding value.
type KV struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ee2fe66594a8f2, []int{0}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// KeyModification -- QueryResult for history query. Holds a transaction ID, value,
// timestamp, and delete marker which resulted from a history query.
type KeyModification struct {
	TxId                 string               `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Value                []byte               `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsDelete             bool                 `protobuf:"varint,4,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeyModification) Reset()         { *m = KeyModification{} }
func (m *KeyModification) String() string { return proto.CompactTextString(m) }
func (*KeyModification) ProtoMessage()    {}
func (*KeyModification) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8ee2fe66594a8f2, []int{1}
}

func (m *KeyModification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyModification.Unmarshal(m, b)
}
func (m *KeyModification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyModification.Marshal(b, m, deterministic)
}
func (m *KeyModification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyModification.Merge(m, src)
}
func (m *KeyModification) XXX_Size() int {
	return xxx_messageInfo_KeyModification.Size(m)
}
func (m *KeyModification) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyModification.DiscardUnknown(m)
}

var xxx_messageInfo_KeyModification proto.InternalMessageInfo

func (m *KeyModification) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *KeyModification) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyModification) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *KeyModification) GetIsDelete() bool {
	if m != nil {
		return m.IsDelete
	}
	return false
}

func init() {
	proto.RegisterType((*KV)(nil), "sdk.queryresult.KV")
	proto.RegisterType((*KeyModification)(nil), "sdk.queryresult.KeyModification")
}

func init() {
	proto.RegisterFile("ledger/queryresult/kv_query_result.proto", fileDescriptor_f8ee2fe66594a8f2)
}

var fileDescriptor_f8ee2fe66594a8f2 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0xfe, 0xc8, 0x9a, 0x09, 0x4a, 0xf4, 0x50, 0xa6, 0x60, 0xd9, 0xa9, 0x97, 0x25,
	0xa2, 0x17, 0xf1, 0x28, 0x5e, 0x74, 0x78, 0x29, 0xe2, 0xc1, 0x4b, 0x49, 0xdb, 0xb7, 0x59, 0x58,
	0xdb, 0xd4, 0x24, 0x1d, 0xeb, 0xe7, 0xf0, 0x0b, 0x8b, 0xc9, 0x66, 0x0b, 0xde, 0xf2, 0xbc, 0xef,
	0xf3, 0x7b, 0x78, 0x78, 0x83, 0xa2, 0x12, 0x72, 0x0e, 0x8a, 0x7e, 0xb5, 0xa0, 0x3a, 0x05, 0xba,
	0x2d, 0x0d, 0xdd, 0xee, 0x12, 0x2b, 0x13, 0xa7, 0x49, 0xa3, 0xa4, 0x91, 0x78, 0x3e, 0xb0, 0x2c,
	0x6e, 0xb8, 0x94, 0xbc, 0x04, 0x6a, 0x57, 0x69, 0x5b, 0x50, 0x23, 0x2a, 0xd0, 0x86, 0x55, 0x8d,
	0x73, 0x2f, 0x5f, 0xd1, 0x68, 0xfd, 0x81, 0xaf, 0x91, 0x5f, 0xb3, 0x0a, 0x74, 0xc3, 0x32, 0x08,
	0xbc, 0xd0, 0x8b, 0xfc, 0xb8, 0x1f, 0xe0, 0x73, 0x34, 0xde, 0x42, 0x17, 0x8c, 0xec, 0xfc, 0xf7,
	0x89, 0x2f, 0xd1, 0x74, 0xc7, 0xca, 0x16, 0x82, 0x71, 0xe8, 0x45, 0xa7, 0xb1, 0x13, 0xcb, 0x6f,
	0x0f, 0x9d, 0xad, 0xa1, 0x7b, 0x93, 0xb9, 0x28, 0x44, 0xc6, 0x8c, 0x90, 0x35, 0xbe, 0x40, 0x53,
	0xb3, 0x4f, 0x44, 0x7e, 0x48, 0x9d, 0x98, 0xfd, 0x4b, 0xde, 0xe3, 0xa3, 0x01, 0x8e, 0x1f, 0x90,
	0xff, 0xd7, 0xce, 0x06, 0xcf, 0xef, 0x16, 0xc4, 0xf5, 0x27, 0xc7, 0xfe, 0xe4, 0xfd, 0xe8, 0x88,
	0x7b, 0x33, 0xbe, 0x42, 0xbe, 0xd0, 0x49, 0x0e, 0x25, 0x18, 0x08, 0x26, 0xa1, 0x17, 0xcd, 0xe2,
	0x99, 0xd0, 0xcf, 0x56, 0x3f, 0xd5, 0xe8, 0x56, 0x2a, 0x4e, 0x36, 0x5d, 0x03, 0xca, 0x1d, 0x91,
	0x14, 0x2c, 0x55, 0x22, 0x73, 0xa1, 0x9a, 0x1c, 0x86, 0x83, 0xb3, 0x7d, 0x3e, 0x72, 0x61, 0x36,
	0x6d, 0x4a, 0x32, 0x59, 0xd1, 0x01, 0x48, 0x1d, 0xb8, 0x72, 0xe0, 0x8a, 0x4b, 0xfa, 0xff, 0x57,
	0xd2, 0x13, 0xbb, 0xbd, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x07, 0x6d, 0xd7, 0x47, 0xb2, 0x01,
	0x00, 0x00,
}
