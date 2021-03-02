package main

import (
	"errors"
	"fmt"
	"github.com/xphip/goargs"
)

func main() {
	args := goargs.Args{}

	args.AddArgs(&goargs.Arg{
		Name:        "search",
		Description: "This is the subcommand for search.",
		Args:        goargs.AddArgs(&goargs.Arg{
			Name:        "all",
			Description: "Returns all results.",
			Args:        nil,
			Exec:        func(args []string) error {

				tmp := `{"id":%s,"username":"%s"}`
				user1 := fmt.Sprintf(tmp, "0", "User1")
				user2 := fmt.Sprintf(tmp, "1", "User2")
				fmt.Printf("[%s,%s]\n", user1, user2)

				return nil
			},
		}, &goargs.Arg{
			Name:        "id",
			Description: "Returns a search by ID.",
			Args:        nil,
			Exec: SearchByID,
		}),
		Exec:        nil,
	})

	if err := args.Parse(); err != nil {
		fmt.Println(err)
	}
}

func SearchByID(args []string) error {
	if len(args) < 1 {
		return errors.New(goargs.MissingParameter)
	}

	id := args[0]
	fmt.Printf(`{"id":%s,"username":"User1"}%s`, id, "\n")

	return nil
}