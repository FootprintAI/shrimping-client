//shritagram is shrimping + instagram

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: shritagram.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ObjectCompressionAlgorithm int32

const (
	ObjectCompressionAlgorithm_NONE ObjectCompressionAlgorithm = 0
	ObjectCompressionAlgorithm_GZIP ObjectCompressionAlgorithm = 1
)

// Enum value maps for ObjectCompressionAlgorithm.
var (
	ObjectCompressionAlgorithm_name = map[int32]string{
		0: "NONE",
		1: "GZIP",
	}
	ObjectCompressionAlgorithm_value = map[string]int32{
		"NONE": 0,
		"GZIP": 1,
	}
)

func (x ObjectCompressionAlgorithm) Enum() *ObjectCompressionAlgorithm {
	p := new(ObjectCompressionAlgorithm)
	*p = x
	return p
}

func (x ObjectCompressionAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectCompressionAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_shritagram_proto_enumTypes[0].Descriptor()
}

func (ObjectCompressionAlgorithm) Type() protoreflect.EnumType {
	return &file_shritagram_proto_enumTypes[0]
}

func (x ObjectCompressionAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectCompressionAlgorithm.Descriptor instead.
func (ObjectCompressionAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{0}
}

type CallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallbackRequest) Reset() {
	*x = CallbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackRequest) ProtoMessage() {}

func (x *CallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackRequest.ProtoReflect.Descriptor instead.
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{0}
}

type ProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usernames []string `protobuf:"bytes,1,rep,name=usernames,proto3" json:"usernames,omitempty"`
	// cacheControl has possible values:
	// no-cache: the request won't go through backend server's cache, it is revalidated with the actualy service (i.g. instagram)
	// max-age=604800: the request's result will be kept in cache with the duration(in s) specified
	//
	// default value is  max-age=86400
	CacheControl string `protobuf:"bytes,2,opt,name=cacheControl,proto3" json:"cacheControl,omitempty"`
	// priority has three possible values: high/median/low.
	// the server will execute jobs according to its priority.
	Priority string `protobuf:"bytes,3,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileRequest) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

func (x *ProfileRequest) GetCacheControl() string {
	if x != nil {
		return x.CacheControl
	}
	return ""
}

func (x *ProfileRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

type PostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shortcodes []string `protobuf:"bytes,1,rep,name=shortcodes,proto3" json:"shortcodes,omitempty"`
	// cacheControl has possible values:
	// no-cache: the request won't go through backend server's cache, it is revalidated with the actualy service (i.g. instagram)
	// max-age=604800: the request's result will be kept in cache with the duration(in s) specified
	//
	// default value is  max-age=86400
	CacheControl string `protobuf:"bytes,2,opt,name=cacheControl,proto3" json:"cacheControl,omitempty"`
	// priority has three possible values: high/median/low.
	// the server will execute jobs according to its priority.
	Priority string `protobuf:"bytes,3,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *PostRequest) Reset() {
	*x = PostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequest) ProtoMessage() {}

func (x *PostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequest.ProtoReflect.Descriptor instead.
func (*PostRequest) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{2}
}

func (x *PostRequest) GetShortcodes() []string {
	if x != nil {
		return x.Shortcodes
	}
	return nil
}

func (x *PostRequest) GetCacheControl() string {
	if x != nil {
		return x.CacheControl
	}
	return ""
}

func (x *PostRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

type ShritagramResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Compression   ObjectCompressionAlgorithm `protobuf:"varint,1,opt,name=compression,proto3,enum=pb.ObjectCompressionAlgorithm" json:"compression,omitempty"`
	RawProfiles   []*RawProfileObject        `protobuf:"bytes,2,rep,name=rawProfiles,proto3" json:"rawProfiles,omitempty"`
	RawPosts      []*RawPostObject           `protobuf:"bytes,3,rep,name=rawPosts,proto3" json:"rawPosts,omitempty"`
	RawTopSearchs []*RawTopSearchObject      `protobuf:"bytes,4,rep,name=rawTopSearchs,proto3" json:"rawTopSearchs,omitempty"`
}

func (x *ShritagramResponse) Reset() {
	*x = ShritagramResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShritagramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShritagramResponse) ProtoMessage() {}

func (x *ShritagramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShritagramResponse.ProtoReflect.Descriptor instead.
func (*ShritagramResponse) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{3}
}

func (x *ShritagramResponse) GetCompression() ObjectCompressionAlgorithm {
	if x != nil {
		return x.Compression
	}
	return ObjectCompressionAlgorithm_NONE
}

func (x *ShritagramResponse) GetRawProfiles() []*RawProfileObject {
	if x != nil {
		return x.RawProfiles
	}
	return nil
}

func (x *ShritagramResponse) GetRawPosts() []*RawPostObject {
	if x != nil {
		return x.RawPosts
	}
	return nil
}

func (x *ShritagramResponse) GetRawTopSearchs() []*RawTopSearchObject {
	if x != nil {
		return x.RawTopSearchs
	}
	return nil
}

type RawProfileObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	RawBytes []byte `protobuf:"bytes,2,opt,name=rawBytes,proto3" json:"rawBytes,omitempty"`
}

func (x *RawProfileObject) Reset() {
	*x = RawProfileObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawProfileObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawProfileObject) ProtoMessage() {}

func (x *RawProfileObject) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawProfileObject.ProtoReflect.Descriptor instead.
func (*RawProfileObject) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{4}
}

func (x *RawProfileObject) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RawProfileObject) GetRawBytes() []byte {
	if x != nil {
		return x.RawBytes
	}
	return nil
}

type RawPostObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shortcode string `protobuf:"bytes,1,opt,name=shortcode,proto3" json:"shortcode,omitempty"`
	RawBytes  []byte `protobuf:"bytes,2,opt,name=rawBytes,proto3" json:"rawBytes,omitempty"`
}

func (x *RawPostObject) Reset() {
	*x = RawPostObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawPostObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawPostObject) ProtoMessage() {}

func (x *RawPostObject) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawPostObject.ProtoReflect.Descriptor instead.
func (*RawPostObject) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{5}
}

func (x *RawPostObject) GetShortcode() string {
	if x != nil {
		return x.Shortcode
	}
	return ""
}

func (x *RawPostObject) GetRawBytes() []byte {
	if x != nil {
		return x.RawBytes
	}
	return nil
}

type TopSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashtags []string `protobuf:"bytes,1,rep,name=hashtags,proto3" json:"hashtags,omitempty"` // without `#`
	// cacheControl has possible values:
	// no-cache: the request won't go through backend server's cache, it is revalidated with the actualy service (i.g. instagram)
	// max-age=604800: the request's result will be kept in cache with the duration(in s) specified
	//
	// default value is  max-age=86400
	CacheControl string `protobuf:"bytes,2,opt,name=cacheControl,proto3" json:"cacheControl,omitempty"`
	// priority has three possible values: high/median/low.
	// the server will execute jobs according to its priority.
	Priority string `protobuf:"bytes,3,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *TopSearchRequest) Reset() {
	*x = TopSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopSearchRequest) ProtoMessage() {}

func (x *TopSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopSearchRequest.ProtoReflect.Descriptor instead.
func (*TopSearchRequest) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{6}
}

func (x *TopSearchRequest) GetHashtags() []string {
	if x != nil {
		return x.Hashtags
	}
	return nil
}

func (x *TopSearchRequest) GetCacheControl() string {
	if x != nil {
		return x.CacheControl
	}
	return ""
}

func (x *TopSearchRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

type RawTopSearchObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashtag  string `protobuf:"bytes,1,opt,name=hashtag,proto3" json:"hashtag,omitempty"`
	RawBytes []byte `protobuf:"bytes,2,opt,name=rawBytes,proto3" json:"rawBytes,omitempty"`
}

func (x *RawTopSearchObject) Reset() {
	*x = RawTopSearchObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shritagram_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawTopSearchObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawTopSearchObject) ProtoMessage() {}

func (x *RawTopSearchObject) ProtoReflect() protoreflect.Message {
	mi := &file_shritagram_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawTopSearchObject.ProtoReflect.Descriptor instead.
func (*RawTopSearchObject) Descriptor() ([]byte, []int) {
	return file_shritagram_proto_rawDescGZIP(), []int{7}
}

func (x *RawTopSearchObject) GetHashtag() string {
	if x != nil {
		return x.Hashtag
	}
	return ""
}

func (x *RawTopSearchObject) GetRawBytes() []byte {
	if x != nil {
		return x.RawBytes
	}
	return nil
}

var File_shritagram_proto protoreflect.FileDescriptor

var file_shritagram_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x68, 0x72, 0x69, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x6d, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xfb, 0x01, 0x0a, 0x12, 0x53, 0x68, 0x72, 0x69, 0x74,
	0x61, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x36, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x61, 0x77, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x72, 0x61, 0x77, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x61, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x72, 0x61,
	0x77, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x72, 0x61, 0x77, 0x54, 0x6f, 0x70,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x61, 0x77, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0d, 0x72, 0x61, 0x77, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x73, 0x22, 0x4a, 0x0a, 0x10, 0x52, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x22, 0x49, 0x0a, 0x0d, 0x52, 0x61, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x72, 0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x10, 0x54,
	0x6f, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x4a, 0x0a, 0x12, 0x52,
	0x61, 0x77, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72,
	0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x2a, 0x30, 0x0a, 0x1a, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x01, 0x32, 0xe1, 0x02, 0x0a, 0x0a, 0x73, 0x68,
	0x72, 0x69, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x59, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x68, 0x72, 0x69, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x3a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3a, 0x01,
	0x2a, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x72, 0x69, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x3a, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x4a, 0x0a, 0x05, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x72, 0x69, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x3a, 0x67, 0x65, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x09, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x72, 0x69,
	0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0xd6, 0x04,
	0x5a, 0x03, 0x2f, 0x70, 0x62, 0x92, 0x41, 0xcd, 0x04, 0x12, 0xf6, 0x01, 0x0a, 0x1d, 0x53, 0x68,
	0x72, 0x69, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x20, 0x41, 0x50, 0x49, 0x20, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6f, 0x0a, 0x20, 0x73,
	0x68, 0x72, 0x69, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x20, 0x72, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x69, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f, 0x6f, 0x74, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x2f,
	0x73, 0x68, 0x72, 0x69, 0x6d, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x1a, 0x1a, 0x73, 0x68, 0x72, 0x69, 0x6d, 0x70, 0x69, 0x6e, 0x67, 0x40, 0x66, 0x6f, 0x6f, 0x74,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x2d, 0x61, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x5f, 0x0a, 0x14,
	0x42, 0x53, 0x44, 0x20, 0x33, 0x2d, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x20, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f, 0x6f, 0x74, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x61, 0x69, 0x2f, 0x73, 0x68, 0x72, 0x69, 0x6d, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2e, 0x74, 0x78, 0x74, 0x32, 0x03, 0x31,
	0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30,
	0x33, 0x12, 0x49, 0x0a, 0x47, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03,
	0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x1d, 0x0a, 0x03, 0x35, 0x30, 0x30,
	0x12, 0x16, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x5a, 0x23, 0x0a, 0x21, 0x0a, 0x0a, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x10, 0x0a,
	0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x72,
	0x45, 0x0a, 0x12, 0x73, 0x68, 0x72, 0x69, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x20, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x69, 0x12, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f, 0x6f, 0x74, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x61, 0x69, 0x2f, 0x73, 0x68, 0x72, 0x69, 0x6d, 0x70, 0x69, 0x6e, 0x67, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shritagram_proto_rawDescOnce sync.Once
	file_shritagram_proto_rawDescData = file_shritagram_proto_rawDesc
)

func file_shritagram_proto_rawDescGZIP() []byte {
	file_shritagram_proto_rawDescOnce.Do(func() {
		file_shritagram_proto_rawDescData = protoimpl.X.CompressGZIP(file_shritagram_proto_rawDescData)
	})
	return file_shritagram_proto_rawDescData
}

var file_shritagram_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shritagram_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_shritagram_proto_goTypes = []interface{}{
	(ObjectCompressionAlgorithm)(0), // 0: pb.ObjectCompressionAlgorithm
	(*CallbackRequest)(nil),         // 1: pb.CallbackRequest
	(*ProfileRequest)(nil),          // 2: pb.ProfileRequest
	(*PostRequest)(nil),             // 3: pb.PostRequest
	(*ShritagramResponse)(nil),      // 4: pb.ShritagramResponse
	(*RawProfileObject)(nil),        // 5: pb.RawProfileObject
	(*RawPostObject)(nil),           // 6: pb.RawPostObject
	(*TopSearchRequest)(nil),        // 7: pb.TopSearchRequest
	(*RawTopSearchObject)(nil),      // 8: pb.RawTopSearchObject
}
var file_shritagram_proto_depIdxs = []int32{
	0, // 0: pb.ShritagramResponse.compression:type_name -> pb.ObjectCompressionAlgorithm
	5, // 1: pb.ShritagramResponse.rawProfiles:type_name -> pb.RawProfileObject
	6, // 2: pb.ShritagramResponse.rawPosts:type_name -> pb.RawPostObject
	8, // 3: pb.ShritagramResponse.rawTopSearchs:type_name -> pb.RawTopSearchObject
	1, // 4: pb.shritagram.Callback:input_type -> pb.CallbackRequest
	2, // 5: pb.shritagram.Profile:input_type -> pb.ProfileRequest
	3, // 6: pb.shritagram.Posts:input_type -> pb.PostRequest
	7, // 7: pb.shritagram.TopSearch:input_type -> pb.TopSearchRequest
	4, // 8: pb.shritagram.Callback:output_type -> pb.ShritagramResponse
	4, // 9: pb.shritagram.Profile:output_type -> pb.ShritagramResponse
	4, // 10: pb.shritagram.Posts:output_type -> pb.ShritagramResponse
	4, // 11: pb.shritagram.TopSearch:output_type -> pb.ShritagramResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_shritagram_proto_init() }
func file_shritagram_proto_init() {
	if File_shritagram_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shritagram_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackRequest); i {
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
		file_shritagram_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileRequest); i {
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
		file_shritagram_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequest); i {
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
		file_shritagram_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShritagramResponse); i {
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
		file_shritagram_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawProfileObject); i {
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
		file_shritagram_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawPostObject); i {
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
		file_shritagram_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopSearchRequest); i {
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
		file_shritagram_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawTopSearchObject); i {
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
			RawDescriptor: file_shritagram_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shritagram_proto_goTypes,
		DependencyIndexes: file_shritagram_proto_depIdxs,
		EnumInfos:         file_shritagram_proto_enumTypes,
		MessageInfos:      file_shritagram_proto_msgTypes,
	}.Build()
	File_shritagram_proto = out.File
	file_shritagram_proto_rawDesc = nil
	file_shritagram_proto_goTypes = nil
	file_shritagram_proto_depIdxs = nil
}
