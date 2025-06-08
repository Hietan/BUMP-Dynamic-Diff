package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	minimalMode bool
)

var infoCmd = &cobra.Command{
	Use:   "info [ID]",
	Short: "Display information for a given ID",
	Long: `The info command retrieves and displays detailed information
for the specified ID. For example:

BUMP-Dynamic-Diff info abc123`,
	Args:          cobra.ExactArgs(1),
	SilenceUsage:  true,
	SilenceErrors: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		dir := filepath.Join("cmd", "data", "bump", "image")

		entries, err := os.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("failed to read directory %q: %w", dir, err)
		}

		found := false
		for _, e := range entries {
			name := e.Name()
			if e.IsDir() || filepath.Ext(name) != ".json" {
				continue
			}
			if !strings.HasPrefix(name, id) {
				continue
			}

			path := filepath.Join(dir, name)
			raw, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %q: %w", path, err)
			}

			if minimalMode {
				var obj map[string]any
				if err := json.Unmarshal(raw, &obj); err != nil {
					return fmt.Errorf("failed to parse JSON %q: %w", path, err)
				}
				fc, ok := obj["failureCategory"]
				if !ok {
					fmt.Printf("%s: field \"failureCategory\" not found\n", name)
				} else {
					fmt.Printf("%s: failureCategory = %v\n", name, fc)
				}
			} else {
				fmt.Printf("----- %s -----\n%s\n\n", name, string(raw))
			}
			found = true
		}

		if !found {
			return fmt.Errorf("no JSON files found with prefix %q in %q", id, dir)
		}
		return nil
	},
}

func init() {
	infoCmd.Flags().BoolVarP(&minimalMode, "minimal", "m", false, "only display the failureCategory field")
	rootCmd.AddCommand(infoCmd)
}
