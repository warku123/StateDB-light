// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: rpc/sdb.proto

package sdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty bool `protobuf:"varint,1,opt,name=Empty,proto3" json:"Empty,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

type SubBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr   string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *SubBalanceRequest) Reset() {
	*x = SubBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubBalanceRequest) ProtoMessage() {}

func (x *SubBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubBalanceRequest.ProtoReflect.Descriptor instead.
func (*SubBalanceRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{2}
}

func (x *SubBalanceRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *SubBalanceRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SubBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty bool `protobuf:"varint,1,opt,name=Empty,proto3" json:"Empty,omitempty"`
}

func (x *SubBalanceResponse) Reset() {
	*x = SubBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubBalanceResponse) ProtoMessage() {}

func (x *SubBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubBalanceResponse.ProtoReflect.Descriptor instead.
func (*SubBalanceResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{3}
}

func (x *SubBalanceResponse) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

type AddBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr   string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *AddBalanceRequest) Reset() {
	*x = AddBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBalanceRequest) ProtoMessage() {}

func (x *AddBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBalanceRequest.ProtoReflect.Descriptor instead.
func (*AddBalanceRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{4}
}

func (x *AddBalanceRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *AddBalanceRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty bool `protobuf:"varint,1,opt,name=Empty,proto3" json:"Empty,omitempty"`
}

func (x *AddBalanceResponse) Reset() {
	*x = AddBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBalanceResponse) ProtoMessage() {}

func (x *AddBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBalanceResponse.ProtoReflect.Descriptor instead.
func (*AddBalanceResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{5}
}

func (x *AddBalanceResponse) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{6}
}

func (x *GetBalanceRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{7}
}

func (x *GetBalanceResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SuicideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *SuicideRequest) Reset() {
	*x = SuicideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuicideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuicideRequest) ProtoMessage() {}

func (x *SuicideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuicideRequest.ProtoReflect.Descriptor instead.
func (*SuicideRequest) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{8}
}

func (x *SuicideRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type SuicideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuicide bool `protobuf:"varint,1,opt,name=Is_suicide,json=IsSuicide,proto3" json:"Is_suicide,omitempty"`
}

func (x *SuicideResponse) Reset() {
	*x = SuicideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_sdb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuicideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuicideResponse) ProtoMessage() {}

func (x *SuicideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_sdb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuicideResponse.ProtoReflect.Descriptor instead.
func (*SuicideResponse) Descriptor() ([]byte, []int) {
	return file_rpc_sdb_proto_rawDescGZIP(), []int{9}
}

func (x *SuicideResponse) GetIsSuicide() bool {
	if x != nil {
		return x.IsSuicide
	}
	return false
}

var File_rpc_sdb_proto protoreflect.FileDescriptor

var file_rpc_sdb_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x73, 0x64, 0x62, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72,
	0x22, 0x2d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x3f, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x2a, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3f, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64,
	0x64, 0x72, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x24, 0x0a, 0x0e, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x22, 0x30, 0x0a, 0x0f, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x49, 0x73, 0x5f,
	0x73, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49,
	0x73, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x32, 0xc0, 0x02, 0x0a, 0x03, 0x53, 0x64, 0x62,
	0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x19, 0x2e, 0x73, 0x64, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x64, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x64, 0x62, 0x2e, 0x53, 0x75, 0x62,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x64, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x64, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x73, 0x64, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73,
	0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65,
	0x12, 0x13, 0x2e, 0x73, 0x64, 0x62, 0x2e, 0x53, 0x75, 0x69, 0x63, 0x69, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x64, 0x62, 0x2e, 0x53, 0x75, 0x69, 0x63,
	0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x73, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_sdb_proto_rawDescOnce sync.Once
	file_rpc_sdb_proto_rawDescData = file_rpc_sdb_proto_rawDesc
)

func file_rpc_sdb_proto_rawDescGZIP() []byte {
	file_rpc_sdb_proto_rawDescOnce.Do(func() {
		file_rpc_sdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_sdb_proto_rawDescData)
	})
	return file_rpc_sdb_proto_rawDescData
}

var file_rpc_sdb_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rpc_sdb_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),  // 0: sdb.CreateAccountRequest
	(*CreateAccountResponse)(nil), // 1: sdb.CreateAccountResponse
	(*SubBalanceRequest)(nil),     // 2: sdb.SubBalanceRequest
	(*SubBalanceResponse)(nil),    // 3: sdb.SubBalanceResponse
	(*AddBalanceRequest)(nil),     // 4: sdb.AddBalanceRequest
	(*AddBalanceResponse)(nil),    // 5: sdb.AddBalanceResponse
	(*GetBalanceRequest)(nil),     // 6: sdb.GetBalanceRequest
	(*GetBalanceResponse)(nil),    // 7: sdb.GetBalanceResponse
	(*SuicideRequest)(nil),        // 8: sdb.SuicideRequest
	(*SuicideResponse)(nil),       // 9: sdb.SuicideResponse
}
var file_rpc_sdb_proto_depIdxs = []int32{
	0, // 0: sdb.Sdb.CreateAccount:input_type -> sdb.CreateAccountRequest
	2, // 1: sdb.Sdb.SubBalance:input_type -> sdb.SubBalanceRequest
	4, // 2: sdb.Sdb.AddBalance:input_type -> sdb.AddBalanceRequest
	6, // 3: sdb.Sdb.GetBalance:input_type -> sdb.GetBalanceRequest
	8, // 4: sdb.Sdb.Suicide:input_type -> sdb.SuicideRequest
	1, // 5: sdb.Sdb.CreateAccount:output_type -> sdb.CreateAccountResponse
	3, // 6: sdb.Sdb.SubBalance:output_type -> sdb.SubBalanceResponse
	5, // 7: sdb.Sdb.AddBalance:output_type -> sdb.AddBalanceResponse
	7, // 8: sdb.Sdb.GetBalance:output_type -> sdb.GetBalanceResponse
	9, // 9: sdb.Sdb.Suicide:output_type -> sdb.SuicideResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_sdb_proto_init() }
func file_rpc_sdb_proto_init() {
	if File_rpc_sdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_sdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuicideRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_sdb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuicideResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_sdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_sdb_proto_goTypes,
		DependencyIndexes: file_rpc_sdb_proto_depIdxs,
		MessageInfos:      file_rpc_sdb_proto_msgTypes,
	}.Build()
	File_rpc_sdb_proto = out.File
	file_rpc_sdb_proto_rawDesc = nil
	file_rpc_sdb_proto_goTypes = nil
	file_rpc_sdb_proto_depIdxs = nil
}
