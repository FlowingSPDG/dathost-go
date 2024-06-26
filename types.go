package dathost

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"net/url"
)

type CS2Settings struct {
	Slots                        int      `json:"slots"`
	SteamGameServerLoginToken    string   `json:"steam_game_server_login_token"`
	Rcon                         string   `json:"rcon"`
	Password                     string   `json:"password"`
	MapsSource                   string   `json:"maps_source"`
	Mapgroup                     string   `json:"mapgroup"`
	MapgroupStartMap             string   `json:"mapgroup_start_map"`
	WorkshopCollectionID         string   `json:"workshop_collection_id"`
	WorkshopCollectionStartMapID string   `json:"workshop_collection_start_map_id"`
	WorkshopSingleMapID          string   `json:"workshop_single_map_id"`
	Insecure                     bool     `json:"insecure"`
	EnableGotv                   bool     `json:"enable_gotv"`
	EnableGotvSecondary          bool     `json:"enable_gotv_secondary"`
	DisableBots                  bool     `json:"disable_bots"`
	GameMode                     string   `json:"game_mode"`
	EnableMetamod                bool     `json:"enable_metamod"`
	MetamodPlugins               []string `json:"metamod_plugins"`
}

type Ports struct {
	Game          int `json:"game"`
	Gotv          int `json:"gotv"`
	GotvSecondary int `json:"gotv_secondary"`
}

type GameServer struct {
	ID                                string      `json:"id"`
	CreatedAt                         int         `json:"created_at"`
	Name                              string      `json:"name"`
	UserData                          any         `json:"user_data"`
	Game                              string      `json:"game"`
	Location                          string      `json:"location"`
	PlayersOnline                     int         `json:"players_online"`
	Status                            []any       `json:"status"` // ?
	Booting                           bool        `json:"booting"`
	ServerError                       string      `json:"server_error"`
	IP                                string      `json:"ip"`
	RawIP                             string      `json:"raw_ip"`
	PrivateIP                         any         `json:"private_ip"` // ?
	MatchID                           any         `json:"match_id"`   // ?
	On                                bool        `json:"on"`
	Ports                             Ports       `json:"ports"`
	Confirmed                         bool        `json:"confirmed"`
	MaxDiskUsageGb                    int         `json:"max_disk_usage_gb"`
	CostPerHour                       float64     `json:"cost_per_hour"`
	MaxCostPerHour                    float64     `json:"max_cost_per_hour"`
	MonthCredits                      float64     `json:"month_credits"`
	MonthResetAt                      int         `json:"month_reset_at"`
	MaxCostPerMonth                   float64     `json:"max_cost_per_month"`
	SubscriptionCycleMonths           int         `json:"subscription_cycle_months"`
	SubscriptionState                 string      `json:"subscription_state"`
	SubscriptionRenewalFailedAttempts int         `json:"subscription_renewal_failed_attempts"`
	SubscriptionRenewalNextAttemptAt  int         `json:"subscription_renewal_next_attempt_at"`
	CycleMonths1DiscountPercentage    int         `json:"cycle_months_1_discount_percentage"`
	CycleMonths3DiscountPercentage    int         `json:"cycle_months_3_discount_percentage"`
	CycleMonths12DiscountPercentage   int         `json:"cycle_months_12_discount_percentage"`
	FirstMonthDiscountPercentage      int         `json:"first_month_discount_percentage"`
	EnableMysql                       bool        `json:"enable_mysql"`
	Autostop                          bool        `json:"autostop"`
	AutostopMinutes                   int         `json:"autostop_minutes"`
	EnableCoreDump                    bool        `json:"enable_core_dump"`
	PreferDedicated                   bool        `json:"prefer_dedicated"`
	EnableSyntropy                    bool        `json:"enable_syntropy"`
	ServerImage                       string      `json:"server_image"`
	RebootOnCrash                     bool        `json:"reboot_on_crash"`
	ManualSortOrder                   int         `json:"manual_sort_order"`
	MysqlUsername                     string      `json:"mysql_username"`
	MysqlPassword                     string      `json:"mysql_password"`
	FtpPassword                       string      `json:"ftp_password"`
	DiskUsageBytes                    int         `json:"disk_usage_bytes"`
	DefaultFileLocations              any         `json:"default_file_locations"` // ?
	CustomDomain                      string      `json:"custom_domain"`
	ScheduledCommands                 []any       `json:"scheduled_commands"`      // ?
	AddedVoiceServer                  any         `json:"added_voice_server"`      // ?
	DuplicateSourceServer             any         `json:"duplicate_source_server"` // ?
	DeletionProtection                bool        `json:"deletion_protection"`
	OngoingMaintenance                bool        `json:"ongoing_maintenance"`
	ArkSettings                       any         `json:"ark_settings"` // ?
	Cs2Settings                       CS2Settings `json:"cs2_settings"`
	CsgoSettings                      any         `json:"csgo_settings"`            // ?
	EnshroudedSettings                any         `json:"enshrouded_settings"`      // ?
	PalworldSettings                  any         `json:"palworld_settings"`        // ?
	SonsoftheforestSettings           any         `json:"sonsoftheforest_settings"` // ?
	Teamfortress2Settings             any         `json:"teamfortress2_settings"`   // ?
	Teamspeak3Settings                any         `json:"teamspeak3_settings"`      // ?
	ValheimSettings                   any         `json:"valheim_settings"`         // ?
}

