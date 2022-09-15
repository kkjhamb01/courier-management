package status

import (
	"github.com/looplab/fsm"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
)

var requestEvents fsm.Events
var requestCallback fsm.Callbacks

func init() {
	requestEvents = fsm.Events{
		{
			Name: deliveryPb.RequestStatus_PROCESSING.String(),
			Src:  []string{},
			Dst:  deliveryPb.RequestStatus_PROCESSING.String(),
		},
		{
			Name: deliveryPb.RequestStatus_SCHEDULED.String(),
			Src: []string{
				deliveryPb.RequestStatus_PROCESSING.String(),
			},
			Dst: deliveryPb.RequestStatus_SCHEDULED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_CREATION_FAILED.String(),
			Src: []string{
				deliveryPb.RequestStatus_PROCESSING.String(),
			},
			Dst: deliveryPb.RequestStatus_CREATION_FAILED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_ACCEPTED.String(),
			Src: []string{
				deliveryPb.RequestStatus_PROCESSING.String(),
			},
			Dst: deliveryPb.RequestStatus_ACCEPTED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_NAVIGATING.String(),
			Src: []string{
				deliveryPb.RequestStatus_ACCEPTED.String(),
				deliveryPb.RequestStatus_DELIVERED.String(),
				deliveryPb.RequestStatus_PICKED_UP.String(),
			},
			Dst: deliveryPb.RequestStatus_NAVIGATING.String(),
		},
		{
			Name: deliveryPb.RequestStatus_COURIER_CANCELLED.String(),
			Src: []string{
				deliveryPb.RequestStatus_PROCESSING.String(),
				deliveryPb.RequestStatus_SCHEDULED.String(),
				deliveryPb.RequestStatus_ACCEPTED.String(),
				deliveryPb.RequestStatus_NAVIGATING.String(),
				deliveryPb.RequestStatus_ARRIVED.String(),
				deliveryPb.RequestStatus_SENDER_NOT_AVAILABLE.String(),
				deliveryPb.RequestStatus_RECEIVER_NOT_AVAILABLE.String(),
			},
			Dst: deliveryPb.RequestStatus_COURIER_CANCELLED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_CUSTOMER_CANCELLED.String(),
			Src: []string{
				deliveryPb.RequestStatus_PROCESSING.String(),
				deliveryPb.RequestStatus_SCHEDULED.String(),
				deliveryPb.RequestStatus_ACCEPTED.String(),
				deliveryPb.RequestStatus_NAVIGATING.String(),
				deliveryPb.RequestStatus_ARRIVED.String(),
				deliveryPb.RequestStatus_SENDER_NOT_AVAILABLE.String(),
				deliveryPb.RequestStatus_RECEIVER_NOT_AVAILABLE.String(),
			},
			Dst: deliveryPb.RequestStatus_CUSTOMER_CANCELLED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_ARRIVED.String(),
			Src: []string{
				deliveryPb.RequestStatus_NAVIGATING.String(),
				deliveryPb.RequestStatus_ARRIVED.String(),
			},
			Dst: deliveryPb.RequestStatus_CUSTOMER_CANCELLED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_PICKED_UP.String(),
			Src: []string{
				deliveryPb.RequestStatus_ARRIVED.String(),
				deliveryPb.RequestStatus_SENDER_NOT_AVAILABLE.String(),
			},
			Dst: deliveryPb.RequestStatus_PICKED_UP.String(),
		},
		{
			Name: deliveryPb.RequestStatus_DELIVERED.String(),
			Src: []string{
				deliveryPb.RequestStatus_ARRIVED.String(),
				deliveryPb.RequestStatus_RECEIVER_NOT_AVAILABLE.String(),
			},
			Dst: deliveryPb.RequestStatus_DELIVERED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_SENDER_NOT_AVAILABLE.String(),
			Src: []string{
				deliveryPb.RequestStatus_ARRIVED.String(),
			},
			Dst: deliveryPb.RequestStatus_SENDER_NOT_AVAILABLE.String(),
		},
		{
			Name: deliveryPb.RequestStatus_RECEIVER_NOT_AVAILABLE.String(),
			Src: []string{
				deliveryPb.RequestStatus_ARRIVED.String(),
			},
			Dst: deliveryPb.RequestStatus_RECEIVER_NOT_AVAILABLE.String(),
		},
		{
			Name: deliveryPb.RequestStatus_COMPLETED.String(),
			Src: []string{
				deliveryPb.RequestStatus_DELIVERED.String(),
			},
			Dst: deliveryPb.RequestStatus_COMPLETED.String(),
		},
		{
			Name: deliveryPb.RequestStatus_NO_DRIVERS_AVAILABLE.String(),
			Src: []string{
				deliveryPb.RequestStatus_PROCESSING.String(),
			},
			Dst: deliveryPb.RequestStatus_NO_DRIVERS_AVAILABLE.String(),
		},
	}

	//requestCallback = fsm.Callbacks{
	//	After + deliveryPb.RequestStatus_COURIER_CANCELLED.String(): func(e *fsm.Event) { /* TODO: required actions on state transition to courier_cancelled */ },
	//}
}

// RequestTransition returns true if transition from status statusA to statusB is legal
// and returns false if it is a illegal transition
func RequestTransition(statusA, statusB string) bool {
	requestTransition := fsm.NewFSM(
		statusA,
		requestEvents,
		requestCallback,
	)

	err := requestTransition.Event(statusB)
	return err == nil
}
