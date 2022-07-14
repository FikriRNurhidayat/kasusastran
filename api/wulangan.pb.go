// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: wulangan.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
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

type ListWulangansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *RequestPagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListWulangansRequest) Reset() {
	*x = ListWulangansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWulangansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWulangansRequest) ProtoMessage() {}

func (x *ListWulangansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWulangansRequest.ProtoReflect.Descriptor instead.
func (*ListWulangansRequest) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{0}
}

func (x *ListWulangansRequest) GetPagination() *RequestPagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListWulangansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *ListWulangansResponse_MetaListWulangans `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Wulangans []*Wulangan                              `protobuf:"bytes,2,rep,name=wulangans,proto3" json:"wulangans,omitempty"`
}

func (x *ListWulangansResponse) Reset() {
	*x = ListWulangansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWulangansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWulangansResponse) ProtoMessage() {}

func (x *ListWulangansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWulangansResponse.ProtoReflect.Descriptor instead.
func (*ListWulangansResponse) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{1}
}

func (x *ListWulangansResponse) GetMeta() *ListWulangansResponse_MetaListWulangans {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListWulangansResponse) GetWulangans() []*Wulangan {
	if x != nil {
		return x.Wulangans
	}
	return nil
}

type GetWulanganRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWulanganRequest) Reset() {
	*x = GetWulanganRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWulanganRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWulanganRequest) ProtoMessage() {}

func (x *GetWulanganRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWulanganRequest.ProtoReflect.Descriptor instead.
func (*GetWulanganRequest) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{2}
}

func (x *GetWulanganRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteWulanganRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWulanganRequest) Reset() {
	*x = DeleteWulanganRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWulanganRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWulanganRequest) ProtoMessage() {}

func (x *DeleteWulanganRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWulanganRequest.ProtoReflect.Descriptor instead.
func (*DeleteWulanganRequest) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteWulanganRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateWulanganRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title             string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description       string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CoverImageUrl     string `protobuf:"bytes,3,opt,name=cover_image_url,json=coverImageUrl,proto3" json:"cover_image_url,omitempty"`
	ThumbnailImageUrl string `protobuf:"bytes,4,opt,name=thumbnail_image_url,json=thumbnailImageUrl,proto3" json:"thumbnail_image_url,omitempty"`
}

func (x *CreateWulanganRequest) Reset() {
	*x = CreateWulanganRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWulanganRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWulanganRequest) ProtoMessage() {}

func (x *CreateWulanganRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWulanganRequest.ProtoReflect.Descriptor instead.
func (*CreateWulanganRequest) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{4}
}

func (x *CreateWulanganRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateWulanganRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateWulanganRequest) GetCoverImageUrl() string {
	if x != nil {
		return x.CoverImageUrl
	}
	return ""
}

func (x *CreateWulanganRequest) GetThumbnailImageUrl() string {
	if x != nil {
		return x.ThumbnailImageUrl
	}
	return ""
}

type UpdateWulanganRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description       string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CoverImageUrl     string `protobuf:"bytes,4,opt,name=cover_image_url,json=coverImageUrl,proto3" json:"cover_image_url,omitempty"`
	ThumbnailImageUrl string `protobuf:"bytes,5,opt,name=thumbnail_image_url,json=thumbnailImageUrl,proto3" json:"thumbnail_image_url,omitempty"`
}

func (x *UpdateWulanganRequest) Reset() {
	*x = UpdateWulanganRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWulanganRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWulanganRequest) ProtoMessage() {}

func (x *UpdateWulanganRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWulanganRequest.ProtoReflect.Descriptor instead.
func (*UpdateWulanganRequest) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateWulanganRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateWulanganRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateWulanganRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateWulanganRequest) GetCoverImageUrl() string {
	if x != nil {
		return x.CoverImageUrl
	}
	return ""
}

func (x *UpdateWulanganRequest) GetThumbnailImageUrl() string {
	if x != nil {
		return x.ThumbnailImageUrl
	}
	return ""
}

type Wulangan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title             string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description       string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CoverImageUrl     string                 `protobuf:"bytes,4,opt,name=cover_image_url,json=coverImageUrl,proto3" json:"cover_image_url,omitempty"`
	ThumbnailImageUrl string                 `protobuf:"bytes,5,opt,name=thumbnail_image_url,json=thumbnailImageUrl,proto3" json:"thumbnail_image_url,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt         *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Wulangan) Reset() {
	*x = Wulangan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wulangan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wulangan) ProtoMessage() {}

