package pubsub

import (
	"gitlab.artin.ai/backend/courier-management/grpc/offering/go"
)

type Event struct {
	Data  interface{}
	Topic string
}

func (e *Event) NewOfferEventData() offeringPb.NewOfferEvent {
	return e.Data.(offeringPb.NewOfferEvent)
}

func (e *Event) AcceptOfferEventData() offeringPb.OfferAcceptedEvent {
	return e.Data.(offeringPb.OfferAcceptedEvent)
}

func (e *Event) CancelOfferEventData() offeringPb.OfferCancelledEvent {
	return e.Data.(offeringPb.OfferCancelledEvent)
}

func (e *Event) RetryOfferEventData() offeringPb.RetryOfferRequestEvent {
	return e.Data.(offeringPb.RetryOfferRequestEvent)
}

func (e *Event) MaxOfferRetriesEventData() offeringPb.MaxOfferRetriesEvent {
	return e.Data.(offeringPb.MaxOfferRetriesEvent)
}

func (e *Event) OfferSentToCouriersEventData() offeringPb.NewOfferSentToCouriersEvent {
	return e.Data.(offeringPb.NewOfferSentToCouriersEvent)
}

func (e *Event) OfferCreationFailedEventData() offeringPb.OfferCreationFailedEvent {
	return e.Data.(offeringPb.OfferCreationFailedEvent)
}

func (e *Event) AcceptOfferFailedEventData() offeringPb.AcceptOfferFailedEvent {
	return e.Data.(offeringPb.AcceptOfferFailedEvent)
}

func (e *Event) RejectOfferFailedEventData() offeringPb.RejectOfferFailedEvent {
	return e.Data.(offeringPb.RejectOfferFailedEvent)
}

func NewOfferEvent(newOfferEvent offeringPb.NewOfferEvent) Event {
	return Event{
		Data:  newOfferEvent,
		Topic: TopicNewOffer,
	}
}

func AcceptOfferEvent(acceptOfferEvent offeringPb.OfferAcceptedEvent) Event {
	return Event{
		Data:  acceptOfferEvent,
		Topic: TopicAcceptOffer,
	}
}

func CancelOfferEvent(cancelOfferEvent offeringPb.OfferCancelledEvent) Event {
	return Event{
		Data:  cancelOfferEvent,
		Topic: TopicCancelOffer,
	}
}

func RetryOfferEventData(retryOfferEvent offeringPb.RetryOfferRequestEvent) Event {
	return Event{
		Data:  retryOfferEvent,
		Topic: TopicRetryOffer,
	}
}

func MaxOfferRetiesEventData(maxOfferRetiesEvent offeringPb.MaxOfferRetriesEvent) Event {
	return Event{
		Data:  maxOfferRetiesEvent,
		Topic: TopicMaxOfferRetries,
	}
}

func OfferSentToCouriersEventData(offerSentToCouriersEvent offeringPb.NewOfferSentToCouriersEvent) Event {
	return Event{
		Data:  offerSentToCouriersEvent,
		Topic: TopicOfferSentToCouriers,
	}
}

func OfferCreationFailedEventData(offerCreationFailedEvent offeringPb.OfferCreationFailedEvent) Event {
	return Event{
		Data:  offerCreationFailedEvent,
		Topic: TopicOfferCreationFailed,
	}
}

func AcceptOfferFailedEventData(acceptOfferFailedEvent offeringPb.AcceptOfferFailedEvent) Event {
	return Event{
		Data:  acceptOfferFailedEvent,
		Topic: TopicAcceptOfferFailed,
	}
}

func RejectOfferFailedEventData(rejectOfferFailedEvent offeringPb.RejectOfferFailedEvent) Event {
	return Event{
		Data:  rejectOfferFailedEvent,
		Topic: TopicRejectOfferFailed,
	}
}
