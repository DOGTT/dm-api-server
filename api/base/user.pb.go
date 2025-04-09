// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.14.0
// source: user.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserPets      []*UserPetInfo         `protobuf:"bytes,3,rep,name=user_pets,json=userPets,proto3" json:"user_pets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetUserPets() []*UserPetInfo {
	if x != nil {
		return x.UserPets
	}
	return nil
}

type UserPetInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pet           *PetInfo               `protobuf:"bytes,1,opt,name=pet,proto3" json:"pet,omitempty"`
	PetTitle      string                 `protobuf:"bytes,2,opt,name=pet_title,json=petTitle,proto3" json:"pet_title,omitempty"`
	PetStatus     int32                  `protobuf:"varint,3,opt,name=pet_status,json=petStatus,proto3" json:"pet_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserPetInfo) Reset() {
	*x = UserPetInfo{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserPetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPetInfo) ProtoMessage() {}

func (x *UserPetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPetInfo.ProtoReflect.Descriptor instead.
func (*UserPetInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserPetInfo) GetPet() *PetInfo {
	if x != nil {
		return x.Pet
	}
	return nil
}

func (x *UserPetInfo) GetPetTitle() string {
	if x != nil {
		return x.PetTitle
	}
	return ""
}

func (x *UserPetInfo) GetPetStatus() int32 {
	if x != nil {
		return x.PetStatus
	}
	return 0
}

// 宠物快速注册信息
type FastRegisterData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 名字
	PetName string `protobuf:"bytes,1,opt,name=pet_name,json=petName,proto3" json:"pet_name,omitempty"`
	// 宠物头像 base64 data
	PetAvatarData string `protobuf:"bytes,2,opt,name=pet_avatar_data,json=petAvatarData,proto3" json:"pet_avatar_data,omitempty"`
	// 宠物和人的关系
	PetTitle      string `protobuf:"bytes,3,opt,name=pet_title,json=petTitle,proto3" json:"pet_title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FastRegisterData) Reset() {
	*x = FastRegisterData{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FastRegisterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FastRegisterData) ProtoMessage() {}

func (x *FastRegisterData) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FastRegisterData.ProtoReflect.Descriptor instead.
func (*FastRegisterData) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *FastRegisterData) GetPetName() string {
	if x != nil {
		return x.PetName
	}
	return ""
}

func (x *FastRegisterData) GetPetAvatarData() string {
	if x != nil {
		return x.PetAvatarData
	}
	return ""
}

func (x *FastRegisterData) GetPetTitle() string {
	if x != nil {
		return x.PetTitle
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x09, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x50, 0x65, 0x74, 0x73, 0x22, 0x69, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x50, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x03, 0x70, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x74, 0x5f,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x74,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x72, 0x0a, 0x10, 0x46, 0x61, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x65, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x65,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x65, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x65, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x4f, 0x47, 0x54, 0x54, 0x2f, 0x64, 0x6d, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_proto_goTypes = []any{
	(*UserInfo)(nil),         // 0: user.UserInfo
	(*UserPetInfo)(nil),      // 1: user.UserPetInfo
	(*FastRegisterData)(nil), // 2: user.FastRegisterData
	(*PetInfo)(nil),          // 3: pet.PetInfo
}
var file_user_proto_depIdxs = []int32{
	1, // 0: user.UserInfo.user_pets:type_name -> user.UserPetInfo
	3, // 1: user.UserPetInfo.pet:type_name -> pet.PetInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_pet_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
