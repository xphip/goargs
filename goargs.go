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
	cmd        *Cmd
	helperFlag string
}

func New() *GoArgs {
	ga := &GoArgs{
		cmd: &Cmd{
			subCmd: make(map[string]*Cmd),
		},
		helperFlag: "help",
	}
	return ga
}

func (ga *GoArgs) Add(name string) *Cmd {
	ga.cmd.Add(name)
	return ga.cmd.subCmd[name]
}

func (ga *GoArgs) Parse() error {

	var _args []string
	var cmd = ga.cmd

	if len(os.Args) <= 1 {
		return errors.New("usage here")
	}

	_args = os.Args[1:]

	if _cmd, ok := cmd.subCmd[_args[0]]; !ok {
		return errors.New(UnknownCommand)
	} else {
		cmd = _cmd
	}

	for len(_args) > 0 && len(cmd.subCmd) != 0 {
		if _cmd, ok := cmd.subCmd[_args[0]]; ok {
			cmd = _cmd
			_args = _args[1:]
		} else {
			return errors.New(UnknownCommand)
		}
	}

	if cmd.exec == nil {
		return nil
	}

	args, err := cmd.parseArgs(_args)
	if err != nil {
		return err
	}

	return cmd.exec(args)
}
