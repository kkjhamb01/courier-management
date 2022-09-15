package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	"gitlab.artin.ai/backend/courier-management/common/push"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	"time"
)

func publishNewRequestEventOnMessaging(ctx context.Context, newRequestEvent deliveryPb.Request) error {

	logger.Infof("publishNewRequestEventOnMessaging newRequestEvent = %+v", newRequestEvent)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeDeliveryNewRequestData(newRequestEvent)
	if err != nil {
		logger.Error("failed to marshal new request event", err, tag.Obj("newRequestEvent", newRequestEvent))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryNewRequest, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish new request event", err)
		return err
	}

	return nil
}

func publishRequestAccepted(ctx context.Context, event deliveryPb.RequestAcceptedEvent,
	senderPhoneNumber string) error {
	logger.Infof("publishRequestAccepted event = %+v", event)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestAcceptedData(event)
	if err != nil {
		logger.Error("failed to marshal new accepted ofer event", err, tag.Obj("AcceptOfferEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestAccepted, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish accepted offer event", err)
		return err
	}

	// push notification
	courierProfile, err := getCourierProfile(ctx, event.CourierId)
	if err != nil {
		logger.Error("NOTIFICATION:: failed to get courier profile", err)
		return err
	}
	courierName := courierProfile.FirstName + " " + courierProfile.LastName

	pushOfferAccepted := &push.OfferAccepted{
		Offer: &push.OfferInfo{
			CourierId:   event.CourierId,
			OfferId:     event.RequestId,
			CourierName: courierName,
			Time:        time.Now().Format("2006-01-02 15:04:05"),
		},
		CustomerPhoneNumber: senderPhoneNumber,
	}
	pushData, err := push.Encode(pushOfferAccepted)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish accepted request to push notification", err)
		return err
	}

	return nil
}

func publishRequestRejected(ctx context.Context, event deliveryPb.RequestRejectedEvent,
	senderPhoneNumber string) error {
	logger.Infof("publishRequestRejected event = %+v", event)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeDeliveryRequestRejectedData(event)
	if err != nil {
		logger.Error("failed to marshal rejected offer event", err, tag.Obj("RejectOfferEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicRequestRejected, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish rejected offer event", err)
		return err
	}

	// push notification
	courierProfile, err := getCourierProfile(ctx, event.CourierId)
	if err != nil {
		logger.Error("failed to get courier profile", err)
		return err
	}
	courierName := courierProfile.FirstName + " " + courierProfile.LastName

	pushOfferRejected := &push.OfferRejected{
		Offer: &push.OfferInfo{
			CourierId:   event.CourierId,
			OfferId:     event.RequestId,
			CourierName: courierName,
			Time:        time.Now().Format("2006-01-02 15:04:05"),
		},
		CustomerPhoneNumber: senderPhoneNumber,
	}
	pushData, err := push.Encode(pushOfferRejected)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish rejected request to push notification", err)
		return err
	}
	return nil
}

