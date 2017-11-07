package zbmsgpack

import (
	"encoding/json"
	"fmt"
)

// TaskSubscription is structure which we use to handle a subscription on a task.
type TaskSubscription struct {
	SubscriberKey uint64 `msgpack:"subscriberKey" json:"subscriberKey"`
	TaskType      string `msgpack:"taskType" json:"taskType"`
	LockDuration  uint64 `msgpack:"lockDuration" json:"lockDuration"`
	LockOwner     string `msgpack:"lockOwner" json:"lockOwner"`
	Credits       int32  `msgpack:"credits" json:"credits"`
}

func (t *TaskSubscription) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}
