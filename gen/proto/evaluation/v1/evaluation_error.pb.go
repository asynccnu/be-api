// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/evaluation/v1/evaluation_error.proto

package evaluationv1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type EvaluationErrorReason int32

const (
	EvaluationErrorReason_INVALID_INPUT EvaluationErrorReason = 0
)

// Enum value maps for EvaluationErrorReason.
var (
	EvaluationErrorReason_name = map[int32]string{
		0: "INVALID_INPUT",
	}
	EvaluationErrorReason_value = map[string]int32{
		"INVALID_INPUT": 0,
	}
)

func (x EvaluationErrorReason) Enum() *EvaluationErrorReason {
	p := new(EvaluationErrorReason)
	*p = x
	return p
}

func (x EvaluationErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvaluationErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_evaluation_v1_evaluation_error_proto_enumTypes[0].Descriptor()
}

func (EvaluationErrorReason) Type() protoreflect.EnumType {
	return &file_proto_evaluation_v1_evaluation_error_proto_enumTypes[0]
}

func (x EvaluationErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvaluationErrorReason.Descriptor instead.
func (EvaluationErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_proto_evaluation_v1_evaluation_error_proto_rawDescGZIP(), []int{0}
}

var File_proto_evaluation_v1_evaluation_error_proto protoreflect.FileDescriptor

var file_proto_evaluation_v1_evaluation_error_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x36, 0x0a, 0x15, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x0d, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0xc2, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x14,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x75, 0x78, 0x69, 0x4b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x62,
	0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58,
	0xaa, 0x02, 0x0d, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0d, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x19, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_evaluation_v1_evaluation_error_proto_rawDescOnce sync.Once
	file_proto_evaluation_v1_evaluation_error_proto_rawDescData = file_proto_evaluation_v1_evaluation_error_proto_rawDesc
)

func file_proto_evaluation_v1_evaluation_error_proto_rawDescGZIP() []byte {
	file_proto_evaluation_v1_evaluation_error_proto_rawDescOnce.Do(func() {
		file_proto_evaluation_v1_evaluation_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_evaluation_v1_evaluation_error_proto_rawDescData)
	})
	return file_proto_evaluation_v1_evaluation_error_proto_rawDescData
}

var file_proto_evaluation_v1_evaluation_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_evaluation_v1_evaluation_error_proto_goTypes = []interface{}{
	(EvaluationErrorReason)(0), // 0: evaluation.v1.EvaluationErrorReason
}
var file_proto_evaluation_v1_evaluation_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_evaluation_v1_evaluation_error_proto_init() }
func file_proto_evaluation_v1_evaluation_error_proto_init() {
	if File_proto_evaluation_v1_evaluation_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_evaluation_v1_evaluation_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_evaluation_v1_evaluation_error_proto_goTypes,
		DependencyIndexes: file_proto_evaluation_v1_evaluation_error_proto_depIdxs,
		EnumInfos:         file_proto_evaluation_v1_evaluation_error_proto_enumTypes,
	}.Build()
	File_proto_evaluation_v1_evaluation_error_proto = out.File
	file_proto_evaluation_v1_evaluation_error_proto_rawDesc = nil
	file_proto_evaluation_v1_evaluation_error_proto_goTypes = nil
	file_proto_evaluation_v1_evaluation_error_proto_depIdxs = nil
}