package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dform",
	Short: "CLI tool to manage Dgraph schema",
	Long:  "CLI tool to manage Dgraph schema",
}

// Execute executes the command.
func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return -1
	}
	return 0
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
