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
			in: "",
			out: &Schema{
				Predicates: []*PredicateSchema{},
			},
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

createdAt: dateTime .
rated: uid @reverse @count .
score: [int] .
			`,
			out: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString},
					{Name: "login", Type: PredicateTypeString, Tokenizers: []string{"exact", "term"}, Index: true},
					{Name: "createdAt", Type: PredicateTypeDateTime},
					{Name: "rated", Type: PredicateTypeUID, Reverse: true, Count: true},
					{Name: "score", Type: PredicateTypeInt, List: true},
				},
			},
		},
	}

	t.Run("with invalid schema", func(t *testing.T) {
		s := &Schema{}
		err := s.UnmarshalText([]byte("name: string"))

		if err == nil {
			t.Error("UnmarshalText() should return an error")
		}

		if s.Predicates != nil {
			t.Errorf("Unmarshaled schema predicates is %v, want nil", s.Predicates)
		}
	})

	t.Run("with valid schema", func(t *testing.T) {
		for _, c := range cases {
			s := &Schema{}
			err := s.UnmarshalText([]byte(c.in))

			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}

			if got, want := s, c.out; !reflect.DeepEqual(got, want) {
				t.Errorf("%q is %v in string, want %v", c.in, got, want)
			}
		}
	})
}

func Test_Schema_MarshalText(t *testing.T) {
	cases := []struct {
		in  *Schema
		out string
	}{
		{
			in:  &Schema{},
			out: "",
		},
		{
			in: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString},
				},
			},
			out: "name: string .",
		},
		{
			in: &Schema{
				Predicates: []*PredicateSchema{
					{Name: "name", Type: PredicateTypeString},
					{Name: "login", Type: PredicateTypeString, Tokenizers: []string{"exact", "term"}, Index: true},
					{Name: "rated", Type: PredicateTypeUID, Reverse: true, Count: true},
					{Name: "score", Type: PredicateTypeInt, List: true},
				},
			},
			out: `name: string .
login: string @index(exact, term) .
rated: uid @reverse @count .
score: [int] .`,
		},
	}

	for _, c := range cases {
		data, err := c.in.MarshalText()

		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if got, want := string(data), c.out; got != want {
			t.Errorf("%v is %q in string, want %q", c.in, got, want)
		}
	}
}
