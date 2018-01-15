package cmd

import (
	"context"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/component"
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/util/log"
)

func newDiffCommand(app component.App) *cobra.Command {
	return &cobra.Command{
		Use:   "diff",
		Short: "Diff schema",
		Long:  "Diff schema",
		RunE: func(c *cobra.Command, _ []string) error {
			fileRepo := app.FileSchemaRepository()
			dgraphRepo, err := app.DgraphSchemaRepository()
			if err != nil {
				log.Error("failed to get repository", "error", err)
				return err
			}

			s1, err := dgraphRepo.GetSchema(context.Background())
			if err != nil {
				log.Error("failed to get schema from Dgraph", "error", err)
				return err
			}
			s2, err := fileRepo.GetSchema(context.Background())
			if err != nil {
				log.Error("failed to get schema from filesystem", "error", err)
				return err
			}

			diff := schema.MakeDiff(s1, s2)
			red := color.New(color.FgRed).FprintfFunc()
			green := color.New(color.FgGreen).FprintfFunc()

			if len(diff.Inserted) > 0 {
				c.Println("Added predicates:")
				c.Println("")
				for _, pred := range diff.Inserted {
					green(c.OutOrStdout(), "    + %s\n", pred)
				}
				c.Println("")
			}

			if len(diff.Deleted) > 0 {
				c.Println("Dropped predicates:")
				c.Println("")
				for _, pred := range diff.Deleted {
					red(c.OutOrStdout(), "    - %s\n", pred)
				}
				c.Println("")
			}

			if len(diff.Modified) > 0 {
				c.Println("Modified predicates:")
				c.Println("")
				for _, pair := range diff.Modified {
					red(c.OutOrStdout(), "    - %s\n", pair.From)
					green(c.OutOrStdout(), "    + %s\n", pair.To)
				}
			}

			return nil
		},
	}
}
