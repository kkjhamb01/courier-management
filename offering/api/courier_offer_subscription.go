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

// some test
func (s serverImpl) CourierSubscriptionOnOffer(_ *empty.Empty, stream offeringPb.Offering_CourierSubscriptionOnOfferServer) error {
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
	courierId := tokenUser.Id

	logger.Infof("CourierSubscriptionOnOffer courierId = %v", courierId)

	isActive, err := business.IsCourierActive(ctx, courierId)
	if err != nil {
		logger.Error("failed to check if the courier is active", err)
		return err
	}
	if !isActive {
		err = fmt.Errorf("the courier %v is not active", courierId)
		logger.Error("failed to subscribe courier", err)
		return err
	}

	newOfferCh := pubsub.SubscribeById(pubsub.TopicNewOffer, courierId)
	acceptOfferCh := pubsub.SubscribeById(pubsub.TopicAcceptOffer, courierId)
	cancelOfferCh := pubsub.SubscribeById(pubsub.TopicCancelOffer, courierId)
	acceptOfferFailedCh := pubsub.SubscribeById(pubsub.TopicAcceptOfferFailed, courierId)
	rejectOfferFailedCh := pubsub.SubscribeById(pubsub.TopicRejectOfferFailed, courierId)

	wg := sync.WaitGroup{}
	wg.Add(5)

	go listenToCourierEvent(ctx, pubsub.TopicNewOffer, newOfferCh, courierId, stream, newCourierOfferEventHandler, &wg)
	go listenToCourierEvent(ctx, pubsub.TopicAcceptOffer, acceptOfferCh, courierId, stream, courierAcceptEventHandler, &wg)
	go listenToCourierEvent(ctx, pubsub.TopicCancelOffer, cancelOfferCh, courierId, stream, courierCancelEventHandler, &wg)
	go listenToCourierEvent(ctx, pubsub.TopicAcceptOfferFailed, acceptOfferFailedCh, courierId, stream, courierAcceptOfferFailedEventHandler, &wg)
	go listenToCourierEvent(ctx, pubsub.TopicRejectOfferFailed, rejectOfferFailedCh, courierId, stream, courierRejectOfferFailedEventHandler, &wg)
	go func(cancelCtx context.Context) {
		for {
			select {
			case <-cancelCtx.Done():
				logger.Infof("closing: listenToCourierEvent courierId = %v Done", courierId)
				pubsub.UnsubscribeById(pubsub.TopicNewOffer, courierId)
				pubsub.UnsubscribeById(pubsub.TopicAcceptOffer, courierId)
				pubsub.UnsubscribeById(pubsub.TopicCancelOffer, courierId)
				pubsub.UnsubscribeById(pubsub.TopicAcceptOfferFailed, courierId)
				pubsub.UnsubscribeById(pubsub.TopicRejectOfferFailed, courierId)
				return
			}
		}
	}(ctx)

	wg.Wait()

	return nil
}

func newCourierOfferEventHandler(event pubsub.Event, stream offeringPb.Offering_CourierSubscriptionOnOfferServer) {
	// get the event type
	newOfferEvent := event.NewOfferEventData()

	logger.Infof("newCourierOfferEventHandler event = %+v", event)

	// create the response
	response := &offeringPb.CourierSubscriptionOnOfferResponse{
		Event: &offeringPb.CourierSubscriptionOnOfferResponse_NewOfferEvent{
			NewOfferEvent: &newOfferEvent,
		},
		ResponseType: offeringPb.CourierSubscriptionOnOfferResponse_TypeNewOfferEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func courierCancelEventHandler(event pubsub.Event, stream offeringPb.Offering_CourierSubscriptionOnOfferServer) {

	logger.Infof("courierCancelEventHandler event = %+v", event)

	// get the event type
	cancelOfferEvent := event.CancelOfferEventData()

	// create the response
	response := &offeringPb.CourierSubscriptionOnOfferResponse{
		Event: &offeringPb.CourierSubscriptionOnOfferResponse_CancelOfferEvent{
			CancelOfferEvent: &cancelOfferEvent,
		},
		ResponseType: offeringPb.CourierSubscriptionOnOfferResponse_TypeOfferCancelledEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func courierAcceptOfferFailedEventHandler(event pubsub.Event, stream offeringPb.Offering_CourierSubscriptionOnOfferServer) {

	logger.Infof("courierAcceptOfferFailedEventHandler event = %+v", event)

	// get the event type
	acceptOfferFailedEvent := event.AcceptOfferFailedEventData()

	// create the response
	response := &offeringPb.CourierSubscriptionOnOfferResponse{
		Event: &offeringPb.CourierSubscriptionOnOfferResponse_AcceptOfferFailedEvent{
			AcceptOfferFailedEvent: &acceptOfferFailedEvent,
		},
		ResponseType: offeringPb.CourierSubscriptionOnOfferResponse_TypeAcceptOfferFailedEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func courierRejectOfferFailedEventHandler(event pubsub.Event, stream offeringPb.Offering_CourierSubscriptionOnOfferServer) {

	logger.Infof("courierRejectOfferFailedEventHandler event = %+v", event)

	// get the event type
	rejectOfferFailedEvent := event.RejectOfferFailedEventData()

	// create the response
	response := &offeringPb.CourierSubscriptionOnOfferResponse{
		Event: &offeringPb.CourierSubscriptionOnOfferResponse_RejectOfferFailedEvent{
			RejectOfferFailedEvent: &rejectOfferFailedEvent,
		},
		ResponseType: offeringPb.CourierSubscriptionOnOfferResponse_TypeRejectOfferFailedEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func courierAcceptEventHandler(event pubsub.Event, stream offeringPb.Offering_CourierSubscriptionOnOfferServer) {

	logger.Infof("courierAcceptEventHandler event = %+v", event)

	// get the event type
	acceptOfferEvent := event.AcceptOfferEventData()

	// create the response
	response := &offeringPb.CourierSubscriptionOnOfferResponse{
		Event: &offeringPb.CourierSubscriptionOnOfferResponse_AcceptOfferEvent{
			AcceptOfferEvent: &acceptOfferEvent,
		},
		ResponseType: offeringPb.CourierSubscriptionOnOfferResponse_TypeOfferAcceptedEvent,
	}

	//send the stream
	err := stream.Send(response)
	if err != nil {
		// ignore error to tolerate failure
		logger.Error("failed to send event", err, tag.Obj("event", event))
	}
}

func listenToCourierEvent(ctx context.Context, topic string, eventCha chan pubsub.Event,
	courierId string, stream offeringPb.Offering_CourierSubscriptionOnOfferServer,
	eventHandler func(pubsub.Event, offeringPb.Offering_CourierSubscriptionOnOfferServer),
	wg *sync.WaitGroup) {

	logger.Infof("listenToCourierEvent courierId = %v", courierId)

	defer wg.Done()
	for {
		select {
		case event := <-eventCha:
			logger.Infof("listenToCourierEvent courierId = %v New Event", courierId)
			eventHandler(event, stream)
		}
	}
}
