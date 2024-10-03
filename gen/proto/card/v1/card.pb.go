// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: card.proto

package cardv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateUserRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type UpdateUserKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *UpdateUserKeyRequest) Reset() {
	*x = UpdateUserKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserKeyRequest) ProtoMessage() {}

func (x *UpdateUserKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserKeyRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserKeyRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *UpdateUserKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetRecordOfConsumptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	StartTime string `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Type      string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetRecordOfConsumptionRequest) Reset() {
	*x = GetRecordOfConsumptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordOfConsumptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordOfConsumptionRequest) ProtoMessage() {}

func (x *GetRecordOfConsumptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordOfConsumptionRequest.ProtoReflect.Descriptor instead.
func (*GetRecordOfConsumptionRequest) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{2}
}

func (x *GetRecordOfConsumptionRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *GetRecordOfConsumptionRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetRecordOfConsumptionRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *GetRecordOfConsumptionRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetRecordOfConsumptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*RecordOfConsumption `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *GetRecordOfConsumptionResponse) Reset() {
	*x = GetRecordOfConsumptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordOfConsumptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordOfConsumptionResponse) ProtoMessage() {}

func (x *GetRecordOfConsumptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordOfConsumptionResponse.ProtoReflect.Descriptor instead.
func (*GetRecordOfConsumptionResponse) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecordOfConsumptionResponse) GetRecords() []*RecordOfConsumption {
	if x != nil {
		return x.Records
	}
	return nil
}

