// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: product_service.proto

package review

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

type ProductDetails struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Asin                 string                 `protobuf:"bytes,1,opt,name=asin,proto3" json:"asin,omitempty"`
	ProductTitle         string                 `protobuf:"bytes,2,opt,name=product_title,json=productTitle,proto3" json:"product_title,omitempty"`
	ProductPrice         string                 `protobuf:"bytes,3,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	ProductOriginalPrice string                 `protobuf:"bytes,4,opt,name=product_original_price,json=productOriginalPrice,proto3" json:"product_original_price,omitempty"`
	Currency             string                 `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Country              string                 `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	ProductUrl           string                 `protobuf:"bytes,7,opt,name=product_url,json=productUrl,proto3" json:"product_url,omitempty"`
	ProductPhoto         string                 `protobuf:"bytes,8,opt,name=product_photo,json=productPhoto,proto3" json:"product_photo,omitempty"`
	ProductAvailability  string                 `protobuf:"bytes,9,opt,name=product_availability,json=productAvailability,proto3" json:"product_availability,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ProductDetails) Reset() {
	*x = ProductDetails{}
	mi := &file_product_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetails) ProtoMessage() {}

func (x *ProductDetails) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetails.ProtoReflect.Descriptor instead.
func (*ProductDetails) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProductDetails) GetAsin() string {
	if x != nil {
		return x.Asin
	}
	return ""
}

func (x *ProductDetails) GetProductTitle() string {
	if x != nil {
		return x.ProductTitle
	}
	return ""
}

func (x *ProductDetails) GetProductPrice() string {
	if x != nil {
		return x.ProductPrice
	}
	return ""
}

func (x *ProductDetails) GetProductOriginalPrice() string {
	if x != nil {
		return x.ProductOriginalPrice
	}
	return ""
}

func (x *ProductDetails) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *ProductDetails) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ProductDetails) GetProductUrl() string {
	if x != nil {
		return x.ProductUrl
	}
	return ""
}

func (x *ProductDetails) GetProductPhoto() string {
	if x != nil {
		return x.ProductPhoto
	}
	return ""
}

func (x *ProductDetails) GetProductAvailability() string {
	if x != nil {
		return x.ProductAvailability
	}
	return ""
}

type GetProductDetailsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Asin          string                 `protobuf:"bytes,1,opt,name=asin,proto3" json:"asin,omitempty"`
	Country       string                 `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductDetailsRequest) Reset() {
	*x = GetProductDetailsRequest{}
	mi := &file_product_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductDetailsRequest) ProtoMessage() {}

func (x *GetProductDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetProductDetailsRequest) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductDetailsRequest) GetAsin() string {
	if x != nil {
		return x.Asin
	}
	return ""
}

func (x *GetProductDetailsRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type GetProductDetailsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *ProductDetails        `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductDetailsResponse) Reset() {
	*x = GetProductDetailsResponse{}
	mi := &file_product_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductDetailsResponse) ProtoMessage() {}

func (x *GetProductDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetProductDetailsResponse) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductDetailsResponse) GetProduct() *ProductDetails {
	if x != nil {
		return x.Product
	}
	return nil
}

type Products struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Asin                 string                 `protobuf:"bytes,1,opt,name=asin,proto3" json:"asin,omitempty"`
	ProductTitle         string                 `protobuf:"bytes,2,opt,name=product_title,json=productTitle,proto3" json:"product_title,omitempty"`
	ProductPrice         string                 `protobuf:"bytes,3,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	ProductOriginalPrice string                 `protobuf:"bytes,4,opt,name=product_original_price,json=productOriginalPrice,proto3" json:"product_original_price,omitempty"`
	Currency             string                 `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	ProductUrl           string                 `protobuf:"bytes,6,opt,name=product_url,json=productUrl,proto3" json:"product_url,omitempty"`
	ProductPhoto         string                 `protobuf:"bytes,7,opt,name=product_photo,json=productPhoto,proto3" json:"product_photo,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Products) Reset() {
	*x = Products{}
	mi := &file_product_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Products) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Products) ProtoMessage() {}

func (x *Products) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Products.ProtoReflect.Descriptor instead.
func (*Products) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{3}
}

func (x *Products) GetAsin() string {
	if x != nil {
		return x.Asin
	}
	return ""
}

func (x *Products) GetProductTitle() string {
	if x != nil {
		return x.ProductTitle
	}
	return ""
}

func (x *Products) GetProductPrice() string {
	if x != nil {
		return x.ProductPrice
	}
	return ""
}

func (x *Products) GetProductOriginalPrice() string {
	if x != nil {
		return x.ProductOriginalPrice
	}
	return ""
}

func (x *Products) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Products) GetProductUrl() string {
	if x != nil {
		return x.ProductUrl
	}
	return ""
}

func (x *Products) GetProductPhoto() string {
	if x != nil {
		return x.ProductPhoto
	}
	return ""
}

type GetProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          string                 `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Country       string                 `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsRequest) Reset() {
	*x = GetProductsRequest{}
	mi := &file_product_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsRequest) ProtoMessage() {}

func (x *GetProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsRequest.ProtoReflect.Descriptor instead.
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductsRequest) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *GetProductsRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type GetProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       []*Products            `protobuf:"bytes,1,rep,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsResponse) Reset() {
	*x = GetProductsResponse{}
	mi := &file_product_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsResponse) ProtoMessage() {}

func (x *GetProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductsResponse) GetProduct() []*Products {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_product_service_proto protoreflect.FileDescriptor

var file_product_service_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22,
	0xd3, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x73, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x73, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x48, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x73, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x73, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x4d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x80,
	0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x73, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x73, 0x69, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x41, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0xaf, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0x1a, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_product_service_proto_rawDescOnce sync.Once
	file_product_service_proto_rawDescData []byte
)

func file_product_service_proto_rawDescGZIP() []byte {
	file_product_service_proto_rawDescOnce.Do(func() {
		file_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_service_proto_rawDesc), len(file_product_service_proto_rawDesc)))
	})
	return file_product_service_proto_rawDescData
}

var file_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_service_proto_goTypes = []any{
	(*ProductDetails)(nil),            // 0: review.ProductDetails
	(*GetProductDetailsRequest)(nil),  // 1: review.GetProductDetailsRequest
	(*GetProductDetailsResponse)(nil), // 2: review.GetProductDetailsResponse
	(*Products)(nil),                  // 3: review.Products
	(*GetProductsRequest)(nil),        // 4: review.GetProductsRequest
	(*GetProductsResponse)(nil),       // 5: review.GetProductsResponse
}
var file_product_service_proto_depIdxs = []int32{
	0, // 0: review.GetProductDetailsResponse.product:type_name -> review.ProductDetails
	3, // 1: review.GetProductsResponse.product:type_name -> review.Products
	1, // 2: review.Product.GetProductDetails:input_type -> review.GetProductDetailsRequest
	4, // 3: review.Product.GetProducts:input_type -> review.GetProductsRequest
	2, // 4: review.Product.GetProductDetails:output_type -> review.GetProductDetailsResponse
	5, // 5: review.Product.GetProducts:output_type -> review.GetProductsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_product_service_proto_init() }
func file_product_service_proto_init() {
	if File_product_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_service_proto_rawDesc), len(file_product_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_service_proto_goTypes,
		DependencyIndexes: file_product_service_proto_depIdxs,
		MessageInfos:      file_product_service_proto_msgTypes,
	}.Build()
	File_product_service_proto = out.File
	file_product_service_proto_goTypes = nil
	file_product_service_proto_depIdxs = nil
}
