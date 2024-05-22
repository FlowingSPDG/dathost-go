package dathost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
func (dc *dathostClientv01) CreateGameServer(data CreateGameServerRequest) (*GameServer, error) {
	ep := "https://dathost.net/api/0.1/game-servers"
	b := &bytes.Buffer{}
	contentType := data.ToFormData(b)

	req, _ := http.NewRequest("PUT", ep, b)
	dc.addHeader(req)
	req.Header.Add("Content-Type", contentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var server GameServer
	if err := json.NewDecoder(res.Body).Decode(&server); err != nil {
		return nil, err
	}
	return &server, nil
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

	var server GameServer
	if err := json.NewDecoder(res.Body).Decode(&server); err != nil {
		return nil, err
	}
	return &server, nil
}

// UpdateGameServer implements DatHostClientv01.
func (dc *dathostClientv01) UpdateGameServer(id string, data CreateGameServerRequest) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s", id)
	b := &bytes.Buffer{}
	contentType := data.ToFormData(b)
	fmt.Println("contentType:", contentType)

	req, _ := http.NewRequest("PUT", ep, b)
	dc.addHeader(req)
	req.Header.Add("Content-Type", contentType)

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

	var metrics GameServerMetrics
	if err := json.NewDecoder(res.Body).Decode(&metrics); err != nil {
		return nil, err
	}
	return &metrics, nil
}
