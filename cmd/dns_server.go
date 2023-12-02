package cmd

import (
	"github.com/spf13/cobra"

	"dnsgo/bootstrap"
)

var DNSServerCmd = &cobra.Command{
	Use:   "dns",
	Short: "start dns server",
	Long:  "start dns server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.DNSServer()
	},
}
