package messaging

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/common/messaging"
	"github.com/kkjhamb01/courier-management/offering/pubsub"
	"github.com/nats-io/nats.go"
)

func onOfferCreationFailed(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: offer creation failed")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on offer creation failed event", err, tag.Obj("msg", msg))
		return err
	}

	event, err := messaging.DecodeOfferCreationFailedData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onOfferCreationFailed event = %+v", event)

	// notify the customer about the event
	pubsub.PublishById(pubsub.OfferCreationFailedEventData(event), event.Offer.CustomerId)

	return nil
}

func onAcceptOfferFailed(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: accept offer failed")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on accept offer failed event", err, tag.Obj("msg", msg))
		return err
	}

	event, err := messaging.DecodeAcceptOfferFailedData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onAcceptOfferFailed event = %+v", event)

	// notify the accepting courier about the event
	pubsub.PublishById(pubsub.AcceptOfferFailedEventData(event), event.CourierId)

	return nil
}

func onRejectOfferFailed(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: reject offer failed")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on reject offer failed event", err, tag.Obj("msg", msg))
		return err
	}

	event, err := messaging.DecodeRejectOfferFailedData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onRejectOfferFailed event = %+v", event)

	// notify the accepting courier about the event
	pubsub.PublishById(pubsub.RejectOfferFailedEventData(event), event.CourierId)

	return nil
}
