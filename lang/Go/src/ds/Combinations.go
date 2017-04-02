////////////////////////////////////////////////////////////////////////////
// Program: Combinations.go
// Purpose: generate all combinations
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: https://rosettacode.org/wiki/Combinations#Go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"strconv"
)

/*

$ go run Combinations.go 5 3
A B C
A B D
A B E
A C D
A C E
A D E
B C D
B C E
B D E
C D E

*/

var ca []rune

func init() {
	ca = make([]rune, 26)
	for i, c := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		ca[i] = c
	}
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("Usage\n  Combinations n k\nto print combination of n pick k")
		os.Exit(0)
	}

	n, err := strconv.Atoi(os.Args[1])
	check(err)
	k, err := strconv.Atoi(os.Args[2])
	check(err)
	comb(n, k, func(c []int) {
		for _, v := range c {
			fmt.Printf("%c ", ca[v])
		}
		fmt.Println()
	})
}

func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
