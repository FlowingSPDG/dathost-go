package dathost

import (
	"context"
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
	ListGameServers(ctx context.Context) ([]GameServer, error)
	CreateGameServer(ctx context.Context, req CreateGameServerRequest) (*GameServer, error)
	DeleteGameServer(ctx context.Context, id string) error
	GetGameServer(ctx context.Context, id string) (*GameServer, error)
	UpdateGameServer(ctx context.Context, id string, req CreateGameServerRequest) error
	GetGameServerMetrics(ctx context.Context, id string) (*GameServerMetrics, error)

	// Actions API
	DuplicateGameServer(ctx context.Context, id string) (*GameServer, error)
	ResetGameServer(ctx context.Context, id string) error
	StartGameServer(ctx context.Context, id string, data StartGameServerBody) error
	StopGameServer(ctx context.Context, id string) error
	SyncFilesGameServer(ctx context.Context, id string) error

	// File Management API
	ListFilesOnGameServer(ctx context.Context, id string, path string) ([]FileInfo, error)
	DeleteFilesFromGameServer(ctx context.Context, id string, path string) error
	DownloadFileFromGameServer(ctx context.Context, id string, path string) ([]byte, error)
	UploadFileToGameServer(ctx context.Context, id string, path string, data []byte) error
	MoveFileOnGameServer(ctx context.Context, id string, from string, to string) error
	RegenerateFTPPasswordForGameServer(ctx context.Context, id string) (*FTPPasswordResponse, error)
	UnzipFileOnGameServer(ctx context.Context, id string, path string) error

	// Console API
	GetLastLineFromConsole(ctx context.Context, id string, maxLines int) (*GetLastLineFromConsoleResponse, error)
	SendCommandToConsole(ctx context.Context, id string, command string) error

	// CS2 Matches API
	StartCS2Match(ctx context.Context, id string, req StartCS2MatchRequest) (*CS2Match, error)
	GetCS2Match(ctx context.Context, id string) (*CS2Match, error)
	CancelCS2Match(ctx context.Context, id string) error
	AddPlayerToCS2Match(ctx context.Context, id string, req AddPlayerToCS2MatchRequest) error

	// Account API
	GetCurrentAccount(ctx context.Context) (*Account, error)
	ListInvoices(ctx context.Context) ([]Invoice, error)
	GetInvoiceAsHTML(ctx context.Context, id string) (string, error)

	// System API
	GetCustomDomains(ctx context.Context) ([]CustomDomain, error)

	// Subscription API
	UpdateSubscription(ctx context.Context, id string, req UpdateSubscriptionRequest) error
}

func (dc *dathostClientv01) addHeader(req *http.Request) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", dc.auth)
}
