package proxy

import (
	"github.com/gorilla/mux"
)

// Server holds the data about a proxy server
type Server struct {
	Routes   map[string]string // Hash of server and port mapping
	PortList map[string]int    // Port existance mapping
	Router   *mux.Router       // The proxy router
}
