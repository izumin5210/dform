package schema

import (
	"strings"

	"github.com/dgraph-io/dgraph/protos/intern"
	dgraphschema "github.com/dgraph-io/dgraph/schema"
)

// Schema represents predicate schemata in Dgrpah.
type Schema struct {
	Predicates []*PredicateSchema
}

// UnmarshalText implements encoding.TextUnmarshaler
func (s *Schema) UnmarshalText(data []byte) error {
	updates, err := dgraphschema.Parse(string(data))
	if err != nil {
		return err
	}
	preds := make([]*PredicateSchema, 0, len(updates))
	for _, u := range updates {
		t, err := PredicateTypeOf(u.GetValueType().String())
		if err != nil {
			return err
		}
		d := u.GetDirective()
		pred := &PredicateSchema{
			Name:       u.GetPredicate(),
			Type:       t,
			Index:      d == intern.SchemaUpdate_INDEX,
			Reverse:    d == intern.SchemaUpdate_REVERSE,
			Tokenizers: u.GetTokenizer(),
			Count:      u.GetCount(),
			List:       u.GetList(),
		}
		preds = append(preds, pred)
	}
	s.Predicates = preds
	return nil
}

// MarshalText implements encoding.TextMarshaler
func (s *Schema) MarshalText() ([]byte, error) {
	lines := make([]string, 0, len(s.Predicates))
	for _, pred := range s.Predicates {
		lines = append(lines, pred.String())
	}
	return []byte(strings.Join(lines, "\n")), nil
}
