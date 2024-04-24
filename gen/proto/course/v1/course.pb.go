// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/course/v1/course.proto

package coursev1

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

// 因为Property，不仅仅要在内部使用，所以改为在proto里面定义枚举类型
type CourseProperty int32

const (
	CourseProperty_CoursePropertyAny             CourseProperty = 0
	CourseProperty_CoursePropertyUnknown         CourseProperty = 1
	CourseProperty_CoursePropertyGeneralCore     CourseProperty = 2 // 通识核心课
	CourseProperty_CoursePropertyGeneralElective CourseProperty = 3 // 通识选修课
	CourseProperty_CoursePropertyGeneralRequired CourseProperty = 4 // 通识必修课
	CourseProperty_CoursePropertyMajorCore       CourseProperty = 5 // 专业主干课程
	CourseProperty_CoursePropertyMajorElective   CourseProperty = 6 // 个性发展课程
)

// Enum value maps for CourseProperty.
var (
	CourseProperty_name = map[int32]string{
		0: "CoursePropertyAny",
		1: "CoursePropertyUnknown",
		2: "CoursePropertyGeneralCore",
		3: "CoursePropertyGeneralElective",
		4: "CoursePropertyGeneralRequired",
		5: "CoursePropertyMajorCore",
		6: "CoursePropertyMajorElective",
	}
	CourseProperty_value = map[string]int32{
		"CoursePropertyAny":             0,
		"CoursePropertyUnknown":         1,
		"CoursePropertyGeneralCore":     2,
		"CoursePropertyGeneralElective": 3,
		"CoursePropertyGeneralRequired": 4,
		"CoursePropertyMajorCore":       5,
		"CoursePropertyMajorElective":   6,
	}
)

func (x CourseProperty) Enum() *CourseProperty {
	p := new(CourseProperty)
	*p = x
	return p
}

func (x CourseProperty) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourseProperty) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_course_v1_course_proto_enumTypes[0].Descriptor()
}

func (CourseProperty) Type() protoreflect.EnumType {
	return &file_proto_course_v1_course_proto_enumTypes[0]
}

func (x CourseProperty) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourseProperty.Descriptor instead.
func (CourseProperty) EnumDescriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{0}
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseCode string         `protobuf:"bytes,2,opt,name=course_code,json=courseCode,proto3" json:"course_code,omitempty"`
	Name       string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Teacher    string         `protobuf:"bytes,5,opt,name=teacher,proto3" json:"teacher,omitempty"`
	School     string         `protobuf:"bytes,6,opt,name=school,proto3" json:"school,omitempty"`                                    // 学院
	Property   CourseProperty `protobuf:"varint,7,opt,name=property,proto3,enum=course.v1.CourseProperty" json:"property,omitempty"` // 课程性质
	Credit     float32        `protobuf:"fixed32,8,opt,name=credit,proto3" json:"credit,omitempty"`                                  // 学分
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetCourseCode() string {
	if x != nil {
		return x.CourseCode
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *Course) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *Course) GetProperty() CourseProperty {
	if x != nil {
		return x.Property
	}
	return CourseProperty_CoursePropertyAny
}

func (x *Course) GetCredit() float32 {
	if x != nil {
		return x.Credit
	}
	return 0
}

type CourseSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	Year   string  `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Term   string  `protobuf:"bytes,3,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *CourseSubscription) Reset() {
	*x = CourseSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseSubscription) ProtoMessage() {}

func (x *CourseSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseSubscription.ProtoReflect.Descriptor instead.
func (*CourseSubscription) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{1}
}

func (x *CourseSubscription) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

