package component

import (
	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/infra/repo"
)

// Grpc contains GrpcPool instances
type Grpc interface {
	GetDgraphPool() repo.GrpcPool
}

func newGrpc(config *system.Config) Grpc {
	return &grpc{
		config: config,
	}
}

type grpc struct {
	config *system.Config
	pool   repo.GrpcPool
}

func (g *grpc) GetDgraphPool() repo.GrpcPool {
	if g.pool == nil {
		g.pool = repo.NewGrpcPool(g.config.GetDgraphURL())
	}
	return g.pool
}
