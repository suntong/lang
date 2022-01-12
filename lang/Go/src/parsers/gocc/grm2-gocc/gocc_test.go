package gocc

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-gocc/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-gocc/parser"
)

var testData = []string{
	"a : 'b' ;",
	"id : (_letter | '_') {_idchar} ;",
	"_letter : 'A'-'Z' | 'a'-'z' | '_' ;",
	`_alpha : _letter | _digit ;`,
}

func TestPass(t *testing.T) {
	for _, ts := range testData {
		_, err := test([]byte(ts))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFail(t *testing.T) {
	_, err := test([]byte("a ::= b"))
	if err == nil {
		t.Fatal("unexpected parse error")
	} else {
		fmt.Printf("  Parsing failed as expected: %v\n", err)
	}
}

func TestFiles(t *testing.T) {
	files := []string{}
	err := filepath.Walk("test",
		func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				// fmt.Printf("  Testing file: %s (%d bytes)\n", path, f.Size())
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
	// fmt.Printf("  Testing files: %v\n", files)
	if len(files) == 0 {
		return
	}

	p := parser.NewParser()
	for _, tf := range files {
		fmt.Printf("  Testing file: %s\n", tf)
		ts, err := ioutil.ReadFile(tf)
		if err != nil {
			panic(err)
		}
		s := lexer.NewLexer(ts)
		_, err = p.Parse(s)
		if err != nil {
			t.Error(err)
		}
	}
}

func test(src []byte) (astree interface{}, err error) {
	fmt.Printf("  Testing: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a
	}
	return
}

/*

$ go test -v .
=== RUN   TestPass
  Testing: a : 'b' ;
  Testing: id : (_letter | '_') {_idchar} ;
  Testing: _letter : 'A'-'Z' | 'a'-'z' | '_' ;
  Testing: _alpha : _letter | _digit ;
--- PASS: TestPass (0.00s)
=== RUN   TestFail
  Testing: a ::= b
  Parsing failed as expected: 1:4: error: expected one of regDefId, ".", char_lit, "[", "{", or "("; got: ":"
--- PASS: TestFail (0.00s)
=== RUN   TestFiles
  Testing file: test/ast.bnf
  Testing file: test/bnf.bnf
  Testing file: test/bools.bnf
  Testing file: test/calc.bnf
  Testing file: test/ebnf.bnf
  Testing file: test/fs.bnf
  Testing file: test/hello.bnf
  Testing file: test/sparql.bnf
  Testing file: test/usercontext.bnf
--- PASS: TestFiles (0.00s)
PASS
ok      github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-gocc      (cached)

*/
