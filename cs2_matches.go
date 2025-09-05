package dathost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// StartCS2Match implements DatHostClientv01.
func (dc *dathostClientv01) StartCS2Match(ctx context.Context, id string, req StartCS2MatchRequest) (*CS2Match, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/cs2/matches", id)

	data := url.Values{}
	data.Set("config", req.Config)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", ep, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	dc.addHeader(httpReq)
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var match CS2Match
	if err := json.NewDecoder(res.Body).Decode(&match); err != nil {
		return nil, err
	}
	return &match, nil
}

// GetCS2Match implements DatHostClientv01.
func (dc *dathostClientv01) GetCS2Match(ctx context.Context, id string) (*CS2Match, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/cs2/matches", id)

	req, err := http.NewRequestWithContext(ctx, "GET", ep, nil)
	if err != nil {
		return nil, err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var match CS2Match
	if err := json.NewDecoder(res.Body).Decode(&match); err != nil {
		return nil, err
	}
	return &match, nil
}

// CancelCS2Match implements DatHostClientv01.
func (dc *dathostClientv01) CancelCS2Match(ctx context.Context, id string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/cs2/matches", id)

	req, err := http.NewRequestWithContext(ctx, "DELETE", ep, nil)
	if err != nil {
		return err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// AddPlayerToCS2Match implements DatHostClientv01.
func (dc *dathostClientv01) AddPlayerToCS2Match(ctx context.Context, id string, req AddPlayerToCS2MatchRequest) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/cs2/matches/players", id)

	data := url.Values{}
	data.Set("steam_id", req.SteamID)
	data.Set("team", req.Team)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", ep, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	dc.addHeader(httpReq)
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
