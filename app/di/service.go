package di

import "github.com/izumin5210/dform/app/service"

// Service contains service implementations.
type Service interface {
	ShowSchemaDiffService() service.ShowSchemaDiffService
}

func newService(system System, dgraph Dgraph, file File) Service {
	return &serviceComponent{
		System: system,
		Dgraph: dgraph,
		File:   file,
	}
}

type serviceComponent struct {
	System
	Dgraph
	File
}

func (c *serviceComponent) ShowSchemaDiffService() service.ShowSchemaDiffService {
	return service.NewShowSchemaDiffService(
		c.Dgraph.DgraphSchemaRepository(),
		c.File.FileSchemaRepository(),
		c.System.UI(),
	)
}
