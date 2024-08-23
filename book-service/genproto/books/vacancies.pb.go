// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: vacancies.proto

package books

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

type VacanciesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Publisher     *PublishersRes `protobuf:"bytes,2,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Title         string         `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description   string         `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status        string         `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	SalaryFrom    float32        `protobuf:"fixed32,6,opt,name=salary_from,json=salaryFrom,proto3" json:"salary_from,omitempty"`
	SalaryTo      float32        `protobuf:"fixed32,7,opt,name=salary_to,json=salaryTo,proto3" json:"salary_to,omitempty"`
	WorkingStyles string         `protobuf:"bytes,8,opt,name=working_styles,json=workingStyles,proto3" json:"working_styles,omitempty"`
	WorkingTypes  string         `protobuf:"bytes,9,opt,name=working_types,json=workingTypes,proto3" json:"working_types,omitempty"`
	PhoneNumber   string         `protobuf:"bytes,10,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Location      string         `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	ViewCount     int64          `protobuf:"varint,12,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
}

func (x *VacanciesRes) Reset() {
	*x = VacanciesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vacancies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacanciesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacanciesRes) ProtoMessage() {}

func (x *VacanciesRes) ProtoReflect() protoreflect.Message {
	mi := &file_vacancies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacanciesRes.ProtoReflect.Descriptor instead.
func (*VacanciesRes) Descriptor() ([]byte, []int) {
	return file_vacancies_proto_rawDescGZIP(), []int{0}
}

func (x *VacanciesRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VacanciesRes) GetPublisher() *PublishersRes {
	if x != nil {
		return x.Publisher
	}
	return nil
}

func (x *VacanciesRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VacanciesRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VacanciesRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VacanciesRes) GetSalaryFrom() float32 {
	if x != nil {
		return x.SalaryFrom
	}
	return 0
}

func (x *VacanciesRes) GetSalaryTo() float32 {
	if x != nil {
		return x.SalaryTo
	}
	return 0
}

func (x *VacanciesRes) GetWorkingStyles() string {
	if x != nil {
		return x.WorkingStyles
	}
	return ""
}

func (x *VacanciesRes) GetWorkingTypes() string {
	if x != nil {
		return x.WorkingTypes
	}
	return ""
}

func (x *VacanciesRes) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *VacanciesRes) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *VacanciesRes) GetViewCount() int64 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

type VacanciesCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublisherId   string  `protobuf:"bytes,1,opt,name=publisher_id,json=publisherId,proto3" json:"publisher_id,omitempty"`
	Title         string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	SalaryFrom    float32 `protobuf:"fixed32,5,opt,name=salary_from,json=salaryFrom,proto3" json:"salary_from,omitempty"`
	SalaryTo      float32 `protobuf:"fixed32,6,opt,name=salary_to,json=salaryTo,proto3" json:"salary_to,omitempty"`
	WorkingStyles string  `protobuf:"bytes,7,opt,name=working_styles,json=workingStyles,proto3" json:"working_styles,omitempty"`
	WorkingTypes  string  `protobuf:"bytes,8,opt,name=working_types,json=workingTypes,proto3" json:"working_types,omitempty"`
	PhoneNumber   string  `protobuf:"bytes,9,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Location      string  `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	ViewCount     int64   `protobuf:"varint,11,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
}

func (x *VacanciesCreateReq) Reset() {
	*x = VacanciesCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vacancies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacanciesCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacanciesCreateReq) ProtoMessage() {}

