// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SubscribedEvent struct {
	PartitionId      uint16
	Position         uint64
	Key              uint64
	SubscriberKey    uint64
	SubscriptionType SubscriptionTypeEnum
	EventType        EventTypeEnum
	Event            []uint8
}

func (s SubscribedEvent) Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := binary.Write(writer, order, s.PartitionId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, s.Position); err != nil {
		return err
	}
	if err := binary.Write(writer, order, s.Key); err != nil {
		return err
	}
	if err := binary.Write(writer, order, s.SubscriberKey); err != nil {
		return err
	}
	if err := s.SubscriptionType.Encode(writer, order); err != nil {
		return err
	}
	if err := s.EventType.Encode(writer, order); err != nil {
		return err
	}
	if err := binary.Write(writer, order, uint16(len(s.Event))); err != nil {
		return err
	}
	if err := binary.Write(writer, order, s.Event); err != nil {
		return err
	}
	return nil
}

func (s *SubscribedEvent) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.PartitionIdInActingVersion(actingVersion) {
		s.PartitionId = s.PartitionIdNullValue()
	} else {
		if err := binary.Read(reader, order, &s.PartitionId); err != nil {
			return err
		}
	}
	if !s.PositionInActingVersion(actingVersion) {
		s.Position = s.PositionNullValue()
	} else {
		if err := binary.Read(reader, order, &s.Position); err != nil {
			return err
		}
	}
	if !s.KeyInActingVersion(actingVersion) {
		s.Key = s.KeyNullValue()
	} else {
		if err := binary.Read(reader, order, &s.Key); err != nil {
			return err
		}
	}
	if !s.SubscriberKeyInActingVersion(actingVersion) {
		s.SubscriberKey = s.SubscriberKeyNullValue()
	} else {
		if err := binary.Read(reader, order, &s.SubscriberKey); err != nil {
			return err
		}
	}
	if s.SubscriptionTypeInActingVersion(actingVersion) {
		if err := s.SubscriptionType.Decode(reader, order, actingVersion); err != nil {
			return err
		}
	}
	if s.EventTypeInActingVersion(actingVersion) {
		if err := s.EventType.Decode(reader, order, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, reader, int64(blockLength-s.SbeBlockLength()))
	}

	if s.EventInActingVersion(actingVersion) {
		var EventLength uint16
		if err := binary.Read(reader, order, &EventLength); err != nil {
			return err
		}
		s.Event = make([]uint8, EventLength)
		if err := binary.Read(reader, order, &s.Event); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s SubscribedEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.PartitionIdInActingVersion(actingVersion) {
		if s.PartitionId < s.PartitionIdMinValue() || s.PartitionId > s.PartitionIdMaxValue() {
			return fmt.Errorf("Range check failed on s.PartitionId (%d < %d > %d)", s.PartitionIdMinValue(), s.PartitionId, s.PartitionIdMaxValue())
		}
	}
	if s.PositionInActingVersion(actingVersion) {
		if s.Position < s.PositionMinValue() || s.Position > s.PositionMaxValue() {
			return fmt.Errorf("Range check failed on s.Position (%d < %d > %d)", s.PositionMinValue(), s.Position, s.PositionMaxValue())
		}
	}
	if s.KeyInActingVersion(actingVersion) {
		if s.Key < s.KeyMinValue() || s.Key > s.KeyMaxValue() {
			return fmt.Errorf("Range check failed on s.Key (%d < %d > %d)", s.KeyMinValue(), s.Key, s.KeyMaxValue())
		}
	}
	if s.SubscriberKeyInActingVersion(actingVersion) {
		if s.SubscriberKey < s.SubscriberKeyMinValue() || s.SubscriberKey > s.SubscriberKeyMaxValue() {
			return fmt.Errorf("Range check failed on s.SubscriberKey (%d < %d > %d)", s.SubscriberKeyMinValue(), s.SubscriberKey, s.SubscriberKeyMaxValue())
		}
	}
	if err := s.SubscriptionType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := s.EventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func (s SubscribedEvent) SbeBlockLength() (blockLength uint16) {
	return 28
}

func (s SubscribedEvent) SbeTemplateId() (templateId uint16) {
	return 30
}

func (s SubscribedEvent) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (s SubscribedEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (s SubscribedEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (s SubscribedEvent) PartitionIdId() uint16 {
	return 1
}

func (s SubscribedEvent) PartitionIdSinceVersion() uint16 {
	return 0
}

func (s SubscribedEvent) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PartitionIdSinceVersion()
}

func (s SubscribedEvent) PartitionIdDeprecated() uint16 {
	return 0
}

func (s SubscribedEvent) PartitionIdMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	}
	return ""
}

