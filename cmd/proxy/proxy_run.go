package main

import (
	"fmt"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/proxy"
)

func main() {
	proxy, err := proxy.StartProxy("8080")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(proxy.PortList["8080"])
}
