package main

import (
	"os"

	"github.com/AhhMonkeyDevs/discordbeat/cmd"

	_ "github.com/AhhMonkeyDevs/discordbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
