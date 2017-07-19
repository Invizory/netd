package netest

import (
	"io"
	"net"
)

// Fire dials to the listener to write the text and disconnect.
func Fire(listener net.Listener, text string) {
	conn := Dial(listener)
	defer conn.Close()
	_, err := io.WriteString(conn, text)
	if err != nil {
		panic(err)
	}
}

// Dial to the listener using its address.
func Dial(listener net.Listener) net.Conn {
	addr := listener.Addr()
	conn, err := net.Dial(addr.Network(), addr.String())
	if err != nil {
		panic(err)
	}
	return conn
}
