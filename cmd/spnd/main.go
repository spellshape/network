package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/spellshape/network/app"
	"github.com/spellshape/network/cmd"
	spntypes "github.com/spellshape/network/pkg/types"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		spntypes.Name,
		spntypes.AccountAddressPrefix,
		app.DefaultNodeHome,
		spntypes.DefaultChainID,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
