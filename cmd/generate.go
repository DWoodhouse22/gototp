package cmd

import (
	"fmt"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/DWoodhouse22/gototp/totp"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate name",
	Short: "Generate a TOTP code",
	Args:  cobra.ExactArgs(1),
	Run:   Generate,
}

func Generate(cmd *cobra.Command, args []string) {
	name := args[0]

	secret, err := storage.LoadAccount(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	code, err := totp.Generate(secret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(code)
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
