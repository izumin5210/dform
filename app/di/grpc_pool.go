package di

import (
	"github.com/izumin5210/dform/infra/repo"
)

// GrpcComponent contains GrpcPool instances
type GrpcComponent interface {
	GetDgraphPool() repo.GrpcPool
}

func newGrpc(system SystemComponent) GrpcComponent {
	return &grpcComponent{
		SystemComponent: system,
	}
}

type grpcComponent struct {
	SystemComponent
	pool repo.GrpcPool
}

func (g *grpcComponent) GetDgraphPool() repo.GrpcPool {
	if g.pool == nil {
		g.pool = repo.NewGrpcPool(g.Config().GetDgraphURL())
	}
	return g.pool
}
