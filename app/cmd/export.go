package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/di"
	"github.com/izumin5210/dform/util/log"
)

func newExportCommand(app di.App) *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export schema information",
		Long:  "Export schema information",
		RunE: func(c *cobra.Command, _ []string) error {
			repo := app.DgraphSchemaRepository()
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
