package cmd

import (
	"github.com/d3ta-go/app-restapi-composite/cmd/db"
	"github.com/d3ta-go/app-restapi-composite/cmd/server"
)

func init() {
	RootCmd.AddCommand(db.DBCmd)
	RootCmd.AddCommand(server.ServerCmd)
	// Add your custom command
}
