package cmd

import (
	"github.com/spf13/cobra"

	"dnsgo/bootstrap"
)

var HttpServerCmd = &cobra.Command{
	Use:   "http",
	Short: "start http server",
	Long:  "start http server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.HttpServer()
	},
}
