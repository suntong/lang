////////////////////////////////////////////////////////////////////////////
// Porgram: DataStructure-Snail
// Purpose: Go data structure demo, the Snail problem
// Authors: Tong Sun (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// Snail-moving Directions
const (
	right = iota
	down
	left
	up
)

var (
	dirs    = []int{right, down, left, up}
	optFlip = flag.Bool("f", false, "Flip the order")
)

////////////////////////////////////////////////////////////
// Function definitions

// Make sure to work in the go playground as well
func init() {
	os.Args = append(os.Args, "-f")
	os.Args = append(os.Args, "5")
}

// Function main
func main() {
	flag.Parse()
	// Get the matrix size to n
	n, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := snailFill(n, n)
	show(n, values, optFlip)
}

// Create & fill the matrix with values
func snailFill(m, n int) (matrix [][]int) {
	// Create the matrix
	res := make([][]int, m)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	// m, n = len(matrix), len(matrix[0])

	f, l, r, t, b := 1, 0, n-1, 0, m-1
	for cnt := 0; l <= r && t <= b; cnt++ {
		switch dirs[cnt%len(dirs)] {
		case right:
			for i, j := t, l; j <= r; j++ {
				res[i][j] = f
				f++
			}
			t++
		case down:
			for i, j := t, r; i <= b; i++ {
				res[i][j] = f
				f++
			}
			r--
		case left:
			for i, j := b, r; j >= l; j-- {
				res[i][j] = f
				f++
			}
			b--
		case up:
			for i, j := b, l; i >= t; i-- {
				res[i][j] = f
				f++
			}
			l++
		}
	}

	return res
}

// Print out the snail
func show(n int, values [][]int, optFlip *bool) {
	flipping := *optFlip
	if flipping {
		fmt.Println("Flipping output order")
	}
	format := fmt.Sprintf("%%%dd ", int(math.Log(float64(n*n))/math.Ln10)+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if flipping {
				fmt.Printf(format, n*n-values[j][i]+1)
			} else {
				fmt.Printf(format, values[i][j])
			}
		}
		fmt.Println("")
	}

}
