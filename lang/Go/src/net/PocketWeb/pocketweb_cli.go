// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package main

import (
	"flag"
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "PocketWeb" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	Directory string // directory holding web files
	Port      string // port used by the pocketweb server
	Debug     int    // debugging level
	Help      bool   // print help then exit
	PrintV    bool   // print version then exit
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {

	// set default values for command line parameters
	flag.StringVar(&Opts.Directory, "d", ".",
		"directory holding web files")
	flag.StringVar(&Opts.Port, "p", "8800",
		"port used by the pocketweb server")
	flag.IntVar(&Opts.Debug, "dbg", 0,
		"debugging level")
	flag.BoolVar(&Opts.Help, "h", false,
		"print help then exit")
	flag.BoolVar(&Opts.PrintV, "ver", false,
		"print version then exit")

	// Now override those default values from environment variables
	if len(Opts.Directory) == 0 ||
		len(os.Getenv("POCKETWEB_D")) != 0 {
		Opts.Directory = os.Getenv("POCKETWEB_D")
	}
	if len(Opts.Port) == 0 ||
		len(os.Getenv("POCKETWEB_P")) != 0 {
		Opts.Port = os.Getenv("POCKETWEB_P")
	}

}

const USAGE_SUMMARY = "  -d\tdirectory holding web files (POCKETWEB_D)\n  -p\tport used by the pocketweb server (POCKETWEB_P)\n  -dbg\tdebugging level (POCKETWEB_DBG)\n  -h\tprint help then exit (POCKETWEB_H)\n  -ver\tprint version then exit (POCKETWEB_VER)\n\nDetails:\n\n"

// Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags]\n\nFlags:\n\n",
		progname)
	fmt.Fprintf(os.Stderr, USAGE_SUMMARY)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\nE.g.:\n\n  PocketWeb &\n  PocketWeb -p 8088 -d /some/where/else\n  POCKETWEB_D=/some/where/else POCKETWEB_P=8088 PocketWeb\n  PocketWeb -ver\n")
	os.Exit(0)
}
