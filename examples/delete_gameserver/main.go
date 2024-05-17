package main

import (
	"fmt"

	"github.com/FlowingSPDG/dathost-go"
)

func main() {
	client := dathost.NewDathostClientv01("", "")

	const serverId = "1234567890"
	res, err := client.DeleteGameServer(serverId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
