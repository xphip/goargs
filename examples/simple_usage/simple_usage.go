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

	a := cmd.Add("a").Usage("Letter A.")
	b := a.Add("b").Usage("Letter B.")

	b.Add("c").Usage("Letter C.").
		Exec(func (_ *goargs.Args) error {
			fmt.Println("Alphabet!")
			return nil
		})

	cmd.Add("test").Usage("Just a test").
		Map([]string{"first", "second"}).
		Exec(Test)

	if err := cmd.Parse(); err != nil {
		fmt.Println(err)
		return
	}
}

func Version(_ *goargs.Args) error {

	fmt.Printf("%s %s %s/%s\n",
		filepath.Base(os.Args[0]),
		"0.0.1-beta",
		runtime.GOOS,
		runtime.GOARCH)

	return nil
}

func Test(args *goargs.Args) error {

	fmt.Printf("First(string): %s\n",
		args.Get("first").String())

	fmt.Printf("Second(number or -1): %d\n",
		args.Get("second").Int(-1))

	fmt.Println(args.GetMapped())
	fmt.Println(args.GetUnmapped())

	return nil
}