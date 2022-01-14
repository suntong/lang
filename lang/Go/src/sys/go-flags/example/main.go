// from: https://github.com/jessevdk/go-flags/blob/master/examples/main.go

package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type EditorOptions struct {
	Input  flags.Filename `short:"i" long:"input" description:"Input file" default:"-"`
	Output flags.Filename `short:"o" long:"output" description:"Output file" default:"-"`
}

type Options struct {
	// Example of verbosity with level
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`

	// Example of optional value
	User string `short:"u" long:"user" description:"User name" optional:"yes" optional-value:"pancake"`

	// Example of map with multiple default values
	Users map[string]string `long:"users" description:"User e-mail map" default:"system:system@example.org" default:"admin:admin@example.org"`

	// Example of option group
	Editor EditorOptions `group:"Editor Options"`

	// Example of using environment variables
	EnvVar1 string `long:"env-var1" default:"Some value" env:"ENV_DEFAULT" description:"Test env-var1 value"`
	Thresholds  []int     `long:"thresholds" default:"1" default:"2" env:"THRESHOLD_VALUES"  env-delim:"," description:"Environment value array"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
