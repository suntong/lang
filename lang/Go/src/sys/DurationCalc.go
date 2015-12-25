////////////////////////////////////////////////////////////////////////////
// Porgram: DurationCalc.go
// Purpose: Go time.Duration calculation demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: Chris Kastorff
//   https://groups.google.com/d/msg/golang-nuts/DgAvKK9uVC0/P7j-YG8bDQAJ
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"time"
)

func main() {

	// you can do math on them with the builtin operators just like an int64,
	// and you can use any untyped integral constants as time.Duration values

	took := time.Second * 20
	estimate := took * 15
	fmt.Println("estimate:", estimate)

	took = time.Second * 120
	estimate = took / 2
	fmt.Println("estimate:", estimate)

	// or convert some integral to time.Duration, like time.Duration(value)
	// to use it in a calculation.

	startTime := time.Now()
	time.Sleep(2 * time.Second)
	fmt.Printf("\nElapsed %s\n", time.Since(startTime))
	fmt.Println("New estimate:", time.Since(startTime)*15)
	var times int64 = 15
	fmt.Println("New estimate:", time.Since(startTime)*time.Duration(times))

	// Depending on the calculations, you might want to convert everything to
	// float64 and then back again to avoid truncation on integer division and
	// integer overflow problems.

	/*

		Algorithm explanation:

		For easy of calculation and understanding, It assumes at *each* step of
		the loop, the elapsed time is 60 seconds (same), and try to estimate the
		remaining time.

		For e.g., when ii==2, two jobs finished in 60 seconds, so the remaining
		3 jobs should finish in 60/2*3=90 seconds, not 1m. And for ii==3, I'm
		expecting 60/3*2=40 seconds, not 0 seconds.

		The wrong values are when I was doing

		time.Duration(int(float32(took)*float32(1.0*(total-ii)/ii))))

	*/

	took = time.Second * 60
	fmt.Printf("Time taken so far %s\n", took)
	total := 5
	for ii := 1; ii <= total; ii++ {
		fmt.Printf("Finishing the remaining %d%% in %s\n",
			(total-ii)*100/total,
			time.Duration(int(float32(took)*float32(total-ii)/float32(ii))))
	}
}
