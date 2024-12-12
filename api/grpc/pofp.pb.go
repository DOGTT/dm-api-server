// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.14.0
// source: pofp.proto

package api

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

// 互动类型
type InteractionType int32

const (
	// 点赞
	InteractionType_LIKE InteractionType = 0
	// 标记
	InteractionType_MARK InteractionType = 1
)

// Enum value maps for InteractionType.
var (
	InteractionType_name = map[int32]string{
		0: "LIKE",
		1: "MARK",
	}
	InteractionType_value = map[string]int32{
		"LIKE": 0,
		"MARK": 1,
	}
)

func (x InteractionType) Enum() *InteractionType {
	p := new(InteractionType)
	*p = x
	return p
}

func (x InteractionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InteractionType) Descriptor() protoreflect.EnumDescriptor {
	return file_pofp_proto_enumTypes[0].Descriptor()
}

func (InteractionType) Type() protoreflect.EnumType {
	return &file_pofp_proto_enumTypes[0]
}

func (x InteractionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InteractionType.Descriptor instead.
func (InteractionType) EnumDescriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{0}
}

type PointCoord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float32 `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *PointCoord) Reset() {
	*x = PointCoord{}
	mi := &file_pofp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PointCoord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointCoord) ProtoMessage() {}

func (x *PointCoord) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PointCoord.ProtoReflect.Descriptor instead.
func (*PointCoord) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{0}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sw *PointCoord `protobuf:"bytes,1,opt,name=sw,proto3" json:"sw,omitempty"`
	Ne *PointCoord `protobuf:"bytes,2,opt,name=ne,proto3" json:"ne,omitempty"`
}

func (x *BoundCoord) Reset() {
	*x = BoundCoord{}
	mi := &file_pofp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoundCoord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundCoord) ProtoMessage() {}

func (x *BoundCoord) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BoundCoord.ProtoReflect.Descriptor instead.
func (*BoundCoord) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{1}
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

type PofpInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	TypeId      uint32                 `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Pid         uint32                 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Title       string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	LatLng      *PointCoord            `protobuf:"bytes,5,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
	Photos      []string               `protobuf:"bytes,6,rep,name=photos,proto3" json:"photos,omitempty"`
	Content     string                 `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	PoiId       string                 `protobuf:"bytes,8,opt,name=poi_id,json=poiId,proto3" json:"poi_id,omitempty"`
	Address     string                 `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	PoiData     map[string]string      `protobuf:"bytes,10,rep,name=poi_data,json=poiData,proto3" json:"poi_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ViewsCnt    int32                  `protobuf:"varint,11,opt,name=views_cnt,json=viewsCnt,proto3" json:"views_cnt,omitempty"`
	LikesCnt    int32                  `protobuf:"varint,12,opt,name=likes_cnt,json=likesCnt,proto3" json:"likes_cnt,omitempty"`
	MarksCnt    int32                  `protobuf:"varint,13,opt,name=marks_cnt,json=marksCnt,proto3" json:"marks_cnt,omitempty"`
	CommentsCnt int32                  `protobuf:"varint,14,opt,name=comments_cnt,json=commentsCnt,proto3" json:"comments_cnt,omitempty"`
	LastView    *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=last_view,json=lastView,proto3" json:"last_view,omitempty"`
	LastMark    *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=last_mark,json=lastMark,proto3" json:"last_mark,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PofpInfo) Reset() {
	*x = PofpInfo{}
	mi := &file_pofp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PofpInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PofpInfo) ProtoMessage() {}

