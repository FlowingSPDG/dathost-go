package dathost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"
)

// ListFilesOnGameServer implements DatHostClientv01.
func (dc *dathostClientv01) ListFilesOnGameServer(id string, path string) ([]FileInfo, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/files", id)
	u, _ := url.Parse(ep)

	q := u.Query()
	if path != "" {
		q.Set("path", path)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var files []FileInfo
	if err := json.NewDecoder(res.Body).Decode(&files); err != nil {
		return nil, err
	}
	return files, nil
}

// DeleteFilesFromGameServer implements DatHostClientv01.
func (dc *dathostClientv01) DeleteFilesFromGameServer(id string, path string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/files", id)
	u, _ := url.Parse(ep)

	q := u.Query()
	q.Set("path", path)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("DELETE", u.String(), nil)
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

// DownloadFileFromGameServer implements DatHostClientv01.
func (dc *dathostClientv01) DownloadFileFromGameServer(id string, path string) ([]byte, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/files", id)
	u, _ := url.Parse(ep)

	q := u.Query()
	q.Set("path", path)
	q.Set("download", "1")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UploadFileToGameServer implements DatHostClientv01.
func (dc *dathostClientv01) UploadFileToGameServer(id string, path string, data []byte) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/files", id)

	var b bytes.Buffer
	mw := multipart.NewWriter(&b)

	// Add path field
	mw.WriteField("path", path)

	// Add file field
	filename := filepath.Base(path)
	fw, err := mw.CreateFormFile("file", filename)
	if err != nil {
		return xerrors.Errorf("failed to create form file: %w", err)
	}

	if _, err := fw.Write(data); err != nil {
		return xerrors.Errorf("failed to write file data: %w", err)
	}

	contentType := mw.FormDataContentType()
	if err := mw.Close(); err != nil {
		return xerrors.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequest("POST", ep, &b)
	if err != nil {
		return err
	}
	dc.addHeader(req)
	req.Header.Set("Content-Type", contentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// MoveFileOnGameServer implements DatHostClientv01.
func (dc *dathostClientv01) MoveFileOnGameServer(id string, from string, to string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/files", id)

	data := url.Values{}
	data.Set("from", from)
	data.Set("to", to)

	req, err := http.NewRequest("PUT", ep, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	dc.addHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// RegenerateFTPPasswordForGameServer implements DatHostClientv01.
func (dc *dathostClientv01) RegenerateFTPPasswordForGameServer(id string) (*FTPPasswordResponse, error) {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/regenerate-ftp-password", id)

	req, err := http.NewRequest("POST", ep, nil)
	if err != nil {
		return nil, err
	}
	dc.addHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response FTPPasswordResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UnzipFileOnGameServer implements DatHostClientv01.
func (dc *dathostClientv01) UnzipFileOnGameServer(id string, path string) error {
	ep := fmt.Sprintf("https://dathost.net/api/0.1/game-servers/%s/unzip", id)

	data := url.Values{}
	data.Set("path", path)

	req, err := http.NewRequest("POST", ep, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	dc.addHeader(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
