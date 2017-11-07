// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type MessageHeader struct {
	BlockLength uint16
	TemplateId  uint16
	SchemaId    uint16
	Version     uint16
}

func (m MessageHeader) Encode(writer io.Writer, order binary.ByteOrder) error {
	if err := binary.Write(writer, order, m.BlockLength); err != nil {
		return err
	}
	if err := binary.Write(writer, order, m.TemplateId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, m.SchemaId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, m.Version); err != nil {
		return err
	}
	return nil
}

func (m *MessageHeader) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16) error {
	if !m.BlockLengthInActingVersion(actingVersion) {
		m.BlockLength = m.BlockLengthNullValue()
	} else {
		if err := binary.Read(reader, order, &m.BlockLength); err != nil {
			return err
		}
	}
	if !m.TemplateIdInActingVersion(actingVersion) {
		m.TemplateId = m.TemplateIdNullValue()
	} else {
		if err := binary.Read(reader, order, &m.TemplateId); err != nil {
			return err
		}
	}
	if !m.SchemaIdInActingVersion(actingVersion) {
		m.SchemaId = m.SchemaIdNullValue()
	} else {
		if err := binary.Read(reader, order, &m.SchemaId); err != nil {
			return err
		}
	}
	if !m.VersionInActingVersion(actingVersion) {
		m.Version = m.VersionNullValue()
	} else {
		if err := binary.Read(reader, order, &m.Version); err != nil {
			return err
		}
	}
	return nil
}

func (m MessageHeader) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.BlockLengthInActingVersion(actingVersion) {
		if m.BlockLength < m.BlockLengthMinValue() || m.BlockLength > m.BlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on m.BlockLength (%d < %d > %d)", m.BlockLengthMinValue(), m.BlockLength, m.BlockLengthMaxValue())
		}
	}
	if m.TemplateIdInActingVersion(actingVersion) {
		if m.TemplateId < m.TemplateIdMinValue() || m.TemplateId > m.TemplateIdMaxValue() {
			return fmt.Errorf("Range check failed on m.TemplateId (%d < %d > %d)", m.TemplateIdMinValue(), m.TemplateId, m.TemplateIdMaxValue())
		}
	}
	if m.SchemaIdInActingVersion(actingVersion) {
		if m.SchemaId < m.SchemaIdMinValue() || m.SchemaId > m.SchemaIdMaxValue() {
			return fmt.Errorf("Range check failed on m.SchemaId (%d < %d > %d)", m.SchemaIdMinValue(), m.SchemaId, m.SchemaIdMaxValue())
		}
	}
	if m.VersionInActingVersion(actingVersion) {
		if m.Version < m.VersionMinValue() || m.Version > m.VersionMaxValue() {
			return fmt.Errorf("Range check failed on m.Version (%d < %d > %d)", m.VersionMinValue(), m.Version, m.VersionMaxValue())
		}
	}
	return nil
}

func MessageHeaderInit(m *MessageHeader) {
	return
}

func (m MessageHeader) EncodedLength() int64 {
	return 8
}

func (m MessageHeader) BlockLengthMinValue() uint16 {
	return 0
}

func (m MessageHeader) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (m MessageHeader) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}

func (m MessageHeader) BlockLengthSinceVersion() uint16 {
	return 0
}

func (m MessageHeader) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BlockLengthSinceVersion()
}

func (m MessageHeader) BlockLengthDeprecated() uint16 {
	return 0
}

func (m MessageHeader) TemplateIdMinValue() uint16 {
	return 0
}

func (m MessageHeader) TemplateIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (m MessageHeader) TemplateIdNullValue() uint16 {
	return math.MaxUint16
}

func (m MessageHeader) TemplateIdSinceVersion() uint16 {
	return 0
}

func (m MessageHeader) TemplateIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TemplateIdSinceVersion()
}

func (m MessageHeader) TemplateIdDeprecated() uint16 {
	return 0
}

func (m MessageHeader) SchemaIdMinValue() uint16 {
	return 0
}

func (m MessageHeader) SchemaIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (m MessageHeader) SchemaIdNullValue() uint16 {
	return math.MaxUint16
}

func (m MessageHeader) SchemaIdSinceVersion() uint16 {
	return 0
}

func (m MessageHeader) SchemaIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SchemaIdSinceVersion()
}

func (m MessageHeader) SchemaIdDeprecated() uint16 {
	return 0
}

func (m MessageHeader) VersionMinValue() uint16 {
	return 0
}

func (m MessageHeader) VersionMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (m MessageHeader) VersionNullValue() uint16 {
	return math.MaxUint16
}

func (m MessageHeader) VersionSinceVersion() uint16 {
	return 0
}

func (m MessageHeader) VersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.VersionSinceVersion()
}

func (m MessageHeader) VersionDeprecated() uint16 {
	return 0
}
