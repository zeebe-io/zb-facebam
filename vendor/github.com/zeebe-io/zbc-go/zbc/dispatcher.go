package zbc

import (
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbsbe"
)

type dispatcher struct {
	transactions  SafeMap
	subscriptions SafeMap
}

func (d *dispatcher) addTransaction(key uint64, value interface{}) {
	d.transactions.Set(fmt.Sprintf("%d", key), value)
}

func (d *dispatcher) dispatchTransaction(key uint64, message *Message) {
	transactionKey := fmt.Sprintf("%d", key)
	if ch, ok := d.transactions.Get(transactionKey); ok {
		requestCh := ch.(chan *Message)
		requestCh <- message
	}
}

func (d *dispatcher) removeTransaction(requestID uint64) {
	d.transactions.Remove(fmt.Sprintf("%d", requestID))
}

func (d *dispatcher) addSubscription(key uint64, value interface{}) {
	d.subscriptions.Set(fmt.Sprintf("%d", key), value)
}

func (d *dispatcher) dispatchTaskEvent(key uint64, message *zbsbe.SubscribedEvent, task *zbmsgpack.Task) {
	subscriberKey := fmt.Sprintf("%d", key)
	if ch, ok := d.subscriptions.Get(subscriberKey); ok {
		requestCh := ch.(chan *SubscriptionEvent)
		requestCh <- &SubscriptionEvent{Task: task, Event: message}
	}
}

func (d *dispatcher) dispatchTopicEvent(key uint64, message *zbsbe.SubscribedEvent) {
	subscriberKey := fmt.Sprintf("%d", key)
	if ch, ok := d.subscriptions.Get(subscriberKey); ok {
		requestCh := ch.(chan *SubscriptionEvent)
		requestCh <- &SubscriptionEvent{Task: nil, Event: message}
	}
}

func (d *dispatcher) removeSubscription(subscriptionKey uint64) {
	d.subscriptions.Remove(fmt.Sprintf("%d", subscriptionKey))
}

func newDispatcher() dispatcher {
	return dispatcher{
		NewSafeMap(),
		NewSafeMap(),
	}
}
