package goargs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: doc
const (
	UsageComplement = "\nRun 'help' for usage."
	MissingParameter = "error: missing parameter" + UsageComplement
	UnknownCommand = "error: unknown command" + UsageComplement
	UnknownError = "error: unknown error" + UsageComplement
)

// TODO: doc
type GoArgs struct {
	cmd        *Cmd
	helperFlag string
	template   func (UsageList) error
}

// TODO: doc
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

// Add adds and returns a new subcommand tree.
func (ga *GoArgs) Add(name string) *Cmd {
	ga.cmd.Add(name)
	return ga.cmd.subCmd[name]
}

// Parse parses the list of arguments.
// Must be called after all commands have been defined.
func (ga *GoArgs) Parse() error {
	_args := os.Args[1:]

	if len(_args) == 0 || _args[0] == ga.helperFlag {
		return ga.parseUsage(_args)
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

func (ga *GoArgs) parseCmd(args []string) (*Cmd, *Args, error) {
	cmd := ga.cmd

	if _, ok := cmd.subCmd[args[0]]; !ok {
		return nil, nil, errors.New(UnknownCommand)
	}

	cmd = cmd.subCmd[args[0]]

	for len(args) > 0 && len(cmd.subCmd) != 0 {
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

	if len(args) > 1 {
		if args[0] == ga.helperFlag {
			args = args[1:]
		}

		if _, ok := cmd.subCmd[args[0]]; !ok {
			return errors.New(UnknownCommand)
		}
		cmd = cmd.subCmd[args[0]]
	}

	usageList := UsageList{
		FileName: filepath.Base(os.Args[0]),
		Path: "",
		SpacingLength: 0,
		StartSpacing: fmt.Sprintf("%4s", ""),
		BetweenSpacing: fmt.Sprintf("%4s", ""),
		List: make([]*Usage, 0),
	}

	for c:=0; c < len(args) - 1; c++ {
		usageList.Path += cmd.name + " "
		if _, ok := cmd.subCmd[args[c]]; !ok {
			return errors.New(UnknownCommand)
		}
		cmd = cmd.subCmd[args[c]]
	}

	if len(cmd.subCmd) != 0 {
		for _, cmdFlags := range cmd.subCmd {
			usageList.List = append(usageList.List, &Usage{
				flag: cmdFlags.name,
				desc: cmdFlags.usage,
			})
			if len(cmdFlags.name) > usageList.SpacingLength {
				usageList.SpacingLength = len(cmdFlags.name)
			}
		}
	} else {
		usageList.List = append(usageList.List, &Usage{
			flag: "",
			desc: cmd.usage,
		})
		usageList.StartSpacing = ""
		usageList.SpacingLength = 0
	}

	if ga.template != nil {
		return ga.template(usageList)
	}

	return defaultTemplate(usageList)
}