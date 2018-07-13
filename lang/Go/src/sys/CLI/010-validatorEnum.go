package main

import (
	"fmt"

	"github.com/go-easygen/cli"
	"github.com/suntong/enum"
)

var (
	gender enum.Enum
	male   = gender.Iota("male")
	female = gender.Iota("female")

	theGender int = -1
)

type argT struct {
	cli.Helper
	Age    int    `cli:"age" usage:"your age"`
	Gender string `cli:"g,gender" usage:"your gender (male/female)" dft:"male"`
}

// Validate implements cli.Validator interface
func (argv *argT) Validate(ctx *cli.Context) error {
	if argv.Age < 0 || argv.Age > 300 {
		return fmt.Errorf("age %d out of range", argv.Age)
	}
	ok := false
	if theGender, ok = gender.Get(argv.Gender); !ok {
		return fmt.Errorf("invalid gender %s", ctx.Color().Yellow(argv.Gender))
	}
	return nil
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.JSONln(ctx.Argv())
		fmt.Printf("%d:%s\n", theGender, gender.String(theGender))
		return nil
	})
}
