package model

import (
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type SavedLocation struct {
	CustomerId          string
	Name                string
	FullName            string
	PhoneNumber         *string
	AddressDetails      *string
	Lat                 float64
	Lon                 float64
	CourierInstructions *string
	Base
}

func (a SavedLocation) ToProto() deliveryPb.SavedLocation {
	savedLocationProto := deliveryPb.SavedLocation{
		Id:        a.ID,
		Name:      a.Name,
		FullName:  a.FullName,
		Lat:       a.Lat,
		Lon:       a.Lon,
		CreatedAt: timestamppb.New(a.CreatedAt),
		UpdatedAt: timestamppb.New(a.UpdatedAt),
	}

	if a.PhoneNumber != nil {
		savedLocationProto.PhoneNumber = &wrapperspb.StringValue{Value: *a.PhoneNumber}
	}
	if a.AddressDetails != nil {
		savedLocationProto.AddressDetails = &wrapperspb.StringValue{Value: *a.AddressDetails}
	}
	if a.CourierInstructions != nil {
		savedLocationProto.CourierInstructions = &wrapperspb.StringValue{Value: *a.CourierInstructions}
	}

	return savedLocationProto
}

func (a SavedLocation) ToProtoP() *deliveryPb.SavedLocation {
	proto := a.ToProto()
	return &proto
}
