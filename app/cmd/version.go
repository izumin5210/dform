package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/component"
)

func newVersionCommand(app component.App) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Long:  "Print version information",
		Run: func(*cobra.Command, []string) {
			c := app.Config()
			fmt.Printf("%s %s (%s)\n", c.Name, c.Version, c.Revision)
		},
	}
}
