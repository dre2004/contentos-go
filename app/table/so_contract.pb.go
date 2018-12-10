// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_contract.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoContract struct {
	Id                   *prototype.ContractId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedTime          *prototype.TimePointSec `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Abi                  string                  `protobuf:"bytes,3,opt,name=abi,proto3" json:"abi,omitempty"`
	Balance              *prototype.Coin         `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
	Code                 []byte                  `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoContract) Reset()         { *m = SoContract{} }
func (m *SoContract) String() string { return proto.CompactTextString(m) }
func (*SoContract) ProtoMessage()    {}
func (*SoContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{0}
}

func (m *SoContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoContract.Unmarshal(m, b)
}
func (m *SoContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoContract.Marshal(b, m, deterministic)
}
func (m *SoContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoContract.Merge(m, src)
}
func (m *SoContract) XXX_Size() int {
	return xxx_messageInfo_SoContract.Size(m)
}
func (m *SoContract) XXX_DiscardUnknown() {
	xxx_messageInfo_SoContract.DiscardUnknown(m)
}

var xxx_messageInfo_SoContract proto.InternalMessageInfo

func (m *SoContract) GetId() *prototype.ContractId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SoContract) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoContract) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *SoContract) GetBalance() *prototype.Coin {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *SoContract) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type SoMemContractById struct {
	Id                   *prototype.ContractId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SoMemContractById) Reset()         { *m = SoMemContractById{} }
func (m *SoMemContractById) String() string { return proto.CompactTextString(m) }
func (*SoMemContractById) ProtoMessage()    {}
func (*SoMemContractById) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{1}
}

func (m *SoMemContractById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemContractById.Unmarshal(m, b)
}
func (m *SoMemContractById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemContractById.Marshal(b, m, deterministic)
}
func (m *SoMemContractById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemContractById.Merge(m, src)
}
func (m *SoMemContractById) XXX_Size() int {
	return xxx_messageInfo_SoMemContractById.Size(m)
}
func (m *SoMemContractById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemContractById.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemContractById proto.InternalMessageInfo

func (m *SoMemContractById) GetId() *prototype.ContractId {
	if m != nil {
		return m.Id
	}
	return nil
}

type SoMemContractByCreatedTime struct {
	CreatedTime          *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoMemContractByCreatedTime) Reset()         { *m = SoMemContractByCreatedTime{} }
func (m *SoMemContractByCreatedTime) String() string { return proto.CompactTextString(m) }
func (*SoMemContractByCreatedTime) ProtoMessage()    {}
func (*SoMemContractByCreatedTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{2}
}

func (m *SoMemContractByCreatedTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemContractByCreatedTime.Unmarshal(m, b)
}
func (m *SoMemContractByCreatedTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemContractByCreatedTime.Marshal(b, m, deterministic)
}
func (m *SoMemContractByCreatedTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemContractByCreatedTime.Merge(m, src)
}
func (m *SoMemContractByCreatedTime) XXX_Size() int {
	return xxx_messageInfo_SoMemContractByCreatedTime.Size(m)
}
func (m *SoMemContractByCreatedTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemContractByCreatedTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemContractByCreatedTime proto.InternalMessageInfo

func (m *SoMemContractByCreatedTime) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

type SoMemContractByAbi struct {
	Abi                  string   `protobuf:"bytes,1,opt,name=abi,proto3" json:"abi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemContractByAbi) Reset()         { *m = SoMemContractByAbi{} }
func (m *SoMemContractByAbi) String() string { return proto.CompactTextString(m) }
func (*SoMemContractByAbi) ProtoMessage()    {}
func (*SoMemContractByAbi) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{3}
}

func (m *SoMemContractByAbi) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemContractByAbi.Unmarshal(m, b)
}
func (m *SoMemContractByAbi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemContractByAbi.Marshal(b, m, deterministic)
}
func (m *SoMemContractByAbi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemContractByAbi.Merge(m, src)
}
func (m *SoMemContractByAbi) XXX_Size() int {
	return xxx_messageInfo_SoMemContractByAbi.Size(m)
}
func (m *SoMemContractByAbi) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemContractByAbi.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemContractByAbi proto.InternalMessageInfo

func (m *SoMemContractByAbi) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

