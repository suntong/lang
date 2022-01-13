package bnf

import (
	"fmt"
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm3-pascal/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm3-pascal/parser"
)

var testData = []string{
	`
	 PROGRAM DEMO0
	 BEGIN
	 END.
`,

	`
	 PROGRAM DEMO1
	 BEGIN
		 A:=3;
		 B:=45;
		 H:=-100023;
		 C:=A;
		 D123:=B34A;
		 BABOON:=GIRAFFE;
		 TEXT:="Hello world!";
	 END.
`,

	`
	PROGRAM DEMO2

	PROCEDURE PrintAnInteger()
	BEGIN
	  A:=3;
	  B:=45;
	END.

	PROCEDURE P2()
	BEGIN
		a := 10;
		b := 'A';
	END.

	BEGIN
		H:=-100023;
		C:=A;
		FOR i := 1 TO 10 DO  // no semicolon here as it would detach the next statement
			D123:=B34A;

		BABOON:=GIRAFFE;
		TEXT:="Hello world!";
	END.
`,
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
