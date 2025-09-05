package dathost

import (
	"encoding/base64"
	"fmt"
	"net/http"
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
	// General APIs
	ListGameServers() ([]GameServer, error)
	CreateGameServer(req CreateGameServerRequest) (*GameServer, error)
	DeleteGameServer(id string) error
	GetGameServer(id string) (*GameServer, error)
	UpdateGameServer(id string, req CreateGameServerRequest) error
	GetGameServerMetrics(id string) (*GameServerMetrics, error)
	// UpdateSubscription(id string) // https://dathost.readme.io/reference/post_api-0-1-game-servers-server-id-subscription // TODO.

	// Actions API
	DuplicateGameServer(id string) (*GameServer, error)
	ResetGameServer(id string) error
	StartGameServer(id string, data StartGameServerBody) error
	StopGameServer(id string) error
	SyncFilesGameServer(id string) error

	// File Management API
	ListFilesOnGameServer(id string, path string) ([]FileInfo, error)
	DeleteFilesFromGameServer(id string, path string) error
	DownloadFileFromGameServer(id string, path string) ([]byte, error)
	UploadFileToGameServer(id string, path string, data []byte) error
	MoveFileOnGameServer(id string, from string, to string) error
	RegenerateFTPPasswordForGameServer(id string) (*FTPPasswordResponse, error)
	UnzipFileOnGameServer(id string, path string) error

	// // Console API
	GetLastLineFromConsole(id string, maxLines int) (*GetLastLineFromConsoleResponse, error)
	SendCommandToConsole(id string, command string) error

	// CS2 Matches API
	StartCS2Match(id string, req StartCS2MatchRequest) (*CS2Match, error)
	GetCS2Match(id string) (*CS2Match, error)
	CancelCS2Match(id string) error
	AddPlayerToCS2Match(id string, req AddPlayerToCS2MatchRequest) error

	// Account API
	GetCurrentAccount() (*Account, error)
	ListInvoices() ([]Invoice, error)
	GetInvoiceAsHTML(id string) (string, error)

	// System API
	GetCustomDomains() ([]CustomDomain, error)

	// Subscription API
	UpdateSubscription(id string, req UpdateSubscriptionRequest) error
}

func (dc *dathostClientv01) addHeader(req *http.Request) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", dc.auth)
}
