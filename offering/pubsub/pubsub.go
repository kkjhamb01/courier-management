package pubsub

import (
	"sync"

	"github.com/kkjhamb01/courier-management/common/logger"
)

var (
	mu       sync.RWMutex
	idSubs   map[string]map[string]chan Event
	typeSubs map[string]map[string]chan Event
)

func init() {
	idSubs = make(map[string]map[string]chan Event)
	typeSubs = make(map[string]map[string]chan Event)
}

func SubscribeById(topic string, id string) chan Event {
	return subscribe(idSubs, topic, id)
}

func SubscribeByType(topic string, tp string) chan Event {
	return subscribe(typeSubs, topic, tp)
}

func subscribe(subs map[string]map[string]chan Event, topic string, key string) chan Event {
	mu.Lock()
	defer mu.Unlock()

	event := make(chan Event, 1)

	if subs[topic] == nil {
		subs[topic] = make(map[string]chan Event, 1000)
	}
	subs[topic][key] = event
	logger.Infof("subscription, subs[topic]: %v, topic: %v, key: %v", subs[topic], topic, key)

	return event
}

func UnsubscribeById(topic string, id string) {
	unsubscribe(idSubs, topic, id)
}

func UnsubscribeByType(topic string, tp string) {
	unsubscribe(typeSubs, topic, tp)
}

func unsubscribe(subs map[string]map[string]chan Event, topic string, key string) {
	mu.Lock()
	defer mu.Unlock()

	delete(subs[topic], key)
	logger.Infof("key %v unsubscribed from topic %v", key, topic)
}

func PublishById(event Event, id string) {
	logger.Infof("PublishById id = %v, event = %+v", id, event)
	publish(idSubs, event, id)
}

func PublishByType(event Event, tp string) {
	publish(typeSubs, event, tp)
}

func publish(subs map[string]map[string]chan Event, event Event, key string) {
	mu.RLock()
	defer mu.RUnlock()

	ch := subs[event.Topic][key]
	if ch != nil {
		ch <- event
	}
	logger.Infof("publish, subs[event.Topic]: %v, ch: %v, key: %v", subs[event.Topic], ch, key)
}
