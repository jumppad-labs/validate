package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var successCmd = &cobra.Command{
	Use:  "success",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(0)
	},
}
