package zbmsgpack

import (
	"encoding/json"
	"fmt"
)

type WorkflowInstance struct {
	State         string                 `yaml:"state" msgpack:"state"`
	BPMNProcessID string                 `yaml:"bpmnProcessId" msgpack:"bpmnProcessId"`
	Version       int                    `yaml:"version" msgpack:"version"`
	Payload       []uint8                `yaml:"-" msgpack:"payload"`
	PayloadJSON   map[string]interface{} `yaml:"payload" msgpack:"-"`
}

func (t *WorkflowInstance) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}
