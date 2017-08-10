package listen_test

import (
	"fmt"

	"github.com/Invizory/netd/listen"
)

func ExampleURI() {
	listener, _ := listen.URI("tcp:[::1]:1337")
	defer listener.Close()
	fmt.Println(listener.Addr())
	// Output: [::1]:1337
}

func ExampleURI_unix() {
	listener, _ := listen.URI("unix:/tmp/netd.sock")
	defer listener.Close()
	fmt.Println(listener.Addr())
	// Output: /tmp/netd.sock
}
