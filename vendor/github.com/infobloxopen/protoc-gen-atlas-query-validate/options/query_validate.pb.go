// Code generated by protoc-gen-gogo.
// source: options/query_validate.proto
// DO NOT EDIT!

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	options/query_validate.proto

It has these top-level messages:
	QueryValidate
	MessageQueryValidate
*/
package options

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type QueryValidate_FilterOperator int32

const (
	QueryValidate_EQ    QueryValidate_FilterOperator = 0
	QueryValidate_MATCH QueryValidate_FilterOperator = 1
	QueryValidate_GT    QueryValidate_FilterOperator = 2
	QueryValidate_GE    QueryValidate_FilterOperator = 3
	QueryValidate_LT    QueryValidate_FilterOperator = 4
	QueryValidate_LE    QueryValidate_FilterOperator = 5
	QueryValidate_ALL   QueryValidate_FilterOperator = 6
	QueryValidate_IEQ   QueryValidate_FilterOperator = 7
	QueryValidate_IN    QueryValidate_FilterOperator = 8
)

var QueryValidate_FilterOperator_name = map[int32]string{
	0: "EQ",
	1: "MATCH",
	2: "GT",
	3: "GE",
	4: "LT",
	5: "LE",
	6: "ALL",
	7: "IEQ",
	8: "IN",
}
var QueryValidate_FilterOperator_value = map[string]int32{
	"EQ":    0,
	"MATCH": 1,
	"GT":    2,
	"GE":    3,
	"LT":    4,
	"LE":    5,
	"ALL":   6,
	"IEQ":   7,
	"IN":    8,
}

func (x QueryValidate_FilterOperator) String() string {
	return proto.EnumName(QueryValidate_FilterOperator_name, int32(x))
}
func (QueryValidate_FilterOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 0}
}

type QueryValidate_ValueType int32

const (
	QueryValidate_DEFAULT QueryValidate_ValueType = 0
	QueryValidate_STRING  QueryValidate_ValueType = 1
	QueryValidate_NUMBER  QueryValidate_ValueType = 2
	QueryValidate_BOOL    QueryValidate_ValueType = 3
)

var QueryValidate_ValueType_name = map[int32]string{
	0: "DEFAULT",
	1: "STRING",
	2: "NUMBER",
	3: "BOOL",
}
var QueryValidate_ValueType_value = map[string]int32{
	"DEFAULT": 0,
	"STRING":  1,
	"NUMBER":  2,
	"BOOL":    3,
}

func (x QueryValidate_ValueType) String() string {
	return proto.EnumName(QueryValidate_ValueType_name, int32(x))
}
func (QueryValidate_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 1}
}

type QueryValidate struct {
	Filtering          *QueryValidate_Filtering      `protobuf:"bytes,1,opt,name=filtering" json:"filtering,omitempty"`
	Sorting            *QueryValidate_Sorting        `protobuf:"bytes,2,opt,name=sorting" json:"sorting,omitempty"`
	FieldSelection     *QueryValidate_FieldSelection `protobuf:"bytes,3,opt,name=field_selection,json=fieldSelection" json:"field_selection,omitempty"`
	ValueType          QueryValidate_ValueType       `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=atlas.query.QueryValidate_ValueType" json:"value_type,omitempty"`
	ValueTypeUrl       string                        `protobuf:"bytes,5,opt,name=value_type_url,json=valueTypeUrl,proto3" json:"value_type_url,omitempty"`
	EnableNestedFields bool                          `protobuf:"varint,6,opt,name=enable_nested_fields,json=enableNestedFields,proto3" json:"enable_nested_fields,omitempty"`
	NestedFields       []string                      `protobuf:"bytes,7,rep,name=nested_fields,json=nestedFields" json:"nested_fields,omitempty"`
}

func (m *QueryValidate) Reset()                    { *m = QueryValidate{} }
func (m *QueryValidate) String() string            { return proto.CompactTextString(m) }
func (*QueryValidate) ProtoMessage()               {}
func (*QueryValidate) Descriptor() ([]byte, []int) { return fileDescriptorQueryValidate, []int{0} }

func (m *QueryValidate) GetFiltering() *QueryValidate_Filtering {
	if m != nil {
		return m.Filtering
	}
	return nil
}

func (m *QueryValidate) GetSorting() *QueryValidate_Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *QueryValidate) GetFieldSelection() *QueryValidate_FieldSelection {
	if m != nil {
		return m.FieldSelection
	}
	return nil
}

func (m *QueryValidate) GetValueType() QueryValidate_ValueType {
	if m != nil {
		return m.ValueType
	}
	return QueryValidate_DEFAULT
}

