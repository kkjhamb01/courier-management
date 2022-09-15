// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: notification.proto

package proto

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

type DeviceOS int32

const (
	DeviceOS_UNKNOWN_DEVICE_OS DeviceOS = 0
	DeviceOS_DEVICE_OS_ANDROID DeviceOS = 1
	DeviceOS_DEVICE_OS_IOS     DeviceOS = 2
	DeviceOS_DEVICE_OS_WINDOWS DeviceOS = 3
	DeviceOS_DEVICE_OS_WEB     DeviceOS = 4
)

// Enum value maps for DeviceOS.
var (
	DeviceOS_name = map[int32]string{
		0: "UNKNOWN_DEVICE_OS",
		1: "DEVICE_OS_ANDROID",
		2: "DEVICE_OS_IOS",
		3: "DEVICE_OS_WINDOWS",
		4: "DEVICE_OS_WEB",
	}
	DeviceOS_value = map[string]int32{
		"UNKNOWN_DEVICE_OS": 0,
		"DEVICE_OS_ANDROID": 1,
		"DEVICE_OS_IOS":     2,
		"DEVICE_OS_WINDOWS": 3,
		"DEVICE_OS_WEB":     4,
	}
)

func (x DeviceOS) Enum() *DeviceOS {
	p := new(DeviceOS)
	*p = x
	return p
}

func (x DeviceOS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceOS) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_proto_enumTypes[0].Descriptor()
}

func (DeviceOS) Type() protoreflect.EnumType {
	return &file_notification_proto_enumTypes[0]
}

func (x DeviceOS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceOS.Descriptor instead.
func (DeviceOS) EnumDescriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken   string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	DeviceId      string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Manufacturer  string   `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	DeviceModel   string   `protobuf:"bytes,4,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
	DeviceOs      DeviceOS `protobuf:"varint,5,opt,name=device_os,json=deviceOs,proto3,enum=notification.DeviceOS" json:"device_os,omitempty"`
	DeviceVersion string   `protobuf:"bytes,6,opt,name=device_version,json=deviceVersion,proto3" json:"device_version,omitempty"`
	DeviceToken   string   `protobuf:"bytes,7,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RegisterRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *RegisterRequest) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *RegisterRequest) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *RegisterRequest) GetDeviceOs() DeviceOS {
	if x != nil {
		return x.DeviceOs
	}
	return DeviceOS_UNKNOWN_DEVICE_OS
}

func (x *RegisterRequest) GetDeviceVersion() string {
	if x != nil {
		return x.DeviceVersion
	}
	return ""
}

func (x *RegisterRequest) GetDeviceToken() string {
	if x != nil {
		return x.DeviceToken
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{1}
}

type UnregisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	DeviceId    string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *UnregisterRequest) Reset() {
	*x = UnregisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterRequest) ProtoMessage() {}

func (x *UnregisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterRequest.ProtoReflect.Descriptor instead.
func (*UnregisterRequest) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{2}
}

func (x *UnregisterRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UnregisterRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type UnregisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnregisterResponse) Reset() {
	*x = UnregisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterResponse) ProtoMessage() {}

func (x *UnregisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterResponse.ProtoReflect.Descriptor instead.
func (*UnregisterResponse) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{3}
}

var File_notification_proto protoreflect.FileDescriptor

var file_notification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x97, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x33, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x53, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4f, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x12, 0x0a, 0x10,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x53, 0x0a, 0x11, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x75, 0x0a, 0x08, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x53, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x53, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x52,
	0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x4f, 0x53, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x4f, 0x53, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x03, 0x12,
	0x11, 0x0a, 0x0d, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4f, 0x53, 0x5f, 0x57, 0x45, 0x42,
	0x10, 0x04, 0x32, 0xb5, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x55, 0x6e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_proto_rawDescOnce sync.Once
	file_notification_proto_rawDescData = file_notification_proto_rawDesc
)

func file_notification_proto_rawDescGZIP() []byte {
	file_notification_proto_rawDescOnce.Do(func() {
		file_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_proto_rawDescData)
	})
	return file_notification_proto_rawDescData
}

var file_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notification_proto_goTypes = []interface{}{
	(DeviceOS)(0),              // 0: notification.DeviceOS
	(*RegisterRequest)(nil),    // 1: notification.RegisterRequest
	(*RegisterResponse)(nil),   // 2: notification.RegisterResponse
	(*UnregisterRequest)(nil),  // 3: notification.UnregisterRequest
	(*UnregisterResponse)(nil), // 4: notification.UnregisterResponse
}
var file_notification_proto_depIdxs = []int32{
	0, // 0: notification.RegisterRequest.device_os:type_name -> notification.DeviceOS
	1, // 1: notification.NotificationService.Register:input_type -> notification.RegisterRequest
	3, // 2: notification.NotificationService.Unregister:input_type -> notification.UnregisterRequest
	2, // 3: notification.NotificationService.Register:output_type -> notification.RegisterResponse
	4, // 4: notification.NotificationService.Unregister:output_type -> notification.UnregisterResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notification_proto_init() }
func file_notification_proto_init() {
	if File_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterRequest); i {
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
		file_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterResponse); i {
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
			RawDescriptor: file_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_proto_goTypes,
		DependencyIndexes: file_notification_proto_depIdxs,
		EnumInfos:         file_notification_proto_enumTypes,
		MessageInfos:      file_notification_proto_msgTypes,
	}.Build()
	File_notification_proto = out.File
	file_notification_proto_rawDesc = nil
	file_notification_proto_goTypes = nil
	file_notification_proto_depIdxs = nil
}
