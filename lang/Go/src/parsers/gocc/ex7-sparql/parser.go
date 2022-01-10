//go:generate gocc -p github.com/suntong/lang/lang/Go/src/parsers/gocc/ex7-sparql -a sparql.bnf
package sparql

import (
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex7-sparql/ast"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex7-sparql/lexer"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/ex7-sparql/parser"
	"sync"
)

var p *parser.Parser
var l sync.Mutex

func init() {
	p = parser.NewParser()
}

func Parse(s string) (*ast.Query, error) {
	l.Lock()
	defer l.Unlock()
	lexed := lexer.NewLexer([]byte(s))
	_q, err := p.Parse(lexed)
	if err != nil {
		return nil, err
	}
	q := _q.(ast.Query)
	return &q, nil
}
