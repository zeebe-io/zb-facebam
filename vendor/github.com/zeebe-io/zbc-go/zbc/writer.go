package zbc

import (
	"bytes"
	"encoding/binary"
	"log"
)

// MessageWriter is builder which will take Message pointer and create valid byte array.
type MessageWriter struct {
	message *Message
}

func (mw *MessageWriter) writeFrameHeader(writer *bytes.Buffer) error {
	err := mw.message.Headers.FrameHeader.Encode(writer)
	if err != nil {
		return err
	}
	return nil
}

func (mw *MessageWriter) writeTransportHeader(writer *bytes.Buffer) error {
	err := mw.message.Headers.TransportHeader.Encode(writer)
	if err != nil {
		return err
	}
	return nil
}

func (mw *MessageWriter) writeRequestResponseHeader(writer *bytes.Buffer) error {
	if mw.message.Headers.IsSingleMessage() {
		return nil
	}
	err := mw.message.Headers.RequestResponseHeader.Encode(writer)
	if err != nil {
		return err
	}
	return nil
}

func (mw *MessageWriter) writeSbeMessageHeader(writer *bytes.Buffer) error {
	if err := mw.message.Headers.SbeMessageHeader.Encode(writer, binary.LittleEndian); err != nil {
		return err
	}
	return nil
}

func (mw *MessageWriter) writeHeaders(writer *bytes.Buffer) error {
	if err := mw.writeFrameHeader(writer); err != nil {
		return err
	}
	if err := mw.writeTransportHeader(writer); err != nil {
		return err
	}
	if err := mw.writeRequestResponseHeader(writer); err != nil {
		return err
	}
	if err := mw.writeSbeMessageHeader(writer); err != nil {
		return err
	}
	return nil
}

func (mw *MessageWriter) writeMessage(writer *bytes.Buffer) error {
	if err := (*mw.message.SbeMessage).Encode(writer, binary.LittleEndian, false); err != nil {
		return err
	}
	return nil
}

func (mw *MessageWriter) align(writer *bytes.Buffer) {
	currentSize := len(writer.Bytes())
	expectedSize := (currentSize + 7) & ^7
	for currentSize < expectedSize {
		writer.Write([]byte{0x00})
		currentSize++
	}
}

func (mw *MessageWriter) Write(writer *bytes.Buffer) {
	err := mw.writeHeaders(writer)
	if err != nil {
		log.Fatal("failed writing header")
	}
	err = mw.writeMessage(writer)
	if err != nil {
		log.Fatalf("failed writing message")
	}
	mw.align(writer)
}

// NewMessageWriter constructor for MessageWriter builder.
func NewMessageWriter(message *Message) *MessageWriter {
	return &MessageWriter{
		message,
	}
}
