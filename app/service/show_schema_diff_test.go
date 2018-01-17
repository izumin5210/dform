package service

import (
	"context"
	"strings"
	"testing"

	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/domain/schema"
)

func Test_ShowSchemaDiff(t *testing.T) {
	predPred := &schema.PredicateSchema{Name: "_predicate_", Type: schema.PredicateTypeString, List: true}
	predLogin := &schema.PredicateSchema{Name: "login", Type: schema.PredicateTypeString, Index: true, Tokenizers: []string{"term"}}
	predUserID := &schema.PredicateSchema{Name: "user_id", Type: schema.PredicateTypeInt, Index: true, Tokenizers: []string{"int"}}
	predFriend1 := &schema.PredicateSchema{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true}
	predFriend2 := &schema.PredicateSchema{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true, Count: true}
	predCreatedAt := &schema.PredicateSchema{Name: "created_at", Type: schema.PredicateTypeDateTime}

	cases := []struct {
		test             string
		schema1, schema2 *schema.Schema
		outputs          [][]string
	}{
		{
			test: "with no diffs",
			schema1: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred},
			},
			schema2: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred},
			},
		},
		{
			test: "with only insertions",
			schema1: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred},
			},
			schema2: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred, predLogin},
			},
			outputs: [][]string{
				[]string{predLogin.String()},
			},
		},
		{
			test: "with only deletions",
			schema1: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred, predLogin},
			},
			schema2: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred},
			},
			outputs: [][]string{
				[]string{predLogin.String()},
			},
		},
		{
			test: "with only modifications",
			schema1: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred, predFriend1},
			},
			schema2: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred, predFriend2},
			},
			outputs: [][]string{
				[]string{predFriend1.String(), predFriend2.String()},
			},
		},
		{
			test: "with complex diff",
			schema1: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred, predCreatedAt, predFriend1},
			},
			schema2: &schema.Schema{
				Predicates: []*schema.PredicateSchema{predPred, predLogin, predUserID, predFriend2},
			},
			outputs: [][]string{
				[]string{predLogin.String(), predUserID.String()},
				[]string{predCreatedAt.String()},
				[]string{predFriend1.String(), predFriend2.String()},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.test, func(t *testing.T) {
			var outputs []string

			repo1 := &fakeSchemaRepo{
				fakeGetSchema: func(context.Context) (*schema.Schema, error) { return c.schema1, nil },
			}
			repo2 := &fakeSchemaRepo{
				fakeGetSchema: func(context.Context) (*schema.Schema, error) { return c.schema2, nil },
			}
			ui := &fakeUI{
				fakeOutput: func(msg string) { outputs = append(outputs, msg) },
			}

			svc := &showSchemaDiffService{
				repo1: repo1,
				repo2: repo2,
				ui:    ui,
			}

			err := svc.Perform(context.Background())

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got, want := len(outputs), len(c.outputs); got != want {
				t.Errorf("Perform() outputs %d items, want %d items", got, want)
			} else {
				for i, got := range outputs {
					for _, want := range c.outputs[i] {
						if !strings.Contains(got, want) {
							t.Errorf("Output %q should include %q", got, want)
						}
					}
				}
			}
		})
	}
}

// Fake implementations
//================================================================
type fakeSchemaRepo struct {
	schema.Repository
	fakeGetSchema func(context.Context) (*schema.Schema, error)
}

func (f *fakeSchemaRepo) GetSchema(c context.Context) (*schema.Schema, error) {
	return f.fakeGetSchema(c)
}

type fakeUI struct {
	system.UI
	fakeOutput func(msg string)
}

func (f *fakeUI) Output(msg string) {
	f.fakeOutput(msg)
}
