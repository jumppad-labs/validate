package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var stringCmd = &cobra.Command{
	Use: "string",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// string values are equal
var stringCompareCmd = &cobra.Command{
	Use:  "compare <first> <second>",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		first := args[0]
		second := args[1]

		if eq {
			if first == second {
				os.Exit(0)
			}
		}

		if lt {
			if first < second {
				os.Exit(0)
			}
		}

		if gt {
			if first > second {
				os.Exit(0)
			}
		}

		os.Exit(1)
	},
}
