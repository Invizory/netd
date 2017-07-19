package netest

import (
	"net"

	"github.com/Invizory/netd"
)

// Once is a decorator for backend that closes the listener after handling.
func Once(backend netd.Backend, listener net.Listener) netd.Backend {
	if backend == nil {
		panic("testing: origin should not be nil")
	}
	if listener == nil {
		panic("testing: listener should not be nil")
	}
	return &once{backend, listener}
}

func (o *once) Handle(conn net.Conn) {
	defer o.listener.Close()
	o.origin.Handle(conn)
}

type once struct {
	origin   netd.Backend
	listener net.Listener
}
