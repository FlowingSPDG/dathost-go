package main

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	servers, err := client.ListGameServers(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("servers:", servers)
}
