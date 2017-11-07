// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type VarDataEncoding struct {
	Length  uint16
	VarData uint8
}

func (v VarDataEncoding) Encode(writer io.Writer, order binary.ByteOrder) error {
	if err := binary.Write(writer, order, v.Length); err != nil {
		return err
	}
	if err := binary.Write(writer, order, v.VarData); err != nil {
		return err
	}
	return nil
}

func (v *VarDataEncoding) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16) error {
	if !v.LengthInActingVersion(actingVersion) {
		v.Length = v.LengthNullValue()
	} else {
		if err := binary.Read(reader, order, &v.Length); err != nil {
			return err
		}
	}
	if !v.VarDataInActingVersion(actingVersion) {
		v.VarData = v.VarDataNullValue()
	} else {
		if err := binary.Read(reader, order, &v.VarData); err != nil {
			return err
		}
	}
	return nil
}

func (v VarDataEncoding) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if v.LengthInActingVersion(actingVersion) {
		if v.Length < v.LengthMinValue() || v.Length > v.LengthMaxValue() {
			return fmt.Errorf("Range check failed on v.Length (%d < %d > %d)", v.LengthMinValue(), v.Length, v.LengthMaxValue())
		}
	}
	if v.VarDataInActingVersion(actingVersion) {
		if v.VarData < v.VarDataMinValue() || v.VarData > v.VarDataMaxValue() {
			return fmt.Errorf("Range check failed on v.VarData (%d < %d > %d)", v.VarDataMinValue(), v.VarData, v.VarDataMaxValue())
		}
	}
	return nil
}

func VarDataEncodingInit(v *VarDataEncoding) {
	return
}

func (v VarDataEncoding) EncodedLength() int64 {
	return -1
}

func (v VarDataEncoding) LengthMinValue() uint16 {
	return 0
}

func (v VarDataEncoding) LengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (v VarDataEncoding) LengthNullValue() uint16 {
	return math.MaxUint16
}

func (v VarDataEncoding) LengthSinceVersion() uint16 {
	return 0
}

func (v VarDataEncoding) LengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.LengthSinceVersion()
}

func (v VarDataEncoding) LengthDeprecated() uint16 {
	return 0
}

func (v VarDataEncoding) VarDataMinValue() uint8 {
	return 0
}

func (v VarDataEncoding) VarDataMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (v VarDataEncoding) VarDataNullValue() uint8 {
	return math.MaxUint8
}

func (v VarDataEncoding) VarDataSinceVersion() uint16 {
	return 0
}

func (v VarDataEncoding) VarDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= v.VarDataSinceVersion()
}

func (v VarDataEncoding) VarDataDeprecated() uint16 {
	return 0
}
