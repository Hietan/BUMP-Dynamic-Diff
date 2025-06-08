package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [ID]",
	Short: "Display information for a given ID",
	Long: `The info command retrieves and displays detailed information
for the specified ID. For example:

BUMP-Dynamic-Diff info abc123`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		fmt.Printf("Retrieving information for ID: %s\n", id)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
