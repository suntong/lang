// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"fmt"
	"strings"

	parseError "github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-gocc/errors"
	"github.com/suntong/lang/lang/Go/src/parsers/gocc/grm2-gocc/token"
)

const (
	numProductions = 40
	numStates      = 119
	numSymbols     = 38
)

// Stack

type stack struct {
	state  []int
	attrib []Attrib
}

const iNITIAL_STACK_SIZE = 100

func newStack() *stack {
	return &stack{
		state:  make([]int, 0, iNITIAL_STACK_SIZE),
		attrib: make([]Attrib, 0, iNITIAL_STACK_SIZE),
	}
}

func (s *stack) reset() {
	s.state = s.state[:0]
	s.attrib = s.attrib[:0]
}

func (s *stack) push(state int, a Attrib) {
	s.state = append(s.state, state)
	s.attrib = append(s.attrib, a)
}

func (s *stack) top() int {
	return s.state[len(s.state)-1]
}

func (s *stack) peek(pos int) int {
	return s.state[pos]
}

func (s *stack) topIndex() int {
	return len(s.state) - 1
}

func (s *stack) popN(items int) []Attrib {
	lo, hi := len(s.state)-items, len(s.state)

	attrib := s.attrib[lo:hi]

	s.state = s.state[:lo]
	s.attrib = s.attrib[:lo]

	return attrib
}

func (s *stack) String() string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "stack:\n")
	for i, st := range s.state {
		fmt.Fprintf(w, "\t%d: %d , ", i, st)
		if s.attrib[i] == nil {
			fmt.Fprintf(w, "nil")
		} else {
			switch attr := s.attrib[i].(type) {
			case *token.Token:
				fmt.Fprintf(w, "%s", attr.Lit)
			default:
				fmt.Fprintf(w, "%v", attr)
			}
		}
		fmt.Fprintf(w, "\n")
	}
	return w.String()
}

// Parser

type Parser struct {
	stack     *stack
	nextToken *token.Token
	pos       int
	Context   Context
}

type Scanner interface {
	Scan() (tok *token.Token)
}

func NewParser() *Parser {
	p := &Parser{stack: newStack()}
	p.Reset()
	return p
}

func (p *Parser) Reset() {
	p.stack.reset()
	p.stack.push(0, nil)
}

func (p *Parser) Error(err error, scanner Scanner) (recovered bool, errorAttrib *parseError.Error) {
	errorAttrib = &parseError.Error{
		Err:            err,
		ErrorToken:     p.nextToken,
		ErrorSymbols:   p.popNonRecoveryStates(),
		ExpectedTokens: make([]string, 0, 8),
	}
	for t, action := range actionTab[p.stack.top()].actions {
		if action != nil {
			errorAttrib.ExpectedTokens = append(errorAttrib.ExpectedTokens, token.TokMap.Id(token.Type(t)))
		}
	}

	if action := actionTab[p.stack.top()].actions[token.TokMap.Type("error")]; action != nil {
		p.stack.push(int(action.(shift)), errorAttrib) // action can only be shift
	} else {
		return
	}

	if action := actionTab[p.stack.top()].actions[p.nextToken.Type]; action != nil {
		recovered = true
	}
	for !recovered && p.nextToken.Type != token.EOF {
		p.nextToken = scanner.Scan()
		if action := actionTab[p.stack.top()].actions[p.nextToken.Type]; action != nil {
			recovered = true
		}
	}

	return
}

func (p *Parser) popNonRecoveryStates() (removedAttribs []parseError.ErrorSymbol) {
	if rs, ok := p.firstRecoveryState(); ok {
		errorSymbols := p.stack.popN(p.stack.topIndex() - rs)
		removedAttribs = make([]parseError.ErrorSymbol, len(errorSymbols))
		for i, e := range errorSymbols {
			removedAttribs[i] = e
		}
	} else {
		removedAttribs = []parseError.ErrorSymbol{}
	}
	return
}

// recoveryState points to the highest state on the stack, which can recover
func (p *Parser) firstRecoveryState() (recoveryState int, canRecover bool) {
	recoveryState, canRecover = p.stack.topIndex(), actionTab[p.stack.top()].canRecover
	for recoveryState > 0 && !canRecover {
		recoveryState--
		canRecover = actionTab[p.stack.peek(recoveryState)].canRecover
	}
	return
}

func (p *Parser) newError(err error) error {
	e := &parseError.Error{
		Err:        err,
		StackTop:   p.stack.top(),
		ErrorToken: p.nextToken,
	}
	actRow := actionTab[p.stack.top()]
	for i, t := range actRow.actions {
		if t != nil {
			e.ExpectedTokens = append(e.ExpectedTokens, token.TokMap.Id(token.Type(i)))
		}
	}
	return e
}

func (p *Parser) Parse(scanner Scanner) (res interface{}, err error) {
	p.Reset()
	p.nextToken = scanner.Scan()
	for acc := false; !acc; {
		action := actionTab[p.stack.top()].actions[p.nextToken.Type]
		if action == nil {
			if recovered, errAttrib := p.Error(nil, scanner); !recovered {
				p.nextToken = errAttrib.ErrorToken
				return nil, p.newError(nil)
			}
			if action = actionTab[p.stack.top()].actions[p.nextToken.Type]; action == nil {
				panic("Error recovery led to invalid action")
			}
		}

		switch act := action.(type) {
		case accept:
			res = p.stack.popN(1)[0]
			acc = true
		case shift:
			p.stack.push(int(act), p.nextToken)
			p.nextToken = scanner.Scan()
		case reduce:
			prod := productionsTable[int(act)]
			attrib, err := prod.ReduceFunc(p.stack.popN(prod.NumSymbols), p.Context)
			if err != nil {
				return nil, p.newError(err)
			} else {
				p.stack.push(gotoTab[p.stack.top()][prod.NTType], attrib)
			}
		default:
			panic("unknown action: " + action.String())
		}
	}
	return res, nil
}
