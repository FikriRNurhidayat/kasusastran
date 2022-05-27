// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: common.pagination.proto

package kasusastran

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

type RequestPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *RequestPagination) Reset() {
	*x = RequestPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPagination) ProtoMessage() {}

func (x *RequestPagination) ProtoReflect() protoreflect.Message {
	mi := &file_common_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPagination.ProtoReflect.Descriptor instead.
func (*RequestPagination) Descriptor() ([]byte, []int) {
	return file_common_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *RequestPagination) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *RequestPagination) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ResponsePagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageCount uint32 `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	Total     uint32 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ResponsePagination) Reset() {
	*x = ResponsePagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePagination) ProtoMessage() {}

func (x *ResponsePagination) ProtoReflect() protoreflect.Message {
	mi := &file_common_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePagination.ProtoReflect.Descriptor instead.
func (*ResponsePagination) Descriptor() ([]byte, []int) {
	return file_common_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsePagination) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ResponsePagination) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ResponsePagination) GetPageCount() uint32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *ResponsePagination) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_common_pagination_proto protoreflect.FileDescriptor

var file_common_pagination_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x61, 0x73, 0x75, 0x73,
	0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x22, 0x44, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7a, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6b, 0x72, 0x69, 0x72, 0x6e, 0x75, 0x72,
	0x68, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x2f, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74,
	0x72, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74,
	0x72, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_pagination_proto_rawDescOnce sync.Once
	file_common_pagination_proto_rawDescData = file_common_pagination_proto_rawDesc
)

func file_common_pagination_proto_rawDescGZIP() []byte {
	file_common_pagination_proto_rawDescOnce.Do(func() {
		file_common_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pagination_proto_rawDescData)
	})
	return file_common_pagination_proto_rawDescData
}

var file_common_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_pagination_proto_goTypes = []interface{}{
	(*RequestPagination)(nil),  // 0: kasusastran.RequestPagination
	(*ResponsePagination)(nil), // 1: kasusastran.ResponsePagination
}
var file_common_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_pagination_proto_init() }
func file_common_pagination_proto_init() {
	if File_common_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPagination); i {
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
		file_common_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePagination); i {
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
			RawDescriptor: file_common_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_pagination_proto_goTypes,
		DependencyIndexes: file_common_pagination_proto_depIdxs,
		MessageInfos:      file_common_pagination_proto_msgTypes,
	}.Build()
	File_common_pagination_proto = out.File
	file_common_pagination_proto_rawDesc = nil
	file_common_pagination_proto_goTypes = nil
	file_common_pagination_proto_depIdxs = nil
}
