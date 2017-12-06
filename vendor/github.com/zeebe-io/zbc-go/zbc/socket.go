package zbc

import (
	"bytes"
	"github.com/zeebe-io/zbc-go/zbc/zbsbe"
	"net"
	"time"
)

type socket struct {
	dispatcher

	connection net.Conn
	stream     []byte
	closeCh    chan bool
}

// responder implements synchronous way of sending ExecuteCommandRequest and waiting for ExecuteCommandResponse.
func (c *socket) responder(request *requestWrapper) (*Message, error) {
	message := request.payload
	respCh := make(chan *Message)

	c.addTransaction(message.Headers.RequestResponseHeader.RequestID, respCh)
	if err := c.sender(message); err != nil {
		return nil, err
	}

	select {
	case resp := <-respCh:
		c.removeTransaction(message.Headers.RequestResponseHeader.RequestID)
		return resp, nil
	case <-time.After(time.Second * RequestTimeout):
		c.removeTransaction(message.Headers.RequestResponseHeader.RequestID)
		return nil, errTimeout
	}
}

func (s *socket) sender(message *Message) error {
	writer := NewMessageWriter(message)
	byteBuff := &bytes.Buffer{}
	writer.Write(byteBuff)

	n, err := s.connection.Write(byteBuff.Bytes())
	if err != nil {
		return err
	}

	if n != len(byteBuff.Bytes()) {
		return errSocketWrite
	}
	return nil
}

func (s *socket) receiver() {
	reader := NewMessageReader(s)
	responseHandler := responseHandler{}

	for {
		select {
		case <-s.closeCh:
			s.connection.Close()
			return

		default:

			headers, tail, err := reader.readHeaders()
			if err != nil {
				continue
			}
			message, err := reader.parseMessage(headers, tail)

			if err != nil && !headers.IsSingleMessage() {
				s.removeTransaction(headers.RequestResponseHeader.RequestID)
				continue
			}

			if !headers.IsSingleMessage() && message != nil && len(message.Data) > 0 {
				s.dispatchTransaction(headers.RequestResponseHeader.RequestID, message)
				continue
			}

			if err != nil && headers.IsSingleMessage() {
				continue
			}

			if headers.IsSingleMessage() && message != nil && len(message.Data) > 0 {
				event := (*message.SbeMessage).(*zbsbe.SubscribedEvent)
				if task := responseHandler.unmarshalTask(message); task != nil {
					s.dispatchTaskEvent(event.SubscriberKey, event, task)
				} else {
					s.dispatchTopicEvent(event.SubscriberKey, event)
				}
			}
		}
	}
}

func (s *socket) teardown() {
	close(s.closeCh)
}

func (s *socket) readChunk() {
	for {
		s.connection.SetReadDeadline(time.Now().Add(time.Millisecond * 100))

		total := make([]byte, SocketChunkSize)
		returnedSize, err := s.connection.Read(total)
		if err != nil {
			continue
		}

		s.stream = append(s.stream, total[:returnedSize]...)
		break
	}
}

func (s *socket) getBytes(start, end int) []byte {
	for {
		if end-start > len(s.stream) || end > len(s.stream) {
			s.readChunk()
		} else {
			break
		}
	}

	frame := s.stream[start:end]
	return frame
}

func (s *socket) popBytes(pos int) {
	s.stream = s.stream[pos:]
}

func (s *socket) dial(addr string) error {
	tcpAddr, wrongAddr := net.ResolveTCPAddr("tcp4", addr)
	if wrongAddr != nil {
		return wrongAddr
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	s.connection = conn
	return nil
}

func newSocketStream(addr string) *socket {
	ss := &socket{
		newDispatcher(),
		nil,
		make([]byte, 0),
		make(chan bool),
	}

	err := ss.dial(addr)
	if err != nil {
		return nil
	}

	go ss.receiver()

	return ss
}
