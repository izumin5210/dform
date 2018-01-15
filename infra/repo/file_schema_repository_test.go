package repo

import (
	"context"
	"testing"

	"github.com/spf13/afero"
)

func Test_FileSchemaRepository_GetSchema(t *testing.T) {
	appFS := afero.NewMemMapFs()
	schemaPath := "dgraph.schema"
	repo := NewFileSchemaRepository(appFS, schemaPath)

	t.Run("when schema file does not exist", func(t *testing.T) {
		s, err := repo.GetSchema(context.Background())

		if err == nil {
			t.Error("GetSchema() should return an error")
		}

		if s != nil {
			t.Errorf("GetSchema() returned %v, want nil", s)
		}
	})

	t.Run("when schema file is directory", func(t *testing.T) {
		appFS.MkdirAll(schemaPath, 0755)
		defer appFS.RemoveAll(schemaPath)

		s, err := repo.GetSchema(context.Background())

		if err == nil {
			t.Error("GetSchema() should return an error")
		}

		if s != nil {
			t.Errorf("GetSchema() returned %v, want nil", s)
		}
	})

	t.Run("when schema file format is invalid", func(t *testing.T) {
		schemaStr := "name: string"
		afero.WriteFile(appFS, schemaPath, []byte(schemaStr), 0644)

		s, err := repo.GetSchema(context.Background())

		if err == nil {
			t.Error("GetSchema() should return an error")
		}

		if s != nil {
			t.Errorf("GetSchema() returned %v, want nil", s)
		}
	})

	t.Run("when schema file is valid", func(t *testing.T) {
		schemaStr := `name: string @index(term) .
user_id: int @index(int) .`
		afero.WriteFile(appFS, schemaPath, []byte(schemaStr), 0644)
		s, err := repo.GetSchema(context.Background())

		if err != nil {
			t.Errorf("GetSchema() should not return errors, but returned %v", err)
		}

		if s == nil {
			t.Error("GetSchema() should return a schema")
		} else {
			if got, want := len(s.Predicates), 2; got != want {
				t.Errorf("GetSchema() returned schema with %d predicates, want %d predicates", got, want)
			}
		}
	})
}
