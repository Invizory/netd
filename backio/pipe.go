package backio

import (
	"io"
	"net"

	"github.com/Invizory/netd"
)

// Pipe creates a backend that copies from input to the specified writer.
//
// Any errors during copy are silently ignored.
func Pipe(writer io.Writer) netd.Backend {
	return pipe{writer}
}

func (p pipe) Handle(conn net.Conn) {
	io.Copy(p, conn)
}

type pipe struct {
	io.Writer
}
