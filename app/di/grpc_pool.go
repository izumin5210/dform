package di

import (
	"github.com/izumin5210/dform/infra/repo"
)

// Grpc contains GrpcPool instances
type Grpc interface {
	GetDgraphPool() repo.GrpcPool
}

func newGrpc(system System) Grpc {
	return &grpc{
		System: system,
	}
}

type grpc struct {
	System
	pool repo.GrpcPool
}

func (g *grpc) GetDgraphPool() repo.GrpcPool {
	if g.pool == nil {
		g.pool = repo.NewGrpcPool(g.Config().GetDgraphURL())
	}
	return g.pool
}
