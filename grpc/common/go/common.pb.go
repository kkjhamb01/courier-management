// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: common.proto

package commonPb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type VehicleType int32

const (
	VehicleType_ANY        VehicleType = 0
	VehicleType_TRUCK      VehicleType = 1
	VehicleType_BICYCLE    VehicleType = 2
	VehicleType_CAR        VehicleType = 3
	VehicleType_SMALL_VAN  VehicleType = 4
	VehicleType_MEDIUM_VAN VehicleType = 5
	VehicleType_LARGE_VAN  VehicleType = 6
	VehicleType_MOTORBIKE  VehicleType = 7
)

// Enum value maps for VehicleType.
var (
	VehicleType_name = map[int32]string{
		0: "ANY",
		1: "TRUCK",
		2: "BICYCLE",
		3: "CAR",
		4: "SMALL_VAN",
		5: "MEDIUM_VAN",
		6: "LARGE_VAN",
		7: "MOTORBIKE",
	}
	VehicleType_value = map[string]int32{
		"ANY":        0,
		"TRUCK":      1,
		"BICYCLE":    2,
		"CAR":        3,
		"SMALL_VAN":  4,
		"MEDIUM_VAN": 5,
		"LARGE_VAN":  6,
		"MOTORBIKE":  7,
	}
)

func (x VehicleType) Enum() *VehicleType {
	p := new(VehicleType)
	*p = x
	return p
}

