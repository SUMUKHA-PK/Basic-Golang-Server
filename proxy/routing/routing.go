package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouting adds all the routes
func SetupRouting(r mux.Router) mux.Router {
	r.HandleFunc("/addServer", AddServerToProxy).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/route/{id:[0-9a-zA-Z]+}", RouteForward)
	return r
}
