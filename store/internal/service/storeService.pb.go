// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: storeService.proto

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

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Size     uint64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_storeService_proto_rawDescGZIP(), []int{0}
}

func (x *StoreRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *StoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StoreRequest) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

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
		mi := &file_storeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffRequest) ProtoMessage() {}

func (x *StaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storeService_proto_msgTypes[1]
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
	return file_storeService_proto_rawDescGZIP(), []int{1}
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

type StoreDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreDetail *StoreModel `protobuf:"bytes,1,opt,name=StoreDetail,proto3" json:"StoreDetail,omitempty"`
	Code        uint32      `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *StoreDetailResponse) Reset() {
	*x = StoreDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storeService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDetailResponse) ProtoMessage() {}

func (x *StoreDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storeService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDetailResponse.ProtoReflect.Descriptor instead.
func (*StoreDetailResponse) Descriptor() ([]byte, []int) {
	return file_storeService_proto_rawDescGZIP(), []int{2}
}

func (x *StoreDetailResponse) GetStoreDetail() *StoreModel {
	if x != nil {
		return x.StoreDetail
	}
	return nil
}

func (x *StoreDetailResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_storeService_proto protoreflect.FileDescriptor

var file_storeService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0c, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x5b, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0b, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x93, 0x02, 0x0a,
	0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a,
	0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x6c, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x11, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x1a, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x53, 0x74, 0x61, 0x66, 0x66, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storeService_proto_rawDescOnce sync.Once
	file_storeService_proto_rawDescData = file_storeService_proto_rawDesc
)

func file_storeService_proto_rawDescGZIP() []byte {
	file_storeService_proto_rawDescOnce.Do(func() {
		file_storeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_storeService_proto_rawDescData)
	})
	return file_storeService_proto_rawDescData
}

var file_storeService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storeService_proto_goTypes = []interface{}{
	(*StoreRequest)(nil),        // 0: pb.StoreRequest
	(*StaffRequest)(nil),        // 1: pb.StaffRequest
	(*StoreDetailResponse)(nil), // 2: pb.StoreDetailResponse
	(*StoreModel)(nil),          // 3: pb.StoreModel
}
var file_storeService_proto_depIdxs = []int32{
	3, // 0: pb.StoreDetailResponse.StoreDetail:type_name -> pb.StoreModel
	0, // 1: pb.StoreService.StoreFill:input_type -> pb.StoreRequest
	0, // 2: pb.StoreService.StoreDetailByIdentity:input_type -> pb.StoreRequest
	0, // 3: pb.StoreService.StoreDetailModify:input_type -> pb.StoreRequest
	1, // 4: pb.StoreService.StoreDetailByStaffIdentity:input_type -> pb.StaffRequest
	2, // 5: pb.StoreService.StoreFill:output_type -> pb.StoreDetailResponse
	2, // 6: pb.StoreService.StoreDetailByIdentity:output_type -> pb.StoreDetailResponse
	2, // 7: pb.StoreService.StoreDetailModify:output_type -> pb.StoreDetailResponse
	2, // 8: pb.StoreService.StoreDetailByStaffIdentity:output_type -> pb.StoreDetailResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storeService_proto_init() }
func file_storeService_proto_init() {
	if File_storeService_proto != nil {
		return
	}
	file_storeModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreRequest); i {
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
		file_storeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_storeService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreDetailResponse); i {
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
			RawDescriptor: file_storeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storeService_proto_goTypes,
		DependencyIndexes: file_storeService_proto_depIdxs,
		MessageInfos:      file_storeService_proto_msgTypes,
	}.Build()
	File_storeService_proto = out.File
	file_storeService_proto_rawDesc = nil
	file_storeService_proto_goTypes = nil
	file_storeService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	StoreFill(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error)
	StoreDetailByIdentity(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error)
	StoreDetailModify(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error)
	StoreDetailByStaffIdentity(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) StoreFill(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error) {
	out := new(StoreDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreService/StoreFill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) StoreDetailByIdentity(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error) {
	out := new(StoreDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreService/StoreDetailByIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) StoreDetailModify(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error) {
	out := new(StoreDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreService/StoreDetailModify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) StoreDetailByStaffIdentity(ctx context.Context, in *StaffRequest, opts ...grpc.CallOption) (*StoreDetailResponse, error) {
	out := new(StoreDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.StoreService/StoreDetailByStaffIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	StoreFill(context.Context, *StoreRequest) (*StoreDetailResponse, error)
	StoreDetailByIdentity(context.Context, *StoreRequest) (*StoreDetailResponse, error)
	StoreDetailModify(context.Context, *StoreRequest) (*StoreDetailResponse, error)
	StoreDetailByStaffIdentity(context.Context, *StaffRequest) (*StoreDetailResponse, error)
}

// UnimplementedStoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (*UnimplementedStoreServiceServer) StoreFill(context.Context, *StoreRequest) (*StoreDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreFill not implemented")
}
func (*UnimplementedStoreServiceServer) StoreDetailByIdentity(context.Context, *StoreRequest) (*StoreDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDetailByIdentity not implemented")
}
func (*UnimplementedStoreServiceServer) StoreDetailModify(context.Context, *StoreRequest) (*StoreDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDetailModify not implemented")
}
func (*UnimplementedStoreServiceServer) StoreDetailByStaffIdentity(context.Context, *StaffRequest) (*StoreDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDetailByStaffIdentity not implemented")
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_StoreFill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).StoreFill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreService/StoreFill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).StoreFill(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_StoreDetailByIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).StoreDetailByIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreService/StoreDetailByIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).StoreDetailByIdentity(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_StoreDetailModify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).StoreDetailModify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreService/StoreDetailModify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).StoreDetailModify(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_StoreDetailByStaffIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).StoreDetailByStaffIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StoreService/StoreDetailByStaffIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).StoreDetailByStaffIdentity(ctx, req.(*StaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreFill",
			Handler:    _StoreService_StoreFill_Handler,
		},
		{
			MethodName: "StoreDetailByIdentity",
			Handler:    _StoreService_StoreDetailByIdentity_Handler,
		},
		{
			MethodName: "StoreDetailModify",
			Handler:    _StoreService_StoreDetailModify_Handler,
		},
		{
			MethodName: "StoreDetailByStaffIdentity",
			Handler:    _StoreService_StoreDetailByStaffIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storeService.proto",
}
