// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_transactionObject.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
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

type SoTransactionObject struct {
	TrxId                *prototype.Sha256       `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	Expiration           *prototype.TimePointSec `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoTransactionObject) Reset()         { *m = SoTransactionObject{} }
func (m *SoTransactionObject) String() string { return proto.CompactTextString(m) }
func (*SoTransactionObject) ProtoMessage()    {}
func (*SoTransactionObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e5dbc6d4c15723, []int{0}
}

func (m *SoTransactionObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoTransactionObject.Unmarshal(m, b)
}
func (m *SoTransactionObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoTransactionObject.Marshal(b, m, deterministic)
}
func (m *SoTransactionObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoTransactionObject.Merge(m, src)
}
func (m *SoTransactionObject) XXX_Size() int {
	return xxx_messageInfo_SoTransactionObject.Size(m)
}
func (m *SoTransactionObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SoTransactionObject.DiscardUnknown(m)
}

var xxx_messageInfo_SoTransactionObject proto.InternalMessageInfo

func (m *SoTransactionObject) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func (m *SoTransactionObject) GetExpiration() *prototype.TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type SoListTransactionObjectByExpiration struct {
	Expiration           *prototype.TimePointSec `protobuf:"bytes,1,opt,name=expiration,proto3" json:"expiration,omitempty"`
	TrxId                *prototype.Sha256       `protobuf:"bytes,2,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListTransactionObjectByExpiration) Reset()         { *m = SoListTransactionObjectByExpiration{} }
func (m *SoListTransactionObjectByExpiration) String() string { return proto.CompactTextString(m) }
func (*SoListTransactionObjectByExpiration) ProtoMessage()    {}
func (*SoListTransactionObjectByExpiration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e5dbc6d4c15723, []int{1}
}

func (m *SoListTransactionObjectByExpiration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Unmarshal(m, b)
}
func (m *SoListTransactionObjectByExpiration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Marshal(b, m, deterministic)
}
func (m *SoListTransactionObjectByExpiration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListTransactionObjectByExpiration.Merge(m, src)
}
func (m *SoListTransactionObjectByExpiration) XXX_Size() int {
	return xxx_messageInfo_SoListTransactionObjectByExpiration.Size(m)
}
func (m *SoListTransactionObjectByExpiration) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListTransactionObjectByExpiration.DiscardUnknown(m)
}

var xxx_messageInfo_SoListTransactionObjectByExpiration proto.InternalMessageInfo

func (m *SoListTransactionObjectByExpiration) GetExpiration() *prototype.TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *SoListTransactionObjectByExpiration) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

type SoUniqueTransactionObjectByTrxId struct {
	TrxId                *prototype.Sha256 `protobuf:"bytes,1,opt,name=trx_id,json=trxId,proto3" json:"trx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoUniqueTransactionObjectByTrxId) Reset()         { *m = SoUniqueTransactionObjectByTrxId{} }
func (m *SoUniqueTransactionObjectByTrxId) String() string { return proto.CompactTextString(m) }
func (*SoUniqueTransactionObjectByTrxId) ProtoMessage()    {}
func (*SoUniqueTransactionObjectByTrxId) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0e5dbc6d4c15723, []int{2}
}

func (m *SoUniqueTransactionObjectByTrxId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Unmarshal(m, b)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Marshal(b, m, deterministic)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Merge(m, src)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_Size() int {
	return xxx_messageInfo_SoUniqueTransactionObjectByTrxId.Size(m)
}
func (m *SoUniqueTransactionObjectByTrxId) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueTransactionObjectByTrxId.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueTransactionObjectByTrxId proto.InternalMessageInfo

func (m *SoUniqueTransactionObjectByTrxId) GetTrxId() *prototype.Sha256 {
	if m != nil {
		return m.TrxId
	}
	return nil
}

func init() {
	proto.RegisterType((*SoTransactionObject)(nil), "table.so_transactionObject")
	proto.RegisterType((*SoListTransactionObjectByExpiration)(nil), "table.so_list_transactionObject_by_expiration")
	proto.RegisterType((*SoUniqueTransactionObjectByTrxId)(nil), "table.so_unique_transactionObject_by_trx_id")
}

func init() {
	proto.RegisterFile("app/table/so_transactionObject.proto", fileDescriptor_b0e5dbc6d4c15723)
}

var fileDescriptor_b0e5dbc6d4c15723 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xd9, 0x42, 0x7b, 0x88, 0x27, 0x97, 0x1e, 0xaa, 0x27, 0x59, 0x14, 0x8b, 0xe8, 0x06,
	0x2a, 0x0a, 0x5e, 0xbd, 0x79, 0x12, 0x7b, 0xf4, 0x12, 0x92, 0x74, 0xe8, 0x8e, 0xb4, 0x99, 0x98,
	0x99, 0x85, 0x2d, 0xfe, 0x07, 0x7f, 0xb3, 0x74, 0x0b, 0x75, 0xd1, 0x45, 0xec, 0x25, 0x10, 0xf2,
	0xbd, 0x79, 0xef, 0x65, 0xd4, 0xb9, 0x8d, 0x51, 0x8b, 0x75, 0x2b, 0xd0, 0x4c, 0x46, 0x92, 0x0d,
	0x6c, 0xbd, 0x20, 0x85, 0x67, 0xf7, 0x06, 0x5e, 0xca, 0x98, 0x48, 0x28, 0x1f, 0xb6, 0xc4, 0xe9,
	0xb8, 0xbd, 0xc9, 0x26, 0x82, 0xde, 0x1e, 0xbb, 0xc7, 0xe2, 0x43, 0x8d, 0xfb, 0xa4, 0xf9, 0x54,
	0x8d, 0x24, 0x35, 0x06, 0x17, 0x93, 0xec, 0x2c, 0x9b, 0x1e, 0xcd, 0x8e, 0xcb, 0xbd, 0xbc, 0xe4,
	0xca, 0xce, 0xee, 0xee, 0xe7, 0x43, 0x49, 0xcd, 0xd3, 0x22, 0x7f, 0x50, 0x0a, 0x9a, 0x88, 0xc9,
	0x6e, 0xd5, 0x93, 0x41, 0x4b, 0x9f, 0x74, 0x68, 0xc1, 0x35, 0x98, 0x48, 0x18, 0xc4, 0x30, 0xf8,
	0x79, 0x07, 0x2e, 0x3e, 0x33, 0x75, 0xc9, 0x64, 0x56, 0xc8, 0xf2, 0x3b, 0x82, 0x71, 0x1b, 0xf3,
	0xcd, 0xfe, 0xb0, 0xc9, 0x0e, 0xb0, 0xe9, 0x74, 0x19, 0xfc, 0xdd, 0xa5, 0x78, 0x51, 0x17, 0x4c,
	0xa6, 0x0e, 0xf8, 0x5e, 0x43, 0x7f, 0xa2, 0xdd, 0xa0, 0xff, 0x7f, 0xcf, 0xe3, 0xf5, 0xeb, 0xd5,
	0x12, 0xa5, 0xaa, 0x5d, 0xe9, 0x69, 0xad, 0x3d, 0xb1, 0xaf, 0x2c, 0x06, 0xed, 0x29, 0x08, 0x04,
	0x21, 0xbe, 0x59, 0x92, 0xde, 0xaf, 0xd1, 0x8d, 0xda, 0x31, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xba, 0x9c, 0xf1, 0x81, 0xda, 0x01, 0x00, 0x00,
}
