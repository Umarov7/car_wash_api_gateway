// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: providers.proto

package providers

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

type NewProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CompanyName   string    `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Description   string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Services      []string  `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
	Availability  []string  `protobuf:"bytes,5,rep,name=availability,proto3" json:"availability,omitempty"`
	AverageRating float32   `protobuf:"fixed32,6,opt,name=average_rating,json=averageRating,proto3" json:"average_rating,omitempty"`
	Location      *Location `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *NewProvider) Reset() {
	*x = NewProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_providers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProvider) ProtoMessage() {}

func (x *NewProvider) ProtoReflect() protoreflect.Message {
	mi := &file_providers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProvider.ProtoReflect.Descriptor instead.
func (*NewProvider) Descriptor() ([]byte, []int) {
	return file_providers_proto_rawDescGZIP(), []int{0}
}

func (x *NewProvider) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NewProvider) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *NewProvider) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewProvider) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *NewProvider) GetAvailability() []string {
	if x != nil {
		return x.Availability
	}
	return nil
}

func (x *NewProvider) GetAverageRating() float32 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

func (x *NewProvider) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type CreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateResp) Reset() {
	*x = CreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_providers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResp) ProtoMessage() {}

func (x *CreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_providers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResp.ProtoReflect.Descriptor instead.
func (*CreateResp) Descriptor() ([]byte, []int) {
	return file_providers_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateResp) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyName   string  `protobuf:"bytes,1,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	AverageRating float32 `protobuf:"fixed32,2,opt,name=average_rating,json=averageRating,proto3" json:"average_rating,omitempty"`
	CreatedAt     string  `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_providers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_providers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_providers_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *Filter) GetAverageRating() float32 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

func (x *Filter) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	City      string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country   string  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Latitude  float32 `protobuf:"fixed32,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_providers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_providers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_providers_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Location) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Location) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Location) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string    `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CompanyName   string    `protobuf:"bytes,3,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Description   string    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Services      []string  `protobuf:"bytes,5,rep,name=services,proto3" json:"services,omitempty"`
	Availability  []string  `protobuf:"bytes,6,rep,name=availability,proto3" json:"availability,omitempty"`
	AverageRating float32   `protobuf:"fixed32,7,opt,name=average_rating,json=averageRating,proto3" json:"average_rating,omitempty"`
	Location      *Location `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt     string    `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string    `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_providers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_providers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_providers_proto_rawDescGZIP(), []int{4}
}

func (x *Provider) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Provider) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Provider) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *Provider) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Provider) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Provider) GetAvailability() []string {
	if x != nil {
		return x.Availability
	}
	return nil
}

func (x *Provider) GetAverageRating() float32 {
	if x != nil {
		return x.AverageRating
	}
	return 0
}

func (x *Provider) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Provider) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Provider) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SearchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Providers []*Provider `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *SearchResp) Reset() {
	*x = SearchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_providers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResp) ProtoMessage() {}

func (x *SearchResp) ProtoReflect() protoreflect.Message {
	mi := &file_providers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResp.ProtoReflect.Descriptor instead.
func (*SearchResp) Descriptor() ([]byte, []int) {
	return file_providers_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResp) GetProviders() []*Provider {
	if x != nil {
		return x.Providers
	}
	return nil
}

var File_providers_proto protoreflect.FileDescriptor

var file_providers_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x83, 0x02, 0x0a,
	0x0b, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x2f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x71, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x22, 0xce, 0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x3f, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x31, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x32, 0x89, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x4e, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_providers_proto_rawDescOnce sync.Once
	file_providers_proto_rawDescData = file_providers_proto_rawDesc
)

func file_providers_proto_rawDescGZIP() []byte {
	file_providers_proto_rawDescOnce.Do(func() {
		file_providers_proto_rawDescData = protoimpl.X.CompressGZIP(file_providers_proto_rawDescData)
	})
	return file_providers_proto_rawDescData
}

var file_providers_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_providers_proto_goTypes = []interface{}{
	(*NewProvider)(nil), // 0: providers.NewProvider
	(*CreateResp)(nil),  // 1: providers.CreateResp
	(*Filter)(nil),      // 2: providers.Filter
	(*Location)(nil),    // 3: providers.Location
	(*Provider)(nil),    // 4: providers.Provider
	(*SearchResp)(nil),  // 5: providers.SearchResp
}
var file_providers_proto_depIdxs = []int32{
	3, // 0: providers.NewProvider.location:type_name -> providers.Location
	3, // 1: providers.Provider.location:type_name -> providers.Location
	4, // 2: providers.SearchResp.providers:type_name -> providers.Provider
	0, // 3: providers.Providers.CreateProvider:input_type -> providers.NewProvider
	2, // 4: providers.Providers.SearchProviders:input_type -> providers.Filter
	1, // 5: providers.Providers.CreateProvider:output_type -> providers.CreateResp
	5, // 6: providers.Providers.SearchProviders:output_type -> providers.SearchResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_providers_proto_init() }
func file_providers_proto_init() {
	if File_providers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_providers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProvider); i {
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
		file_providers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResp); i {
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
		file_providers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_providers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_providers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_providers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResp); i {
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
			RawDescriptor: file_providers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_providers_proto_goTypes,
		DependencyIndexes: file_providers_proto_depIdxs,
		MessageInfos:      file_providers_proto_msgTypes,
	}.Build()
	File_providers_proto = out.File
	file_providers_proto_rawDesc = nil
	file_providers_proto_goTypes = nil
	file_providers_proto_depIdxs = nil
}
