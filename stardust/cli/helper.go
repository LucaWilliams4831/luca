package cli

import (
	"fmt"

	"github.com/fatih/color"
)

type Args struct {
	args []string
}

// len is the lenth of the remaining command line arguments
func (a *Args) len() int {
	return len(a.args)
}

// pop returns the first of the remaining command line arguments to the caller
// and deletes if from the arguments stack
func (a *Args) pop() string {
	if len(a.args) == 0 {
		return ""
	}

	item := a.args[0]
	a.args = a.args[1:]

	return item
}

func ErrorFromString(str string) error {
	return fmt.Errorf(color.RedString(str))
}
