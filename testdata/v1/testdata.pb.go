// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: testdata.proto

package testdata

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

type GlobalEnum int32

const (
	GlobalEnum_DEFAULT  GlobalEnum = 0
	GlobalEnum_OPTION_A GlobalEnum = 1
)

// Enum value maps for GlobalEnum.
var (
	GlobalEnum_name = map[int32]string{
		0: "DEFAULT",
		1: "OPTION_A",
	}
	GlobalEnum_value = map[string]int32{
		"DEFAULT":  0,
		"OPTION_A": 1,
	}
)

func (x GlobalEnum) Enum() *GlobalEnum {
	p := new(GlobalEnum)
	*p = x
	return p
}

func (x GlobalEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GlobalEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_proto_enumTypes[0].Descriptor()
}

func (GlobalEnum) Type() protoreflect.EnumType {
	return &file_testdata_proto_enumTypes[0]
}

func (x GlobalEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GlobalEnum.Descriptor instead.
func (GlobalEnum) EnumDescriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0}
}

type TestConfig_InternalEnum int32

const (
	TestConfig_DEFAULT  TestConfig_InternalEnum = 0
	TestConfig_OPTION_A TestConfig_InternalEnum = 1
)

// Enum value maps for TestConfig_InternalEnum.
var (
	TestConfig_InternalEnum_name = map[int32]string{
		0: "DEFAULT",
		1: "OPTION_A",
	}
	TestConfig_InternalEnum_value = map[string]int32{
		"DEFAULT":  0,
		"OPTION_A": 1,
	}
)

func (x TestConfig_InternalEnum) Enum() *TestConfig_InternalEnum {
	p := new(TestConfig_InternalEnum)
	*p = x
	return p
}

func (x TestConfig_InternalEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestConfig_InternalEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_testdata_proto_enumTypes[1].Descriptor()
}

func (TestConfig_InternalEnum) Type() protoreflect.EnumType {
	return &file_testdata_proto_enumTypes[1]
}

func (x TestConfig_InternalEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestConfig_InternalEnum.Descriptor instead.
func (TestConfig_InternalEnum) EnumDescriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0, 0}
}

type TestConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str              string                    `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	StrArr           []string                  `protobuf:"bytes,2,rep,name=str_arr,json=strArr,proto3" json:"str_arr,omitempty"`
	Numeric32        int32                     `protobuf:"varint,3,opt,name=numeric32,proto3" json:"numeric32,omitempty"`
	Numeric32Arr     []int32                   `protobuf:"varint,4,rep,packed,name=numeric32_arr,json=numeric32Arr,proto3" json:"numeric32_arr,omitempty"`
	Numeric64        int64                     `protobuf:"varint,5,opt,name=numeric64,proto3" json:"numeric64,omitempty"`
	Numeric64Arr     []int64                   `protobuf:"varint,6,rep,packed,name=numeric64_arr,json=numeric64Arr,proto3" json:"numeric64_arr,omitempty"`
	Unsigned32       uint32                    `protobuf:"varint,7,opt,name=unsigned32,proto3" json:"unsigned32,omitempty"`
	Unsigned32Arr    []uint32                  `protobuf:"varint,8,rep,packed,name=unsigned32_arr,json=unsigned32Arr,proto3" json:"unsigned32_arr,omitempty"`
	Unsigned64       uint64                    `protobuf:"varint,9,opt,name=unsigned64,proto3" json:"unsigned64,omitempty"`
	Unsigned64Arr    []uint64                  `protobuf:"varint,10,rep,packed,name=unsigned64_arr,json=unsigned64Arr,proto3" json:"unsigned64_arr,omitempty"`
	FloatingPoint    float32                   `protobuf:"fixed32,11,opt,name=floating_point,json=floatingPoint,proto3" json:"floating_point,omitempty"`
	FloatingPointArr []float32                 `protobuf:"fixed32,12,rep,packed,name=floating_point_arr,json=floatingPointArr,proto3" json:"floating_point_arr,omitempty"`
	GlobalEnum       GlobalEnum                `protobuf:"varint,13,opt,name=global_enum,json=globalEnum,proto3,enum=libprotoconf.testdata.v1.GlobalEnum" json:"global_enum,omitempty"`
	GlobalEnumArr    []GlobalEnum              `protobuf:"varint,14,rep,packed,name=global_enum_arr,json=globalEnumArr,proto3,enum=libprotoconf.testdata.v1.GlobalEnum" json:"global_enum_arr,omitempty"`
	InternalEnum     TestConfig_InternalEnum   `protobuf:"varint,15,opt,name=internal_enum,json=internalEnum,proto3,enum=libprotoconf.testdata.v1.TestConfig_InternalEnum" json:"internal_enum,omitempty"`
	InternalEnumArr  []TestConfig_InternalEnum `protobuf:"varint,16,rep,packed,name=internal_enum_arr,json=internalEnumArr,proto3,enum=libprotoconf.testdata.v1.TestConfig_InternalEnum" json:"internal_enum_arr,omitempty"`
	SubMessage       *TestConfig_SubMessage    `protobuf:"bytes,17,opt,name=sub_message,json=subMessage,proto3" json:"sub_message,omitempty"`
	SubMessageArr    []*TestConfig_SubMessage  `protobuf:"bytes,18,rep,name=sub_message_arr,json=subMessageArr,proto3" json:"sub_message_arr,omitempty"`
}

func (x *TestConfig) Reset() {
	*x = TestConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConfig) ProtoMessage() {}

func (x *TestConfig) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConfig.ProtoReflect.Descriptor instead.
func (*TestConfig) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0}
}

func (x *TestConfig) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *TestConfig) GetStrArr() []string {
	if x != nil {
		return x.StrArr
	}
	return nil
}

func (x *TestConfig) GetNumeric32() int32 {
	if x != nil {
		return x.Numeric32
	}
	return 0
}

func (x *TestConfig) GetNumeric32Arr() []int32 {
	if x != nil {
		return x.Numeric32Arr
	}
	return nil
}

func (x *TestConfig) GetNumeric64() int64 {
	if x != nil {
		return x.Numeric64
	}
	return 0
}

func (x *TestConfig) GetNumeric64Arr() []int64 {
	if x != nil {
		return x.Numeric64Arr
	}
	return nil
}

func (x *TestConfig) GetUnsigned32() uint32 {
	if x != nil {
		return x.Unsigned32
	}
	return 0
}

func (x *TestConfig) GetUnsigned32Arr() []uint32 {
	if x != nil {
		return x.Unsigned32Arr
	}
	return nil
}

func (x *TestConfig) GetUnsigned64() uint64 {
	if x != nil {
		return x.Unsigned64
	}
	return 0
}

func (x *TestConfig) GetUnsigned64Arr() []uint64 {
	if x != nil {
		return x.Unsigned64Arr
	}
	return nil
}

func (x *TestConfig) GetFloatingPoint() float32 {
	if x != nil {
		return x.FloatingPoint
	}
	return 0
}

func (x *TestConfig) GetFloatingPointArr() []float32 {
	if x != nil {
		return x.FloatingPointArr
	}
	return nil
}

func (x *TestConfig) GetGlobalEnum() GlobalEnum {
	if x != nil {
		return x.GlobalEnum
	}
	return GlobalEnum_DEFAULT
}

func (x *TestConfig) GetGlobalEnumArr() []GlobalEnum {
	if x != nil {
		return x.GlobalEnumArr
	}
	return nil
}

func (x *TestConfig) GetInternalEnum() TestConfig_InternalEnum {
	if x != nil {
		return x.InternalEnum
	}
	return TestConfig_DEFAULT
}

func (x *TestConfig) GetInternalEnumArr() []TestConfig_InternalEnum {
	if x != nil {
		return x.InternalEnumArr
	}
	return nil
}

func (x *TestConfig) GetSubMessage() *TestConfig_SubMessage {
	if x != nil {
		return x.SubMessage
	}
	return nil
}

func (x *TestConfig) GetSubMessageArr() []*TestConfig_SubMessage {
	if x != nil {
		return x.SubMessageArr
	}
	return nil
}

type TestConfig_SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // TestConfig recusrive = 2;
}

func (x *TestConfig_SubMessage) Reset() {
	*x = TestConfig_SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConfig_SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConfig_SubMessage) ProtoMessage() {}

func (x *TestConfig_SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConfig_SubMessage.ProtoReflect.Descriptor instead.
func (*TestConfig_SubMessage) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TestConfig_SubMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_testdata_proto protoreflect.FileDescriptor

var file_testdata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x22, 0xe4, 0x07, 0x0a, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x73,
	0x74, 0x72, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x41, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x33,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63,
	0x33, 0x32, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x33, 0x32, 0x5f,
	0x61, 0x72, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x33, 0x32, 0x41, 0x72, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x36, 0x34, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63,
	0x36, 0x34, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x36, 0x34, 0x41, 0x72, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0d, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x33, 0x32, 0x41, 0x72,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x36, 0x34, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x36,
	0x34, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x36, 0x34, 0x5f,
	0x61, 0x72, 0x72, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0d, 0x75, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x36, 0x34, 0x41, 0x72, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0d, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x02, 0x52, 0x10, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x12, 0x45, 0x0a,
	0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0a, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x4c, 0x0a, 0x0f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x0d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x41,
	0x72, 0x72, 0x12, 0x56, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6c, 0x69, 0x62, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x5d, 0x0a, 0x11, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x61, 0x72, 0x72, 0x18,
	0x10, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x41, 0x72, 0x72, 0x12, 0x50, 0x0a, 0x0b, 0x73, 0x75, 0x62,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x73,
	0x75, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x12,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x72, 0x72, 0x1a, 0x20, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x10,
	0x01, 0x2a, 0x27, 0x0a, 0x0a, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x10, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6e, 0x66, 0x2f, 0x6c, 0x69, 0x62, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_proto_rawDescOnce sync.Once
	file_testdata_proto_rawDescData = file_testdata_proto_rawDesc
)

func file_testdata_proto_rawDescGZIP() []byte {
	file_testdata_proto_rawDescOnce.Do(func() {
		file_testdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_proto_rawDescData)
	})
	return file_testdata_proto_rawDescData
}

var file_testdata_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_testdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testdata_proto_goTypes = []interface{}{
	(GlobalEnum)(0),               // 0: libprotoconf.testdata.v1.GlobalEnum
	(TestConfig_InternalEnum)(0),  // 1: libprotoconf.testdata.v1.TestConfig.InternalEnum
	(*TestConfig)(nil),            // 2: libprotoconf.testdata.v1.TestConfig
	(*TestConfig_SubMessage)(nil), // 3: libprotoconf.testdata.v1.TestConfig.SubMessage
}
var file_testdata_proto_depIdxs = []int32{
	0, // 0: libprotoconf.testdata.v1.TestConfig.global_enum:type_name -> libprotoconf.testdata.v1.GlobalEnum
	0, // 1: libprotoconf.testdata.v1.TestConfig.global_enum_arr:type_name -> libprotoconf.testdata.v1.GlobalEnum
	1, // 2: libprotoconf.testdata.v1.TestConfig.internal_enum:type_name -> libprotoconf.testdata.v1.TestConfig.InternalEnum
	1, // 3: libprotoconf.testdata.v1.TestConfig.internal_enum_arr:type_name -> libprotoconf.testdata.v1.TestConfig.InternalEnum
	3, // 4: libprotoconf.testdata.v1.TestConfig.sub_message:type_name -> libprotoconf.testdata.v1.TestConfig.SubMessage
	3, // 5: libprotoconf.testdata.v1.TestConfig.sub_message_arr:type_name -> libprotoconf.testdata.v1.TestConfig.SubMessage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_testdata_proto_init() }
func file_testdata_proto_init() {
	if File_testdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConfig); i {
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
		file_testdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConfig_SubMessage); i {
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
			RawDescriptor: file_testdata_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_proto_goTypes,
		DependencyIndexes: file_testdata_proto_depIdxs,
		EnumInfos:         file_testdata_proto_enumTypes,
		MessageInfos:      file_testdata_proto_msgTypes,
	}.Build()
	File_testdata_proto = out.File
	file_testdata_proto_rawDesc = nil
	file_testdata_proto_goTypes = nil
	file_testdata_proto_depIdxs = nil
}
