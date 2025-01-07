// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.14.0
// source: pofp.proto

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

// 足迹点类型
type PofpTypeInfo struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CoverageRadius int32                  `protobuf:"varint,3,opt,name=coverage_radius,json=coverageRadius,proto3" json:"coverage_radius,omitempty"`
	ThemeColor     string                 `protobuf:"bytes,4,opt,name=theme_color,json=themeColor,proto3" json:"theme_color,omitempty"`
	CreatedAt      int64                  `protobuf:"varint,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      int64                  `protobuf:"varint,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PofpTypeInfo) Reset() {
	*x = PofpTypeInfo{}
	mi := &file_pofp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PofpTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PofpTypeInfo) ProtoMessage() {}

func (x *PofpTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pofp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PofpTypeInfo.ProtoReflect.Descriptor instead.
func (*PofpTypeInfo) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{0}
}

func (x *PofpTypeInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PofpTypeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PofpTypeInfo) GetCoverageRadius() int32 {
	if x != nil {
		return x.CoverageRadius
	}
	return 0
}

func (x *PofpTypeInfo) GetThemeColor() string {
	if x != nil {
		return x.ThemeColor
	}
	return ""
}

func (x *PofpTypeInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PofpTypeInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type PofpInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 足迹 ID
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// 足迹类型, 不可更新
	TypeId uint32 `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	// 足迹作者, 不可更新
	Pid uint64 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	// 足迹名称, 可更新
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// 足迹位置, 不可更新
	LngLat *PointCoord `protobuf:"bytes,5,opt,name=lng_lat,json=lngLat,proto3" json:"lng_lat,omitempty"`
	// 媒体信息
	Media []*MediaInfo `protobuf:"bytes,6,rep,name=media,proto3" json:"media,omitempty"`
	// 内容, 可更新
	Content string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	// POI ID, 不可更新
	PoiId string `protobuf:"bytes,8,opt,name=poi_id,json=poiId,proto3" json:"poi_id,omitempty"`
	// POI 地址, 不可更新
	Address string `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	// POI 详细信息, 不可更新
	PoiData map[string]string `protobuf:"bytes,10,rep,name=poi_data,json=poiData,proto3" json:"poi_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// 查看数，只读
	ViewsCnt int32 `protobuf:"varint,11,opt,name=views_cnt,json=viewsCnt,proto3" json:"views_cnt,omitempty"`
	// 喜欢数，只读
	LikesCnt int32 `protobuf:"varint,12,opt,name=likes_cnt,json=likesCnt,proto3" json:"likes_cnt,omitempty"`
	// 标记数，只读
	MarksCnt int32 `protobuf:"varint,13,opt,name=marks_cnt,json=marksCnt,proto3" json:"marks_cnt,omitempty"`
	// 评论数，只读
	CommentsCnt   int32 `protobuf:"varint,14,opt,name=comments_cnt,json=commentsCnt,proto3" json:"comments_cnt,omitempty"`
	LastView      int64 `protobuf:"varint,15,opt,name=last_view,json=lastView,proto3" json:"last_view,omitempty"`
	LastMark      int64 `protobuf:"varint,16,opt,name=last_mark,json=lastMark,proto3" json:"last_mark,omitempty"`
	CreatedAt     int64 `protobuf:"varint,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     int64 `protobuf:"varint,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PofpInfo) Reset() {
	*x = PofpInfo{}
	mi := &file_pofp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PofpInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PofpInfo) ProtoMessage() {}

func (x *PofpInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pofp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PofpInfo.ProtoReflect.Descriptor instead.
func (*PofpInfo) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{1}
}

func (x *PofpInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PofpInfo) GetTypeId() uint32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *PofpInfo) GetPid() uint64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *PofpInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PofpInfo) GetLngLat() *PointCoord {
	if x != nil {
		return x.LngLat
	}
	return nil
}

func (x *PofpInfo) GetMedia() []*MediaInfo {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *PofpInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PofpInfo) GetPoiId() string {
	if x != nil {
		return x.PoiId
	}
	return ""
}

func (x *PofpInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PofpInfo) GetPoiData() map[string]string {
	if x != nil {
		return x.PoiData
	}
	return nil
}

func (x *PofpInfo) GetViewsCnt() int32 {
	if x != nil {
		return x.ViewsCnt
	}
	return 0
}

func (x *PofpInfo) GetLikesCnt() int32 {
	if x != nil {
		return x.LikesCnt
	}
	return 0
}

