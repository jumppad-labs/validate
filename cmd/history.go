package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jumppad-labs/validate/data"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use: "history",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// file contains
var historyContainsCmd = &cobra.Command{
	Use:  "contains <command> [flags]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := args[0]

		histfile := "/var/lib/jumppad/history"
		content := []byte{}
		if _, err := os.Stat(histfile); err == nil {
			content, err = os.ReadFile(histfile)
			if err != nil {
				data.Error(err)
			}
		}

		reA := regexp.MustCompile(`\\\n`)
		withoutNewlines := reA.ReplaceAllString(string(content), " ")

		reB := regexp.MustCompile(`:\s?\d+:\d;`)
		withoutTimestamps := reB.ReplaceAllString(withoutNewlines, "")

		lines := strings.Split(withoutTimestamps, "\n")
		for _, line := range lines {
			if exact {
				// if we need to match the line exactly
				if line == value {
					os.Exit(0)
				}
			} else if loose {
				// if we can match parts of the value in any order
				found := 0
				parts := strings.Split(value, " ")
				for _, part := range parts {
					fmt.Println(part)
					fmt.Println(line)
					if strings.Contains(line, part) {
						found++
					}
				}

				if found == len(parts) {
					os.Exit(0)
				}
			} else if strings.Contains(line, value) {
				os.Exit(0)
			}
		}

		os.Exit(1)
	},
}
