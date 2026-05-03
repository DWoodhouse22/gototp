package cmd

import (
	"fmt"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove a registered account",
	Args:  cobra.ExactArgs(1),
	Run:   Remove,
}

func Remove(cmd *cobra.Command, args []string) {
	name := args[0]
	effectiveGroup, err := getEffectiveGroup()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	accounts, err := storage.FindAccounts(name, effectiveGroup)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(accounts) == 0 {
		fmt.Printf("'%s' (%s) not found\n", name, effectiveGroup)
		return
	}

	if !confirm(fmt.Sprintf("Remove '%s' (%s)? (y/N): ", name, effectiveGroup)) {
		fmt.Println("Cancelled")
		return
	}

	if err := storage.RemoveAccount(accounts[0]); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("'%s' (%s) removed\n", name, effectiveGroup)
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
