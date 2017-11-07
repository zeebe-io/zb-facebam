// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type ExecuteCommandResponse struct {
	PartitionId uint16
	Position    uint64
	Key         uint64
	Event       []uint8
}

func (e ExecuteCommandResponse) Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error {
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
	if err := binary.Write(writer, order, uint16(len(e.Event))); err != nil {
		return err
	}
	if err := binary.Write(writer, order, e.Event); err != nil {
		return err
	}
	return nil
}

func (e *ExecuteCommandResponse) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, reader, int64(blockLength-e.SbeBlockLength()))
	}

	if e.EventInActingVersion(actingVersion) {
		var EventLength uint16
		if err := binary.Read(reader, order, &EventLength); err != nil {
			return err
		}
		e.Event = make([]uint8, EventLength)
		if err := binary.Read(reader, order, &e.Event); err != nil {
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

func (e ExecuteCommandResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	return nil
}

func (e ExecuteCommandResponse) SbeBlockLength() (blockLength uint16) {
	return 18
}

func (e ExecuteCommandResponse) SbeTemplateId() (templateId uint16) {
	return 21
}

func (e ExecuteCommandResponse) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (e ExecuteCommandResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (e ExecuteCommandResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (e ExecuteCommandResponse) PartitionIdId() uint16 {
	return 1
}

func (e ExecuteCommandResponse) PartitionIdSinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandResponse) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PartitionIdSinceVersion()
}

func (e ExecuteCommandResponse) PartitionIdDeprecated() uint16 {
	return 0
}

func (e ExecuteCommandResponse) PartitionIdMetaAttribute(meta int) string {
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

func (e ExecuteCommandResponse) PartitionIdMinValue() uint16 {
	return 0
}

func (e ExecuteCommandResponse) PartitionIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (e ExecuteCommandResponse) PartitionIdNullValue() uint16 {
	return math.MaxUint16
}

func (e ExecuteCommandResponse) PositionMinValue() uint64 {
	return 0
}

func (e ExecuteCommandResponse) PositionMaxValue() uint64 {
	return math.MaxInt64 - 1
}

func (e ExecuteCommandResponse) PositionNullValue() uint64 {
	return math.MaxInt64
}

func (e ExecuteCommandResponse) KeyId() uint16 {
	return 2
}

func (e ExecuteCommandResponse) KeySinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandResponse) PositionSinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandResponse) KeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.KeySinceVersion()
}

func (e ExecuteCommandResponse) PositionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PositionSinceVersion()
}

func (e ExecuteCommandResponse) KeyDeprecated() uint16 {
	return 0
}

func (e ExecuteCommandResponse) KeyMetaAttribute(meta int) string {
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

func (e ExecuteCommandResponse) KeyMinValue() uint64 {
	return 0
}

func (e ExecuteCommandResponse) KeyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (e ExecuteCommandResponse) KeyNullValue() uint64 {
	return math.MaxUint64
}

func (e ExecuteCommandResponse) EventMetaAttribute(meta int) string {
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

func (e ExecuteCommandResponse) EventSinceVersion() uint16 {
	return 0
}

func (e ExecuteCommandResponse) EventInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.EventSinceVersion()
}

func (e ExecuteCommandResponse) EventDeprecated() uint16 {
	return 0
}

func (e ExecuteCommandResponse) EventCharacterEncoding() string {
	return "UTF-8"
}

func (e ExecuteCommandResponse) EventHeaderLength() uint64 {
	return 2
}
