// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: s12/protobuf/workato/annotations.proto

package workato

import (
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

type PicklistOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PicklistOptions) Reset() {
	*x = PicklistOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s12_protobuf_workato_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PicklistOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PicklistOptions) ProtoMessage() {}

func (x *PicklistOptions) ProtoReflect() protoreflect.Message {
	mi := &file_s12_protobuf_workato_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PicklistOptions.ProtoReflect.Descriptor instead.
func (*PicklistOptions) Descriptor() ([]byte, []int) {
	return file_s12_protobuf_workato_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *PicklistOptions) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PicklistOptions) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// MethodOptionsWorkatoTrigger
type MethodOptionsWorkato struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Trigger  bool             `protobuf:"varint,1,opt,name=trigger,proto3" json:"trigger,omitempty"`
	Picklist *PicklistOptions `protobuf:"bytes,2,opt,name=picklist,proto3" json:"picklist,omitempty"`
	// @deprecated use google.api.method_visibility instead
	//
	// Deprecated: Do not use.
	Excluded bool `protobuf:"varint,3,opt,name=excluded,proto3" json:"excluded,omitempty"`
}

func (x *MethodOptionsWorkato) Reset() {
	*x = MethodOptionsWorkato{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s12_protobuf_workato_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptionsWorkato) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptionsWorkato) ProtoMessage() {}

func (x *MethodOptionsWorkato) ProtoReflect() protoreflect.Message {
	mi := &file_s12_protobuf_workato_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptionsWorkato.ProtoReflect.Descriptor instead.
func (*MethodOptionsWorkato) Descriptor() ([]byte, []int) {
	return file_s12_protobuf_workato_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *MethodOptionsWorkato) GetTrigger() bool {
	if x != nil {
		return x.Trigger
	}
	return false
}

func (x *MethodOptionsWorkato) GetPicklist() *PicklistOptions {
	if x != nil {
		return x.Picklist
	}
	return nil
}

// Deprecated: Do not use.
func (x *MethodOptionsWorkato) GetExcluded() bool {
	if x != nil {
		return x.Excluded
	}
	return false
}

// FieldOptionsWorkato
type FieldOptionsWorkato struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DynamicPicklist string `protobuf:"bytes,1,opt,name=dynamic_picklist,json=dynamicPicklist,proto3" json:"dynamic_picklist,omitempty"`
	// @deprecated use google.api.field_visibility instead
	//
	// Deprecated: Do not use.
	Excluded bool   `protobuf:"varint,2,opt,name=excluded,proto3" json:"excluded,omitempty"`
	Picklist string `protobuf:"bytes,3,opt,name=picklist,proto3" json:"picklist,omitempty"`
	// marks a field `sticky`
	Important bool `protobuf:"varint,4,opt,name=important,proto3" json:"important,omitempty"`
}

func (x *FieldOptionsWorkato) Reset() {
	*x = FieldOptionsWorkato{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s12_protobuf_workato_annotations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptionsWorkato) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptionsWorkato) ProtoMessage() {}

func (x *FieldOptionsWorkato) ProtoReflect() protoreflect.Message {
	mi := &file_s12_protobuf_workato_annotations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptionsWorkato.ProtoReflect.Descriptor instead.
func (*FieldOptionsWorkato) Descriptor() ([]byte, []int) {
	return file_s12_protobuf_workato_annotations_proto_rawDescGZIP(), []int{2}
}

func (x *FieldOptionsWorkato) GetDynamicPicklist() string {
	if x != nil {
		return x.DynamicPicklist
	}
	return ""
}

// Deprecated: Do not use.
func (x *FieldOptionsWorkato) GetExcluded() bool {
	if x != nil {
		return x.Excluded
	}
	return false
}

func (x *FieldOptionsWorkato) GetPicklist() string {
	if x != nil {
		return x.Picklist
	}
	return ""
}

func (x *FieldOptionsWorkato) GetImportant() bool {
	if x != nil {
		return x.Important
	}
	return false
}

var file_s12_protobuf_workato_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptionsWorkato)(nil),
		Field:         50009,
		Name:          "s12.protobuf.workato.method",
		Tag:           "bytes,50009,opt,name=method",
		Filename:      "s12/protobuf/workato/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptionsWorkato)(nil),
		Field:         50009,
		Name:          "s12.protobuf.workato.field",
		Tag:           "bytes,50009,opt,name=field",
		Filename:      "s12/protobuf/workato/annotations.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// See `MethodOptionsWorkato`.
	//
	// optional s12.protobuf.workato.MethodOptionsWorkato method = 50009;
	E_Method = &file_s12_protobuf_workato_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// See `FieldOptionsWorkato`.
	//
	// optional s12.protobuf.workato.FieldOptionsWorkato field = 50009;
	E_Field = &file_s12_protobuf_workato_annotations_proto_extTypes[1]
)

