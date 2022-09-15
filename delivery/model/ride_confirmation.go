package model

type RideConfirmation struct {
	RideLocationId string
	RequestId      string
	Name           string
	Signature      string
	Base
}