type SoMemContractByBalance struct {
	Balance              *prototype.Coin `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SoMemContractByBalance) Reset()         { *m = SoMemContractByBalance{} }
func (m *SoMemContractByBalance) String() string { return proto.CompactTextString(m) }
func (*SoMemContractByBalance) ProtoMessage()    {}
func (*SoMemContractByBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{4}
}

func (m *SoMemContractByBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemContractByBalance.Unmarshal(m, b)
}
func (m *SoMemContractByBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemContractByBalance.Marshal(b, m, deterministic)
}
func (m *SoMemContractByBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemContractByBalance.Merge(m, src)
}
func (m *SoMemContractByBalance) XXX_Size() int {
	return xxx_messageInfo_SoMemContractByBalance.Size(m)
}
func (m *SoMemContractByBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemContractByBalance.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemContractByBalance proto.InternalMessageInfo

func (m *SoMemContractByBalance) GetBalance() *prototype.Coin {
	if m != nil {
		return m.Balance
	}
	return nil
}

type SoMemContractByCode struct {
	Code                 []byte   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoMemContractByCode) Reset()         { *m = SoMemContractByCode{} }
func (m *SoMemContractByCode) String() string { return proto.CompactTextString(m) }
func (*SoMemContractByCode) ProtoMessage()    {}
func (*SoMemContractByCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{5}
}

func (m *SoMemContractByCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoMemContractByCode.Unmarshal(m, b)
}
func (m *SoMemContractByCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoMemContractByCode.Marshal(b, m, deterministic)
}
func (m *SoMemContractByCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoMemContractByCode.Merge(m, src)
}
func (m *SoMemContractByCode) XXX_Size() int {
	return xxx_messageInfo_SoMemContractByCode.Size(m)
}
func (m *SoMemContractByCode) XXX_DiscardUnknown() {
	xxx_messageInfo_SoMemContractByCode.DiscardUnknown(m)
}

var xxx_messageInfo_SoMemContractByCode proto.InternalMessageInfo

func (m *SoMemContractByCode) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type SoListContractByCreatedTime struct {
	CreatedTime          *prototype.TimePointSec `protobuf:"bytes,1,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Id                   *prototype.ContractId   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SoListContractByCreatedTime) Reset()         { *m = SoListContractByCreatedTime{} }
func (m *SoListContractByCreatedTime) String() string { return proto.CompactTextString(m) }
func (*SoListContractByCreatedTime) ProtoMessage()    {}
func (*SoListContractByCreatedTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{6}
}

func (m *SoListContractByCreatedTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListContractByCreatedTime.Unmarshal(m, b)
}
func (m *SoListContractByCreatedTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListContractByCreatedTime.Marshal(b, m, deterministic)
}
func (m *SoListContractByCreatedTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListContractByCreatedTime.Merge(m, src)
}
func (m *SoListContractByCreatedTime) XXX_Size() int {
	return xxx_messageInfo_SoListContractByCreatedTime.Size(m)
}
func (m *SoListContractByCreatedTime) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListContractByCreatedTime.DiscardUnknown(m)
}

var xxx_messageInfo_SoListContractByCreatedTime proto.InternalMessageInfo

func (m *SoListContractByCreatedTime) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoListContractByCreatedTime) GetId() *prototype.ContractId {
	if m != nil {
		return m.Id
	}
	return nil
}

type SoUniqueContractById struct {
	Id                   *prototype.ContractId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SoUniqueContractById) Reset()         { *m = SoUniqueContractById{} }
func (m *SoUniqueContractById) String() string { return proto.CompactTextString(m) }
func (*SoUniqueContractById) ProtoMessage()    {}
func (*SoUniqueContractById) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ccfb8b33a116ac5, []int{7}
}

func (m *SoUniqueContractById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueContractById.Unmarshal(m, b)
}
func (m *SoUniqueContractById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueContractById.Marshal(b, m, deterministic)
}
func (m *SoUniqueContractById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueContractById.Merge(m, src)
}
func (m *SoUniqueContractById) XXX_Size() int {
	return xxx_messageInfo_SoUniqueContractById.Size(m)
}
func (m *SoUniqueContractById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueContractById.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueContractById proto.InternalMessageInfo

func (m *SoUniqueContractById) GetId() *prototype.ContractId {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*SoContract)(nil), "table.so_contract")
	proto.RegisterType((*SoMemContractById)(nil), "table.so_mem_contract_by_id")
	proto.RegisterType((*SoMemContractByCreatedTime)(nil), "table.so_mem_contract_by_created_time")
	proto.RegisterType((*SoMemContractByAbi)(nil), "table.so_mem_contract_by_abi")
	proto.RegisterType((*SoMemContractByBalance)(nil), "table.so_mem_contract_by_balance")
	proto.RegisterType((*SoMemContractByCode)(nil), "table.so_mem_contract_by_code")
	proto.RegisterType((*SoListContractByCreatedTime)(nil), "table.so_list_contract_by_created_time")
	proto.RegisterType((*SoUniqueContractById)(nil), "table.so_unique_contract_by_id")
}

