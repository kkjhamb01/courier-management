package push

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"gitlab.artin.ai/backend/courier-management/common/logger"
)

func Encode(e proto.Message) ([]byte, error){
	serialized,_ := proto.Marshal(e)
	return encodeEventData(serialized, proto.MessageName(e))
}

func encodeEventData(serialized []byte, name string) ([]byte, error) {
	a := PushEvent{
		Event: &any.Any{
			TypeUrl: name,
			Value:   serialized,
		},
	}
	return proto.Marshal(&a)
}

func Decode(eventByte []byte) (proto.Message, error) {
	logger.Debugf("Decode event %v", string(eventByte))
	var event PushEvent
	err := proto.Unmarshal(eventByte, &event)
	if err != nil{
		logger.Errorf("Decode cannot decode push event", err)
		return nil, err
	}

	switch event.Event.TypeUrl {
	case "push.OfferAccepted":
		result := OfferAccepted{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode OfferAccepted %v", result)
		return &result, err
	case "push.OfferRejected":
		result := OfferRejected{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode OfferRejected %v", result)
		return &result, err
	case "push.RideStateArrived":
		result := RideStateArrived{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideStateArrived %v", result)
		return &result, err
	case "push.RideStateCancelled":
		result := RideStateCancelled{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideStateCancelled %v", result)
		return &result, err
	case "push.RideStatePickedup":
		result := RideStatePickedup{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideStatePickedup %v", result)
		return &result, err
	case "push.RideStateDelivered":
		result := RideStateDelivered{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideStateDelivered %v", result)
		return &result, err
	case "push.RideStateSenderIsNotAvailable":
		result := RideStateSenderIsNotAvailable{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideStateSenderIsNotAvailable %v", result)
		return &result, err
	case "push.RideStateReceiverIsNotAvailable":
		result := RideStateReceiverIsNotAvailable{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideStateReceiverIsNotAvailable %v", result)
		return &result, err
	case "push.RideStateFinished":
		result := RideStateFinished{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideStateFinished %v", result)
		return &result, err
	case "push.RideNewLocation":
		result := RideNewLocation{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode RideNewLocation %v", result)
		return &result, err
	case "push.AnnouncementReceived":
		result := AnnouncementReceived{}
		err = ptypes.UnmarshalAny(event.Event, &result)
		logger.Debugf("Decode AnnouncementReceived %v", result)
		return &result, err
	}

	err = errors.New(fmt.Sprintf("invalid event %v", event.Event.TypeUrl))

	logger.Debugf("Decode cannot decode message")

	return nil, err
}