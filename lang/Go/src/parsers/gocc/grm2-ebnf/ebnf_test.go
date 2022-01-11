package bnf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-ebnf/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-ebnf/parser"
)

var testData = []string{
	"a : b ;",
	//	"id : (_letter | '_') {_idchar} ;",
	"_letter : 'A'-'Z' | 'a'-'z' | '_' ;",
	`_alpha : _letter | _digit ;`,
	//	"!lineComment  : '/' '/' { . } '\n' ;",
	"!whitespace : ' ' | '\t' | '\v' | '\f' | '\r' | '\n';",
	"_whitespace : ' ' | '\t' | '\v' | '\f' | '\r' | '\n';",
	"_whitespace : ' ' | '\t' | '\v' | '\f' | '\r' | '\n';\n",
}

/*

 */

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

*/
