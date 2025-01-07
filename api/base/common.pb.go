// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.14.0
// source: common.proto

package api

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

type MediaType int32

const (
	MediaType_MT_DEFAULT    MediaType = 0
	MediaType_MT_POFP_IMAGE MediaType = 1
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0: "MT_DEFAULT",
		1: "MT_POFP_IMAGE",
	}
	MediaType_value = map[string]int32{
		"MT_DEFAULT":    0,
		"MT_POFP_IMAGE": 1,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type PointCoord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Lat           float32                `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng           float32                `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PointCoord) Reset() {
	*x = PointCoord{}
	mi := &file_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PointCoord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointCoord) ProtoMessage() {}

func (x *PointCoord) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointCoord.ProtoReflect.Descriptor instead.
func (*PointCoord) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *PointCoord) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *PointCoord) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

type BoundCoord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sw            *PointCoord            `protobuf:"bytes,1,opt,name=sw,proto3" json:"sw,omitempty"`
	Ne            *PointCoord            `protobuf:"bytes,2,opt,name=ne,proto3" json:"ne,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BoundCoord) Reset() {
	*x = BoundCoord{}
	mi := &file_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoundCoord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundCoord) ProtoMessage() {}

func (x *BoundCoord) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundCoord.ProtoReflect.Descriptor instead.
func (*BoundCoord) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *BoundCoord) GetSw() *PointCoord {
	if x != nil {
		return x.Sw
	}
	return nil
}

func (x *BoundCoord) GetNe() *PointCoord {
	if x != nil {
		return x.Ne
	}
	return nil
}

type MediaInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID, 可写入
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// 类型
	Type MediaType `protobuf:"varint,2,opt,name=type,proto3,enum=common.MediaType" json:"type,omitempty"`
	// 读取URL
	GetUrl string `protobuf:"bytes,3,opt,name=get_url,json=getUrl,proto3" json:"get_url,omitempty"`
	// 写入URL
	PutUrl        string `protobuf:"bytes,4,opt,name=put_url,json=putUrl,proto3" json:"put_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MediaInfo) Reset() {
	*x = MediaInfo{}
	mi := &file_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MediaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaInfo) ProtoMessage() {}

func (x *MediaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaInfo.ProtoReflect.Descriptor instead.
func (*MediaInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *MediaInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MediaInfo) GetType() MediaType {
	if x != nil {
		return x.Type
	}
	return MediaType_MT_DEFAULT
}

func (x *MediaInfo) GetGetUrl() string {
	if x != nil {
		return x.GetUrl
	}
	return ""
}

func (x *MediaInfo) GetPutUrl() string {
	if x != nil {
		return x.PutUrl
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x0a, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x54, 0x0a, 0x0a, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x02, 0x73, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x52, 0x02, 0x73, 0x77, 0x12, 0x22, 0x0a, 0x02, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x52, 0x02, 0x6e, 0x65, 0x22, 0x78,
	0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x2a, 0x2e, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x54, 0x5f, 0x50, 0x4f, 0x46, 0x50,
	0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x4f, 0x47, 0x54, 0x54, 0x2f, 0x64, 0x6d, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_proto_goTypes = []any{
	(MediaType)(0),     // 0: common.MediaType
	(*PointCoord)(nil), // 1: common.PointCoord
	(*BoundCoord)(nil), // 2: common.BoundCoord
	(*MediaInfo)(nil),  // 3: common.MediaInfo
}
var file_common_proto_depIdxs = []int32{
	1, // 0: common.BoundCoord.sw:type_name -> common.PointCoord
	1, // 1: common.BoundCoord.ne:type_name -> common.PointCoord
	0, // 2: common.MediaInfo.type:type_name -> common.MediaType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
