package messaging

import (
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
)

func EncodeMaxOfferRetriesData(event offeringPb.MaxOfferRetriesEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeMaxOfferRetriesData(eventByte []byte) (offeringPb.MaxOfferRetriesEvent, error) {
	var event offeringPb.MaxOfferRetriesEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRetryOfferRequestData(event offeringPb.RetryOfferRequestEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRetryOfferRequestData(eventByte []byte) (offeringPb.RetryOfferRequestEvent, error) {
	var event offeringPb.RetryOfferRequestEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeNewOfferSentToCouriersData(event offeringPb.NewOfferSentToCouriersEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeNewOfferSentToCouriersData(eventByte []byte) (offeringPb.NewOfferSentToCouriersEvent, error) {
	var event offeringPb.NewOfferSentToCouriersEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeOfferCreationFailedData(event offeringPb.OfferCreationFailedEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeOfferCreationFailedData(eventByte []byte) (offeringPb.OfferCreationFailedEvent, error) {
	var event offeringPb.OfferCreationFailedEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeAcceptOfferFailedData(event offeringPb.AcceptOfferFailedEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeAcceptOfferFailedData(eventByte []byte) (offeringPb.AcceptOfferFailedEvent, error) {
	var event offeringPb.AcceptOfferFailedEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRejectOfferFailedData(event offeringPb.RejectOfferFailedEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRejectOfferFailedData(eventByte []byte) (offeringPb.RejectOfferFailedEvent, error) {
	var event offeringPb.RejectOfferFailedEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestAcceptedData(event deliveryPb.RequestAcceptedEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestAcceptedData(eventByte []byte) (deliveryPb.RequestAcceptedEvent, error) {
	var event deliveryPb.RequestAcceptedEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestArrivedOriginData(event deliveryPb.RequestArrivedOriginEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestArrivedOriginData(eventByte []byte) (deliveryPb.RequestArrivedOriginEvent, error) {
	var event deliveryPb.RequestArrivedOriginEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestArrivedDestinationData(event deliveryPb.RequestArrivedDestinationEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestArrivedDestinationData(eventByte []byte) (deliveryPb.RequestArrivedDestinationEvent, error) {
	var event deliveryPb.RequestArrivedDestinationEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestPickedUpData(event deliveryPb.RequestPickedUpEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestPickedUpData(eventByte []byte) (deliveryPb.RequestPickedUpEvent, error) {
	var event deliveryPb.RequestPickedUpEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestDeliveredData(event deliveryPb.RequestDeliveredEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestDeliveredData(eventByte []byte) (deliveryPb.RequestDeliveredEvent, error) {
	var event deliveryPb.RequestDeliveredEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestCompletedData(event deliveryPb.RequestCompletedEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestCompletedData(eventByte []byte) (deliveryPb.RequestCompletedEvent, error) {
	var event deliveryPb.RequestCompletedEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestNavigatingToReceiverData(event deliveryPb.RequestNavigatingToReceiverEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestNavigatingToReceiverData(eventByte []byte) (deliveryPb.RequestNavigatingToReceiverEvent, error) {
	var event deliveryPb.RequestNavigatingToReceiverEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestNavigatingToSenderData(event deliveryPb.RequestNavigatingToSenderEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestNavigatingToSenderData(eventByte []byte) (deliveryPb.RequestNavigatingToSenderEvent, error) {
	var event deliveryPb.RequestNavigatingToSenderEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestSenderNotAvailable(event deliveryPb.RequestSenderNotAvailableEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestSenderNotAvailable(eventByte []byte) (deliveryPb.RequestSenderNotAvailableEvent, error) {
	var event deliveryPb.RequestSenderNotAvailableEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeRequestReceiverNotAvailable(event deliveryPb.RequestReceiverNotAvailableEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeRequestReceiverNotAvailable(eventByte []byte) (deliveryPb.RequestReceiverNotAvailableEvent, error) {
	var event deliveryPb.RequestReceiverNotAvailableEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}
