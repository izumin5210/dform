package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/component"
	"github.com/izumin5210/dform/app/service"
)

func newDiffCommand(app component.App) *cobra.Command {
	return &cobra.Command{
		Use:   "diff",
		Short: "Diff schema",
		Long:  "Diff schema",
		RunE: func(c *cobra.Command, _ []string) error {
			svc := service.NewShowSchemaDiffService(
				app.DgraphSchemaRepository(),
				app.FileSchemaRepository(),
				app.UI(),
			)

			return svc.Perform(context.Background())
		},
	}
}