var File_s12_protobuf_workato_annotations_proto protoreflect.FileDescriptor

var file_s12_protobuf_workato_annotations_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x31, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x31, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0f, 0x50,
	0x69, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x61, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x41, 0x0a,
	0x08, 0x70, 0x69, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x73, 0x31, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x70, 0x69, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x22, 0x9a, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x57, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x50, 0x69, 0x63, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x69, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x74, 0x3a, 0x64, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x73, 0x31, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x3a, 0x60, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x31, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x52, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x43, 0x75, 0x6c, 0x74, 0x75, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x77, 0x6f, 0x72,
	0x6b, 0x61, 0x74, 0x6f, 0x2f, 0x73, 0x31, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74, 0x6f, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x61, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s12_protobuf_workato_annotations_proto_rawDescOnce sync.Once
	file_s12_protobuf_workato_annotations_proto_rawDescData = file_s12_protobuf_workato_annotations_proto_rawDesc
)

func file_s12_protobuf_workato_annotations_proto_rawDescGZIP() []byte {
	file_s12_protobuf_workato_annotations_proto_rawDescOnce.Do(func() {
		file_s12_protobuf_workato_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_s12_protobuf_workato_annotations_proto_rawDescData)
	})
	return file_s12_protobuf_workato_annotations_proto_rawDescData
}

var file_s12_protobuf_workato_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_s12_protobuf_workato_annotations_proto_goTypes = []interface{}{
	(*PicklistOptions)(nil),            // 0: s12.protobuf.workato.PicklistOptions
	(*MethodOptionsWorkato)(nil),       // 1: s12.protobuf.workato.MethodOptionsWorkato
	(*FieldOptionsWorkato)(nil),        // 2: s12.protobuf.workato.FieldOptionsWorkato
	(*descriptorpb.MethodOptions)(nil), // 3: google.protobuf.MethodOptions
	(*descriptorpb.FieldOptions)(nil),  // 4: google.protobuf.FieldOptions
}
var file_s12_protobuf_workato_annotations_proto_depIdxs = []int32{
	0, // 0: s12.protobuf.workato.MethodOptionsWorkato.picklist:type_name -> s12.protobuf.workato.PicklistOptions
	3, // 1: s12.protobuf.workato.method:extendee -> google.protobuf.MethodOptions
	4, // 2: s12.protobuf.workato.field:extendee -> google.protobuf.FieldOptions
	1, // 3: s12.protobuf.workato.method:type_name -> s12.protobuf.workato.MethodOptionsWorkato
	2, // 4: s12.protobuf.workato.field:type_name -> s12.protobuf.workato.FieldOptionsWorkato
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_s12_protobuf_workato_annotations_proto_init() }
func file_s12_protobuf_workato_annotations_proto_init() {
	if File_s12_protobuf_workato_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s12_protobuf_workato_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PicklistOptions); i {
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
		file_s12_protobuf_workato_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptionsWorkato); i {
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
		file_s12_protobuf_workato_annotations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptionsWorkato); i {
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
			RawDescriptor: file_s12_protobuf_workato_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_s12_protobuf_workato_annotations_proto_goTypes,
		DependencyIndexes: file_s12_protobuf_workato_annotations_proto_depIdxs,
		MessageInfos:      file_s12_protobuf_workato_annotations_proto_msgTypes,
		ExtensionInfos:    file_s12_protobuf_workato_annotations_proto_extTypes,
	}.Build()
	File_s12_protobuf_workato_annotations_proto = out.File
	file_s12_protobuf_workato_annotations_proto_rawDesc = nil
	file_s12_protobuf_workato_annotations_proto_goTypes = nil
	file_s12_protobuf_workato_annotations_proto_depIdxs = nil
}
