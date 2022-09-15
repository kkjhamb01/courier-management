// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: pricing_message.proto

package pricingPb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_go "gitlab.artin.ai/backend/courier-management/grpc/common/go"
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

type CalculateCourierPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleType     _go.VehicleType `protobuf:"varint,1,opt,name=vehicle_type,json=vehicleType,proto3,enum=artin.couriermanagement.common.VehicleType" json:"vehicle_type,omitempty"`
	RequiredWorkers int32           `protobuf:"varint,2,opt,name=required_workers,json=requiredWorkers,proto3" json:"required_workers,omitempty"`
	Source          *_go.Location   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destinations    []*_go.Location `protobuf:"bytes,4,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *CalculateCourierPriceRequest) Reset() {
	*x = CalculateCourierPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateCourierPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateCourierPriceRequest) ProtoMessage() {}

func (x *CalculateCourierPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateCourierPriceRequest.ProtoReflect.Descriptor instead.
func (*CalculateCourierPriceRequest) Descriptor() ([]byte, []int) {
	return file_pricing_message_proto_rawDescGZIP(), []int{0}
}

func (x *CalculateCourierPriceRequest) GetVehicleType() _go.VehicleType {
	if x != nil {
		return x.VehicleType
	}
	return _go.VehicleType_ANY
}

func (x *CalculateCourierPriceRequest) GetRequiredWorkers() int32 {
	if x != nil {
		return x.RequiredWorkers
	}
	return 0
}

func (x *CalculateCourierPriceRequest) GetSource() *_go.Location {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CalculateCourierPriceRequest) GetDestinations() []*_go.Location {
	if x != nil {
		return x.Destinations
	}
	return nil
}

type CalculateCourierPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EstimatedDuration int64   `protobuf:"varint,1,opt,name=estimated_duration,json=estimatedDuration,proto3" json:"estimated_duration,omitempty"`
	EstimatedDistance int32   `protobuf:"varint,2,opt,name=estimated_distance,json=estimatedDistance,proto3" json:"estimated_distance,omitempty"`
	Amount            float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency          string  `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *CalculateCourierPriceResponse) Reset() {
	*x = CalculateCourierPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateCourierPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateCourierPriceResponse) ProtoMessage() {}

