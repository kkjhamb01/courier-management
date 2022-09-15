package push

func (e *OfferAccepted) GetTitle() string{
	return "Request accepted"
}

func (e *OfferRejected) GetTitle() string{
	return "Request rejected"
}

func (e *RideStateArrived) GetTitle() string{
	return "Courier arrived"
}

func (e *RideStateCancelled) GetTitle() string{
	return "Delivery canceled"
}

func (e *RideStatePickedup) GetTitle() string{
	return "Pickup confirmed"
}

func (e *RideStateDelivered) GetTitle() string{
	return "Package dropped off"
}

func (e *RideStateSenderIsNotAvailable) GetTitle() string{
	return "No one is at the pickup off point"
}

func (e *RideStateReceiverIsNotAvailable) GetTitle() string{
	return "No one is at the drop off point"
}

func (e *RideStateFinished) GetTitle() string{
	return "Delivery finished"
}

func (e *RideNewLocation) GetTitle() string{
	return "New location to the ride"
}
