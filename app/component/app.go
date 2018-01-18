package component

import (
	"github.com/izumin5210/dform/app/system"
)

// App contains dependencies for this app.
type App interface {
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
		Service: serviceComponent,
	}
}

type app struct {
	Service
}
