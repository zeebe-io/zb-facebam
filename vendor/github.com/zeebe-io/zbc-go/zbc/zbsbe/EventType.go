// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

type EventTypeEnum uint8
type EventTypeValues struct {
	TASK_EVENT              EventTypeEnum
	RAFT_EVENT              EventTypeEnum
	SUBSCRIPTION_EVENT      EventTypeEnum
	SUBSCRIBER_EVENT        EventTypeEnum
	DEPLOYMENT_EVENT        EventTypeEnum
	WORKFLOW_INSTANCE_EVENT EventTypeEnum
	INCIDENT_EVENT          EventTypeEnum
	WORKFLOW_EVENT          EventTypeEnum
	NOOP_EVENT              EventTypeEnum
	TOPIC_EVENT             EventTypeEnum
	NullValue               EventTypeEnum
}

var EventType = EventTypeValues{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

func (e EventTypeEnum) Encode(writer io.Writer, order binary.ByteOrder) error {
	if err := binary.Write(writer, order, e); err != nil {
		return err
	}
	return nil
}

func (e *EventTypeEnum) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16) error {
	if err := binary.Read(reader, order, e); err != nil {
		return err
	}
	return nil
}

func (e EventTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(EventType)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on EventType, unknown enumeration value %d", e)
}

func (e EventTypeEnum) EncodedLength() int64 {
	return 1
}

func (e EventTypeEnum) TASK_EVENTSinceVersion() uint16 {
	return 0
}

func (e EventTypeEnum) TASK_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TASK_EVENTSinceVersion()
}

func (e EventTypeEnum) TASK_EVENTDeprecated() uint16 {
	return 0
}

func (e EventTypeEnum) RAFT_EVENTSinceVersion() uint16 {
	return 0
}

func (e EventTypeEnum) RAFT_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RAFT_EVENTSinceVersion()
}

func (e EventTypeEnum) RAFT_EVENTDeprecated() uint16 {
	return 0
}

func (e EventTypeEnum) SUBSCRIPTION_EVENTSinceVersion() uint16 {
	return 0
}

func (e EventTypeEnum) SUBSCRIPTION_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SUBSCRIPTION_EVENTSinceVersion()
}

func (e EventTypeEnum) SUBSCRIPTION_EVENTDeprecated() uint16 {
	return 0
}

func (e EventTypeEnum) SUBSCRIBER_EVENTSinceVersion() uint16 {
	return 0
}

func (e EventTypeEnum) SUBSCRIBER_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SUBSCRIBER_EVENTSinceVersion()
}

func (e EventTypeEnum) SUBSCRIBER_EVENTDeprecated() uint16 {
	return 0
}

func (e EventTypeEnum) DEPLOYMENT_EVENTSinceVersion() uint16 {
	return 0
}

func (e EventTypeEnum) DEPLOYMENT_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.DEPLOYMENT_EVENTSinceVersion()
}

func (e EventTypeEnum) DEPLOYMENT_EVENTDeprecated() uint16 {
	return 0
}

func (e EventTypeEnum) WORKFLOW_EVENTSinceVersion() uint16 {
	return 0
}

func (e EventTypeEnum) WORKFLOW_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.WORKFLOW_EVENTSinceVersion()
}

func (e EventTypeEnum) WORKFLOW_EVENTDeprecated() uint16 {
	return 0
}

func (e EventTypeEnum) INCIDENT_EVENTSinceVersion() uint16 {
	return 0
}

func (e EventTypeEnum) INCIDENT_EVENTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.INCIDENT_EVENTSinceVersion()
}

func (e EventTypeEnum) INCIDENT_EVENTDeprecated() uint16 {
	return 0
}
