package di

import "github.com/izumin5210/dform/app/service"

// ServiceComponent contains service implementations.
type ServiceComponent interface {
	ShowSchemaDiffService() service.ShowSchemaDiffService
}

func newService(system SystemComponent, dgraph DgraphComponent, file FileComponent) ServiceComponent {
	return &serviceComponent{
		SystemComponent: system,
		DgraphComponent: dgraph,
		FileComponent:   file,
	}
}

type serviceComponent struct {
	SystemComponent
	DgraphComponent
	FileComponent
}

func (c *serviceComponent) ShowSchemaDiffService() service.ShowSchemaDiffService {
	return service.NewShowSchemaDiffService(
		c.DgraphSchemaRepository(),
		c.FileSchemaRepository(),
		c.UI(),
	)
}
