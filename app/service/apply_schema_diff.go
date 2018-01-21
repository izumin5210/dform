package service

import (
	"context"

	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/util/log"
	"golang.org/x/sync/errgroup"
)

// ApplySchemaDiffService applies diff between 2 schemata taken from given repositories.
type ApplySchemaDiffService interface {
	Perform(context.Context) error
}

// NewApplySchemaDiffService creates a ApplySchemaDiffService instance.
func NewApplySchemaDiffService(base, target schema.Repository, ui system.UI) ApplySchemaDiffService {
	return &applySchemaDiffService{
		ui:         ui,
		baseRepo:   base,
		targetRepo: target,
	}
}

type applySchemaDiffService struct {
	baseRepo, targetRepo schema.Repository
	ui                   system.UI
}

func (s *applySchemaDiffService) Perform(ctx context.Context) error {
	var s1, s2 *schema.Schema

	eg := errgroup.Group{}

	eg.Go(func() error {
		s, err := s.baseRepo.GetSchema(ctx)
		if err != nil {
			log.Error("failed to get schema from base schema repository", "error", err)
			return err
		}
		log.Debug("succeeded to get base schema", "schema", s)
		s1 = s
		return nil
	})
	eg.Go(func() error {
		s, err := s.targetRepo.GetSchema(ctx)
		if err != nil {
			log.Error("failed to get schema from target schema repository", "error", err)
			return err
		}
		log.Debug("succeeded to get target schema", "schema", s)
		s2 = s
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Error("failed to get schemata", "error", err)
		return err
	}

	diff := schema.MakeDiff(s1, s2)
	if diff.Empty() {
		log.Debug("diff is empty")
		return nil
	}

	log.Debug("succeeded to get diff", "diff", diff)

	err = s.targetRepo.Update(ctx, diff)
	if err != nil {
		log.Error("failed to update schema for target schema repository", "error", err)
		return err
	}

	return nil
}
