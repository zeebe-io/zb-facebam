package zbc

import "time"

// RequestTimeout specifies default timeout for responder.
const RequestTimeout = 30

// TopologyRefreshInterval defines time to live of topology object.
const TopologyRefreshInterval = 30

// Retry constants
const (
	BackoffMin      = 20 * time.Millisecond
	BackoffMax      = 1000 * time.Millisecond
	BackoffDeadline = 30 * time.Second
)

// Sbe template ID constants
const (
	templateIDExecuteCommandRequest  = 20
	templateIDExecuteCommandResponse = 21
	templateIDControlMessageResponse = 11
	templateIDSubscriptionEvent      = 30
)

// Zeebe protocol constants
const (
	FrameHeaderSize           = 12
	TransportHeaderSize       = 2
	RequestResponseHeaderSize = 8
	SBEMessageHeaderSize      = 8

	TotalHeaderSizeNoFrame = 18
	TotalHeaderSize        = 30
	LengthFieldSize        = 2
)

// Message pack states for Task, Deployment and WorkflowInstance
const (
	TaskCreate  = "CREATE"
	TaskCreated = "CREATED"

	TaskComplete  = "COMPLETE"
	TaskCompleted = "COMPLETED"

	CreateDeployment   = "CREATE_DEPLOYMENT"
	DeployementCreated = "DEPLOYMENT_CREATED"

	CreateWorkflowInstance   = "CREATE_WORKFLOW_INSTANCE"
	WorkflowInstanceCreated  = "WORKFLOW_INSTANCE_CREATED"
	WorkflowInstanceRejected = "WORKFLOW_INSTANCE_REJECTED"
)

//
const (
	TopicCreate   = "CREATE"
	TopicCreated  = "CREATED"
	TopicRejected = "CREATE_REJECTED"
)

// TopicSubscription states
const (
	TopicSubscriptionSubscribeState  = "SUBSCRIBE"
	TopicSubscriptionSubscribedState = "SUBSCRIBED"
)

// TopicSubscriptionAck states
const (
	TopicSubscriptionAckState          = "ACKNOWLEDGE"
	TopicSubscriptionAcknowledgedState = "ACKNOWLEDGED"
)

const (
	SystemTopic = "internal-system"
)

// Workflow resource types
const (
	BpmnXml      = "BPMN_XML"
	YamlWorkflow = "YAML_WORKFLOW"
)

const (
	SocketChunkSize = 4096
)