type RecordOfConsumption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SMT_TIMES        uint32                 `protobuf:"varint,1,opt,name=SMT_TIMES,json=SMTTIMES,proto3" json:"SMT_TIMES,omitempty"`
	SMT_DEALDATETIME *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=SMT_DEALDATETIME,json=SMTDEALDATETIME,proto3" json:"SMT_DEALDATETIME,omitempty"`
	SMT_ORG_NAME     string                 `protobuf:"bytes,3,opt,name=SMT_ORG_NAME,json=SMTORGNAME,proto3" json:"SMT_ORG_NAME,omitempty"`  //消费窗口
	SMT_DEALNAME     string                 `protobuf:"bytes,4,opt,name=SMT_DEALNAME,json=SMTDEALNAME,proto3" json:"SMT_DEALNAME,omitempty"` //消费方式
	AfterMoney       float32                `protobuf:"fixed32,5,opt,name=after_money,json=afterMoney,proto3" json:"after_money,omitempty"`
	Money            float32                `protobuf:"fixed32,6,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *RecordOfConsumption) Reset() {
	*x = RecordOfConsumption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordOfConsumption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordOfConsumption) ProtoMessage() {}

func (x *RecordOfConsumption) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordOfConsumption.ProtoReflect.Descriptor instead.
func (*RecordOfConsumption) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{4}
}

func (x *RecordOfConsumption) GetSMT_TIMES() uint32 {
	if x != nil {
		return x.SMT_TIMES
	}
	return 0
}

func (x *RecordOfConsumption) GetSMT_DEALDATETIME() *timestamppb.Timestamp {
	if x != nil {
		return x.SMT_DEALDATETIME
	}
	return nil
}

func (x *RecordOfConsumption) GetSMT_ORG_NAME() string {
	if x != nil {
		return x.SMT_ORG_NAME
	}
	return ""
}

func (x *RecordOfConsumption) GetSMT_DEALNAME() string {
	if x != nil {
		return x.SMT_DEALNAME
	}
	return ""
}

func (x *RecordOfConsumption) GetAfterMoney() float32 {
	if x != nil {
		return x.AfterMoney
	}
	return 0
}

func (x *RecordOfConsumption) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

type OperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OperationResponse) Reset() {
	*x = OperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResponse) ProtoMessage() {}

func (x *OperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResponse.ProtoReflect.Descriptor instead.
func (*OperationResponse) Descriptor() ([]byte, []int) {
	return file_card_proto_rawDescGZIP(), []int{5}
}

func (x *OperationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *OperationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_card_proto protoreflect.FileDescriptor

var file_card_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x47, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x83, 0x01, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x5c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x66, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xf5, 0x01,
	0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x53, 0x4d, 0x54, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x4d, 0x54, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x12, 0x45, 0x0a, 0x10, 0x53, 0x4d, 0x54, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x44, 0x41,
	0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x53, 0x4d, 0x54, 0x44, 0x45, 0x41,
	0x4c, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x12, 0x20, 0x0a, 0x0c, 0x53, 0x4d, 0x54,
	0x5f, 0x4f, 0x52, 0x47, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x53, 0x4d, 0x54, 0x4f, 0x52, 0x47, 0x4e, 0x41, 0x4d, 0x45, 0x12, 0x21, 0x0a, 0x0c, 0x53,
	0x4d, 0x54, 0x5f, 0x44, 0x45, 0x41, 0x4c, 0x4e, 0x41, 0x4d, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x53, 0x4d, 0x54, 0x44, 0x45, 0x41, 0x4c, 0x4e, 0x41, 0x4d, 0x45, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x61, 0x66, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9b,
	0x02, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x4c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x66, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x66, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63,
	0x63, 0x6e, 0x75, 0x2f, 0x62, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x72,
	0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_card_proto_rawDescOnce sync.Once
	file_card_proto_rawDescData = file_card_proto_rawDesc
)

func file_card_proto_rawDescGZIP() []byte {
	file_card_proto_rawDescOnce.Do(func() {
		file_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_proto_rawDescData)
	})
	return file_card_proto_rawDescData
}

var file_card_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_card_proto_goTypes = []any{
	(*CreateUserRequest)(nil),              // 0: api.card.v1.CreateUserRequest
	(*UpdateUserKeyRequest)(nil),           // 1: api.card.v1.UpdateUserKeyRequest
	(*GetRecordOfConsumptionRequest)(nil),  // 2: api.card.v1.GetRecordOfConsumptionRequest
	(*GetRecordOfConsumptionResponse)(nil), // 3: api.card.v1.GetRecordOfConsumptionResponse
	(*RecordOfConsumption)(nil),            // 4: api.card.v1.RecordOfConsumption
	(*OperationResponse)(nil),              // 5: api.card.v1.OperationResponse
	(*timestamppb.Timestamp)(nil),          // 6: google.protobuf.Timestamp
}
var file_card_proto_depIdxs = []int32{
	4, // 0: api.card.v1.GetRecordOfConsumptionResponse.records:type_name -> api.card.v1.RecordOfConsumption
	6, // 1: api.card.v1.RecordOfConsumption.SMT_DEALDATETIME:type_name -> google.protobuf.Timestamp
	0, // 2: api.card.v1.card.CreateUser:input_type -> api.card.v1.CreateUserRequest
	1, // 3: api.card.v1.card.UpdateUserKey:input_type -> api.card.v1.UpdateUserKeyRequest
	2, // 4: api.card.v1.card.GetRecordOfConsumption:input_type -> api.card.v1.GetRecordOfConsumptionRequest
	5, // 5: api.card.v1.card.CreateUser:output_type -> api.card.v1.OperationResponse
	5, // 6: api.card.v1.card.UpdateUserKey:output_type -> api.card.v1.OperationResponse
	3, // 7: api.card.v1.card.GetRecordOfConsumption:output_type -> api.card.v1.GetRecordOfConsumptionResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_card_proto_init() }
func file_card_proto_init() {
	if File_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_card_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateUserRequest); i {
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
		file_card_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserKeyRequest); i {
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
		file_card_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetRecordOfConsumptionRequest); i {
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
		file_card_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetRecordOfConsumptionResponse); i {
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
		file_card_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RecordOfConsumption); i {
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
		file_card_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OperationResponse); i {
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
			RawDescriptor: file_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_card_proto_goTypes,
		DependencyIndexes: file_card_proto_depIdxs,
		MessageInfos:      file_card_proto_msgTypes,
	}.Build()
	File_card_proto = out.File
	file_card_proto_rawDesc = nil
	file_card_proto_goTypes = nil
	file_card_proto_depIdxs = nil
}
