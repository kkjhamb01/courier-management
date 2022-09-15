package messaging

import (
	"context"
	"errors"
	"github.com/nats-io/nats.go"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"gitlab.artin.ai/backend/courier-management/offering/business"
	"gitlab.artin.ai/backend/courier-management/offering/pubsub"
)

func onRequestCancelled(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: offer cancelled")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on offer cancelled event", err, tag.Obj("msg", msg))
		return err
	}

	cancelledEvent, err := messaging.DecodeDeliveryRequestCancelledData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}

	logger.Infof("onRequestCancelled event = %+v", cancelledEvent)

	var cancelReason offeringPb.OfferCancelledEvent_Reason
	switch cancelledEvent.CancelReason {
	case deliveryPb.CancelReason_BY_CUSTOMER:
		cancelReason = offeringPb.OfferCancelledEvent_BY_CLIENT
	case deliveryPb.CancelReason_ACCEPTED_BY_ANOTHER_COURIER:
		cancelReason = offeringPb.OfferCancelledEvent_ACCEPTED_BY_ANOTHER_COURIER
	case deliveryPb.CancelReason_TIMEOUT:
		cancelReason = offeringPb.OfferCancelledEvent_TIMEOUT
	}

	var cancelledBy offeringPb.OfferCancelledEvent_CancelledBy
	switch cancelledEvent.CancelledBy {
	case deliveryPb.CancelledBy_COURIER:
		cancelledBy = offeringPb.OfferCancelledEvent_COURIER
	case deliveryPb.CancelledBy_CUSTOMER:
		cancelledBy = offeringPb.OfferCancelledEvent_CUSTOMER
	case deliveryPb.CancelledBy_SYSTEM:
		cancelledBy = offeringPb.OfferCancelledEvent_SYSTEM
	}

	offerCancelledEvent := offeringPb.OfferCancelledEvent{
		OfferId:      cancelledEvent.RequestId,
		CustomerId:   cancelledEvent.CustomerId,
		UpdatedBy:    cancelledEvent.UpdatedBy,
		CancelReason: cancelReason,
		CancelledBy:  cancelledBy,
	}
	err = business.OnOfferCancelled(ctx, offerCancelledEvent)
	if err != nil {
		logger.Error("failed to handle message", err)
		return err
	}

	// notify the customer about the event
	if offerCancelledEvent.CancelReason == offeringPb.OfferCancelledEvent_BY_CLIENT {
		pubsub.PublishById(pubsub.CancelOfferEvent(offerCancelledEvent), cancelledEvent.CustomerId)
		if len(cancelledEvent.CourierId) > 0 {
			pubsub.PublishById(pubsub.CancelOfferEvent(offerCancelledEvent), cancelledEvent.CourierId)
		}
	}

	//pubsub.UnsubscribeById(pubsub.TopicNewOffer, cancelledEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicAcceptOffer, cancelledEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicCancelOffer, cancelledEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicAcceptOfferFailed, cancelledEvent.CourierId)
	//pubsub.UnsubscribeById(pubsub.TopicRejectOfferFailed, cancelledEvent.CourierId)

	//if cancelReason != offeringPb.OfferCancelledEvent_ACCEPTED_BY_ANOTHER_COURIER &&
	//	cancelReason != offeringPb.OfferCancelledEvent_TIMEOUT {
	//	pubsub.UnsubscribeById(pubsub.TopicRetryOffer, cancelledEvent.CustomerId)
	//	pubsub.UnsubscribeById(pubsub.TopicMaxOfferRetries, cancelledEvent.CustomerId)
	//	pubsub.UnsubscribeById(pubsub.TopicAcceptOffer, cancelledEvent.CustomerId)
	//	pubsub.UnsubscribeById(pubsub.TopicCancelOffer, cancelledEvent.CustomerId)
	//	pubsub.UnsubscribeById(pubsub.TopicOfferSentToCouriers, cancelledEvent.CustomerId)
	//	pubsub.UnsubscribeById(pubsub.TopicOfferCreationFailed, cancelledEvent.CustomerId)
	//}

	return nil
}
