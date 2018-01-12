package schema

import (
	"reflect"
	"testing"
)

func Test_Schema_UnmarshalText(t *testing.T) {
	cases := []struct {
		in  string
		out *Schema
	}{
		{
			in:  "",
			out: &Schema{},
		},
		{
			in: "name: string .",
			out: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString},
				},
			},
		},
		{
			in: `
			name: string .
			login: string @index(exact, term) .

			rated: uid @reverse @count .
			score: [int] .
			`,
			out: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString},
					{Name: "login", Type: PredicateTypeString, Tokenizers: []string{"exact", "term"}, Index: true},
					{Name: "rated", Type: PredicateTypeUID, Reverse: true, Count: true},
					{Name: "score", Type: PredicateTypeInt, List: true},
				},
			},
		},
	}

	for _, c := range cases {
		s := &Schema{}
		err := s.UnmarshalText([]byte(c.in))

		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if got, want := s, c.out; reflect.DeepEqual(got, want) {
			t.Errorf("%q is %v in string, want %v", c.in, got, want)
		}
	}
}