func (x *CalculateCourierPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateCourierPriceResponse.ProtoReflect.Descriptor instead.
func (*CalculateCourierPriceResponse) Descriptor() ([]byte, []int) {
	return file_pricing_message_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateCourierPriceResponse) GetEstimatedDuration() int64 {
	if x != nil {
		return x.EstimatedDuration
	}
	return 0
}

func (x *CalculateCourierPriceResponse) GetEstimatedDistance() int32 {
	if x != nil {
		return x.EstimatedDistance
	}
	return 0
}

func (x *CalculateCourierPriceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CalculateCourierPriceResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type ReviewCourierPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequiredWorkers int32           `protobuf:"varint,1,opt,name=required_workers,json=requiredWorkers,proto3" json:"required_workers,omitempty"`
	Source          *_go.Location   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Destinations    []*_go.Location `protobuf:"bytes,3,rep,name=destinations,proto3" json:"destinations,omitempty"`
}

func (x *ReviewCourierPriceRequest) Reset() {
	*x = ReviewCourierPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewCourierPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewCourierPriceRequest) ProtoMessage() {}

func (x *ReviewCourierPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewCourierPriceRequest.ProtoReflect.Descriptor instead.
func (*ReviewCourierPriceRequest) Descriptor() ([]byte, []int) {
	return file_pricing_message_proto_rawDescGZIP(), []int{2}
}

func (x *ReviewCourierPriceRequest) GetRequiredWorkers() int32 {
	if x != nil {
		return x.RequiredWorkers
	}
	return 0
}

func (x *ReviewCourierPriceRequest) GetSource() *_go.Location {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *ReviewCourierPriceRequest) GetDestinations() []*_go.Location {
	if x != nil {
		return x.Destinations
	}
	return nil
}

type ReviewCourierPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EstimatedDuration int64                               `protobuf:"varint,1,opt,name=estimated_duration,json=estimatedDuration,proto3" json:"estimated_duration,omitempty"`
	EstimatedDistance int32                               `protobuf:"varint,2,opt,name=estimated_distance,json=estimatedDistance,proto3" json:"estimated_distance,omitempty"`
	Prices            []*ReviewCourierPriceResponse_Price `protobuf:"bytes,3,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (x *ReviewCourierPriceResponse) Reset() {
	*x = ReviewCourierPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewCourierPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewCourierPriceResponse) ProtoMessage() {}

func (x *ReviewCourierPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewCourierPriceResponse.ProtoReflect.Descriptor instead.
func (*ReviewCourierPriceResponse) Descriptor() ([]byte, []int) {
	return file_pricing_message_proto_rawDescGZIP(), []int{3}
}

func (x *ReviewCourierPriceResponse) GetEstimatedDuration() int64 {
	if x != nil {
		return x.EstimatedDuration
	}
	return 0
}

func (x *ReviewCourierPriceResponse) GetEstimatedDistance() int32 {
	if x != nil {
		return x.EstimatedDistance
	}
	return 0
}

func (x *ReviewCourierPriceResponse) GetPrices() []*ReviewCourierPriceResponse_Price {
	if x != nil {
		return x.Prices
	}
	return nil
}

type ReviewCourierPriceResponse_Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleType _go.VehicleType `protobuf:"varint,4,opt,name=vehicle_type,json=vehicleType,proto3,enum=artin.couriermanagement.common.VehicleType" json:"vehicle_type,omitempty"`
	Amount      float64         `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency    string          `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *ReviewCourierPriceResponse_Price) Reset() {
	*x = ReviewCourierPriceResponse_Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pricing_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewCourierPriceResponse_Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewCourierPriceResponse_Price) ProtoMessage() {}

func (x *ReviewCourierPriceResponse_Price) ProtoReflect() protoreflect.Message {
	mi := &file_pricing_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewCourierPriceResponse_Price.ProtoReflect.Descriptor instead.
func (*ReviewCourierPriceResponse_Price) Descriptor() ([]byte, []int) {
	return file_pricing_message_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ReviewCourierPriceResponse_Price) GetVehicleType() _go.VehicleType {
	if x != nil {
		return x.VehicleType
	}
	return _go.VehicleType_ANY
}

func (x *ReviewCourierPriceResponse_Price) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReviewCourierPriceResponse_Price) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

var File_pricing_message_proto protoreflect.FileDescriptor

var file_pricing_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbd, 0x02, 0x0a, 0x1c, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4e, 0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x4a, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xb1, 0x01, 0x0a, 0x1d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x12, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0xea, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x4a, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xe3, 0x02, 0x0a, 0x1a, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x12, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x12, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x59, 0x0a,
	0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x8b, 0x01, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x61, 0x69, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pricing_message_proto_rawDescOnce sync.Once
	file_pricing_message_proto_rawDescData = file_pricing_message_proto_rawDesc
)

func file_pricing_message_proto_rawDescGZIP() []byte {
	file_pricing_message_proto_rawDescOnce.Do(func() {
		file_pricing_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pricing_message_proto_rawDescData)
	})
	return file_pricing_message_proto_rawDescData
}

var file_pricing_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pricing_message_proto_goTypes = []interface{}{
	(*CalculateCourierPriceRequest)(nil),     // 0: artin.couriermanagement.pricing.CalculateCourierPriceRequest
	(*CalculateCourierPriceResponse)(nil),    // 1: artin.couriermanagement.pricing.CalculateCourierPriceResponse
	(*ReviewCourierPriceRequest)(nil),        // 2: artin.couriermanagement.pricing.ReviewCourierPriceRequest
	(*ReviewCourierPriceResponse)(nil),       // 3: artin.couriermanagement.pricing.ReviewCourierPriceResponse
	(*ReviewCourierPriceResponse_Price)(nil), // 4: artin.couriermanagement.pricing.ReviewCourierPriceResponse.Price
	(_go.VehicleType)(0),                     // 5: artin.couriermanagement.common.VehicleType
	(*_go.Location)(nil),                     // 6: artin.couriermanagement.common.Location
}
var file_pricing_message_proto_depIdxs = []int32{
	5, // 0: artin.couriermanagement.pricing.CalculateCourierPriceRequest.vehicle_type:type_name -> artin.couriermanagement.common.VehicleType
	6, // 1: artin.couriermanagement.pricing.CalculateCourierPriceRequest.source:type_name -> artin.couriermanagement.common.Location
	6, // 2: artin.couriermanagement.pricing.CalculateCourierPriceRequest.destinations:type_name -> artin.couriermanagement.common.Location
	6, // 3: artin.couriermanagement.pricing.ReviewCourierPriceRequest.source:type_name -> artin.couriermanagement.common.Location
	6, // 4: artin.couriermanagement.pricing.ReviewCourierPriceRequest.destinations:type_name -> artin.couriermanagement.common.Location
	4, // 5: artin.couriermanagement.pricing.ReviewCourierPriceResponse.prices:type_name -> artin.couriermanagement.pricing.ReviewCourierPriceResponse.Price
	5, // 6: artin.couriermanagement.pricing.ReviewCourierPriceResponse.Price.vehicle_type:type_name -> artin.couriermanagement.common.VehicleType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pricing_message_proto_init() }
func file_pricing_message_proto_init() {
	if File_pricing_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pricing_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateCourierPriceRequest); i {
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
		file_pricing_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateCourierPriceResponse); i {
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
		file_pricing_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewCourierPriceRequest); i {
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
		file_pricing_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewCourierPriceResponse); i {
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
		file_pricing_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewCourierPriceResponse_Price); i {
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
			RawDescriptor: file_pricing_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pricing_message_proto_goTypes,
		DependencyIndexes: file_pricing_message_proto_depIdxs,
		MessageInfos:      file_pricing_message_proto_msgTypes,
	}.Build()
	File_pricing_message_proto = out.File
	file_pricing_message_proto_rawDesc = nil
	file_pricing_message_proto_goTypes = nil
	file_pricing_message_proto_depIdxs = nil
}
