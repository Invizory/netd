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

// BrokenListener creates fake net.Listener that returns error on every Accept.
func BrokenListener(err error) net.Listener {
	return broken{err}
}

func (b broken) Accept() (net.Conn, error) {
	return nil, b.error
}

func (broken) Addr() net.Addr {
	return nil
}

func (broken) Close() error {
	return nil
}

type broken struct {
	error
}
