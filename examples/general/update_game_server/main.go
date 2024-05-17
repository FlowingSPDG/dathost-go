package main

import (
	"os"

	"github.com/FlowingSPDG/dathost-go"
	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	req := dathost.CreateGameServerRequest{
		AutoStop:                     false,
		AutoStopMinutes:              0,
		Confirmed:                    true,
		CustomDomain:                 "",
		DeleteProtection:             false,
		EnableCoreDump:               false,
		EnableMySQL:                  false,
		Game:                         "cs2",
		Location:                     "tokyo",
		ManualSortOrder:              0,
		MaxDiskUsageGb:               50,
		Name:                         "test_API_Updated",
		PreferDedicated:              false,
		RebootOnCrash:                false,
		ScheduledCommands:            "",
		ServerImage:                  "default",
		UserData:                     "",
		SteamGameServerLoginToken:    "2053B813D172D1AC4F3C840515E77531",
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
		Password:                     "lo3jp",
		RCON:                         "matukodx1",
		Slots:                        10,
	}

	serverID := os.Args[1]

	if err := client.UpdateGameServer(serverID, req); err != nil {
		panic(err)
	}
}
