package main

import (
	"os"

	"github.com/FlowingSPDG/dathost-go"
	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	req := dathost.CreateGameServerRequest{
		AddedVoiceServer:  "",
		AutoStop:          false,
		EnableSyntropy:    false,
		AutoStopMinutes:   0,
		Confirmed:         true,
		CustomDomain:      "",
		DeleteProtection:  false,
		EnableCoreDump:    false,
		EnableMySQL:       false,
		Game:              "cs2",
		Location:          "tokyo",
		ManualSortOrder:   0,
		MaxDiskUsageGb:    50,
		Name:              "test_API_Updated",
		PreferDedicated:   false,
		RebootOnCrash:     false,
		ScheduledCommands: "",
		ServerImage:       "default",
		UserData:          "",
		CS2Settings: dathost.CS2SettingsForm{
			SteamGameServerLoginToken:    "",
			DisableBots:                  false,
			EnableGOTV:                   true,
			EnableGOTVSecondary:          false,
			GameMode:                     "competitive",
			Insecure:                     false,
			MapsSource:                   "mapgroup",
			MapGroup:                     "",
			MapGroupStartMap:             "",
			WorkshopCollectionID:         "",
			WorkshopCollectionStartMapID: "",
			WorkshopSingleMapID:          "",
			Password:                     "lo3jp",
			RCON:                         "dathost",
			Slots:                        10,
			EnableMetamod:                true,
			MetamodPlugins:               []string{"654a32ea452c94f085961b91", "65538a41f7b98449a05fe3b9"}, // CounterStrikeSharp / MatchZy
		},
	}

	serverID := os.Args[1]

	if err := client.UpdateGameServer(serverID, req); err != nil {
		panic(err)
	}
}
