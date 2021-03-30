package main

import (
	"fmt"
	"github.com/xphip/goargs"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	cmd := goargs.New()

	cmd.Add("version").Usage("Print app version").Exec(Version)

	cmd.
		Add("test").
		Usage("Print app version").
		Map([]string{"first", "second"}).
		Exec(func(args *goargs.Args) error {
			fmt.Printf("First(string): %s\nSecond(number or -1): %d\n\n",
				args.Get("first").String(),
				args.Get("second").Int(-1))
			fmt.Println(args.GetMapped())
			fmt.Println(args.GetUnmapped())

			return nil
	})

	if err := cmd.Parse(); err != nil {
		fmt.Println(err)
		return
	}
}

func Version(_ *goargs.Args) error {
	filename := filepath.Base(os.Args[0])
	version := "0.0.1-beta"

	fmt.Printf("%s %s %s/%s\n",
		filename,
		version,
		runtime.GOOS,
		runtime.GOARCH)

	return nil
}
