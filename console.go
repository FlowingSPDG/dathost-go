package dathost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

// SendCommandToConsole implements DatHostClientv01.
func (dc *dathostClientv01) SendCommandToConsole(id string, command string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/console", id)
	b := &bytes.Buffer{}
	mw := multipart.NewWriter(b)
	mw.WriteField("line", command)
	contentType := mw.FormDataContentType()
	mw.Close()

	req, _ := http.NewRequest("POST", ep, b)
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

type GetLastLineFromConsoleResponse struct {
	Lines []string `json:"lines"`
}

// GetLastLineFromConsole implements DatHostClientv01.
func (dc *dathostClientv01) GetLastLineFromConsole(id string, maxLines int) (*GetLastLineFromConsoleResponse, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/console", id)
	u, _ := url.Parse(ep)

	q := u.Query()
	q.Set("__a", "1")
	q.Set("max_lines", strconv.Itoa(maxLines))
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resp := GetLastLineFromConsoleResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
