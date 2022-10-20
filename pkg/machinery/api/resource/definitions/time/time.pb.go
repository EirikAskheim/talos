// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: resource/definitions/time/time.proto

package time

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// StatusSpec describes time sync state.
type StatusSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Synced       bool  `protobuf:"varint,1,opt,name=synced,proto3" json:"synced,omitempty"`
	Epoch        int64 `protobuf:"varint,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	SyncDisabled bool  `protobuf:"varint,3,opt,name=sync_disabled,json=syncDisabled,proto3" json:"sync_disabled,omitempty"`
}

func (x *StatusSpec) Reset() {
	*x = StatusSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resource_definitions_time_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusSpec) ProtoMessage() {}

func (x *StatusSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_time_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusSpec.ProtoReflect.Descriptor instead.
func (*StatusSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_time_time_proto_rawDescGZIP(), []int{0}
}

func (x *StatusSpec) GetSynced() bool {
	if x != nil {
		return x.Synced
	}
	return false
}

func (x *StatusSpec) GetEpoch() int64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *StatusSpec) GetSyncDisabled() bool {
	if x != nil {
		return x.SyncDisabled
	}
	return false
}

var File_resource_definitions_time_time_proto protoreflect.FileDescriptor

var file_resource_definitions_time_time_proto_rawDesc = []byte{
	0x0a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2d, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x2f, 0x74, 0x61, 0x6c, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resource_definitions_time_time_proto_rawDescOnce sync.Once
	file_resource_definitions_time_time_proto_rawDescData = file_resource_definitions_time_time_proto_rawDesc
)

func file_resource_definitions_time_time_proto_rawDescGZIP() []byte {
	file_resource_definitions_time_time_proto_rawDescOnce.Do(func() {
		file_resource_definitions_time_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_resource_definitions_time_time_proto_rawDescData)
	})
	return file_resource_definitions_time_time_proto_rawDescData
}

var file_resource_definitions_time_time_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resource_definitions_time_time_proto_goTypes = []interface{}{
	(*StatusSpec)(nil), // 0: talos.resource.definitions.time.StatusSpec
}
var file_resource_definitions_time_time_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resource_definitions_time_time_proto_init() }
func file_resource_definitions_time_time_proto_init() {
	if File_resource_definitions_time_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resource_definitions_time_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusSpec); i {
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
			RawDescriptor: file_resource_definitions_time_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_time_time_proto_goTypes,
		DependencyIndexes: file_resource_definitions_time_time_proto_depIdxs,
		MessageInfos:      file_resource_definitions_time_time_proto_msgTypes,
	}.Build()
	File_resource_definitions_time_time_proto = out.File
	file_resource_definitions_time_time_proto_rawDesc = nil
	file_resource_definitions_time_time_proto_goTypes = nil
	file_resource_definitions_time_time_proto_depIdxs = nil
}
