package messaging

import (
	"context"
	"errors"
	"github.com/nats-io/nats.go"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"gitlab.artin.ai/backend/courier-management/offering/business"
)

func onRequestRejected(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: offer rejected")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on offer rejected event", err, tag.Obj("msg", msg))
		return err
	}

	rejectedEvent, err := messaging.DecodeDeliveryRequestRejectedData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}

	logger.Infof("onRequestRejected rejectedEvent = %+v", rejectedEvent)

	err = business.OnOfferRejected(ctx, rejectedEvent.RequestId, rejectedEvent.CourierId)
	if err != nil {
		logger.Error("failed to handle message", err)
		publishErr := publishErrorDuringRejectOffer(ctx, offeringPb.RejectOfferFailedEvent{
			OfferId:    rejectedEvent.RequestId,
			CustomerId: rejectedEvent.CustomerId,
			CourierId:  rejectedEvent.CourierId,
			Msg:        err.Error(),
		})
		if publishErr != nil {
			logger.Error("important: failed to publish offer creation failure", err)
		}
		return err
	}

	//pubsub.UnsubscribeById(pubsub.TopicNewOffer, rejectedEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicAcceptOffer, rejectedEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicCancelOffer, rejectedEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicAcceptOfferFailed, rejectedEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicRejectOfferFailed, rejectedEvent.CourierId)
	//
	return nil
}

func publishErrorDuringRejectOffer(ctx context.Context, failureEvent offeringPb.RejectOfferFailedEvent) error {
	failureEventByte, err := messaging.EncodeRejectOfferFailedData(failureEvent)
	if err != nil {
		logger.Error("failed to marshal failure event", err, tag.Obj("failureEvent", failureEvent))
		return err
	}

	logger.Infof("publishErrorDuringRejectOffer failureEventByte = %+v", failureEventByte)

	err = messaging.NatsClient().Publish(messaging.TopicRejectOfferFailed, failureEventByte)
	if err != nil {
		logger.Error("failed to publish 'offer reject failed' event", err, tag.Obj("failureEvent", failureEvent))
		return err
	}

	return nil
}
