package goargs

import (
	"errors"
)

type Cmd struct {
	name    string
	usage   string
	mapping []string
	subCmd  map[string]*Cmd
	exec    func (args *Args) error
}

func (cmd *Cmd) Add(name string) *Cmd {
	cmd.subCmd[name] = &Cmd{
		name:    name,
		usage:   "",
		mapping: make([]string, 0),
		subCmd:  make(map[string]*Cmd),
		exec:    nil,
	}
	return cmd
}

func (cmd *Cmd) Usage(usage string) *Cmd {
	cmd.usage = usage
	return cmd
}

func (cmd *Cmd) Map(mapping []string) *Cmd {
	cmd.mapping = mapping
	return cmd
}

func (cmd *Cmd) Exec(fn func (args *Args) error) {
	cmd.exec = fn
}

func (cmd *Cmd) parseArgs(a []string) (*Args, error) {
	a = a[1:]

	if len(cmd.mapping) > len(a) {
		return nil, errors.New(MissingParameter)
	}

	args := NewArgs()

	for _, arg := range cmd.mapping {
		args.AddMapped(arg, a[0])
		a = a[1:]
	}

	for _, arg := range a {
		args.AddUnmapped(arg)
	}

	return args, nil
}