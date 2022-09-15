package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
)

func publishOfferCancelledEventOnMessaging(ctx context.Context, cancelEvent offeringPb.OfferCancelledEvent) error {

	logger.Infof("publishOfferCancelledEventOnMessaging cancelEvent = %+v", cancelEvent)

	// publish on messaging system
	var cancelReason deliveryPb.CancelReason
	switch cancelEvent.CancelReason {
	case offeringPb.OfferCancelledEvent_BY_CLIENT:
		cancelReason = deliveryPb.CancelReason_BY_CUSTOMER
	case offeringPb.OfferCancelledEvent_TIMEOUT:
		cancelReason = deliveryPb.CancelReason_TIMEOUT
	case offeringPb.OfferCancelledEvent_ACCEPTED_BY_ANOTHER_COURIER:
		cancelReason = deliveryPb.CancelReason_ACCEPTED_BY_ANOTHER_COURIER
	}

	var cancelledBy deliveryPb.CancelledBy
	switch cancelEvent.CancelledBy {
	case offeringPb.OfferCancelledEvent_COURIER:
		cancelledBy = deliveryPb.CancelledBy_COURIER
	case offeringPb.OfferCancelledEvent_CUSTOMER:
		cancelledBy = deliveryPb.CancelledBy_CUSTOMER
	case offeringPb.OfferCancelledEvent_SYSTEM:
		cancelledBy = deliveryPb.CancelledBy_SYSTEM
	}

	cancelEventByte, err := messaging.EncodeDeliveryRequestCancelledData(deliveryPb.RequestCancelledEvent{
		RequestId:    cancelEvent.OfferId,
		CustomerId:   cancelEvent.CustomerId,
		CancelledBy:  cancelledBy,
		UpdatedBy:    cancelEvent.UpdatedBy,
		CancelReason: cancelReason,
	})
	if err != nil {
		logger.Error("failed to marshal cancel event", err, tag.Obj("cancelEvent", cancelEvent))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicRequestCancelled, cancelEventByte)
	if err != nil {
		logger.Error("failed to publish cancel offer event", err)
		return err
	}

	return nil
}

func publishMaxRetriesOnMessaging(ctx context.Context, maxRetiesEvent offeringPb.MaxOfferRetriesEvent) error {

	logger.Infof("publishMaxRetriesOnMessaging maxRetiesEvent = %+v", maxRetiesEvent)

	// publish on messaging system
	maxRetriesEventByte, err := messaging.EncodeMaxOfferRetriesData(maxRetiesEvent)
	if err != nil {
		logger.Error("failed to marshal max retries event", err, tag.Obj("maxRetiesEvent", maxRetiesEvent))
		return err
	}

	err = messaging.NatsClient().Publish(messaging.TopicOfferMaxRetries, maxRetriesEventByte)
	if err != nil {
		logger.Error("failed to publish  offer max retries event", err, tag.Obj("maxRetiesEvent", maxRetiesEvent))
		return err
	}

	return nil
}

func publishRetryOfferRequestedOnMessaging(ctx context.Context, retryRequestEvent offeringPb.RetryOfferRequestEvent) error {

	logger.Infof("publishRetryOfferRequestedOnMessaging retryRequestEvent = %+v", retryRequestEvent)

	// publish on messaging system
	maxRetriesEventByte, err := messaging.EncodeRetryOfferRequestData(retryRequestEvent)
	if err != nil {
		logger.Error("failed to marshal retry request event", err, tag.Obj("retryRequestEvent", retryRequestEvent))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicRetryOfferRequest, maxRetriesEventByte)
	if err != nil {
		logger.Error("failed to publish offer retry request event", err)
		return err
	}

	return nil
}

func publishNewOfferSentToCouriersEventOnMessaging(ctx context.Context, newOfferSentToCouriersEvent offeringPb.NewOfferSentToCouriersEvent) error {

	logger.Infof("publishNewOfferSentToCouriersEventOnMessaging newOfferSentToCouriersEvent = %+v", newOfferSentToCouriersEvent)

	// publish new offer sent to couriers notification
	offerSentToCouriersEventByte, err := messaging.EncodeNewOfferSentToCouriersData(newOfferSentToCouriersEvent)
	if err != nil {
		logger.Error("failed to marshal cancel event", err, tag.Obj("offerSentToCouriersEventByte", offerSentToCouriersEventByte))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicNewOfferSentToCouriers, offerSentToCouriersEventByte)
	if err != nil {
		logger.Error("failed to publish 'new offer sent to couriers' event", err, tag.Obj("courierIds", newOfferSentToCouriersEvent.CourierIds))
		return err
	}

	return nil
}
