package ast

import (
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex3-hello/token"
)

type (
	ID string
)

func NewHello(id interface{}) (Hello, error) {
	return ID{id.(ID)}, nil
}
