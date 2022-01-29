package main

import (
	"fmt"
	"io/ioutil"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/parser"
)

func main() {
	dat, err := ioutil.ReadFile("./example.md")
	check(err)

	l := lexer.NewLexer([]byte(dat))
	p := parser.NewParser()
	ast, err := p.Parse(l)
	check(err)

	fmt.Printf("ast value: %#+v\n", ast)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
