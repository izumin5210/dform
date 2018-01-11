package schema

import "testing"

func Test_PredicateTypeOf(t *testing.T) {
	cases := []struct {
		in    string
		out   PredicateType
		isErr bool
	}{
		{in: "default", out: PredicateTypeDefault},
		{in: "int", out: PredicateTypeInt},
		{in: "float", out: PredicateTypeFloat},
		{in: "string", out: PredicateTypeString},
		{in: "bool", out: PredicateTypeBool},
		{in: "dateTime", out: PredicateTypeDateTime},
		{in: "geo", out: PredicateTypeGeo},
		{in: "password", out: PredicateTypePassword},
		{in: "uid", out: PredicateTypeUID},
		{in: "unknown", isErr: true},
	}

	for _, c := range cases {
		out, err := PredicateTypeOf(c.in)

		if got, want := err != nil, c.isErr; got != want {
			if c.isErr {
				t.Errorf("PredicateTypeOf(%s) should return an error", c.in)
			} else {
				t.Errorf("PredicateTypeOf(%s) should not return an error", c.in)
			}
		}

		if got, want := out, c.out; got != want {
			t.Errorf("PredicateTypeOf(%s) returned %v, want %v", c.in, got, want)
		}
	}
}
