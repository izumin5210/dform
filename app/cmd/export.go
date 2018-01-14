package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/component"
)

func newExportCommand(app component.App) *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export schema information",
		Long:  "Export schema information",
		Run: func(c *cobra.Command, _ []string) {
			repo, err := app.DgraphSchemaRepository()
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to get repository: %v", err))
			}
			s, err := repo.GetSchema(context.Background())
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to get schema: %v", err))
			}
			data, err := s.MarshalText()
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to marshal schema: %v", err))
			}
			c.Println(string(data))
		},
	}
}
