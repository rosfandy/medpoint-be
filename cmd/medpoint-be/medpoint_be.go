package main

import (
	"medpointbe/internal/bootstrap"

	"github.com/sev-2/raiden"
)

func main() {
	config, err := raiden.LoadConfig(nil)
	if err != nil {
		raiden.Error("load configuration", err.Error())
	}

	server := raiden.NewServer(config)

	bootstrap.RegisterRoute(server)

	server.Run()
}
