package zbc

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/zeebe-io/zbc-go/zbc/zbmsgpack"
	"github.com/zeebe-io/zbc-go/zbc/zbprotocol"
	"github.com/zeebe-io/zbc-go/zbc/zbsbe"
	"io"
)

// Headers is aggregator for all headers. It holds pointer to every layer. If RequestResponseHeader is nil, then IsSingleMessage will always return true.
type Headers struct {
	FrameHeader           *zbprotocol.FrameHeader
	TransportHeader       *zbprotocol.TransportHeader
	RequestResponseHeader *zbprotocol.RequestResponseHeader // If this is nil then struct is equal to SingleMessage
	SbeMessageHeader      *zbsbe.MessageHeader
}

// SetFrameHeader is a setter for FrameHeader.
func (h *Headers) SetFrameHeader(header *zbprotocol.FrameHeader) {
	h.FrameHeader = header
}

// SetTransportHeader is a setter for TransportHeader.
func (h *Headers) SetTransportHeader(header *zbprotocol.TransportHeader) {
	h.TransportHeader = header
}

// SetRequestResponseHeader is a setting for RequestResponseHeader.
func (h *Headers) SetRequestResponseHeader(header *zbprotocol.RequestResponseHeader) {
	h.RequestResponseHeader = header
}

// IsSingleMessage is helper to determine which model of communication we use.
func (h *Headers) IsSingleMessage() bool {
	return h.RequestResponseHeader == nil
}

// SetSbeMessageHeader is a setter for SBEMessageHeader.
func (h *Headers) SetSbeMessageHeader(header *zbsbe.MessageHeader) {
	h.SbeMessageHeader = header
}

// SBE interface is abstraction over all SBE Messages.
type SBE interface {
	Encode(writer io.Writer, order binary.ByteOrder, doRangeCheck bool) error
	Decode(reader io.Reader, order binary.ByteOrder, actingVersion uint16, blockLength uint16, doRangeCheck bool) error
}

// Message is Zeebe message which will contain pointers to all parts of the message. Data is Message Pack layer.
type Message struct {
	Headers    *Headers
	SbeMessage *SBE
	Data       []byte
}

// SetHeaders is a setter for Headers attribute.
func (m *Message) SetHeaders(headers *Headers) {
	m.Headers = headers
}

// SetSbeMessage is a setter for SBE attribute.
func (m *Message) SetSbeMessage(data SBE) {
	m.SbeMessage = &data
}

// SetData is a setter for unmarshaled message pack data.
func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) jsonString(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("json marshaling failed\n")
	}
	return fmt.Sprintf("%+v", string(b))
}

// SubscriptionEvent is used on task and topic subscription.
type SubscriptionEvent struct {
	Task  *zbmsgpack.Task
	Event *zbsbe.SubscribedEvent
}

func (se *SubscriptionEvent) String() string {
	b, _ := json.MarshalIndent(se, "", "  ")
	return fmt.Sprintf("%+v", string(b))
}
