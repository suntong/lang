package bnf

import (
	"fmt"
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-gocc1/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-gocc1/parser"
)

var testData = []string{
	"a : 'b' ;",
	"grm : (_a | _b);", // OK
	//"grm : (a | b);",   // NOK!
}

func TestPass(t *testing.T) {
	p := parser.NewParser()
	for _, ts := range testData {
		fmt.Printf("  Testing: %s\n", ts)
		s := lexer.NewLexer([]byte(ts))
		_, err := p.Parse(s)
		if err != nil {
			t.Error(err)
		}
	}
}
