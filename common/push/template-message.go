package push

import "fmt"

func (e *OfferAccepted) GetMessage() string{
	return fmt.Sprintf("%v accepted your delivery request!", e.GetOffer().GetCourierName())
}

func (e *OfferRejected) GetMessage() string{
	return fmt.Sprintf("%v rejected your delivery request!", e.GetOffer().GetCourierName())
}

func (e *RideStateArrived) GetMessage() string{
	return "The courier is waiting to receive your package at the pickup location."
}

func (e *RideStateCancelled) GetMessage() string{
	return "Your request has been canceled by the courier."
}

func (e *RideStatePickedup) GetMessage() string{
	return "The courier has picked up your package."
}

func (e *RideStateDelivered) GetMessage() string{
	return fmt.Sprintf("Your package has been dropped off by %v (%v).",
		e.GetRide().GetFullName(),
		e.GetRide().GetPhoneNumber())
}

func (e *RideStateSenderIsNotAvailable) GetMessage() string{
	return "The courier can not find anyone at pickup point."
}

func (e *RideStateReceiverIsNotAvailable) GetMessage() string{
	return "The courier can not find anyone at drop off point."
}

func (e *RideStateFinished) GetMessage() string{
	return "Your request has been finished."
}

func (e *RideNewLocation) GetMessage() string{
	return fmt.Sprintf("Location %v is added to the ride.", e.Location)
}

func (e *AnnouncementReceived) GetMessage() string{
	return e.GetText()
}