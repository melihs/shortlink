// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: eventsourcing/v1/command.proto

package v1

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

// BaseCommand contains the basic info
// that all commands should have
type BaseCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// AggregateID represents the id of the aggregate to apply to
	AggregateId   string `protobuf:"bytes,2,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateType string `protobuf:"bytes,3,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	Version       int32  `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Payload       string `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *BaseCommand) Reset() {
	*x = BaseCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventsourcing_v1_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseCommand) ProtoMessage() {}

func (x *BaseCommand) ProtoReflect() protoreflect.Message {
	mi := &file_eventsourcing_v1_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseCommand.ProtoReflect.Descriptor instead.
func (*BaseCommand) Descriptor() ([]byte, []int) {
	return file_eventsourcing_v1_command_proto_rawDescGZIP(), []int{0}
}

func (x *BaseCommand) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BaseCommand) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *BaseCommand) GetAggregateType() string {
	if x != nil {
		return x.AggregateType
	}
	return ""
}

func (x *BaseCommand) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BaseCommand) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_eventsourcing_v1_command_proto protoreflect.FileDescriptor

var file_eventsourcing_v1_command_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x22, 0x9f, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventsourcing_v1_command_proto_rawDescOnce sync.Once
	file_eventsourcing_v1_command_proto_rawDescData = file_eventsourcing_v1_command_proto_rawDesc
)

func file_eventsourcing_v1_command_proto_rawDescGZIP() []byte {
	file_eventsourcing_v1_command_proto_rawDescOnce.Do(func() {
		file_eventsourcing_v1_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventsourcing_v1_command_proto_rawDescData)
	})
	return file_eventsourcing_v1_command_proto_rawDescData
}

var file_eventsourcing_v1_command_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eventsourcing_v1_command_proto_goTypes = []interface{}{
	(*BaseCommand)(nil), // 0: eventsourcing.v1.BaseCommand
}
var file_eventsourcing_v1_command_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eventsourcing_v1_command_proto_init() }
func file_eventsourcing_v1_command_proto_init() {
	if File_eventsourcing_v1_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventsourcing_v1_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseCommand); i {
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
			RawDescriptor: file_eventsourcing_v1_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventsourcing_v1_command_proto_goTypes,
		DependencyIndexes: file_eventsourcing_v1_command_proto_depIdxs,
		MessageInfos:      file_eventsourcing_v1_command_proto_msgTypes,
	}.Build()
	File_eventsourcing_v1_command_proto = out.File
	file_eventsourcing_v1_command_proto_rawDesc = nil
	file_eventsourcing_v1_command_proto_goTypes = nil
	file_eventsourcing_v1_command_proto_depIdxs = nil
}
