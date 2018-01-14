package component

import (
	"github.com/izumin5210/dform/app/system"
)

// App contains dependencies for this app.
type App interface {
	Dgraph
	Config() *system.Config
}

// New creates a new app.
func New(config *system.Config) App {
	return &app{
		Dgraph: newDgraph(config),
		config: config,
	}
}

type app struct {
	Dgraph
	config *system.Config
}

func (a *app) Config() *system.Config {
	return a.config
}
