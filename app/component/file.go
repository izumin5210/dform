package component

import (
	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
	"github.com/spf13/afero"
)

// File containes dependencies for accessing to local filesystem.
type File interface {
	FileSchemaRepository() schema.Repository
}

func newFile(config *system.Config) File {
	return &file{
		config: config,
		fs:     afero.NewOsFs(),
	}
}

type file struct {
	config     *system.Config
	fs         afero.Fs
	schemaRepo schema.Repository
}

func (f *file) FileSchemaRepository() schema.Repository {
	if f.schemaRepo == nil {
		f.schemaRepo = repo.NewFileSchemaRepository(f.fs, f.config.GetSchemaPath())
	}

	return f.schemaRepo
}
