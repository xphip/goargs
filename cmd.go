package goargs

import "errors"

type Cmd struct {
	name   string
	usage  string
	subCmd CmdMap
	exec   func(args []string) error
}

type CmdMap map[string]*Cmd

func (cmd *Cmd) parse(args []string) error {
	var arg = ""
	if len(args) > 0 {
		arg = args[0]
	}

	if cmd.exec != nil {
		return cmd.exec(args)

	} else if len(args) == 0 {
		return errors.New(MissingParameter)

	} else if subCmd, ok := cmd.subCmd[arg]; ok {
		return subCmd.parse(args[1:])

	} else {
		return errors.New(UnknownCommand)
	}
}

func (cmd *Cmd) Add(name string) *Cmd {
	newCmd := &Cmd{
		name: name,
		subCmd: make(CmdMap),
	}
	cmd.subCmd[name] = newCmd
	return newCmd
}

func (cmd *Cmd) Usage(usage string) *Cmd {
	cmd.usage = usage
	return cmd
}

func (cmd *Cmd) Exec(fn func(args []string) error) *Cmd {
	cmd.exec = fn
	return cmd
}
