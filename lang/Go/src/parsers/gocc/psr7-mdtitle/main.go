package main

import (
	"fmt"
	"io/ioutil"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/parser"
	//"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr7-mdtitle/token"
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

/*

$ psr7-mdtitle
ast value: "# This is title\n\n## This is another title\n\n\n\n# Heading level 1\n## Heading level 2\n\n### Heading level 3\n\n\n#### Heading level 4\n\n\n\n##### Heading level 5\n\n\n###### Heading level 6\n\n"

*/
