package proxy

import (
	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/gorilla/mux"
)

func newProxy() Server {
	return Server{
		Routes:   make(map[string]string),
		PortList: make(map[string]int),
		Router:   mux.NewRouter(),
	}
}

// StartProxy is used to start a single proxy server.
// Creates a new router, assigns the port to the proxy,
// and starts the server. The user cannot handle routing
// of the proxy server. All routing is handled by special
// mechanisms as described in adding and removing servers.
func StartProxy(port string) (Server, error) {
	proxy := newProxy()
	err := server.Server(proxy.Router, port)
	if err != nil {
		return Server{}, err
	}
	proxy.PortList[port] = 1
	return proxy, nil
}

// AddServerToProxy adds a server to one of the routes of the proxy.
// Function returns a "hash" for the new server added. This is the
// route to which the proxy sends the requests to for this server.
func AddServerToProxy() (string, error) {
	return "", nil
}

// RemoveServerFromProxy removes a server from the proxy and its route
func RemoveServerFromProxy() error {
	return nil
}
