package model

import (
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type CourierStatusLog struct {
	CourierId string
	Status    string
	Time      time.Time
	Base
}

func (s CourierStatusLog) ToProto() (offeringPb.CourierStatusLog, error) {
	statusProto, ok := commonPb.CourierStatus_value[s.Status]
	if !ok {
		err := errors.New("failed to match courier status in CourierStatus_value map")
		logger.Error("the courier status is not valid", err)
		return offeringPb.CourierStatusLog{}, err
	}

	return offeringPb.CourierStatusLog{
		CourierId: s.CourierId,
		Status:    commonPb.CourierStatus(statusProto),
		Time:      timestamppb.New(s.Time),
	}, nil
}

func (s CourierStatusLog) ToProtoP() (*offeringPb.CourierStatusLog, error) {
	proto, err := s.ToProto()
	return &proto, err
}