func (x VehicleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VehicleType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (VehicleType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x VehicleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VehicleType.Descriptor instead.
func (VehicleType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type UserType int32

const (
	UserType_COURIER  UserType = 0
	UserType_CUSTOMER UserType = 1
)

// Enum value maps for UserType.
var (
	UserType_name = map[int32]string{
		0: "COURIER",
		1: "CUSTOMER",
	}
	UserType_value = map[string]int32{
		"COURIER":  0,
		"CUSTOMER": 1,
	}
)

func (x UserType) Enum() *UserType {
	p := new(UserType)
	*p = x
	return p
}

func (x UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (UserType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserType.Descriptor instead.
func (UserType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type CourierStatus int32

const (
	CourierStatus_AVAILABLE   CourierStatus = 0
	CourierStatus_UNAVAILABLE CourierStatus = 1
	CourierStatus_BLOCKED     CourierStatus = 2
	CourierStatus_ENABLED     CourierStatus = 3
	CourierStatus_DISABLED    CourierStatus = 4
	CourierStatus_UNKNOWN     CourierStatus = 5
	CourierStatus_ON_RIDE     CourierStatus = 6
)

// Enum value maps for CourierStatus.
var (
	CourierStatus_name = map[int32]string{
		0: "AVAILABLE",
		1: "UNAVAILABLE",
		2: "BLOCKED",
		3: "ENABLED",
		4: "DISABLED",
		5: "UNKNOWN",
		6: "ON_RIDE",
	}
	CourierStatus_value = map[string]int32{
		"AVAILABLE":   0,
		"UNAVAILABLE": 1,
		"BLOCKED":     2,
		"ENABLED":     3,
		"DISABLED":    4,
		"UNKNOWN":     5,
		"ON_RIDE":     6,
	}
)

func (x CourierStatus) Enum() *CourierStatus {
	p := new(CourierStatus)
	*p = x
	return p
}

func (x CourierStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CourierStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (CourierStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x CourierStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CourierStatus.Descriptor instead.
func (CourierStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

type Courier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VehicleType VehicleType `protobuf:"varint,2,opt,name=vehicle_type,json=vehicleType,proto3,enum=artin.couriermanagement.common.VehicleType" json:"vehicle_type,omitempty"`
}

func (x *Courier) Reset() {
	*x = Courier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Courier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Courier) ProtoMessage() {}

func (x *Courier) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Courier.ProtoReflect.Descriptor instead.
func (*Courier) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Courier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Courier) GetVehicleType() VehicleType {
	if x != nil {
		return x.VehicleType
	}
	return VehicleType_ANY
}

type CourierLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courier  *Courier  `protobuf:"bytes,1,opt,name=courier,proto3" json:"courier,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CourierLocation) Reset() {
	*x = CourierLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierLocation) ProtoMessage() {}

func (x *CourierLocation) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierLocation.ProtoReflect.Descriptor instead.
func (*CourierLocation) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *CourierLocation) GetCourier() *Courier {
	if x != nil {
		return x.Courier
	}
	return nil
}

func (x *CourierLocation) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type CourierETA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courier  *Courier             `protobuf:"bytes,1,opt,name=courier,proto3" json:"courier,omitempty"`
	Duration *durationpb.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Meters   int32                `protobuf:"varint,3,opt,name=meters,proto3" json:"meters,omitempty"`
}

func (x *CourierETA) Reset() {
	*x = CourierETA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierETA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierETA) ProtoMessage() {}

func (x *CourierETA) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierETA.ProtoReflect.Descriptor instead.
func (*CourierETA) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *CourierETA) GetCourier() *Courier {
	if x != nil {
		return x.Courier
	}
	return nil
}

func (x *CourierETA) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *CourierETA) GetMeters() int32 {
	if x != nil {
		return x.Meters
	}
	return 0
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
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
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type LocationRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat   float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon   float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	Range float64 `protobuf:"fixed64,3,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *LocationRange) Reset() {
	*x = LocationRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationRange) ProtoMessage() {}

func (x *LocationRange) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationRange.ProtoReflect.Descriptor instead.
func (*LocationRange) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *LocationRange) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *LocationRange) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *LocationRange) GetRange() float64 {
	if x != nil {
		return x.Range
	}
	return 0
}

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *TimeRange) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimeRange) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4e, 0x0a,
	0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9a, 0x01,
	0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x43,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x45, 0x54, 0x41, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x60, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x80, 0x56, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x56, 0xc0, 0x52, 0x03, 0x6c,
	0x61, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42,
	0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x66, 0x40, 0x29,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x66, 0xc0, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22, 0x7b, 0x0a,
	0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x29,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14,
	0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x56, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x80, 0x56, 0xc0, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x6c, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x80, 0x66, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x66, 0xc0, 0x52,
	0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x67, 0x0a, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x02, 0x74, 0x6f, 0x2a, 0x74, 0x0a, 0x0b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54,
	0x52, 0x55, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x43, 0x59, 0x43, 0x4c,
	0x45, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x41, 0x52, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x5f, 0x56, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x45, 0x44, 0x49, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x4c,
	0x41, 0x52, 0x47, 0x45, 0x5f, 0x56, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f,
	0x54, 0x4f, 0x52, 0x42, 0x49, 0x4b, 0x45, 0x10, 0x07, 0x2a, 0x25, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x52, 0x49, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x01,
	0x2a, 0x71, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x5f, 0x52, 0x49, 0x44,
	0x45, 0x10, 0x06, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x2e, 0x61, 0x69, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x6f,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(VehicleType)(0),              // 0: artin.couriermanagement.common.VehicleType
	(UserType)(0),                 // 1: artin.couriermanagement.common.UserType
	(CourierStatus)(0),            // 2: artin.couriermanagement.common.CourierStatus
	(*Courier)(nil),               // 3: artin.couriermanagement.common.Courier
	(*CourierLocation)(nil),       // 4: artin.couriermanagement.common.CourierLocation
	(*CourierETA)(nil),            // 5: artin.couriermanagement.common.CourierETA
	(*Location)(nil),              // 6: artin.couriermanagement.common.Location
	(*LocationRange)(nil),         // 7: artin.couriermanagement.common.LocationRange
	(*TimeRange)(nil),             // 8: artin.couriermanagement.common.TimeRange
	(*durationpb.Duration)(nil),   // 9: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_common_proto_depIdxs = []int32{
	0,  // 0: artin.couriermanagement.common.Courier.vehicle_type:type_name -> artin.couriermanagement.common.VehicleType
	3,  // 1: artin.couriermanagement.common.CourierLocation.courier:type_name -> artin.couriermanagement.common.Courier
	6,  // 2: artin.couriermanagement.common.CourierLocation.location:type_name -> artin.couriermanagement.common.Location
	3,  // 3: artin.couriermanagement.common.CourierETA.courier:type_name -> artin.couriermanagement.common.Courier
	9,  // 4: artin.couriermanagement.common.CourierETA.duration:type_name -> google.protobuf.Duration
	10, // 5: artin.couriermanagement.common.TimeRange.from:type_name -> google.protobuf.Timestamp
	10, // 6: artin.couriermanagement.common.TimeRange.to:type_name -> google.protobuf.Timestamp
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Courier); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierLocation); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierETA); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationRange); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}