package cmd

import (
	"fmt"
	"sort"

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
	groups := make(map[string][]string)

	for _, act := range accounts {
		group := act.Group
		if group == "" {
			group = "default"
		}
		groups[group] = append(groups[group], act.Name)
	}

	for g, list := range groups {
		sort.Strings(list)
		groups[g] = list
	}

	for group, names := range groups {
		fmt.Println(group + ":")
		for _, name := range names {
			fmt.Println("  " + name)
		}
		fmt.Println()
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
