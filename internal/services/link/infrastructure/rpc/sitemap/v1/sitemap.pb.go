// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: infrastructure/rpc/sitemap/v1/sitemap.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ParseRequest) Reset() {
	*x = ParseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_sitemap_v1_sitemap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseRequest) ProtoMessage() {}

func (x *ParseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_sitemap_v1_sitemap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseRequest.ProtoReflect.Descriptor instead.
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescGZIP(), []int{0}
}

func (x *ParseRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_infrastructure_rpc_sitemap_v1_sitemap_proto protoreflect.FileDescriptor

var file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x50, 0x61, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x60, 0x0a, 0x0e, 0x53,
	0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a,
	0x05, 0x50, 0x61, 0x72, 0x73, 0x65, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x69, 0x74, 0x65,
	0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x53, 0x5a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61,
	0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescOnce sync.Once
	file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescData = file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDesc
)

func file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescGZIP() []byte {
	file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescOnce.Do(func() {
		file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescData)
	})
	return file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDescData
}

var file_infrastructure_rpc_sitemap_v1_sitemap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infrastructure_rpc_sitemap_v1_sitemap_proto_goTypes = []interface{}{
	(*ParseRequest)(nil),  // 0: infrastructure.rpc.sitemap.v1.ParseRequest
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_infrastructure_rpc_sitemap_v1_sitemap_proto_depIdxs = []int32{
	0, // 0: infrastructure.rpc.sitemap.v1.SitemapService.Parse:input_type -> infrastructure.rpc.sitemap.v1.ParseRequest
	1, // 1: infrastructure.rpc.sitemap.v1.SitemapService.Parse:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infrastructure_rpc_sitemap_v1_sitemap_proto_init() }
func file_infrastructure_rpc_sitemap_v1_sitemap_proto_init() {
	if File_infrastructure_rpc_sitemap_v1_sitemap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_rpc_sitemap_v1_sitemap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseRequest); i {
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
			RawDescriptor: file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_rpc_sitemap_v1_sitemap_proto_goTypes,
		DependencyIndexes: file_infrastructure_rpc_sitemap_v1_sitemap_proto_depIdxs,
		MessageInfos:      file_infrastructure_rpc_sitemap_v1_sitemap_proto_msgTypes,
	}.Build()
	File_infrastructure_rpc_sitemap_v1_sitemap_proto = out.File
	file_infrastructure_rpc_sitemap_v1_sitemap_proto_rawDesc = nil
	file_infrastructure_rpc_sitemap_v1_sitemap_proto_goTypes = nil
	file_infrastructure_rpc_sitemap_v1_sitemap_proto_depIdxs = nil
}
