package system

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

// UI is an interface to abstract interactions with users.
type UI interface {
	Output(string)
	Warn(string)
	Error(string)
	Confirm(string) (bool, error)
}

type uiImpl struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

var (
	fprintlnWarn  = color.New(color.FgYellow).FprintlnFunc()
	fprintlnError = color.New(color.FgYellow).FprintlnFunc()
)

// NewUI creates new UI object.
func NewUI(in io.Reader, out, err io.Writer) UI {
	return &uiImpl{
		in:  in,
		out: out,
		err: err,
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
	return false, nil
}
