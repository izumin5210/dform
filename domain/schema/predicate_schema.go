package schema

// PredicateSchema represents predicate schema in Dgraph.
type PredicateSchema struct {
	Name       string
	Type       PredicateType
	Tokenizers []string
	Index      bool
	Reverse    bool
}