func (m *QueryValidate) GetValueTypeUrl() string {
	if m != nil {
		return m.ValueTypeUrl
	}
	return ""
}

func (m *QueryValidate) GetEnableNestedFields() bool {
	if m != nil {
		return m.EnableNestedFields
	}
	return false
}

func (m *QueryValidate) GetNestedFields() []string {
	if m != nil {
		return m.NestedFields
	}
	return nil
}

type QueryValidate_Filtering struct {
	Allow []QueryValidate_FilterOperator `protobuf:"varint,1,rep,packed,name=allow,enum=atlas.query.QueryValidate_FilterOperator" json:"allow,omitempty"`
	Deny  []QueryValidate_FilterOperator `protobuf:"varint,2,rep,packed,name=deny,enum=atlas.query.QueryValidate_FilterOperator" json:"deny,omitempty"`
}

func (m *QueryValidate_Filtering) Reset()         { *m = QueryValidate_Filtering{} }
func (m *QueryValidate_Filtering) String() string { return proto.CompactTextString(m) }
func (*QueryValidate_Filtering) ProtoMessage()    {}
func (*QueryValidate_Filtering) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 0}
}

func (m *QueryValidate_Filtering) GetAllow() []QueryValidate_FilterOperator {
	if m != nil {
		return m.Allow
	}
	return nil
}

func (m *QueryValidate_Filtering) GetDeny() []QueryValidate_FilterOperator {
	if m != nil {
		return m.Deny
	}
	return nil
}

type QueryValidate_Sorting struct {
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (m *QueryValidate_Sorting) Reset()         { *m = QueryValidate_Sorting{} }
func (m *QueryValidate_Sorting) String() string { return proto.CompactTextString(m) }
func (*QueryValidate_Sorting) ProtoMessage()    {}
func (*QueryValidate_Sorting) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 1}
}

func (m *QueryValidate_Sorting) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

type QueryValidate_FieldSelection struct {
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (m *QueryValidate_FieldSelection) Reset()         { *m = QueryValidate_FieldSelection{} }
func (m *QueryValidate_FieldSelection) String() string { return proto.CompactTextString(m) }
func (*QueryValidate_FieldSelection) ProtoMessage()    {}
func (*QueryValidate_FieldSelection) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 2}
}

func (m *QueryValidate_FieldSelection) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

type MessageQueryValidate struct {
	Validate []*MessageQueryValidate_QueryValidateEntry `protobuf:"bytes,1,rep,name=validate" json:"validate,omitempty"`
}

func (m *MessageQueryValidate) Reset()         { *m = MessageQueryValidate{} }
func (m *MessageQueryValidate) String() string { return proto.CompactTextString(m) }
func (*MessageQueryValidate) ProtoMessage()    {}
func (*MessageQueryValidate) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{1}
}

func (m *MessageQueryValidate) GetValidate() []*MessageQueryValidate_QueryValidateEntry {
	if m != nil {
		return m.Validate
	}
	return nil
}