type CreateGameServerRequest struct {
	AddedVoiceServer  string
	AutoStop          bool
	AutoStopMinutes   int
	Confirmed         bool // ?
	CustomDomain      string
	DeleteProtection  bool
	EnableCoreDump    bool
	EnableMySQL       bool
	EnableSyntropy    bool
	Game              string // cs2. required
	Location          string // https://dathost.net/reference/server-locations-mapping
	ManualSortOrder   int
	MaxDiskUsageGb    int
	Name              string // required.
	PreferDedicated   bool
	RebootOnCrash     bool
	ScheduledCommands string // ?
	ServerImage       string // default is "default", or "ubuntu-20.04"
	UserData          string // custom metadata
	CS2Settings       CS2SettingsForm
}

type CS2SettingsForm struct {
	// CS2
	SteamGameServerLoginToken    string
	DisableBots                  bool
	EnableGOTV                   bool
	EnableGOTVSecondary          bool
	GameMode                     string // competitive
	Insecure                     bool
	MapsSource                   string // mapgroup
	MapGroup                     string
	MapGroupStartMap             string
	WorkshopCollectionID         string
	WorkshopCollectionStartMapID string
	WorkshopSingleMapID          string
	Password                     string
	RCON                         string
	Slots                        int
}

func (cgsr *CreateGameServerRequest) ToFormData(b *bytes.Buffer) string {
	mw := multipart.NewWriter(b)

	mw.WriteField("added_voice_server", cgsr.AddedVoiceServer)
	mw.WriteField("auto_stop", fmt.Sprintf("%t", cgsr.AutoStop))
	mw.WriteField("auto_stop_minutes", fmt.Sprintf("%d", cgsr.AutoStopMinutes))
	mw.WriteField("confirmed", fmt.Sprintf("%t", cgsr.Confirmed))
	mw.WriteField("custom_domain", cgsr.CustomDomain)
	mw.WriteField("delete_protection", fmt.Sprintf("%t", cgsr.DeleteProtection))
	mw.WriteField("enable_core_dump", fmt.Sprintf("%t", cgsr.EnableCoreDump))
	mw.WriteField("enable_mysql", fmt.Sprintf("%t", cgsr.EnableMySQL))
	mw.WriteField("enable_syntropy", fmt.Sprintf("%t", cgsr.EnableSyntropy))
	mw.WriteField("game", cgsr.Game)
	mw.WriteField("location", cgsr.Location)
	mw.WriteField("manual_sort_order", fmt.Sprintf("%d", cgsr.ManualSortOrder))
	mw.WriteField("max_disk_usage_gb", fmt.Sprintf("%d", cgsr.MaxDiskUsageGb))
	mw.WriteField("name", cgsr.Name)
	mw.WriteField("prefer_dedicated", fmt.Sprintf("%t", cgsr.PreferDedicated))
	mw.WriteField("reboot_on_crash", fmt.Sprintf("%t", cgsr.RebootOnCrash))
	mw.WriteField("scheduled_commands", cgsr.ScheduledCommands)
	mw.WriteField("server_image", cgsr.ServerImage)
	mw.WriteField("user_data", cgsr.UserData)

	// CS2
	mh := make(textproto.MIMEHeader)
	mh.Set("content-type", "multipart/form-data")
	mw.CreatePart(mh)
	mw.WriteField("cs2_settings.steam_game_server_login_token", cgsr.CS2Settings.SteamGameServerLoginToken)
	mw.WriteField("cs2_settings.disable_bots", fmt.Sprintf("%t", cgsr.CS2Settings.DisableBots))
	mw.WriteField("cs2_settings.enable_gotv", fmt.Sprintf("%t", cgsr.CS2Settings.EnableGOTV))
	mw.WriteField("cs2_settings.enable_gotv_secondary", fmt.Sprintf("%t", cgsr.CS2Settings.EnableGOTVSecondary))
	mw.WriteField("cs2_settings.game_mode", cgsr.CS2Settings.GameMode)
	mw.WriteField("cs2_settings.insecure", fmt.Sprintf("%t", cgsr.CS2Settings.Insecure))
	mw.WriteField("cs2_settings.maps_source", cgsr.CS2Settings.MapsSource)
	mw.WriteField("cs2_settings.mapgroup", cgsr.CS2Settings.MapGroup)
	mw.WriteField("cs2_settings.mapgroup_start_map", cgsr.CS2Settings.MapGroupStartMap)
	mw.WriteField("cs2_settings.workshop_collection_id", cgsr.CS2Settings.WorkshopCollectionID)
	mw.WriteField("cs2_settings.workshop_collection_start_map_id", cgsr.CS2Settings.WorkshopCollectionStartMapID)
	mw.WriteField("cs2_settings.workshop_single_map_id", cgsr.CS2Settings.WorkshopSingleMapID)
	mw.WriteField("cs2_settings.password", cgsr.CS2Settings.Password)
	mw.WriteField("cs2_settings.rcon", cgsr.CS2Settings.RCON)
	mw.WriteField("cs2_settings.slots", fmt.Sprintf("%d", cgsr.CS2Settings.Slots))

	contentType := mw.FormDataContentType()

	mw.Close()

	return contentType

}

type PlayerOnlineGraph struct {
	Timestamp int `json:"timestamp"`
	Value     int `json:"value"`
}

type MapPlayed struct {
	Map     string  `json:"map"`
	Seconds float64 `json:"seconds"`
}

type GameServerMetrics struct {
	PlayersOnlineGraph []PlayerOnlineGraph `json:"players_online_graph"`
	PlayersOnline      []any               `json:"players_online"`   // TODO...
	AllTimePlayers     []any               `json:"all_time_players"` // TODO...
	MapsPlayed         []MapPlayed         `json:"maps_played"`
}

type StartGameServerBody struct {
	AllowHostReassignment bool
}

func (sgsb *StartGameServerBody) ToFormData() *url.Values {
	ret := &url.Values{}
	ret.Add("allow_host_reassignment", fmt.Sprintf("%t", sgsb.AllowHostReassignment))
	return ret
}
