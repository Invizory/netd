package backio

import (
	"io"
	"net"

	"github.com/Invizory/netd"
)

// Echo creates a backend that copies input to output.
//
// Any errors during copy are silently ignored.
func Echo() netd.Backend {
	return echo{}
}

func (echo) Handle(conn net.Conn) {
	io.Copy(conn, conn)
}

type echo struct{}
