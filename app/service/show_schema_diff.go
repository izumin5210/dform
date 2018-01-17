package service

import (
	"context"
	"strings"

	"github.com/fatih/color"

	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/util/log"
)

// ShowSchemaDiffService shows diff between 2 schemata taken from given repositories.
type ShowSchemaDiffService interface {
	Perform(context.Context) error
}

// NewShowSchemaDiffService creates a ShowSchemaDiffService instance.
func NewShowSchemaDiffService(repo1, repo2 schema.Repository, ui system.UI) ShowSchemaDiffService {
	return &showSchemaDiffService{
		repo1: repo1,
		repo2: repo2,
		ui:    ui,
	}
}

type showSchemaDiffService struct {
	repo1, repo2 schema.Repository
	ui           system.UI
}

func (s *showSchemaDiffService) Perform(ctx context.Context) error {
	s1, err := s.repo1.GetSchema(context.Background())
	if err != nil {
		log.Error("failed to get schema from Dgraph", "error", err)
		return err
	}
	s2, err := s.repo2.GetSchema(context.Background())
	if err != nil {
		log.Error("failed to get schema from filesystem", "error", err)
		return err
	}

	diff := schema.MakeDiff(s1, s2)
	red := color.New(color.FgRed).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()

	if len(diff.Inserted) > 0 {
		msgs := make([]string, 0, len(diff.Inserted)+3)
		msgs = append(msgs, "Added predicates:", "")
		for _, pred := range diff.Inserted {
			msgs = append(msgs, green("    + %s", pred))
		}
		msgs = append(msgs, "")
		s.ui.Output(strings.Join(msgs, "\n"))
	}

	if len(diff.Deleted) > 0 {
		msgs := make([]string, 0, len(diff.Deleted)+3)
		msgs = append(msgs, "Dropped predicates:", "")
		for _, pred := range diff.Deleted {
			msgs = append(msgs, red("    - %s", pred))
		}
		msgs = append(msgs, "")
		s.ui.Output(strings.Join(msgs, "\n"))
	}

	if len(diff.Modified) > 0 {
		msgs := make([]string, 0, len(diff.Modified)+3)
		msgs = append(msgs, "Modified predicates:", "")
		for _, pair := range diff.Modified {
			msgs = append(msgs, red("    - %s", pair.From))
			msgs = append(msgs, green("    + %s", pair.To))
		}
		s.ui.Output(strings.Join(msgs, "\n"))
	}

	return nil
}
