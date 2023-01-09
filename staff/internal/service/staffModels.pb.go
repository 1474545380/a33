// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: staffModels.proto

package service

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

type StaffModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Identity      string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Email         string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Position      string `protobuf:"bytes,5,opt,name=position,proto3" json:"position,omitempty"`
	StoreIdentity string `protobuf:"bytes,6,opt,name=store_identity,json=storeIdentity,proto3" json:"store_identity,omitempty"`
}

func (x *StaffModel) Reset() {
	*x = StaffModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffModel) ProtoMessage() {}

func (x *StaffModel) ProtoReflect() protoreflect.Message {
	mi := &file_staffModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffModel.ProtoReflect.Descriptor instead.
func (*StaffModel) Descriptor() ([]byte, []int) {
	return file_staffModels_proto_rawDescGZIP(), []int{0}
}

func (x *StaffModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StaffModel) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *StaffModel) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *StaffModel) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *StaffModel) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *StaffModel) GetStoreIdentity() string {
	if x != nil {
		return x.StoreIdentity
	}
	return ""
}

type StaffPreferenceModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffPreferenceIdentity string `protobuf:"bytes,1,opt,name=staff_preference_identity,json=staffPreferenceIdentity,proto3" json:"staff_preference_identity,omitempty"`
	PreferenceType          string `protobuf:"bytes,2,opt,name=preference_type,json=preferenceType,proto3" json:"preference_type,omitempty"`
	StaffIdentity           string `protobuf:"bytes,3,opt,name=staff_identity,json=staffIdentity,proto3" json:"staff_identity,omitempty"`
	PreferenceValue         string `protobuf:"bytes,4,opt,name=preference_value,json=preferenceValue,proto3" json:"preference_value,omitempty"`
}

func (x *StaffPreferenceModel) Reset() {
	*x = StaffPreferenceModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffModels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffPreferenceModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffPreferenceModel) ProtoMessage() {}

func (x *StaffPreferenceModel) ProtoReflect() protoreflect.Message {
	mi := &file_staffModels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffPreferenceModel.ProtoReflect.Descriptor instead.
func (*StaffPreferenceModel) Descriptor() ([]byte, []int) {
	return file_staffModels_proto_rawDescGZIP(), []int{1}
}

func (x *StaffPreferenceModel) GetStaffPreferenceIdentity() string {
	if x != nil {
		return x.StaffPreferenceIdentity
	}
	return ""
}

func (x *StaffPreferenceModel) GetPreferenceType() string {
	if x != nil {
		return x.PreferenceType
	}
	return ""
}

func (x *StaffPreferenceModel) GetStaffIdentity() string {
	if x != nil {
		return x.StaffIdentity
	}
	return ""
}

func (x *StaffPreferenceModel) GetPreferenceValue() string {
	if x != nil {
		return x.PreferenceValue
	}
	return ""
}

var File_staffModels_proto protoreflect.FileDescriptor

var file_staffModels_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x61, 0x66, 0x66, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xb1, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xcd, 0x01, 0x0a, 0x14,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x73, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staffModels_proto_rawDescOnce sync.Once
	file_staffModels_proto_rawDescData = file_staffModels_proto_rawDesc
)

func file_staffModels_proto_rawDescGZIP() []byte {
	file_staffModels_proto_rawDescOnce.Do(func() {
		file_staffModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_staffModels_proto_rawDescData)
	})
	return file_staffModels_proto_rawDescData
}

var file_staffModels_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_staffModels_proto_goTypes = []interface{}{
	(*StaffModel)(nil),           // 0: pb.StaffModel
	(*StaffPreferenceModel)(nil), // 1: pb.StaffPreferenceModel
}
var file_staffModels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_staffModels_proto_init() }
func file_staffModels_proto_init() {
	if File_staffModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staffModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffModel); i {
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
		file_staffModels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffPreferenceModel); i {
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
			RawDescriptor: file_staffModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_staffModels_proto_goTypes,
		DependencyIndexes: file_staffModels_proto_depIdxs,
		MessageInfos:      file_staffModels_proto_msgTypes,
	}.Build()
	File_staffModels_proto = out.File
	file_staffModels_proto_rawDesc = nil
	file_staffModels_proto_goTypes = nil
	file_staffModels_proto_depIdxs = nil
}
