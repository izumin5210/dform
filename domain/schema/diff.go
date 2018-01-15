package schema

import (
	"reflect"
)

// Diff represents a changes in 2 schemata
type Diff struct {
	Inserted []*PredicateSchema
	Deleted  []*PredicateSchema
	Modified []*ModifiedPredicate
}

// ModifiedPredicate represents a changes in 2 predicates
type ModifiedPredicate struct {
	From, To *PredicateSchema
}

// MakeDiff creates Diff in 2 schemata
func MakeDiff(s1, s2 *Schema) *Diff {
	predByName := make(map[string]*PredicateSchema, len(s1.Predicates))
	for _, p := range s1.Predicates {
		predByName[p.Name] = p
	}

	diff := &Diff{}

	for _, p2 := range s2.Predicates {
		if p1, ok := predByName[p2.Name]; ok {
			delete(predByName, p2.Name)
			if !reflect.DeepEqual(p1, p2) {
				diff.Modified = append(diff.Modified, &ModifiedPredicate{From: p1, To: p2})
			}
		} else {
			diff.Inserted = append(diff.Inserted, p2)
		}
	}

	for _, p := range s1.Predicates {
		if _, ok := predByName[p.Name]; ok {
			diff.Deleted = append(diff.Deleted, p)
		}
	}

	return diff
}
