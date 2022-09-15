package push

func (e *OfferAccepted) GetCategory() string{
	return "Accepted"
}

func (e *OfferRejected) GetCategory() string{
	return "Rejected"
}

func (e *RideStateArrived) GetCategory() string{
	return "Arrived"
}

func (e *RideStateCancelled) GetCategory() string{
	return "Canceled"
}

func (e *RideStatePickedup) GetCategory() string{
	return "Picked up"
}

func (e *RideStateDelivered) GetCategory() string{
	return "Dropped off"
}

func (e *RideStateSenderIsNotAvailable) GetCategory() string{
	return "No one is here"
}

func (e *RideStateReceiverIsNotAvailable) GetCategory() string{
	return "No one is here"
}

func (e *RideStateFinished) GetCategory() string{
	return "Finished"
}

func (e *RideNewLocation) GetCategory() string{
	return "New location"
}

func (e *AnnouncementReceived) GetCategory() string{
	return "New announcement"
}