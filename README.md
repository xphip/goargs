# GOARGS
_(WIP)_ Light implementation for command-line flag parsing. 

Example:

> ./main help
>
> ./main help search
>
> ./main search all
```go
package main

import (
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
			Exec:        SearchAll,
		}),
		Exec:        nil,
	})

	if err := args.Parse(); err != nil {
		fmt.Println(err)
	}
}

func SearchAll(args []string) error {

	tmp := `{"id":%s,"username":"%s"}`
	user1 := fmt.Sprintf(tmp, "0", "User1")
	user2 := fmt.Sprintf(tmp, "1", "User2")
	fmt.Printf("[%s,%s]\n", user1, user2)

	return nil
}
```