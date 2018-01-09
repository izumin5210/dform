package cmd

import (
	"github.com/spf13/cobra"
)

func (r *root) newExportCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export schema information",
		Long:  "Export schema information",
		Run: func(*cobra.Command, []string) {
		},
	}
}
