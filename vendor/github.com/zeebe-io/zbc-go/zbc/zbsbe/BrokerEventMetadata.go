// Generated SBE (Simple Binary Encoding) message codec

package zbsbe

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type BrokerEventMetadata struct {
	ReqStreamId     int32
	ReqRequestId    uint64
	RaftTermId      int32
	SubscriptionId  uint64
	ProtocolVersion uint16
	EventType       EventTypeEnum
	IncidentKey     uint64
}

func (b BrokerEventMetadata) Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error {
	if doRangeCheck {
		if err := b.RangeCheck(b.SbeSchemaVersion(), b.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := binary.Write(writer, order, b.ReqStreamId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, b.ReqRequestId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, b.RaftTermId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, b.SubscriptionId); err != nil {
		return err
	}
	if err := binary.Write(writer, order, b.ProtocolVersion); err != nil {
		return err
	}
	if err := b.EventType.Encode(writer, order); err != nil {
		return err
	}
	if err := binary.Write(writer, order, b.IncidentKey); err != nil {
		return err
	}
	return nil
}

func (b *BrokerEventMetadata) Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !b.ReqChannelIdInActingVersion(actingVersion) {
		b.ReqStreamId = b.ReqChannelIdNullValue()
	} else {
		if err := binary.Read(reader, order, &b.ReqStreamId); err != nil {
			return err
		}
	}
	if !b.ReqRequestIdInActingVersion(actingVersion) {
		b.ReqRequestId = b.ReqRequestIdNullValue()
	} else {
		if err := binary.Read(reader, order, &b.ReqRequestId); err != nil {
			return err
		}
	}
	if !b.RaftTermIdInActingVersion(actingVersion) {
		b.RaftTermId = b.RaftTermIdNullValue()
	} else {
		if err := binary.Read(reader, order, &b.RaftTermId); err != nil {
			return err
		}
	}
	if !b.SubscriptionIdInActingVersion(actingVersion) {
		b.SubscriptionId = b.SubscriptionIdNullValue()
	} else {
		if err := binary.Read(reader, order, &b.SubscriptionId); err != nil {
			return err
		}
	}
	if !b.ProtocolVersionInActingVersion(actingVersion) {
		b.ProtocolVersion = b.ProtocolVersionNullValue()
	} else {
		if err := binary.Read(reader, order, &b.ProtocolVersion); err != nil {
			return err
		}
	}
	if b.EventTypeInActingVersion(actingVersion) {
		if err := b.EventType.Decode(reader, order, actingVersion); err != nil {
			return err
		}
	}
	if !b.IncidentKeyInActingVersion(actingVersion) {
		b.IncidentKey = b.IncidentKeyNullValue()
	} else {
		if err := binary.Read(reader, order, &b.IncidentKey); err != nil {
			return err
		}
	}
	if actingVersion > b.SbeSchemaVersion() && blockLength > b.SbeBlockLength() {
		io.CopyN(ioutil.Discard, reader, int64(blockLength-b.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := b.RangeCheck(actingVersion, b.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (b BrokerEventMetadata) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if b.ReqChannelIdInActingVersion(actingVersion) {
		if b.ReqStreamId < b.ReqChannelIdMinValue() || b.ReqStreamId > b.ReqChannelIdMaxValue() {
			return fmt.Errorf("Range check failed on b.ReqStreamId (%d < %d > %d)", b.ReqChannelIdMinValue(), b.ReqStreamId, b.ReqChannelIdMaxValue())
		}
	}
	if b.ReqRequestIdInActingVersion(actingVersion) {
		if b.ReqRequestId < b.ReqRequestIdMinValue() || b.ReqRequestId > b.ReqRequestIdMaxValue() {
			return fmt.Errorf("Range check failed on b.ReqRequestId (%d < %d > %d)", b.ReqRequestIdMinValue(), b.ReqRequestId, b.ReqRequestIdMaxValue())
		}
	}
	if b.RaftTermIdInActingVersion(actingVersion) {
		if b.RaftTermId < b.RaftTermIdMinValue() || b.RaftTermId > b.RaftTermIdMaxValue() {
			return fmt.Errorf("Range check failed on b.RaftTermId (%d < %d > %d)", b.RaftTermIdMinValue(), b.RaftTermId, b.RaftTermIdMaxValue())
		}
	}
	if b.SubscriptionIdInActingVersion(actingVersion) {
		if b.SubscriptionId < b.SubscriptionIdMinValue() || b.SubscriptionId > b.SubscriptionIdMaxValue() {
			return fmt.Errorf("Range check failed on b.SubscriptionId (%d < %d > %d)", b.SubscriptionIdMinValue(), b.SubscriptionId, b.SubscriptionIdMaxValue())
		}
	}
	if b.ProtocolVersionInActingVersion(actingVersion) {
		if b.ProtocolVersion < b.ProtocolVersionMinValue() || b.ProtocolVersion > b.ProtocolVersionMaxValue() {
			return fmt.Errorf("Range check failed on b.ProtocolVersion (%d < %d > %d)", b.ProtocolVersionMinValue(), b.ProtocolVersion, b.ProtocolVersionMaxValue())
		}
	}
	if err := b.EventType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if b.IncidentKeyInActingVersion(actingVersion) {
		if b.IncidentKey < b.IncidentKeyMinValue() || b.IncidentKey > b.IncidentKeyMaxValue() {
			return fmt.Errorf("Range check failed on b.IncidentKey (%d < %d > %d)", b.IncidentKeyMinValue(), b.IncidentKey, b.IncidentKeyMaxValue())
		}
	}
	return nil
}

func BrokerEventMetadataInit(b *BrokerEventMetadata) {
	return
}

func (b BrokerEventMetadata) SbeBlockLength() (blockLength uint16) {
	return 43
}

func (b BrokerEventMetadata) SbeTemplateId() (templateId uint16) {
	return 200
}

func (b BrokerEventMetadata) SbeSchemaId() (schemaId uint16) {
	return 0
}

func (b BrokerEventMetadata) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (b BrokerEventMetadata) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (b BrokerEventMetadata) ReqChannelIdId() uint16 {
	return 1
}

func (b BrokerEventMetadata) ReqChannelIdSinceVersion() uint16 {
	return 0
}

func (b BrokerEventMetadata) ReqChannelIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.ReqChannelIdSinceVersion()
}

func (b BrokerEventMetadata) ReqChannelIdDeprecated() uint16 {
	return 0
}

func (b BrokerEventMetadata) ReqChannelIdMetaAttribute(meta int) string {
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

func (b BrokerEventMetadata) ReqChannelIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (b BrokerEventMetadata) ReqChannelIdMaxValue() int32 {
	return math.MaxInt32
}

func (b BrokerEventMetadata) ReqChannelIdNullValue() int32 {
	return math.MinInt32
}

func (b BrokerEventMetadata) ReqRequestIdId() uint16 {
	return 3
}

func (b BrokerEventMetadata) ReqRequestIdSinceVersion() uint16 {
	return 0
}

func (b BrokerEventMetadata) ReqRequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.ReqRequestIdSinceVersion()
}

func (b BrokerEventMetadata) ReqRequestIdDeprecated() uint16 {
	return 0
}

func (b BrokerEventMetadata) ReqRequestIdMetaAttribute(meta int) string {
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

func (b BrokerEventMetadata) ReqRequestIdMinValue() uint64 {
	return 0
}

func (b BrokerEventMetadata) ReqRequestIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (b BrokerEventMetadata) ReqRequestIdNullValue() uint64 {
	return math.MaxUint64
}

func (b BrokerEventMetadata) RaftTermIdId() uint16 {
	return 4
}

func (b BrokerEventMetadata) RaftTermIdSinceVersion() uint16 {
	return 0
}

func (b BrokerEventMetadata) RaftTermIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.RaftTermIdSinceVersion()
}

func (b BrokerEventMetadata) RaftTermIdDeprecated() uint16 {
	return 0
}

func (b BrokerEventMetadata) RaftTermIdMetaAttribute(meta int) string {
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

func (b BrokerEventMetadata) RaftTermIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (b BrokerEventMetadata) RaftTermIdMaxValue() int32 {
	return math.MaxInt32
}

func (b BrokerEventMetadata) RaftTermIdNullValue() int32 {
	return math.MinInt32
}

func (b BrokerEventMetadata) SubscriptionIdId() uint16 {
	return 5
}

func (b BrokerEventMetadata) SubscriptionIdSinceVersion() uint16 {
	return 0
}

func (b BrokerEventMetadata) SubscriptionIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.SubscriptionIdSinceVersion()
}

func (b BrokerEventMetadata) SubscriptionIdDeprecated() uint16 {
	return 0
}

func (b BrokerEventMetadata) SubscriptionIdMetaAttribute(meta int) string {
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

func (b BrokerEventMetadata) SubscriptionIdMinValue() uint64 {
	return 0
}

func (b BrokerEventMetadata) SubscriptionIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (b BrokerEventMetadata) SubscriptionIdNullValue() uint64 {
	return math.MaxUint64
}

func (b BrokerEventMetadata) ProtocolVersionId() uint16 {
	return 6
}

func (b BrokerEventMetadata) ProtocolVersionSinceVersion() uint16 {
	return 0
}

func (b BrokerEventMetadata) ProtocolVersionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.ProtocolVersionSinceVersion()
}

func (b BrokerEventMetadata) ProtocolVersionDeprecated() uint16 {
	return 0
}

func (b BrokerEventMetadata) ProtocolVersionMetaAttribute(meta int) string {
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

func (b BrokerEventMetadata) ProtocolVersionMinValue() uint16 {
	return 0
}

func (b BrokerEventMetadata) ProtocolVersionMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (b BrokerEventMetadata) ProtocolVersionNullValue() uint16 {
	return math.MaxUint16
}

func (b BrokerEventMetadata) EventTypeId() uint16 {
	return 7
}

func (b BrokerEventMetadata) EventTypeSinceVersion() uint16 {
	return 0
}

func (b BrokerEventMetadata) EventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.EventTypeSinceVersion()
}

func (b BrokerEventMetadata) EventTypeDeprecated() uint16 {
	return 0
}

func (b BrokerEventMetadata) EventTypeMetaAttribute(meta int) string {
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

func (b BrokerEventMetadata) IncidentKeyId() uint16 {
	return 8
}

func (b BrokerEventMetadata) IncidentKeySinceVersion() uint16 {
	return 0
}

func (b BrokerEventMetadata) IncidentKeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= b.IncidentKeySinceVersion()
}

func (b BrokerEventMetadata) IncidentKeyDeprecated() uint16 {
	return 0
}

func (b BrokerEventMetadata) IncidentKeyMetaAttribute(meta int) string {
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

func (b BrokerEventMetadata) IncidentKeyMinValue() uint64 {
	return 0
}

func (b BrokerEventMetadata) IncidentKeyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (b BrokerEventMetadata) IncidentKeyNullValue() uint64 {
	return math.MaxUint64
}