func (x *PofpInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PofpInfo.ProtoReflect.Descriptor instead.
func (*PofpInfo) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{2}
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

func (x *PofpInfo) GetPid() uint32 {
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

func (x *PofpInfo) GetLatLng() *PointCoord {
	if x != nil {
		return x.LatLng
	}
	return nil
}

func (x *PofpInfo) GetPhotos() []string {
	if x != nil {
		return x.Photos
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

func (x *PofpInfo) GetLastView() *timestamppb.Timestamp {
	if x != nil {
		return x.LastView
	}
	return nil
}

func (x *PofpInfo) GetLastMark() *timestamppb.Timestamp {
	if x != nil {
		return x.LastMark
	}
	return nil
}

func (x *PofpInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PofpInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PofpDynamicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *PofpDynamicInfo) Reset() {
	*x = PofpDynamicInfo{}
	mi := &file_pofp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PofpDynamicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PofpDynamicInfo) ProtoMessage() {}

func (x *PofpDynamicInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PofpDynamicInfo.ProtoReflect.Descriptor instead.
func (*PofpDynamicInfo) Descriptor() ([]byte, []int) {
	return file_pofp_proto_rawDescGZIP(), []int{3}
}

func (x *PofpDynamicInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type PofpCommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ParentUuid string                 `protobuf:"bytes,2,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"`
	Content    string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *PofpCommentInfo) Reset() {
	*x = PofpCommentInfo{}
	mi := &file_pofp_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PofpCommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PofpCommentInfo) ProtoMessage() {}

func (x *PofpCommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pofp_proto_msgTypes[4]
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
	return file_pofp_proto_rawDescGZIP(), []int{4}
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

func (x *PofpCommentInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PofpCommentInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_pofp_proto protoreflect.FileDescriptor

var file_pofp_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x66, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f,
	0x66, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0a, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6c, 0x6e, 0x67, 0x22, 0x50, 0x0a, 0x0a, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x02, 0x73, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x6f, 0x66, 0x70, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x52, 0x02, 0x73, 0x77, 0x12, 0x20, 0x0a, 0x02, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x66, 0x70, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f,
	0x6f, 0x72, 0x64, 0x52, 0x02, 0x6e, 0x65, 0x22, 0xc3, 0x05, 0x0a, 0x08, 0x50, 0x6f, 0x66, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x6c, 0x61, 0x74,
	0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x6f, 0x66,
	0x70, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x6c, 0x61,
	0x74, 0x4c, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6f, 0x66, 0x70,
	0x2e, 0x50, 0x6f, 0x66, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6f, 0x69, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1b, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x76, 0x69, 0x65, 0x77, 0x73, 0x43, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x43, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72,
	0x6b, 0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61,
	0x72, 0x6b, 0x73, 0x43, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x50, 0x6f, 0x69, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x25, 0x0a,
	0x0f, 0x50, 0x6f, 0x66, 0x70, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x66, 0x70, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x25, 0x0a,
	0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41,
	0x52, 0x4b, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x4f, 0x47, 0x54, 0x54, 0x2f, 0x64, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pofp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pofp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pofp_proto_goTypes = []any{
	(InteractionType)(0),          // 0: pofp.InteractionType
	(*PointCoord)(nil),            // 1: pofp.PointCoord
	(*BoundCoord)(nil),            // 2: pofp.BoundCoord
	(*PofpInfo)(nil),              // 3: pofp.PofpInfo
	(*PofpDynamicInfo)(nil),       // 4: pofp.PofpDynamicInfo
	(*PofpCommentInfo)(nil),       // 5: pofp.PofpCommentInfo
	nil,                           // 6: pofp.PofpInfo.PoiDataEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_pofp_proto_depIdxs = []int32{
	1,  // 0: pofp.BoundCoord.sw:type_name -> pofp.PointCoord
	1,  // 1: pofp.BoundCoord.ne:type_name -> pofp.PointCoord
	1,  // 2: pofp.PofpInfo.lat_lng:type_name -> pofp.PointCoord
	6,  // 3: pofp.PofpInfo.poi_data:type_name -> pofp.PofpInfo.PoiDataEntry
	7,  // 4: pofp.PofpInfo.last_view:type_name -> google.protobuf.Timestamp
	7,  // 5: pofp.PofpInfo.last_mark:type_name -> google.protobuf.Timestamp
	7,  // 6: pofp.PofpInfo.created_at:type_name -> google.protobuf.Timestamp
	7,  // 7: pofp.PofpInfo.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 8: pofp.PofpCommentInfo.created_at:type_name -> google.protobuf.Timestamp
	7,  // 9: pofp.PofpCommentInfo.updated_at:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pofp_proto_init() }
func file_pofp_proto_init() {
	if File_pofp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pofp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pofp_proto_goTypes,
		DependencyIndexes: file_pofp_proto_depIdxs,
		EnumInfos:         file_pofp_proto_enumTypes,
		MessageInfos:      file_pofp_proto_msgTypes,
	}.Build()
	File_pofp_proto = out.File
	file_pofp_proto_rawDesc = nil
	file_pofp_proto_goTypes = nil
	file_pofp_proto_depIdxs = nil
}