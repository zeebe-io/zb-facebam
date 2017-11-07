// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type ControlMessageRequest struct {
	MessageType ControlMessageTypeEnum
	PartitionId uint16
	Data        []uint8
}

func (c ControlMessageRequest) Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error {
	if doRangeCheck {
		if err := c.RangeCheck(c.SbeSchemaVersion(), c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := c.MessageType.Encode(writer, order); err != nil {
		return err
	}
	if err := binary.Write(writer, order, c.PartitionId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, uint16(len(c.Data))); err != nil {
		return err
	}
	if err := binary.Write(writer, order, c.Data); err != nil {
		return err
	}
	return nil
}

func (c *ControlMessageRequest) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if c.MessageTypeInActingVersion(actingVersion) {
		if err := c.MessageType.Decode(reader, order, actingVersion); err != nil {
			return err
		}
	}
	if c.PartitionIdInActingVersion(actingVersion) {
		if err := binary.Read(reader, order, &c.PartitionId); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, reader, int64(blockLength-c.SbeBlockLength()))
	}

	if c.DataInActingVersion(actingVersion) {
		var DataLength uint16
		if err := binary.Read(reader, order, &DataLength); err != nil {
			return err
		}
		c.Data = make([]uint8, DataLength)
		if err := binary.Read(reader, order, &c.Data); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := c.RangeCheck(actingVersion, c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (c ControlMessageRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := c.MessageType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(c.Data[:]) {
		return errors.New("c.Data failed UTF-8 validation")
	}
	return nil
}

func (c ControlMessageRequest) SbeBlockLength() (blockLength uint16) {
	return 3
}

func (c ControlMessageRequest) SbeTemplateId() (templateId uint16) {
	return 10
}

func (c ControlMessageRequest) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (c ControlMessageRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (c ControlMessageRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (c ControlMessageRequest) MessageTypeId() uint16 {
	return 1
}

func (c ControlMessageRequest) MessageTypeSinceVersion() uint16 {
	return 0
}

func (c ControlMessageRequest) MessageTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.MessageTypeSinceVersion()
}

func (c ControlMessageRequest) MessageTypeDeprecated() uint16 {
	return 0
}

func (c ControlMessageRequest) MessageTypeMetaAttribute(meta int) string {
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

func (c ControlMessageRequest) DataMetaAttribute(meta int) string {
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

func (c ControlMessageRequest) DataSinceVersion() uint16 {
	return 0
}

func (c ControlMessageRequest) DataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.DataSinceVersion()
}

func (c ControlMessageRequest) DataDeprecated() uint16 {
	return 0
}

func (c ControlMessageRequest) DataCharacterEncoding() string {
	return "UTF-8"
}

func (c ControlMessageRequest) DataHeaderLength() uint64 {
	return 2
}

func (c *ControlMessageRequest) PartitionIdSinceVersion() uint16 {
	return 0
}

func (c *ControlMessageRequest) PartitionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.PartitionIdSinceVersion()
}
