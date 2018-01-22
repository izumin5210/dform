package system

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	input "github.com/tcnksm/go-input"

	"github.com/izumin5210/dform/util/log"
)

// UI is an interface to abstract interactions with users.
type UI interface {
	Output(string)
	Warn(string)
	Error(string)
	Confirm(string) (bool, error)
}

type uiImpl struct {
	in      io.Reader
	out     io.Writer
	err     io.Writer
	inputUI *input.UI
}

var (
	fprintlnWarn  = color.New(color.FgYellow).FprintlnFunc()
	fprintlnError = color.New(color.FgRed).FprintlnFunc()
)

// NewUI creates new UI object.
func NewUI(in io.Reader, out, err io.Writer) UI {
	return &uiImpl{
		in:  in,
		out: out,
		err: err,
		inputUI: &input.UI{
			Reader: in,
			Writer: out,
		},
	}
}

func (i *uiImpl) Output(msg string) {
	fmt.Fprintln(i.out, msg)
}

func (i *uiImpl) Warn(msg string) {
	fprintlnWarn(i.err, msg)
}

func (i *uiImpl) Error(msg string) {
	fprintlnError(i.err, msg)
}

func (i *uiImpl) Confirm(msg string) (bool, error) {
	ans, err := i.inputUI.Ask(fmt.Sprintf("%s [Y/n]", msg), &input.Options{
		HideOrder: true,
		Loop:      true,
		ValidateFunc: func(ans string) error {
			log.Debug("receive user input", "query", msg, "input", ans)
			if ans != "Y" && ans != "n" {
				return fmt.Errorf("input must be Y or n")
			}
			return nil
		},
	})
	if err != nil {
		return false, err
	}
	return ans == "Y", nil
}
