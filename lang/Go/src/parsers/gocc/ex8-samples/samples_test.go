package bnf

import (
	"fmt"
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex8-samples/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex8-samples/parser"
)

var testData = []string{
	"a ::= 'b' ;",
	"grm ::= (a | b);",     // OK
	"grm ::= c (a | b);",   // OK
	"grm ::= c d a;",       // OK
	"grm ::= c d (a | b);", // OK
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
