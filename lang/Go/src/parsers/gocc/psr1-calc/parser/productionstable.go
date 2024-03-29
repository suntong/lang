// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr1-calc/ast"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr1-calc/token"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr1-calc/util"
)

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Calcs	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Calcs : Calc	<<  >>`,
		Id:         "Calcs",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Calcs : Calcs ";" Calc	<< X[2], nil >>`,
		Id:         "Calcs",
		NTType:     1,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[2], nil
		},
	},
	ProdTabEntry{
		String: `Calc : Expr	<<  >>`,
		Id:         "Calc",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Calc : reg "=" Calc	<< ast.Assign(X[2].(int64)) >>`,
		Id:         "Calc",
		NTType:     2,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Assign(X[2].(int64))
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "+" Term	<< X[0].(int64) + X[2].(int64), nil >>`,
		Id:         "Expr",
		NTType:     3,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(int64) + X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Expr : Expr "-" Term	<< X[0].(int64) - X[2].(int64), nil >>`,
		Id:         "Expr",
		NTType:     3,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(int64) - X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Expr : Term	<<  >>`,
		Id:         "Expr",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Term : Term "*" Factor	<< X[0].(int64) * X[2].(int64), nil >>`,
		Id:         "Term",
		NTType:     4,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(int64) * X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Term : Term "/" Factor	<< X[0].(int64) / X[2].(int64), nil >>`,
		Id:         "Term",
		NTType:     4,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0].(int64) / X[2].(int64), nil
		},
	},
	ProdTabEntry{
		String: `Term : Factor	<<  >>`,
		Id:         "Term",
		NTType:     4,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : "(" Expr ")"	<< X[1], nil >>`,
		Id:         "Factor",
		NTType:     5,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Factor : int64	<< util.IntValue(X[0].(*token.Token).Lit) >>`,
		Id:         "Factor",
		NTType:     5,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return util.IntValue(X[0].(*token.Token).Lit)
		},
	},
	ProdTabEntry{
		String: `Factor : reg	<< ast.Reg, nil >>`,
		Id:         "Factor",
		NTType:     5,
		Index:      13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.Reg, nil
		},
	},
}
