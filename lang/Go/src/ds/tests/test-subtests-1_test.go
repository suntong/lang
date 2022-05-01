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

// https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721
func TestTLog(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Here you test tc.value against a test function.
			// Let's use t.Log as our test function :-)
			t.Log(tc.value)
		})
	}
}

func TestGroupedParallel(t *testing.T) {
	//t.Parallel()
	/*

	   The outer test will not complete until all parallel tests started by
	   Run have completed. As a result, no other parallel tests can run in
	   parallel to these parallel tests.

	   Note that we need to capture the range variable to ensure that tc gets
	   bound to the correct instance.

	*/
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// Here you test tc.value against a test function.
			// Let's use t.Log as our test function :-)
			t.Log(tc.value)
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


$ go test -v -run=TestTLog test-subtests-1_test.go
=== RUN   TestTLog
=== PAUSE TestTLog
=== CONT  TestTLog
=== RUN   TestTLog/test_1
    test-subtests-1_test.go:60: 1
=== RUN   TestTLog/test_2
    test-subtests-1_test.go:60: 2
=== RUN   TestTLog/test_3
    test-subtests-1_test.go:60: 3
=== RUN   TestTLog/test_4
    test-subtests-1_test.go:60: 4
--- PASS: TestTLog (0.00s)
    --- PASS: TestTLog/test_1 (0.00s)
    --- PASS: TestTLog/test_2 (0.00s)
    --- PASS: TestTLog/test_3 (0.00s)
    --- PASS: TestTLog/test_4 (0.00s)
PASS
ok      command-line-arguments  0.003s


$ go test -v -run=TestGroupedParallel test-subtests-1_test.go
=== RUN   TestGroupedParallel
=== RUN   TestGroupedParallel/test_1
=== PAUSE TestGroupedParallel/test_1
=== RUN   TestGroupedParallel/test_2
=== PAUSE TestGroupedParallel/test_2
=== RUN   TestGroupedParallel/test_3
=== PAUSE TestGroupedParallel/test_3
=== RUN   TestGroupedParallel/test_4
=== PAUSE TestGroupedParallel/test_4
=== CONT  TestGroupedParallel/test_1
=== CONT  TestGroupedParallel/test_4
    test-subtests-1_test.go:92: 4
=== CONT  TestGroupedParallel/test_1
    test-subtests-1_test.go:92: 1
=== CONT  TestGroupedParallel/test_3
    test-subtests-1_test.go:92: 3
=== CONT  TestGroupedParallel/test_2
    test-subtests-1_test.go:92: 2
--- PASS: TestGroupedParallel (0.00s)
    --- PASS: TestGroupedParallel/test_4 (0.00s)
    --- PASS: TestGroupedParallel/test_1 (0.00s)
    --- PASS: TestGroupedParallel/test_3 (0.00s)
    --- PASS: TestGroupedParallel/test_2 (0.00s)
PASS
ok      command-line-arguments  0.004s


*/
