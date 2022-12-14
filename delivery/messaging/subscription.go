package messaging

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/common/messaging"
	"github.com/nats-io/nats.go"
)

var subscriptions []*nats.Subscription

func StartSubscriptions() {
	addSubscription(messaging.TopicOfferCreationFailed, func(m *nats.Msg) {
		err := onOfferCreationFailed(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})
}

func addSubscription(topic string, msgHandler func(m *nats.Msg)) {
	if subscriptions == nil {
		subscriptions = make([]*nats.Subscription, 0, 10)
	}

	subscription, err := messaging.NatsClient().Subscribe(topic, msgHandler)
	if err != nil {
		logger.Fatal("failed to subscribe on the topic", tag.Str("topic", topic))
		panic("failed to subscribe")
	}
	subscriptions = append(subscriptions, subscription)
}

func StopSubscriptions() {
	for _, subscription := range subscriptions {
		err := subscription.Unsubscribe()
		if err != nil {
			logger.Error("failed to unsubscribe from subscribed topic", err, tag.Obj("subscription", subscription))
		}
	}
}
