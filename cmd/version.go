package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (r *root) newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Long:  "Print version information",
		Run: func(*cobra.Command, []string) {
			fmt.Printf("%s %s (%s)\n", r.name, r.version, r.revision)
		},
	}
}
