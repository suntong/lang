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

	// {
	//         c, _ := parser.AddCommand("hello", "hello", "hello", &struct{}{})
	//         c.AddGroup("hello.config", "", &config)
	// }

	parser.Parse()

	ini := flags.NewIniParser(parser)
	ini.Write(os.Stdout, flags.IniIncludeDefaults)
}
