package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var failureCmd = &cobra.Command{
	Use:  "fail <message>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]
		os.Stdout.WriteString(message)
		os.Exit(1)
	},
}
