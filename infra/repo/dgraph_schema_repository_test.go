package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/izumin5210/dform/domain/schema"
)

func Test_DgraphSchemaRepository_GetSchema(t *testing.T) {
	defer testDgraph.MustCleanup(t)

	repo := NewDgraphSchemaRepository(testDgraph.GetConn())

	t.Run("when schema has not defined", func(t *testing.T) {
		schema, err := repo.GetSchema(context.Background())

		if err != nil {
			t.Errorf("GetSchema() returned an error: %v", err)
		}

		if schema == nil {
			t.Fatal("GetSchema() returned nil")
		}

		if got, want := len(schema.Predicates), 1; got != want {
			t.Errorf("GetSchema() returned %d predicate schemata, want %d", got, want)
		} else if got, want := schema.Predicates[0].Name, "_predicate_"; got != want {
			t.Errorf("Returned predicate is %q, want %q", got, want)
		}
	})

	t.Run("when schema has defined", func(t *testing.T) {
		defer testDgraph.MustCleanup(t)
		testDgraph.MustAlter(t, `
			name: string .
			login: string @index(exact, term) .

			rated: uid @reverse @count .
			score: [int] .
		`)

		wantPreds := []*schema.PredicateSchema{
			{Name: "name", Type: schema.PredicateTypeString},
			{Name: "login", Type: schema.PredicateTypeString, Tokenizers: []string{"exact", "term"}, Index: true},
			{Name: "rated", Type: schema.PredicateTypeUID, Reverse: true, Count: true},
			{Name: "score", Type: schema.PredicateTypeInt, List: true},
			{Name: "_predicate_", Type: schema.PredicateTypeString, List: true},
		}
		wantPredByName := map[string]*schema.PredicateSchema{}
		for _, p := range wantPreds {
			wantPredByName[p.Name] = p
		}

		schema, err := repo.GetSchema(context.Background())

		if err != nil {
			t.Errorf("GetSchema() returned an error: %v", err)
		}

		if schema == nil {
			t.Fatal("GetSchema() returned nil")
		}

		if got, want := len(schema.Predicates), len(wantPreds); got != want {
			t.Errorf("GetSchema() returned %d predicate schemata, want %d", got, want)
		} else {
			for _, pred := range schema.Predicates {
				if got, want := pred, wantPredByName[pred.Name]; !reflect.DeepEqual(got, want) {
					t.Errorf("Returned predicate is %q, want %q", got, want)
				}
			}
		}
	})
}
