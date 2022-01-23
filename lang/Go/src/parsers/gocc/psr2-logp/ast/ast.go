package ast

import (
	//"strings"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr2-logp/token"
)

type Attrib interface{}

func AttribToString(a Attrib) string {
	return string(a.(*token.Token).Lit)
}

func Pair(s, e string) string {
	r := s + " ==> " + e
	println(r)
	return r
}
