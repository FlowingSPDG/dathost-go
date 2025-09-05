package dathost

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// DuplicateGameServer implements DatHostClientv01.
func (dc *dathostClientv01) DuplicateGameServer(ctx context.Context, id string) (*GameServer, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/duplicate", id)
	req, _ := http.NewRequestWithContext(ctx, "POST", ep, nil)
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

// ResetGameServer implements DatHostClientv01.
func (dc *dathostClientv01) ResetGameServer(ctx context.Context, id string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/reset", id)
	req, _ := http.NewRequestWithContext(ctx, "POST", ep, nil)
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

// StartGameServer implements DatHostClientv01.
func (dc *dathostClientv01) StartGameServer(ctx context.Context, id string, data StartGameServerBody) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/start", id)
	encoded := data.ToFormData().Encode()
	req, _ := http.NewRequestWithContext(ctx, "POST", ep, strings.NewReader(encoded))

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

// StopGameServer implements DatHostClientv01.
func (dc *dathostClientv01) StopGameServer(ctx context.Context, id string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/stop", id)
	req, _ := http.NewRequestWithContext(ctx, "POST", ep, nil)

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

// SyncFilesGameServer implements DatHostClientv01.
func (dc *dathostClientv01) SyncFilesGameServer(ctx context.Context, id string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/sync-files", id)
	req, _ := http.NewRequestWithContext(ctx, "POST", ep, nil)

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
