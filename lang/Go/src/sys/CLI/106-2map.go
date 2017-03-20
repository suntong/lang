package main

import (
	"strconv"
	"strings"

	"github.com/mkideal/cli"
)

type mapT map[string]int64

type argT struct {
	Macros mapT `cli:"D" usage:"define macros"`
}

var keys = []string{}

// Decode implements cli.Decoder interface
func (maps *mapT) Decode(s string) error {
	*maps = make(map[string]int64)
	keys = []string{}
	i := strings.Index(s, "=")
	if i >= 0 {
		key := s[:i]
		val, err := strconv.ParseInt(s[i+1:], 10, 64)
		if err != nil {
			return err
		}
		keys = append(keys, key)
		(*maps)[key] = val
	}
	return nil
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.JSONln(ctx.Argv())
		return nil
	})
}
