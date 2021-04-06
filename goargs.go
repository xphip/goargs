package goargs

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	UsageComplement = "\nRun 'help' for usage."
	MissingParameter = "error: missing parameter" + UsageComplement
	UnknownCommand = "error: unknown command" + UsageComplement
)

// GoArgs is the framework instance.
type GoArgs struct {
	cmd        *Cmd
	helperFlag string
	template   func (UsageList) error
}

// New returns a new instance of GoArgs.
func New() *GoArgs {
	ga := &GoArgs{
		cmd: &Cmd{
			subCmd: make(map[string]*Cmd),
		},
		helperFlag: "help",
		template:   nil,
	}
	return ga
}

// Add creates and attach a new Cmd instance.
func (ga *GoArgs) Add(name string) *Cmd {
	ga.cmd.Add(name)
	return ga.cmd.subCmd[name]
}

// Parse parses the list of arguments.
// Must be called after all commands have been defined.
func (ga *GoArgs) Parse() error {
	_args := os.Args[1:]

	if len(_args) == 0 {
		return ga.parseUsage(_args)
	}

	if _args[0] == ga.helperFlag {
		return ga.parseUsage(_args[1:])
	}

	cmd, args, err := ga.parseCmd(_args)
	if err != nil {
		return err
	}

	if cmd.exec == nil {
		return nil
	}

	return cmd.exec(args)
}

// SetTemplate defines the function that will be called to generate the customized output.
// The function receives a UsageList containing information needed to assemble the output.
func (ga *GoArgs) SetTemplate(templateFn func (UsageList) error) *GoArgs {
	ga.template = templateFn
	return ga
}

// SetHelperFlag sets the flag that will invoke usage method.
func (ga *GoArgs) SetHelperFlag(helperFlag string) *GoArgs {
	ga.helperFlag = helperFlag
	return ga
}

func (ga *GoArgs) parseCmd(args []string) (*Cmd, *Args, error) {
	cmd := ga.cmd

	if _, ok := cmd.subCmd[args[0]]; !ok {
		return nil, nil, errors.New(UnknownCommand)
	}

	for len(cmd.subCmd) > 0 {
		if len(args) == 0 {
			return nil, nil, errors.New(MissingParameter)
		}

		if _, ok := cmd.subCmd[args[0]]; !ok {
			return nil, nil, errors.New(UnknownCommand)
		}

		cmd = cmd.subCmd[args[0]]
		args = args[1:]
	}

	_args, err := cmd.parseArgs(args)
	if err != nil {
		return nil, nil, err
	}

	return cmd, _args, nil
}

func (ga *GoArgs) parseUsage(args []string) error {
	cmd := ga.cmd
	argsLength := len(args)

	if argsLength > 0 {
		if _, ok := cmd.subCmd[args[0]]; !ok {
			return errors.New(UnknownCommand)
		}
	}

	usageList := UsageList{
		FileName: filepath.Base(os.Args[0]),
		Path: "",
		SpacingLength: 0,
		List: make([]*Usage, 0),
	}

	for c := 0; c < argsLength; c++ {
		index := args[c]
		if _, ok := cmd.subCmd[index]; !ok {
			return errors.New(UnknownCommand)
		}

		cmd = cmd.subCmd[index]
		usageList.Path += cmd.name + " "
	}

	usageList.CurrentUsage = cmd.usage

	for _, cmdFlags := range cmd.subCmd {
		usageList.List = append(usageList.List, &Usage{
			flag: cmdFlags.name,
			desc: cmdFlags.usage,
		})

		if len(cmdFlags.name) > usageList.SpacingLength {
			usageList.SpacingLength = len(cmdFlags.name)
		}
	}

	if ga.template != nil {
		return ga.template(usageList)
	}

	return defaultTemplate(usageList)
}