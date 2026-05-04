package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCommand = cobra.Command{
	Use:     "app",
	Version: "1.0.0",
	Short:   "Run Application",
	Run: func(_ *cobra.Command, args []string) {
		fmt.Println("Running application")
	},
}

func Execute() error {
	return rootCommand.Execute()
}

func init() {
	rootCommand.AddCommand(
		restCommand,
	)
}
