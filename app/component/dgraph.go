package component

import (
	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
	"google.golang.org/grpc"
)

// Dgraph containes dependencies for accessing to Dgraph.
type Dgraph interface {
	DgraphSchemaRepository() (schema.Repository, error)
}

func newDgraph(config *system.Config) Dgraph {
	return &dgraph{
		config: config,
	}
}

type dgraph struct {
	config     *system.Config
	conn       *grpc.ClientConn
	schemaRepo schema.Repository
}

func (d *dgraph) DgraphSchemaRepository() (schema.Repository, error) {
	if d.schemaRepo != nil {
		return d.schemaRepo, nil
	}
	conn, err := d.getConn()
	if err != nil {
		return nil, err
	}
	d.schemaRepo = repo.NewDgraphSchemaRepository(conn)
	return d.schemaRepo, nil
}

func (d *dgraph) getConn() (*grpc.ClientConn, error) {
	if d.conn != nil {
		return d.conn, nil
	}
	conn, err := grpc.Dial(d.config.GetDgraphURL(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	d.conn = conn
	return d.conn, nil
}
