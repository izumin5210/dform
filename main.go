package main

import (
	"os"

	"github.com/izumin5210/dform/cmd"
)

func main() {
	code := cmd.Execute()
	os.Exit(code)
}