type MessageQueryValidate_QueryValidateEntry struct {
	Name  string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *QueryValidate `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *MessageQueryValidate_QueryValidateEntry) Reset() {
	*m = MessageQueryValidate_QueryValidateEntry{}
}
func (m *MessageQueryValidate_QueryValidateEntry) String() string { return proto.CompactTextString(m) }
func (*MessageQueryValidate_QueryValidateEntry) ProtoMessage()    {}
func (*MessageQueryValidate_QueryValidateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{1, 0}
}

func (m *MessageQueryValidate_QueryValidateEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageQueryValidate_QueryValidateEntry) GetValue() *QueryValidate {
	if m != nil {
		return m.Value
	}
	return nil
}

var E_Validate = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*QueryValidate)(nil),
	Field:         52121,
	Name:          "atlas.query.validate",
	Tag:           "bytes,52121,opt,name=validate",
	Filename:      "options/query_validate.proto",
}

var E_Message = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*MessageQueryValidate)(nil),
	Field:         52121,
	Name:          "atlas.query.message",
	Tag:           "bytes,52121,opt,name=message",
	Filename:      "options/query_validate.proto",
}

func init() {
	proto.RegisterType((*QueryValidate)(nil), "atlas.query.QueryValidate")
	proto.RegisterType((*QueryValidate_Filtering)(nil), "atlas.query.QueryValidate.Filtering")
	proto.RegisterType((*QueryValidate_Sorting)(nil), "atlas.query.QueryValidate.Sorting")
	proto.RegisterType((*QueryValidate_FieldSelection)(nil), "atlas.query.QueryValidate.FieldSelection")
	proto.RegisterType((*MessageQueryValidate)(nil), "atlas.query.MessageQueryValidate")
	proto.RegisterType((*MessageQueryValidate_QueryValidateEntry)(nil), "atlas.query.MessageQueryValidate.QueryValidateEntry")
	proto.RegisterEnum("atlas.query.QueryValidate_FilterOperator", QueryValidate_FilterOperator_name, QueryValidate_FilterOperator_value)
	proto.RegisterEnum("atlas.query.QueryValidate_ValueType", QueryValidate_ValueType_name, QueryValidate_ValueType_value)
	proto.RegisterExtension(E_Validate)
	proto.RegisterExtension(E_Message)
}

func init() { proto.RegisterFile("options/query_validate.proto", fileDescriptorQueryValidate) }

var fileDescriptorQueryValidate = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0x97, 0xa6, 0x6d, 0x9a, 0xd3, 0x6f, 0xfd, 0x2c, 0x6b, 0x17, 0x51, 0x05, 0x22, 0x74,
	0xbb, 0x28, 0x48, 0x4d, 0xa7, 0xc1, 0xd5, 0x00, 0xa1, 0x75, 0x64, 0xa3, 0x52, 0xd7, 0x32, 0xaf,
	0x1b, 0xd2, 0x24, 0xa8, 0xd2, 0xc6, 0x2d, 0x91, 0xbc, 0x38, 0x24, 0xee, 0xa0, 0xcf, 0xc0, 0x03,
	0x20, 0x2e, 0x79, 0x17, 0x1e, 0x0c, 0xd9, 0x69, 0xba, 0x85, 0xb1, 0x4d, 0x5c, 0xd9, 0xe9, 0xf9,
	0xff, 0x7f, 0x3d, 0x3e, 0xc7, 0x3e, 0xf0, 0x80, 0x47, 0x22, 0xe0, 0x61, 0xd2, 0xfe, 0x3c, 0xa7,
	0xf1, 0x62, 0x74, 0xe9, 0xb1, 0xc0, 0xf7, 0x04, 0x75, 0xa2, 0x98, 0x0b, 0x8e, 0xab, 0x9e, 0x60,
	0x5e, 0xe2, 0xa8, 0x58, 0xdd, 0x9e, 0x71, 0x3e, 0x63, 0xb4, 0xad, 0x42, 0xe3, 0xf9, 0xb4, 0xed,
	0xd3, 0x64, 0x12, 0x07, 0x91, 0xe0, 0x71, 0x2a, 0x6f, 0xfc, 0x2c, 0xc3, 0xfa, 0xb1, 0xd4, 0x9e,
	0x2d, 0x31, 0xb8, 0x03, 0xe6, 0x34, 0x60, 0x82, 0xc6, 0x41, 0x38, 0xb3, 0x34, 0x5b, 0x6b, 0x56,
	0x77, 0xb6, 0x9c, 0x6b, 0x50, 0x27, 0x27, 0x77, 0x0e, 0x32, 0x2d, 0xb9, 0xb2, 0xe1, 0x97, 0x60,
	0x24, 0x3c, 0x16, 0x92, 0x50, 0x50, 0x84, 0xc6, 0x1d, 0x84, 0x93, 0x54, 0x49, 0x32, 0x0b, 0x26,
	0xf0, 0xff, 0x34, 0xa0, 0xcc, 0x1f, 0x25, 0x94, 0xd1, 0x89, 0x3c, 0xab, 0xa5, 0x2b, 0xca, 0x93,
	0x3b, 0xf3, 0xa0, 0xcc, 0x3f, 0xc9, 0x0c, 0xa4, 0x36, 0xcd, 0x7d, 0xe3, 0x7d, 0x80, 0x4b, 0x8f,
	0xcd, 0xe9, 0x48, 0x2c, 0x22, 0x6a, 0x15, 0x6d, 0xad, 0x59, 0xbb, 0xf3, 0x58, 0x67, 0x52, 0x3c,
	0x5c, 0x44, 0x94, 0x98, 0x97, 0xd9, 0x16, 0x6f, 0x41, 0xed, 0x0a, 0x32, 0x9a, 0xc7, 0xcc, 0x2a,
	0xd9, 0x5a, 0xd3, 0x24, 0xff, 0xad, 0x24, 0xa7, 0x31, 0xc3, 0xdb, 0xb0, 0x41, 0x43, 0x6f, 0xcc,
	0xe8, 0x28, 0xa4, 0x89, 0xa0, 0xfe, 0x48, 0xa5, 0x92, 0x58, 0x65, 0x5b, 0x6b, 0x56, 0x08, 0x4e,
	0x63, 0x7d, 0x15, 0x52, 0x49, 0x27, 0x78, 0x13, 0xd6, 0xf3, 0x52, 0xc3, 0xd6, 0x25, 0x36, 0xbc,
	0x26, 0xaa, 0x7f, 0xd3, 0xc0, 0x5c, 0x15, 0x1b, 0xbf, 0x86, 0x92, 0xc7, 0x18, 0xff, 0x62, 0x69,
	0xb6, 0xde, 0xac, 0xdd, 0x53, 0x19, 0x69, 0x1a, 0x44, 0x34, 0xf6, 0x04, 0x8f, 0x49, 0xea, 0xc3,
	0xaf, 0xa0, 0xe8, 0xd3, 0x70, 0x61, 0x15, 0xfe, 0xd5, 0xaf, 0x6c, 0xf5, 0x4d, 0x30, 0x96, 0x7d,
	0xc3, 0x16, 0x18, 0x7e, 0x90, 0xc8, 0x43, 0xa9, 0xeb, 0x52, 0x21, 0xd9, 0x67, 0xfd, 0x29, 0xd4,
	0xf2, 0x6d, 0xb9, 0x5d, 0xdb, 0xf8, 0x20, 0xb5, 0xd7, 0xff, 0x08, 0x97, 0xa1, 0xe0, 0x1e, 0xa3,
	0x35, 0x6c, 0x42, 0xe9, 0x68, 0x6f, 0xb8, 0xff, 0x16, 0x69, 0xf2, 0xa7, 0xc3, 0x21, 0x2a, 0xa8,
	0xd5, 0x45, 0xba, 0x5c, 0x7b, 0x43, 0x54, 0x54, 0xab, 0x8b, 0x4a, 0xd8, 0x00, 0x7d, 0xaf, 0xd7,
	0x43, 0x65, 0xb9, 0xe9, 0xba, 0xc7, 0xc8, 0x90, 0x91, 0x6e, 0x1f, 0x55, 0x1a, 0xbb, 0x60, 0xae,
	0x5a, 0x8a, 0xab, 0x60, 0xbc, 0x71, 0x0f, 0xf6, 0x4e, 0x7b, 0x43, 0xb4, 0x86, 0x01, 0xca, 0x27,
	0x43, 0xd2, 0xed, 0x1f, 0x22, 0x4d, 0xee, 0xfb, 0xa7, 0x47, 0x1d, 0x97, 0xa0, 0x02, 0xae, 0x40,
	0xb1, 0x33, 0x18, 0xf4, 0x90, 0xde, 0xf8, 0xa5, 0xc1, 0xc6, 0x11, 0x4d, 0x12, 0x6f, 0x46, 0xf3,
	0x4f, 0xe5, 0x1d, 0x54, 0xb2, 0xd7, 0xa7, 0xfa, 0x50, 0xdd, 0x79, 0x9e, 0xab, 0xe3, 0xdf, 0x4c,
	0xf9, 0xe2, 0xba, 0xa1, 0x88, 0x17, 0x64, 0x45, 0xa9, 0x9f, 0x03, 0xbe, 0x19, 0xc7, 0x18, 0x8a,
	0xa1, 0x77, 0x91, 0x96, 0xcc, 0x24, 0x6a, 0x8f, 0xb7, 0xa1, 0xa4, 0x6e, 0xdd, 0xf2, 0x81, 0xd5,
	0x6f, 0x6f, 0x20, 0x49, 0x85, 0xbb, 0xef, 0xaf, 0xb2, 0xc5, 0x0f, 0x9d, 0x74, 0x32, 0x38, 0xd9,
	0x64, 0x48, 0xdf, 0xcf, 0x20, 0x9d, 0x2c, 0xd6, 0x8f, 0xef, 0xfa, 0xbd, 0xd4, 0x15, 0x6c, 0xf7,
	0x23, 0x18, 0x17, 0xe9, 0x49, 0xf1, 0xa3, 0x1b, 0xdc, 0x65, 0x0d, 0xfe, 0x24, 0x3f, 0xbe, 0xb7,
	0x50, 0x24, 0x83, 0x76, 0xba, 0xe7, 0x87, 0xb3, 0x40, 0x7c, 0x9a, 0x8f, 0x9d, 0x09, 0xbf, 0x68,
	0x07, 0xe1, 0x94, 0x8f, 0x19, 0xff, 0xca, 0x23, 0x1a, 0xa6, 0x83, 0x6d, 0xd2, 0x9a, 0xd1, 0xb0,
	0xa5, 0x78, 0x2d, 0xc5, 0x6b, 0x65, 0xa9, 0xb5, 0x97, 0xa3, 0xf2, 0xc5, 0x72, 0x1d, 0x97, 0x95,
	0xe1, 0xd9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xe6, 0x34, 0x1e, 0x44, 0x05, 0x00, 0x00,
}