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
		pool: repo.NewGrpcPool(config.GetDgraphURL()),
	}
}

type grpc struct {
	pool repo.GrpcPool
}

func (g *grpc) GetDgraphPool() repo.GrpcPool {
	return g.pool
}
