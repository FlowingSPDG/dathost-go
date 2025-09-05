package main

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/dathost-go"
	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	req := dathost.CreateGameServerRequest{
		AutoStop:          false,
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
		Name:              "test_API",
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
			MapsSource:                   "",
			MapGroup:                     "",
			MapGroupStartMap:             "",
			WorkshopCollectionID:         "",
			WorkshopCollectionStartMapID: "",
			WorkshopSingleMapID:          "",
			Password:                     "",
			RCON:                         "",
			Slots:                        10,
		},
	}

	server, err := client.CreateGameServer(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("server:", server)
}
