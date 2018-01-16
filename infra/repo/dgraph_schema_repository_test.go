package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/izumin5210/dform/domain/schema"
)

func Test_DgraphSchemaRepository_GetSchema(t *testing.T) {
	defer testDgraph.MustCleanup(t)

	repo := NewDgraphSchemaRepository(testDgraph)

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

func Test_DgraphSchemaRepository_Update(t *testing.T) {
	defer testDgraph.MustCleanup(t)

	repo := NewDgraphSchemaRepository(testDgraph)

	mustGetSchema := func(t *testing.T) *schema.Schema {
		s, err := repo.GetSchema(context.Background())
		if err != nil {
			t.Fatalf("Failed GetSchema(): %v", err)
		}
		if s == nil {
			t.Fatalf("GetSchema() returned nil")
		}
		return s
	}

	t.Run("when diff is empty", func(t *testing.T) {
		before := mustGetSchema(t)

		err := repo.Update(context.Background(), &schema.Diff{})

		if err != nil {
			t.Errorf("Update() returned an error: %v", err)
		}

		after := mustGetSchema(t)

		if !reflect.DeepEqual(before, after) {
			t.Errorf("Update() should not modify schema, before: %v, after: %v", before, after)
		}
	})

	t.Run("when diff exists", func(t *testing.T) {
		diff := &schema.Diff{
			Inserted: []*schema.PredicateSchema{
				{Name: "name", Type: schema.PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				{Name: "created_at", Type: schema.PredicateTypeDateTime},
			},
			Deleted: []*schema.PredicateSchema{
				{Name: "score", Type: schema.PredicateTypeFloat},
				{Name: "last_logged_in_at", Type: schema.PredicateTypeDateTime},
			},
			Modified: []*schema.ModifiedPredicate{
				{
					From: &schema.PredicateSchema{Name: "login", Type: schema.PredicateTypeString},
					To:   &schema.PredicateSchema{Name: "login", Type: schema.PredicateTypeString, Index: true, Tokenizers: []string{"exact"}},
				},
				{
					From: &schema.PredicateSchema{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true},
					To:   &schema.PredicateSchema{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true, Count: true},
				},
			},
		}

		// -- setup
		before := []*schema.PredicateSchema{}
		remainByName := map[string]*schema.PredicateSchema{}
		for _, pred := range diff.Inserted {
			remainByName[pred.Name] = pred
		}
		for _, pred := range diff.Deleted {
			before = append(before, pred)
		}
		for _, pair := range diff.Modified {
			remainByName[pair.To.Name] = pair.To
			before = append(before, pair.From)
		}

		q, err := (&schema.Schema{Predicates: before}).MarshalText()
		if err != nil {
			t.Fatalf("Failed to setup test schema: %v", err)
		}
		testDgraph.MustAlter(t, string(q))

		// -- testing
		err = repo.Update(context.Background(), diff)

		if err != nil {
			t.Errorf("Update() returned an error: %v", err)
		}

		after := mustGetSchema(t).Predicates

		if got, want := len(after), len(remainByName)+1; got != want {
			t.Errorf("Predicate count after updated is %d, want %d", got, want)
		}

		for _, pred2 := range after {
			name := pred2.Name
			if pred1, ok := remainByName[name]; ok {
				if got, want := pred2, pred1; !reflect.DeepEqual(got, want) {
					t.Errorf("Predicate %v existed, want %v", got, want)
				}
				delete(remainByName, name)
			} else if name != "_predicate_" {
				t.Errorf("Unexpected predicate exists: %v", pred2)
			}
		}

		if got, want := len(remainByName), 0; got != want {
			t.Errorf("predicate(s) should exists: %v", remainByName)
		}
	})
}
