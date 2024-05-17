package main

import (
	"fmt"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	servers, err := client.ListGameServers()
	if err != nil {
		panic(err)
	}
	fmt.Println("servers:", servers)
}
