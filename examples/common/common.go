package common

import (
	"os"

	"github.com/FlowingSPDG/dathost-go"
)

// common functions for examples

func MustGetClient() dathost.DatHostClientv01 {
	username := os.Getenv("DATHOST_API_USERNAME")
	if username == "" {
		panic("DATHOST_API_USERNAME not set")
	}
	password := os.Getenv("DATHOST_API_PASSWORD")
	if password == "" {
		panic("DATHOST_API_PASSWORD not set")
	}
	client := dathost.NewDathostClientv01(username, password)
	return client
}



