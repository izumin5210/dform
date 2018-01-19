package di

import (
	"github.com/spf13/afero"

	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
)

// FileComponent containes dependencies for accessing to local filesystem.
type FileComponent interface {
	FileSchemaRepository() schema.Repository
}

func newFile(system SystemComponent) FileComponent {
	return &fileComponent{
		SystemComponent: system,
		fs:              afero.NewOsFs(),
	}
}

type fileComponent struct {
	SystemComponent
	fs         afero.Fs
	schemaRepo schema.Repository
}

func (f *fileComponent) FileSchemaRepository() schema.Repository {
	if f.schemaRepo == nil {
		f.schemaRepo = repo.NewFileSchemaRepository(f.fs, f.Config().GetSchemaPath())
	}

	return f.schemaRepo
}
