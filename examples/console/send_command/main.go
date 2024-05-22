package main

import (
	"os"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	serverID := os.Args[1]
	command := os.Args[2]

	if err := client.SendCommandToConsole(serverID, command); err != nil {
		panic(err)
	}
}
