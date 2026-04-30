package cmd

import (
	"fmt"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered accounts",
	Run:   List,
}

func List(cmd *cobra.Command, args []string) {
	accounts, err := storage.ListAccounts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, name := range accounts {
		fmt.Println(name)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
