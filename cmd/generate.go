package cmd

import (
	"fmt"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/DWoodhouse22/gototp/totp"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate <name>",
	Short: "Generate a TOTP code",
	Args:  cobra.ExactArgs(1),
	Run:   Generate,
}

func Generate(cmd *cobra.Command, args []string) {
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
		fmt.Printf("no account found for '%s'\n", name)
		return
	}

	if len(accounts) == 1 {
		generateCode(accounts[0].Secret)
		return
	}

	fmt.Printf("Multiple accounts found for '%s':\n\n", name)
	for i, act := range accounts {
		fmt.Printf("%d) %s (%s)\n", i+1, act.Name, act.Group)
	}

	fmt.Print("\n Select an option: ")
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(accounts) {
		fmt.Println("Invalid selection")
		return
	}

	selected := accounts[choice-1]
	generateCode(selected.Secret)
}

func generateCode(secret string) {
	code, err := totp.Generate(secret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if copyFlag {
		if err := clipboard.WriteAll(code); err != nil {
			fmt.Println("Warning: failed to copy to clipboard:", err)
		} else {
			fmt.Println("(copied to clipboard)")
		}
	}
	fmt.Println(code)
}

func init() {
	generateCmd.Flags().StringVarP(&flagGroup, "group", "g", "", "Group name (default: 'default')")
	generateCmd.Flags().BoolVarP(&copyFlag, "copy", "c", false, "Copy code to clipboard")
	rootCmd.AddCommand(generateCmd)
}
