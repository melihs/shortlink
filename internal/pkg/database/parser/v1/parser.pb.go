// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: database/parser/v1/parser.proto

package v1

import (
	v1 "github.com/batazor/shortlink/internal/pkg/database/query/v1"
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

type Step int32

const (
	Step_STEP_UNSPECIFIED                               Step = 0
	Step_STEP_SELECT_FIELD                              Step = 1
	Step_STEP_SELECT_FROM                               Step = 2
	Step_STEP_SELECT_COMMA                              Step = 3
	Step_STEP_SELECT_FROM_TABLE                         Step = 4
	Step_STEP_INSERT_TABLE                              Step = 5
	Step_STEP_INSERT_FIELD_OPENING_PARENTS              Step = 6
	Step_STEP_INSERT_FIELDS                             Step = 7
	Step_STEP_INSERT_FIELDS_COMMA_OR_CLOSING_PARENTS    Step = 8
	Step_STEP_INSERT_VALUES_OPENING_PARENS              Step = 9
	Step_STEP_INSERT_RWORD                              Step = 10
	Step_STEP_INSERT_VALUES                             Step = 11
	Step_STEP_INSERT_VALUES_COMMA_OR_CLOSING_PARENS     Step = 12
	Step_STEP_INSERT_VALUES_COMMA_BEFORE_OPENING_PARENS Step = 13
	Step_STEP_UPDATE_TABLE                              Step = 14
	Step_STEP_UPDATE_SET                                Step = 15
	Step_STEP_UPDATE_FIELD                              Step = 16
	Step_STEP_UPDATE_EQUALS                             Step = 17
	Step_STEP_UPDATE_VALUE                              Step = 18
	Step_STEP_UPDATE_COMMA                              Step = 19
	Step_STEP_DELETE_FROM_TABLE                         Step = 20
	Step_STEP_WHERE                                     Step = 21
	Step_STEP_WHERE_FIELD                               Step = 22
	Step_STEP_WHERE_OPERATOR                            Step = 23
	Step_STEP_WHERE_VALUE                               Step = 24
	Step_STEP_WHERE_AND                                 Step = 25
	Step_STEP_ORDER                                     Step = 26
	Step_STEP_ORDER_FIELD                               Step = 27
	Step_STEP_ORDER_DIRECTION_OR_COMMA                  Step = 28
	Step_STEP_JOIN                                      Step = 29
	Step_STEP_JOIN_TABLE                                Step = 30
	Step_STEP_JOIN_CONDITION                            Step = 31
)

// Enum value maps for Step.
var (
	Step_name = map[int32]string{
		0:  "STEP_UNSPECIFIED",
		1:  "STEP_SELECT_FIELD",
		2:  "STEP_SELECT_FROM",
		3:  "STEP_SELECT_COMMA",
		4:  "STEP_SELECT_FROM_TABLE",
		5:  "STEP_INSERT_TABLE",
		6:  "STEP_INSERT_FIELD_OPENING_PARENTS",
		7:  "STEP_INSERT_FIELDS",
		8:  "STEP_INSERT_FIELDS_COMMA_OR_CLOSING_PARENTS",
		9:  "STEP_INSERT_VALUES_OPENING_PARENS",
		10: "STEP_INSERT_RWORD",
		11: "STEP_INSERT_VALUES",
		12: "STEP_INSERT_VALUES_COMMA_OR_CLOSING_PARENS",
		13: "STEP_INSERT_VALUES_COMMA_BEFORE_OPENING_PARENS",
		14: "STEP_UPDATE_TABLE",
		15: "STEP_UPDATE_SET",
		16: "STEP_UPDATE_FIELD",
		17: "STEP_UPDATE_EQUALS",
		18: "STEP_UPDATE_VALUE",
		19: "STEP_UPDATE_COMMA",
		20: "STEP_DELETE_FROM_TABLE",
		21: "STEP_WHERE",
		22: "STEP_WHERE_FIELD",
		23: "STEP_WHERE_OPERATOR",
		24: "STEP_WHERE_VALUE",
		25: "STEP_WHERE_AND",
		26: "STEP_ORDER",
		27: "STEP_ORDER_FIELD",
		28: "STEP_ORDER_DIRECTION_OR_COMMA",
		29: "STEP_JOIN",
		30: "STEP_JOIN_TABLE",
		31: "STEP_JOIN_CONDITION",
	}
	Step_value = map[string]int32{
		"STEP_UNSPECIFIED":                               0,
		"STEP_SELECT_FIELD":                              1,
		"STEP_SELECT_FROM":                               2,
		"STEP_SELECT_COMMA":                              3,
		"STEP_SELECT_FROM_TABLE":                         4,
		"STEP_INSERT_TABLE":                              5,
		"STEP_INSERT_FIELD_OPENING_PARENTS":              6,
		"STEP_INSERT_FIELDS":                             7,
		"STEP_INSERT_FIELDS_COMMA_OR_CLOSING_PARENTS":    8,
		"STEP_INSERT_VALUES_OPENING_PARENS":              9,
		"STEP_INSERT_RWORD":                              10,
		"STEP_INSERT_VALUES":                             11,
		"STEP_INSERT_VALUES_COMMA_OR_CLOSING_PARENS":     12,
		"STEP_INSERT_VALUES_COMMA_BEFORE_OPENING_PARENS": 13,
		"STEP_UPDATE_TABLE":                              14,
		"STEP_UPDATE_SET":                                15,
		"STEP_UPDATE_FIELD":                              16,
		"STEP_UPDATE_EQUALS":                             17,
		"STEP_UPDATE_VALUE":                              18,
		"STEP_UPDATE_COMMA":                              19,
		"STEP_DELETE_FROM_TABLE":                         20,
		"STEP_WHERE":                                     21,
		"STEP_WHERE_FIELD":                               22,
		"STEP_WHERE_OPERATOR":                            23,
		"STEP_WHERE_VALUE":                               24,
		"STEP_WHERE_AND":                                 25,
		"STEP_ORDER":                                     26,
		"STEP_ORDER_FIELD":                               27,
		"STEP_ORDER_DIRECTION_OR_COMMA":                  28,
		"STEP_JOIN":                                      29,
		"STEP_JOIN_TABLE":                                30,
		"STEP_JOIN_CONDITION":                            31,
	}
)

func (x Step) Enum() *Step {
	p := new(Step)
	*p = x
	return p
}

func (x Step) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Step) Descriptor() protoreflect.EnumDescriptor {
	return file_database_parser_v1_parser_proto_enumTypes[0].Descriptor()
}

