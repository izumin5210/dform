package di

import (
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
)

// DgraphComponent containes dependencies for accessing to Dgraph.
type DgraphComponent interface {
	DgraphSchemaRepository() schema.Repository
}

func newDgraph(system SystemComponent) DgraphComponent {
	return &dgraphComponent{
		GrpcComponent: newGrpc(system),
	}
}

type dgraphComponent struct {
	GrpcComponent
}

func (d *dgraphComponent) DgraphSchemaRepository() schema.Repository {
	return repo.NewDgraphSchemaRepository(d.GetDgraphPool())
}
