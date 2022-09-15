package messaging

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/common/messaging"
	offeringPb "github.com/kkjhamb01/courier-management/grpc/offering/go"
	"github.com/kkjhamb01/courier-management/offering/business"
	"github.com/kkjhamb01/courier-management/offering/pubsub"
	"github.com/kkjhamb01/courier-management/offering/storage"
	"github.com/nats-io/nats.go"
)

var createdRequests map[string]bool

func init() {
	createdRequests = make(map[string]bool, 100)
}

func onNewRequestCreated(ctx context.Context, msg *nats.Msg) error {
	// TODO move ack method call to the end of the func
	msg.Ack()

	logger.Infof("msg received: new request created")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on new offer event", err, tag.Obj("msg", msg))
		return err
	}

	newRequest, err := messaging.DecodeDeliveryNewRequestData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onNewRequestCreated newRequest = %+v", newRequest)

	// check whether we are already processing the request
	if createdRequests[newRequest.Id] {
		return nil
	}
	createdRequests[newRequest.Id] = true
	isOfferCreated, err := storage.CreateTx().IsOfferCreatedEver(ctx, newRequest.Id)
	if err != nil {
		logger.Error("failed to check whether the offer is created", err)
		return err
	}
	if isOfferCreated {
		return nil
	}

	destinations := make([]*offeringPb.Location, len(newRequest.Destinations))
	for i, requestDest := range newRequest.Destinations {
		destinations[i] = &offeringPb.Location{
			PhoneNumber:         requestDest.PhoneNumber,
			FullName:            requestDest.FullName,
			AddressDetails:      requestDest.AddressDetails,
			Lat:                 requestDest.Lat,
			Lon:                 requestDest.Lon,
			Order:               requestDest.Order,
			CourierInstructions: requestDest.CourierInstructions,
		}
	}

	offer := offeringPb.Offer{
		Id:         newRequest.Id,
		CustomerId: newRequest.CustomerId,
		Source: &offeringPb.Location{
			PhoneNumber:         newRequest.Origin.PhoneNumber,
			FullName:            newRequest.Origin.FullName,
			AddressDetails:      newRequest.Origin.AddressDetails,
			Lat:                 newRequest.Origin.Lat,
			Lon:                 newRequest.Origin.Lon,
			CourierInstructions: newRequest.Origin.CourierInstructions,
		},
		Destinations:    destinations,
		VehicleType:     newRequest.VehicleType,
		RequiredWorkers: newRequest.RequiredWorkers,
		Price:           newRequest.FinalPrice,
		Currency:        newRequest.FinalPriceCurrency,
	}
	err = business.OnNewOfferCreated(ctx, offer)
	if err != nil {
		logger.Error("failed to handle message", err)
		publishErr := publishErrorDuringOfferCreation(ctx, offeringPb.OfferCreationFailedEvent{
			Offer: &offer,
			Msg:   err.Error(),
		})
		if publishErr != nil {
			logger.Error("important: failed to publish offer creation failure", err)
		}
		return err
	}

	delete(createdRequests, newRequest.Id)

	return nil
}

func onOfferRetryRequested(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: offer retry requested")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on offer retry event", err, tag.Obj("msg", msg))
		return err
	}

	event, err := messaging.DecodeRetryOfferRequestData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onNewRequestCreated event = %+v", event)

	// notify the customer about the event
	pubsub.PublishById(pubsub.RetryOfferEventData(event), event.Offer.CustomerId)

	err = business.OnNewOfferCreated(ctx, *event.Offer)
	if err != nil {
		logger.Error("failed to handle message", err)
		publishErr := publishErrorDuringOfferCreation(ctx, offeringPb.OfferCreationFailedEvent{
			Offer: event.Offer,
			Msg:   err.Error(),
		})
		if publishErr != nil {
			logger.Error("important: failed to publish offer creation failure", err)
		}
		return err
	}

	return nil
}

func onOfferMaxRetries(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: offer max retries")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on offer retry event", err, tag.Obj("msg", msg))
		return err
	}

	event, err := messaging.DecodeMaxOfferRetriesData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onOfferMaxRetries event = %+v", event)

	// notify the customer about the event
	pubsub.PublishById(pubsub.MaxOfferRetiesEventData(event), event.CustomerId)

	return nil
}

func onOfferSentToCouriers(ctx context.Context, msg *nats.Msg) error {
	logger.Infof("msg received: offer sent to couriers")

	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on offer retry event", err, tag.Obj("msg", msg))
		return err
	}

	event, err := messaging.DecodeNewOfferSentToCouriersData(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return err
	}
	logger.Infof("onOfferSentToCouriers event = %+v", event)

	// notify the customer about the event
	pubsub.PublishById(pubsub.OfferSentToCouriersEventData(event), event.Offer.CustomerId)

	return nil
}

func publishErrorDuringOfferCreation(ctx context.Context, failureEvent offeringPb.OfferCreationFailedEvent) error {
	// publish new offer sent to couriers notification
	failureEventByte, err := messaging.EncodeOfferCreationFailedData(failureEvent)
	if err != nil {
		logger.Error("failed to marshal failure event", err, tag.Obj("failureEvent", failureEvent))
		return err
	}

	logger.Infof("publishErrorDuringOfferCreation failureEventByte = %+v", failureEventByte)

	err = messaging.NatsClient().Publish(messaging.TopicOfferCreationFailed, failureEventByte)
	if err != nil {
		logger.Error("failed to publish 'offer creation failed' event", err, tag.Obj("failureEvent", failureEvent))
		return err
	}

	return nil
}
