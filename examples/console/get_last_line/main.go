package main

import (
	"fmt"
	"os"

	"github.com/FlowingSPDG/dathost-go/examples/common"
)

func main() {
	client := common.MustGetClient()

	serverID := os.Args[1]

	resp, err := client.GetLastLineFromConsole(serverID, 100)
	if err != nil {
		panic(err)
	}
	for i, line := range resp.Lines {
		fmt.Printf("LINE[%d]: %s\n", i+1, line)
	}
}
