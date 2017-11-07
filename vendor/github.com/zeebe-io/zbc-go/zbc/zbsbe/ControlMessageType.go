// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

type ControlMessageTypeEnum uint8
type ControlMessageTypeValues struct {
	ADD_TASK_SUBSCRIPTION              ControlMessageTypeEnum
	REMOVE_TASK_SUBSCRIPTION           ControlMessageTypeEnum
	INCREASE_TASK_SUBSCRIPTION_CREDITS ControlMessageTypeEnum
	REMOVE_TOPIC_SUBSCRIPTION          ControlMessageTypeEnum
	REQUEST_TOPOLOGY                   ControlMessageTypeEnum
	NullValue                          ControlMessageTypeEnum
}

var ControlMessageType = ControlMessageTypeValues{0, 1, 2, 3, 4, 255}

func (c ControlMessageTypeEnum) Encode(writer io.Writer, order binary.ByteOrder) error {
	if err := binary.Write(writer, order, c); err != nil {
		return err
	}
	return nil
}

func (c *ControlMessageTypeEnum) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16) error {
	if err := binary.Read(reader, order, c); err != nil {
		return err
	}
	return nil
}

func (c ControlMessageTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ControlMessageType)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ControlMessageType, unknown enumeration value %d", c)
}

func (c ControlMessageTypeEnum) EncodedLength() int64 {
	return 1
}

func (c ControlMessageTypeEnum) ADD_TASK_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (c ControlMessageTypeEnum) ADD_TASK_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ADD_TASK_SUBSCRIPTIONSinceVersion()
}

func (c ControlMessageTypeEnum) ADD_TASK_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}

func (c ControlMessageTypeEnum) REMOVE_TASK_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (c ControlMessageTypeEnum) REMOVE_TASK_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.REMOVE_TASK_SUBSCRIPTIONSinceVersion()
}

func (c ControlMessageTypeEnum) REMOVE_TASK_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}

func (c ControlMessageTypeEnum) INCREASE_TASK_SUBSCRIPTION_CREDITSSinceVersion() uint16 {
	return 0
}

func (c ControlMessageTypeEnum) INCREASE_TASK_SUBSCRIPTION_CREDITSInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.INCREASE_TASK_SUBSCRIPTION_CREDITSSinceVersion()
}

func (c ControlMessageTypeEnum) INCREASE_TASK_SUBSCRIPTION_CREDITSDeprecated() uint16 {
	return 0
}

func (c ControlMessageTypeEnum) REMOVE_TOPIC_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (c ControlMessageTypeEnum) REMOVE_TOPIC_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.REMOVE_TOPIC_SUBSCRIPTIONSinceVersion()
}

func (c ControlMessageTypeEnum) REMOVE_TOPIC_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}
