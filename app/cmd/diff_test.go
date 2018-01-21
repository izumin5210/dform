package cmd

import (
	"context"
	"errors"
	"testing"

	"github.com/izumin5210/dform/app/di"
	"github.com/izumin5210/dform/app/service"
	"github.com/izumin5210/dform/app/system"
	"github.com/spf13/cobra"
)

func Test_Diff(t *testing.T) {
	t.Run("when succeeded in applying", func(t *testing.T) {
		var showCalled, applyCalled, confirmCalled bool

		cmd := newTestDiffCommand(
			func(context.Context) error {
				showCalled = true
				return nil
			},
			func(context.Context) error {
				applyCalled = true
				return nil
			},
			func(msg string) (bool, error) {
				confirmCalled = true
				return true, nil
			},
		)

		err := cmd.Execute()

		if err != nil {
			t.Errorf("Execute() returned %v, want nil", err)
		}

		if !showCalled {
			t.Errorf("ShowSchemaDiffService#Perform() should be called")
		}

		if !confirmCalled {
			t.Errorf("UI#Confirm() should be called")
		}

		if !applyCalled {
			t.Errorf("ApplySchemaDiffService#Perform() should be called")
		}
	})

	t.Run("when stop applying", func(t *testing.T) {
		var showCalled, applyCalled, confirmCalled bool

		cmd := newTestDiffCommand(
			func(context.Context) error {
				showCalled = true
				return nil
			},
			func(context.Context) error {
				applyCalled = true
				return nil
			},
			func(msg string) (bool, error) {
				confirmCalled = true
				return false, nil
			},
		)

		err := cmd.Execute()

		if err != nil {
			t.Errorf("Execute() returned %v, want nil", err)
		}

		if !showCalled {
			t.Errorf("ShowSchemaDiffService#Perform() should be called")
		}

		if !confirmCalled {
			t.Errorf("UI#Confirm() should be called")
		}

		if applyCalled {
			t.Errorf("ApplySchemaDiffService#Perform() should not be called")
		}
	})

	t.Run("when failed to show diff", func(t *testing.T) {
		var showCalled, applyCalled, confirmCalled bool

		cmd := newTestDiffCommand(
			func(context.Context) error {
				showCalled = true
				return errors.New("an error")
			},
			func(context.Context) error {
				applyCalled = true
				return nil
			},
			func(msg string) (bool, error) {
				confirmCalled = true
				return false, nil
			},
		)

		err := cmd.Execute()

		if err == nil {
			t.Error("Execute() should return an error")
		}

		if !showCalled {
			t.Errorf("ShowSchemaDiffService#Perform() should be called")
		}

		if confirmCalled {
			t.Errorf("UI#Confirm() should not be called")
		}

		if applyCalled {
			t.Errorf("ApplySchemaDiffService#Perform() should not be called")
		}
	})

	t.Run("when failed to confirm", func(t *testing.T) {
		var showCalled, applyCalled, confirmCalled bool

		cmd := newTestDiffCommand(
			func(context.Context) error {
				showCalled = true
				return nil
			},
			func(context.Context) error {
				applyCalled = true
				return nil
			},
			func(msg string) (bool, error) {
				confirmCalled = true
				return false, errors.New("an error")
			},
		)

		err := cmd.Execute()

		if err == nil {
			t.Error("Execute() should return an error")
		}

		if !showCalled {
			t.Errorf("ShowSchemaDiffService#Perform() should be called")
		}

		if !confirmCalled {
			t.Errorf("UI#Confirm() should be called")
		}

		if applyCalled {
			t.Errorf("ApplySchemaDiffService#Perform() should not be called")
		}
	})

	t.Run("when failed to apply diff", func(t *testing.T) {
		var showCalled, applyCalled, confirmCalled bool

		cmd := newTestDiffCommand(
			func(context.Context) error {
				showCalled = true
				return nil
			},
			func(context.Context) error {
				applyCalled = true
				return errors.New("an error")
			},
			func(msg string) (bool, error) {
				confirmCalled = true
				return true, nil
			},
		)

		err := cmd.Execute()

		if err == nil {
			t.Error("Execute() should return an error")
		}

		if !showCalled {
			t.Errorf("ShowSchemaDiffService#Perform() should be called")
		}

		if !confirmCalled {
			t.Errorf("UI#Confirm() should be called")
		}

		if !applyCalled {
			t.Errorf("ApplySchemaDiffService#Perform() should be called")
		}
	})
}

// Fake implementations
//================================================================
func newTestDiffCommand(
	fakeShow func(context.Context) error,
	fakeApply func(context.Context) error,
	fakeConfirm func(string) (bool, error),
) *cobra.Command {
	component := &fakeRootComponent{
		fakeShowSchemaDiffService: func() service.ShowSchemaDiffService {
			return &fakeShowSchemaDiffService{
				fakePerform: fakeShow,
			}
		},
		fakeApplySchemaDiffService: func() service.ApplySchemaDiffService {
			return &fakeApplySchemaDiffService{
				fakePerform: fakeApply,
			}
		},
		fakeUI: func() system.UI {
			return &fakeUI{
				fakeConfirm: fakeConfirm,
			}
		},
	}
	return newDiffCommand(component)
}

type fakeRootComponent struct {
	di.RootComponent
	fakeShowSchemaDiffService  func() service.ShowSchemaDiffService
	fakeApplySchemaDiffService func() service.ApplySchemaDiffService
	fakeUI                     func() system.UI
}

func (c *fakeRootComponent) ShowSchemaDiffService() service.ShowSchemaDiffService {
	return c.fakeShowSchemaDiffService()
}

func (c *fakeRootComponent) ApplySchemaDiffService() service.ApplySchemaDiffService {
	return c.fakeApplySchemaDiffService()
}

func (c *fakeRootComponent) UI() system.UI {
	return c.fakeUI()
}

type fakeShowSchemaDiffService struct {
	fakePerform func(context.Context) error
}

func (s *fakeShowSchemaDiffService) Perform(c context.Context) error {
	return s.fakePerform(c)
}

type fakeApplySchemaDiffService struct {
	fakePerform func(context.Context) error
}

func (s *fakeApplySchemaDiffService) Perform(c context.Context) error {
	return s.fakePerform(c)
}

type fakeUI struct {
	system.UI
	fakeConfirm func(string) (bool, error)
}

func (u *fakeUI) Confirm(msg string) (bool, error) {
	return u.fakeConfirm(msg)
}