func (x *VacanciesCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_vacancies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacanciesCreateReq.ProtoReflect.Descriptor instead.
func (*VacanciesCreateReq) Descriptor() ([]byte, []int) {
	return file_vacancies_proto_rawDescGZIP(), []int{1}
}

func (x *VacanciesCreateReq) GetPublisherId() string {
	if x != nil {
		return x.PublisherId
	}
	return ""
}

func (x *VacanciesCreateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VacanciesCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VacanciesCreateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VacanciesCreateReq) GetSalaryFrom() float32 {
	if x != nil {
		return x.SalaryFrom
	}
	return 0
}

func (x *VacanciesCreateReq) GetSalaryTo() float32 {
	if x != nil {
		return x.SalaryTo
	}
	return 0
}

func (x *VacanciesCreateReq) GetWorkingStyles() string {
	if x != nil {
		return x.WorkingStyles
	}
	return ""
}

func (x *VacanciesCreateReq) GetWorkingTypes() string {
	if x != nil {
		return x.WorkingTypes
	}
	return ""
}

func (x *VacanciesCreateReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *VacanciesCreateReq) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *VacanciesCreateReq) GetViewCount() int64 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

type VacanciesGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vacancies []*VacanciesRes `protobuf:"bytes,1,rep,name=vacancies,proto3" json:"vacancies,omitempty"`
}

func (x *VacanciesGetAllRes) Reset() {
	*x = VacanciesGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vacancies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacanciesGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacanciesGetAllRes) ProtoMessage() {}

func (x *VacanciesGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_vacancies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacanciesGetAllRes.ProtoReflect.Descriptor instead.
func (*VacanciesGetAllRes) Descriptor() ([]byte, []int) {
	return file_vacancies_proto_rawDescGZIP(), []int{2}
}

func (x *VacanciesGetAllRes) GetVacancies() []*VacanciesRes {
	if x != nil {
		return x.Vacancies
	}
	return nil
}

type VacanciesUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vacancy *VacanciesCreateReq `protobuf:"bytes,1,opt,name=vacancy,proto3" json:"vacancy,omitempty"`
	Id      string              `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VacanciesUpdateReq) Reset() {
	*x = VacanciesUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vacancies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacanciesUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacanciesUpdateReq) ProtoMessage() {}

func (x *VacanciesUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_vacancies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacanciesUpdateReq.ProtoReflect.Descriptor instead.
func (*VacanciesUpdateReq) Descriptor() ([]byte, []int) {
	return file_vacancies_proto_rawDescGZIP(), []int{3}
}

func (x *VacanciesUpdateReq) GetVacancy() *VacanciesCreateReq {
	if x != nil {
		return x.Vacancy
	}
	return nil
}

func (x *VacanciesUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VacanciesGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SalaryFrom float32 `protobuf:"fixed32,1,opt,name=salary_from,json=salaryFrom,proto3" json:"salary_from,omitempty"`
	SalaryTo   float32 `protobuf:"fixed32,2,opt,name=salary_to,json=salaryTo,proto3" json:"salary_to,omitempty"`
	Status     string  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Page       *Filter `protobuf:"bytes,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *VacanciesGetAllReq) Reset() {
	*x = VacanciesGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vacancies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VacanciesGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VacanciesGetAllReq) ProtoMessage() {}

