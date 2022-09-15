package model

import (
	"time"
)

type AcceptedOffer struct {
	CourierId  string
	CustomerId string
	OfferId    string
	Time       time.Time
	Base
}
