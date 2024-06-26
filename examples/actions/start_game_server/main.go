package main

import (
	"os"

	"github.com/FlowingSPDG/dathost-go"
	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	req := dathost.StartGameServerBody{
		AllowHostReassignment: true,
	}

	serverID := os.Args[1]

	if err := client.StartGameServer(serverID, req); err != nil {
		panic(err)
	}
}
