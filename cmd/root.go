package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "dnsgo",
	Short: "dnsgo is a demo",
	Long:  "dnsgo is a demo",
}

func Execute(cmd ...*cobra.Command) {
	if cmd != nil {
		rootCmd.AddCommand(cmd...)
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
