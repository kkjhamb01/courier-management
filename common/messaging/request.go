package messaging

import deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"

func EncodeDeliveryRequestRejectedData(event deliveryPb.RequestRejectedEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeDeliveryRequestRejectedData(eventByte []byte) (deliveryPb.RequestRejectedEvent, error) {
	var event deliveryPb.RequestRejectedEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeDeliveryRequestCancelledData(event deliveryPb.RequestCancelledEvent) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeDeliveryRequestCancelledData(eventByte []byte) (deliveryPb.RequestCancelledEvent, error) {
	var event deliveryPb.RequestCancelledEvent
	var err = decodeProto(eventByte, &event)
	return event, err
}

func EncodeDeliveryNewRequestData(event deliveryPb.Request) ([]byte, error) {
	return encodeProto(&event)
}

func DecodeDeliveryNewRequestData(eventByte []byte) (deliveryPb.Request, error) {
	var event deliveryPb.Request
	var err = decodeProto(eventByte, &event)
	return event, err
}
