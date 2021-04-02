package goargs

// Args is the structure for storing the parsed arguments.
type Args struct {
	mapped  map[string]*Arg
	list    []*Arg
}

// newArgs returns a new structure of Args.
func newArgs() *Args {
	return &Args{
		mapped: make(map[string]*Arg),
		list:   make([]*Arg, 0),
	}
}

// addMapped maps a string-type index to a value.
func (args *Args) addMapped(name string, value string) {
	val := Arg(value)
	args.mapped[name] = &val
}

// addMapped adds an unmapped argument.
func (args *Args) addUnmapped(value string) {
	val := Arg(value)
	args.list = append(args.list, &val)
}

// Get returns a mapped argument.
func (args *Args) Get(name string) *Arg {
	arg, ok := args.mapped[name]
	if !ok {
		return nil
	}
	return arg
}

// GetMapped returns only the mapped arguments.
func (args *Args) GetMapped() map[string]*Arg {
	return args.mapped
}

// GetUnmapped returns only unmapped arguments.
func (args *Args) GetUnmapped() []*Arg {
	return args.list
}

// GetPos returns an unmapped argument by the position.
func (args *Args) GetPos(position int) *Arg {
	if position < 0 || position > len(args.list) - 1 {
		return nil
	}
	return args.list[position]
}