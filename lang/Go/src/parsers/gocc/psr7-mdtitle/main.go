package main

import (
	"fmt"
	"io/ioutil"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/parser"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/token"
)

func main() {
	dat, err := ioutil.ReadFile("./example.md")
	check(err)

	l := lexer.NewLexer([]byte(dat))
	p := parser.NewParser()
	ast, err := p.Parse(l)
	check(err)

	fmt.Printf("ast value: %#+v\n", ast)
	fmt.Printf("ast Lit value: %s\n", AttribToString(ast))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Attrib interface{}

func AttribToString(a Attrib) string {
	return string(a.(*token.Token).Lit)
}
