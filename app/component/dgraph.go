package component

import (
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
)

// Dgraph containes dependencies for accessing to Dgraph.
type Dgraph interface {
	DgraphSchemaRepository() schema.Repository
}

func newDgraph(system System) Dgraph {
	return &dgraph{
		Grpc: newGrpc(system),
	}
}

type dgraph struct {
	Grpc
}

func (d *dgraph) DgraphSchemaRepository() schema.Repository {
	return repo.NewDgraphSchemaRepository(d.GetDgraphPool())
}
