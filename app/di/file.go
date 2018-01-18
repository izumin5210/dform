package di

import (
	"github.com/spf13/afero"

	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
)

// File containes dependencies for accessing to local filesystem.
type File interface {
	FileSchemaRepository() schema.Repository
}

func newFile(system System) File {
	return &file{
		System: system,
		fs:     afero.NewOsFs(),
	}
}

type file struct {
	System
	fs         afero.Fs
	schemaRepo schema.Repository
}

func (f *file) FileSchemaRepository() schema.Repository {
	if f.schemaRepo == nil {
		f.schemaRepo = repo.NewFileSchemaRepository(f.fs, f.Config().GetSchemaPath())
	}

	return f.schemaRepo
}
