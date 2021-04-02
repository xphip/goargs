package goargs

import "errors"

// Cmd is the structure of each subcommand.
type Cmd struct {
	name    string
	usage   string
	mapping []string
	subCmd  map[string]*Cmd
	exec    func (args *Args) error
}

// Add adds a subcommand to the parent node and is returned.
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

// Usage defines the current subcommand description.
func (cmd *Cmd) Usage(usage string) *Cmd {
	cmd.usage = usage
	return cmd
}

// Map maps the arguments that are not part of the subcommand tree.
// The mapping order is defined from left to right.
// The number of arguments mapped must be less than or equal to the number of arguments informed or MissingParameter warning will be returned.
func (cmd *Cmd) Map(mapping []string) *Cmd {
	cmd.mapping = mapping
	return cmd
}

// Exec defines the function for the subcommand.
// If subCmd still has subcommands, Exec will be ignored.
func (cmd *Cmd) Exec(fn func (args *Args) error) {
	cmd.exec = fn
}

// parseArgs parses the arguments and returns an Args type.
func (cmd *Cmd) parseArgs(a []string) (*Args, error) {
	a = a[1:]

	if len(cmd.mapping) > len(a) {
		return nil, errors.New(MissingParameter)
	}

	args := newArgs()

	for _, arg := range cmd.mapping {
		args.addMapped(arg, a[0])
		a = a[1:]
	}

	for _, arg := range a {
		args.addUnmapped(arg)
	}

	return args, nil
}