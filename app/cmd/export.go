package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/component"
	"github.com/izumin5210/dform/util/log"
)

func newExportCommand(app component.App) *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export schema information",
		Long:  "Export schema information",
		RunE: func(c *cobra.Command, _ []string) error {
			repo, err := app.DgraphSchemaRepository()
			if err != nil {
				log.Error("failed to get repository: %v", "error", err)
				return err
			}
			s, err := repo.GetSchema(context.Background())
			if err != nil {
				log.Error("failed to get schema", "error", err)
				return err
			}
			data, err := s.MarshalText()
			if err != nil {
				log.Error("failed to marshal schema", "error", err)
				return err
			}
			c.Println(string(data))
			return nil
		},
	}
}
