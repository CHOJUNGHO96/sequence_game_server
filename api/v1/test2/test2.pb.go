// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.3
// source: test2.proto

package test2

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

type Test2Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestIntData int32 `protobuf:"varint,1,opt,name=test_int_data,json=testIntData,proto3" json:"test_int_data,omitempty"`
}

func (x *Test2Request) Reset() {
	*x = Test2Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2Request) ProtoMessage() {}

func (x *Test2Request) ProtoReflect() protoreflect.Message {
	mi := &file_test2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2Request.ProtoReflect.Descriptor instead.
func (*Test2Request) Descriptor() ([]byte, []int) {
	return file_test2_proto_rawDescGZIP(), []int{0}
}

func (x *Test2Request) GetTestIntData() int32 {
	if x != nil {
		return x.TestIntData
	}
	return 0
}

type Test2Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestMapData map[string]string `protobuf:"bytes,1,rep,name=test_map_data,json=testMapData,proto3" json:"test_map_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Test2Response) Reset() {
	*x = Test2Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2Response) ProtoMessage() {}

func (x *Test2Response) ProtoReflect() protoreflect.Message {
	mi := &file_test2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2Response.ProtoReflect.Descriptor instead.
func (*Test2Response) Descriptor() ([]byte, []int) {
	return file_test2_proto_rawDescGZIP(), []int{1}
}

func (x *Test2Response) GetTestMapData() map[string]string {
	if x != nil {
		return x.TestMapData
	}
	return nil
}

var File_test2_proto protoreflect.FileDescriptor

var file_test2_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x65, 0x73, 0x74, 0x32, 0x22, 0x32, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x73,
	0x74, 0x49, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x73,
	0x74, 0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0d, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x61,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x3e, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x44, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x32, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0x12, 0x13,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x32, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test2_proto_rawDescOnce sync.Once
	file_test2_proto_rawDescData = file_test2_proto_rawDesc
)

func file_test2_proto_rawDescGZIP() []byte {
	file_test2_proto_rawDescOnce.Do(func() {
		file_test2_proto_rawDescData = protoimpl.X.CompressGZIP(file_test2_proto_rawDescData)
	})
	return file_test2_proto_rawDescData
}

var file_test2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_test2_proto_goTypes = []interface{}{
	(*Test2Request)(nil),  // 0: test2.Test2Request
	(*Test2Response)(nil), // 1: test2.Test2Response
	nil,                   // 2: test2.Test2Response.TestMapDataEntry
}
var file_test2_proto_depIdxs = []int32{
	2, // 0: test2.Test2Response.test_map_data:type_name -> test2.Test2Response.TestMapDataEntry
	0, // 1: test2.Test2Service.Test2:input_type -> test2.Test2Request
	1, // 2: test2.Test2Service.Test2:output_type -> test2.Test2Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test2_proto_init() }
func file_test2_proto_init() {
	if File_test2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2Request); i {
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
		file_test2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2Response); i {
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
			RawDescriptor: file_test2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test2_proto_goTypes,
		DependencyIndexes: file_test2_proto_depIdxs,
		MessageInfos:      file_test2_proto_msgTypes,
	}.Build()
	File_test2_proto = out.File
	file_test2_proto_rawDesc = nil
	file_test2_proto_goTypes = nil
	file_test2_proto_depIdxs = nil
}
