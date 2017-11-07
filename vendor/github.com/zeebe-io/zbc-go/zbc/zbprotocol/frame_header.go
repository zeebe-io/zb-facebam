package zbprotocol

import (
	"encoding/binary"
	"io"
)

const (
	// FrameTypeMessage is used to determine FrameType inside Message.
	FrameTypeMessage = iota
)

const (
	_ = iota
	// ControlClose representation of header value.
	ControlClose = 100 + iota

	// ControlEndOfStream representation of header value.
	ControlEndOfStream

	// ControlKeepAlive representation of header value.
	ControlKeepAlive

	// ProtocolControlFrame representation of header value.
	ProtocolControlFrame
)

// FrameHeader is first layer which we use in framing.
type FrameHeader struct {
	Length   uint32
	Version  uint8
	Flags    uint8
	TypeID   uint16 // One of the above defined constants.
	StreamID uint32
}

// Encode is used to serialize structure to byte array.
func (fh FrameHeader) Encode(writer io.Writer) error {
	return binary.Write(writer, binary.LittleEndian, fh)
}

// Decode is use to deserialize byte array to structure.
func (fh *FrameHeader) Decode(reader io.Reader, order binary.ByteOrder, _ uint16) error {
	return binary.Read(reader, order, fh)
}

// NewFrameHeader is constructor used to construct new FrameHeader object. Used mainly for writing purposes.
func NewFrameHeader(length uint32, version uint8, flags uint8, typeID uint16, streamID uint32) *FrameHeader {
	return &FrameHeader{
		length,
		version,
		flags,
		typeID,
		streamID,
	}
}
