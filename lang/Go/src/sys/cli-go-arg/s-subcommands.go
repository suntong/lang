package main

import (
	//"fmt"
	"os"

	arg "github.com/alexflint/go-arg"
)

type CheckoutCmd struct {
	Branch string `arg:"positional"`
	Track  bool   `arg:"-t"`
}

type CommitCmd struct {
	All     bool   `arg:"-a"`
	Message string `arg:"-m"`
}

type PushCmd struct {
	Remote      string `arg:"positional"`
	Branch      string `arg:"positional"`
	SetUpstream bool   `arg:"-u"`
}

type argsT struct {
	Checkout *CheckoutCmd `arg:"subcommand:checkout"`
	Commit   *CommitCmd   `arg:"subcommand:commit"`
	Push     *PushCmd     `arg:"subcommand:push"`
	Quiet    bool         `arg:"-q"` // this flag is global to all subcommands
}

func (argsT) Description() string {
	return "This program does this and that"
}

var args argsT

func main() {
	p := arg.MustParse(&args)
	if p.Subcommand() == nil {
		p.WriteUsage(os.Stdout)
		os.Exit(1)
	}
}
