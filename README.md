# Basic-Golang-Server


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


PS: Contributions and thoughts on how to improve/ add more features is appreciated!
