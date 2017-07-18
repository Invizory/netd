// Package netd provides useful primitives for building network daemons.
package netd

import "net"

// Backend handles incoming connections.
type Backend interface {
	Handle(net.Conn)
}

// Server handles multiple incoming connections using the specified Backend.
type Server interface {
	Serve(Backend)
}
