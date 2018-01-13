package main

import (
	"fmt"
	"log"
	"os"

	"github.com/izumin5210/dform/app/cmd"
)

var (
	// Name is application name
	Name string
	// Version is application version
	Version string
	// Revision describes current commit hash generated by `git describe --always`.
	Revision string
)

func main() {
	c := cmd.New(Name, Version, Revision)
	err := c.Execute()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to execute command: %v", err))
		os.Exit(-1)
	}
}
