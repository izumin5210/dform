package schema

import (
	"context"
)

// Repository is an interface for operating Dgraph schema.
type Repository interface {
	GetSchema(context.Context) (*Schema, error)
}
