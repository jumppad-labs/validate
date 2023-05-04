package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var exact bool
var loose bool

var eq bool
var lt bool
var gt bool

var rootCmd = &cobra.Command{
	Use: "validate",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(successCmd)
	rootCmd.AddCommand(failureCmd)

	rootCmd.AddCommand(historyCmd)
	historyCmd.AddCommand(historyContainsCmd)
	historyContainsCmd.Flags().BoolVar(&exact, "match-line", false, "must match entire line")
	historyContainsCmd.Flags().BoolVar(&loose, "any-order", false, "value parts can match in any order")

	rootCmd.AddCommand(terraformCmd)
	rootCmd.AddCommand(dockerCmd)

	rootCmd.AddCommand(jsonCmd)

	rootCmd.AddCommand(stringCmd)
	stringCmd.AddCommand(stringCompareCmd)
	stringCompareCmd.Flags().BoolVar(&eq, "eq", true, "first == second")
	stringCompareCmd.Flags().BoolVar(&lt, "lt", false, "first < second")
	stringCompareCmd.Flags().BoolVar(&gt, "gt", false, "first > second")

	rootCmd.AddCommand(fileCmd)
	fileCmd.AddCommand(fileExistsCmd)
	fileCmd.AddCommand(fileContainsCmd)
	fileContainsCmd.Flags().BoolVar(&exact, "match-line", false, "must match entire line")

	rootCmd.AddCommand(commandCmd)
	commandCmd.AddCommand(commandExistsCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
