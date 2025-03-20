// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: dapplink/keylocker.proto

package keylocker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReturnCode int32

const (
	ReturnCode_SUCCESS ReturnCode = 0
	ReturnCode_ERROR   ReturnCode = 1
)

// Enum value maps for ReturnCode.
var (
	ReturnCode_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
	}
	ReturnCode_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   1,
	}
)

func (x ReturnCode) Enum() *ReturnCode {
	p := new(ReturnCode)
	*p = x
	return p
}

func (x ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_dapplink_keylocker_proto_enumTypes[0].Descriptor()
}

func (ReturnCode) Type() protoreflect.EnumType {
	return &file_dapplink_keylocker_proto_enumTypes[0]
}

func (x ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnCode.Descriptor instead.
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{0}
}

type SocialKey struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key           string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SocialKey) Reset() {
	*x = SocialKey{}
	mi := &file_dapplink_keylocker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SocialKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialKey) ProtoMessage() {}

func (x *SocialKey) ProtoReflect() protoreflect.Message {
	mi := &file_dapplink_keylocker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialKey.ProtoReflect.Descriptor instead.
func (*SocialKey) Descriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{0}
}

func (x *SocialKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SocialKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SupportChainReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConsumerToken string                 `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string                 `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	Network       string                 `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SupportChainReq) Reset() {
	*x = SupportChainReq{}
	mi := &file_dapplink_keylocker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportChainReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportChainReq) ProtoMessage() {}

func (x *SupportChainReq) ProtoReflect() protoreflect.Message {
	mi := &file_dapplink_keylocker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportChainReq.ProtoReflect.Descriptor instead.
func (*SupportChainReq) Descriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{1}
}

func (x *SupportChainReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SupportChainReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *SupportChainReq) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

type SupportChainRep struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          ReturnCode             `protobuf:"varint,1,opt,name=code,proto3,enum=dapplink.keylocker.ReturnCode" json:"code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Support       bool                   `protobuf:"varint,3,opt,name=support,proto3" json:"support,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SupportChainRep) Reset() {
	*x = SupportChainRep{}
	mi := &file_dapplink_keylocker_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportChainRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportChainRep) ProtoMessage() {}

func (x *SupportChainRep) ProtoReflect() protoreflect.Message {
	mi := &file_dapplink_keylocker_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportChainRep.ProtoReflect.Descriptor instead.
func (*SupportChainRep) Descriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{2}
}

func (x *SupportChainRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *SupportChainRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SupportChainRep) GetSupport() bool {
	if x != nil {
		return x.Support
	}
	return false
}

type SetSocialKeyReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConsumerToken string                 `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string                 `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	WalletUuid    string                 `protobuf:"bytes,3,opt,name=wallet_uuid,json=walletUuid,proto3" json:"wallet_uuid,omitempty"`
	Key           string                 `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	SocialCode    string                 `protobuf:"bytes,6,opt,name=social_code,json=socialCode,proto3" json:"social_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetSocialKeyReq) Reset() {
	*x = SetSocialKeyReq{}
	mi := &file_dapplink_keylocker_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetSocialKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSocialKeyReq) ProtoMessage() {}

func (x *SetSocialKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_dapplink_keylocker_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSocialKeyReq.ProtoReflect.Descriptor instead.
func (*SetSocialKeyReq) Descriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{3}
}

func (x *SetSocialKeyReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SetSocialKeyReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *SetSocialKeyReq) GetWalletUuid() string {
	if x != nil {
		return x.WalletUuid
	}
	return ""
}

func (x *SetSocialKeyReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetSocialKeyReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SetSocialKeyReq) GetSocialCode() string {
	if x != nil {
		return x.SocialCode
	}
	return ""
}

