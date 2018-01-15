package component

import (
	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
)

// Dgraph containes dependencies for accessing to Dgraph.
type Dgraph interface {
	DgraphSchemaRepository() schema.Repository
}

func newDgraph(config *system.Config) Dgraph {
	return &dgraph{
		Grpc: newGrpc(config),
	}
}

type dgraph struct {
	Grpc
}

func (d *dgraph) DgraphSchemaRepository() schema.Repository {
	return repo.NewDgraphSchemaRepository(d.GetDgraphPool())
}
