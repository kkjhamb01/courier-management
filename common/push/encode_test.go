package push

import (
	"fmt"
	"testing"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}

func TestEncode(t *testing.T) {
	evt1Data := &RideStateArrived{
		RequesterPhoneNumber: "req",
		SenderPhoneNumber:    "sen",
		ReceiverPhoneNumber:  "rec",
		Ride: &RideInfo{
			RequestId:       "rid1",
			LocationId:      "lid1",
			HumanReadableId: "hrid1",
			FullName:        "f1 l1",
			Time:            "time1",
		},
	}
	serializedA, _ := Encode(evt1Data)

	p, _ := Decode(serializedA)
	fmt.Printf("p %v", p.String())

}
