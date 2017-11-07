package zbmsgpack

import (
	"encoding/json"
	"fmt"
)

// Workflow is msgpack structure used when creating a workflow
type Workflow struct {
	State        string `yaml:"state" msgpack:"state"`
	ResourceType string `yaml:"resourceType" msgpack:"resourceType"`
	TopicName    string `yaml:"topicName" msgpack:"topicName"`
	Resource     []byte `yaml:"resource" msgpack:"resource"`
}

func (t *Workflow) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}
