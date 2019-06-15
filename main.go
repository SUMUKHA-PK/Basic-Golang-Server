package main

import (
	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/Database-password-management-system/routing"
	"github.com/gorilla/mux"
)

// description of how to start a server
func main() {
	r := mux.NewRouter()

	*r = routing.SetupRouting(*r)

	// call the function based on the port needed and your own routing function
	server.Server(r, "10000")
}