func (x *Wulangan) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wulangan.ProtoReflect.Descriptor instead.
func (*Wulangan) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{6}
}

func (x *Wulangan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Wulangan) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Wulangan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Wulangan) GetCoverImageUrl() string {
	if x != nil {
		return x.CoverImageUrl
	}
	return ""
}

func (x *Wulangan) GetThumbnailImageUrl() string {
	if x != nil {
		return x.ThumbnailImageUrl
	}
	return ""
}

func (x *Wulangan) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Wulangan) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ListWulangansResponse_MetaListWulangans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *ResponsePagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListWulangansResponse_MetaListWulangans) Reset() {
	*x = ListWulangansResponse_MetaListWulangans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wulangan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWulangansResponse_MetaListWulangans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWulangansResponse_MetaListWulangans) ProtoMessage() {}

func (x *ListWulangansResponse_MetaListWulangans) ProtoReflect() protoreflect.Message {
	mi := &file_wulangan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWulangansResponse_MetaListWulangans.ProtoReflect.Descriptor instead.
func (*ListWulangansResponse_MetaListWulangans) Descriptor() ([]byte, []int) {
	return file_wulangan_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListWulangansResponse_MetaListWulangans) GetPagination() *ResponsePagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_wulangan_proto protoreflect.FileDescriptor

var file_wulangan_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x1a, 0x17, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74,
	0x72, 0x61, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xec, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67,
	0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6b, 0x61, 0x73, 0x75,
	0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x75, 0x6c, 0x61,
	0x6e, 0x67, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x09, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73,
	0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x52,
	0x09, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x1a, 0x54, 0x0a, 0x11, 0x4d, 0x65,
	0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12,
	0x3f, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x00, 0xb0, 0x01, 0x01, 0x92,
	0x41, 0x0f, 0x32, 0x0d, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x27, 0x73, 0x20, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57,
	0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x08, 0x72,
	0x06, 0xd0, 0x01, 0x00, 0xb0, 0x01, 0x01, 0x92, 0x41, 0x0f, 0x32, 0x0d, 0x57, 0x75, 0x6c, 0x61,
	0x6e, 0x67, 0x61, 0x6e, 0x27, 0x73, 0x20, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x02,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x02, 0x18,
	0xff, 0x01, 0xd0, 0x01, 0x00, 0x92, 0x41, 0x12, 0x32, 0x10, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67,
	0x61, 0x6e, 0x27, 0x73, 0x20, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x02, 0x18,
	0xff, 0x01, 0xd0, 0x01, 0x00, 0x92, 0x41, 0x18, 0x32, 0x16, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67,
	0x61, 0x6e, 0x27, 0x73, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a,
	0x0f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x00,
	0x88, 0x01, 0x01, 0x92, 0x41, 0x22, 0x32, 0x1a, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x27, 0x73, 0x20, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55,
	0x52, 0x4c, 0xa2, 0x02, 0x03, 0x75, 0x72, 0x69, 0x52, 0x0d, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x64, 0x0a, 0x13, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x00, 0x88, 0x01,
	0x01, 0x92, 0x41, 0x26, 0x32, 0x1e, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x27, 0x73,
	0x20, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x20, 0x55, 0x52, 0x4c, 0xa2, 0x02, 0x03, 0x75, 0x72, 0x69, 0x52, 0x11, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x8c, 0x03,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x00, 0xb0, 0x01, 0x01,
	0x92, 0x41, 0x0f, 0x32, 0x0d, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x27, 0x73, 0x20,
	0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x02, 0x18, 0xff,
	0x01, 0xd0, 0x01, 0x00, 0x92, 0x41, 0x12, 0x32, 0x10, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61,
	0x6e, 0x27, 0x73, 0x20, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x4a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10, 0x02, 0x18, 0xff,
	0x01, 0xd0, 0x01, 0x00, 0x92, 0x41, 0x18, 0x32, 0x16, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61,
	0x6e, 0x27, 0x73, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x0f,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x00, 0x88,
	0x01, 0x01, 0x92, 0x41, 0x22, 0x32, 0x1a, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x27,
	0x73, 0x20, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x55, 0x52,
	0x4c, 0xa2, 0x02, 0x03, 0x75, 0x72, 0x69, 0x52, 0x0d, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x64, 0x0a, 0x13, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x34, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xd0, 0x01, 0x00, 0x88, 0x01, 0x01,
	0x92, 0x41, 0x26, 0x32, 0x1e, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x27, 0x73, 0x20,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20,
	0x55, 0x52, 0x4c, 0xa2, 0x02, 0x03, 0x75, 0x72, 0x69, 0x52, 0x11, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xa0, 0x02, 0x0a,
	0x08, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32,
	0x97, 0x07, 0x0a, 0x09, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12, 0xb6, 0x01,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12,
	0x21, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x92, 0x41, 0x46,
	0x0a, 0x09, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x1a, 0x1a, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x77, 0x75,
	0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x2a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x75, 0x6c,
	0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12, 0xa4, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x75,
	0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x12, 0x1f, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73,
	0x74, 0x72, 0x61, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61,
	0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x22, 0x5d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x75, 0x6c, 0x61,
	0x6e, 0x67, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x40, 0x0a, 0x09, 0x57,
	0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12, 0x0c, 0x47, 0x65, 0x74, 0x20, 0x57, 0x75,
	0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x1a, 0x18, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x20, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x12, 0xb7, 0x01,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x12, 0x22, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x69, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67,
	0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x4c, 0x0a, 0x09, 0x57, 0x75, 0x6c,
	0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x57,
	0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x1a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61,
	0x6e, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x2a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57,
	0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x12, 0xb6, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x12, 0x22, 0x2e, 0x6b, 0x61, 0x73,
	0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x57, 0x75, 0x6c,
	0x61, 0x6e, 0x67, 0x61, 0x6e, 0x22, 0x69, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a,
	0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x92,
	0x41, 0x4e, 0x0a, 0x09, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x1a, 0x20,
	0x41, 0x64, 0x64, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x12, 0xb6, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e,
	0x67, 0x61, 0x6e, 0x12, 0x22, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61, 0x73, 0x74, 0x72, 0x61,
	0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61,
	0x73, 0x74, 0x72, 0x61, 0x6e, 0x2e, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x22, 0x69,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x1a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0x49,
	0x0a, 0x09, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x73, 0x12, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x1a, 0x1b, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x20, 0x77, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x20, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x42, 0x9d, 0x01, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6b, 0x72, 0x69, 0x72, 0x6e,
	0x75, 0x72, 0x68, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x2f, 0x6b, 0x61, 0x73, 0x75, 0x73, 0x61,
	0x73, 0x74, 0x72, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x92, 0x41, 0x6d, 0x12, 0x6b, 0x0a, 0x08,
	0x57, 0x75, 0x6c, 0x61, 0x6e, 0x67, 0x61, 0x6e, 0x22, 0x5a, 0x0a, 0x17, 0x46, 0x69, 0x6b, 0x72,
	0x69, 0x20, 0x52, 0x61, 0x68, 0x6d, 0x61, 0x74, 0x20, 0x4e, 0x75, 0x72, 0x68, 0x69, 0x64, 0x61,
	0x79, 0x61, 0x74, 0x12, 0x23, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6b, 0x72, 0x69, 0x72, 0x6e, 0x75,
	0x72, 0x68, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x1a, 0x1a, 0x66, 0x69, 0x6b, 0x72, 0x69, 0x72,
	0x6e, 0x75, 0x72, 0x68, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wulangan_proto_rawDescOnce sync.Once
	file_wulangan_proto_rawDescData = file_wulangan_proto_rawDesc
)

