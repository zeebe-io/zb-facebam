// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

type SubscriptionTypeEnum uint8
type SubscriptionTypeValues struct {
	TASK_SUBSCRIPTION  SubscriptionTypeEnum
	TOPIC_SUBSCRIPTION SubscriptionTypeEnum
	NullValue          SubscriptionTypeEnum
}

var SubscriptionType = SubscriptionTypeValues{0, 1, 255}

func (s SubscriptionTypeEnum) Encode(writer io.Writer, order binary.ByteOrder) error {
	if err := binary.Write(writer, order, s); err != nil {
		return err
	}
	return nil
}

func (s *SubscriptionTypeEnum) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16) error {
	if err := binary.Read(reader, order, s); err != nil {
		return err
	}
	return nil
}

func (s SubscriptionTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SubscriptionType)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SubscriptionType, unknown enumeration value %d", s)
}

func (s SubscriptionTypeEnum) EncodedLength() int64 {
	return 1
}

func (s SubscriptionTypeEnum) TASK_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (s SubscriptionTypeEnum) TASK_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TASK_SUBSCRIPTIONSinceVersion()
}

func (s SubscriptionTypeEnum) TASK_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}

func (s SubscriptionTypeEnum) TOPIC_SUBSCRIPTIONSinceVersion() uint16 {
	return 0
}

func (s SubscriptionTypeEnum) TOPIC_SUBSCRIPTIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TOPIC_SUBSCRIPTIONSinceVersion()
}

func (s SubscriptionTypeEnum) TOPIC_SUBSCRIPTIONDeprecated() uint16 {
	return 0
}
