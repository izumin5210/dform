package main

import (
	"os"

	"github.com/izumin5210/dform/app/cmd"
	"github.com/izumin5210/dform/app/di"
	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/util/log"
)

var (
	// Name is application name
	Name string
	// Version is application version
	Version string
	// Revision describes current commit hash generated by `git describe --always`.
	Revision string

	inReader  = os.Stdin
	outWriter = os.Stdout
	errWriter = os.Stderr
)

func main() {
	os.Exit(run())
}

func run() int {
	defer log.Close()

	config := system.NewConfig(
		Name,
		Version,
		Revision,
		inReader,
		outWriter,
		errWriter,
	)
	rootComponent := di.New(config)
	c := cmd.New(rootComponent)
	err := c.Execute()
	if err != nil {
		log.Error("failed to execute command", "error", err)
		return 1
	}
	return 0
}
