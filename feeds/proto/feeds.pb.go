// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: proto/feeds.proto

package feeds

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rss feed name
	// eg. a16z
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// rss feed url
	// eg. http://a16z.com/feed/
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// category of the feed
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{0}
}

func (x *Feed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feed) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Feed) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain   string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Url      string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Date     int64  `protobuf:"varint,6,opt,name=date,proto3" json:"date,omitempty"`
	Category string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entry) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Entry) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Entry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Entry) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Entry) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Entry) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rss feed name
	// eg. a16z
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// rss feed url
	// eg. http://a16z.com/feed/
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// category to add
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{3}
}

type EntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rss feed url
	// eg. http://a16z.com/feed/
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *EntriesRequest) Reset() {
	*x = EntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntriesRequest) ProtoMessage() {}

func (x *EntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntriesRequest.ProtoReflect.Descriptor instead.
func (*EntriesRequest) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{4}
}

func (x *EntriesRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type EntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *EntriesResponse) Reset() {
	*x = EntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntriesResponse) ProtoMessage() {}

func (x *EntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntriesResponse.ProtoReflect.Descriptor instead.
func (*EntriesResponse) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{5}
}

func (x *EntriesResponse) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{6}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feeds []*Feed `protobuf:"bytes,1,rep,name=feeds,proto3" json:"feeds,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetFeeds() []*Feed {
	if x != nil {
		return x.Feeds
	}
	return nil
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rss feed name
	// eg. a16z
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feeds_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feeds_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_proto_feeds_proto_rawDescGZIP(), []int{9}
}

var File_proto_feeds_proto protoreflect.FileDescriptor

var file_proto_feeds_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x22, 0x48, 0x0a, 0x04, 0x46, 0x65,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x4e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x39, 0x0a, 0x0f, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xdf, 0x01, 0x0a, 0x05, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x03, 0x41, 0x64, 0x64,
	0x12, 0x11, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x15, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x66, 0x65, 0x65, 0x64, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_feeds_proto_rawDescOnce sync.Once
	file_proto_feeds_proto_rawDescData = file_proto_feeds_proto_rawDesc
)

func file_proto_feeds_proto_rawDescGZIP() []byte {
	file_proto_feeds_proto_rawDescOnce.Do(func() {
		file_proto_feeds_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_feeds_proto_rawDescData)
	})
	return file_proto_feeds_proto_rawDescData
}

var file_proto_feeds_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_feeds_proto_goTypes = []interface{}{
	(*Feed)(nil),            // 0: feeds.Feed
	(*Entry)(nil),           // 1: feeds.Entry
	(*AddRequest)(nil),      // 2: feeds.AddRequest
	(*AddResponse)(nil),     // 3: feeds.AddResponse
	(*EntriesRequest)(nil),  // 4: feeds.EntriesRequest
	(*EntriesResponse)(nil), // 5: feeds.EntriesResponse
	(*ListRequest)(nil),     // 6: feeds.ListRequest
	(*ListResponse)(nil),    // 7: feeds.ListResponse
	(*RemoveRequest)(nil),   // 8: feeds.RemoveRequest
	(*RemoveResponse)(nil),  // 9: feeds.RemoveResponse
}
var file_proto_feeds_proto_depIdxs = []int32{
	1, // 0: feeds.EntriesResponse.entries:type_name -> feeds.Entry
	0, // 1: feeds.ListResponse.feeds:type_name -> feeds.Feed
	2, // 2: feeds.Feeds.Add:input_type -> feeds.AddRequest
	8, // 3: feeds.Feeds.Remove:input_type -> feeds.RemoveRequest
	4, // 4: feeds.Feeds.Entries:input_type -> feeds.EntriesRequest
	6, // 5: feeds.Feeds.List:input_type -> feeds.ListRequest
	3, // 6: feeds.Feeds.Add:output_type -> feeds.AddResponse
	9, // 7: feeds.Feeds.Remove:output_type -> feeds.RemoveResponse
	5, // 8: feeds.Feeds.Entries:output_type -> feeds.EntriesResponse
	7, // 9: feeds.Feeds.List:output_type -> feeds.ListResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_feeds_proto_init() }
func file_proto_feeds_proto_init() {
	if File_proto_feeds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_feeds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feed); i {
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
		file_proto_feeds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_proto_feeds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_proto_feeds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_proto_feeds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntriesRequest); i {
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
		file_proto_feeds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntriesResponse); i {
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
		file_proto_feeds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_feeds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_proto_feeds_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_proto_feeds_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveResponse); i {
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
			RawDescriptor: file_proto_feeds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_feeds_proto_goTypes,
		DependencyIndexes: file_proto_feeds_proto_depIdxs,
		MessageInfos:      file_proto_feeds_proto_msgTypes,
	}.Build()
	File_proto_feeds_proto = out.File
	file_proto_feeds_proto_rawDesc = nil
	file_proto_feeds_proto_goTypes = nil
	file_proto_feeds_proto_depIdxs = nil
}
