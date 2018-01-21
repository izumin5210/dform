package schema

import (
	"reflect"
	"testing"
)

func Test_Diff(t *testing.T) {
	cases := []struct {
		s1, s2 *Schema
		out    *Diff
		empty  bool
	}{
		{
			s1: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
				},
			},
			s2: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
				},
			},
			out:   &Diff{},
			empty: true,
		},
		{
			s1: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
				},
			},
			s2: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
					{Name: "name", Type: PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				},
			},
			out: &Diff{
				Inserted: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				},
			},
			empty: false,
		},
		{
			s1: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
					{Name: "name", Type: PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				},
			},
			s2: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
				},
			},
			out: &Diff{
				Deleted: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
				},
			},
			empty: false,
		},
		{
			s1: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
					{Name: "friend", Type: PredicateTypeUID, Reverse: true},
				},
			},
			s2: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
					{Name: "friend", Type: PredicateTypeUID, Reverse: true, Count: true},
				},
			},
			out: &Diff{
				Modified: []*ModifiedPredicate{
					{
						From: &PredicateSchema{Name: "friend", Type: PredicateTypeUID, Reverse: true},
						To:   &PredicateSchema{Name: "friend", Type: PredicateTypeUID, Reverse: true, Count: true},
					},
				},
			},
			empty: false,
		},
		{
			s1: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
					{Name: "login", Type: PredicateTypeString},
					{Name: "score", Type: PredicateTypeFloat},
					{Name: "last_logged_in_at", Type: PredicateTypeDateTime},
					{Name: "friend", Type: PredicateTypeUID, Reverse: true},
				},
			},
			s2: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "_predicate_", Type: PredicateTypeString, List: true},
					{Name: "login", Type: PredicateTypeString, Index: true, Tokenizers: []string{"exact"}},
					{Name: "name", Type: PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
					{Name: "created_at", Type: PredicateTypeDateTime},
					{Name: "friend", Type: PredicateTypeUID, Reverse: true, Count: true},
				},
			},
			out: &Diff{
				Inserted: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString, Index: true, Tokenizers: []string{"term"}},
					{Name: "created_at", Type: PredicateTypeDateTime},
				},
				Deleted: []*PredicateSchema{
					{Name: "score", Type: PredicateTypeFloat},
					{Name: "last_logged_in_at", Type: PredicateTypeDateTime},
				},
				Modified: []*ModifiedPredicate{
					{
						From: &PredicateSchema{Name: "login", Type: PredicateTypeString},
						To:   &PredicateSchema{Name: "login", Type: PredicateTypeString, Index: true, Tokenizers: []string{"exact"}},
					},
					{
						From: &PredicateSchema{Name: "friend", Type: PredicateTypeUID, Reverse: true},
						To:   &PredicateSchema{Name: "friend", Type: PredicateTypeUID, Reverse: true, Count: true},
					},
				},
			},
			empty: false,
		},
	}

	for _, c := range cases {
		diff := MakeDiff(c.s1, c.s2)

		if got, want := diff, c.out; !reflect.DeepEqual(got, want) {
			t.Errorf("Diff in %v and %v is %v, want %v", c.s1, c.s2, got, want)
		}

		if got, want := diff.Empty(), c.empty; got != want {
			t.Errorf("Empty() returned %t, want %t", got, want)
		}
	}
}
