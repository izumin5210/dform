package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/izumin5210/dform/app/component"
	"github.com/izumin5210/dform/domain/schema"
)

func newDiffCommand(app component.App) *cobra.Command {
	return &cobra.Command{
		Use:   "diff",
		Short: "Diff schema",
		Long:  "Diff schema",
		Run: func(c *cobra.Command, _ []string) {
			fileRepo := app.FileSchemaRepository()
			dgraphRepo, err := app.DgraphSchemaRepository()
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to get repository: %v", err))
			}

			s1, err := dgraphRepo.GetSchema(context.Background())
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to get schema: %v", err))
			}
			s2, err := fileRepo.GetSchema(context.Background())
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to get schema: %v", err))
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
		},
	}
}