func (x *PofpInfo) GetMarksCnt() int32 {
	if x != nil {
		return x.MarksCnt
	}
	return 0
}

func (x *PofpInfo) GetCommentsCnt() int32 {
	if x != nil {
		return x.CommentsCnt
	}
	return 0
}

func (x *PofpInfo) GetLastView() int64 {
	if x != nil {
		return x.LastView
	}
	return 0
}

func (x *PofpInfo) GetLastMark() int64 {
	if x != nil {
		return x.LastMark
	}
	return 0
}

func (x *PofpInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PofpInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type PofpDynamicInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PofpDynamicInfo) Reset() {
	*x = PofpDynamicInfo{}
	mi := &file_pofp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PofpDynamicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PofpDynamicInfo) ProtoMessage() {}

func (x *PofpDynamicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pofp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PofpDynamicInfo.ProtoReflect.Descriptor instead.
func (*PofpDynamicInfo) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{2}
}

func (x *PofpDynamicInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 足迹评论
type PofpCommentInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ParentUuid    string                 `protobuf:"bytes,2,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     int64                  `protobuf:"varint,32,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PofpCommentInfo) Reset() {
	*x = PofpCommentInfo{}
	mi := &file_pofp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PofpCommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PofpCommentInfo) ProtoMessage() {}

func (x *PofpCommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pofp_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PofpCommentInfo.ProtoReflect.Descriptor instead.
func (*PofpCommentInfo) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{3}
}

func (x *PofpCommentInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PofpCommentInfo) GetParentUuid() string {
	if x != nil {
		return x.ParentUuid
	}
	return ""
}

func (x *PofpCommentInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PofpCommentInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PofpCommentInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_pofp_proto protoreflect.FileDescriptor

var file_pofp_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x66, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f,
	0x66, 0x70, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xba, 0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x66, 0x70, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe6, 0x04,
	0x0a, 0x08, 0x50, 0x6f, 0x66, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x6c, 0x6e, 0x67, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x6c, 0x6e, 0x67, 0x4c, 0x61, 0x74, 0x12, 0x27, 0x0a, 0x05,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6f, 0x69, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x36, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x66, 0x70, 0x2e, 0x50, 0x6f, 0x66, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x50, 0x6f, 0x69, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x70, 0x6f, 0x69, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x43, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x63,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x43,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x43, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x50, 0x6f,
	0x69, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x25, 0x0a, 0x0f, 0x50, 0x6f, 0x66, 0x70, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x9e, 0x01,
	0x0a, 0x0f, 0x50, 0x6f, 0x66, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x4f, 0x47,
	0x54, 0x54, 0x2f, 0x64, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pofp_proto_rawDescOnce sync.Once
	file_pofp_proto_rawDescData = file_pofp_proto_rawDesc
)

func file_pofp_proto_rawDescGZIP() []byte {
	file_pofp_proto_rawDescOnce.Do(func() {
		file_pofp_proto_rawDescData = protoimpl.X.CompressGZIP(file_pofp_proto_rawDescData)
	})
	return file_pofp_proto_rawDescData
}

var file_pofp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pofp_proto_goTypes = []any{
	(*PofpTypeInfo)(nil),    // 0: pofp.PofpTypeInfo
	(*PofpInfo)(nil),        // 1: pofp.PofpInfo
	(*PofpDynamicInfo)(nil), // 2: pofp.PofpDynamicInfo
	(*PofpCommentInfo)(nil), // 3: pofp.PofpCommentInfo
	nil,                     // 4: pofp.PofpInfo.PoiDataEntry
	(*PointCoord)(nil),      // 5: common.PointCoord
	(*MediaInfo)(nil),       // 6: common.MediaInfo
}
var file_pofp_proto_depIdxs = []int32{
	5, // 0: pofp.PofpInfo.lng_lat:type_name -> common.PointCoord
	6, // 1: pofp.PofpInfo.media:type_name -> common.MediaInfo
	4, // 2: pofp.PofpInfo.poi_data:type_name -> pofp.PofpInfo.PoiDataEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pofp_proto_init() }
func file_pofp_proto_init() {
	if File_pofp_proto != nil {
		return
	}
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pofp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pofp_proto_goTypes,
		DependencyIndexes: file_pofp_proto_depIdxs,
		MessageInfos:      file_pofp_proto_msgTypes,
	}.Build()
	File_pofp_proto = out.File
	file_pofp_proto_rawDesc = nil
	file_pofp_proto_goTypes = nil
	file_pofp_proto_depIdxs = nil
}
