package backio

import (
	"io"
	"net"

	"github.com/Invizory/netd"
)

// Motd creates a backend that copies from the specified publisher to output.
//
// Any errors during copy are silently ignored.
func Motd(publisher io.WriterTo) netd.Backend {
	return motd{publisher}
}

func (m motd) Handle(conn net.Conn) {
	m.WriteTo(conn)
}

type motd struct {
	io.WriterTo
}
