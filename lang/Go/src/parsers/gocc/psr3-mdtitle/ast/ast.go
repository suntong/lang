package ast

import (
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr3-mdtitle/token"
)

type Attrib interface{}

func AttribToString(a Attrib) string {
	return string(a.(*token.Token).Lit)
}