func file_wulangan_proto_rawDescGZIP() []byte {
	file_wulangan_proto_rawDescOnce.Do(func() {
		file_wulangan_proto_rawDescData = protoimpl.X.CompressGZIP(file_wulangan_proto_rawDescData)
	})
	return file_wulangan_proto_rawDescData
}

var file_wulangan_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_wulangan_proto_goTypes = []interface{}{
	(*ListWulangansRequest)(nil),                    // 0: kasusastran.ListWulangansRequest
	(*ListWulangansResponse)(nil),                   // 1: kasusastran.ListWulangansResponse
	(*GetWulanganRequest)(nil),                      // 2: kasusastran.GetWulanganRequest
	(*DeleteWulanganRequest)(nil),                   // 3: kasusastran.DeleteWulanganRequest
	(*CreateWulanganRequest)(nil),                   // 4: kasusastran.CreateWulanganRequest
	(*UpdateWulanganRequest)(nil),                   // 5: kasusastran.UpdateWulanganRequest
	(*Wulangan)(nil),                                // 6: kasusastran.Wulangan
	(*ListWulangansResponse_MetaListWulangans)(nil), // 7: kasusastran.ListWulangansResponse.MetaListWulangans
	(*RequestPagination)(nil),                       // 8: kasusastran.RequestPagination
	(*timestamppb.Timestamp)(nil),                   // 9: google.protobuf.Timestamp
	(*ResponsePagination)(nil),                      // 10: kasusastran.ResponsePagination
	(*emptypb.Empty)(nil),                           // 11: google.protobuf.Empty
}
var file_wulangan_proto_depIdxs = []int32{
	8,  // 0: kasusastran.ListWulangansRequest.pagination:type_name -> kasusastran.RequestPagination
	7,  // 1: kasusastran.ListWulangansResponse.meta:type_name -> kasusastran.ListWulangansResponse.MetaListWulangans
	6,  // 2: kasusastran.ListWulangansResponse.wulangans:type_name -> kasusastran.Wulangan
	9,  // 3: kasusastran.Wulangan.created_at:type_name -> google.protobuf.Timestamp
	9,  // 4: kasusastran.Wulangan.updated_at:type_name -> google.protobuf.Timestamp
	10, // 5: kasusastran.ListWulangansResponse.MetaListWulangans.pagination:type_name -> kasusastran.ResponsePagination
	0,  // 6: kasusastran.Wulangans.ListWulangans:input_type -> kasusastran.ListWulangansRequest
	2,  // 7: kasusastran.Wulangans.GetWulangan:input_type -> kasusastran.GetWulanganRequest
	3,  // 8: kasusastran.Wulangans.DeleteWulangan:input_type -> kasusastran.DeleteWulanganRequest
	4,  // 9: kasusastran.Wulangans.CreateWulangan:input_type -> kasusastran.CreateWulanganRequest
	5,  // 10: kasusastran.Wulangans.UpdateWulangan:input_type -> kasusastran.UpdateWulanganRequest
	1,  // 11: kasusastran.Wulangans.ListWulangans:output_type -> kasusastran.ListWulangansResponse
	6,  // 12: kasusastran.Wulangans.GetWulangan:output_type -> kasusastran.Wulangan
	11, // 13: kasusastran.Wulangans.DeleteWulangan:output_type -> google.protobuf.Empty
	6,  // 14: kasusastran.Wulangans.CreateWulangan:output_type -> kasusastran.Wulangan
	6,  // 15: kasusastran.Wulangans.UpdateWulangan:output_type -> kasusastran.Wulangan
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_wulangan_proto_init() }
func file_wulangan_proto_init() {
	if File_wulangan_proto != nil {
		return
	}
	file_common_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wulangan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWulangansRequest); i {
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
		file_wulangan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWulangansResponse); i {
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
		file_wulangan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWulanganRequest); i {
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
		file_wulangan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWulanganRequest); i {
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
		file_wulangan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWulanganRequest); i {
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
		file_wulangan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWulanganRequest); i {
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
		file_wulangan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wulangan); i {
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
		file_wulangan_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWulangansResponse_MetaListWulangans); i {
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
			RawDescriptor: file_wulangan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wulangan_proto_goTypes,
		DependencyIndexes: file_wulangan_proto_depIdxs,
		MessageInfos:      file_wulangan_proto_msgTypes,
	}.Build()
	File_wulangan_proto = out.File
	file_wulangan_proto_rawDesc = nil
	file_wulangan_proto_goTypes = nil
	file_wulangan_proto_depIdxs = nil
}
