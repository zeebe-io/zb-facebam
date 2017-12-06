package zbmsgpack

import (
	"encoding/json"
	"fmt"
)

// Resource is message pack structure used to represent a workflow definition.
type Resource struct {
	Resource     []byte `yaml:"resource" msgpack:"resource"`
	ResourceType string `yaml:"resourceType" msgpack:"resourceType"`
	ResourceName string `yaml:"resourceName" msgpack:"resourceName"`
}

// Workflow is message pack structure used when creating a workflow.
type Workflow struct {
	State string `yaml:"state" msgpack:"state"`

	TopicName string      `yaml:"topicName" msgpack:"topicName"`
	Resources []*Resource `yaml:"resources" msgpack:"resources"`
}

func (t *Workflow) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}
