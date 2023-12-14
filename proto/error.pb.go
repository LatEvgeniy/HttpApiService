// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/error.proto

package proto

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

type ErrorCode int32

const (
	ErrorCode_ERROR_NONE                ErrorCode = 0
	ErrorCode_ERROR_INTERNAL            ErrorCode = 1
	ErrorCode_ERROR_INVALID_REQUEST     ErrorCode = 2
	ErrorCode_ERROR_REDIS_PROCESSING    ErrorCode = 3
	ErrorCode_ERROR_POSTGRES_PROCESSING ErrorCode = 4
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_NONE",
		1: "ERROR_INTERNAL",
		2: "ERROR_INVALID_REQUEST",
		3: "ERROR_REDIS_PROCESSING",
		4: "ERROR_POSTGRES_PROCESSING",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_NONE":                0,
		"ERROR_INTERNAL":            1,
		"ERROR_INVALID_REQUEST":     2,
		"ERROR_REDIS_PROCESSING":    3,
		"ERROR_POSTGRES_PROCESSING": 4,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_proto_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_error_proto_rawDescGZIP(), []int{0}
}

type ErrorDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=proto.ErrorCode" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorDto) Reset() {
	*x = ErrorDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDto) ProtoMessage() {}

func (x *ErrorDto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDto.ProtoReflect.Descriptor instead.
func (*ErrorDto) Descriptor() ([]byte, []int) {
	return file_proto_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDto) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_NONE
}

func (x *ErrorDto) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_error_proto protoreflect.FileDescriptor

var file_proto_error_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x08, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x85, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44,
	0x49, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12,
	0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45,
	0x53, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_error_proto_rawDescOnce sync.Once
	file_proto_error_proto_rawDescData = file_proto_error_proto_rawDesc
)

func file_proto_error_proto_rawDescGZIP() []byte {
	file_proto_error_proto_rawDescOnce.Do(func() {
		file_proto_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_error_proto_rawDescData)
	})
	return file_proto_error_proto_rawDescData
}

var file_proto_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_error_proto_goTypes = []interface{}{
	(ErrorCode)(0),   // 0: proto.ErrorCode
	(*ErrorDto)(nil), // 1: proto.ErrorDto
}
var file_proto_error_proto_depIdxs = []int32{
	0, // 0: proto.ErrorDto.code:type_name -> proto.ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_error_proto_init() }
func file_proto_error_proto_init() {
	if File_proto_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDto); i {
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
			RawDescriptor: file_proto_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_error_proto_goTypes,
		DependencyIndexes: file_proto_error_proto_depIdxs,
		EnumInfos:         file_proto_error_proto_enumTypes,
		MessageInfos:      file_proto_error_proto_msgTypes,
	}.Build()
	File_proto_error_proto = out.File
	file_proto_error_proto_rawDesc = nil
	file_proto_error_proto_goTypes = nil
	file_proto_error_proto_depIdxs = nil
}
