// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.3
// source: app/protobuf/message/promotion.proto

package message

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

type Active struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Active) Reset() {
	*x = Active{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protobuf_message_promotion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Active) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Active) ProtoMessage() {}

func (x *Active) ProtoReflect() protoreflect.Message {
	mi := &file_app_protobuf_message_promotion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Active.ProtoReflect.Descriptor instead.
func (*Active) Descriptor() ([]byte, []int) {
	return file_app_protobuf_message_promotion_proto_rawDescGZIP(), []int{0}
}

func (x *Active) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type PromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Image     string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Active    *Active `protobuf:"bytes,5,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt int64   `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64   `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt int64   `protobuf:"varint,8,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	Limit     int32   `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32   `protobuf:"varint,10,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PromotionRequest) Reset() {
	*x = PromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protobuf_message_promotion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionRequest) ProtoMessage() {}

func (x *PromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_protobuf_message_promotion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionRequest.ProtoReflect.Descriptor instead.
func (*PromotionRequest) Descriptor() ([]byte, []int) {
	return file_app_protobuf_message_promotion_proto_rawDescGZIP(), []int{1}
}

func (x *PromotionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PromotionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PromotionRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PromotionRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PromotionRequest) GetActive() *Active {
	if x != nil {
		return x.Active
	}
	return nil
}

func (x *PromotionRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PromotionRequest) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *PromotionRequest) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *PromotionRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PromotionRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type PromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                   string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	PromotionType          string  `protobuf:"bytes,3,opt,name=promotionType,proto3" json:"promotionType,omitempty"`
	Value                  float32 `protobuf:"fixed32,4,opt,name=value,proto3" json:"value,omitempty"`
	Image                  string  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Active                 *Active `protobuf:"bytes,6,opt,name=active,proto3" json:"active,omitempty"`
	CreatedBy              int32   `protobuf:"varint,7,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedAt              int64   `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt              int64   `protobuf:"varint,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt              int64   `protobuf:"varint,10,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	ActiveFrom             int64   `protobuf:"varint,11,opt,name=activeFrom,proto3" json:"activeFrom,omitempty"`
	ActiveTo               int64   `protobuf:"varint,12,opt,name=activeTo,proto3" json:"activeTo,omitempty"`
	DailyActiveFrom        int64   `protobuf:"varint,13,opt,name=dailyActiveFrom,proto3" json:"dailyActiveFrom,omitempty"`
	DailyActiveTo          int32   `protobuf:"varint,14,opt,name=dailyActiveTo,proto3" json:"dailyActiveTo,omitempty"`
	MaxActiveTime          int32   `protobuf:"varint,15,opt,name=maxActiveTime,proto3" json:"maxActiveTime,omitempty"`
	MaxDailyActiveTime     int32   `protobuf:"varint,16,opt,name=maxDailyActiveTime,proto3" json:"maxDailyActiveTime,omitempty"`
	PerUserActiveTime      int32   `protobuf:"varint,17,opt,name=perUserActiveTime,proto3" json:"perUserActiveTime,omitempty"`
	PerUserDailyActiveTime int32   `protobuf:"varint,18,opt,name=perUserDailyActiveTime,proto3" json:"perUserDailyActiveTime,omitempty"`
}

func (x *PromotionResponse) Reset() {
	*x = PromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protobuf_message_promotion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionResponse) ProtoMessage() {}

func (x *PromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_protobuf_message_promotion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionResponse.ProtoReflect.Descriptor instead.
func (*PromotionResponse) Descriptor() ([]byte, []int) {
	return file_app_protobuf_message_promotion_proto_rawDescGZIP(), []int{2}
}

func (x *PromotionResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PromotionResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PromotionResponse) GetPromotionType() string {
	if x != nil {
		return x.PromotionType
	}
	return ""
}

func (x *PromotionResponse) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PromotionResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PromotionResponse) GetActive() *Active {
	if x != nil {
		return x.Active
	}
	return nil
}

func (x *PromotionResponse) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *PromotionResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PromotionResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *PromotionResponse) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *PromotionResponse) GetActiveFrom() int64 {
	if x != nil {
		return x.ActiveFrom
	}
	return 0
}

func (x *PromotionResponse) GetActiveTo() int64 {
	if x != nil {
		return x.ActiveTo
	}
	return 0
}

func (x *PromotionResponse) GetDailyActiveFrom() int64 {
	if x != nil {
		return x.DailyActiveFrom
	}
	return 0
}

