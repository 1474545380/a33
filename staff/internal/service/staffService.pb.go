// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: staffService.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StaffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password      string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Position      string `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	StoreIdentity string `protobuf:"bytes,5,opt,name=store_identity,json=storeIdentity,proto3" json:"store_identity,omitempty"`
	Identity      string `protobuf:"bytes,6,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *StaffRequest) Reset() {
	*x = StaffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffRequest) ProtoMessage() {}

func (x *StaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staffService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffRequest.ProtoReflect.Descriptor instead.
func (*StaffRequest) Descriptor() ([]byte, []int) {
	return file_staffService_proto_rawDescGZIP(), []int{0}
}

func (x *StaffRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *StaffRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StaffRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *StaffRequest) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *StaffRequest) GetStoreIdentity() string {
	if x != nil {
		return x.StoreIdentity
	}
	return ""
}

func (x *StaffRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type StaffPreferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffPreferenceIdentity string `protobuf:"bytes,1,opt,name=staff_preference_identity,json=staffPreferenceIdentity,proto3" json:"staff_preference_identity,omitempty"`
	PreferenceType          string `protobuf:"bytes,2,opt,name=preference_type,json=preferenceType,proto3" json:"preference_type,omitempty"`
	StaffIdentity           string `protobuf:"bytes,3,opt,name=staff_identity,json=staffIdentity,proto3" json:"staff_identity,omitempty"`
	PreferenceValue         string `protobuf:"bytes,4,opt,name=preference_value,json=preferenceValue,proto3" json:"preference_value,omitempty"`
}

func (x *StaffPreferenceRequest) Reset() {
	*x = StaffPreferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffPreferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffPreferenceRequest) ProtoMessage() {}

func (x *StaffPreferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staffService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffPreferenceRequest.ProtoReflect.Descriptor instead.
func (*StaffPreferenceRequest) Descriptor() ([]byte, []int) {
	return file_staffService_proto_rawDescGZIP(), []int{1}
}

func (x *StaffPreferenceRequest) GetStaffPreferenceIdentity() string {
	if x != nil {
		return x.StaffPreferenceIdentity
	}
	return ""
}

func (x *StaffPreferenceRequest) GetPreferenceType() string {
	if x != nil {
		return x.PreferenceType
	}
	return ""
}

func (x *StaffPreferenceRequest) GetStaffIdentity() string {
	if x != nil {
		return x.StaffIdentity
	}
	return ""
}

func (x *StaffPreferenceRequest) GetPreferenceValue() string {
	if x != nil {
		return x.PreferenceValue
	}
	return ""
}

type StaffDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffDetail *StaffModel `protobuf:"bytes,1,opt,name=StaffDetail,proto3" json:"StaffDetail,omitempty"`
	Code        uint32      `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *StaffDetailResponse) Reset() {
	*x = StaffDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffDetailResponse) ProtoMessage() {}

