// https://leetcode.com/problems/zigzag-conversion/
// https://medium.com/@acinom/zigzag-conversion-c6827126e941

package main

import (
	"fmt"

	"github.com/suntong/testing"
)

type testCase struct {
	s      string
	nRows  int
	result string
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	var t *testing.T = testing.NewT()
	TestIt(t)
	t.Report()
}

//==========================================================================
// test functions

func TestIt(t *testing.T) {
	testData := []testCase{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"PAYPALISHIRINGPAYPALISHIRING", 4, "PINARALSIGPLIIYAHRPYIHNPIASG"},
		{"THIS_IS_HOW_STRING_IS_ARRANGED", 4, "TSS_RHI__TGIRADI_HWRNSANESOI_G"},
		{"THIS_IS_HOW_STRING_IS_ARRANGED", 5, "THNRH_OIGRAISWR_ANSI_TI_GD_SSE"},
		{"A", 1, "A"},
	}

	for _, tc := range testData {
		mustEqual(t, convert(tc.s, tc.nRows), tc.result)
	}
}

func mustEqual(t *testing.T, res, result string) {
	if res != result {
		t.Errorf(`expected "%s" but got "%s"\n`, result, res)
	} else {
		fmt.Println("matched")
	}
}

func convert(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			//fmt.Print(string(item))
			solution = append(solution, item)
		}
		//fmt.Println()
	}
	return string(solution)
}

/*

T     S     S     _     R
H   I _   _ T   G I   R A   D
I _   H W   R N   S A   N E
S     O     I     _     G

T        H       N       R
H      _ O     I G     R A
I    S   W   R   _   A   N
S  I     _ T     I _     G D
_        S       S       E


*/
