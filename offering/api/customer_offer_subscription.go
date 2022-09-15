package api

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/offering/business"
	"github.com/kkjhamb01/courier-management/offering/pubsub"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func (s serverImpl) CustomerSubscriptionOnOffer(_ *empty.Empty, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer) error {
	ctx := stream.Context()

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return err
	}

	tokenUser, ok := ctx.Value("user").(security.User)
	if !ok {
		err := errors.New("failed to cast user value found in context to security.User")
		logger.Error("failed to get user", err)
		return err
	}
	customerId := tokenUser.Id

	logger.Infof("CustomerSubscriptionOnOffer customerId = %v", customerId)

	isActive, err := business.IsCustomerActive(ctx, customerId)
	if err != nil {
		logger.Error("failed to check if the customer is active", err)
		return err
	}
	if !isActive {
		err = fmt.Errorf("the customer %v is not active", customerId)
		logger.Error("failed to subscribe customer", err)
		return err
	}

	retryOfferCh := pubsub.SubscribeById(pubsub.TopicRetryOffer, customerId)
	maxOfferRetriesCh := pubsub.SubscribeById(pubsub.TopicMaxOfferRetries, customerId)
	acceptOfferCh := pubsub.SubscribeById(pubsub.TopicAcceptOffer, customerId)
	cancelOfferCh := pubsub.SubscribeById(pubsub.TopicCancelOffer, customerId)
	offerSentToCouriersCh := pubsub.SubscribeById(pubsub.TopicOfferSentToCouriers, customerId)
	offerCreationFailedCh := pubsub.SubscribeById(pubsub.TopicOfferCreationFailed, customerId)

	wg := sync.WaitGroup{}
	wg.Add(6)

	go listenToCustomerEvent(ctx, pubsub.TopicRetryOffer, retryOfferCh, customerId, stream, courierRetryOfferRequestEventHandler, &wg)
	go listenToCustomerEvent(ctx, pubsub.TopicMaxOfferRetries, maxOfferRetriesCh, customerId, stream, courierMaxOfferRetriesEventHandler, &wg)
	go listenToCustomerEvent(ctx, pubsub.TopicAcceptOffer, acceptOfferCh, customerId, stream, customerOfferAcceptedEventHandler, &wg)
	go listenToCustomerEvent(ctx, pubsub.TopicCancelOffer, cancelOfferCh, customerId, stream, customerCancelEventHandler, &wg)
	go listenToCustomerEvent(ctx, pubsub.TopicOfferSentToCouriers, offerSentToCouriersCh, customerId, stream, customerOfferSentToCouriersEventHandler, &wg)
	go listenToCustomerEvent(ctx, pubsub.TopicOfferCreationFailed, offerCreationFailedCh, customerId, stream, customerOfferCreationFailedEventHandler, &wg)

	wg.Wait()

	return nil
}

func courierRetryOfferRequestEventHandler(event pubsub.Event, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer) {

	logger.Infof("courierRetryOfferRequestEventHandler event = %+v", event)

	// get the event type
	retryOfferEvent := event.RetryOfferEventData()

	// create the response
	response := &offeringPb.CustomerSubscriptionOnOfferResponse{
		Event: &offeringPb.CustomerSubscriptionOnOfferResponse_RetryOfferEvent{
			RetryOfferEvent: &retryOfferEvent,
		},
		ResponseType: offeringPb.CustomerSubscriptionOnOfferResponse_TypeRetryOfferRequestEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func courierMaxOfferRetriesEventHandler(event pubsub.Event, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer) {

	logger.Infof("courierMaxOfferRetriesEventHandler event = %+v", event)

	// get the event type
	maxOfferRetriesEvent := event.MaxOfferRetriesEventData()

	// create the response
	response := &offeringPb.CustomerSubscriptionOnOfferResponse{
		Event: &offeringPb.CustomerSubscriptionOnOfferResponse_MaxOfferRetries{
			MaxOfferRetries: &maxOfferRetriesEvent,
		},
		ResponseType: offeringPb.CustomerSubscriptionOnOfferResponse_TypeMaxOfferRetriesEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func customerOfferAcceptedEventHandler(event pubsub.Event, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer) {

	logger.Infof("customerOfferAcceptedEventHandler event = %+v", event)

	// get the event type
	acceptOfferEvent := event.AcceptOfferEventData()

	// create the response
	response := &offeringPb.CustomerSubscriptionOnOfferResponse{
		Event: &offeringPb.CustomerSubscriptionOnOfferResponse_OfferAcceptedEvent{
			OfferAcceptedEvent: &acceptOfferEvent,
		},
		ResponseType: offeringPb.CustomerSubscriptionOnOfferResponse_TypeOfferAcceptedEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func customerCancelEventHandler(event pubsub.Event, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer) {

	logger.Infof("customerCancelEventHandler event = %+v", event)

	// get the event type
	cancelOfferEvent := event.CancelOfferEventData()

	// create the response
	response := &offeringPb.CustomerSubscriptionOnOfferResponse{
		Event: &offeringPb.CustomerSubscriptionOnOfferResponse_OfferCancelledEvent{
			OfferCancelledEvent: &cancelOfferEvent,
		},
		ResponseType: offeringPb.CustomerSubscriptionOnOfferResponse_TypeOfferCancelledEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func customerOfferSentToCouriersEventHandler(event pubsub.Event, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer) {

	logger.Infof("customerOfferSentToCouriersEventHandler event = %+v", event)

	// get the event type
	offerSentToCouriersEvent := event.OfferSentToCouriersEventData()

	// create the response
	response := &offeringPb.CustomerSubscriptionOnOfferResponse{
		Event: &offeringPb.CustomerSubscriptionOnOfferResponse_OfferSentToCouriers{
			OfferSentToCouriers: &offerSentToCouriersEvent,
		},
		ResponseType: offeringPb.CustomerSubscriptionOnOfferResponse_TypeNewOfferSentToCouriersEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func customerOfferCreationFailedEventHandler(event pubsub.Event, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer) {

	logger.Infof("customerOfferCreationFailedEventHandler event = %+v", event)

	// get the event type
	OfferCreationFailedEvent := event.OfferCreationFailedEventData()

	// create the response
	response := &offeringPb.CustomerSubscriptionOnOfferResponse{
		Event: &offeringPb.CustomerSubscriptionOnOfferResponse_OfferCreationFailedEvent{
			OfferCreationFailedEvent: &OfferCreationFailedEvent,
		},
		ResponseType: offeringPb.CustomerSubscriptionOnOfferResponse_TypeOfferCreationFailedEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func listenToCustomerEvent(ctx context.Context, topic string, eventCha chan pubsub.Event,
	customerId string, stream offeringPb.Offering_CustomerSubscriptionOnOfferServer,
	eventHandler func(pubsub.Event, offeringPb.Offering_CustomerSubscriptionOnOfferServer),
	wg *sync.WaitGroup) {

	logger.Infof("listenToCustomerEvent customerId = %v", customerId)

	defer wg.Done()
	for {
		select {
		case event := <-eventCha:
			logger.Infof("listenToCustomerEvent customerId = %v New Event", customerId)
			eventHandler(event, stream)
		case <-ctx.Done():
			logger.Infof("listenToCustomerEvent customerId = %v Done", customerId)
			pubsub.UnsubscribeById(topic, customerId)
			return
		}
	}
}
