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
}

type ui struct {
	out io.Writer
	err io.Writer
}

var (
	fprintlnWarn  = color.New(color.FgYellow).FprintlnFunc()
	fprintlnError = color.New(color.FgYellow).FprintlnFunc()
)

// NewUI creates new UI object.
func NewUI(out, err io.Writer) UI {
	return &ui{
		out: out,
		err: err,
	}
}

func (i *ui) Output(msg string) {
	fmt.Fprintln(i.out, msg)
}

func (i *ui) Warn(msg string) {
	fprintlnWarn(i.err, msg)
}

func (i *ui) Error(msg string) {
	fprintlnError(i.err, msg)
}
