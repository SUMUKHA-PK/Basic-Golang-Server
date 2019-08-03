package main

import (
	"fmt"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/proxy"
	"github.com/SUMUKHA-PK/Basic-Golang-Server/proxy/routing"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	*r = routing.SetupRouting(*r)

	proxyServer, err := proxy.StartProxy(r, "8080")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(proxyServer)
}
