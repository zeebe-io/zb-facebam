package zbmsgpack

import (
	"encoding/json"
	"fmt"
)

// Task structure is used when creating or read a task.
type Task struct {
	State        string                 `yaml:"state" msgpack:"state"`
	LockTime     uint64                 `yaml:"lockTime" msgpack:"lockTime"`
	LockOwner    string                 `yaml:"lockOwner" msgpack:"lockOwner"`
	Headers      map[string]interface{} `yaml:"headers" msgpack:"headers"`
	CustomHeader map[string]interface{} `yaml:"customHeaders" msgpack:"customHeaders"`
	Retries      int                    `yaml:"retries" msgpack:"retries"`
	Type         string                 `yaml:"type" msgpack:"type"`
	Payload      []uint8                `yaml:"-" msgpack:"payload"`
	PayloadJSON  map[string]interface{} `yaml:"payload" msgpack:"-" json:"-"`
}

func (t *Task) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}
