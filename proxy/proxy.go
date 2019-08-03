// Package proxy is an implementation of a proxy server.
// This package allows you to start a proxy server
// using a function and then provides API end-points
// to add and remove independent servers from this
// proxy server.
// The servers attached to the proxy must have a function
// in them to send a HTTP request to the proxy server to
// register itself in the proxy server. The server will
// have an option to be visible/not from direct access to
// it. Once the server sends a request to the proxy server
// to register itself, the proxy responds with a special
// hasb that is verified everytime a request goes through.
// Without this hash matching from the proxy server, the
// "normal" server rejects requests.
package proxy

import (
	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/gorilla/mux"
)

func newProxy() Server {
	return Server{
		Routes: make(map[string]string),
		Router: mux.NewRouter(),
	}
}

// StartProxy is used to start a single proxy server.
// Creates a new router, assigns the port to the proxy,
// and starts the server. The user cannot handle routing
// of the proxy server. All routing is handled by special
// mechanisms as described in adding and removing servers.
func StartProxy(r *mux.Router, port string) (Server, error) {
	ProxyServer = newProxy()
	ProxyServer.Router = r
	err := server.Server(ProxyServer.Router, port)
	if err != nil {
		return Server{}, err
	}
	return ProxyServer, nil
}
