// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: serats.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListSeratsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *RequestPagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListSeratsRequest) Reset() {
	*x = ListSeratsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSeratsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSeratsRequest) ProtoMessage() {}

func (x *ListSeratsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSeratsRequest.ProtoReflect.Descriptor instead.
func (*ListSeratsRequest) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{0}
}

func (x *ListSeratsRequest) GetPagination() *RequestPagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListSeratsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *ListSeratsResponse_MetaListSerats `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Serats []*Serat                           `protobuf:"bytes,2,rep,name=serats,proto3" json:"serats,omitempty"`
}

func (x *ListSeratsResponse) Reset() {
	*x = ListSeratsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSeratsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSeratsResponse) ProtoMessage() {}

func (x *ListSeratsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSeratsResponse.ProtoReflect.Descriptor instead.
func (*ListSeratsResponse) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{1}
}

func (x *ListSeratsResponse) GetMeta() *ListSeratsResponse_MetaListSerats {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListSeratsResponse) GetSerats() []*Serat {
	if x != nil {
		return x.Serats
	}
	return nil
}

type GetSeratRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSeratRequest) Reset() {
	*x = GetSeratRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeratRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeratRequest) ProtoMessage() {}

func (x *GetSeratRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeratRequest.ProtoReflect.Descriptor instead.
func (*GetSeratRequest) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{2}
}

func (x *GetSeratRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSeratRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSeratRequest) Reset() {
	*x = DeleteSeratRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSeratRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSeratRequest) ProtoMessage() {}

func (x *DeleteSeratRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSeratRequest.ProtoReflect.Descriptor instead.
func (*DeleteSeratRequest) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSeratRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateSeratRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title             string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CoverImageUrl     string `protobuf:"bytes,3,opt,name=cover_image_url,json=coverImageUrl,proto3" json:"cover_image_url,omitempty"`
	ThumbnailImageUrl string `protobuf:"bytes,4,opt,name=thumbnail_image_url,json=thumbnailImageUrl,proto3" json:"thumbnail_image_url,omitempty"`
}

func (x *CreateSeratRequest) Reset() {
	*x = CreateSeratRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSeratRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSeratRequest) ProtoMessage() {}

func (x *CreateSeratRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSeratRequest.ProtoReflect.Descriptor instead.
func (*CreateSeratRequest) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSeratRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateSeratRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateSeratRequest) GetCoverImageUrl() string {
	if x != nil {
		return x.CoverImageUrl
	}
	return ""
}

func (x *CreateSeratRequest) GetThumbnailImageUrl() string {
	if x != nil {
		return x.ThumbnailImageUrl
	}
	return ""
}

type UpdateSeratRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description       string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CoverImageUrl     string `protobuf:"bytes,4,opt,name=cover_image_url,json=coverImageUrl,proto3" json:"cover_image_url,omitempty"`
	ThumbnailImageUrl string `protobuf:"bytes,5,opt,name=thumbnail_image_url,json=thumbnailImageUrl,proto3" json:"thumbnail_image_url,omitempty"`
}

func (x *UpdateSeratRequest) Reset() {
	*x = UpdateSeratRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSeratRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSeratRequest) ProtoMessage() {}

func (x *UpdateSeratRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSeratRequest.ProtoReflect.Descriptor instead.
func (*UpdateSeratRequest) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSeratRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSeratRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateSeratRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateSeratRequest) GetCoverImageUrl() string {
	if x != nil {
		return x.CoverImageUrl
	}
	return ""
}

func (x *UpdateSeratRequest) GetThumbnailImageUrl() string {
	if x != nil {
		return x.ThumbnailImageUrl
	}
	return ""
}

type Serat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description       string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CoverImageUrl     string `protobuf:"bytes,5,opt,name=cover_image_url,json=coverImageUrl,proto3" json:"cover_image_url,omitempty"`
	ThumbnailImageUrl string `protobuf:"bytes,6,opt,name=thumbnail_image_url,json=thumbnailImageUrl,proto3" json:"thumbnail_image_url,omitempty"`
}

func (x *Serat) Reset() {
	*x = Serat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Serat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Serat) ProtoMessage() {}

func (x *Serat) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Serat.ProtoReflect.Descriptor instead.
func (*Serat) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{6}
}

func (x *Serat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Serat) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Serat) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Serat) GetCoverImageUrl() string {
	if x != nil {
		return x.CoverImageUrl
	}
	return ""
}

func (x *Serat) GetThumbnailImageUrl() string {
	if x != nil {
		return x.ThumbnailImageUrl
	}
	return ""
}

type ListSeratsResponse_MetaListSerats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *ResponsePagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListSeratsResponse_MetaListSerats) Reset() {
	*x = ListSeratsResponse_MetaListSerats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSeratsResponse_MetaListSerats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSeratsResponse_MetaListSerats) ProtoMessage() {}

func (x *ListSeratsResponse_MetaListSerats) ProtoReflect() protoreflect.Message {
	mi := &file_serats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSeratsResponse_MetaListSerats.ProtoReflect.Descriptor instead.
func (*ListSeratsResponse_MetaListSerats) Descriptor() ([]byte, []int) {
	return file_serats_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListSeratsResponse_MetaListSerats) GetPagination() *ResponsePagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_serats_proto protoreflect.FileDescriptor

var file_serats_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x1a, 0x17, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd7, 0x01, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x61, 0x74,
	0x73, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x61, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61,
	0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x61, 0x74, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x61, 0x74, 0x73, 0x1a, 0x51, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x61, 0x73, 0x75,
	0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xce, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0x92, 0x41, 0x0f, 0x32, 0x0d, 0x53, 0x65, 0x72, 0x61,
	0x74, 0x27, 0x73, 0x20, 0x74, 0x69, 0x74, 0x6c, 0x65, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x02,
	0x18, 0xff, 0x01, 0xd0, 0x01, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x47, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x25, 0x92, 0x41, 0x15, 0x32, 0x13, 0x53, 0x65, 0x72, 0x61, 0x74, 0x27, 0x73,
	0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0xfa, 0x42, 0x0a, 0x72,
	0x08, 0x10, 0x02, 0x18, 0xff, 0x01, 0xd0, 0x01, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2d, 0x92, 0x41, 0x1f, 0x32, 0x17, 0x53, 0x65, 0x72, 0x61, 0x74, 0x27, 0x73, 0x20, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55, 0x52, 0x4c, 0xa2, 0x02, 0x03,
	0x75, 0x72, 0x69, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x88, 0x01, 0x01, 0xd0, 0x01, 0x00, 0x52, 0x0d,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x61, 0x0a,
	0x13, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0x92, 0x41, 0x23, 0x32,
	0x1b, 0x53, 0x65, 0x72, 0x61, 0x74, 0x27, 0x73, 0x20, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55, 0x52, 0x4c, 0xa2, 0x02, 0x03, 0x75,
	0x72, 0x69, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x88, 0x01, 0x01, 0xd0, 0x01, 0x00, 0x52, 0x11, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x22, 0x84, 0x03, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x24, 0x92, 0x41, 0x16, 0x32, 0x0d, 0x53, 0x65, 0x72, 0x61, 0x74, 0x27,
	0x73, 0x20, 0x74, 0x69, 0x74, 0x6c, 0x65, 0xa2, 0x02, 0x04, 0x75, 0x75, 0x69, 0x64, 0xfa, 0x42,
	0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1f, 0x92, 0x41,
	0x0f, 0x32, 0x0d, 0x53, 0x65, 0x72, 0x61, 0x74, 0x27, 0x73, 0x20, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x02, 0x18, 0xff, 0x01, 0xd0, 0x01, 0x00, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0x92, 0x41, 0x15, 0x32, 0x13,
	0x53, 0x65, 0x72, 0x61, 0x74, 0x27, 0x73, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x02, 0x18, 0xff, 0x01, 0xd0, 0x01, 0x00,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a,
	0x0f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0x92, 0x41, 0x1f, 0x32, 0x17, 0x53, 0x65, 0x72,
	0x61, 0x74, 0x27, 0x73, 0x20, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x20, 0x55, 0x52, 0x4c, 0xa2, 0x02, 0x03, 0x75, 0x72, 0x69, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x88,
	0x01, 0x01, 0xd0, 0x01, 0x00, 0x52, 0x0d, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x61, 0x0a, 0x13, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x31, 0x92, 0x41, 0x23, 0x32, 0x1b, 0x53, 0x65, 0x72, 0x61, 0x74, 0x27, 0x73, 0x20,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20,
	0x55, 0x52, 0x4c, 0xa2, 0x02, 0x03, 0x75, 0x72, 0x69, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x88, 0x01,
	0x01, 0xd0, 0x01, 0x00, 0x52, 0x11, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xa7, 0x01, 0x0a, 0x05, 0x53, 0x65, 0x72, 0x61,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x32, 0x9f, 0x06, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x9e, 0x01, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x6b, 0x61,
	0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6b, 0x61,
	0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x61, 0x74, 0x73,
	0x92, 0x41, 0x3a, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x1a, 0x17, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x72, 0x61, 0x74,
	0x73, 0x2a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x8c, 0x01,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x61, 0x74, 0x12, 0x1c, 0x2e, 0x6b, 0x61, 0x73,
	0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73,
	0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x61, 0x74, 0x22, 0x4e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x61, 0x74, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73,
	0x12, 0x09, 0x47, 0x65, 0x74, 0x20, 0x53, 0x65, 0x72, 0x61, 0x74, 0x1a, 0x15, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x73, 0x65, 0x72, 0x61, 0x74, 0x20, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x61, 0x74, 0x12, 0xa2, 0x01, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x6b,
	0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41,
	0x40, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x20, 0x53, 0x65, 0x72, 0x61, 0x74, 0x1a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x73, 0x65, 0x72, 0x61, 0x74, 0x20, 0x62,
	0x79, 0x20, 0x69, 0x64, 0x2a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61,
	0x74, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61,
	0x74, 0x12, 0x1f, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e,
	0x2e, 0x53, 0x65, 0x72, 0x61, 0x74, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01,
	0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x61, 0x74, 0x73, 0x92, 0x41, 0x42,
	0x0a, 0x06, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x20, 0x53, 0x65, 0x72, 0x61, 0x74, 0x1a, 0x1d, 0x41, 0x64, 0x64, 0x20, 0x6e, 0x65, 0x77, 0x20,
	0x73, 0x65, 0x72, 0x61, 0x74, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x61, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x61, 0x74, 0x12, 0x1f, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61,
	0x6e, 0x2e, 0x53, 0x65, 0x72, 0x61, 0x74, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a,
	0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x61, 0x74, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x92, 0x41, 0x3d, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73, 0x12, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x53, 0x65, 0x72, 0x61, 0x74, 0x1a, 0x18, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x20, 0x73, 0x65, 0x72, 0x61, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x61, 0x74, 0x42, 0x9b, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6b, 0x72, 0x69, 0x72, 0x6e, 0x75, 0x72, 0x68, 0x69, 0x64, 0x61,
	0x79, 0x61, 0x74, 0x2f, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x92, 0x41, 0x6b, 0x12, 0x69, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x61, 0x74, 0x73,
	0x22, 0x5a, 0x0a, 0x17, 0x46, 0x69, 0x6b, 0x72, 0x69, 0x20, 0x52, 0x61, 0x68, 0x6d, 0x61, 0x74,
	0x20, 0x4e, 0x75, 0x72, 0x68, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x12, 0x23, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x69, 0x6b, 0x72, 0x69, 0x72, 0x6e, 0x75, 0x72, 0x68, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74,
	0x1a, 0x1a, 0x66, 0x69, 0x6b, 0x72, 0x69, 0x72, 0x6e, 0x75, 0x72, 0x68, 0x69, 0x64, 0x61, 0x79,
	0x61, 0x74, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serats_proto_rawDescOnce sync.Once
	file_serats_proto_rawDescData = file_serats_proto_rawDesc
)

func file_serats_proto_rawDescGZIP() []byte {
	file_serats_proto_rawDescOnce.Do(func() {
		file_serats_proto_rawDescData = protoimpl.X.CompressGZIP(file_serats_proto_rawDescData)
	})
	return file_serats_proto_rawDescData
}

var file_serats_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_serats_proto_goTypes = []interface{}{
	(*ListSeratsRequest)(nil),                 // 0: kasusastran.ListSeratsRequest
	(*ListSeratsResponse)(nil),                // 1: kasusastran.ListSeratsResponse
	(*GetSeratRequest)(nil),                   // 2: kasusastran.GetSeratRequest
	(*DeleteSeratRequest)(nil),                // 3: kasusastran.DeleteSeratRequest
	(*CreateSeratRequest)(nil),                // 4: kasusastran.CreateSeratRequest
	(*UpdateSeratRequest)(nil),                // 5: kasusastran.UpdateSeratRequest
	(*Serat)(nil),                             // 6: kasusastran.Serat
	(*ListSeratsResponse_MetaListSerats)(nil), // 7: kasusastran.ListSeratsResponse.MetaListSerats
	(*RequestPagination)(nil),                 // 8: kasusastran.RequestPagination
	(*ResponsePagination)(nil),                // 9: kasusastran.ResponsePagination
	(*emptypb.Empty)(nil),                     // 10: google.protobuf.Empty
}
var file_serats_proto_depIdxs = []int32{
	8,  // 0: kasusastran.ListSeratsRequest.pagination:type_name -> kasusastran.RequestPagination
	7,  // 1: kasusastran.ListSeratsResponse.meta:type_name -> kasusastran.ListSeratsResponse.MetaListSerats
	6,  // 2: kasusastran.ListSeratsResponse.serats:type_name -> kasusastran.Serat
	9,  // 3: kasusastran.ListSeratsResponse.MetaListSerats.pagination:type_name -> kasusastran.ResponsePagination
	0,  // 4: kasusastran.Serats.ListSerats:input_type -> kasusastran.ListSeratsRequest
	2,  // 5: kasusastran.Serats.GetSerat:input_type -> kasusastran.GetSeratRequest
	3,  // 6: kasusastran.Serats.DeleteSerat:input_type -> kasusastran.DeleteSeratRequest
	4,  // 7: kasusastran.Serats.CreateSerat:input_type -> kasusastran.CreateSeratRequest
	5,  // 8: kasusastran.Serats.UpdateSerat:input_type -> kasusastran.UpdateSeratRequest
	1,  // 9: kasusastran.Serats.ListSerats:output_type -> kasusastran.ListSeratsResponse
	6,  // 10: kasusastran.Serats.GetSerat:output_type -> kasusastran.Serat
	10, // 11: kasusastran.Serats.DeleteSerat:output_type -> google.protobuf.Empty
	6,  // 12: kasusastran.Serats.CreateSerat:output_type -> kasusastran.Serat
	6,  // 13: kasusastran.Serats.UpdateSerat:output_type -> kasusastran.Serat
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_serats_proto_init() }
func file_serats_proto_init() {
	if File_serats_proto != nil {
		return
	}
	file_common_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_serats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSeratsRequest); i {
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
		file_serats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSeratsResponse); i {
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
		file_serats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeratRequest); i {
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
		file_serats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSeratRequest); i {
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
		file_serats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSeratRequest); i {
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
		file_serats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSeratRequest); i {
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
		file_serats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Serat); i {
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
		file_serats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSeratsResponse_MetaListSerats); i {
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
			RawDescriptor: file_serats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serats_proto_goTypes,
		DependencyIndexes: file_serats_proto_depIdxs,
		MessageInfos:      file_serats_proto_msgTypes,
	}.Build()
	File_serats_proto = out.File
	file_serats_proto_rawDesc = nil
	file_serats_proto_goTypes = nil
	file_serats_proto_depIdxs = nil
}
