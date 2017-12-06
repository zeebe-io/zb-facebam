package zbc

import (
	"github.com/vmihailenco/msgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbprotocol"
	"github.com/zeebe-io/zbc-go/zbc/zbsbe"
)

type requestFactory struct{}

func (rf *requestFactory) headers(t interface{}) *Headers {
	switch v := t.(type) {

	case *zbsbe.ExecuteCommandRequest:
		length := uint32(LengthFieldSize + len(v.Command))
		length += uint32(v.SbeBlockLength()) + TotalHeaderSize //TotalHeaderSizeNoFrame + 12

		var headers Headers
		headers.SetSbeMessageHeader(&zbsbe.MessageHeader{
			BlockLength: v.SbeBlockLength(),
			TemplateId:  v.SbeTemplateId(),
			SchemaId:    v.SbeSchemaId(),
			Version:     v.SbeSchemaVersion(),
		})

		headers.SetRequestResponseHeader(zbprotocol.NewRequestResponseHeader())
		headers.SetTransportHeader(zbprotocol.NewTransportHeader(zbprotocol.RequestResponse))

		headers.SetFrameHeader(zbprotocol.NewFrameHeader(uint32(length), 0, 0, 0, 0))
		return &headers

	case *zbsbe.ControlMessageRequest:
		length := uint32(LengthFieldSize + len(v.Data))
		length += uint32(v.SbeBlockLength()) + TotalHeaderSize //TotalHeaderSizeNoFrame + 12

		var headers Headers
		headers.SetSbeMessageHeader(&zbsbe.MessageHeader{
			BlockLength: v.SbeBlockLength(),
			TemplateId:  v.SbeTemplateId(),
			SchemaId:    v.SbeSchemaId(),
			Version:     v.SbeSchemaVersion(),
		})
		headers.SetRequestResponseHeader(zbprotocol.NewRequestResponseHeader())
		headers.SetTransportHeader(zbprotocol.NewTransportHeader(zbprotocol.RequestResponse))

		// Writer will set FrameHeader after serialization to byte array.
		headers.SetFrameHeader(zbprotocol.NewFrameHeader(uint32(length), 0, 0, 0, 0))
		return &headers
	}

	return nil
}

func (rf *requestFactory) newCommandMessage(commandRequest *zbsbe.ExecuteCommandRequest, command interface{}) *Message {
	var msg Message

	b, err := msgpack.Marshal(command)
	if err != nil {
		return nil
	}
	commandRequest.Command = b
	msg.SetSbeMessage(commandRequest)
	msg.SetHeaders(rf.headers(commandRequest))

	return &msg
}

func (rf *requestFactory) newControlMessage(req *zbsbe.ControlMessageRequest, payload interface{}) *Message {
	var msg Message

	b, err := msgpack.Marshal(payload)
	if err != nil {
		return nil
	}
	req.Data = b
	msg.SetSbeMessage(req)
	msg.SetHeaders(rf.headers(req))

	return &msg
}

func (rf *requestFactory) createTaskRequest(partition uint16, position uint64, task *zbmsgpack.Task) *Message {
	commandRequest := &zbsbe.ExecuteCommandRequest{
		PartitionId: partition,
		Position:    position,
		Command:     []uint8{},
	}
	commandRequest.Key = commandRequest.KeyNullValue()

	commandRequest.EventType = zbsbe.EventTypeEnum(0)

	if task.Payload == nil {
		b, err := msgpack.Marshal(task.PayloadJSON)
		if err != nil {
			return nil
		}
		task.Payload = b
	}

	return rf.newCommandMessage(commandRequest, task)
}

func (rf *requestFactory) completeTaskRequest(taskMessage *SubscriptionEvent) *Message {
	taskMessage.Task.State = TaskComplete
	cmdReq := &zbsbe.ExecuteCommandRequest{
		PartitionId: taskMessage.Event.PartitionId,
		Position:    taskMessage.Event.Position,
		Key:         taskMessage.Event.Key,
	}
	return rf.newCommandMessage(cmdReq, taskMessage.Task)
}

func (rf *requestFactory) createWorkflowInstanceRequest(partition uint16, position uint64, topic string, wf *zbmsgpack.WorkflowInstance) *Message {
	commandRequest := &zbsbe.ExecuteCommandRequest{
		PartitionId: partition,
		Position:    position,
		Command:     []uint8{},
	}

	commandRequest.Key = commandRequest.KeyNullValue()
	commandRequest.EventType = zbsbe.EventTypeEnum(5)

	if wf.Payload == nil {
		b, err := msgpack.Marshal(wf.PayloadJSON)
		if err != nil {
			return nil
		}
		wf.Payload = b
	}

	return rf.newCommandMessage(commandRequest, wf)

}

func (rf *requestFactory) topologyRequest() *Message {
	t := &zbmsgpack.TopologyRequest{}
	cmr := &zbsbe.ControlMessageRequest{
		MessageType: zbsbe.ControlMessageType.REQUEST_TOPOLOGY,

		Data: nil,
	}
	return rf.newControlMessage(cmr, t)
}

func (rf *requestFactory) deployWorkflowRequest(topic string, resources []*zbmsgpack.Resource) *Message {
	deployment := zbmsgpack.Workflow{
		State:     CreateDeployment,
		Resources: resources,
		TopicName: topic,
	}
	commandRequest := &zbsbe.ExecuteCommandRequest{
		PartitionId: 0,
		Position:    0,
		Command:     []uint8{},
	}
	commandRequest.Key = commandRequest.KeyNullValue()
	commandRequest.EventType = zbsbe.EventTypeEnum(4)
	return rf.newCommandMessage(commandRequest, deployment)
}