func (Step) Type() protoreflect.EnumType {
	return &file_database_parser_v1_parser_proto_enumTypes[0]
}

func (x Step) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Step.Descriptor instead.
func (Step) EnumDescriptor() ([]byte, []int) {
	return file_database_parser_v1_parser_proto_rawDescGZIP(), []int{0}
}

type Parser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I               int32     `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	Sql             string    `protobuf:"bytes,2,opt,name=sql,proto3" json:"sql,omitempty"`
	Step            Step      `protobuf:"varint,3,opt,name=step,proto3,enum=database.parser.v1.Step" json:"step,omitempty"`
	Query           *v1.Query `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	Error           string    `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	NextUpdateField string    `protobuf:"bytes,6,opt,name=next_update_field,json=nextUpdateField,proto3" json:"next_update_field,omitempty"`
}

func (x *Parser) Reset() {
	*x = Parser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_database_parser_v1_parser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Parser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parser) ProtoMessage() {}

func (x *Parser) ProtoReflect() protoreflect.Message {
	mi := &file_database_parser_v1_parser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parser.ProtoReflect.Descriptor instead.
func (*Parser) Descriptor() ([]byte, []int) {
	return file_database_parser_v1_parser_proto_rawDescGZIP(), []int{0}
}

func (x *Parser) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

func (x *Parser) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

func (x *Parser) GetStep() Step {
	if x != nil {
		return x.Step
	}
	return Step_STEP_UNSPECIFIED
}

func (x *Parser) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *Parser) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Parser) GetNextUpdateField() string {
	if x != nil {
		return x.NextUpdateField
	}
	return ""
}

var File_database_parser_v1_parser_proto protoreflect.FileDescriptor

var file_database_parser_v1_parser_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12,
	0x0c, 0x0a, 0x01, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x69, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x71, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12,
	0x2c, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x2e, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a,
	0xd1, 0x06, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x45, 0x50,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x53, 0x45,
	0x4c, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x54, 0x45, 0x50, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43,
	0x54, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x15,
	0x0a, 0x11, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x54, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e,
	0x53, 0x45, 0x52, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x53, 0x10, 0x07, 0x12, 0x2f, 0x0a, 0x2b, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e, 0x53,
	0x45, 0x52, 0x54, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x5f, 0x4f, 0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x45,
	0x4e, 0x54, 0x53, 0x10, 0x08, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e,
	0x53, 0x45, 0x52, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x5f, 0x4f, 0x50, 0x45, 0x4e,
	0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x53, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11,
	0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x52, 0x57, 0x4f, 0x52,
	0x44, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x45,
	0x52, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x10, 0x0b, 0x12, 0x2e, 0x0a, 0x2a, 0x53,
	0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x5f, 0x4f, 0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x49,
	0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x53, 0x10, 0x0c, 0x12, 0x32, 0x0a, 0x2e, 0x53,
	0x54, 0x45, 0x50, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x5f, 0x42, 0x45, 0x46, 0x4f, 0x52, 0x45, 0x5f, 0x4f,
	0x50, 0x45, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x53, 0x10, 0x0d, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x0f, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x54, 0x45, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x11, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54,
	0x45, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10,
	0x12, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x10, 0x13, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x45, 0x50,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x14, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x57, 0x48, 0x45,
	0x52, 0x45, 0x10, 0x15, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x57, 0x48, 0x45,
	0x52, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x16, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54,
	0x45, 0x50, 0x5f, 0x57, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f,
	0x52, 0x10, 0x17, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x57, 0x48, 0x45, 0x52,
	0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x18, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x45,
	0x50, 0x5f, 0x57, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x10, 0x19, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x1a, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x10, 0x1b, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x10, 0x1c, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x4a,
	0x4f, 0x49, 0x4e, 0x10, 0x1d, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x45, 0x50, 0x5f, 0x4a, 0x4f,
	0x49, 0x4e, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x1e, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54,
	0x45, 0x50, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x1f, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c,
	0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_database_parser_v1_parser_proto_rawDescOnce sync.Once
	file_database_parser_v1_parser_proto_rawDescData = file_database_parser_v1_parser_proto_rawDesc
)

func file_database_parser_v1_parser_proto_rawDescGZIP() []byte {
	file_database_parser_v1_parser_proto_rawDescOnce.Do(func() {
		file_database_parser_v1_parser_proto_rawDescData = protoimpl.X.CompressGZIP(file_database_parser_v1_parser_proto_rawDescData)
	})
	return file_database_parser_v1_parser_proto_rawDescData
}

var file_database_parser_v1_parser_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_database_parser_v1_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_database_parser_v1_parser_proto_goTypes = []interface{}{
	(Step)(0),        // 0: database.parser.v1.Step
	(*Parser)(nil),   // 1: database.parser.v1.Parser
	(*v1.Query)(nil), // 2: database.query.v1.Query
}
var file_database_parser_v1_parser_proto_depIdxs = []int32{
	0, // 0: database.parser.v1.Parser.step:type_name -> database.parser.v1.Step
	2, // 1: database.parser.v1.Parser.query:type_name -> database.query.v1.Query
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_database_parser_v1_parser_proto_init() }
func file_database_parser_v1_parser_proto_init() {
	if File_database_parser_v1_parser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_database_parser_v1_parser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Parser); i {
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
			RawDescriptor: file_database_parser_v1_parser_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_database_parser_v1_parser_proto_goTypes,
		DependencyIndexes: file_database_parser_v1_parser_proto_depIdxs,
		EnumInfos:         file_database_parser_v1_parser_proto_enumTypes,
		MessageInfos:      file_database_parser_v1_parser_proto_msgTypes,
	}.Build()
	File_database_parser_v1_parser_proto = out.File
	file_database_parser_v1_parser_proto_rawDesc = nil
	file_database_parser_v1_parser_proto_goTypes = nil
	file_database_parser_v1_parser_proto_depIdxs = nil
}
