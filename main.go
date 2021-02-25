package main

import (
	"github.com/AhhMonkeyDevs/discordbeat/cmd"
	_ "github.com/AhhMonkeyDevs/discordbeat/include"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
