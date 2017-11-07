// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type ErrorResponse struct {
	ErrorCode     ErrorCodeEnum
	ErrorData     []uint8
	FailedRequest []uint8
}

func (e ErrorResponse) Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := e.ErrorCode.Encode(writer, order); err != nil {
		return err
	}
	if err := binary.Write(writer, order, uint16(len(e.ErrorData))); err != nil {
		return err
	}
	if err := binary.Write(writer, order, e.ErrorData); err != nil {
		return err
	}
	if err := binary.Write(writer, order, uint16(len(e.FailedRequest))); err != nil {
		return err
	}
	if err := binary.Write(writer, order, e.FailedRequest); err != nil {
		return err
	}
	return nil
}

func (e *ErrorResponse) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if e.ErrorCodeInActingVersion(actingVersion) {
		if err := e.ErrorCode.Decode(reader, order, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, reader, int64(blockLength-e.SbeBlockLength()))
	}

	if e.ErrorDataInActingVersion(actingVersion) {
		var ErrorDataLength uint16
		if err := binary.Read(reader, order, &ErrorDataLength); err != nil {
			return err
		}
		e.ErrorData = make([]uint8, ErrorDataLength)
		if err := binary.Read(reader, order, &e.ErrorData); err != nil {
			return err
		}
	}

	if e.FailedRequestInActingVersion(actingVersion) {
		var FailedRequestLength uint16
		if err := binary.Read(reader, order, &FailedRequestLength); err != nil {
			return err
		}
		e.FailedRequest = make([]uint8, FailedRequestLength)
		if err := binary.Read(reader, order, &e.FailedRequest); err != nil {
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

func (e ErrorResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := e.ErrorCode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(e.ErrorData[:]) {
		return errors.New("e.ErrorData failed UTF-8 validation")
	}
	if !utf8.Valid(e.FailedRequest[:]) {
		return errors.New("e.FailedRequest failed UTF-8 validation")
	}
	return nil
}

func ErrorResponseInit(e *ErrorResponse) {
	return
}

func (e ErrorResponse) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (e ErrorResponse) SbeTemplateId() (templateId uint16) {
	return 0
}

func (e ErrorResponse) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (e ErrorResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (e ErrorResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (e ErrorResponse) ErrorCodeId() uint16 {
	return 1
}

func (e ErrorResponse) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (e ErrorResponse) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorCodeSinceVersion()
}

func (e ErrorResponse) ErrorCodeDeprecated() uint16 {
	return 0
}

func (e ErrorResponse) ErrorCodeMetaAttribute(meta int) string {
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

func (e ErrorResponse) ErrorDataMetaAttribute(meta int) string {
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

func (e ErrorResponse) ErrorDataSinceVersion() uint16 {
	return 0
}

func (e ErrorResponse) ErrorDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorDataSinceVersion()
}

func (e ErrorResponse) ErrorDataDeprecated() uint16 {
	return 0
}

func (e ErrorResponse) ErrorDataCharacterEncoding() string {
	return "UTF-8"
}

func (e ErrorResponse) ErrorDataHeaderLength() uint64 {
	return 2
}

func (e ErrorResponse) FailedRequestMetaAttribute(meta int) string {
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

func (e ErrorResponse) FailedRequestSinceVersion() uint16 {
	return 0
}

func (e ErrorResponse) FailedRequestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FailedRequestSinceVersion()
}

func (e ErrorResponse) FailedRequestDeprecated() uint16 {
	return 0
}

func (e ErrorResponse) FailedRequestCharacterEncoding() string {
	return "UTF-8"
}

func (e ErrorResponse) FailedRequestHeaderLength() uint64 {
	return 2
}
