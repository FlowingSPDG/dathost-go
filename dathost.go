package dathost

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// dathostClientv01 Client API v0.1
type dathostClientv01 struct {
	// TODO: 共通処理をプライベートなメソッドに切り出す
	auth string // Basic auth with Base64-encoded username/password pair
}

func NewDathostClientv01(username, password string) DatHostClientv01 {
	base64Auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return &dathostClientv01{
		auth: fmt.Sprintf("Basic %s", base64Auth),
	}
}

// DatHostClientv01 Client API v0.1
type DatHostClientv01 interface {
	ListGameServers() ([]GameServer, error)
	CreateGameServer(req *CreateGameServerRequest) (*GameServer, error)
	DeleteGameServer(id string) error
	GetGameServer(id string) (*GameServer, error)
	UpdateGameServer(id string, req *CreateGameServerRequest) error
	GetGameServerMetrics(id string) (*GameServerMetrics, error)
	// UpdateSubscription(id string) // https://dathost.readme.io/reference/post_api-0-1-game-servers-server-id-subscription
}

func (dc *dathostClientv01) addHeader(req *http.Request) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", dc.auth)
}

// ListGameServers implements DatHostClientv01.
func (dc *dathostClientv01) ListGameServers() ([]GameServer, error) {
	ep := "https://dathost.net/api/0.1/game-servers"
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var servers []GameServer
	if err := json.NewDecoder(res.Body).Decode(&servers); err != nil {
		return nil, err
	}
	return servers, nil
}

// CreateGameServer implements DatHostClientv01.
func (dc *dathostClientv01) CreateGameServer(data *CreateGameServerRequest) (*GameServer, error) {
	ep := "https://dathost.net/api/0.1/game-servers"
	encoded := data.ToFormData().Encode()
	req, _ := http.NewRequest("POST", ep, strings.NewReader(encoded))

	dc.addHeader(req)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var server *GameServer
	if err := json.NewDecoder(res.Body).Decode(&server); err != nil {
		return nil, err
	}
	return server, nil
}

// DeleteGameServer implements DatHostClientv01.
func (dc *dathostClientv01) DeleteGameServer(id string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s", id)
	req, _ := http.NewRequest("DELETE", ep, nil)

	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	_ = body

	return nil
}

// GetGameServer implements DatHostClientv01.
func (dc *dathostClientv01) GetGameServer(id string) (*GameServer, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s", id)
	req, _ := http.NewRequest("GET", ep, nil)
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var server *GameServer
	if err := json.NewDecoder(res.Body).Decode(&server); err != nil {
		return nil, err
	}
	return server, nil
}

// UpdateGameServer implements DatHostClientv01.
func (dc *dathostClientv01) UpdateGameServer(id string, data *CreateGameServerRequest) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s", id)
	encoded := data.ToFormData().Encode()
	req, _ := http.NewRequest("PUT", ep, strings.NewReader(encoded))
	dc.addHeader(req)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	_ = body

	return nil
}

// GetGameServerMetrics implements DatHostClientv01.
func (dc *dathostClientv01) GetGameServerMetrics(id string) (*GameServerMetrics, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/metrics", id)
	req, _ := http.NewRequest("GET", ep, nil)
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var metrics *GameServerMetrics
	if err := json.NewDecoder(res.Body).Decode(&metrics); err != nil {
		return nil, err
	}
	return metrics, nil
}
