package schema

import "fmt"

// PredicateType represents schema types in Dgraph.
type PredicateType int

// PredicateType enum values
// See https://docs.dgraph.io/query-language/#schemav
const (
	PredicateTypeDefault PredicateType = iota
	PredicateTypeInt
	PredicateTypeFloat
	PredicateTypeString
	PredicateTypeBool
	PredicateTypeDateTime
	PredicateTypeGeo
	PredicateTypePassword
	PredicateTypeUID
)

var (
	nameByPredType map[PredicateType]string
	predTypeByName map[string]PredicateType
)

func init() {
	nameByPredType = map[PredicateType]string{
		PredicateTypeDefault:  "default",
		PredicateTypeInt:      "int",
		PredicateTypeFloat:    "float",
		PredicateTypeString:   "string",
		PredicateTypeBool:     "bool",
		PredicateTypeDateTime: "dateTime",
		PredicateTypeGeo:      "geo",
		PredicateTypePassword: "password",
		PredicateTypeUID:      "uid",
	}
	predTypeByName = map[string]PredicateType{}
	for t, n := range nameByPredType {
		predTypeByName[n] = t
	}
}

func (t PredicateType) String() string {
	return nameByPredType[t]
}

// PredicateTypeOf returns PredicateType corresponds to given name
func PredicateTypeOf(name string) (PredicateType, error) {
	t, ok := predTypeByName[name]
	if !ok {
		return PredicateTypeDefault, fmt.Errorf("unexpected type: %s", name)
	}
	return t, nil
}
