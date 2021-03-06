package main

import (
	"errors"
	"fmt"
	"github.com/xphip/goargs"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	args := goargs.New()

	// $ ./simple_usage version
	args.
		Add("version").
		Usage("Print app version").
		Exec(Version)

	search := args.Add("search").Usage("Search for something")

	// $ ./simple_usage search all
	search.
		Add("all").
		Usage("Return all matches").
		Exec(SearchAll)

	// $ ./simple_usage search id <userID>
	search.
		Add("id").
		Usage("Search by user ID").
		Exec(SearchByID)

	if err := args.Parse(); err != nil {
		fmt.Println(err)
	}
}

func Version(args []string) error {
	filename := filepath.Base(os.Args[0])
	version := "0.0.1-beta"

	fmt.Printf("%s %s %s/%s\n",
		filename,
		version,
		runtime.GOOS,
		runtime.GOARCH)

	return nil
}

func SearchAll(args []string) error {
	tmp := `{"id":%d,"username":"%s"}`

	fmt.Printf("[%s,%s]\n",
		fmt.Sprintf(tmp, 0, "User1"),
		fmt.Sprintf(tmp, 1, "User2"))

	return nil
}

func SearchByID(args []string) error {
	if len(args) == 0 {
		return errors.New("error: missing userID")
	}

	fmt.Printf(`{"id":%s,"username":"User1"}%s`, args[0], "\n")

	return nil
}