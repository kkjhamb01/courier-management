// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: offering.proto

package offeringPb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_offering_proto protoreflect.FileDescriptor

var file_offering_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x1a, 0x16, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x09, 0x0a, 0x08, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0xa1, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x76,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x40, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69,
	0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x9f, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x4c, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x7e, 0x0a, 0x1a, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x44, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x80, 0x01, 0x0a, 0x1b, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x45, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x6e, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6b, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x8e, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73,
	0x12, 0x3a, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x43, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa9, 0x01, 0x0a, 0x1a,
	0x48, 0x61, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x69, 0x64, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x43, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x61,
	0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x69, 0x64, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x44, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x48, 0x61, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x69,
	0x64, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa9, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x43, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x41, 0x6e,
	0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x6e, 0x2e, 0x61, 0x69, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x67, 0x6f, 0x3b, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_offering_proto_goTypes = []interface{}{
	(*SetCourierLiveLocationRequest)(nil),       // 0: artin.couriermanagement.offering.SetCourierLiveLocationRequest
	(*GetCourierLiveLocationRequest)(nil),       // 1: artin.couriermanagement.offering.GetCourierLiveLocationRequest
	(*emptypb.Empty)(nil),                       // 2: google.protobuf.Empty
	(*SetCourierLocationRequest)(nil),           // 3: artin.couriermanagement.offering.SetCourierLocationRequest
	(*GetNearbyCouriersRequest)(nil),            // 4: artin.couriermanagement.offering.GetNearbyCouriersRequest
	(*HadCustomerRideWithCourierRequest)(nil),   // 5: artin.couriermanagement.offering.HadCustomerRideWithCourierRequest
	(*GetOfferCourierAndCustomerRequest)(nil),   // 6: artin.couriermanagement.offering.GetOfferCourierAndCustomerRequest
	(*SetCourierLiveLocationResponse)(nil),      // 7: artin.couriermanagement.offering.SetCourierLiveLocationResponse
	(*GetCourierLiveLocationResponse)(nil),      // 8: artin.couriermanagement.offering.GetCourierLiveLocationResponse
	(*CourierSubscriptionOnOfferResponse)(nil),  // 9: artin.couriermanagement.offering.CourierSubscriptionOnOfferResponse
	(*CustomerSubscriptionOnOfferResponse)(nil), // 10: artin.couriermanagement.offering.CustomerSubscriptionOnOfferResponse
	(*GetNearbyCouriersResponse)(nil),           // 11: artin.couriermanagement.offering.GetNearbyCouriersResponse
	(*HadCustomerRideWithCourierResponse)(nil),  // 12: artin.couriermanagement.offering.HadCustomerRideWithCourierResponse
	(*GetOfferCourierAndCustomerResponse)(nil),  // 13: artin.couriermanagement.offering.GetOfferCourierAndCustomerResponse
}
var file_offering_proto_depIdxs = []int32{
	0,  // 0: artin.couriermanagement.offering.Offering.SetCourierLiveLocation:input_type -> artin.couriermanagement.offering.SetCourierLiveLocationRequest
	1,  // 1: artin.couriermanagement.offering.Offering.GetCourierLiveLocation:input_type -> artin.couriermanagement.offering.GetCourierLiveLocationRequest
	2,  // 2: artin.couriermanagement.offering.Offering.CourierSubscriptionOnOffer:input_type -> google.protobuf.Empty
	2,  // 3: artin.couriermanagement.offering.Offering.CustomerSubscriptionOnOffer:input_type -> google.protobuf.Empty
	3,  // 4: artin.couriermanagement.offering.Offering.SetCourierLocation:input_type -> artin.couriermanagement.offering.SetCourierLocationRequest
	4,  // 5: artin.couriermanagement.offering.Offering.GetNearbyCouriers:input_type -> artin.couriermanagement.offering.GetNearbyCouriersRequest
	5,  // 6: artin.couriermanagement.offering.Offering.HadCustomerRideWithCourier:input_type -> artin.couriermanagement.offering.HadCustomerRideWithCourierRequest
	6,  // 7: artin.couriermanagement.offering.Offering.GetOfferCourierAndCustomer:input_type -> artin.couriermanagement.offering.GetOfferCourierAndCustomerRequest
	7,  // 8: artin.couriermanagement.offering.Offering.SetCourierLiveLocation:output_type -> artin.couriermanagement.offering.SetCourierLiveLocationResponse
	8,  // 9: artin.couriermanagement.offering.Offering.GetCourierLiveLocation:output_type -> artin.couriermanagement.offering.GetCourierLiveLocationResponse
	9,  // 10: artin.couriermanagement.offering.Offering.CourierSubscriptionOnOffer:output_type -> artin.couriermanagement.offering.CourierSubscriptionOnOfferResponse
	10, // 11: artin.couriermanagement.offering.Offering.CustomerSubscriptionOnOffer:output_type -> artin.couriermanagement.offering.CustomerSubscriptionOnOfferResponse
	2,  // 12: artin.couriermanagement.offering.Offering.SetCourierLocation:output_type -> google.protobuf.Empty
	11, // 13: artin.couriermanagement.offering.Offering.GetNearbyCouriers:output_type -> artin.couriermanagement.offering.GetNearbyCouriersResponse
	12, // 14: artin.couriermanagement.offering.Offering.HadCustomerRideWithCourier:output_type -> artin.couriermanagement.offering.HadCustomerRideWithCourierResponse
	13, // 15: artin.couriermanagement.offering.Offering.GetOfferCourierAndCustomer:output_type -> artin.couriermanagement.offering.GetOfferCourierAndCustomerResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_offering_proto_init() }
func file_offering_proto_init() {
	if File_offering_proto != nil {
		return
	}
	file_offering_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_offering_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_offering_proto_goTypes,
		DependencyIndexes: file_offering_proto_depIdxs,
	}.Build()
	File_offering_proto = out.File
	file_offering_proto_rawDesc = nil
	file_offering_proto_goTypes = nil
	file_offering_proto_depIdxs = nil
}
