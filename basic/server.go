package basic

import (
	"net"

	"github.com/Invizory/netd"
)

// Server creates basic server that accepts connections using the specified
// listener.
func Server(listener net.Listener, errors chan<- error) netd.Server {
	if listener == nil {
		panic("simple: listener should not be nil")
	}
	return &server{listener, errors}
}

func (srv server) Serve(backend netd.Backend) {
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			if isClosed(err) {
				break
			}
			srv.errors <- err
			continue
		}
		backend.Handle(conn)
	}
}

func isClosed(err error) bool {
	// https://github.com/golang/go/issues/4373
	if err, ok := err.(*net.OpError); ok {
		return err.Err.Error() == "use of closed network connection"
	}
	return false
}

type server struct {
	listener net.Listener
	errors   chan<- error
}
