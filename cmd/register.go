package cmd

import (
	"fmt"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register <name> <secret>",
	Short: "Register a TOTP secret",
	Args:  cobra.ExactArgs(2),
	Run:   Register,
}

func Register(cmd *cobra.Command, args []string) {
	name := args[0]
	secret := args[1]
	effectiveGroup, err := getEffectiveGroup()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = storage.SaveAccount(name, secret, effectiveGroup)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Secret saved")
}

func init() {
	registerCmd.Flags().StringVarP(&flagGroup, "group", "g", "", "Group name (default: 'default')")
	rootCmd.AddCommand(registerCmd)
}
