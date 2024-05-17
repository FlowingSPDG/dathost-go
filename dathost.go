package dathost

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://dathost.net/api/0.1"

// dathostClientv01 Client API v0.1
type dathostClientv01 struct {
	auth string // Basic auth with Base64-encoded username/password pair
}

// ListGameServers implements DatHostClientv01.
func (dc *dathostClientv01) ListGameServers() {
	ep := baseURL + "/game-servers"
	body, err := dc.sendRequest("GET", ep, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// sendRequest sends an HTTP request to the given endpoint with the specified method and body.
func (dc *dathostClientv01) sendRequest(method, endpoint string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", dc.auth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// NewDathostClientv01 creates a new instance of DatHostClientv01.
func NewDathostClientv01(username, password string) DatHostClientv01 {
	base64Auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return &dathostClientv01{
		auth: fmt.Sprintf("Basic %s", base64Auth),
	}
}

// DatHostClientv01 Client API v0.1
type DatHostClientv01 interface {
	ListGameServers() // WIP: types
}
