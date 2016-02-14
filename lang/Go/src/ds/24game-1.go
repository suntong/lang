////////////////////////////////////////////////////////////////////////////
// Porgram: 24game-1
// Purpose: Solve 24 game in Go, single solution
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: http://rosettacode.org/wiki/24_game/Solve#Go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	op_num = iota
	op_add
	op_sub
	op_mul
	op_div
)

type frac struct {
	num, denom int
}

// Expression: can either be a single number, or a result of binary
// operation from left and right node
type Expr struct {
	op          int
	left, right *Expr
	value       frac
}

var n_cards = 4
var goal = 24
var digit_range = 9

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (x *Expr) String() string {
	if x.op == op_num {
		return fmt.Sprintf("%d", x.value.num)
	}

	var bl1, br1, bl2, br2, opstr string
	switch {
	case x.left.op == op_num:
	case x.left.op >= x.op:
	case x.left.op == op_add && x.op == op_sub:
		bl1, br1 = "", ""
	default:
		bl1, br1 = "(", ")"
	}

	if x.right.op == op_num || x.op < x.right.op {
		bl2, br2 = "", ""
	} else {
		bl2, br2 = "(", ")"
	}

	switch {
	case x.op == op_add:
		opstr = " + "
	case x.op == op_sub:
		opstr = " - "
	case x.op == op_mul:
		opstr = " * "
	case x.op == op_div:
		opstr = " / "
	}

	return bl1 + x.left.String() + br1 + opstr +
		bl2 + x.right.String() + br2
}

func expr_eval(x *Expr) (f frac) {
	if x.op == op_num {
		return x.value
	}

	l, r := expr_eval(x.left), expr_eval(x.right)

	switch x.op {
	case op_add:
		f.num = l.num*r.denom + l.denom*r.num
		f.denom = l.denom * r.denom
		return

	case op_sub:
		f.num = l.num*r.denom - l.denom*r.num
		f.denom = l.denom * r.denom
		return

	case op_mul:
		f.num = l.num * r.num
		f.denom = l.denom * r.denom
		return

	case op_div:
		f.num = l.num * r.denom
		f.denom = l.denom * r.num
		return
	}
	return
}

func solve(ex_in []*Expr) bool {
	// only one expression left, meaning all numbers are arranged into
	// a binary tree, so evaluate and see if we get 24
	if len(ex_in) == 1 {
		f := expr_eval(ex_in[0])
		if f.denom != 0 && f.num == f.denom*goal {
			fmt.Println(ex_in[0].String())
			return true
		}
		return false
	}

	var node Expr
	ex := make([]*Expr, len(ex_in)-1)

	// try to combine a pair of expressions into one, thus reduce
	// the list length by 1, and recurse down
	for i := range ex {
		copy(ex[i:len(ex)], ex_in[i+1:len(ex_in)])

		ex[i] = &node
		for j := i + 1; j < len(ex_in); j++ {
			node.left = ex_in[i]
			node.right = ex_in[j]

			// try all 4 operators
			for o := op_add; o <= op_div; o++ {
				node.op = o
				if solve(ex) {
					return true
				}
			}

			// also - and / are not commutative, so swap arguments
			node.left = ex_in[j]
			node.right = ex_in[i]

			node.op = op_sub
			if solve(ex) {
				return true
			}

			node.op = op_div
			if solve(ex) {
				return true
			}

			if j < len(ex) {
				ex[j] = ex_in[j]
			}
		}
		ex[i] = ex_in[i]
	}
	return false
}

func main() {
	cards := make([]*Expr, n_cards)

	// If there are command line arguments
	if len(os.Args) > 1 {
		if len(os.Args) != n_cards+1 {
			fmt.Printf("Error: Needs exactly %d numbers from command line", n_cards)
			os.Exit(1)
		}

		for i := 0; i < n_cards; i++ {
			n, err := strconv.Atoi(os.Args[i+1])
			check(err)
			cards[i] = &Expr{op_num, nil, nil,
				frac{n, 1}}
			fmt.Printf(" %d", cards[i].value.num)
		}
		fmt.Print(":  ")
		if !solve(cards) {
			fmt.Println("No solution")
		}
		os.Exit(0)
	}

	rand.Seed(time.Now().Unix())

	for k := 0; k < 10; k++ {
		for i := 0; i < n_cards; i++ {
			cards[i] = &Expr{op_num, nil, nil,
				frac{rand.Intn(digit_range-1) + 1, 1}}
			fmt.Printf(" %d", cards[i].value.num)
		}
		fmt.Print(":  ")
		if !solve(cards) {
			fmt.Println("No solution")
		}
	}
}

/*

To run, either with no command line arguments:

    $ go run 24game-1.go
     3 5 3 6:  (3 + 5) * (6 - 3)
     5 7 2 2:  5 * 2 + 7 * 2
     2 4 8 7:  7 / (2 / 8) - 4
     2 2 3 4:  (2 + 2 + 4) * 3
     1 7 3 5:  (3 - 1) * (7 + 5)
     1 6 3 3:  (1 + 6) * 3 + 3
     2 1 4 1:  No solution
     5 3 8 2:  (5 + 3) * 2 + 8
     8 2 8 4:  (8 - 2) * (8 - 4)
     7 5 2 6:  7 + 5 + 2 * 6

Or, with exactly 4 numbers from command line

    $ go run 24game-1.go 1 3 4 6
     1 3 4 6:  6 / (1 - 3 / 4)

*/
