package basic_test

import (
	"bytes"
	"net"
	"testing"

	"github.com/Invizory/netd"
	"github.com/Invizory/netd/backio"
	"github.com/Invizory/netd/basic"
	"github.com/Invizory/netd/internal/netest"
)

func TestClosedListener(t *testing.T) {
	listener, server := spawn()
	go listener.Close()
	server.Serve(backio.Echo())
}

func TestPipe(t *testing.T) {
	listener, server := spawn()
	const sent = "hello"
	go netest.Fire(listener, sent)
	buffer := new(bytes.Buffer)
	server.Serve(netest.Once(backio.Pipe(buffer), listener))
	if received := buffer.String(); sent != received {
		t.Error("sent and received strings should match")
	}
}

func spawn() (listener net.Listener, server netd.Server) {
	listener = netest.Listen()
	server = basic.Server(listener, nil)
	return
}