func (rf *requestFactory) openTaskSubscriptionRequest(partitionId uint16, lockOwner, taskType string, credits int32) *Message {
	taskSub := &zbmsgpack.TaskSubscription{
		Credits:       credits,
		LockDuration:  300000,
		LockOwner:     lockOwner,
		SubscriberKey: 0,
		TaskType:      taskType,
	}

	var msg Message
	b, err := msgpack.Marshal(taskSub)
	if err != nil {
		return nil
	}
	controlRequest := &zbsbe.ControlMessageRequest{
		MessageType: zbsbe.ControlMessageType.ADD_TASK_SUBSCRIPTION,
		PartitionId: partitionId,
		Data:        b,
	}
	msg.SetSbeMessage(controlRequest)
	msg.SetHeaders(rf.headers(controlRequest))

	return &msg
}

func (rf *requestFactory) increaseTaskSubscriptionCreditsRequest(ts *zbmsgpack.TaskSubscription) *Message {
	var msg Message

	b, err := msgpack.Marshal(ts)
	if err != nil {
		return nil
	}
	controlRequest := &zbsbe.ControlMessageRequest{
		MessageType: zbsbe.ControlMessageType.INCREASE_TASK_SUBSCRIPTION_CREDITS,
		Data:        b,
	}
	msg.SetSbeMessage(controlRequest)
	msg.SetHeaders(rf.headers(controlRequest))
	return &msg
}

func (rf *requestFactory) closeTaskSubscriptionRequest(ts *zbmsgpack.TaskSubscription) *Message {
	var msg Message

	b, err := msgpack.Marshal(ts)
	if err != nil {
		return nil
	}
	controlRequest := &zbsbe.ControlMessageRequest{
		MessageType: zbsbe.ControlMessageType.REMOVE_TASK_SUBSCRIPTION,
		Data:        b,
	}
	msg.SetSbeMessage(controlRequest)
	msg.SetHeaders(rf.headers(controlRequest))
	return &msg
}

func (rf *requestFactory) closeTopicSubscriptionRequest(ts *zbmsgpack.TopicSubscription) *Message {
	var msg Message

	b, err := msgpack.Marshal(ts)
	if err != nil {
		return nil
	}
	controlRequest := &zbsbe.ControlMessageRequest{
		MessageType: zbsbe.ControlMessageType.REMOVE_TOPIC_SUBSCRIPTION,
		PartitionId: ts.PartitionID,
		Data:        b,
	}
	msg.SetSbeMessage(controlRequest)
	msg.SetHeaders(rf.headers(controlRequest))
	return &msg
}

func (rf *requestFactory) openTopicSubscriptionRequest(partitionID uint16, topic, subName string, startPosition int64) *Message {
	ts := &zbmsgpack.OpenTopicSubscription{
		StartPosition:    startPosition,
		Name:             subName,
		PrefetchCapacity: 0,
		ForceStart:       true,
		State:            TopicSubscriptionSubscribeState,
	}
	execCommandRequest := &zbsbe.ExecuteCommandRequest{
		PartitionId: partitionID,
		Position:    0,
		EventType:   zbsbe.EventType.SUBSCRIBER_EVENT,
	}
	execCommandRequest.Key = execCommandRequest.KeyNullValue()

	var msg Message

	b, err := msgpack.Marshal(ts)
	if err != nil {
		return nil
	}
	execCommandRequest.Command = b
	msg.SetSbeMessage(execCommandRequest)
	msg.SetHeaders(rf.headers(execCommandRequest))
	return &msg
}

func (rf *requestFactory) topicSubscriptionAckRequest(ts *zbmsgpack.TopicSubscription, s *SubscriptionEvent) *Message {
	tsa := &zbmsgpack.TopicSubscriptionAck{
		Name:        ts.SubscriptionName,
		AckPosition: s.Event.Position,
		State:       TopicSubscriptionAckState,
	}
	execCommandRequest := &zbsbe.ExecuteCommandRequest{
		PartitionId: s.Event.PartitionId,
		Position:    0,
		EventType:   zbsbe.EventType.SUBSCRIPTION_EVENT,
	}
	execCommandRequest.Key = execCommandRequest.KeyNullValue()

	var msg Message

	b, err := msgpack.Marshal(tsa)
	if err != nil {
		return nil
	}
	execCommandRequest.Command = b
	msg.SetSbeMessage(execCommandRequest)
	msg.SetHeaders(rf.headers(execCommandRequest))
	return &msg
}

func (rf *requestFactory) createTopicRequest(topic *zbmsgpack.Topic) *Message {
	execCommandRequest := &zbsbe.ExecuteCommandRequest{
		PartitionId: 0,
		Position:    0,
		EventType:   zbsbe.EventType.TOPIC_EVENT,
	}
	execCommandRequest.Key = execCommandRequest.KeyNullValue()

	var msg Message

	b, err := msgpack.Marshal(topic)
	if err != nil {
		return nil
	}
	execCommandRequest.Command = b
	msg.SetSbeMessage(execCommandRequest)
	msg.SetHeaders(rf.headers(execCommandRequest))
	return &msg
}

func newRequestFactory() *requestFactory {
	return &requestFactory{}
}
