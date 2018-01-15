package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/izumin5210/dform/domain/schema"
	"github.com/spf13/afero"
)

type fileSchemaRepository struct {
	fs         afero.Fs
	schemaPath string
}

// NewFileSchemaRepository creates new schema repository interface for accessing Dgraph.
func NewFileSchemaRepository(fs afero.Fs, schemaPath string) schema.Repository {
	return &fileSchemaRepository{
		fs:         fs,
		schemaPath: schemaPath,
	}
}

func (r *fileSchemaRepository) GetSchema(ctx context.Context) (*schema.Schema, error) {
	if info, err := r.fs.Stat(r.schemaPath); err != nil {
		return nil, fmt.Errorf("%s does not exist", r.schemaPath)
	} else if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory", r.schemaPath)
	}

	body, err := afero.ReadFile(r.fs, r.schemaPath)

	if err != nil {
		return nil, fmt.Errorf("failed to open %s", r.schemaPath)
	}

	s := &schema.Schema{}
	err = s.UnmarshalText(body)

	if err != nil {
		return nil, fmt.Errorf("failed to parse %s contents: %v", r.schemaPath, err)
	}

	return s, nil
}

func (r *fileSchemaRepository) Update(ctx context.Context, diff *schema.Diff) error {
	return errors.New("not yet implemented")
}
