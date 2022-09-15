package messaging

import (
	"fmt"
	"testing"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/messaging"
	"github.com/kkjhamb01/courier-management/common/push"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}

func TestSubscription(t *testing.T) {
	//StartSubscriptions()
	/*	subscriptionClient := messaging.NatsClient()
		_, err := subscriptionClient.Subscribe(messaging.TopicPushNotification, func(m *nats.Msg) {
			fmt.Println(m)
		})
		if err != nil {
			panic(err)
		}*/

	evt1Data := &push.RideStateArrived{
		RequesterPhoneNumber: "req",
		SenderPhoneNumber:    "sen",
		ReceiverPhoneNumber:  "rec",
		Ride: &push.RideInfo{
			RequestId:       "rid1",
			LocationId:      "lid1",
			HumanReadableId: "hrid1",
			FirstName:       "f1",
			LastName:        "l1",
			Time:            "time1",
		},
	}
	client := messaging.NatsClient()

	data1, _ := push.Encode(evt1Data)
	err := client.Publish(messaging.TopicPushNotification, data1)
	if err != nil {
		fmt.Printf("error in publish %v", err)
	}

	/*	client.Drain()
		client.Close()*/

	for {
	}
}
