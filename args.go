package goargs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	UsageComplement = "\nRun 'help' for usage."
	MissingParameter = "error: missing parameter" + UsageComplement
	UnknownCommand = "error: unknown command" + UsageComplement
	UnknownError = "error: unknown error" + UsageComplement
)

type Args map[string]*Arg

type Arg struct {
	Name        string
	Description string
	Args        Args
	Exec        func(args []string) error
}

func (a *Args) AddArgs(args ...*Arg) *Args {

	for _, arg := range args {
		(*a)[arg.Name] = arg
	}

	return a
}

func (a *Args) Parse() error {

	args := os.Args[1:]

	if len(args) <= 0 {
		return errors.New(MissingParameter)

	} else if act, ok := (*a)[args[0]]; ok {
		_, err := act.parse(args[1:])
		return err

	} else if args[0] == "help" {
		return a.usage(args[1:])
	}

	return errors.New(UnknownCommand)
}

func AddArgs(args ...*Arg) Args {
	newArgs := Args{}

	for _, arg := range args {
		newArgs[arg.Name] = arg
	}

	return newArgs
}

func (a *Arg) parse(args []string) (*Arg, error) {

	if len(args) <= 0 && a.Exec == nil {
		return nil, errors.New(MissingParameter)

	} else if a.Exec != nil {
		return nil, a.Exec(args)

	} else if act, ok := a.Args[args[0]]; ok {
		return act.parse(args[1:])

	} else {
		return nil, errors.New(UnknownError)
	}
}

func (a *Args) usage(args []string) error {

	var pointer = *a
	var list = a
	var breadCrumb = filepath.Base(os.Args[0])

	maxArgs := len(args) - 1

	for c := 0; c <= maxArgs; c++ {
		currentCommand := args[c]

		if arg, ok := pointer[currentCommand]; ok {
			breadCrumb += " " + arg.Name
			pointer = arg.Args

			if c == maxArgs && arg.Args == nil {
				list = &Args{
					arg.Name: arg,
				}
			} else if c == maxArgs && arg.Args != nil {
				list = &arg.Args
			}

		} else if !ok {
			return errors.New(UnknownCommand)
		}

	}

	fmt.Printf("Usage: %s\n\n", breadCrumb)

	for _, arg := range *list {
		fmt.Printf(" %s\t\t%s\n", (*arg).Name, (*arg).Description)
	}

	fmt.Printf("\n")

	return nil
}






