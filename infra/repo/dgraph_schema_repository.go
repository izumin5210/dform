package repo

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/api"
	"google.golang.org/grpc"

	"github.com/izumin5210/dform/domain/schema"
)

type dgraphSchemaRepository struct {
	dgraph *client.Dgraph
}

func NewDgraphSchemaRepository(conn *grpc.ClientConn) schema.Repository {
	dgraph := client.NewDgraphClient(api.NewDgraphClient(conn))

	return &dgraphSchemaRepository{
		dgraph: dgraph,
	}
}

func (r *dgraphSchemaRepository) GetSchema(ctx context.Context) (*schema.Schema, error) {
	txn := r.dgraph.NewTxn()
	defer txn.Discard(ctx)

	q := "schema {}"
	resp, err := txn.Query(ctx, q)

	if err != nil {
		return nil, fmt.Errorf("failed to query: %v", err)
	}

	s := &schema.Schema{}

	for _, n := range resp.GetSchema() {
		t, err := schema.PredicateTypeOf(n.GetType())
		if err != nil {
			return nil, fmt.Errorf("unknown type: %v", err)
		}
		ps := &schema.PredicateSchema{
			Name:       n.GetPredicate(),
			Type:       t,
			Tokenizers: n.GetTokenizer(),
			Index:      n.GetIndex(),
			Reverse:    n.GetReverse(),
			Count:      n.GetCount(),
			List:       n.GetList(),
		}
		s.Predicates = append(s.Predicates, ps)
	}

	return s, nil
}
