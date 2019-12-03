package main

import (
	"log"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/gorilla/mux"
)

// description of how to start a server
func main() {
	r := mux.NewRouter()
	m := make(map[string]int)
	// r = routing.SetupRouting(r)
	serverData := server.Data{
		Router:        r,
		Port:          "5555",
		IP:            "127.0.0.1",
		HTTPS:         false,
		ConnectionMap: m,
	}

	err := server.Server(&serverData)
	if err != nil {
		log.Fatalf("Could not start server : %v", err)
	}
}
