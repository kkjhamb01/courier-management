package model

import (
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RequestLocation struct {
	FullName            string
	PhoneNumber         string
	AddressDetails      *string
	Lat                 float64
	Lon                 float64
	CourierInstructions *string
	RequestId           string `gorm:"type:BINARY(36);"`
	IsOrigin            bool
	Order               int
	Base
}

func (r RequestLocation) ToProto() deliveryPb.RequestLocation {
	requestLocation := deliveryPb.RequestLocation{
		Id:          r.ID,
		FullName:    r.FullName,
		PhoneNumber: r.PhoneNumber,
		Lat:         r.Lat,
		Lon:         r.Lon,
		Order:       int32(r.Order),
		CreatedAt:   timestamppb.New(r.CreatedAt),
		UpdatedAt:   timestamppb.New(r.UpdatedAt),
	}

	if r.AddressDetails != nil {
		requestLocation.AddressDetails = *r.AddressDetails
	}

	if r.CourierInstructions != nil {
		requestLocation.CourierInstructions = *r.CourierInstructions
	}

	return requestLocation
}

func (r RequestLocation) ToProtoP() *deliveryPb.RequestLocation {
	proto := r.ToProto()
	return &proto
}
