// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.1
// source: test_messages.proto

package proto_path

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

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId             int32  `protobuf:"varint,1,opt,name=pageId,proto3" json:"pageId,omitempty"`
	PageTitle          string `protobuf:"bytes,2,opt,name=pageTitle,proto3" json:"pageTitle,omitempty"`
	TitleClass         int32  `protobuf:"varint,3,opt,name=titleClass,proto3" json:"titleClass,omitempty"`
	RevisionTextLength int32  `protobuf:"varint,4,opt,name=revisionTextLength,proto3" json:"revisionTextLength,omitempty"`
	RevisionTextLines  int32  `protobuf:"varint,5,opt,name=revisionTextLines,proto3" json:"revisionTextLines,omitempty"`
	RevisionDatetime   int64  `protobuf:"varint,6,opt,name=revisionDatetime,proto3" json:"revisionDatetime,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_test_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_test_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *Page) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *Page) GetTitleClass() int32 {
	if x != nil {
		return x.TitleClass
	}
	return 0
}

func (x *Page) GetRevisionTextLength() int32 {
	if x != nil {
		return x.RevisionTextLength
	}
	return 0
}

func (x *Page) GetRevisionTextLines() int32 {
	if x != nil {
		return x.RevisionTextLines
	}
	return 0
}

func (x *Page) GetRevisionDatetime() int64 {
	if x != nil {
		return x.RevisionDatetime
	}
	return 0
}

type FrameHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameId     uint32 `protobuf:"fixed32,1,opt,name=frameId,proto3" json:"frameId,omitempty"`
	Timestamp   uint64 `protobuf:"fixed64,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	FrameLength uint32 `protobuf:"fixed32,3,opt,name=frameLength,proto3" json:"frameLength,omitempty"`
}

func (x *FrameHeader) Reset() {
	*x = FrameHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameHeader) ProtoMessage() {}

func (x *FrameHeader) ProtoReflect() protoreflect.Message {
	mi := &file_test_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameHeader.ProtoReflect.Descriptor instead.
func (*FrameHeader) Descriptor() ([]byte, []int) {
	return file_test_messages_proto_rawDescGZIP(), []int{1}
}

func (x *FrameHeader) GetFrameId() uint32 {
	if x != nil {
		return x.FrameId
	}
	return 0
}

func (x *FrameHeader) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *FrameHeader) GetFrameLength() uint32 {
	if x != nil {
		return x.FrameLength
	}
	return 0
}

var File_test_messages_proto protoreflect.FileDescriptor

var file_test_messages_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x65, 0x78, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x78, 0x74, 0x4c, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x67,
	0x0a, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_messages_proto_rawDescOnce sync.Once
	file_test_messages_proto_rawDescData = file_test_messages_proto_rawDesc
)

func file_test_messages_proto_rawDescGZIP() []byte {
	file_test_messages_proto_rawDescOnce.Do(func() {
		file_test_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_messages_proto_rawDescData)
	})
	return file_test_messages_proto_rawDescData
}

var file_test_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_messages_proto_goTypes = []interface{}{
	(*Page)(nil),        // 0: Page
	(*FrameHeader)(nil), // 1: frameHeader
}
var file_test_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_messages_proto_init() }
func file_test_messages_proto_init() {
	if File_test_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_test_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameHeader); i {
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
			RawDescriptor: file_test_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_messages_proto_goTypes,
		DependencyIndexes: file_test_messages_proto_depIdxs,
		MessageInfos:      file_test_messages_proto_msgTypes,
	}.Build()
	File_test_messages_proto = out.File
	file_test_messages_proto_rawDesc = nil
	file_test_messages_proto_goTypes = nil
	file_test_messages_proto_depIdxs = nil
}
