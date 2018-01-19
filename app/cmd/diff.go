package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/di"
)

func newDiffCommand(app di.RootComponent) *cobra.Command {
	return &cobra.Command{
		Use:   "diff",
		Short: "Diff schema",
		Long:  "Diff schema",
		RunE: func(c *cobra.Command, _ []string) error {
			return app.ShowSchemaDiffService().Perform(context.Background())
		},
	}
}
