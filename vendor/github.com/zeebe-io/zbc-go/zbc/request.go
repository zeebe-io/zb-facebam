package zbc

import (
	"github.com/vmihailenco/msgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbprotocol"
	"github.com/zeebe-io/zbc-go/zbc/zbsbe"
)

type requestHandler struct{}

func (rf *requestHandler) headers(t interface{}) *Headers {
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

func (rf *requestHandler) newCommandMessage(commandRequest *zbsbe.ExecuteCommandRequest, command interface{}) *Message {
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

func (rf *requestHandler) newControlMessage(req *zbsbe.ControlMessageRequest, payload interface{}) *Message {
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

func (rf *requestHandler) createTaskRequest(commandRequest *zbsbe.ExecuteCommandRequest, task *zbmsgpack.Task) *Message {
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

func (rf *requestHandler) completeTaskRequest(taskMessage *SubscriptionEvent) *Message {
	taskMessage.Task.State = TaskComplete
	cmdReq := &zbsbe.ExecuteCommandRequest{
		PartitionId: taskMessage.Event.PartitionId,
		Position:    taskMessage.Event.Position,
		Key:         taskMessage.Event.Key,
	}
	return rf.newCommandMessage(cmdReq, taskMessage.Task)
}

func (rf *requestHandler) createWorkflowInstanceRequest(commandRequest *zbsbe.ExecuteCommandRequest, wf *zbmsgpack.WorkflowInstance) *Message {
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

func (rf *requestHandler) topologyRequest() *Message {
	t := &zbmsgpack.TopologyRequest{}
	cmr := &zbsbe.ControlMessageRequest{
		MessageType: zbsbe.ControlMessageType.REQUEST_TOPOLOGY,

		Data: nil,
	}
	return rf.newControlMessage(cmr, t)
}

func (rf *requestHandler) newWorkflowRequest(commandRequest *zbsbe.ExecuteCommandRequest, d *zbmsgpack.Workflow) *Message {
	commandRequest.EventType = zbsbe.EventTypeEnum(4)
	return rf.newCommandMessage(commandRequest, d)
}

func (rf *requestHandler) openTaskSubscriptionRequest(partitionId uint16, ts *zbmsgpack.TaskSubscription) *Message {
	var msg Message

	b, err := msgpack.Marshal(ts)
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

func (rf *requestHandler) increaseTaskSubscriptionCreditsRequest(ts *zbmsgpack.TaskSubscription) *Message {
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

func (rf *requestHandler) closeTaskSubscriptionRequest(ts *zbmsgpack.TaskSubscription) *Message {
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

func (rf *requestHandler) closeTopicSubscriptionRequest(ts *zbmsgpack.TopicSubscription) *Message {
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

func (rf *requestHandler) openTopicSubscriptionRequest(cmdReq *zbsbe.ExecuteCommandRequest, ts *zbmsgpack.OpenTopicSubscription) *Message {
	var msg Message

	b, err := msgpack.Marshal(ts)
	if err != nil {
		return nil
	}
	cmdReq.Command = b
	msg.SetSbeMessage(cmdReq)
	msg.SetHeaders(rf.headers(cmdReq))
	return &msg
}

func (rf *requestHandler) topicSubscriptionAckRequest(cmdReq *zbsbe.ExecuteCommandRequest, tsa *zbmsgpack.TopicSubscriptionAck) *Message {
	var msg Message

	b, err := msgpack.Marshal(tsa)
	if err != nil {
		return nil
	}
	cmdReq.Command = b
	msg.SetSbeMessage(cmdReq)
	msg.SetHeaders(rf.headers(cmdReq))
	return &msg
}

func (rf *requestHandler) createTopicRequest(cmdReq *zbsbe.ExecuteCommandRequest, t *zbmsgpack.Topic) *Message {
	var msg Message

	b, err := msgpack.Marshal(t)
	if err != nil {
		return nil
	}
	cmdReq.Command = b
	msg.SetSbeMessage(cmdReq)
	msg.SetHeaders(rf.headers(cmdReq))
	return &msg
}
