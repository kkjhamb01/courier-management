package messaging

import (
	"context"
	"errors"
	"github.com/nats-io/nats.go"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	"gitlab.artin.ai/backend/courier-management/delivery/business"
)

func onOfferCreationFailed(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: onOfferCreationFailed")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on new offer event", err, tag.Obj("msg", msg))
		return err
	}

	failedEvent, err := messaging.DecodeOfferCreationFailedData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onOfferCreationFailed failedEvent = %+v", failedEvent)

	err = business.OnRequestCreationFailed(ctx, failedEvent.Offer.Id)
	if err != nil {
		logger.Error("changing delivery data after offer creation failure, failed", err)
		return err
	}
	return nil
}
