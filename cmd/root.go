package cmd

import (
	"os"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gototp",
	Short: "Simple TOTP CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var flagGroup string

func getEffectiveGroup() (string, error) {
	effectiveGroup := flagGroup

	if effectiveGroup == "" {
		cfgGroup, err := storage.GetCurrentGroup()
		if err != nil {
			return "", nil
		}

		if cfgGroup != "" {
			effectiveGroup = cfgGroup
		} else {
			effectiveGroup = "default"
		}
	}
	return effectiveGroup, nil
}
