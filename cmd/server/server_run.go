package main

import (
	"fmt"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/Database-password-management-system/routing"
	"github.com/gorilla/mux"
)

// description of how to start a server
func main() {
	r := mux.NewRouter()

	*r = routing.SetupRouting(*r)

	data := server.Data{
		r,
		"8080",
		false,
		0,
	}
	// call the function based on the port needed and your own routing function
	err := server.Server(data)
	if err != nil {
		fmt.Println(err)
	}
}
