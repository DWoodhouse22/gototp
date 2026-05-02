package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [group]",
	Short: "Set the active group",
	Args:  cobra.MaximumNArgs(1),
	Run:   Use,
}

func Use(cmd *cobra.Command, args []string) {
	currentGroup, err := getCurrentGroup()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Current group: %s\n\n", currentGroup)

	groupSet, err := getGroupSet(currentGroup)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(args) == 1 {
		handleGroupArg(args[0], groupSet)
		return
	}

	var groups []string
	for g := range groupSet {
		groups = append(groups, g)
	}
	sort.Strings(groups)

	selected, ok := promptForGroup(groups)
	if !ok {
		return
	}

	switchToGroup(selected, false)
}

func getCurrentGroup() (string, error) {
	g, err := storage.GetCurrentGroup()
	if err != nil {
		return "", err
	}
	if g == "" {
		g = "default"
	}
	return g, nil
}

func getGroupSet(current string) (map[string]any, error) {
	accounts, err := storage.ListAccounts()
	if err != nil {
		return nil, err
	}

	out := make(map[string]any)
	out[current] = struct{}{}
	for _, act := range accounts {
		out[act.Group] = struct{}{}
	}

	return out, nil
}

func handleGroupArg(target string, groupSet map[string]any) {
	if _, ok := groupSet[target]; ok {
		switchToGroup(target, false)
		return
	}

	fmt.Printf("Group '%s' does not exist.\n", target)
	if !confirm("Create and switch to it? (y/N): ") {
		fmt.Println("Cancelled")
		return
	}

	switchToGroup(target, true)
}

func promptForGroup(groups []string) (string, bool) {
	fmt.Println("Available groups:")
	for i, g := range groups {
		fmt.Printf("%d) %s\n", i+1, g)
	}

	fmt.Print("\nSelect a group: ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(groups) {
		fmt.Println("Invalid selection")
		return "", false
	}

	return groups[choice-1], true
}

func switchToGroup(group string, created bool) {
	err := storage.SetCurrentGroup(group)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if created {
		fmt.Printf("Created and switched to group '%s'\n", group)
	} else {
		fmt.Printf("Switched to group '%s'\n", group)
	}
}

func confirm(prompt string) bool {
	fmt.Print(prompt)
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y"
}

func init() {
	rootCmd.AddCommand(useCmd)
}