type SetSocialKeyRep struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          ReturnCode             `protobuf:"varint,1,opt,name=code,proto3,enum=dapplink.keylocker.ReturnCode" json:"code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Pub           string                 `protobuf:"bytes,3,opt,name=pub,proto3" json:"pub,omitempty"`
	Priv          string                 `protobuf:"bytes,4,opt,name=priv,proto3" json:"priv,omitempty"`
	CryptoWay     string                 `protobuf:"bytes,5,opt,name=crypto_way,json=cryptoWay,proto3" json:"crypto_way,omitempty"`
	FileCid       string                 `protobuf:"bytes,6,opt,name=file_cid,json=fileCid,proto3" json:"file_cid,omitempty"`
	Contract      string                 `protobuf:"bytes,7,opt,name=contract,proto3" json:"contract,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetSocialKeyRep) Reset() {
	*x = SetSocialKeyRep{}
	mi := &file_dapplink_keylocker_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetSocialKeyRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSocialKeyRep) ProtoMessage() {}

func (x *SetSocialKeyRep) ProtoReflect() protoreflect.Message {
	mi := &file_dapplink_keylocker_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSocialKeyRep.ProtoReflect.Descriptor instead.
func (*SetSocialKeyRep) Descriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{4}
}

func (x *SetSocialKeyRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *SetSocialKeyRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SetSocialKeyRep) GetPub() string {
	if x != nil {
		return x.Pub
	}
	return ""
}

func (x *SetSocialKeyRep) GetPriv() string {
	if x != nil {
		return x.Priv
	}
	return ""
}

func (x *SetSocialKeyRep) GetCryptoWay() string {
	if x != nil {
		return x.CryptoWay
	}
	return ""
}

func (x *SetSocialKeyRep) GetFileCid() string {
	if x != nil {
		return x.FileCid
	}
	return ""
}

func (x *SetSocialKeyRep) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type GetSocialKeyReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConsumerToken string                 `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Chain         string                 `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	WalletUuid    string                 `protobuf:"bytes,3,opt,name=wallet_uuid,json=walletUuid,proto3" json:"wallet_uuid,omitempty"`
	FileCid       string                 `protobuf:"bytes,4,opt,name=file_cid,json=fileCid,proto3" json:"file_cid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSocialKeyReq) Reset() {
	*x = GetSocialKeyReq{}
	mi := &file_dapplink_keylocker_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSocialKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSocialKeyReq) ProtoMessage() {}

func (x *GetSocialKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_dapplink_keylocker_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSocialKeyReq.ProtoReflect.Descriptor instead.
func (*GetSocialKeyReq) Descriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{5}
}

func (x *GetSocialKeyReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *GetSocialKeyReq) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *GetSocialKeyReq) GetWalletUuid() string {
	if x != nil {
		return x.WalletUuid
	}
	return ""
}

func (x *GetSocialKeyReq) GetFileCid() string {
	if x != nil {
		return x.FileCid
	}
	return ""
}

type GetSocialKeyRep struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          ReturnCode             `protobuf:"varint,1,opt,name=code,proto3,enum=dapplink.keylocker.ReturnCode" json:"code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	KeyList       []*SocialKey           `protobuf:"bytes,3,rep,name=key_list,json=keyList,proto3" json:"key_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSocialKeyRep) Reset() {
	*x = GetSocialKeyRep{}
	mi := &file_dapplink_keylocker_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSocialKeyRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSocialKeyRep) ProtoMessage() {}

func (x *GetSocialKeyRep) ProtoReflect() protoreflect.Message {
	mi := &file_dapplink_keylocker_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSocialKeyRep.ProtoReflect.Descriptor instead.
func (*GetSocialKeyRep) Descriptor() ([]byte, []int) {
	return file_dapplink_keylocker_proto_rawDescGZIP(), []int{6}
}

func (x *GetSocialKeyRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *GetSocialKeyRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetSocialKeyRep) GetKeyList() []*SocialKey {
	if x != nil {
		return x.KeyList
	}
	return nil
}

var File_dapplink_keylocker_proto protoreflect.FileDescriptor

var file_dapplink_keylocker_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x6b, 0x65, 0x79, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x2d,
	0x0a, 0x09, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x68, 0x0a,
	0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x71, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x0f, 0x53,
	0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x0f,
	0x53, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x12,
	0x32, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x75, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x69, 0x76, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x72, 0x69, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x61, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x43, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x22, 0x8a, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x69, 0x64, 0x22, 0x91,
	0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x70, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x2a, 0x24, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32, 0xa9, 0x02, 0x0a, 0x10, 0x4c, 0x65, 0x79,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a,
	0x0f, 0x67, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x23, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0c,
	0x73, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x2e, 0x64,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x23, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e,
	0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x70, 0x22, 0x00, 0x42, 0x2b, 0x0a, 0x16, 0x78, 0x79, 0x7a, 0x2e, 0x64, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5a, 0x11,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_dapplink_keylocker_proto_rawDescOnce sync.Once
	file_dapplink_keylocker_proto_rawDescData []byte
)

