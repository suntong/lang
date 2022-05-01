// https://go.dev/blog/subtests

package main

import (
	"fmt"
	"testing"
	"time"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//==========================================================================
// functions to test

//==========================================================================
// test functions

func TestTime(t *testing.T) {
	testCases := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},
		{"12:31", "America/New_York", "7:31"},
		{"08:08", "Australia/Sydney", "18:08"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
			loc, err := time.LoadLocation(tc.loc)
			if err != nil {
				t.Fatal("could not load location")
			}
			gmt, _ := time.Parse("15:04", tc.gmt)
			if got := gmt.In(loc).Format("15:04"); got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

/*

$ go test -v test-subtests-1_test.go
=== RUN   TestTime
=== RUN   TestTime/12:31_in_Europe/Zuri
    test-subtests-1_test.go:36: could not load location
=== RUN   TestTime/12:31_in_America/New_York
    test-subtests-1_test.go:40: got 07:34; want 7:31
=== RUN   TestTime/08:08_in_Australia/Sydney
    test-subtests-1_test.go:40: got 18:12; want 18:08
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
    --- FAIL: TestTime/08:08_in_Australia/Sydney (0.00s)
FAIL
FAIL    command-line-arguments  0.005s
FAIL


$ go test -v -run=TestTime/"in Europe" test-subtests-1_test.go
=== RUN   TestTime
=== RUN   TestTime/12:31_in_Europe/Zuri
    test-subtests-1_test.go:36: could not load location
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
FAIL
FAIL    command-line-arguments  0.003s
FAIL


$ go test -v -run=TestTime//New_York test-subtests-1_test.go
=== RUN   TestTime
=== RUN   TestTime/12:31_in_America/New_York
    test-subtests-1_test.go:40: got 07:34; want 7:31
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
FAIL
FAIL    command-line-arguments  0.003s
FAIL


   Note the // in the string passed to -run. The / in time zone name
   America/New_York is handled as if it were a separator resulting from a
   subtest. The first regular expression of the pattern (TestTime) matches
   the top-level test. The second regular expression (the empty string)
   matches anything, in this case the time and the continent part of the
   location. The third regular expression (New_York) matches the city part
   of the location.

   Treating slashes in names as separators allows the user to refactor
   hierarchies of tests without the need to change the naming. It also
   simplifies the escaping rules. The user should escape slashes in names,
   for instance by replacing them with backslashes, if this poses a problem.

   A unique sequence number is appended to test names that are not
   unique. So one could just pass an empty string to Run if there is no
   obvious naming scheme for subtests and the subtests can easily be
   identified by their sequence number.


*/
