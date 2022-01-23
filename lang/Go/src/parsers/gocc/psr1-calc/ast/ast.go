package ast

import (
// "github.com/suntong/lang/lang/Go/src/parsers/gocc/psr1-calc/token"
// "github.com/suntong/lang/lang/Go/src/parsers/gocc/psr1-calc/util"
)

var Reg int64

func Assign(v int64) (int64, error) {
	//println("Assign!")
	Reg = v
	return Reg, nil
}