func file_dapplink_keylocker_proto_rawDescGZIP() []byte {
	file_dapplink_keylocker_proto_rawDescOnce.Do(func() {
		file_dapplink_keylocker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dapplink_keylocker_proto_rawDesc), len(file_dapplink_keylocker_proto_rawDesc)))
	})
	return file_dapplink_keylocker_proto_rawDescData
}

var file_dapplink_keylocker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dapplink_keylocker_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_dapplink_keylocker_proto_goTypes = []any{
	(ReturnCode)(0),         // 0: dapplink.keylocker.ReturnCode
	(*SocialKey)(nil),       // 1: dapplink.keylocker.SocialKey
	(*SupportChainReq)(nil), // 2: dapplink.keylocker.SupportChainReq
	(*SupportChainRep)(nil), // 3: dapplink.keylocker.SupportChainRep
	(*SetSocialKeyReq)(nil), // 4: dapplink.keylocker.SetSocialKeyReq
	(*SetSocialKeyRep)(nil), // 5: dapplink.keylocker.SetSocialKeyRep
	(*GetSocialKeyReq)(nil), // 6: dapplink.keylocker.GetSocialKeyReq
	(*GetSocialKeyRep)(nil), // 7: dapplink.keylocker.GetSocialKeyRep
}
var file_dapplink_keylocker_proto_depIdxs = []int32{
	0, // 0: dapplink.keylocker.SupportChainRep.code:type_name -> dapplink.keylocker.ReturnCode
	0, // 1: dapplink.keylocker.SetSocialKeyRep.code:type_name -> dapplink.keylocker.ReturnCode
	0, // 2: dapplink.keylocker.GetSocialKeyRep.code:type_name -> dapplink.keylocker.ReturnCode
	1, // 3: dapplink.keylocker.GetSocialKeyRep.key_list:type_name -> dapplink.keylocker.SocialKey
	2, // 4: dapplink.keylocker.LeyLockerService.getSupportChain:input_type -> dapplink.keylocker.SupportChainReq
	4, // 5: dapplink.keylocker.LeyLockerService.setSocialKey:input_type -> dapplink.keylocker.SetSocialKeyReq
	6, // 6: dapplink.keylocker.LeyLockerService.getSocialKey:input_type -> dapplink.keylocker.GetSocialKeyReq
	3, // 7: dapplink.keylocker.LeyLockerService.getSupportChain:output_type -> dapplink.keylocker.SupportChainRep
	5, // 8: dapplink.keylocker.LeyLockerService.setSocialKey:output_type -> dapplink.keylocker.SetSocialKeyRep
	7, // 9: dapplink.keylocker.LeyLockerService.getSocialKey:output_type -> dapplink.keylocker.GetSocialKeyRep
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dapplink_keylocker_proto_init() }
func file_dapplink_keylocker_proto_init() {
	if File_dapplink_keylocker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dapplink_keylocker_proto_rawDesc), len(file_dapplink_keylocker_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dapplink_keylocker_proto_goTypes,
		DependencyIndexes: file_dapplink_keylocker_proto_depIdxs,
		EnumInfos:         file_dapplink_keylocker_proto_enumTypes,
		MessageInfos:      file_dapplink_keylocker_proto_msgTypes,
	}.Build()
	File_dapplink_keylocker_proto = out.File
	file_dapplink_keylocker_proto_goTypes = nil
	file_dapplink_keylocker_proto_depIdxs = nil
}
