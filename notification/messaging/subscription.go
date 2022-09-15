package messaging

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/common/messaging"
	"gitlab.artin.ai/backend/courier-management/common/push"
)

var subscriptions []*nats.Subscription

func StartSubscriptions() {
	service := NewPushClient(config.Notification())
	addSubscription(messaging.TopicPushNotification, func(m *nats.Msg) {
		logger.Debugf("msg : %v", m)
		onNewPushEvent(context.Background(), service, m)
	})

}

func addSubscription(topic string, msgHandler func(m *nats.Msg)) {
	if subscriptions == nil {
		subscriptions = make([]*nats.Subscription, 0, 10)
	}

	client := messaging.NatsClient()
	subscription, err := client.Subscribe(topic, msgHandler)
	if err != nil {
		logger.Fatal("failed to subscribe on the topic", tag.Str("topic", topic))
		panic("failed to subscribe")
	}

	subscriptions = append(subscriptions, subscription)
}

func onNewPushEvent(ctx context.Context, client *PushClient, msg *nats.Msg) {
	logger.Debugf("onNewPushEvent")
	if msg.Data == nil {
		err := errors.New("msg.Data is nil")
		logger.Error("no data is supplied on new offer event", err, tag.Obj("msg", msg))
		return
	}

	newPush, err := push.Decode(msg.Data)
	if err != nil {
		logger.Error("failed to decode data", err)
		return
	}

	if templater, ok := newPush.(push.Templater); ok{
		go client.OnNewPushEvent(ctx, templater)
	} else {
		logger.Debugf("event is not templater %v", newPush.String())
	}

}

func StopSubscriptions() {
	for _, subscription := range subscriptions {
		err := subscription.Unsubscribe()
		if err != nil {
			logger.Error("failed to unsubscribe from subscribed topic", err, tag.Obj("subscription", subscription))
		}
	}
}
