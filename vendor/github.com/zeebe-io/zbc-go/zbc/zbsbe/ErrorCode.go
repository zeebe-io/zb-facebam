// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

type ErrorCodeEnum uint8
type ErrorCodeValues struct {
	MESSAGE_NOT_SUPPORTED      ErrorCodeEnum
	PARTITION_NOT_FOUND        ErrorCodeEnum
	REQUEST_WRITE_FAILURE      ErrorCodeEnum
	INVALID_CLIENT_VERSION     ErrorCodeEnum
	REQUEST_TIMEOUT            ErrorCodeEnum
	REQUEST_PROCESSING_FAILURE ErrorCodeEnum
	NullValue                  ErrorCodeEnum
}

var ErrorCode = ErrorCodeValues{0, 1, 2, 3, 4, 5, 255}

func (e ErrorCodeEnum) Encode(writer io.Writer, order binary.ByteOrder) error {
	if err := binary.Write(writer, order, e); err != nil {
		return err
	}
	return nil
}

func (e *ErrorCodeEnum) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16) error {
	if err := binary.Read(reader, order, e); err != nil {
		return err
	}
	return nil
}

func (e ErrorCodeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ErrorCode)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ErrorCode, unknown enumeration value %d", e)
}

func (e ErrorCodeEnum) EncodedLength() int64 {
	return 1
}

func (e ErrorCodeEnum) MESSAGE_NOT_SUPPORTEDSinceVersion() uint16 {
	return 0
}

func (e ErrorCodeEnum) MESSAGE_NOT_SUPPORTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MESSAGE_NOT_SUPPORTEDSinceVersion()
}

func (e ErrorCodeEnum) MESSAGE_NOT_SUPPORTEDDeprecated() uint16 {
	return 0
}

func (e ErrorCodeEnum) TOPIC_NOT_FOUNDSinceVersion() uint16 {
	return 0
}

func (e ErrorCodeEnum) TOPIC_NOT_FOUNDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TOPIC_NOT_FOUNDSinceVersion()
}

func (e ErrorCodeEnum) TOPIC_NOT_FOUNDDeprecated() uint16 {
	return 0
}

func (e ErrorCodeEnum) REQUEST_WRITE_FAILURESinceVersion() uint16 {
	return 0
}

func (e ErrorCodeEnum) REQUEST_WRITE_FAILUREInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.REQUEST_WRITE_FAILURESinceVersion()
}

func (e ErrorCodeEnum) REQUEST_WRITE_FAILUREDeprecated() uint16 {
	return 0
}

func (e ErrorCodeEnum) INVALID_CLIENT_VERSIONSinceVersion() uint16 {
	return 0
}

func (e ErrorCodeEnum) INVALID_CLIENT_VERSIONInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.INVALID_CLIENT_VERSIONSinceVersion()
}

func (e ErrorCodeEnum) INVALID_CLIENT_VERSIONDeprecated() uint16 {
	return 0
}

func (e ErrorCodeEnum) REQUEST_TIMEOUTSinceVersion() uint16 {
	return 0
}

func (e ErrorCodeEnum) REQUEST_TIMEOUTInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.REQUEST_TIMEOUTSinceVersion()
}

func (e ErrorCodeEnum) REQUEST_TIMEOUTDeprecated() uint16 {
	return 0
}

func (e ErrorCodeEnum) REQUEST_PROCESSING_FAILURESinceVersion() uint16 {
	return 0
}

func (e ErrorCodeEnum) REQUEST_PROCESSING_FAILUREInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.REQUEST_PROCESSING_FAILURESinceVersion()
}

func (e ErrorCodeEnum) REQUEST_PROCESSING_FAILUREDeprecated() uint16 {
	return 0
}