func publishRequestCancelled(ctx context.Context, event deliveryPb.RequestCancelledEvent,
	cancellationNote string, location deliveryPb.RequestLocation, requestHumanReadableId string) error {
	logger.Infof("publishRequestCancelled event = %+v, cancellationNote = %v, location = %+v, requestHumanReadableId = %v", event, cancellationNote, location, requestHumanReadableId)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeDeliveryRequestCancelledData(event)
	if err != nil {
		logger.Error("failed to marshal cancel offer event", err, tag.Obj("CancelOfferEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicRequestCancelled, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish cancelled offer event", err)
		return err
	}

	// push notification
	if event.CancelledBy != deliveryPb.CancelledBy_CUSTOMER {
		customerProfile, err := getCustomerProfile(ctx, event.CustomerId)
		if err != nil {
			logger.Error("failed to get customer profile", err)
			return err
		}
		pushOfferCancelled := &push.RideStateCancelled{
			RequesterPhoneNumber: customerProfile.PhoneNumber,
			Ride: &push.RideInfo{
				RequestId:       event.RequestId,
				LocationId:      location.Id,
				HumanReadableId: requestHumanReadableId,
				FullName:        location.FullName,
				PhoneNumber:     location.PhoneNumber,
				Time:            time.Now().Format("2006-01-02 15:04:05"),
			},
			CancellationNote: cancellationNote,
		}
		pushData, err := push.Encode(pushOfferCancelled)
		err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
		if err != nil {
			logger.Error("failed to publish cancelled request to push notification", err)
			return err
		}
	}

	return nil
}

func publishRequestArrivedOrigin(ctx context.Context, event deliveryPb.RequestArrivedOriginEvent,
	customerId string, originLocation deliveryPb.RequestLocation, requestHumanReadableId string) error {
	logger.Infof("publishRequestArrivedOrigin event = %+v, customerId = %v, originLocation = %+v, requestHumanReadableId = %v", event, customerId, originLocation, requestHumanReadableId)
	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestArrivedOriginData(event)
	if err != nil {
		logger.Error("failed to marshal request arrived event", err, tag.Obj("publishRequestArrivedOrigin", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestArrivedOrigin, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish arrived request event", err)
		return err
	}

	// push notification
	customerProfile, err := getCustomerProfile(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer profile", err)
		return err
	}
	pushOfferArrived := &push.RideStateArrived{
		RequesterPhoneNumber: customerProfile.PhoneNumber,
		SenderPhoneNumber:    originLocation.PhoneNumber,
		ReceiverPhoneNumber:  "",
		Ride: &push.RideInfo{
			RequestId:       event.RequestId,
			LocationId:      originLocation.Id,
			HumanReadableId: requestHumanReadableId,
			FullName:        originLocation.FullName,
			PhoneNumber:     originLocation.PhoneNumber,
			Time:            time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	pushData, err := push.Encode(pushOfferArrived)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish arrived request to push notification", err)
		return err
	}

	return nil
}

func publishRequestArrivedDestination(ctx context.Context, event deliveryPb.RequestArrivedDestinationEvent,
	customerId string, senderPhoneNumber string, destinationLocation deliveryPb.RequestLocation, requestHumanReadableId string) error {
	logger.Infof("publishRequestArrivedDestination event = %+v, customerId = %v, senderPhoneNumber = %v, destinationLocation = %+v, requestHumanReadableId = %v", event, customerId, senderPhoneNumber, destinationLocation, requestHumanReadableId)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestArrivedDestinationData(event)
	if err != nil {
		logger.Error("failed to marshal request arrived event", err, tag.Obj("publishRequestArrivedDestination", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestArrivedDestination, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish arrived request event", err)
		return err
	}

	// push notification
	customerProfile, err := getCustomerProfile(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer profile", err)
		return err
	}
	pushOfferArrived := &push.RideStateArrived{
		RequesterPhoneNumber: customerProfile.PhoneNumber,
		SenderPhoneNumber:    senderPhoneNumber,
		ReceiverPhoneNumber:  destinationLocation.PhoneNumber,
		Ride: &push.RideInfo{
			RequestId:       event.RequestId,
			LocationId:      destinationLocation.Id,
			HumanReadableId: requestHumanReadableId,
			FullName:        destinationLocation.FullName,
			PhoneNumber:     destinationLocation.PhoneNumber,
			Time:            time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	pushData, err := push.Encode(pushOfferArrived)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish arrived request to push notification", err)
		return err
	}

	return nil
}

func publishRequestPickedUp(ctx context.Context, event deliveryPb.RequestPickedUpEvent,
	customerId string, originLocation deliveryPb.RequestLocation, requestHumanReadableId string) error {
	logger.Infof("publishRequestPickedUp event = %+v, customerId = %v, originLocation = %+v, requestHumanReadableId = %v", event, customerId, originLocation, requestHumanReadableId)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestPickedUpData(event)
	if err != nil {
		logger.Error("failed to marshal request picked up event", err, tag.Obj("RequestPickedUpEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestPickedUp, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish picked up request event", err)
		return err
	}

	// push notification
	customerProfile, err := getCustomerProfile(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer profile", err)
		return err
	}
	pushOfferPickedUp := &push.RideStatePickedup{
		RequesterPhoneNumber: customerProfile.PhoneNumber,
		Ride: &push.RideInfo{
			RequestId:       event.RequestId,
			LocationId:      originLocation.Id,
			HumanReadableId: requestHumanReadableId,
			FullName:        originLocation.FullName,
			PhoneNumber:     originLocation.PhoneNumber,
			Time:            time.Now().Format("2006-01-02 15:04:05"),
		},
		SenderPhoneNumber: originLocation.PhoneNumber,
	}
	pushData, err := push.Encode(pushOfferPickedUp)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish picked up request to push notification", err)
		return err
	}

	return nil
}

func publishRequestDelivered(ctx context.Context, event deliveryPb.RequestDeliveredEvent,
	customerId string, destinationLocation deliveryPb.RequestLocation, requestHumanReadableId string) error {
	logger.Infof("publishRequestDelivered event = %+v, customerId = %v, destinationLocation = %+v, requestHumanReadableId = %v", event, customerId, destinationLocation, requestHumanReadableId)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestDeliveredData(event)
	if err != nil {
		logger.Error("failed to marshal request delivered event", err, tag.Obj("RequestDeliveredEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestDelivered, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish delivered request event", err)
		return err
	}

	// push notification
	customerProfile, err := getCustomerProfile(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer profile", err)
		return err
	}
	pushOfferDelivered := &push.RideStateDelivered{
		RequesterPhoneNumber: customerProfile.PhoneNumber,
		Ride: &push.RideInfo{
			RequestId:       event.RequestId,
			LocationId:      destinationLocation.Id,
			HumanReadableId: requestHumanReadableId,
			FullName:        destinationLocation.FullName,
			PhoneNumber:     destinationLocation.PhoneNumber,
			Time:            time.Now().Format("2006-01-02 15:04:05"),
		},
		ReceiverPhoneNumber: destinationLocation.PhoneNumber,
	}
	pushData, err := push.Encode(pushOfferDelivered)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish delivered request to push notification", err)
		return err
	}

	return nil
}

func publishRequestNavigatingToReceiver(event deliveryPb.RequestNavigatingToReceiverEvent) error {
	logger.Infof("publishRequestNavigatingToReceiver event = %+v", event)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestNavigatingToReceiverData(event)
	if err != nil {
		logger.Error("failed to marshal request navigating to receiver event", err, tag.Obj("RequestNavigatingToReceiverEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestNavigatingToReceiver, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish navigating to receiver request event", err)
		return err
	}

	return nil
}

func publishRequestNavigatingToSender(event deliveryPb.RequestNavigatingToSenderEvent) error {
	logger.Infof("publishRequestNavigatingToSender event = %+v", event)
	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestNavigatingToSenderData(event)
	if err != nil {
		logger.Error("failed to marshal request navigating to sender event", err, tag.Obj("RequestNavigatingToSenderEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestNavigatingToSender, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish navigating to sender request event", err)
		return err
	}

	return nil
}

func publishRequestSenderNotAvailable(ctx context.Context, event deliveryPb.RequestSenderNotAvailableEvent,
	originLocation deliveryPb.RequestLocation, customerId string, requestHumanReadableId string) error {
	logger.Infof("publishRequestSenderNotAvailable event = %+v, customerId = %v, originLocation = %+v, requestHumanReadableId = %v", event, customerId, originLocation, requestHumanReadableId)

	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestSenderNotAvailable(event)
	if err != nil {
		logger.Error("failed to marshal request sender not available event", err, tag.Obj("RequestSenderNotAvailableEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestSenderNotAvailable, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish request sender not available event", err)
		return err
	}

	// push notification
	customerProfile, err := getCustomerProfile(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer profile", err)
		return err
	}
	pushOfferSenderNotAvailable := &push.RideStateSenderIsNotAvailable{
		RequesterPhoneNumber: customerProfile.PhoneNumber,
		Ride: &push.RideInfo{
			RequestId:       event.RequestId,
			LocationId:      originLocation.Id,
			HumanReadableId: requestHumanReadableId,
			FullName:        originLocation.FullName,
			PhoneNumber:     originLocation.PhoneNumber,
			Time:            time.Now().Format("2006-01-02 15:04:05"),
		},
		SenderPhoneNumber: originLocation.PhoneNumber,
	}
	pushData, err := push.Encode(pushOfferSenderNotAvailable)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish sender not available to push notification", err)
		return err
	}

	return nil
}

func publishRequestReceiverNotAvailable(ctx context.Context, event deliveryPb.RequestReceiverNotAvailableEvent,
	customerId string, destinationLocation deliveryPb.RequestLocation, requestHumanReadableId string) error {
	logger.Infof("publishRequestReceiverNotAvailable event = %+v, customerId = %v, destinationLocation = %+v, requestHumanReadableId = %v", event, customerId, destinationLocation, requestHumanReadableId)
	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestReceiverNotAvailable(event)
	if err != nil {
		logger.Error("failed to marshal request receiver not available event", err, tag.Obj("RequestReceiverNotAvailableEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestReceiverNotAvailable, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish request receiver not available event", err)
		return err
	}

	// push notification
	customerProfile, err := getCustomerProfile(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer profile", err)
		return err
	}
	pushOfferReceiverNotAvailable := &push.RideStateReceiverIsNotAvailable{
		RequesterPhoneNumber: customerProfile.PhoneNumber,
		Ride: &push.RideInfo{
			RequestId:       event.RequestId,
			LocationId:      destinationLocation.Id,
			HumanReadableId: requestHumanReadableId,
			FullName:        destinationLocation.FullName,
			PhoneNumber:     destinationLocation.PhoneNumber,
			Time:            time.Now().Format("2006-01-02 15:04:05"),
		},
		ReceiverPhoneNumber: destinationLocation.PhoneNumber,
	}
	pushData, err := push.Encode(pushOfferReceiverNotAvailable)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish receiver not available to push notification", err)
		return err
	}

	return nil
}

func publishRequestCompleted(ctx context.Context, event deliveryPb.RequestCompletedEvent,
	customerId string, lastDestinationLocation deliveryPb.RequestLocation, requestHumanReadableId string) error {
	logger.Infof("publishRequestCompleted event = %+v, customerId = %v, lastDestinationLocation = %+v, requestHumanReadableId = %v", event, customerId, lastDestinationLocation, requestHumanReadableId)
	// publish on messaging system
	newRequestEventByte, err := messaging.EncodeRequestCompletedData(event)
	if err != nil {
		logger.Error("failed to marshal request completed event", err, tag.Obj("RequestCompletedEvent", event))
		return err
	}
	err = messaging.NatsClient().Publish(messaging.TopicDeliveryRequestCompleted, newRequestEventByte)
	if err != nil {
		logger.Error("failed to publish completed request event", err)
		return err
	}

	// push notification
	customerProfile, err := getCustomerProfile(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer profile", err)
		return err
	}
	pushOfferFinished := &push.RideStateFinished{
		RequesterPhoneNumber: customerProfile.PhoneNumber,
		Ride: &push.RideInfo{
			RequestId:       event.RequestId,
			LocationId:      lastDestinationLocation.Id,
			HumanReadableId: requestHumanReadableId,
			FullName:        lastDestinationLocation.FullName,
			PhoneNumber:     lastDestinationLocation.PhoneNumber,
			Time:            time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	pushData, err := push.Encode(pushOfferFinished)
	err = messaging.NatsClient().Publish(messaging.TopicPushNotification, pushData)
	if err != nil {
		logger.Error("failed to publish finished request to push notification", err)
		return err
	}

	return nil
}
