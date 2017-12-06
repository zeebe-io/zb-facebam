package zbc

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/vmihailenco/msgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
)

var (
	errTimeout             = errors.New("requestWrapper timeout")
	errSocketWrite         = errors.New("tried to write more bytes to socket")
	errTopicLeaderNotFound = errors.New("topic leader not found")
	errResourceNotFound    = errors.New("resource not found")
)

// Client for Zeebe broker with support for clustered deployment.
type Client struct {
	*requestManager
}

// CreateTask will create new task on specified topic.
func (c *Client) CreateTask(topic string, task *zbmsgpack.Task) (*zbmsgpack.Task, error) {
	return c.createTask(topic, task)
}

// CreateWorkflow will deploy process to the broker.
func (c *Client) CreateWorkflow(topic string, resources ...*zbmsgpack.Resource) (*zbmsgpack.Workflow, error) {
	return c.createWorkflow(topic, resources)
}

func (c *Client) CreateWorkflowFromFile(topic, resourceType, path string) (*zbmsgpack.Workflow, error) {
	if len(path) == 0 {
		return nil, errResourceNotFound
	}

	filename, _ := filepath.Abs(path)
	definition, err := ioutil.ReadFile(filename)
	resource := NewResource(path, resourceType, definition)
	if err != nil {
		return nil, errResourceNotFound
	}
	return c.CreateWorkflow(topic, resource)
}

// CreateWorkflowInstance will create new workflow instance on the broker.
func (c *Client) CreateWorkflowInstance(topic string, workflowInstance *zbmsgpack.WorkflowInstance) (*zbmsgpack.WorkflowInstance, error) {
	return c.createWorkflowInstance(topic, workflowInstance)
}

// TaskConsumer opens a subscription on task and returns a channel where all the SubscribedEvents will arrive.
func (c *Client) TaskConsumer(topic, lockOwner, taskType string) (chan *SubscriptionEvent, *zbmsgpack.TaskSubscription, error) {
	return c.taskConsumer(topic, lockOwner, taskType, 32)
}

// CompleteTask will notify broker about finished task.
func (c *Client) CompleteTask(task *SubscriptionEvent) (*zbmsgpack.Task, error) {
	return c.completeTask(task)
}

//  IncreaseTaskSubscriptionCredits will increase the current credits of the task subscription.
func (c *Client) IncreaseTaskSubscriptionCredits(task *zbmsgpack.TaskSubscription) (*zbmsgpack.TaskSubscription, error) {
	return c.increaseTaskSubscriptionCredits(task)
}

// CloseTaskSubscription will tear down currently active task subscription.
func (c *Client) CloseTaskSubscription(task *zbmsgpack.TaskSubscription) (*Message, error) {
	return c.closeTaskSubscription(task)
}

// CloseTopicSubscription will tear down currently active topic subscription.
func (c *Client) CloseTopicSubscription(topicSub *zbmsgpack.TopicSubscription) (*Message, error) {
	return c.closeTopicSubscription(topicSub)
}

// TopicSubscriptionAck will ACK received events from the broker.
func (c *Client) TopicSubscriptionAck(ts *zbmsgpack.TopicSubscription, s *SubscriptionEvent) (*zbmsgpack.TopicSubscriptionAck, error) {
	return c.topicSubscriptionAck(ts, s)
}

// TopicConsumer opens a subscription on topic and returns a channel where all the SubscribedEvents will arrive.
func (c *Client) TopicConsumer(topic, subName string, startPosition int64) (chan *SubscriptionEvent, *zbmsgpack.TopicSubscription, error) {
	return c.topicConsumer(topic, subName, startPosition)
}

func (c *Client) CreateTopic(name string, partitionNum int) (*zbmsgpack.Topic, error) {
	return c.createTopic(name, partitionNum)
}

func (c *Client) UnmarshalFromFile(path string) (*Message, error) {
	data, fsErr := ioutil.ReadFile(path)

	if fsErr != nil {
		return nil, fsErr
	}

	messageReader := MessageReader{nil}
	return messageReader.readMessage(data)
}

func (c *Client) Topology() (*zbmsgpack.ClusterTopology, error) {
	return c.refreshTopology()
}

// NewClient is constructor for Client structure. It will resolve IP address and dial the provided tcp address.
func NewClient(bootstrapAddr string) (*Client, error) {
	c := &Client{
		newRequestManager(bootstrapAddr),
	}
	return c, nil
}

// NewTask is constructor for Task object. Function signature denotes mandatory fields.
func NewTask(typeName, lockOwner string) *zbmsgpack.Task {
	return &zbmsgpack.Task{
		State:        TaskCreate,
		Headers:      make(map[string]interface{}),
		CustomHeader: make(map[string]interface{}),

		Type:      typeName,
		LockOwner: lockOwner,
		Retries:   3,
	}
}

func NewWorkflowInstance(bpmnProcessId string, version int, payload map[string]interface{}) *zbmsgpack.WorkflowInstance {
	b, err := msgpack.Marshal(payload)
	if err != nil {
		return nil
	}
	return &zbmsgpack.WorkflowInstance{
		State:         CreateWorkflowInstance,
		BPMNProcessID: bpmnProcessId,
		Version:       version,
		Payload:       b,
	}
}

func NewResource(resourceName, resourceType string, resource []byte) *zbmsgpack.Resource {
	return &zbmsgpack.Resource{
		ResourceName: resourceName,
		ResourceType: resourceType,
		Resource:     resource,
	}
}
