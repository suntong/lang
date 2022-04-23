////////////////////////////////////////////////////////////////////////////
// Porgram: DataStructure-Snail
// Purpose: Go data structure demo, the Snail problem
// Authors: Tong Sun (c) 2014, All rights reserved
////////////////////////////////////////////////////////////////////////////

// Style: gofmt -tabs=false -tabwidth=2 -w

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// Snail-moving Directions
const (
	dRight = iota
	dDown
	dLeft
	dUp
	dNumbers // total number of directions
)

// Snail-moving Boundary
type tBoundary struct {
	xmin, xmax, ymin, ymax int
}

// Snail-moving Movement instruction
type tMove struct {
	// Direction
	d int
	// Position adjustment
	x, y int
}

// Option for output format
type tOptFmt int

// A list of valid formats
const (
	fNormal tOptFmt = iota
	fFlip           // Flip the snail direction and flip the count from up to down
	fMirror         // Mirror the starting porint and count direction
)

////////////////////////////////////////////////////////////
// Global variables definitions

// snail movement & boundary
var sb tBoundary
var sm = []tMove{
	{dRight, 1, 0},
	{dDown, 0, 1},
	{dLeft, -1, 0},
	{dUp, 0, -1},
}

////////////////////////////////////////////////////////////
// Function definitions

// Make sure to work in the go playground as well
func init() {
	os.Args = append(os.Args, "5")
}

// Function main
func main() {

	// Get the matrix size to n
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var optFmt tOptFmt = fNormal
	if os.Args[2] == "-f" {
		optFmt = fFlip
	} else if os.Args[2] == "-m" {
		optFmt = fMirror
	}
	fmt.Printf("r: %v\n", optFmt)

	//fmt.Printf("sm: %v\n", sm)
	sb = tBoundary{0, n - 1, 0, n - 1}
	//fmt.Printf("sb: %v\n", sb)

	// Create the matrix
	values := make([][]int, n)
	for i := 0; i < n; i++ {
		values[i] = make([]int, n)
	}

	draw(n, values)
	show(n, values, optFmt)

}

// Fill the matrix with values
func draw(n int, values [][]int) {
	// Direction & position x & y
	d, x, y := 0, 0, 0

	for value := 1; value <= n*n; value++ {
		//fmt.Println(value, x, y)
		values[y][x] = value
		x += sm[d].x
		y += sm[d].y
		if x < sb.xmin || x > sb.xmax || y < sb.ymin || y > sb.ymax {
			x -= sm[d].x
			y -= sm[d].y
			d, x, y = turn(d, x, y)
		}
	}
}

// Turn the snail-moving direction each time when hitting the boundary
func turn(d, x, y int) (od, ox, oy int) {
	ox, oy = x, y
	od = (d + 1) % dNumbers

	switch d {
	case dRight:
		sb.ymin += 1
		oy += sm[od].y
	case dDown:
		sb.xmax -= 1
		ox += sm[od].x
	case dLeft:
		sb.ymax -= 1
		oy += sm[od].y
	case dUp:
		sb.xmin += 1
		ox += sm[od].x
	}

	//fmt.Printf("Turning: %d => %d, (%d,%d) => (%d,%d)\n", d, od, x, y, ox, oy)
	return
}

// Print out the snail
func show(n int, values [][]int, optFmt tOptFmt) {
	format := fmt.Sprintf("%%%dd ", int(math.Log(float64(n*n))/math.Ln10)+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if optFmt == fFlip {
				fmt.Printf(format, n*n-values[j][i]+1)
			} else if optFmt == fMirror {
				fmt.Printf(format, n*n-values[n-j-1][n-i-1]+1)
			} else {
				fmt.Printf(format, values[i][j])
			}
		}
		fmt.Println("")
	}

}
