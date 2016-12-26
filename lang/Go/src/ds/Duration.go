// https://tour.golang.org/moretypes/5
// https://godoc.org/time

package main

import (
	"fmt"
	"time"
)

func main() {
	test1()
	test2()
	test3()
	test4()
}

func test1() {
	// ParseDuration
	var passTime time.Duration
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "1000"))
	fmt.Printf("Delay: %v\n", passTime)
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "20"))
	fmt.Printf("Delay: %v\n", passTime)
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "1200"))
	fmt.Printf("Delay: %v\n", passTime)
	fmt.Println()
}

func test2() {
	// == Multiple and division
	var passTime time.Duration
	passTime, _ = time.ParseDuration(
		fmt.Sprintf("%ss", "1200"))
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
	fmt.Println()
}

func test3() {
	// Difference in days, simple solution
	// By Dave Cheney https://play.golang.org/p/qbx_J1jVhV
	d1, _ := time.Parse(time.RFC3339, "2015-06-27T18:24:49Z")
	fmt.Println(d1)
	d2, _ := time.Parse(time.RFC3339, "2016-05-18T15:07:40Z")
	fmt.Println(d2)
	delta := d2.Sub(d1)
	fmt.Printf("difference was %d days\n", int(delta.Hours()/24))
	// Acutally better do int(d2.Sub(d1) / (24 * time.Hour))
	// https://groups.google.com/d/msg/golang-nuts/O2NaRAH94GI/w_pcfI1AKE8J
	fmt.Println()
}

func test4() {
	// Difference in days, convoluted solution
	d1, _ := time.Parse(time.RFC3339, "2015-06-27T18:24:49Z")
	fmt.Println(d1)
	d2, _ := time.Parse(time.RFC3339, "2016-05-18T15:07:40Z")
	fmt.Println(d2)
	fmt.Println(daysDiff(d2, d1))

}

// a - b in days
// By Karthik Krishnamurthy https://play.golang.org/p/nTcjGZQKAa
func daysDiff(a, b time.Time) (days int) {
	cur := b
	for cur.Year() < a.Year() {
		// add 1 to count the last day of the year too.
		days += lastDayOfYear(cur).YearDay() - cur.YearDay() + 1
		cur = firstDayOfNextYear(cur)
	}
	days += a.YearDay() - cur.YearDay()
	if b.AddDate(0, 0, days).After(a) {
		days -= 1
	}
	return days
}

func lastDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 0, 0, 0, 0, t.Location())
}

func firstDayOfNextYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
}
