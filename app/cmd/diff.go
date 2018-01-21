package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/di"
	"github.com/izumin5210/dform/util/log"
)

func newDiffCommand(component di.RootComponent) *cobra.Command {
	return &cobra.Command{
		Use:   "diff",
		Short: "Diff schema",
		Long:  "Diff schema",
		RunE: func(c *cobra.Command, _ []string) error {
			var err error
			ctx := context.Background()

			err = component.ShowSchemaDiffService().Perform(ctx)
			if err != nil {
				log.Error("failed to show schema diff", "error", err)
				return err
			}

			ok, err := component.UI().Confirm("Would you like to apply it?")
			if err != nil {
				log.Error("failed to confirm to apply", "error", err)
				return err
			}

			if ok {
				err = component.ApplySchemaDiffService().Perform(ctx)
				if err != nil {
					log.Error("failed to apply schema diff", "error", err)
					return err
				}
			}

			return nil
		},
	}
}
