# dform
[![Build Status](https://travis-ci.org/izumin5210/dform.svg?branch=master)](https://travis-ci.org/izumin5210/dform)
[![codecov](https://codecov.io/gh/izumin5210/dform/branch/master/graph/badge.svg)](https://codecov.io/gh/izumin5210/dform)
[![GoDoc](https://godoc.org/github.com/izumin5210/dform?status.svg)](https://godoc.org/github.com/izumin5210/dform)
[![Go Report Card](https://goreportcard.com/badge/github.com/izumin5210/dform)](https://goreportcard.com/report/github.com/izumin5210/dform)
[![license](https://img.shields.io/github/license/izumin5210/dform.svg)](./LICENSE)

CLI to manage [Dgraph](https://dgraph.io/) schema.

## Usage
TBD

## Development

### with rid (run-in-docker)
You can develop dform with [rid](https://github.com/creasty/rid).
 
#### Getting started

```
# Bootstrap the project
$ rid bootstrap

# Start Dgraph server
$ rid dgraph start
```

#### Executing app

```
# `rid run` execute `go build` and `./bin/dform` in docker container
$ rid run
izumin5210dform_app_1 is up-to-date
make: Nothing to be done for 'all'.
CLI tool to manage Dgraph schema

Usage:
  dform [command]

Available Commands:
  export      Export schema information
  help        Help about any command
  version     Print version information

Flags:
      --config string   config file (default is $PWD/.dform.toml)
  -h, --help            help for dform

Use "dform [command] --help" for more information about a command.
```

#### Run tests

```
# Start Dgraph server for testing
$ rid dgraph test start

# Run tests
$ rid make test
```