func (x *CourseSubscription) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *CourseSubscription) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type SubscriptionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 这里uid和下面的学号密码是两种查询方式所需的参数
	// uid 是从数据查， 学号密码是从教务系统查
	// 其实这里用了强制参数携带uid，其实就是向上游显式暴露了自己的一部分实现
	// 也就是除了从教务系统查询，还做了容错，教务系统没查到可以去自己系统数据查旧的数据
	// 我觉得这部分的显式暴露是不得已的做法，而且这个容错我认为也是必须有的
	// 不适合下游自己来觉得要不要使用容错策略
	Uid       int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	StudentId string `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Year      string `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
	Term      string `protobuf:"bytes,5,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *SubscriptionListRequest) Reset() {
	*x = SubscriptionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionListRequest) ProtoMessage() {}

func (x *SubscriptionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionListRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionListRequest) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{2}
}

func (x *SubscriptionListRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SubscriptionListRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *SubscriptionListRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SubscriptionListRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *SubscriptionListRequest) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

type SubscriptionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseSubscriptions []*CourseSubscription `protobuf:"bytes,1,rep,name=course_subscriptions,json=courseSubscriptions,proto3" json:"course_subscriptions,omitempty"`
}

func (x *SubscriptionListResponse) Reset() {
	*x = SubscriptionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionListResponse) ProtoMessage() {}

func (x *SubscriptionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionListResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionListResponse) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionListResponse) GetCourseSubscriptions() []*CourseSubscription {
	if x != nil {
		return x.CourseSubscriptions
	}
	return nil
}

type GetDetailByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId int64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *GetDetailByIdRequest) Reset() {
	*x = GetDetailByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailByIdRequest) ProtoMessage() {}

func (x *GetDetailByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailByIdRequest.ProtoReflect.Descriptor instead.
func (*GetDetailByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{4}
}

func (x *GetDetailByIdRequest) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type GetDetailByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *GetDetailByIdResponse) Reset() {
	*x = GetDetailByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailByIdResponse) ProtoMessage() {}

func (x *GetDetailByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailByIdResponse.ProtoReflect.Descriptor instead.
func (*GetDetailByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{5}
}

func (x *GetDetailByIdResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type GetSubscriberUidsByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId int64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CurUid   int64 `protobuf:"varint,2,opt,name=cur_uid,json=curUid,proto3" json:"cur_uid,omitempty"`
	Limit    int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetSubscriberUidsByIdRequest) Reset() {
	*x = GetSubscriberUidsByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriberUidsByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriberUidsByIdRequest) ProtoMessage() {}

func (x *GetSubscriberUidsByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriberUidsByIdRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriberUidsByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{6}
}

func (x *GetSubscriberUidsByIdRequest) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *GetSubscriberUidsByIdRequest) GetCurUid() int64 {
	if x != nil {
		return x.CurUid
	}
	return 0
}

func (x *GetSubscriberUidsByIdRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetSubscriberUidsByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviteeUids []int64 `protobuf:"varint,1,rep,packed,name=invitee_uids,json=inviteeUids,proto3" json:"invitee_uids,omitempty"`
}

func (x *GetSubscriberUidsByIdResponse) Reset() {
	*x = GetSubscriberUidsByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriberUidsByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriberUidsByIdResponse) ProtoMessage() {}

func (x *GetSubscriberUidsByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriberUidsByIdResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriberUidsByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{7}
}

func (x *GetSubscriberUidsByIdResponse) GetInviteeUids() []int64 {
	if x != nil {
		return x.InviteeUids
	}
	return nil
}

type SubscribedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	CourseId int64 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *SubscribedRequest) Reset() {
	*x = SubscribedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribedRequest) ProtoMessage() {}

func (x *SubscribedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribedRequest.ProtoReflect.Descriptor instead.
func (*SubscribedRequest) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{8}
}

func (x *SubscribedRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SubscribedRequest) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type SubscribedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscribed bool `protobuf:"varint,1,opt,name=subscribed,proto3" json:"subscribed,omitempty"`
}

func (x *SubscribedResponse) Reset() {
	*x = SubscribedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_course_v1_course_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribedResponse) ProtoMessage() {}

func (x *SubscribedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_course_v1_course_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribedResponse.ProtoReflect.Descriptor instead.
func (*SubscribedResponse) Descriptor() ([]byte, []int) {
	return file_proto_course_v1_course_proto_rawDescGZIP(), []int{9}
}

func (x *SubscribedResponse) GetSubscribed() bool {
	if x != nil {
		return x.Subscribed
	}
	return false
}

var File_proto_course_v1_course_proto protoreflect.FileDescriptor

var file_proto_course_v1_course_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xce, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x35, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x22, 0x67, 0x0a, 0x12, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x29, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x22, 0x8e, 0x01, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x22, 0x6c, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x14, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x33, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x6a, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x55, 0x69, 0x64, 0x73,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x55, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x42, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x55, 0x69, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x55, 0x69, 0x64, 0x73, 0x22, 0x42, 0x0a, 0x11, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22,
	0x34, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x64, 0x2a, 0xe5, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x41, 0x6e, 0x79, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x43, 0x6f, 0x72, 0x65, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x10, 0x04, 0x12,
	0x1b, 0x0a, 0x17, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4d, 0x61,
	0x6a, 0x6f, 0x72, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x06, 0x32, 0xf7, 0x02,
	0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5b, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x55, 0x69, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x55, 0x69, 0x64, 0x73,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x9d, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x75, 0x78, 0x69, 0x4b, 0x65, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_course_v1_course_proto_rawDescOnce sync.Once
	file_proto_course_v1_course_proto_rawDescData = file_proto_course_v1_course_proto_rawDesc
)

func file_proto_course_v1_course_proto_rawDescGZIP() []byte {
	file_proto_course_v1_course_proto_rawDescOnce.Do(func() {
		file_proto_course_v1_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_course_v1_course_proto_rawDescData)
	})
	return file_proto_course_v1_course_proto_rawDescData
}

var file_proto_course_v1_course_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_course_v1_course_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_course_v1_course_proto_goTypes = []interface{}{
	(CourseProperty)(0),                   // 0: course.v1.CourseProperty
	(*Course)(nil),                        // 1: course.v1.Course
	(*CourseSubscription)(nil),            // 2: course.v1.CourseSubscription
	(*SubscriptionListRequest)(nil),       // 3: course.v1.SubscriptionListRequest
	(*SubscriptionListResponse)(nil),      // 4: course.v1.SubscriptionListResponse
	(*GetDetailByIdRequest)(nil),          // 5: course.v1.GetDetailByIdRequest
	(*GetDetailByIdResponse)(nil),         // 6: course.v1.GetDetailByIdResponse
	(*GetSubscriberUidsByIdRequest)(nil),  // 7: course.v1.GetSubscriberUidsByIdRequest
	(*GetSubscriberUidsByIdResponse)(nil), // 8: course.v1.GetSubscriberUidsByIdResponse
	(*SubscribedRequest)(nil),             // 9: course.v1.SubscribedRequest
	(*SubscribedResponse)(nil),            // 10: course.v1.SubscribedResponse
}
var file_proto_course_v1_course_proto_depIdxs = []int32{
	0,  // 0: course.v1.Course.property:type_name -> course.v1.CourseProperty
	1,  // 1: course.v1.CourseSubscription.course:type_name -> course.v1.Course
	2,  // 2: course.v1.SubscriptionListResponse.course_subscriptions:type_name -> course.v1.CourseSubscription
	1,  // 3: course.v1.GetDetailByIdResponse.course:type_name -> course.v1.Course
	3,  // 4: course.v1.CourseService.SubscriptionList:input_type -> course.v1.SubscriptionListRequest
	5,  // 5: course.v1.CourseService.GetDetailById:input_type -> course.v1.GetDetailByIdRequest
	7,  // 6: course.v1.CourseService.GetSubscriberUidsById:input_type -> course.v1.GetSubscriberUidsByIdRequest
	9,  // 7: course.v1.CourseService.Subscribed:input_type -> course.v1.SubscribedRequest
	4,  // 8: course.v1.CourseService.SubscriptionList:output_type -> course.v1.SubscriptionListResponse
	6,  // 9: course.v1.CourseService.GetDetailById:output_type -> course.v1.GetDetailByIdResponse
	8,  // 10: course.v1.CourseService.GetSubscriberUidsById:output_type -> course.v1.GetSubscriberUidsByIdResponse
	10, // 11: course.v1.CourseService.Subscribed:output_type -> course.v1.SubscribedResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_course_v1_course_proto_init() }
func file_proto_course_v1_course_proto_init() {
	if File_proto_course_v1_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_course_v1_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_proto_course_v1_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseSubscription); i {
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
		file_proto_course_v1_course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionListRequest); i {
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
		file_proto_course_v1_course_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionListResponse); i {
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
		file_proto_course_v1_course_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailByIdRequest); i {
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
		file_proto_course_v1_course_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDetailByIdResponse); i {
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
		file_proto_course_v1_course_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriberUidsByIdRequest); i {
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
		file_proto_course_v1_course_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriberUidsByIdResponse); i {
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
		file_proto_course_v1_course_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribedRequest); i {
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
		file_proto_course_v1_course_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribedResponse); i {
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
			RawDescriptor: file_proto_course_v1_course_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_course_v1_course_proto_goTypes,
		DependencyIndexes: file_proto_course_v1_course_proto_depIdxs,
		EnumInfos:         file_proto_course_v1_course_proto_enumTypes,
		MessageInfos:      file_proto_course_v1_course_proto_msgTypes,
	}.Build()
	File_proto_course_v1_course_proto = out.File
	file_proto_course_v1_course_proto_rawDesc = nil
	file_proto_course_v1_course_proto_goTypes = nil
	file_proto_course_v1_course_proto_depIdxs = nil
}
