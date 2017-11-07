// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"io"
	"io/ioutil"
)

type ControlMessageResponse struct {
	Data []uint8
}

func (c ControlMessageResponse) Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error {
	if doRangeCheck {
		if err := c.RangeCheck(c.SbeSchemaVersion(), c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := binary.Write(writer, order, uint16(len(c.Data))); err != nil {
		return err
	}
	if err := binary.Write(writer, order, c.Data); err != nil {
		return err
	}
	return nil
}

func (c *ControlMessageResponse) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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

func (c ControlMessageResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	return nil
}

func ControlMessageResponseInit(c *ControlMessageResponse) {
	return
}

func (c ControlMessageResponse) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (c ControlMessageResponse) SbeTemplateId() (templateId uint16) {
	return 11
}

func (c ControlMessageResponse) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (c ControlMessageResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (c ControlMessageResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (c ControlMessageResponse) DataMetaAttribute(meta int) string {
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

func (c ControlMessageResponse) DataSinceVersion() uint16 {
	return 0
}

func (c ControlMessageResponse) DataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.DataSinceVersion()
}

func (c ControlMessageResponse) DataDeprecated() uint16 {
	return 0
}

func (c ControlMessageResponse) DataCharacterEncoding() string {
	return "UTF-8"
}

func (c ControlMessageResponse) DataHeaderLength() uint64 {
	return 2
}
