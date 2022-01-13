// https://github.com/jessevdk/go-flags/issues/294

package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {

	config := struct {
		Config string `long:"config" default:"default"`
	}{}

	parser := flags.NewParser(nil, flags.Default)

	{
		c, _ := parser.AddCommand("hi", "hi", "hi", &struct{}{})
		c.AddGroup("hi.config", "", &config)
	}

	config2 := struct {
		Config string `long:"config" default:"default"`
	}{}
	{
		c, _ := parser.AddCommand("hello", "hello", "hello", &struct{}{})
		c.AddGroup("hello.config", "", &config2)
	}

	parser.Parse()

	ini := flags.NewIniParser(parser)
	ini.Write(os.Stdout, flags.IniIncludeDefaults)
}

/*

$ go run addcommand2.go hi --config hi
[hi.hi.config]
Config = hi

[hello.hello.config]
Config = default


$ go run addcommand2.go hello --config hi
[hi.hi.config]
Config = default

[hello.hello.config]
Config = hi


*/
