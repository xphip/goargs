# GOARGS
_(WIP)_ Light implementation for command-line flag parsing. 

Example:
> ./simple_usage version 
```go
package main

import (
	"fmt"
	"github.com/xphip/goargs"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	args := goargs.New()

	args.Add("version").
	     Usage("Print app version").
	     Exec(Version)

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
```