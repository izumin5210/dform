package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/component"
)

// New creates a new command object
func New(app component.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   app.Config().Name,
		Short: "CLI tool to manage Dgraph schema",
		Long:  "CLI tool to manage Dgraph schema",
	}

	cobra.OnInitialize(app.Config().Init)

	rootCmd.PersistentFlags().StringVar(
		&(app.Config().ConfigFilePath),
		"config",
		"",
		fmt.Sprintf("config file (default is $PWD/%s.toml)", app.Config().GetDefaultConfigName()),
	)
	rootCmd.AddCommand(
		newExportCommand(app),
		newVersionCommand(app),
	)

	return rootCmd
}
