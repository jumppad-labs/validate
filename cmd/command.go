package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var commandCmd = &cobra.Command{
	Use: "command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// command exists on the path
var commandExistsCmd = &cobra.Command{
	Use:  "exists <command> [flags]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		if _, err := exec.LookPath(command); err == nil {
			os.Exit(0)
		}

		os.Exit(1)
	},
}
