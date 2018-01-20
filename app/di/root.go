package di

import (
	"github.com/izumin5210/dform/app/system"
)

// RootComponent contains dependencies for this app.
type RootComponent interface {
	SystemComponent
	DgraphComponent
	ServiceComponent
}

// New creates a new app.
func New(config *system.Config) RootComponent {
	systemComponent := newSystem(config)
	dgraphComponent := newDgraph(systemComponent)
	fileComponent := newFile(systemComponent)
	serviceComponent := newService(
		systemComponent,
		dgraphComponent,
		fileComponent,
	)
	return &rootComponent{
		SystemComponent:  systemComponent,
		DgraphComponent:  dgraphComponent,
		ServiceComponent: serviceComponent,
	}
}

type rootComponent struct {
	SystemComponent
	DgraphComponent
	ServiceComponent
}
