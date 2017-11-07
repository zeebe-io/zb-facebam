package zbc

import (
	"github.com/vmihailenco/msgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
	"time"
)

type responseHandler struct{}

func (rf *responseHandler) unmarshalTopology(msg *Message) *zbmsgpack.ClusterTopology {
	var resp zbmsgpack.ClusterTopologyResponse
	msgpack.Unmarshal(msg.Data, &resp)

	ct := &zbmsgpack.ClusterTopology{
		Brokers:      resp.Brokers,
		TopicLeaders: make(map[string][]zbmsgpack.TopicLeader),
		UpdatedAt:    time.Now(),
	}

	for _, leader := range resp.TopicLeaders {
		ct.TopicLeaders[leader.TopicName] = append(ct.TopicLeaders[leader.TopicName], leader)
	}
	return ct
}

func (rf *responseHandler) unmarshalTopic(msg *Message) *zbmsgpack.Topic {
	var topic zbmsgpack.Topic
	msgpack.Unmarshal(msg.Data, &topic)

	return &topic
}

func (rf *responseHandler) unmarshalTask(m *Message) *zbmsgpack.Task {
	var d zbmsgpack.Task
	err := msgpack.Unmarshal(m.Data, &d)
	if err != nil {
		return nil
	}
	if len(d.Type) > 0 {
		return &d
	}
	return nil
}

func (rf *responseHandler) unmarshalTaskSubscription(m *Message) *zbmsgpack.TaskSubscription {
	var d zbmsgpack.TaskSubscription
	err := msgpack.Unmarshal(m.Data, &d)
	if err != nil {
		return nil
	}
	return &d
}

func (rf *responseHandler) unmarshalWorkflow(m *Message) *zbmsgpack.Workflow {
	var d zbmsgpack.Workflow
	err := msgpack.Unmarshal(m.Data, &d)
	if err != nil {
		return nil
	}
	if len(d.State) > 0 && len(d.ResourceType) > 0 {
		return &d
	}
	return nil
}

func (rf *responseHandler) unmarshalWorkflowInstance(m *Message) *zbmsgpack.WorkflowInstance {
	var d zbmsgpack.WorkflowInstance
	err := msgpack.Unmarshal(m.Data, &d)
	if err != nil {
		return nil
	}
	if len(d.BPMNProcessID) > 0 {
		return &d
	}
	return nil
}

func (rf *responseHandler) unmarshalTopicSubAck(m *Message) *zbmsgpack.TopicSubscriptionAck {
	var d zbmsgpack.TopicSubscriptionAck
	err := msgpack.Unmarshal(m.Data, &d)
	if err != nil {
		return nil
	}
	if d.State == TopicSubscriptionAcknowledgedState {
		return &d
	}
	return nil
}
