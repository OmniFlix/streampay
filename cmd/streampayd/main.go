package main

import (
	"os"

	"github.com/OmniFlix/streampay/app"
	"github.com/OmniFlix/streampay/cmd/streampayd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
