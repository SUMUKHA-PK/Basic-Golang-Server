package server

import "github.com/gorilla/mux"

//Data is the current state of the server
type Data struct {
	Router *mux.Router
	Port   string
	HTTPS  bool
	// Map   map[string]int
	Count int
}