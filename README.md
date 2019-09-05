# Basic-Golang-Server AND Proxy (Reverse)


## 1. The simple server
This is a simple Go server that can be the skeleton for all your servers you build.

Key features:
* Uses gorilla mux for routing. 
* Can provide custom routing functions which are of type \*mux.Router
* Can provide custom ports to create server.
* Uses go.mod, so just include *github.com/SUMUKHA-PK/Basic-Golang-Server* in your imports to get GOing.

Usage:

```
import (
  .
  .
  Server "github.com/SUMUKHA-PK/Basic-Golang-Server/server"
  .
  .
)

func main(){
  .
  .
  err:= server.Server(routing_function, port_number)  //Server starts!
  //Handle err accordingly.
  .
  .
}
```

## 2. The proxy :


The idea of the proxy (reverse proxy) is to be able to host it as a separate server with its own routing.
The server has separate API end points which when hit will allow us to add or remove servers from the 
proxy!


Since they all are independently managed, all servers can be independently started or shut down and added to 
the proxy easily using some simple code, which can be in any language! (Or just cURL or wget commands)

Further plans are to provide support to the proxy and the servers which enable the the servers to accept the
request ONLY through the proxy to increase security in the servers.



PS: Contributions and thoughts on how to improve/ add more features is appreciated!

