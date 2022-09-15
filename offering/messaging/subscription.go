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
	addSubscription(messaging.TopicDeliveryNewRequest, func(m *nats.Msg) {
		err := onNewRequestCreated(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicDeliveryRequestAccepted, func(m *nats.Msg) {
		err := onRequestAccepted(context.Background(), m)
		if err != nil {
			logger.Error("NOTIFICATION:: failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicRequestRejected, func(m *nats.Msg) {
		err := onRequestRejected(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicRequestCancelled, func(m *nats.Msg) {
		err := onRequestCancelled(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicRetryOfferRequest, func(m *nats.Msg) {
		err := onOfferRetryRequested(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicNewOfferSentToCouriers, func(m *nats.Msg) {
		err := onOfferSentToCouriers(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicOfferMaxRetries, func(m *nats.Msg) {
		err := onOfferMaxRetries(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicOfferCreationFailed, func(m *nats.Msg) {
		err := onOfferCreationFailed(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicAcceptOfferFailed, func(m *nats.Msg) {
		err := onAcceptOfferFailed(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicRejectOfferFailed, func(m *nats.Msg) {
		err := onRejectOfferFailed(context.Background(), m)
		if err != nil {
			logger.Error("failed to handle message", err, tag.Obj("msg", m))
		}
	})

	addSubscription(messaging.TopicCourierStatusUpdate, func(m *nats.Msg) {
		err := onCourierStatusChange(context.Background(), m)
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
