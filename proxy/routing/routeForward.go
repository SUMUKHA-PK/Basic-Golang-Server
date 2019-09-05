package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/proxy"
)

// RouteForward forwards the request to the server
// and provides its response back to the client
func RouteForward(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	url = strings.ReplaceAll(url, "/route/", "")
	IP := strings.Split(url, "/")
	url = IP[0]
	forwardingIP := proxy.ProxyServer.RouteMap.Routes[url]
	forwardingIP += "/" + IP[1]
	fmt.Println(forwardingIP)

	go keepConnectionOpen(forwardingIP)
	return
}

func keepConnectionOpen(forwardingIP string) {

}
