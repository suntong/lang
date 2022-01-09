package astx

import (
	"fmt"
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex3-hello/ast"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex3-hello/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex3-hello/parser"
)

func TestPass(t *testing.T) {
	sml, err := test([]byte("hello abc"))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("output: %s\n", sml)
}

func TestFail(t *testing.T) {
	_, err := test([]byte("hello ab;"))
	if err == nil {
		t.Fatal("expected parse error")
	} else {
		fmt.Printf("Parsing failed as expected: %v\n", err)
	}
}

func test(src []byte) (astree ast.Hello, err error) {
	fmt.Printf("input: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.Hello)
	}
	return
}
