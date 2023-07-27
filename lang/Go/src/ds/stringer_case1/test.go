//go:generate stringer -type=solfa

package main

import (
	"bytes"
)

type solfa int

const (
	_ solfa = iota
	Do
	Re
	Mi
	Fa
	So
	La
	Ti
)

func main() {
	buf := new(bytes.Buffer)
	_ = buf
}
