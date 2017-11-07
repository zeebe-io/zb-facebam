package zbprotocol

import (
	"encoding/binary"
	"io"
)

const (
	// RequestResponse header type value.
	RequestResponse = iota
	// FullDuplexSingleMessage header type value.
	FullDuplexSingleMessage
)

// TransportHeader contains information about ProtocolID in use.
type TransportHeader struct {
	ProtocolID uint16
}

// Encode is used to serialize structure to byte array.
func (fh TransportHeader) Encode(writer io.Writer) error {
	return binary.Write(writer, binary.LittleEndian, fh)
}

// Decode is use to deserialize byte array to structure.
func (fh *TransportHeader) Decode(reader io.Reader, order binary.ByteOrder, _ uint16) error {
	return binary.Read(reader, order, fh)
}

// NewTransportHeader constructor
func NewTransportHeader(pid uint16) *TransportHeader {
	return &TransportHeader{
		pid,
	}
}