func init() { proto.RegisterFile("app/table/so_contract.proto", fileDescriptor_6ccfb8b33a116ac5) }

var fileDescriptor_6ccfb8b33a116ac5 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xd9, 0xb4, 0xfd, 0xff, 0x71, 0x5a, 0x50, 0x82, 0xd6, 0x58, 0x0f, 0x86, 0x1c, 0x4a,
	0x14, 0x9a, 0x80, 0x5e, 0x05, 0xa1, 0x17, 0xef, 0xc1, 0x93, 0x97, 0x65, 0xb3, 0xbb, 0xb4, 0x03,
	0xcd, 0x4e, 0xec, 0x6e, 0x0f, 0xfd, 0x06, 0x7e, 0x30, 0x3f, 0x98, 0x24, 0x7d, 0x31, 0x62, 0xf0,
	0x0d, 0xbc, 0x84, 0x49, 0xe6, 0x79, 0x66, 0xf8, 0x3d, 0x43, 0xe0, 0x5c, 0x94, 0x65, 0xea, 0x44,
	0xbe, 0xd0, 0xa9, 0x25, 0x2e, 0xc9, 0xb8, 0xa5, 0x90, 0x2e, 0x29, 0x97, 0xe4, 0xc8, 0xef, 0xd5,
	0x8d, 0xd1, 0x71, 0xfd, 0xe6, 0xd6, 0xa5, 0x4e, 0xab, 0xc7, 0xa6, 0x19, 0xbd, 0x30, 0xe8, 0x37,
	0x2c, 0xfe, 0x18, 0x3c, 0x54, 0x01, 0x0b, 0x59, 0xdc, 0xbf, 0x1e, 0x26, 0x7b, 0x4b, 0xb2, 0x13,
	0x70, 0x54, 0x99, 0x87, 0xca, 0xbf, 0x85, 0x81, 0x5c, 0x6a, 0xe1, 0xb4, 0xe2, 0x0e, 0x0b, 0x1d,
	0x78, 0xb5, 0xe3, 0xac, 0xe1, 0xa8, 0x3e, 0xf3, 0x92, 0xd0, 0x38, 0x6e, 0xb5, 0xcc, 0xfa, 0x5b,
	0xf9, 0x03, 0x16, 0xda, 0x3f, 0x82, 0x8e, 0xc8, 0x31, 0xe8, 0x84, 0x2c, 0x3e, 0xc8, 0xaa, 0xd2,
	0xbf, 0x84, 0xff, 0xb9, 0x58, 0x08, 0x23, 0x75, 0xd0, 0xad, 0x47, 0x1d, 0xbe, 0x5b, 0x8e, 0x26,
	0xdb, 0xf5, 0x7d, 0x1f, 0xba, 0x92, 0x94, 0x0e, 0x7a, 0x21, 0x8b, 0x07, 0x59, 0x5d, 0x47, 0x77,
	0x70, 0x62, 0x89, 0x17, 0xba, 0xd8, 0x93, 0xf0, 0x7c, 0xcd, 0x51, 0x7d, 0x97, 0x27, 0xe2, 0x70,
	0xd1, 0x32, 0xa0, 0x89, 0xf8, 0x01, 0x99, 0xfd, 0x04, 0x39, 0xba, 0x82, 0x61, 0xcb, 0x82, 0x0a,
	0x7d, 0x1b, 0x06, 0xdb, 0x87, 0x11, 0xdd, 0xc3, 0xa8, 0x45, 0xbb, 0xe3, 0x6f, 0x44, 0xc5, 0x3e,
	0x8f, 0x2a, 0x9a, 0xc0, 0x69, 0x1b, 0x15, 0xa9, 0xb7, 0x14, 0x59, 0x23, 0xc5, 0x67, 0x06, 0xa1,
	0x25, 0xbe, 0x40, 0xeb, 0xfe, 0x28, 0x86, 0xed, 0x3d, 0xbc, 0x2f, 0xef, 0x31, 0x85, 0xc0, 0x12,
	0x5f, 0x19, 0x7c, 0x5a, 0xe9, 0x5f, 0xde, 0x74, 0x1a, 0x3f, 0x8e, 0x67, 0xe8, 0xe6, 0xab, 0x3c,
	0x91, 0x54, 0xa4, 0x92, 0xac, 0x9c, 0x0b, 0x34, 0x69, 0x25, 0xd3, 0xc6, 0x91, 0x9d, 0xcc, 0x68,
	0xf3, 0xd3, 0xe4, 0xff, 0xea, 0x21, 0x37, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xdc, 0xa2,
	0x22, 0x48, 0x03, 0x00, 0x00,
}