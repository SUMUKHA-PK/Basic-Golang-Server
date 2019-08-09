package proxy

import (
	"sync"

	"github.com/gorilla/mux"
)

// Server holds the data about a proxy server
type Server struct {
	Routes map[string]string // Hash of server and IP mapping
	Router *mux.Router       // The proxy router
	Mutex  sync.Mutex
}

// ProxyServer is the server datatype
var ProxyServer Server

// AddServerReq is the incoming request
// for adding a server to the proxy
type AddServerReq struct {
	IP string
}

// AddServerRes is the response for
// /addServer request
type AddServerRes struct {
	Hash string
}
