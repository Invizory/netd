package listen

import (
	"net"
	"net/url"
)

// URI creates net.Listener from address in the URI format.
//
// Examples:
//     URI("tcp:127.0.0.1:1337")
//     URI("tcp:[::1]:1337")
//     URI("unix:/tmp/netd.sock")
//
// See net.Listen for more details about listening and address syntax.
func URI(address string) (net.Listener, error) {
	uri, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	if uri.Scheme == "" {
		return nil, newError("network should be present as URI scheme")
	}
	return net.Listen(qualified(uri))
}

func qualified(url *url.URL) (network, address string) {
	network, address = url.Scheme, url.Opaque
	if address == "" {
		address = url.EscapedPath()
	}
	return
}
