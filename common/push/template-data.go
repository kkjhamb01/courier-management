package push

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
)


func (e *OfferAccepted) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "OfferAccepted"},
			},
			"Id": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetOfferId()},
			},
			"CourierId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetCourierId()},
			},
			"CourierName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetCourierName()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetTime()},
			},
			"CustomerPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetCustomerPhoneNumber()},
			},
		},
	}
}

func (e *OfferRejected) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "OfferRejected"},
			},
			"Id": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetOfferId()},
			},
			"CourierId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetCourierId()},
			},
			"CourierName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetCourierName()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetOffer().GetTime()},
			},
			"CustomerPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetCustomerPhoneNumber()},
			},
		},
	}
}

func (e *RideStateArrived) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStateArrived"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"RequesterPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRequesterPhoneNumber()},
			},
			"SenderPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetSenderPhoneNumber()},
			},
			"ReceiverPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetReceiverPhoneNumber()},
			},
		},
	}
}

func (e *RideStateCancelled) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStateCancelled"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"RequesterPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRequesterPhoneNumber()},
			},
			"CancellationNote": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetCancellationNote()},
			},
		},
	}
}

func (e *RideStatePickedup) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStatePickedup"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"RequesterPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRequesterPhoneNumber()},
			},
			"SenderPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetSenderPhoneNumber()},
			},
		},
	}
}

func (e *RideStateDelivered) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStateDelivered"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"RequesterPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRequesterPhoneNumber()},
			},
			"ReceiverPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetReceiverPhoneNumber()},
			},
		},
	}
}

func (e *RideStateSenderIsNotAvailable) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStateSenderIsNotAvailable"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"RequesterPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRequesterPhoneNumber()},
			},
			"SenderPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetSenderPhoneNumber()},
			},
		},
	}
}

func (e *RideStateReceiverIsNotAvailable) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStateReceiverIsNotAvailable"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"RequesterPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRequesterPhoneNumber()},
			},
			"ReceiverPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetReceiverPhoneNumber()},
			},
		},
	}
}

func (e *RideStateFinished) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStateFinished"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"RequesterPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRequesterPhoneNumber()},
			},
		},
	}
}

func (e *RideNewLocation) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "RideStateFinished"},
			},
			"RequestId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetRequestId()},
			},
			"LocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetLocationId()},
			},
			"HumanReadableId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetHumanReadableId()},
			},
			"FullName": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetFullName()},
			},
			"PhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetPhoneNumber()},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetRide().GetTime()},
			},
			"CourierPhoneNumber": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetCourierPhoneNumber()},
			},
			"NewLocation": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetLocation()},
			},
			"NewLocationId": {
				Kind: &structpb.Value_StringValue{StringValue: e.GetLocationId()},
			},
		},
	}
}

func (e *AnnouncementReceived) GetData() *structpb.Struct{
	return &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"NotificationType": {
				Kind: &structpb.Value_StringValue{StringValue: "AnnouncementReceived"},
			},
			"Id": {
				Kind: &structpb.Value_NumberValue{NumberValue: float64(e.Id)},
			},
			"Title": {
				Kind: &structpb.Value_StringValue{StringValue: e.Title},
			},
			"Type": {
				Kind: &structpb.Value_NumberValue{NumberValue: float64(e.Type)},
			},
			"MessageType": {
				Kind: &structpb.Value_NumberValue{NumberValue: float64(e.MessageType)},
			},
			"Text": {
				Kind: &structpb.Value_StringValue{StringValue: e.Text},
			},
			"Time": {
				Kind: &structpb.Value_StringValue{StringValue: e.Time},
			},
		},
	}
}
