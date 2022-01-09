package ast

import (
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex3-hello/token"
)

type Attrib interface{}
type ID string

func NewID(id Attrib) (ID, error) {
	if id == nil {
		return ID(""), nil
	}
	id_lit := string(id.(*token.Token).Lit)
	return ID(id_lit), nil
}
