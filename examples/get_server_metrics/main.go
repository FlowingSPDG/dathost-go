package main

import (
	"fmt"
	"os"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()
	serverID := os.Args[1]

	server, err := client.GetGameServerMetrics(serverID)
	if err != nil {
		panic(err)
	}
	fmt.Println("server:", server)
}
