// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: savour_rpc/l3staking.proto

package l3staking

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

type StakingNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Address       string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	EthIncome     string `protobuf:"bytes,3,opt,name=eth_income,json=ethIncome,proto3" json:"eth_income,omitempty"`
	EthIncomeRate string `protobuf:"bytes,4,opt,name=eth_income_rate,json=ethIncomeRate,proto3" json:"eth_income_rate,omitempty"`
	DpIncome      string `protobuf:"bytes,5,opt,name=dp_income,json=dpIncome,proto3" json:"dp_income,omitempty"`
	DpIncomeRate  string `protobuf:"bytes,6,opt,name=dp_income_rate,json=dpIncomeRate,proto3" json:"dp_income_rate,omitempty"`
	EthEvil       string `protobuf:"bytes,7,opt,name=eth_evil,json=ethEvil,proto3" json:"eth_evil,omitempty"`
	EthEvilRate   string `protobuf:"bytes,8,opt,name=eth_evil_rate,json=ethEvilRate,proto3" json:"eth_evil_rate,omitempty"`
	DpEvil        string `protobuf:"bytes,9,opt,name=dp_evil,json=dpEvil,proto3" json:"dp_evil,omitempty"`
	DpEvilRate    string `protobuf:"bytes,10,opt,name=dp_evil_rate,json=dpEvilRate,proto3" json:"dp_evil_rate,omitempty"`
	Tvl           string `protobuf:"bytes,11,opt,name=tvl,proto3" json:"tvl,omitempty"`
}

func (x *StakingNodeReq) Reset() {
	*x = StakingNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savour_rpc_l3staking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingNodeReq) ProtoMessage() {}

func (x *StakingNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_savour_rpc_l3staking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingNodeReq.ProtoReflect.Descriptor instead.
func (*StakingNodeReq) Descriptor() ([]byte, []int) {
	return file_savour_rpc_l3staking_proto_rawDescGZIP(), []int{0}
}

func (x *StakingNodeReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *StakingNodeReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StakingNodeReq) GetEthIncome() string {
	if x != nil {
		return x.EthIncome
	}
	return ""
}

func (x *StakingNodeReq) GetEthIncomeRate() string {
	if x != nil {
		return x.EthIncomeRate
	}
	return ""
}

func (x *StakingNodeReq) GetDpIncome() string {
	if x != nil {
		return x.DpIncome
	}
	return ""
}

func (x *StakingNodeReq) GetDpIncomeRate() string {
	if x != nil {
		return x.DpIncomeRate
	}
	return ""
}

func (x *StakingNodeReq) GetEthEvil() string {
	if x != nil {
		return x.EthEvil
	}
	return ""
}

func (x *StakingNodeReq) GetEthEvilRate() string {
	if x != nil {
		return x.EthEvilRate
	}
	return ""
}

func (x *StakingNodeReq) GetDpEvil() string {
	if x != nil {
		return x.DpEvil
	}
	return ""
}

func (x *StakingNodeReq) GetDpEvilRate() string {
	if x != nil {
		return x.DpEvilRate
	}
	return ""
}

func (x *StakingNodeReq) GetTvl() string {
	if x != nil {
		return x.Tvl
	}
	return ""
}

type StakingNodeRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *StakingNodeRep) Reset() {
	*x = StakingNodeRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_savour_rpc_l3staking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakingNodeRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakingNodeRep) ProtoMessage() {}

func (x *StakingNodeRep) ProtoReflect() protoreflect.Message {
	mi := &file_savour_rpc_l3staking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakingNodeRep.ProtoReflect.Descriptor instead.
func (*StakingNodeRep) Descriptor() ([]byte, []int) {
	return file_savour_rpc_l3staking_proto_rawDescGZIP(), []int{1}
}

func (x *StakingNodeRep) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StakingNodeRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_savour_rpc_l3staking_proto protoreflect.FileDescriptor

var file_savour_rpc_l3staking_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x33, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x6c, 0x33, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0xe7, 0x02, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x74, 0x68, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x65, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x74, 0x68, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x70, 0x5f, 0x69, 0x6e, 0x63,
	0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x70, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x70, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x70, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x74, 0x68,
	0x5f, 0x65, 0x76, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x74, 0x68,
	0x45, 0x76, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x65, 0x74, 0x68, 0x5f, 0x65, 0x76, 0x69, 0x6c,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x74, 0x68,
	0x45, 0x76, 0x69, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x70, 0x5f, 0x65,
	0x76, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x70, 0x45, 0x76, 0x69,
	0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x70, 0x5f, 0x65, 0x76, 0x69, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x70, 0x45, 0x76, 0x69, 0x6c, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x76, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x76, 0x6c, 0x22, 0x36, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x8d, 0x01,
	0x0a, 0x10, 0x4c, 0x33, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x79, 0x0a, 0x17, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x2d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x6c, 0x33, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x6c, 0x33, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x22, 0x00, 0x42, 0x2b, 0x0a,
	0x16, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x73, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x2e, 0x6c, 0x33,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x33, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_savour_rpc_l3staking_proto_rawDescOnce sync.Once
	file_savour_rpc_l3staking_proto_rawDescData = file_savour_rpc_l3staking_proto_rawDesc
)

func file_savour_rpc_l3staking_proto_rawDescGZIP() []byte {
	file_savour_rpc_l3staking_proto_rawDescOnce.Do(func() {
		file_savour_rpc_l3staking_proto_rawDescData = protoimpl.X.CompressGZIP(file_savour_rpc_l3staking_proto_rawDescData)
	})
	return file_savour_rpc_l3staking_proto_rawDescData
}

var file_savour_rpc_l3staking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_savour_rpc_l3staking_proto_goTypes = []interface{}{
	(*StakingNodeReq)(nil), // 0: services.savour_rpc.l3staking.StakingNodeReq
	(*StakingNodeRep)(nil), // 1: services.savour_rpc.l3staking.StakingNodeRep
}
var file_savour_rpc_l3staking_proto_depIdxs = []int32{
	0, // 0: services.savour_rpc.l3staking.L3StakingService.updateStakingNodeIncome:input_type -> services.savour_rpc.l3staking.StakingNodeReq
	1, // 1: services.savour_rpc.l3staking.L3StakingService.updateStakingNodeIncome:output_type -> services.savour_rpc.l3staking.StakingNodeRep
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_savour_rpc_l3staking_proto_init() }
func file_savour_rpc_l3staking_proto_init() {
	if File_savour_rpc_l3staking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_savour_rpc_l3staking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingNodeReq); i {
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
		file_savour_rpc_l3staking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakingNodeRep); i {
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
			RawDescriptor: file_savour_rpc_l3staking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_savour_rpc_l3staking_proto_goTypes,
		DependencyIndexes: file_savour_rpc_l3staking_proto_depIdxs,
		MessageInfos:      file_savour_rpc_l3staking_proto_msgTypes,
	}.Build()
	File_savour_rpc_l3staking_proto = out.File
	file_savour_rpc_l3staking_proto_rawDesc = nil
	file_savour_rpc_l3staking_proto_goTypes = nil
	file_savour_rpc_l3staking_proto_depIdxs = nil
}
