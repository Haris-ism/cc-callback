// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/merchant/trans.proto

package merchant

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

type ReqTransItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID     string `protobuf:"bytes,1,opt,name=ItemID,json=item_id,proto3" json:"ItemID,omitempty"`
	Discount   string `protobuf:"bytes,2,opt,name=Discount,json=discount,proto3" json:"Discount,omitempty"`
	Quantity   string `protobuf:"bytes,3,opt,name=Quantity,json=quantity,proto3" json:"Quantity,omitempty"`
	CCNumber   string `protobuf:"bytes,4,opt,name=CCNumber,json=cc_number,proto3" json:"CCNumber,omitempty"`
	CVV        string `protobuf:"bytes,5,opt,name=CVV,json=cvv,proto3" json:"CVV,omitempty"`
	Amount     string `protobuf:"bytes,6,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	Price      string `protobuf:"bytes,7,opt,name=Price,json=price,proto3" json:"Price,omitempty"`
	Name       string `protobuf:"bytes,8,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Type       string `protobuf:"bytes,9,opt,name=Type,json=type,proto3" json:"Type,omitempty"`
	Percentage string `protobuf:"bytes,10,opt,name=Percentage,json=percentage,proto3" json:"Percentage,omitempty"`
}

func (x *ReqTransItems) Reset() {
	*x = ReqTransItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_trans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqTransItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqTransItems) ProtoMessage() {}

func (x *ReqTransItems) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_trans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqTransItems.ProtoReflect.Descriptor instead.
func (*ReqTransItems) Descriptor() ([]byte, []int) {
	return file_proto_merchant_trans_proto_rawDescGZIP(), []int{0}
}

func (x *ReqTransItems) GetItemID() string {
	if x != nil {
		return x.ItemID
	}
	return ""
}

func (x *ReqTransItems) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

func (x *ReqTransItems) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *ReqTransItems) GetCCNumber() string {
	if x != nil {
		return x.CCNumber
	}
	return ""
}

func (x *ReqTransItems) GetCVV() string {
	if x != nil {
		return x.CVV
	}
	return ""
}

func (x *ReqTransItems) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ReqTransItems) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ReqTransItems) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqTransItems) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReqTransItems) GetPercentage() string {
	if x != nil {
		return x.Percentage
	}
	return ""
}

type ResTransItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,json=item_id,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,json=item_name,proto3" json:"Name,omitempty"`
	Quantity string `protobuf:"bytes,3,opt,name=Quantity,json=quantity,proto3" json:"Quantity,omitempty"`
	CC       string `protobuf:"bytes,4,opt,name=CC,json=cc_number,proto3" json:"CC,omitempty"`
}

func (x *ResTransItems) Reset() {
	*x = ResTransItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_trans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResTransItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResTransItems) ProtoMessage() {}

func (x *ResTransItems) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_trans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResTransItems.ProtoReflect.Descriptor instead.
func (*ResTransItems) Descriptor() ([]byte, []int) {
	return file_proto_merchant_trans_proto_rawDescGZIP(), []int{1}
}

func (x *ResTransItems) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ResTransItems) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResTransItems) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *ResTransItems) GetCC() string {
	if x != nil {
		return x.CC
	}
	return ""
}

type ReqCallbackItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID   string `protobuf:"bytes,1,opt,name=ItemID,json=item_id,proto3" json:"ItemID,omitempty"`
	Discount string `protobuf:"bytes,2,opt,name=Discount,json=discount,proto3" json:"Discount,omitempty"`
	Quantity string `protobuf:"bytes,3,opt,name=Quantity,json=quantity,proto3" json:"Quantity,omitempty"`
	CCNumber string `protobuf:"bytes,4,opt,name=CCNumber,json=cc_number,proto3" json:"CCNumber,omitempty"`
	Amount   string `protobuf:"bytes,5,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
}

func (x *ReqCallbackItems) Reset() {
	*x = ReqCallbackItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_trans_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqCallbackItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqCallbackItems) ProtoMessage() {}

func (x *ReqCallbackItems) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_trans_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqCallbackItems.ProtoReflect.Descriptor instead.
func (*ReqCallbackItems) Descriptor() ([]byte, []int) {
	return file_proto_merchant_trans_proto_rawDescGZIP(), []int{2}
}

func (x *ReqCallbackItems) GetItemID() string {
	if x != nil {
		return x.ItemID
	}
	return ""
}

func (x *ReqCallbackItems) GetDiscount() string {
	if x != nil {
		return x.Discount
	}
	return ""
}

func (x *ReqCallbackItems) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *ReqCallbackItems) GetCCNumber() string {
	if x != nil {
		return x.CCNumber
	}
	return ""
}

func (x *ReqCallbackItems) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type ResCallbackItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,json=item_id,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,json=item_name,proto3" json:"Name,omitempty"`
	Quantity string `protobuf:"bytes,3,opt,name=Quantity,json=quantity,proto3" json:"Quantity,omitempty"`
	CCNumber string `protobuf:"bytes,4,opt,name=CCNumber,json=cc_number,proto3" json:"CCNumber,omitempty"`
	Code     string `protobuf:"bytes,5,opt,name=Code,json=code,proto3" json:"Code,omitempty"`
}

func (x *ResCallbackItems) Reset() {
	*x = ResCallbackItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_trans_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResCallbackItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResCallbackItems) ProtoMessage() {}

