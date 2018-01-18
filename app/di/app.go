package di

import (
	"github.com/izumin5210/dform/app/system"
)

// App contains dependencies for this app.
type App interface {
	System
	Dgraph
	Service
}

// New creates a new app.
func New(config *system.Config) App {
	systemComponent := newSystem(config)
	dgraphComponent := newDgraph(systemComponent)
	fileComponent := newFile(systemComponent)
	serviceComponent := newService(
		systemComponent,
		dgraphComponent,
		fileComponent,
	)
	return &app{
		System:  systemComponent,
		Dgraph:  dgraphComponent,
		Service: serviceComponent,
	}
}

type app struct {
	System
	Dgraph
	Service
}
