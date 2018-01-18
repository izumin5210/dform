package di

import "github.com/izumin5210/dform/app/system"

// System provides accessors for moodules in system package.
type System interface {
	Config() *system.Config
	UI() system.UI
}

func newSystem(config *system.Config) System {
	return &systemComponent{
		config: config,
	}
}

type systemComponent struct {
	config *system.Config
}

func (c *systemComponent) Config() *system.Config {
	return c.config
}

func (c *systemComponent) UI() system.UI {
	return system.NewUI(
		c.config.InReader,
		c.config.OutWriter,
		c.config.ErrWriter,
	)
}