func (x *PromotionResponse) GetDailyActiveTo() int32 {
	if x != nil {
		return x.DailyActiveTo
	}
	return 0
}

func (x *PromotionResponse) GetMaxActiveTime() int32 {
	if x != nil {
		return x.MaxActiveTime
	}
	return 0
}

func (x *PromotionResponse) GetMaxDailyActiveTime() int32 {
	if x != nil {
		return x.MaxDailyActiveTime
	}
	return 0
}

func (x *PromotionResponse) GetPerUserActiveTime() int32 {
	if x != nil {
		return x.PerUserActiveTime
	}
	return 0
}

func (x *PromotionResponse) GetPerUserDailyActiveTime() int32 {
	if x != nil {
		return x.PerUserDailyActiveTime
	}
	return 0
}

type PromotionResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromotionResponses []*PromotionResponse `protobuf:"bytes,1,rep,name=promotionResponses,proto3" json:"promotionResponses,omitempty"`
}

func (x *PromotionResponses) Reset() {
	*x = PromotionResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_protobuf_message_promotion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionResponses) ProtoMessage() {}

func (x *PromotionResponses) ProtoReflect() protoreflect.Message {
	mi := &file_app_protobuf_message_promotion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionResponses.ProtoReflect.Descriptor instead.
func (*PromotionResponses) Descriptor() ([]byte, []int) {
	return file_app_protobuf_message_promotion_proto_rawDescGZIP(), []int{3}
}

func (x *PromotionResponses) GetPromotionResponses() []*PromotionResponse {
	if x != nil {
		return x.PromotionResponses
	}
	return nil
}

var File_app_protobuf_message_promotion_proto protoreflect.FileDescriptor

var file_app_protobuf_message_promotion_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x9d, 0x02, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x31, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xfc, 0x04, 0x0a, 0x11,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x12, 0x28,
	0x0a, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x69, 0x6c,
	0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x6f, 0x12, 0x24,
	0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x6d, 0x61, 0x78, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x11, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x16, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x12, 0x54, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_protobuf_message_promotion_proto_rawDescOnce sync.Once
	file_app_protobuf_message_promotion_proto_rawDescData = file_app_protobuf_message_promotion_proto_rawDesc
)

func file_app_protobuf_message_promotion_proto_rawDescGZIP() []byte {
	file_app_protobuf_message_promotion_proto_rawDescOnce.Do(func() {
		file_app_protobuf_message_promotion_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_protobuf_message_promotion_proto_rawDescData)
	})
	return file_app_protobuf_message_promotion_proto_rawDescData
}

var file_app_protobuf_message_promotion_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_protobuf_message_promotion_proto_goTypes = []interface{}{
	(*Active)(nil),             // 0: promotion.message.Active
	(*PromotionRequest)(nil),   // 1: promotion.message.PromotionRequest
	(*PromotionResponse)(nil),  // 2: promotion.message.PromotionResponse
	(*PromotionResponses)(nil), // 3: promotion.message.PromotionResponses
}
var file_app_protobuf_message_promotion_proto_depIdxs = []int32{
	0, // 0: promotion.message.PromotionRequest.active:type_name -> promotion.message.Active
	0, // 1: promotion.message.PromotionResponse.active:type_name -> promotion.message.Active
	2, // 2: promotion.message.PromotionResponses.promotionResponses:type_name -> promotion.message.PromotionResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_protobuf_message_promotion_proto_init() }
func file_app_protobuf_message_promotion_proto_init() {
	if File_app_protobuf_message_promotion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_protobuf_message_promotion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Active); i {
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
		file_app_protobuf_message_promotion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionRequest); i {
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
		file_app_protobuf_message_promotion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionResponse); i {
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
		file_app_protobuf_message_promotion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionResponses); i {
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
			RawDescriptor: file_app_protobuf_message_promotion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_protobuf_message_promotion_proto_goTypes,
		DependencyIndexes: file_app_protobuf_message_promotion_proto_depIdxs,
		MessageInfos:      file_app_protobuf_message_promotion_proto_msgTypes,
	}.Build()
	File_app_protobuf_message_promotion_proto = out.File
	file_app_protobuf_message_promotion_proto_rawDesc = nil
	file_app_protobuf_message_promotion_proto_goTypes = nil
	file_app_protobuf_message_promotion_proto_depIdxs = nil
}
