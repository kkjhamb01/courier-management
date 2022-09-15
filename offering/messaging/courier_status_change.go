package messaging

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nats-io/nats.go"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	"gitlab.artin.ai/backend/courier-management/offering/business"
	partyBusiness "gitlab.artin.ai/backend/courier-management/party/business"
	"gitlab.artin.ai/backend/courier-management/party/proto"
	"time"
)

func onCourierStatusChange(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: onCourierStatusChange")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on CourierStatusChange event", err, tag.Obj("msg", msg))
		return err
	}

	var event partyBusiness.UpdateCourierStatusEvent
	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("CourierStatusChange event: %v", event)

	logger.Infof("onCourierStatusChange event = %+v", event)

	var status commonPb.CourierStatus
	switch event.Status {
	case proto.UserStatus_USER_STATUS_ENABLED:
		status = commonPb.CourierStatus_ENABLED
	case proto.UserStatus_USER_STATUS_DISABLED:
		status = commonPb.CourierStatus_DISABLED
	case proto.UserStatus_USER_STATUS_BLOCKED:
		status = commonPb.CourierStatus_BLOCKED
	case proto.UserStatus_USER_STATUS_AVAILABLE:
		status = commonPb.CourierStatus_AVAILABLE
	case proto.UserStatus_UNKNOWN_USER_STATUS:
		status = commonPb.CourierStatus_UNKNOWN
	}

	eventTime, err := time.Parse("2006-01-02 15:04:05", event.Time)
	if err != nil {
		logger.Error("failed to parse time", err, tag.Str("event.Time", event.Time))
		return err
	}
	err = business.OnCourierStatusChange(ctx, event.CourierId, status, eventTime)
	if err != nil {
		logger.Error("failed to handle courier status change", err)
		return err
	}

	logger.Infof("onCourierStatusChange event handled successfully")
	return nil
}
