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
	"gitlab.artin.ai/backend/courier-management/offering/pubsub"
)

func onRequestAccepted(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("onRequestAccepted")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("NOTIFICATION:: no data is supplied on offer accepted event", err, tag.Obj("msg", msg))
		return err
	}

	event, err := messaging.DecodeRequestAcceptedData(msg.Data)
	if err != nil {
		logger.Error("NOTIFICATION:: failed to decode data", err)
		return err
	}

	logger.Infof("onRequestAccepted event = %+v", event)

	err = business.OnOfferAccepted(ctx, offeringPb.OfferAcceptedEvent{
		OfferId:    event.RequestId,
		CustomerId: event.CustomerId,
		CourierId:  event.CourierId,
		Desc:       event.Desc,
	})
	if err != nil {
		logger.Error("NOTIFICATION:: failed to handle offer accepted message", err)
		publishErr := publishErrorDuringAcceptOffer(ctx, offeringPb.AcceptOfferFailedEvent{
			OfferId:    event.RequestId,
			CustomerId: event.CustomerId,
			CourierId:  event.CourierId,
			Msg:        err.Error(),
		})
		if publishErr != nil {
			logger.Error("important: failed to publish offer creation failure", err)
		}
		return err
	}

	// notify the customer about the event
	pubsub.PublishById(pubsub.AcceptOfferEvent(offeringPb.OfferAcceptedEvent{
		OfferId:    event.RequestId,
		CustomerId: event.CustomerId,
		CourierId:  event.CourierId,
		Desc:       event.Desc,
	}), event.CustomerId)

	// notify the courier about the event
	pubsub.PublishById(pubsub.AcceptOfferEvent(offeringPb.OfferAcceptedEvent{
		OfferId:    event.RequestId,
		CustomerId: event.CustomerId,
		CourierId:  event.CourierId,
		Desc:       event.Desc,
	}), event.CourierId)

	return nil
}

func publishErrorDuringAcceptOffer(ctx context.Context, failureEvent offeringPb.AcceptOfferFailedEvent) error {

	logger.Infof("publishErrorDuringAcceptOffer failureEvent = %+v", failureEvent)

	failureEventByte, err := messaging.EncodeAcceptOfferFailedData(failureEvent)
	if err != nil {
		logger.Error("failed to marshal failure event", err, tag.Obj("failureEvent", failureEvent))
		return err
	}

	err = messaging.NatsClient().Publish(messaging.TopicAcceptOfferFailed, failureEventByte)
	if err != nil {
		logger.Error("failed to publish 'offer accept failed' event", err, tag.Obj("failureEvent", failureEvent))
		return err
	}

	return nil
}
