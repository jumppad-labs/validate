package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/jumppad-labs/validate/data"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use: "file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// file exists on path
var fileExistsCmd = &cobra.Command{
	Use:  "exists <file>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		if _, err := os.Stat(filename); err == nil {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	},
}

// file contains
var fileContainsCmd = &cobra.Command{
	Use:  "contains <file> <value> [flags]",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		value := args[1]

		f, err := os.Open(filename)
		if err != nil {
			data.Error(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			// remote leading/trailing whitespace
			content := strings.TrimSpace(scanner.Text())

			if exact && content == value {
				os.Exit(0)
			} else if !exact && strings.Contains(content, value) {
				os.Exit(0)
			}
		}

		if err := scanner.Err(); err != nil {
			data.Error(err)
		}

		os.Exit(1)
	},
}
