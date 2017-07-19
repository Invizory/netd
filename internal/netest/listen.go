package netest

import "net"

// Listen creates local net.Listener on ephemeral port.
func Listen() net.Listener {
	listener, err := net.Listen("tcp", "[::1]:")
	if err != nil {
		panic(err)
	}
	return listener
}
