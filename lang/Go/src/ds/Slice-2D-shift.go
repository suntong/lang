// https://leetcode.com/problems/shift-2d-grid/

package main

import "fmt"

var grid [][]int = [][]int{
	[]int{3, 8, 1, 9},
	[]int{19, 7, 2, 5},
	[]int{4, 6, 11, 10},
	[]int{12, 0, 21, 13},
}

func main() {
	fmt.Println(grid)
	fmt.Println()
	for k := 1; k <= 4; k++ {
		fmt.Println(shiftGrid(grid, k))
		fmt.Println(shiftGrid2(grid, k))
	}
}

func shiftGrid(grid [][]int, k int) [][]int {
	x, y := len(grid[0]), len(grid)
	// Create a new 2D grid
	newGrid := make([][]int, y)
	for i := 0; i < y; i++ {
		newGrid[i] = make([]int, x)
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			ny := (k / x) + i
			if (j + (k % x)) >= x {
				ny++
			}
			newGrid[ny%y][(j+(k%x))%x] = grid[i][j]
		}
	}
	return newGrid
}

func shiftGrid2(grid [][]int, k int) [][]int {
	newGrid := slice2Dto1D(grid)
	x, y := len(grid[0]), len(grid)
	ng := newGrid[x*y-k:]
	return slice1Dto2D(append(ng, newGrid[:x*y-k]...), x, y)
}

func slice2Dto1D(grid [][]int) []int {
	x, y := len(grid[0]), len(grid)
	newGrid := make([]int, x*y)
	for i := 0; i < y; i++ {
		//newGrid[x*i : x*(i+1)] = grid[i] // cannot assign to
		for j := 0; j < x; j++ {
			newGrid[x*i+j] = grid[i][j]
		}
	}
	return newGrid
}

func slice1Dto2D(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = original[n*i : n*(i+1)]
	}
	return res
}
