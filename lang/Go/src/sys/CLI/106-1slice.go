////////////////////////////////////////////////////////////////////////////
// Program: 106-1slice.go
// Purpose: 006-map extension demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: as credited below
////////////////////////////////////////////////////////////////////////////

/*

How to obtain/preserver output orders to that given on command line?
https://github.com/mkideal/cli/issues/29

Q: The go map is known to not preserver map orders. I.e., if we run 006-map.go
a couple of times, the output might not always be the same.
How to obtain/preserver output orders to that given on command line?

go-easygen: You can replace map with slice. Here is an example

*/

package main

import "github.com/mkideal/cli"
import "strings"
import "strconv"

type KV struct {
	Key   string
	Value int
}

type KVs []KV

// Decode implements cli.Decoder interface
func (kvs *KVs) Decode(s string) error {
	i := strings.Index(s, "=")
	if i >= 0 {
		key := s[:i]
		val, err := strconv.ParseInt(s[i+1:], 10, 64)
		if err != nil {
			return err
		}
		*kvs = append(*kvs, KV{key, int(val)})
	}
	return nil
}

type argT struct {
	cli.Helper
	Values KVs `cli:"D" usage:"define macro values"`
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		ctx.JSONIndentln(ctx.Argv(), "", "    ")
		return nil
	})
}

/*

$ go run 106-slice.go -Dx=1 -D y=2
{
    "Help": false,
    "Values": [
        {
            "Key": "x",
            "Value": 1
        },
        {
            "Key": "y",
            "Value": 2
        }
    ]
}

*/
