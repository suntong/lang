package bnf

import (
	"fmt"
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm5-line/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm5-line/parser"
)

var testData = []string{
	"`[21] These to ignore, `",
	//	"/* [21] These are strings */2022-01-18 11:33:21.9885 [21] These are strings that I need to ignore, until - MYKW - Start Active One: 1/18/2022 11:33:21 AM\n",
	//	"2022-01-18 2022-01-18 2022-01-18 11:33:21.9885 [21] These are strings that I need to egnore, until - MYKW - Start Active One: 1/18/2022 11:33:21 AM\n",
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
