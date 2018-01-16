package repo

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/api"

	"github.com/izumin5210/dform/domain/schema"
)

type dgraphSchemaRepository struct {
	pool GrpcPool
}

// NewDgraphSchemaRepository creates new schema repository interface for accessing Dgraph.
func NewDgraphSchemaRepository(pool GrpcPool) schema.Repository {
	return &dgraphSchemaRepository{
		pool: pool,
	}
}

func (r *dgraphSchemaRepository) GetSchema(ctx context.Context) (*schema.Schema, error) {
	conn, err := r.pool.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Dgraph: %v", err)
	}
	defer conn.Close()

	dgraph := client.NewDgraphClient(api.NewDgraphClient(conn))

	txn := dgraph.NewTxn()
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

func (r *dgraphSchemaRepository) Update(ctx context.Context, diff *schema.Diff) error {
	alteredPreds := make([]*schema.PredicateSchema, 0, len(diff.Inserted)+len(diff.Modified))
	droppedPreds := make([]*schema.PredicateSchema, 0, len(diff.Deleted))

	for _, pred := range diff.Inserted {
		alteredPreds = append(alteredPreds, pred)
	}
	for _, pair := range diff.Modified {
		alteredPreds = append(alteredPreds, pair.To)
	}
	for _, pred := range diff.Deleted {
		droppedPreds = append(droppedPreds, pred)
	}

	if len(alteredPreds)+len(droppedPreds) == 0 {
		return nil
	}

	conn, err := r.pool.Get()
	if err != nil {
		return fmt.Errorf("failed to connect to Dgraph: %v", err)
	}
	defer conn.Close()

	dgraph := client.NewDgraphClient(api.NewDgraphClient(conn))

	q, err := (&schema.Schema{Predicates: alteredPreds}).MarshalText()
	if err != nil {
		return fmt.Errorf("failed to marshal schema: %v", err)
	}
	err = dgraph.Alter(ctx, &api.Operation{Schema: string(q)})
	if err != nil {
		return fmt.Errorf("failed to alter schema: %v", err)
	}
	for _, pred := range diff.Deleted {
		err := dgraph.Alter(ctx, &api.Operation{DropAttr: pred.Name})
		if err != nil {
			return fmt.Errorf("failed to alter schema: %v", err)
		}
	}

	return nil
}
