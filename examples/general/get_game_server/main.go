package main

import (
	"context"
	"fmt"
	"os"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()
	serverID := os.Args[1]

	server, err := client.GetGameServer(context.Background(), serverID)
	if err != nil {
		panic(err)
	}
	fmt.Println("server:", server)
}



