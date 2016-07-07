// https://tour.golang.org/moretypes/5

package main

import (
	"fmt"
	"time"
)

func main() {
	var passTime time.Duration
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "1000"))
	fmt.Printf("Delay: %v\n", passTime)
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "1200"))
	fmt.Printf("Delay: %v\n", passTime)
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "20"))

	// == Multiple and division
	fmt.Printf("Delay*3: %v\n", passTime*3)
	fmt.Printf("Delay/2: %v\n", passTime/2)
	divideBy := 2
	//fmt.Printf("Delay/2: %v\n", passTime/divideBy)
	// invalid operation: passTime / divideBy (mismatched types time.Duration and int)
	fmt.Printf("Delay/2: %v\n", passTime/time.Duration(divideBy))

	// == Conver the divide result to float
	fmt.Printf("%v\n", float32(3/2))          // Output: 1
	fmt.Printf("%v\n", float32(float32(3)/2)) // Output: 1.5

	// Type convert *prior* to the division. Dan Kortschak
	hourTime, _ := time.ParseDuration("1h")
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "68"))
	fmt.Printf("%v\n", float32(hourTime)/float32(passTime))

	// Or, use the built in methods to get seconds in floating point. Jakob Borg
	t0 := time.Hour
	t1 := 68 * time.Second
	fmt.Println(t0.Seconds() / t1.Seconds())
}