func (x *StaffDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staffService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffDetailResponse.ProtoReflect.Descriptor instead.
func (*StaffDetailResponse) Descriptor() ([]byte, []int) {
	return file_staffService_proto_rawDescGZIP(), []int{2}
}

func (x *StaffDetailResponse) GetStaffDetail() *StaffModel {
	if x != nil {
		return x.StaffDetail
	}
	return nil
}

func (x *StaffDetailResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type StaffPreferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffPreferenceDetail []*StaffPreferenceModel `protobuf:"bytes,1,rep,name=StaffPreferenceDetail,proto3" json:"StaffPreferenceDetail,omitempty"`
	Code                  uint32                  `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *StaffPreferenceResponse) Reset() {
	*x = StaffPreferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffPreferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffPreferenceResponse) ProtoMessage() {}

func (x *StaffPreferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staffService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffPreferenceResponse.ProtoReflect.Descriptor instead.
func (*StaffPreferenceResponse) Descriptor() ([]byte, []int) {
	return file_staffService_proto_rawDescGZIP(), []int{3}
}

func (x *StaffPreferenceResponse) GetStaffPreferenceDetail() []*StaffPreferenceModel {
	if x != nil {
		return x.StaffPreferenceDetail
	}
	return nil
}

func (x *StaffPreferenceResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_staffService_proto protoreflect.FileDescriptor

var file_staffService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x73, 0x74, 0x61, 0x66, 0x66, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0xcf, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x19,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x73, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x5b, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x0b, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x7d, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x15, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x32,
	0xe3, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staffService_proto_rawDescOnce sync.Once
	file_staffService_proto_rawDescData = file_staffService_proto_rawDesc
)

func file_staffService_proto_rawDescGZIP() []byte {
	file_staffService_proto_rawDescOnce.Do(func() {
		file_staffService_proto_rawDescData = protoimpl.X.CompressGZIP(file_staffService_proto_rawDescData)
	})
	return file_staffService_proto_rawDescData
}

var file_staffService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_staffService_proto_goTypes = []interface{}{
	(*StaffRequest)(nil),            // 0: pb.StaffRequest
	(*StaffPreferenceRequest)(nil),  // 1: pb.StaffPreferenceRequest
	(*StaffDetailResponse)(nil),     // 2: pb.StaffDetailResponse
	(*StaffPreferenceResponse)(nil), // 3: pb.StaffPreferenceResponse
	(*StaffModel)(nil),              // 4: pb.StaffModel
	(*StaffPreferenceModel)(nil),    // 5: pb.StaffPreferenceModel
}
var file_staffService_proto_depIdxs = []int32{
	4, // 0: pb.StaffDetailResponse.StaffDetail:type_name -> pb.StaffModel
	5, // 1: pb.StaffPreferenceResponse.StaffPreferenceDetail:type_name -> pb.StaffPreferenceModel
	0, // 2: pb.StaffService.StaffLogin:input_type -> pb.StaffRequest
	0, // 3: pb.StaffService.StaffRegister:input_type -> pb.StaffRequest
	0, // 4: pb.StaffService.StaffDetails:input_type -> pb.StaffRequest
	0, // 5: pb.StaffService.StaffDetailsModify:input_type -> pb.StaffRequest
	0, // 6: pb.StaffService.StaffPreference:input_type -> pb.StaffRequest
	1, // 7: pb.StaffService.StaffPreferenceFill:input_type -> pb.StaffPreferenceRequest
	1, // 8: pb.StaffService.StaffPreferenceModify:input_type -> pb.StaffPreferenceRequest
	2, // 9: pb.StaffService.StaffLogin:output_type -> pb.StaffDetailResponse
	2, // 10: pb.StaffService.StaffRegister:output_type -> pb.StaffDetailResponse
	2, // 11: pb.StaffService.StaffDetails:output_type -> pb.StaffDetailResponse
	2, // 12: pb.StaffService.StaffDetailsModify:output_type -> pb.StaffDetailResponse
	3, // 13: pb.StaffService.StaffPreference:output_type -> pb.StaffPreferenceResponse
	3, // 14: pb.StaffService.StaffPreferenceFill:output_type -> pb.StaffPreferenceResponse
	3, // 15: pb.StaffService.StaffPreferenceModify:output_type -> pb.StaffPreferenceResponse
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_staffService_proto_init() }
func file_staffService_proto_init() {
	if File_staffService_proto != nil {
		return
	}
	file_staffModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_staffService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffRequest); i {
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
		file_staffService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffPreferenceRequest); i {
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
		file_staffService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffDetailResponse); i {
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
		file_staffService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffPreferenceResponse); i {
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
			RawDescriptor: file_staffService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staffService_proto_goTypes,
		DependencyIndexes: file_staffService_proto_depIdxs,
		MessageInfos:      file_staffService_proto_msgTypes,
	}.Build()
	File_staffService_proto = out.File
	file_staffService_proto_rawDesc = nil
	file_staffService_proto_goTypes = nil
	file_staffService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StaffServiceClient interface {
	StaffLogin(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error)
	StaffRegister(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error)
	StaffDetails(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error)
	StaffDetailsModify(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error)
	StaffPreference(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffPreferenceResponse, error)
	StaffPreferenceFill(ctx context.Context, in *StaffPreferenceRequest, opts ...grpc.CallOption) (*StaffPreferenceResponse, error)
	StaffPreferenceModify(ctx context.Context, in *StaffPreferenceRequest, opts ...grpc.CallOption) (*StaffPreferenceResponse, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) StaffLogin(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error) {
	out := new(StaffDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StaffService/StaffLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) StaffRegister(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error) {
	out := new(StaffDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StaffService/StaffRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) StaffDetails(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error) {
	out := new(StaffDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StaffService/StaffDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) StaffDetailsModify(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffDetailResponse, error) {
	out := new(StaffDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StaffService/StaffDetailsModify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) StaffPreference(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StaffPreferenceResponse, error) {
	out := new(StaffPreferenceResponse)
	err := c.cc.Invoke(ctx, "/pb.StaffService/StaffPreference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) StaffPreferenceFill(ctx context.Context, in *StaffPreferenceRequest, opts ...grpc.CallOption) (*StaffPreferenceResponse, error) {
	out := new(StaffPreferenceResponse)
	err := c.cc.Invoke(ctx, "/pb.StaffService/StaffPreferenceFill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) StaffPreferenceModify(ctx context.Context, in *StaffPreferenceRequest, opts ...grpc.CallOption) (*StaffPreferenceResponse, error) {
	out := new(StaffPreferenceResponse)
	err := c.cc.Invoke(ctx, "/pb.StaffService/StaffPreferenceModify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
type StaffServiceServer interface {
	StaffLogin(context.Context, *StaffRequest) (*StaffDetailResponse, error)
	StaffRegister(context.Context, *StaffRequest) (*StaffDetailResponse, error)
	StaffDetails(context.Context, *StaffRequest) (*StaffDetailResponse, error)
	StaffDetailsModify(context.Context, *StaffRequest) (*StaffDetailResponse, error)
	StaffPreference(context.Context, *StaffRequest) (*StaffPreferenceResponse, error)
	StaffPreferenceFill(context.Context, *StaffPreferenceRequest) (*StaffPreferenceResponse, error)
	StaffPreferenceModify(context.Context, *StaffPreferenceRequest) (*StaffPreferenceResponse, error)
}

// UnimplementedStaffServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (*UnimplementedStaffServiceServer) StaffLogin(context.Context, *StaffRequest) (*StaffDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffLogin not implemented")
}
func (*UnimplementedStaffServiceServer) StaffRegister(context.Context, *StaffRequest) (*StaffDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffRegister not implemented")
}
func (*UnimplementedStaffServiceServer) StaffDetails(context.Context, *StaffRequest) (*StaffDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffDetails not implemented")
}
func (*UnimplementedStaffServiceServer) StaffDetailsModify(context.Context, *StaffRequest) (*StaffDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffDetailsModify not implemented")
}
func (*UnimplementedStaffServiceServer) StaffPreference(context.Context, *StaffRequest) (*StaffPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffPreference not implemented")
}
func (*UnimplementedStaffServiceServer) StaffPreferenceFill(context.Context, *StaffPreferenceRequest) (*StaffPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffPreferenceFill not implemented")
}
func (*UnimplementedStaffServiceServer) StaffPreferenceModify(context.Context, *StaffPreferenceRequest) (*StaffPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffPreferenceModify not implemented")
}

func RegisterStaffServiceServer(s *grpc.Server, srv StaffServiceServer) {
	s.RegisterService(&_StaffService_serviceDesc, srv)
}

func _StaffService_StaffLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StaffService/StaffLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffLogin(ctx, req.(*StaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_StaffRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StaffService/StaffRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffRegister(ctx, req.(*StaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_StaffDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StaffService/StaffDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffDetails(ctx, req.(*StaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_StaffDetailsModify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffDetailsModify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StaffService/StaffDetailsModify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffDetailsModify(ctx, req.(*StaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_StaffPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StaffService/StaffPreference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffPreference(ctx, req.(*StaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_StaffPreferenceFill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffPreferenceFill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StaffService/StaffPreferenceFill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffPreferenceFill(ctx, req.(*StaffPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_StaffPreferenceModify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).StaffPreferenceModify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StaffService/StaffPreferenceModify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).StaffPreferenceModify(ctx, req.(*StaffPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StaffService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StaffLogin",
			Handler:    _StaffService_StaffLogin_Handler,
		},
		{
			MethodName: "StaffRegister",
			Handler:    _StaffService_StaffRegister_Handler,
		},
		{
			MethodName: "StaffDetails",
			Handler:    _StaffService_StaffDetails_Handler,
		},
		{
			MethodName: "StaffDetailsModify",
			Handler:    _StaffService_StaffDetailsModify_Handler,
		},
		{
			MethodName: "StaffPreference",
			Handler:    _StaffService_StaffPreference_Handler,
		},
		{
			MethodName: "StaffPreferenceFill",
			Handler:    _StaffService_StaffPreferenceFill_Handler,
		},
		{
			MethodName: "StaffPreferenceModify",
			Handler:    _StaffService_StaffPreferenceModify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staffService.proto",
}
