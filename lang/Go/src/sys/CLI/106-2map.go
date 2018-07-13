package main

import (
	"strconv"
	"strings"

	"github.com/go-easygen/cli"
)

type MapStringToInt64 struct {
	Keys   []string
	Values map[string]int64
}

type argT struct {
	Macros MapStringToInt64 `cli:"D" usage:"define macros"`
}

// DecodeSlice implements cli.SliceDecoder
// NOTE: if SliceDecoder not implemented, the Decode method would be only invoked once
func (MapStringToInt64) DecodeSlice() {}

// Decode implements cli.Decoder interface
func (m *MapStringToInt64) Decode(s string) error {
	if (m.Values) == nil {
		m.Values = make(map[string]int64)
	}
	i := strings.Index(s, "=")
	if i >= 0 {
		key := s[:i]
		val, err := strconv.ParseInt(s[i+1:], 10, 64)
		if err != nil {
			return err
		}
		m.Keys = append(m.Keys, key)
		m.Values[key] = val
	}
	return nil
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.JSONln(ctx.Argv())
		return nil
	})
}
