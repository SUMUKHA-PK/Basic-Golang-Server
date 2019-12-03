package server

import "github.com/gorilla/mux"

//Data is the current state of the server
type Data struct {
	Router        *mux.Router
	IP            string
	Port          string
	HTTPS         bool
	ConnectionMap map[string]int
	Count         int
}
