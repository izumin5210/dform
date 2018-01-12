package schema

import (
	"strings"
)

// PredicateSchema represents predicate schema in Dgraph.
type PredicateSchema struct {
	Name       string
	Type       PredicateType
	Tokenizers []string
	Index      bool
	Reverse    bool
	Count      bool
	List       bool
}

func (s *PredicateSchema) String() string {
	cols := []string{}
	cols = append(cols, s.Name+":")
	if s.List {
		cols = append(cols, "["+s.Type.String()+"]")
	} else {
		cols = append(cols, s.Type.String())
	}
	if s.Index {
		cols = append(cols, "@index("+strings.Join(s.Tokenizers, ", ")+")")
	}
	if s.Reverse {
		cols = append(cols, "@reverse")
	}
	if s.Count {
		cols = append(cols, "@count")
	}
	cols = append(cols, ".")

	return strings.Join(cols, " ")
}
