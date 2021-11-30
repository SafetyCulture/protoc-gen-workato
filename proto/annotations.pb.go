// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: proto/annotations.proto

package workato

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// MethodOptionsWorkatoTrigger
type MethodOptionsWorkatoTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethodOptionsWorkatoTrigger) Reset() {
	*x = MethodOptionsWorkatoTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptionsWorkatoTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptionsWorkatoTrigger) ProtoMessage() {}

func (x *MethodOptionsWorkatoTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_proto_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptionsWorkatoTrigger.ProtoReflect.Descriptor instead.
func (*MethodOptionsWorkatoTrigger) Descriptor() ([]byte, []int) {
	return file_proto_annotations_proto_rawDescGZIP(), []int{0}
}

// MethodOptionsWorkatoTrigger
type MethodOptionsWorkatoPickList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MethodOptionsWorkatoPickList) Reset() {
	*x = MethodOptionsWorkatoPickList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptionsWorkatoPickList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptionsWorkatoPickList) ProtoMessage() {}

func (x *MethodOptionsWorkatoPickList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptionsWorkatoPickList.ProtoReflect.Descriptor instead.
func (*MethodOptionsWorkatoPickList) Descriptor() ([]byte, []int) {
	return file_proto_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *MethodOptionsWorkatoPickList) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MethodOptionsWorkatoPickList) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var file_proto_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptionsWorkatoTrigger)(nil),
		Field:         50009,
		Name:          "s12.protobuf.workato.trigger",
		Tag:           "bytes,50009,opt,name=trigger",
		Filename:      "proto/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptionsWorkatoPickList)(nil),
		Field:         50010,
		Name:          "s12.protobuf.workato.pick_list",
		Tag:           "bytes,50010,opt,name=pick_list",
		Filename:      "proto/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// See `MethodOptionsWorkatoTrigger`.
	//
	// optional s12.protobuf.workato.MethodOptionsWorkatoTrigger trigger = 50009;
	E_Trigger = &file_proto_annotations_proto_extTypes[0]
	// See `MethodOptionsWorkatoPickList`.
	//
	// optional s12.protobuf.workato.MethodOptionsWorkatoPickList pick_list = 50010;
	E_PickList = &file_proto_annotations_proto_extTypes[1]
)

var File_proto_annotations_proto protoreflect.FileDescriptor

var file_proto_annotations_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x31, 0x32, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x6f, 0x72,
	0x6b, 0x61, 0x74, 0x6f, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x1c, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x61, 0x74, 0x6f, 0x50, 0x69, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x6d, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd9, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x31, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74,
	0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57,
	0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x07, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x3a, 0x71, 0x0a, 0x09, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xda, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x31,
	0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x61,
	0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x57, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x50, 0x69, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x08, 0x70, 0x69, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x43, 0x75,
	0x6c, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x77,
	0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_annotations_proto_rawDescOnce sync.Once
	file_proto_annotations_proto_rawDescData = file_proto_annotations_proto_rawDesc
)

func file_proto_annotations_proto_rawDescGZIP() []byte {
	file_proto_annotations_proto_rawDescOnce.Do(func() {
		file_proto_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_annotations_proto_rawDescData)
	})
	return file_proto_annotations_proto_rawDescData
}

var file_proto_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_annotations_proto_goTypes = []interface{}{
	(*MethodOptionsWorkatoTrigger)(nil),  // 0: s12.protobuf.workato.MethodOptionsWorkatoTrigger
	(*MethodOptionsWorkatoPickList)(nil), // 1: s12.protobuf.workato.MethodOptionsWorkatoPickList
	(*descriptorpb.MethodOptions)(nil),   // 2: google.protobuf.MethodOptions
}
var file_proto_annotations_proto_depIdxs = []int32{
	2, // 0: s12.protobuf.workato.trigger:extendee -> google.protobuf.MethodOptions
	2, // 1: s12.protobuf.workato.pick_list:extendee -> google.protobuf.MethodOptions
	0, // 2: s12.protobuf.workato.trigger:type_name -> s12.protobuf.workato.MethodOptionsWorkatoTrigger
	1, // 3: s12.protobuf.workato.pick_list:type_name -> s12.protobuf.workato.MethodOptionsWorkatoPickList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_annotations_proto_init() }
func file_proto_annotations_proto_init() {
	if File_proto_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptionsWorkatoTrigger); i {
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
		file_proto_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptionsWorkatoPickList); i {
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
			RawDescriptor: file_proto_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_proto_annotations_proto_goTypes,
		DependencyIndexes: file_proto_annotations_proto_depIdxs,
		MessageInfos:      file_proto_annotations_proto_msgTypes,
		ExtensionInfos:    file_proto_annotations_proto_extTypes,
	}.Build()
	File_proto_annotations_proto = out.File
	file_proto_annotations_proto_rawDesc = nil
	file_proto_annotations_proto_goTypes = nil
	file_proto_annotations_proto_depIdxs = nil
}
