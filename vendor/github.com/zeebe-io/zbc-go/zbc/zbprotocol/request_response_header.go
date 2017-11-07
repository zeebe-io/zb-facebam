package zbprotocol

import (
	"encoding/binary"
	"io"
	"math/rand"
	"time"
)

// RequestResponseHeader is layer to represent Request-Response model of communication. With it we keep transaction house keeping.
type RequestResponseHeader struct {
	RequestID uint64
}

// Encode is used to serialize structure to byte array.
func (fh RequestResponseHeader) Encode(writer io.Writer) error {
	return binary.Write(writer, binary.LittleEndian, fh)
}

// Decode is use to deserialize byte array to structure.
func (fh *RequestResponseHeader) Decode(reader io.Reader, order binary.ByteOrder, _ uint16) error {
	return binary.Read(reader, order, fh)
}

// NewRequestResponseHeader is constructor for RequestResponseHeader object. Constructor will generate random ID's for fields.
func NewRequestResponseHeader() *RequestResponseHeader {
	max := ^uint64(0)
	var s float64 = 10
	var v float64 = 1000
	zipf := rand.NewZipf(rand.New(rand.NewSource(time.Now().UnixNano())), s, v, max)
	return &RequestResponseHeader{
		zipf.Uint64(),
	}
}
