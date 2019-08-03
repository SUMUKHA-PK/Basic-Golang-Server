package responses

import "github.com/SUMUKHA-PK/Basic-Golang-Server/proxy"

// AddServerResponse is the JSON output for /addServer
type AddServerResponse struct {
	StatusCode     int
	SuccessMessage string
	Data           proxy.AddServerRes
}
