package model

import (
	"time"

	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Request struct {
	CustomerId               string
	CourierId                string
	Locations                []RequestLocation `gorm:"foreignKey:RequestId"`
	EstimatedDuration        int64
	EstimatedDistanceMeter   int
	FinalPrice               float64
	FinalPriceCurrency       string
	VehicleType              string
	RequiredWorkers          int32
	HumanReadableId          string
	Status                   string
	LastProcessedDestination int
	ScheduleOn               *time.Time
	CancellingReason         string
	CancelledBy              string
	UpdatedBy                string
	Base
}

func (r Request) ToProto() deliveryPb.Request {
	var locationsSize int
	if len(r.Locations) > 0 {
		locationsSize = len(r.Locations) - 1
	} else {
		locationsSize = 0
	}

	destinationsProto := make([]*deliveryPb.RequestLocation, 0, locationsSize)
	var originProto *deliveryPb.RequestLocation
	for _, destination := range r.Locations {
		if destination.IsOrigin {
			originProto = destination.ToProtoP()
		} else {
			destinationsProto = append(destinationsProto, destination.ToProtoP())
		}
	}

	vehicleTypeProto := commonPb.VehicleType_value[r.VehicleType]

	statusProto := deliveryPb.RequestStatus_value[r.Status]

	request := deliveryPb.Request{
		Id:                       r.ID,
		HumanReadableId:          r.HumanReadableId,
		CustomerId:               r.CustomerId,
		CourierId:                r.CourierId,
		Origin:                   originProto,
		Destinations:             destinationsProto,
		EstimatedDuration:        r.EstimatedDuration,
		EstimatedDistanceMeter:   int32(r.EstimatedDistanceMeter),
		FinalPrice:               r.FinalPrice,
		FinalPriceCurrency:       r.FinalPriceCurrency,
		VehicleType:              commonPb.VehicleType(vehicleTypeProto),
		RequiredWorkers:          r.RequiredWorkers,
		Status:                   deliveryPb.RequestStatus(statusProto),
		LastProcessedDestination: int32(r.LastProcessedDestination),
		CreatedAt:                timestamppb.New(r.CreatedAt),
		UpdatedAt:                timestamppb.New(r.UpdatedAt),
		CalculatedPrice:          taxation(r),
	}

	if r.ScheduleOn != nil {
		request.ScheduledOn = timestamppb.New(*r.ScheduleOn)
	}

	return request
}

func taxation(r Request) float64 {
	return r.FinalPrice * 0.9
}

func (r Request) ToProtoP() *deliveryPb.Request {
	proto := r.ToProto()
	return &proto
}
