package zbc

import (
	"net"
	"time"
)

type Socket struct {
	connection net.Conn
	stream     []byte
}

func (ss *Socket) ReadChunk() {
	for {
		ss.connection.SetReadDeadline(time.Now().Add(time.Millisecond * 100))

		total := make([]byte, SocketChunkSize)
		returnedSize, err := ss.connection.Read(total)
		if err != nil {
			continue
		}

		ss.stream = append(ss.stream, total[:returnedSize]...)
		break
	}
}

func (ss *Socket) getBytes(start, end int) []byte {
	for {
		if end-start > len(ss.stream) || end > len(ss.stream) {
			ss.ReadChunk()
		} else {
			break
		}
	}

	frame := ss.stream[start:end]
	return frame
}

func (ss *Socket) PopBytes(pos int) {
	ss.stream = ss.stream[pos:]
}

func NewSocketStream(conn net.Conn) *Socket {
	ss := &Socket{
		conn,
		make([]byte, 0),
	}
	return ss
}
