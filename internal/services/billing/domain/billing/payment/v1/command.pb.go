// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: domain/billing/payment/v1/command.proto

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

type Command int32

const (
	Command_COMMAND_UNSPECIFIED     Command = 0
	Command_COMMAND_PAYMENT_CREATE  Command = 1
	Command_COMMAND_PAYMENT_APPROVE Command = 2
	Command_COMMAND_PAYMENT_CLOSE   Command = 3
	Command_COMMAND_PAYMENT_REJECT  Command = 4
	Command_COMMAND_BALANCE_UPDATE  Command = 5
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "COMMAND_UNSPECIFIED",
		1: "COMMAND_PAYMENT_CREATE",
		2: "COMMAND_PAYMENT_APPROVE",
		3: "COMMAND_PAYMENT_CLOSE",
		4: "COMMAND_PAYMENT_REJECT",
		5: "COMMAND_BALANCE_UPDATE",
	}
	Command_value = map[string]int32{
		"COMMAND_UNSPECIFIED":     0,
		"COMMAND_PAYMENT_CREATE":  1,
		"COMMAND_PAYMENT_APPROVE": 2,
		"COMMAND_PAYMENT_CLOSE":   3,
		"COMMAND_PAYMENT_REJECT":  4,
		"COMMAND_BALANCE_UPDATE":  5,
	}
)

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}

func (x Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_billing_payment_v1_command_proto_enumTypes[0].Descriptor()
}

func (Command) Type() protoreflect.EnumType {
	return &file_domain_billing_payment_v1_command_proto_enumTypes[0]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command.Descriptor instead.
func (Command) EnumDescriptor() ([]byte, []int) {
	return file_domain_billing_payment_v1_command_proto_rawDescGZIP(), []int{0}
}

var File_domain_billing_payment_v1_command_proto protoreflect.FileDescriptor

var file_domain_billing_payment_v1_command_proto_rawDesc = []byte{
	0x0a, 0x27, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xae, 0x01,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4d,
	0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x41,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e,
	0x44, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x42, 0x41,
	0x4c, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x05, 0x42, 0x52,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74,
	0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_billing_payment_v1_command_proto_rawDescOnce sync.Once
	file_domain_billing_payment_v1_command_proto_rawDescData = file_domain_billing_payment_v1_command_proto_rawDesc
)

func file_domain_billing_payment_v1_command_proto_rawDescGZIP() []byte {
	file_domain_billing_payment_v1_command_proto_rawDescOnce.Do(func() {
		file_domain_billing_payment_v1_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_billing_payment_v1_command_proto_rawDescData)
	})
	return file_domain_billing_payment_v1_command_proto_rawDescData
}

var file_domain_billing_payment_v1_command_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_billing_payment_v1_command_proto_goTypes = []interface{}{
	(Command)(0), // 0: domain.billing.payment.v1.Command
}
var file_domain_billing_payment_v1_command_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_billing_payment_v1_command_proto_init() }
func file_domain_billing_payment_v1_command_proto_init() {
	if File_domain_billing_payment_v1_command_proto != nil {
		return
	}
	file_domain_billing_payment_v1_payment_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_billing_payment_v1_command_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_billing_payment_v1_command_proto_goTypes,
		DependencyIndexes: file_domain_billing_payment_v1_command_proto_depIdxs,
		EnumInfos:         file_domain_billing_payment_v1_command_proto_enumTypes,
	}.Build()
	File_domain_billing_payment_v1_command_proto = out.File
	file_domain_billing_payment_v1_command_proto_rawDesc = nil
	file_domain_billing_payment_v1_command_proto_goTypes = nil
	file_domain_billing_payment_v1_command_proto_depIdxs = nil
}
