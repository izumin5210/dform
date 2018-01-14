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
		Run: func(c *cobra.Command, _ []string) {
			conn, err := grpc.Dial(r.getDgraphURL(), grpc.WithInsecure())
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to connect with Dgraph server: %v", err))
			}
			repo := repo.NewDgraphSchemaRepository(conn)
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
