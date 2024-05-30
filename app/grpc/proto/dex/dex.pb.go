// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: app/grpc/proto/dex/dex.proto

package dex

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

type DexEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DexEventRequest) Reset() {
	*x = DexEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_proto_dex_dex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DexEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DexEventRequest) ProtoMessage() {}

func (x *DexEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_proto_dex_dex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DexEventRequest.ProtoReflect.Descriptor instead.
func (*DexEventRequest) Descriptor() ([]byte, []int) {
	return file_app_grpc_proto_dex_dex_proto_rawDescGZIP(), []int{0}
}

func (x *DexEventRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DexEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId     int64   `protobuf:"varint,1,opt,name=EventId,proto3" json:"EventId,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Level       string  `protobuf:"bytes,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Created     string  `protobuf:"bytes,4,opt,name=Created,proto3" json:"Created,omitempty"`
	Updated     string  `protobuf:"bytes,5,opt,name=Updated,proto3" json:"Updated,omitempty"`
	IsUsed      bool    `protobuf:"varint,6,opt,name=IsUsed,proto3" json:"IsUsed,omitempty"`
	EventDetail *Detail `protobuf:"bytes,7,opt,name=EventDetail,proto3" json:"EventDetail,omitempty"`
}

func (x *DexEventResponse) Reset() {
	*x = DexEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_proto_dex_dex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DexEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DexEventResponse) ProtoMessage() {}

func (x *DexEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_proto_dex_dex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DexEventResponse.ProtoReflect.Descriptor instead.
func (*DexEventResponse) Descriptor() ([]byte, []int) {
	return file_app_grpc_proto_dex_dex_proto_rawDescGZIP(), []int{1}
}

func (x *DexEventResponse) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *DexEventResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DexEventResponse) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *DexEventResponse) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *DexEventResponse) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *DexEventResponse) GetIsUsed() bool {
	if x != nil {
		return x.IsUsed
	}
	return false
}

func (x *DexEventResponse) GetEventDetail() *Detail {
	if x != nil {
		return x.EventDetail
	}
	return nil
}

type Detail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Define     string `protobuf:"bytes,1,opt,name=Define,proto3" json:"Define,omitempty"`
	Outline    string `protobuf:"bytes,2,opt,name=Outline,proto3" json:"Outline,omitempty"`
	Place      string `protobuf:"bytes,3,opt,name=Place,proto3" json:"Place,omitempty"`
	Background string `protobuf:"bytes,4,opt,name=Background,proto3" json:"Background,omitempty"`
	ImageUrl   string `protobuf:"bytes,5,opt,name=ImageUrl,proto3" json:"ImageUrl,omitempty"`
}

func (x *Detail) Reset() {
	*x = Detail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_grpc_proto_dex_dex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detail) ProtoMessage() {}

func (x *Detail) ProtoReflect() protoreflect.Message {
	mi := &file_app_grpc_proto_dex_dex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detail.ProtoReflect.Descriptor instead.
func (*Detail) Descriptor() ([]byte, []int) {
	return file_app_grpc_proto_dex_dex_proto_rawDescGZIP(), []int{2}
}

func (x *Detail) GetDefine() string {
	if x != nil {
		return x.Define
	}
	return ""
}

func (x *Detail) GetOutline() string {
	if x != nil {
		return x.Outline
	}
	return ""
}

func (x *Detail) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Detail) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *Detail) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_app_grpc_proto_dex_dex_proto protoreflect.FileDescriptor

var file_app_grpc_proto_dex_dex_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x78, 0x2f, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0xd2, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x78, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x0b,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x8c, 0x01, 0x0a,
	0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x32, 0x52, 0x0a, 0x0f, 0x44,
	0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x78,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x19, 0x5a, 0x17, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_app_grpc_proto_dex_dex_proto_rawDescOnce sync.Once
	file_app_grpc_proto_dex_dex_proto_rawDescData = file_app_grpc_proto_dex_dex_proto_rawDesc
)

func file_app_grpc_proto_dex_dex_proto_rawDescGZIP() []byte {
	file_app_grpc_proto_dex_dex_proto_rawDescOnce.Do(func() {
		file_app_grpc_proto_dex_dex_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_grpc_proto_dex_dex_proto_rawDescData)
	})
	return file_app_grpc_proto_dex_dex_proto_rawDescData
}

var file_app_grpc_proto_dex_dex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_grpc_proto_dex_dex_proto_goTypes = []interface{}{
	(*DexEventRequest)(nil),  // 0: grpc.DexEventRequest
	(*DexEventResponse)(nil), // 1: grpc.DexEventResponse
	(*Detail)(nil),           // 2: grpc.Detail
}
var file_app_grpc_proto_dex_dex_proto_depIdxs = []int32{
	2, // 0: grpc.DexEventResponse.EventDetail:type_name -> grpc.Detail
	0, // 1: grpc.DexEventService.FindDexEvent:input_type -> grpc.DexEventRequest
	1, // 2: grpc.DexEventService.FindDexEvent:output_type -> grpc.DexEventResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_grpc_proto_dex_dex_proto_init() }
func file_app_grpc_proto_dex_dex_proto_init() {
	if File_app_grpc_proto_dex_dex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_grpc_proto_dex_dex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DexEventRequest); i {
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
		file_app_grpc_proto_dex_dex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DexEventResponse); i {
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
		file_app_grpc_proto_dex_dex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detail); i {
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
			RawDescriptor: file_app_grpc_proto_dex_dex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_grpc_proto_dex_dex_proto_goTypes,
		DependencyIndexes: file_app_grpc_proto_dex_dex_proto_depIdxs,
		MessageInfos:      file_app_grpc_proto_dex_dex_proto_msgTypes,
	}.Build()
	File_app_grpc_proto_dex_dex_proto = out.File
	file_app_grpc_proto_dex_dex_proto_rawDesc = nil
	file_app_grpc_proto_dex_dex_proto_goTypes = nil
	file_app_grpc_proto_dex_dex_proto_depIdxs = nil
}
