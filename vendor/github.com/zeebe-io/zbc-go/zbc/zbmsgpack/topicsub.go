package zbmsgpack

import (
	"encoding/json"
	"fmt"
)

// OpenTopicSubscription is used to open a topic subscription.
type OpenTopicSubscription struct {
	StartPosition    int64  `msgpack:"startPosition"`
	PrefetchCapacity int32  `msgpack:"prefetchCapacity"`
	Name             string `msgpack:"name"`

	ForceStart bool   `msgpack:"forceStart"`
	State      string `msgpack:"state"`
}

func (t *OpenTopicSubscription) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}

// TopicSubscriptionAck is used to acknowledge receiving of an event.
type TopicSubscriptionAck struct {
	Name        string `msgpack:"name"`
	AckPosition uint64 `msgpack:"ackPosition"`
	State       string `msgpack:"state"`
}

func (t *TopicSubscriptionAck) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}

type TopicSubscription struct {
	TopicName     string `msgpack:"topicName"`
	PartitionID   uint16 `msgpack:"partitionId"`
	SubscriberKey uint64 `msgpack:"subscriberKey"`

	SubscriptionName string `msgpack:"-"`
}

func (t *TopicSubscription) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}
