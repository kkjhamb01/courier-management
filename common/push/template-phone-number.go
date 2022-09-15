package push

func (e *OfferAccepted) GetPhoneNumbers() []string{
	return []string{e.GetCustomerPhoneNumber()}
}

func (e *OfferRejected) GetPhoneNumbers() []string{
	return []string{e.GetCustomerPhoneNumber()}
}

func (e *RideStateArrived) GetPhoneNumbers() []string{
	return []string{e.GetRequesterPhoneNumber(), e.GetSenderPhoneNumber(), e.GetReceiverPhoneNumber()}
}

func (e *RideStateCancelled) GetPhoneNumbers() []string{
	return []string{e.GetRequesterPhoneNumber()}
}

func (e *RideStatePickedup) GetPhoneNumbers() []string{
	return []string{e.GetRequesterPhoneNumber(), e.GetSenderPhoneNumber()}
}

func (e *RideStateDelivered) GetPhoneNumbers() []string{
	return []string{e.GetRequesterPhoneNumber(), e.GetReceiverPhoneNumber()}
}

func (e *RideStateSenderIsNotAvailable) GetPhoneNumbers() []string{
	return []string{e.GetRequesterPhoneNumber(), e.GetSenderPhoneNumber()}
}

func (e *RideStateReceiverIsNotAvailable) GetPhoneNumbers() []string{
	return []string{e.GetRequesterPhoneNumber(), e.GetReceiverPhoneNumber()}
}

func (e *RideStateFinished) GetPhoneNumbers() []string{
	return []string{e.GetRequesterPhoneNumber()}
}

func (e *RideNewLocation) GetPhoneNumbers() []string{
	return []string{e.GetCourierPhoneNumber()}
}

func (e *AnnouncementReceived) GetPhoneNumbers() []string{
	return []string{e.GetUserPhoneNumber()}
}