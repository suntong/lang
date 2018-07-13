package main

import (
	"regexp"

	"github.com/go-easygen/cli"
)

type MapStringString struct {
	Keys   []string
	Values map[string]string
}

type argT struct {
	Macros MapStringString `cli:"D" usage:"define macros"`
}

// DecodeSlice implements cli.SliceDecoder
// NOTE: if SliceDecoder not implemented, the Decode method would be only invoked once
func (MapStringString) DecodeSlice() {}

// Decode implements cli.Decoder interface
func (m *MapStringString) Decode(s string) error {
	if (m.Values) == nil {
		m.Values = make(map[string]string)
	}
	matches := regexp.MustCompile("(.*)=(.*)").FindStringSubmatch(s)
	key := matches[1]
	val := matches[2]
	m.Keys = append(m.Keys, key)
	m.Values[key] = val
	return nil
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.JSONln(ctx.Argv())
		return nil
	})
}

/*

$ go run 106-3map.go -Dx=1 -D y=2
{"Macros":{"Keys":["x","y"],"Values":{"x":"1","y":"2"}}}

$ go run 106-3map.go -Dx='1 1' -D y="20 01"
{"Macros":{"Keys":["x","y"],"Values":{"x":"1 1","y":"20 01"}}}

*/
