package dathost

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

// dathostClientv01 Client API v0.1
type dathostClientv01 struct {
	auth string // Basic auth with Base64-encoded username/password pair
}

// ListGameServers implements DatHostClientv01.
func (dc *dathostClientv01) ListGameServers() {
	ep := "https://dathost.net/api/0.1/game-servers"
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", dc.auth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

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
