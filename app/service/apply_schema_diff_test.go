package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/izumin5210/dform/domain/schema"
)

func Test_ApplySchemaDiffService(t *testing.T) {
	t.Run("without diff", func(t *testing.T) {
		calledRepo1GetSchema := false
		calledRepo2GetSchema := false

		repo1 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) {
				calledRepo1GetSchema = true
				return &schema.Schema{}, nil
			},
		}
		repo2 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) {
				calledRepo2GetSchema = true
				return &schema.Schema{}, nil
			},
			fakeUpdate: func(_ context.Context, diff *schema.Diff) error {
				t.Error("Update() should not be called")
				return nil
			},
		}
		ui := &fakeUI{fakeOutput: func(msg string) {}}

		svc := NewApplySchemaDiffService(repo1, repo2, ui)
		err := svc.Perform(context.Background())

		if err != nil {
			t.Errorf("Perform() returned error %v, want nil", err)
		}

		if !calledRepo1GetSchema {
			t.Error("GetSchema() of base repository should be called")
		}

		if !calledRepo2GetSchema {
			t.Error("GetSchema() of target repository should be called")
		}
	})

	t.Run("with diff", func(t *testing.T) {
		s1 := &schema.Schema{
			Predicates: []*schema.PredicateSchema{
				{Name: "_predicate_", Type: schema.PredicateTypeString, List: true},
				{Name: "name", Type: schema.PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true},
			},
		}
		s2 := &schema.Schema{
			Predicates: []*schema.PredicateSchema{
				{Name: "_predicate_", Type: schema.PredicateTypeString, List: true},
				{Name: "login", Type: schema.PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true, Count: true},
			},
		}
		var gotDiff *schema.Diff

		repo1 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) { return s1, nil },
		}
		repo2 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) { return s2, nil },
			fakeUpdate: func(_ context.Context, diff *schema.Diff) error {
				gotDiff = diff
				return nil
			},
		}
		ui := &fakeUI{fakeOutput: func(msg string) {}}

		svc := NewApplySchemaDiffService(repo1, repo2, ui)
		err := svc.Perform(context.Background())

		if err != nil {
			t.Errorf("Perform() returned error %v, want nil", err)
		}

		if got, want := gotDiff, schema.MakeDiff(s1, s2); !reflect.DeepEqual(got, want) {
			t.Errorf("Applied diff is %v, want %v", got, want)
		}
	})

	t.Run("when GetSchema() of base repository return an error", func(t *testing.T) {
		calledRepoGetSchema := false
		s1 := &schema.Schema{
			Predicates: []*schema.PredicateSchema{
				{Name: "_predicate_", Type: schema.PredicateTypeString, List: true},
				{Name: "name", Type: schema.PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true},
			},
		}

		repo1 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) {
				calledRepoGetSchema = true
				return s1, nil
			},
		}
		repo2 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) { return nil, errors.New("an error") },
			fakeUpdate: func(_ context.Context, diff *schema.Diff) error {
				t.Error("Update() should not be called")
				return nil
			},
		}
		ui := &fakeUI{fakeOutput: func(msg string) {}}

		svc := NewApplySchemaDiffService(repo1, repo2, ui)
		err := svc.Perform(context.Background())

		if err == nil {
			t.Errorf("Perform() returned error %v, want nil", err)
		}

		if !calledRepoGetSchema {
			t.Error("GetSchema() of base repository should be called")
		}
	})

	t.Run("when GetSchema() of target repository return an error", func(t *testing.T) {
		calledRepoGetSchema := false
		s2 := &schema.Schema{
			Predicates: []*schema.PredicateSchema{
				{Name: "_predicate_", Type: schema.PredicateTypeString, List: true},
				{Name: "login", Type: schema.PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				{Name: "friend", Type: schema.PredicateTypeUID, Reverse: true, Count: true},
			},
		}

		repo1 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) { return nil, errors.New("an error") },
		}
		repo2 := &fakeSchemaRepo{
			fakeGetSchema: func(context.Context) (*schema.Schema, error) {
				calledRepoGetSchema = true
				return s2, nil
			},
			fakeUpdate: func(_ context.Context, diff *schema.Diff) error {
				t.Error("Update() should not be called")
				return nil
			},
		}
		ui := &fakeUI{fakeOutput: func(msg string) {}}

		svc := NewApplySchemaDiffService(repo1, repo2, ui)
		err := svc.Perform(context.Background())

		if err == nil {
			t.Errorf("Perform() returned error %v, want nil", err)
		}

		if !calledRepoGetSchema {
			t.Error("GetSchema() of target repository should be called")
		}
	})
}
