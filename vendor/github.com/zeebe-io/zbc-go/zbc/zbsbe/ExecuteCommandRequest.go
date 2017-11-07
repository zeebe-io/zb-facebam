// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecuteCommandRequest struct {
	PartitionId uint16
	Position    uint64
	Key         uint64
	EventType   EventTypeEnum // alias for uint8 or byte
	Command     []uint8
}

func (e ExecuteCommandRequest) Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := binary.Write(writer, order, e.PartitionId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, e.Position); err != nil {
		return err
	}
	if err := binary.Write(writer, order, e.Key); err != nil {
		return err
	}
	if err := e.EventType.Encode(writer, order); err != nil {
		return err
	}
	if err := binary.Write(writer, order, uint16(len(e.Command))); err != nil {
		return err
	}
	if err := binary.Write(writer, order, e.Command); err != nil {
		return err
	}
	return nil
}

func (e *ExecuteCommandRequest) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.PartitionIdInActingVersion(actingVersion) {
		e.PartitionId = e.PartitionIdNullValue()
	} else {
		if err := binary.Read(reader, order, &e.PartitionId); err != nil {
			return err
		}
	}

	if !e.PositionInActingVersion(actingVersion) {
		e.Position = e.PositionNullValue()
	} else {
		if err := binary.Read(reader, order, &e.Position); err != nil {
			return err
		}
	}
	if !e.KeyInActingVersion(actingVersion) {
		e.Key = e.KeyNullValue()
	} else {
		if err := binary.Read(reader, order, &e.Key); err != nil {
			return err
		}
	}
	if e.EventTypeInActingVersion(actingVersion) {
		if err := e.EventType.Decode(reader, order, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, reader, int64(blockLength-e.SbeBlockLength()))
	}

	if e.CommandInActingVersion(actingVersion) {
		var CommandLength uint16
		if err := binary.Read(reader, order, &CommandLength); err != nil {
			return err
		}
		e.Command = make([]uint8, CommandLength)
		if err := binary.Read(reader, order, &e.Command); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e ExecuteCommandRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.PartitionIdInActingVersion(actingVersion) {
		if e.PartitionId < e.PartitionIdMinValue() || e.PartitionId > e.PartitionIdMaxValue() {
			return fmt.Errorf("Range check failed on e.PartitionId (%d < %d > %d)", e.PartitionIdMinValue(), e.PartitionId, e.PartitionIdMaxValue())
		}
	}
	if e.KeyInActingVersion(actingVersion) {
		if e.Key < e.KeyMinValue() || e.Key > e.KeyMaxValue() {
			return fmt.Errorf("Range check failed on e.Key (%d < %d > %d)", e.KeyMinValue(), e.Key, e.KeyMaxValue())
		}
	}
	if e.PositionInActingVersion(actingVersion) {
		if e.Position < e.PositionMinValue() || e.Position > e.PositionMaxValue() {
			return fmt.Errorf("Range check failed on e.Position (%d < %d > %d)", e.PositionMinValue(), e.Position, e.PositionMaxValue())
		}
	}
	if err := e.EventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func (e ExecuteCommandRequest) SbeBlockLength() (blockLength uint16) {
	return 19
}

func (e ExecuteCommandRequest) SbeTemplateId() (templateId uint16) {
	return 20
}

func (e ExecuteCommandRequest) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (e ExecuteCommandRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (e ExecuteCommandRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (e ExecuteCommandRequest) PartitionIdId() uint16 {
	return 1
}

func (e ExecuteCommandRequest) PartitionIdSinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandRequest) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartitionIdSinceVersion()
}

func (e ExecuteCommandRequest) PartitionIdDeprecated() uint16 {
	return 0
}

func (e ExecuteCommandRequest) PartitionIdMetaAttribute(meta int) string {
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

func (e ExecuteCommandRequest) PartitionIdMinValue() uint16 {
	return 0
}

func (e ExecuteCommandRequest) PartitionIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (e ExecuteCommandRequest) PartitionIdNullValue() uint16 {
	return math.MaxUint16
}

func (e ExecuteCommandRequest) PositionMinValue() uint64 {
	return 0
}

func (e ExecuteCommandRequest) PositionMaxValue() uint64 {
	return math.MaxInt64 - 1
}

func (e ExecuteCommandRequest) PositionNullValue() uint64 {
	return math.MaxInt64
}

func (e ExecuteCommandRequest) KeyId() uint16 {
	return 2
}

func (e ExecuteCommandRequest) KeySinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandRequest) PositionSinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandRequest) KeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.KeySinceVersion()
}

func (e ExecuteCommandRequest) PositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PositionSinceVersion()
}

func (e ExecuteCommandRequest) KeyDeprecated() uint16 {
	return 0
}

func (e ExecuteCommandRequest) KeyMetaAttribute(meta int) string {
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

func (e ExecuteCommandRequest) KeyMinValue() uint64 {
	return 0
}

func (e ExecuteCommandRequest) KeyMaxValue() uint64 {
	return math.MaxUint64
}

func (e ExecuteCommandRequest) KeyNullValue() uint64 {
	return math.MaxUint64
}

func (e ExecuteCommandRequest) EventTypeId() uint16 {
	return 3
}

func (e ExecuteCommandRequest) EventTypeSinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandRequest) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.EventTypeSinceVersion()
}

func (e ExecuteCommandRequest) EventTypeDeprecated() uint16 {
	return 0
}

func (e ExecuteCommandRequest) EventTypeMetaAttribute(meta int) string {
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

func (e ExecuteCommandRequest) CommandMetaAttribute(meta int) string {
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

func (e ExecuteCommandRequest) CommandSinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandRequest) CommandInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CommandSinceVersion()
}

func (e ExecuteCommandRequest) CommandDeprecated() uint16 {
	return 0
}

func (e ExecuteCommandRequest) CommandCharacterEncoding() string {
	return "UTF-8"
}

func (e ExecuteCommandRequest) CommandHeaderLength() uint64 {
	return 2
}