func (s SubscribedEvent) PartitionIdMinValue() uint16 {
	return 0
}

func (s SubscribedEvent) PartitionIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (s SubscribedEvent) PartitionIdNullValue() uint16 {
	return math.MaxUint16
}

func (s SubscribedEvent) PositionId() uint16 {
	return 2
}

func (s SubscribedEvent) PositionSinceVersion() uint16 {
	return 0
}

func (s SubscribedEvent) PositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PositionSinceVersion()
}

func (s SubscribedEvent) PositionDeprecated() uint16 {
	return 0
}

func (s SubscribedEvent) PositionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	}
	return ""
}

func (s SubscribedEvent) PositionMinValue() uint64 {
	return 0
}

func (s SubscribedEvent) PositionMaxValue() uint64 {
	return math.MaxUint64
}

func (s SubscribedEvent) PositionNullValue() uint64 {
	return math.MaxUint64
}

func (s SubscribedEvent) KeyId() uint16 {
	return 3
}

func (s SubscribedEvent) KeySinceVersion() uint16 {
	return 0
}

func (s SubscribedEvent) KeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.KeySinceVersion()
}

func (s SubscribedEvent) KeyDeprecated() uint16 {
	return 0
}

func (s SubscribedEvent) KeyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	}
	return ""
}

func (s SubscribedEvent) KeyMinValue() uint64 {
	return 0
}

func (s SubscribedEvent) KeyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (s SubscribedEvent) KeyNullValue() uint64 {
	return math.MaxUint64
}

func (s SubscribedEvent) SubscriberKeyId() uint16 {
	return 4
}

func (s SubscribedEvent) SubscriberKeySinceVersion() uint16 {
	return 0
}

func (s SubscribedEvent) SubscriberKeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SubscriberKeySinceVersion()
}

func (s SubscribedEvent) SubscriberKeyDeprecated() uint16 {
	return 0
}

func (s SubscribedEvent) SubscriberKeyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	}
	return ""
}

func (s SubscribedEvent) SubscriberKeyMinValue() uint64 {
	return 0
}

func (s SubscribedEvent) SubscriberKeyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (s SubscribedEvent) SubscriberKeyNullValue() uint64 {
	return math.MaxUint64
}

func (s SubscribedEvent) SubscriptionTypeId() uint16 {
	return 5
}

func (s SubscribedEvent) SubscriptionTypeSinceVersion() uint16 {
	return 0
}

func (s SubscribedEvent) SubscriptionTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SubscriptionTypeSinceVersion()
}

func (s SubscribedEvent) SubscriptionTypeDeprecated() uint16 {
	return 0
}

func (s SubscribedEvent) SubscriptionTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	}
	return ""
}

func (s SubscribedEvent) EventTypeId() uint16 {
	return 6
}

func (s SubscribedEvent) EventTypeSinceVersion() uint16 {
	return 0
}

func (s SubscribedEvent) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.EventTypeSinceVersion()
}

func (s SubscribedEvent) EventTypeDeprecated() uint16 {
	return 0
}

func (s SubscribedEvent) EventTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	}
	return ""
}

func (s SubscribedEvent) EventMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return "unix"
	case 2:
		return "nanosecond"
	case 3:
		return ""
	}
	return ""
}

func (s SubscribedEvent) EventSinceVersion() uint16 {
	return 0
}

func (s SubscribedEvent) EventInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.EventSinceVersion()
}

func (s SubscribedEvent) EventDeprecated() uint16 {
	return 0
}

func (s SubscribedEvent) EventCharacterEncoding() string {
	return "UTF-8"
}

func (s SubscribedEvent) EventHeaderLength() uint64 {
	return 2
}
