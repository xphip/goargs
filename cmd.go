package goargs

type SubCmd struct {
	mapped  map[string]*Cmd
	list  []*Cmd
}

func (sc SubCmd) GetIndex(index int) *Cmd {
	if index >= len(sc.list) || index <= 0 {
		return nil
	}
	return sc.list[index]
}

// Cmd is the instance for a command.
type Cmd struct {
	name    string
	usage   string
	mapping []string
	subCmd  SubCmd
	exec    func (args *Args) error
}

// Add attach a new Cmd instance to the parent node and returns it.
func (cmd *Cmd) Add(name string) *Cmd {
	subCmd := &Cmd{
		name:    name,
		usage:   "",
		mapping: make([]string, 0),
		subCmd:  SubCmd{
			mapped: make(map[string]*Cmd),
			list:   make([]*Cmd, 0),
		},
		exec:    nil,
	}
	cmd.subCmd.mapped[name] = subCmd
	cmd.subCmd.list = append(cmd.subCmd.list, subCmd)
	return subCmd
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

func (cmd *Cmd) parseArgs(a []string) (*Args, bool) {

	if len(cmd.mapping) > len(a) {
		return nil, false
	}

	args := newArgs()

	for _, arg := range cmd.mapping {
		args.addMapped(arg, a[0])
		a = a[1:]
	}

	for _, arg := range a {
		args.addUnmapped(arg)
	}

	return args, true
}