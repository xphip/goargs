package goargs

type Args struct {
	mapped  map[string]*Arg
	list    []*Arg
}

func NewArgs() *Args {
	return &Args{
		mapped: make(map[string]*Arg),
		list:   make([]*Arg, 0),
	}
}

func (args *Args) AddMapped(name string, value string) {
	val := Arg(value)
	args.mapped[name] = &val
}

func (args *Args) AddUnmapped(value string) {
	val := Arg(value)
	args.list = append(args.list, &val)
}

func (args *Args) Get(name string) *Arg {
	arg, ok := args.mapped[name]
	if !ok {
		return nil
	}
	return arg
}

func (args *Args) GetMapped() map[string]*Arg {
	return args.mapped
}

func (args *Args) GetUnmapped() []*Arg {
	return args.list
}

func (args *Args) GetPos(position int) *Arg {
	if position <= 0 || position > args.Size() - 1 {
		return nil
	}
	return args.list[position]
}

func (args *Args) Size() int {
	return len(args.list)
}
