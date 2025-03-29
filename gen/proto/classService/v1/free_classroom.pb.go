// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.26.1
// source: free_classroom.proto

package classServicev1

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

type QueryFreeClassroomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year        string  `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	Semester    string  `protobuf:"bytes,2,opt,name=semester,proto3" json:"semester,omitempty"`
	Week        int32   `protobuf:"varint,3,opt,name=week,proto3" json:"week,omitempty"`                //哪一周
	Day         int32   `protobuf:"varint,4,opt,name=day,proto3" json:"day,omitempty"`                  //哪一天
	Sections    []int32 `protobuf:"varint,5,rep,packed,name=sections,proto3" json:"sections,omitempty"` //哪几节课
	WherePrefix string  `protobuf:"bytes,6,opt,name=wherePrefix,proto3" json:"wherePrefix,omitempty"`   //查询的地点前缀，比如南湖1楼，就是 "n1"，7号教学楼2楼就是"72"
	StuID       string  `protobuf:"bytes,7,opt,name=stuID,proto3" json:"stuID,omitempty"`               //学号
}

func (x *QueryFreeClassroomReq) Reset() {
	*x = QueryFreeClassroomReq{}
	mi := &file_free_classroom_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryFreeClassroomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFreeClassroomReq) ProtoMessage() {}

func (x *QueryFreeClassroomReq) ProtoReflect() protoreflect.Message {
	mi := &file_free_classroom_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFreeClassroomReq.ProtoReflect.Descriptor instead.
func (*QueryFreeClassroomReq) Descriptor() ([]byte, []int) {
	return file_free_classroom_proto_rawDescGZIP(), []int{0}
}

func (x *QueryFreeClassroomReq) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *QueryFreeClassroomReq) GetSemester() string {
	if x != nil {
		return x.Semester
	}
	return ""
}

func (x *QueryFreeClassroomReq) GetWeek() int32 {
	if x != nil {
		return x.Week
	}
	return 0
}

func (x *QueryFreeClassroomReq) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *QueryFreeClassroomReq) GetSections() []int32 {
	if x != nil {
		return x.Sections
	}
	return nil
}

func (x *QueryFreeClassroomReq) GetWherePrefix() string {
	if x != nil {
		return x.WherePrefix
	}
	return ""
}

func (x *QueryFreeClassroomReq) GetStuID() string {
	if x != nil {
		return x.StuID
	}
	return ""
}

type QueryFreeClassroomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat []*ClassroomAvailableStat `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
}

func (x *QueryFreeClassroomResp) Reset() {
	*x = QueryFreeClassroomResp{}
	mi := &file_free_classroom_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryFreeClassroomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFreeClassroomResp) ProtoMessage() {}

func (x *QueryFreeClassroomResp) ProtoReflect() protoreflect.Message {
	mi := &file_free_classroom_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFreeClassroomResp.ProtoReflect.Descriptor instead.
func (*QueryFreeClassroomResp) Descriptor() ([]byte, []int) {
	return file_free_classroom_proto_rawDescGZIP(), []int{1}
}

func (x *QueryFreeClassroomResp) GetStat() []*ClassroomAvailableStat {
	if x != nil {
		return x.Stat
	}
	return nil
}

type ClassroomAvailableStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classroom     string `protobuf:"bytes,1,opt,name=classroom,proto3" json:"classroom,omitempty"`
	AvailableStat []bool `protobuf:"varint,2,rep,packed,name=availableStat,proto3" json:"availableStat,omitempty"` //空闲情况，顺序和请求的sections对应, true表示空闲
}

func (x *ClassroomAvailableStat) Reset() {
	*x = ClassroomAvailableStat{}
	mi := &file_free_classroom_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClassroomAvailableStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassroomAvailableStat) ProtoMessage() {}

func (x *ClassroomAvailableStat) ProtoReflect() protoreflect.Message {
	mi := &file_free_classroom_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassroomAvailableStat.ProtoReflect.Descriptor instead.
func (*ClassroomAvailableStat) Descriptor() ([]byte, []int) {
	return file_free_classroom_proto_rawDescGZIP(), []int{2}
}

func (x *ClassroomAvailableStat) GetClassroom() string {
	if x != nil {
		return x.Classroom
	}
	return ""
}

func (x *ClassroomAvailableStat) GetAvailableStat() []bool {
	if x != nil {
		return x.AvailableStat
	}
	return nil
}

var File_free_classroom_proto protoreflect.FileDescriptor

var file_free_classroom_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xc1, 0x01, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x77, 0x65, 0x65, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x68, 0x65, 0x72, 0x65, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x68, 0x65, 0x72, 0x65, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x75, 0x49, 0x44, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x75, 0x49, 0x44, 0x22, 0x55, 0x0a, 0x16, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73, 0x74,
	0x61, 0x74, 0x22, 0x5c, 0x0a, 0x16, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x08, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x32, 0x79, 0x0a, 0x10, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x53, 0x76, 0x63, 0x12, 0x65, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x72, 0x65,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x1a, 0x27, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x63,
	0x6e, 0x75, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_free_classroom_proto_rawDescOnce sync.Once
	file_free_classroom_proto_rawDescData = file_free_classroom_proto_rawDesc
)

func file_free_classroom_proto_rawDescGZIP() []byte {
	file_free_classroom_proto_rawDescOnce.Do(func() {
		file_free_classroom_proto_rawDescData = protoimpl.X.CompressGZIP(file_free_classroom_proto_rawDescData)
	})
	return file_free_classroom_proto_rawDescData
}

var file_free_classroom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_free_classroom_proto_goTypes = []any{
	(*QueryFreeClassroomReq)(nil),  // 0: classService.v1.QueryFreeClassroomReq
	(*QueryFreeClassroomResp)(nil), // 1: classService.v1.QueryFreeClassroomResp
	(*ClassroomAvailableStat)(nil), // 2: classService.v1.ClassroomAvailableStat
}
var file_free_classroom_proto_depIdxs = []int32{
	2, // 0: classService.v1.QueryFreeClassroomResp.stat:type_name -> classService.v1.ClassroomAvailableStat
	0, // 1: classService.v1.FreeClassroomSvc.QueryFreeClassroom:input_type -> classService.v1.QueryFreeClassroomReq
	1, // 2: classService.v1.FreeClassroomSvc.QueryFreeClassroom:output_type -> classService.v1.QueryFreeClassroomResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_free_classroom_proto_init() }
func file_free_classroom_proto_init() {
	if File_free_classroom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_free_classroom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_free_classroom_proto_goTypes,
		DependencyIndexes: file_free_classroom_proto_depIdxs,
		MessageInfos:      file_free_classroom_proto_msgTypes,
	}.Build()
	File_free_classroom_proto = out.File
	file_free_classroom_proto_rawDesc = nil
	file_free_classroom_proto_goTypes = nil
	file_free_classroom_proto_depIdxs = nil
}
