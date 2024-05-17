package main

import (
	"fmt"

	"github.com/FlowingSPDG/dathost-go"
)

func main() {
	client := dathost.NewDathostClientv01("", "")

	const cs2token = "YOUR_STEAM_GAMESERVER_TOKEN"
	const serverName = "testserver"
	res, err := client.CreateGameServer(serverName, cs2token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
