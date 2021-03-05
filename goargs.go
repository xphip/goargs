package goargs

import (
	"errors"
	"os"
)

var (
	UsageComplement = "\nRun 'help' for usage."
	MissingParameter = "error: missing parameter" + UsageComplement
	UnknownCommand = "error: unknown command" + UsageComplement
	UnknownError = "error: unknown error" + UsageComplement
)

type GoArgs struct {
	subCmd     CmdMap
	helperFlag string
}

func New() *GoArgs {
	return &GoArgs{
		subCmd: make(CmdMap),
		helperFlag: "help",
	}
}

func (a *GoArgs) Add(name string) *Cmd {
	newCmd := &Cmd{
		name: name,
		subCmd: make(CmdMap),
	}
	a.subCmd[name] = newCmd
	return newCmd
}

func (a *GoArgs) Parse() error {
	args := os.Args[1:]

	if len(args) == 0 {
		return errors.New(MissingParameter)

	} else if cmd, ok := a.subCmd[args[0]]; ok {
		return cmd.parse(args[1:])

	} else {
		return errors.New(UnknownCommand)
	}

}