func (x *VacanciesGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_vacancies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VacanciesGetAllReq.ProtoReflect.Descriptor instead.
func (*VacanciesGetAllReq) Descriptor() ([]byte, []int) {
	return file_vacancies_proto_rawDescGZIP(), []int{4}
}

func (x *VacanciesGetAllReq) GetSalaryFrom() float32 {
	if x != nil {
		return x.SalaryFrom
	}
	return 0
}

func (x *VacanciesGetAllReq) GetSalaryTo() float32 {
	if x != nil {
		return x.SalaryTo
	}
	return 0
}

func (x *VacanciesGetAllReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *VacanciesGetAllReq) GetPage() *Filter {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_vacancies_proto protoreflect.FileDescriptor

var file_vacancies_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x0c, 0x56, 0x61,
	0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x54, 0x6f, 0x12, 0x25, 0x0a, 0x0e,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x69, 0x65,
	0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xef, 0x02, 0x0a, 0x12, 0x56, 0x61, 0x63, 0x61, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x54, 0x6f, 0x12, 0x25,
	0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x79, 0x6c, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x12, 0x56, 0x61, 0x63, 0x61,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x31,
	0x0a, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x52, 0x09, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x22, 0x59, 0x0a, 0x12, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x61, 0x63, 0x61, 0x6e,
	0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x52, 0x07, 0x76, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8d, 0x01, 0x0a,
	0x12, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x5f, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x54,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x87, 0x02, 0x0a,
	0x10, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x12, 0x3e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x56, 0x61,
	0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vacancies_proto_rawDescOnce sync.Once
	file_vacancies_proto_rawDescData = file_vacancies_proto_rawDesc
)

func file_vacancies_proto_rawDescGZIP() []byte {
	file_vacancies_proto_rawDescOnce.Do(func() {
		file_vacancies_proto_rawDescData = protoimpl.X.CompressGZIP(file_vacancies_proto_rawDescData)
	})
	return file_vacancies_proto_rawDescData
}

var file_vacancies_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vacancies_proto_goTypes = []interface{}{
	(*VacanciesRes)(nil),       // 0: books.VacanciesRes
	(*VacanciesCreateReq)(nil), // 1: books.VacanciesCreateReq
	(*VacanciesGetAllRes)(nil), // 2: books.VacanciesGetAllRes
	(*VacanciesUpdateReq)(nil), // 3: books.VacanciesUpdateReq
	(*VacanciesGetAllReq)(nil), // 4: books.VacanciesGetAllReq
	(*PublishersRes)(nil),      // 5: books.PublishersRes
	(*Filter)(nil),             // 6: books.Filter
	(*ById)(nil),               // 7: books.ById
	(*Void)(nil),               // 8: books.Void
}
var file_vacancies_proto_depIdxs = []int32{
	5, // 0: books.VacanciesRes.publisher:type_name -> books.PublishersRes
	0, // 1: books.VacanciesGetAllRes.vacancies:type_name -> books.VacanciesRes
	1, // 2: books.VacanciesUpdateReq.vacancy:type_name -> books.VacanciesCreateReq
	6, // 3: books.VacanciesGetAllReq.page:type_name -> books.Filter
	1, // 4: books.VacanciesService.Create:input_type -> books.VacanciesCreateReq
	7, // 5: books.VacanciesService.GetById:input_type -> books.ById
	4, // 6: books.VacanciesService.GetAll:input_type -> books.VacanciesGetAllReq
	3, // 7: books.VacanciesService.Update:input_type -> books.VacanciesUpdateReq
	7, // 8: books.VacanciesService.Delete:input_type -> books.ById
	8, // 9: books.VacanciesService.Create:output_type -> books.Void
	0, // 10: books.VacanciesService.GetById:output_type -> books.VacanciesRes
	2, // 11: books.VacanciesService.GetAll:output_type -> books.VacanciesGetAllRes
	8, // 12: books.VacanciesService.Update:output_type -> books.Void
	8, // 13: books.VacanciesService.Delete:output_type -> books.Void
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vacancies_proto_init() }
func file_vacancies_proto_init() {
	if File_vacancies_proto != nil {
		return
	}
	file_common2_proto_init()
	file_publishers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vacancies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacanciesRes); i {
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
		file_vacancies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacanciesCreateReq); i {
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
		file_vacancies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacanciesGetAllRes); i {
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
		file_vacancies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacanciesUpdateReq); i {
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
		file_vacancies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VacanciesGetAllReq); i {
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
			RawDescriptor: file_vacancies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vacancies_proto_goTypes,
		DependencyIndexes: file_vacancies_proto_depIdxs,
		MessageInfos:      file_vacancies_proto_msgTypes,
	}.Build()
	File_vacancies_proto = out.File
	file_vacancies_proto_rawDesc = nil
	file_vacancies_proto_goTypes = nil
	file_vacancies_proto_depIdxs = nil
}
