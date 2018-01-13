package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/izumin5210/dform/infra/repo"
)

func (r *root) newExportCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export schema information",
		Long:  "Export schema information",
		Run: func(*cobra.Command, []string) {
			conn, err := grpc.Dial(r.getDgraphURL(), grpc.WithInsecure())
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to connect with Dgraph server: %v", err))
			}
			repo := repo.NewDgraphSchemaRepository(conn)
			s, err := repo.GetSchema(context.Background())
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to get schema: %v", err))
			}
			fmt.Println(s.Predicates)
		},
	}
}
