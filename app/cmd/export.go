package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/di"
	"github.com/izumin5210/dform/util/log"
)

func newExportCommand(component di.RootComponent) *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export schema information",
		Long:  "Export schema information",
		RunE: func(c *cobra.Command, _ []string) error {
			repo := component.DgraphSchemaRepository()
			s, err := repo.GetSchema(context.Background())
			if err != nil {
				component.UI().Error("Failed to get schema diff")
				log.Error("failed to get schema", "error", err)
				return err
			}
			data, err := s.MarshalText()
			if err != nil {
				log.Error("failed to marshal schema", "error", err)
				return err
			}
			component.UI().Output(string(data))
			return nil
		},
	}
}
