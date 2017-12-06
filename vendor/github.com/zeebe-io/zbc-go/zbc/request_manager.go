package zbc

import (
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbsbe"
)

type requestWrapper struct {
	addr string
	sock       *socket
	responseCh chan *Message
	errorCh    chan error
	payload    *Message
}

func newRequestWrapper(payload *Message) *requestWrapper {
	return &requestWrapper{
		"",
		nil,
		make(chan *Message),
		make(chan error),
		payload,
	}
}

type requestManager struct {
	*requestFactory
	*responseHandler

	*topologyManager
}

func (rm *requestManager) createTask(topic string, task *zbmsgpack.Task) (*zbmsgpack.Task, error) {
	partitionID, err := rm.partitionID(topic)

	if err != nil {
		return nil, err
	}

	message := rm.createTaskRequest(partitionID, 0, task)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)
	if err != nil {
		return nil, err
	}
	return rm.unmarshalTask(resp), nil
}

func (rm *requestManager) createWorkflow(topic string, resource []*zbmsgpack.Resource) (*zbmsgpack.Workflow, error) {
	message := rm.deployWorkflowRequest(topic, resource)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)
	if err != nil {
		return nil, err
	}
	return rm.unmarshalWorkflow(resp), nil
}

func (rm *requestManager) createWorkflowInstance(topic string, wfi *zbmsgpack.WorkflowInstance) (*zbmsgpack.WorkflowInstance, error) {
	partitionID, err := rm.partitionID(topic)

	if err != nil {
		return nil, err
	}

	message := rm.createWorkflowInstanceRequest(partitionID, 0, topic, wfi)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)
	if err != nil {
		return nil, err
	}
	return rm.unmarshalWorkflowInstance(resp), nil
}

func (rm *requestManager) completeTask(task *SubscriptionEvent) (*zbmsgpack.Task, error) {
	message := rm.completeTaskRequest(task)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)
	if err != nil {
		return nil, err
	}
	return rm.unmarshalTask(resp), nil
}

func (rm *requestManager) taskConsumer(topic, lockOwner, taskType string, credits int32) (chan *SubscriptionEvent, *zbmsgpack.TaskSubscription, error) {
	partitionID, err := rm.partitionID(topic)
	if err != nil {
		return nil, nil, err
	}

	subscriptionCh := make(chan *SubscriptionEvent, credits)
	message := rm.openTaskSubscriptionRequest(partitionID, lockOwner, taskType, credits)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)

	if err != nil {
		return nil, nil, err
	}

	taskSubInfo := rm.unmarshalTaskSubscription(resp)
	if taskSubInfo == nil {
		// TODO:
	}

	request.sock.addSubscription(taskSubInfo.SubscriberKey, subscriptionCh)
	return subscriptionCh, taskSubInfo, nil
}

func (rm *requestManager) increaseTaskSubscriptionCredits(task *zbmsgpack.TaskSubscription) (*zbmsgpack.TaskSubscription, error) {
	message := rm.increaseTaskSubscriptionCreditsRequest(task)

	request := newRequestWrapper(message)

	resp, err := rm.executeRequest(request)
	if err != nil {
		return nil, err
	}

	return rm.unmarshalTaskSubscription(resp), nil
}

func (rm *requestManager) closeTaskSubscription(task *zbmsgpack.TaskSubscription) (*Message, error) {
	message := rm.closeTaskSubscriptionRequest(task)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)
	return resp, err
}

func (rm *requestManager) closeTopicSubscription(task *zbmsgpack.TopicSubscription) (*Message, error) {
	message := rm.closeTopicSubscriptionRequest(task)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)
	return resp, err
}

func (rm *requestManager) topicSubscriptionAck(ts *zbmsgpack.TopicSubscription, s *SubscriptionEvent) (*zbmsgpack.TopicSubscriptionAck, error) {
	message := rm.topicSubscriptionAckRequest(ts, s)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)
	return rm.unmarshalTopicSubAck(resp), err
}

func (rm *requestManager) createTopic(name string, partitionNum int) (*zbmsgpack.Topic, error) {
	topic := zbmsgpack.NewTopic(name, TopicCreate, partitionNum)
	message := rm.createTopicRequest(topic)
	request := newRequestWrapper(message)

	resp, err := rm.executeRequest(request)
	if err != nil {
		return nil, err
	}
	return rm.unmarshalTopic(resp), nil
}

func (rm *requestManager) topicConsumer(topic, subName string, startPosition int64) (chan *SubscriptionEvent, *zbmsgpack.TopicSubscription, error) {
	partitionID, err := rm.partitionID(topic)
	if err != nil {
		return nil, nil, err
	}

	subscriptionCh := make(chan *SubscriptionEvent, 1000)
	message := rm.openTopicSubscriptionRequest(partitionID, topic, subName, startPosition)
	request := newRequestWrapper(message)
	resp, err := rm.executeRequest(request)

	cmdResponse := (*resp.SbeMessage).(*zbsbe.ExecuteCommandResponse)
	subscriberKey := cmdResponse.Key

	request.sock.addSubscription(cmdResponse.Key, subscriptionCh)

	subscriptionInfo := &zbmsgpack.TopicSubscription{
		TopicName:        topic,
		PartitionID:      partitionID,
		SubscriberKey:    subscriberKey,
		SubscriptionName: subName,
	}

	return subscriptionCh, subscriptionInfo, nil
}

func newRequestManager(bootstrapAddr string) *requestManager {
	return &requestManager{
		newRequestFactory(),
		newResponseHandler(),
		newTopologyManager(bootstrapAddr),
	}
}
