package model

import (
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RideStatus struct {
	RideLocationId   string
	RequestId        string
	Status           string
	CancellationNote string
	Base
}

func (r RideStatus) ToProto() deliveryPb.RideStatus {
	return deliveryPb.RideStatus{
		Id:               r.ID,
		RequestId:        r.RequestId,
		RideLocationId:   r.RideLocationId,
		Status:           deliveryPb.RequestStatus(deliveryPb.RequestStatus_value[r.Status]),
		CancellationNote: r.CancellationNote,
		CreatedAt:        timestamppb.New(r.CreatedAt),
	}
}

func (r RideStatus) ToProtoP() *deliveryPb.RideStatus {
	proto := r.ToProto()
	return &proto
}
