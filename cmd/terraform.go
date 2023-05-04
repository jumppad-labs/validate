package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var terraformCmd = &cobra.Command{
	Use: "terraform",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}
