package calc

import (
	"testing"

	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr1-calc/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/psr1-calc/parser"
)

type TI struct {
	src    string
	expect int64
}

var testData = []*TI{
	{"1 + 1", 2},
	{"1 * 1", 1},
	{"1 + 2 * 3", 7},
	{"(1 + 2) * 3", 9},
	{"2 * (3 + 2)", 10},
	{"1 - 2 + 3", 2},
	{"(1 + 2) / 3", 1},
	{"1 ; 2 +1", 3},
	{"r = 2; 2 + 1", 3},
	{"r = 2; r + 3", 5},
	{"r = 2; r = r + 3; r + 5", 10},
	{"r = 2; r = r + 3; r = r + 5; r = r-3", 7},
	{"r = 2; r = r + 3; r = r + 5; r = r=r-3", 7},
}

func Test1(t *testing.T) {
	p := parser.NewParser()
	for _, ts := range testData {
		s := lexer.NewLexer([]byte(ts.src))
		sum, err := p.Parse(s)
		if err != nil {
			t.Error(err)
		}
		if sum != ts.expect {
			t.Errorf("Error: %s = %d. Expected %d\n", ts.src, sum, ts.expect)
		}
	}
}
