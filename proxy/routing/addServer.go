package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/proxy/responses"

	random "math/rand"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/proxy"
	"github.com/SUMUKHA-PK/Database-password-management-system/crypto"
	"github.com/SUMUKHA-PK/Database-password-management-system/util"
)

// AddServerToProxy adds a server to one of the routes of the proxy.
// It DOES NOT create a server. The already running server is added. (Checked by /healthCheck)
// Function writes a JSON response with a "hash" for the new server
// added. This is the route to which the proxy sends the requests
// to for this server. It checks for already existing ports on the
// proxy and rejects accordingly if a new server is requested on it.
func AddServerToProxy(w http.ResponseWriter, r *http.Request) {

	log.Println("/addServer request received")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in proxy/routing/addServer.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newReq proxy.AddServerReq
	if err = json.Unmarshal(body, &newReq); err != nil {
		log.Printf("Coudn't Unmarshal data in proxy/routing/addServer.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxy.ProxyServer.Mutex.Lock()
	hash := getHash(newReq.IP, proxy.ProxyServer)
	proxy.ProxyServer.Mutex.Unlock()

	proxy.ProxyServer.Mutex.Lock()
	// Add this to the routing of the proxy
	proxy.ProxyServer.Routes[hash] = newReq.IP
	proxy.ProxyServer.Mutex.Unlock()

	outData := &responses.AddServerResponse{200, "Successfully added server to proxy", proxy.AddServerRes{Hash: hash}}
	outJSON, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in proxy/routing/addServer.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}

// getHash checks for existing hashes
func getHash(IP string, server proxy.Server) string {
	var hash string
	for {
		hash = crypto.CreateMD5Hash(util.StringWithCharset(random.Intn(20)+1, util.Charset) + IP)
		if _, ok := server.Routes[hash]; !ok {
			server.Routes[hash] = IP
			break
		}
	}
	return hash
}
