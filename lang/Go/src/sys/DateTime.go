////////////////////////////////////////////////////////////////////////////
// Porgram: Composition.go
// Purpose: Go composition & field promotion demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: Go by Example
//          https://gobyexample.com/
////////////////////////////////////////////////////////////////////////////

// Go's offers extensive support for times and durations;
// here are some examples.

package main

import (
	"fmt"
	"time"
)

/*

Parsing and formatting date/time in Go
https://pauladamsmith.com/blog/2011/05/go_time.html

Go takes an interesting approach to parsing strings to time objects, and formatting time objects as strings. Instead of using codes like most languages to represent component parts of a date/time string representation—like %Y for a 4-digit year like “2011” or %b for an abbreviated month name like “Feb”—Go uses a mnemonic device: there is a standard time, which is:

    Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)

Or put another way:

    01/02 03:04:05PM '06 -0700

Instead of having to remember or lookup the traditional formatting codes for functions like strftime, you just count one-two-three-four and each place in the standard time corresponds to a component of a date/time object (the Time type in Go): one for day of the month, two for the month, three for the hour (in 12-hour time), four for the minutes, etc.

*/

func test_now() {

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	/*

		Friday
		it's a weekday
		it's after noon

	*/

}

func test_parts() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)
	// 2015-12-04 23:02:51.941177635 -0500 EST

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// 2009-11-17 20:34:58.651387237 +0000 UTC

	// You can extract the various components of the time
	// value as expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// The Monday-Sunday `Weekday` is also available.
	p(then.Weekday())
	/*
		2009
		November
		17
		20
		34
		58
		651387237
		UTC
		Tuesday
	*/

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	/*
		true
		false
		false
	*/

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	p(diff)

	// We can compute the length of the duration in various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	/*
		52999h27m53.289790398s
		52999.46480271956
		3.179967888163173e+06
		1.907980732897904e+08
		190798073289790398
	*/

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	p(then.Add(diff))
	p(then.Add(-diff))
	// 2015-12-05 04:02:51.941177635 +0000 UTC
	// 2003-11-01 13:07:05.361596839 +0000 UTC
}

func test_parse() {
	p := fmt.Println

	// Here's a basic example of formatting a time
	// according to RFC3339.
	t := time.Now()
	p(t.Format("2006-01-02T15:04:05Z07:00"))
	p(t.Format(time.RFC3339))
	// 2015-12-04T23:02:52-05:00
	// 2015-12-04T23:02:52-05:00

	// `Format` uses an example-based layout approach; it
	// takes a formatted version of the reference time
	// `Mon Jan 2 15:04:05 MST 2006` to determine the
	// general pattern with which to format the given
	// time. Here are a few more examples of time
	// formatting.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	/*
		11:02PM
		Fri Dec  4 23:02:52 2015
		2015-12-04T23:02:52.94202-05:00
	*/

	// For purely numeric representations you can also
	// use standard string formatting with the extracted
	// components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// 2015-12-04T23:02:52-00:00

	// Time parsing uses the same example-based approach
	// as `Format`ing. These examples parse times rendered
	// with some of the layouts used above.
	withNanos := "2006-01-02T15:04:05.999999999-07:00"
	t1, e := time.Parse(
		withNanos,
		"2012-11-01T22:08:41.117442+00:00")
	p(t1)
	kitchen := "3:04PM"
	t2, e := time.Parse(kitchen, "8:41PM")
	p(t2)
	// 2012-11-01 22:08:41.117442 +0000 +0000
	// 0000-01-01 20:41:00 +0000 UTC

	// `Parse` will return an error on malformed input
	// explaining the parsing problem.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
	// parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"

	// There are several predefined formats that you can
	// use for both formatting and parsing.
	p(t.Format(time.Kitchen))
	// 11:02PM
}

func main() {
	test_now()
	time.Sleep(time.Second)
	test_parts()
	time.Sleep(time.Second * 1)
	test_parse()
}
