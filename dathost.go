package dathost

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseURL = "https://dathost.net/api/0.1"

// dathostClientv01 Client API v0.1
type dathostClientv01 struct {
	auth string // Basic auth with Base64-encoded username/password pair
}

// ListGameServers implements DatHostClientv01.
func (dc *dathostClientv01) ListGameServers() (string, error) {
	ep := baseURL + "/game-servers"
	body, err := dc.sendRequest("GET", ep, nil, "application/json")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}

func (dc *dathostClientv01) CreateGameServer(serverName string, cs2token string) (string, error) {
	ep := baseURL + "/game-servers"
	// https://dathost.readme.io/reference/post_game_servers
	payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"game\"\r\n\r\ncs2\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"name\"\r\n\r\n" + serverName + "\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"csgo_settings.steam_game_server_login_token\"\r\n\r\n" + cs2token + "\r\n-----011000010111000001101001--")
	body, err := dc.sendRequest("POST", ep, payload, "multipart/form-data")
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (dc *dathostClientv01) DeleteGameServer(serverId string) (string, error) {
	ep := baseURL + "/game-servers/" + serverId
	body, err := dc.sendRequest("DELETE", ep, nil, "application/json")
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// sendRequest sends an HTTP request to the given endpoint with the specified method and body.
func (dc *dathostClientv01) sendRequest(method, endpoint string, body io.Reader, acceptType string) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", acceptType)
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
	ListGameServers() (string, error)
	CreateGameServer(serverName string, cs2token string) (string, error)
	DeleteGameServer(serverId string) (string, error)
}
