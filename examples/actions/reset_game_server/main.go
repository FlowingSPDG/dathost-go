package main

import (
	"os"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()
	serverID := os.Args[1]

	if err := client.ResetGameServer(serverID); err != nil {
		panic(err)
	}
}
