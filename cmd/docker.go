package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use: "docker",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}
