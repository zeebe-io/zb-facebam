package zbc

import (
	"errors"
)

var brokerNotFound = errors.New("cannot contact the broker")

type transportManager struct {
	transportWorkload chan *requestWrapper

	connections map[string]*socket
}

// TODO: check if socket in map is unreachable/unhealty than delete that record from the map
// TODO: keep-alive messages

func (tm *transportManager) getSocket(addr string) (*socket, error) {
	if conn, ok := tm.connections[addr]; ok {
		return conn, nil
	}

	sock := newSocketStream(addr)
	if sock == nil {
		return nil, brokerNotFound
	}

	tm.connections[addr] = sock
	return tm.connections[addr], nil
}


func (tm *transportManager) execTransport(request *requestWrapper) {
	tm.transportWorkload <- request
}

func (tm *transportManager) transportWorker() {
	for {
		select {
		case request := <- tm.transportWorkload:
			sock, err := tm.getSocket(request.addr)
			if err != nil {
				request.errorCh <- err
			}
			request.sock = sock
			resp, err := MessageRetry(func() (*Message, error) {
				return sock.responder(request)
			})

			if err != nil {
				request.errorCh <- err
				continue
			}

			request.responseCh <- resp
		}
	}
}

func newTransportManager() *transportManager {
	tm := &transportManager{
		make(chan *requestWrapper, requestQueueSize),
		make(map[string]*socket),
	}

	go tm.transportWorker()
	return tm
}
