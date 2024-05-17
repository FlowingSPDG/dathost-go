package main

import "github.com/FlowingSPDG/dathost-go"

func main() {
	client := dathost.NewDathostClientv01("", "")

	client.ListGameServers()
}