func (x *ResCallbackItems) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_trans_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResCallbackItems.ProtoReflect.Descriptor instead.
func (*ResCallbackItems) Descriptor() ([]byte, []int) {
	return file_proto_merchant_trans_proto_rawDescGZIP(), []int{3}
}

func (x *ResCallbackItems) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ResCallbackItems) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResCallbackItems) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

func (x *ResCallbackItems) GetCCNumber() string {
	if x != nil {
		return x.CCNumber
	}
	return ""
}

func (x *ResCallbackItems) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ResMerchantCallbackModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string            `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
	Code    int32             `protobuf:"varint,2,opt,name=Code,json=code,proto3" json:"Code,omitempty"`
	Data    *ResCallbackItems `protobuf:"bytes,3,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (x *ResMerchantCallbackModel) Reset() {
	*x = ResMerchantCallbackModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_trans_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMerchantCallbackModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMerchantCallbackModel) ProtoMessage() {}

func (x *ResMerchantCallbackModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_trans_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMerchantCallbackModel.ProtoReflect.Descriptor instead.
func (*ResMerchantCallbackModel) Descriptor() ([]byte, []int) {
	return file_proto_merchant_trans_proto_rawDescGZIP(), []int{4}
}

func (x *ResMerchantCallbackModel) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResMerchantCallbackModel) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResMerchantCallbackModel) GetData() *ResCallbackItems {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResMerchantTransModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string         `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
	Code    int32          `protobuf:"varint,2,opt,name=Code,json=code,proto3" json:"Code,omitempty"`
	Data    *ResTransItems `protobuf:"bytes,3,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
}

func (x *ResMerchantTransModel) Reset() {
	*x = ResMerchantTransModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_trans_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResMerchantTransModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResMerchantTransModel) ProtoMessage() {}

func (x *ResMerchantTransModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_trans_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResMerchantTransModel.ProtoReflect.Descriptor instead.
func (*ResMerchantTransModel) Descriptor() ([]byte, []int) {
	return file_proto_merchant_trans_proto_rawDescGZIP(), []int{5}
}

func (x *ResMerchantTransModel) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResMerchantTransModel) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResMerchantTransModel) GetData() *ResTransItems {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_merchant_trans_proto protoreflect.FileDescriptor

var file_proto_merchant_trans_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x17,
	0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x1b, 0x0a, 0x08, 0x43, 0x43, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x43, 0x56, 0x56, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x76, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x22, 0x70, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x15, 0x0a, 0x02, 0x43, 0x43, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x63, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x49,
	0x74, 0x65, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x08,
	0x43, 0x43, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x1b, 0x0a, 0x08, 0x43, 0x43, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x6f, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x69, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x89, 0x01,
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x0e, 0x2e,
	0x52, 0x65, 0x71, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x16, 0x2e,
	0x52, 0x65, 0x73, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x42, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x11, 0x2e, 0x52, 0x65,
	0x71, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x19,
	0x2e, 0x52, 0x65, 0x73, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_merchant_trans_proto_rawDescOnce sync.Once
	file_proto_merchant_trans_proto_rawDescData = file_proto_merchant_trans_proto_rawDesc
)

func file_proto_merchant_trans_proto_rawDescGZIP() []byte {
	file_proto_merchant_trans_proto_rawDescOnce.Do(func() {
		file_proto_merchant_trans_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_merchant_trans_proto_rawDescData)
	})
	return file_proto_merchant_trans_proto_rawDescData
}

var file_proto_merchant_trans_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_merchant_trans_proto_goTypes = []interface{}{
	(*ReqTransItems)(nil),            // 0: ReqTransItems
	(*ResTransItems)(nil),            // 1: ResTransItems
	(*ReqCallbackItems)(nil),         // 2: ReqCallbackItems
	(*ResCallbackItems)(nil),         // 3: ResCallbackItems
	(*ResMerchantCallbackModel)(nil), // 4: ResMerchantCallbackModel
	(*ResMerchantTransModel)(nil),    // 5: ResMerchantTransModel
}
var file_proto_merchant_trans_proto_depIdxs = []int32{
	3, // 0: ResMerchantCallbackModel.Data:type_name -> ResCallbackItems
	1, // 1: ResMerchantTransModel.Data:type_name -> ResTransItems
	0, // 2: TransServices.TransItems:input_type -> ReqTransItems
	2, // 3: TransServices.CallbackTransItems:input_type -> ReqCallbackItems
	5, // 4: TransServices.TransItems:output_type -> ResMerchantTransModel
	4, // 5: TransServices.CallbackTransItems:output_type -> ResMerchantCallbackModel
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_merchant_trans_proto_init() }
func file_proto_merchant_trans_proto_init() {
	if File_proto_merchant_trans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_merchant_trans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqTransItems); i {
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
		file_proto_merchant_trans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResTransItems); i {
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
		file_proto_merchant_trans_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqCallbackItems); i {
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
		file_proto_merchant_trans_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResCallbackItems); i {
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
		file_proto_merchant_trans_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResMerchantCallbackModel); i {
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
		file_proto_merchant_trans_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResMerchantTransModel); i {
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
			RawDescriptor: file_proto_merchant_trans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_merchant_trans_proto_goTypes,
		DependencyIndexes: file_proto_merchant_trans_proto_depIdxs,
		MessageInfos:      file_proto_merchant_trans_proto_msgTypes,
	}.Build()
	File_proto_merchant_trans_proto = out.File
	file_proto_merchant_trans_proto_rawDesc = nil
	file_proto_merchant_trans_proto_goTypes = nil
	file_proto_merchant_trans_proto_depIdxs = nil
}
