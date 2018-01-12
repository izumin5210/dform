package schema

import (
	"testing"
)

func Test_PredicateSchema_String(t *testing.T) {
	cases := []struct {
		in  *PredicateSchema
		out string
	}{
		{
			in: &PredicateSchema{
				Name: "name",
				Type: PredicateTypeString,
			},
			out: "name: string .",
		},
		{
			in: &PredicateSchema{
				Name:       "login",
				Type:       PredicateTypeString,
				Tokenizers: []string{"exact", "term"},
				Index:      true,
			},
			out: "login: string @index(exact, term) .",
		},
		{
			in: &PredicateSchema{
				Name:    "rated",
				Type:    PredicateTypeUID,
				Reverse: true,
				Count:   true,
			},
			out: "rated: uid @reverse @count .",
		},
		{
			in: &PredicateSchema{
				Name: "score",
				Type: PredicateTypeInt,
				List: true,
			},
			out: "score: [int] .",
		},
	}

	for _, c := range cases {
		if got, want := c.in.String(), c.out; got != want {
			t.Errorf("%v is %q in string, want %q", c.in, got, want)
		}
	}
}
