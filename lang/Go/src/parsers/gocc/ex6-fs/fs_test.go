package fs

//go:generate gocc -a -p github.com/suntong/lang/lang/Go/src/parsers/gocc/ex6-fs fs.bnf

import (
	"fmt"
	"testing"

	//"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex6-fs/ast"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex6-fs/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex6-fs/parser"
)

func TestPass(t *testing.T) {
	sml, err := test([]byte("0 + 1 < 23 "))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("output: %s\n", sml)
}

func TestFail(t *testing.T) {
	_, err := test([]byte("a b"))
	if err == nil {
		t.Fatal("expected parse error")
	} else {
		fmt.Printf("Parsing failed as expected: %v\n", err)
	}
}

func test(src []byte) (astree interface{}, err error) {
	fmt.Printf("input: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a
	}
	return
}

type TI struct {
	src    string
	expect int64
}

var testData = []*TI{
	{"1 + 1", 2},
	{"1 * 1", 1},
	{"1 + 2 * 3", 7},
	{"0", 0},
	{"1", 0},
	{"23", 0},
	{"345", 0},
	{"1 || 22", 0},
	{"23 && 567", 0},
	{"1 == 2", 0},
	{"1 != 2", 0},
	{"fn1(1 <= 2)", 0},
	{"fn2.fn3(1 <= 2)", 0},
}

func TestExp(t *testing.T) {
	p := parser.NewParser()
	for _, ts := range testData {
		s := lexer.NewLexer([]byte(ts.src))
		sum, err := p.Parse(s)
		if err != nil {
			t.Error(err)
		}
		if sum != ts.expect {
			// t.Errorf("Error: %s = %d. Expected %d\n", ts.src, sum, ts.expect)
		}
	}
}

/*

$ go test -v .
=== RUN   TestPass
input: 0 + 1 < 23
output: Pos(offset=0, line=1, column=1)
--- PASS: TestPass (0.00s)
=== RUN   TestFail
input: a b
Parsing failed as expected: 1:3: error: expected "("; got: "b"
--- PASS: TestFail (0.00s)
=== RUN   TestExp
--- PASS: TestExp (0.00s)
PASS
ok      github.com/suntong/lang/lang/Go/src/parsers/gocc/ex6-fs (cached)

*/
