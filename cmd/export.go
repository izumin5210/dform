package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func (r *root) newExportCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export schema information",
		Long:  "Export schema information",
		Run: func(*cobra.Command, []string) {
			d, err := grpc.Dial(r.getDgraphURL(), grpc.WithInsecure())
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to connect with Dgraph server: %v", err))
			}
			fmt.Println(r.getDgraphURL())
			c := client.NewDgraphClient(api.NewDgraphClient(d))
			ctx := context.Background()
			txn := c.NewTxn()
			defer txn.Discard(ctx)
			q := `
					schema {
						type
						index
						reverse
						tokenizer
					}
			`
			resp, err := txn.Query(ctx, q)
			if err != nil {
				log.Fatalln(fmt.Errorf("failed to query: %v", err))
			}
			fmt.Println(resp.GetJson())
			fmt.Println(resp.GetSchema())
		},
	}
}
