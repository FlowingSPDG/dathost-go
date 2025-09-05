package main

import (
	"context"
	"os"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()
	serverID := os.Args[1]

	if err := client.DeleteGameServer(context.Background(), serverID); err != nil {
		panic(err)
	}
}



