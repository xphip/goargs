# GOARGS
_(WIP)_ Light implementation for command-line flag parsing. 

```
# Example:
< ./simple_usage version
> simple_usage 0.0.1-beta linux/amd64

< ./simple_usage test a b c
> First(string): a
  Second(number or -1): -1
  map[first:a second:b]
  [c]
``` 
```go
func main() {
    cmd := goargs.New()

    cmd.Add("version").Usage("Print app version").Exec(Version)

    cmd.Add("test").
        Usage("Just a test").
        Map([]string{"first", "second"}).
        Exec(func (args *goargs.Args) error {

            fmt.Printf("First(string): %s\n",
                args.Get("first").String())
            fmt.Printf("Second(number or -1): %d\n",
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

    fmt.Printf("%s %s %s/%s\n",
        filepath.Base(os.Args[0]),
        "0.0.1-beta",
        runtime.GOOS,
        runtime.GOARCH)

    return nil
}
```
[Full example](https://github.com/xphip/goargs/blob/main/examples/simple_usage/simple_usage.go)