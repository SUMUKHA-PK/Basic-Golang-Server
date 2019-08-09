package proxy

import (
	"sync"

	"github.com/gorilla/mux"
)

// Server holds the data about a proxy server
type Server struct {
	RouteMap SyncMap     // A custom sync map
	Router   *mux.Router // The proxy router
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

// SyncMap is a custom sync map
type SyncMap struct {
	Routes map[string]string // Hash of server and IP mapping
	Mutex  sync.Mutex        //Lock
}
