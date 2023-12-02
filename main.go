package main

import (
	"flag"

	"dnsgo/cmd"
)

func main() {
	flag.Parse()
	cmd.Execute(
		cmd.HttpServerCmd,
		cmd.DNSServerCmd,
	)
}
