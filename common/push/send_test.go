package push

import (
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	"testing"
)

func init(){
	config.InitTestConfig()
	logger.InitLogger()
}

func TestSend(t *testing.T) {

	evt1Data := &RideStateArrived{
		RequesterPhoneNumber: "req",
		SenderPhoneNumber: "sen",
		ReceiverPhoneNumber: "rec",
		Ride: &RideInfo{
			RequestId: "rid1",
			LocationId: "lid1",
			HumanReadableId: "hrid1",
			FullName: "f1 l1",
			Time: "time1",
		},
	}
	client := messaging.NatsClient()

	data1,_ := Encode(evt1Data)
	err := client.Publish(messaging.TopicPushNotification, data1)
	if err != nil{
		fmt.Printf("error in publish %v", err)
	}

	for {}
}