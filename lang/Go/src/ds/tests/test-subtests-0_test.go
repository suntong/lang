// https://ieftimov.com/posts/testing-in-go-subtests/
// https://ieftimov.com/posts/testing-in-go-table-driven-tests/
//
// See also:
//   https://www.infoworld.com/article/3633659/get-started-with-go-testing.html

package main

import (
	"errors"

	"testing"
	//"github.com/suntong/testing"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

// func main() {
// 	var t *testing.T = testing.NewT()
// 	TestOlder(t)
// 	t.Report()
// }

//==========================================================================
// functions to test

var (
	AgeTooLowError  = errors.New("A person must be at least 1 years old")
	AgeTooHighError = errors.New("A person cannot be older than 130 years")
)

type Person struct {
	age int
}

func NewPerson(age int) (error, *Person) {
	if age < 1 {
		return AgeTooLowError, nil
	}

	if age >= 130 {
		return AgeTooHighError, nil
	}

	return nil, &Person{
		age: age,
	}
}

func (p *Person) older(other *Person) bool {
	return p.age > other.age
}

//==========================================================================
// test functions

func setupSubtest(t *testing.T) {
	t.Logf("[SETUP] Hello")
}

func teardownSubtest(t *testing.T) {
	t.Logf("[TEARDOWN] Bye, bye!")
}

func TestOlder(t *testing.T) {
	cases := []struct {
		name     string
		age1     int
		age2     int
		expected bool
	}{
		{
			name:     "FirstOlderThanSecond",
			age1:     1,
			age2:     2,
			expected: false,
		},
		{
			name:     "SecondOlderThanFirst",
			age1:     2,
			age2:     1,
			expected: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setupSubtest(t)
			defer teardownSubtest(t)

			_, p1 := NewPerson(c.age1)
			_, p2 := NewPerson(c.age2)
			//t.Logf("] %v vs %v \n", p1, p2)

			got := p1.older(p2)

			t.Logf("[TEST] Hello from subtest %s \n", c.name)
			if got != c.expected {
				t.Errorf("Expected %v > %v, got %v", p1.age, p2.age, got)
			}
		})
	}

}

/*

$ go test -v test-subtests-0_test.go
=== RUN   TestOlder
=== RUN   TestOlder/FirstOlderThanSecond
    test-subtests-0_test.go:59: [SETUP] Hello
    test-subtests-0_test.go:102: [TEST] Hello from subtest FirstOlderThanSecond
    test-subtests-0_test.go:63: [TEARDOWN] Bye, bye!
=== RUN   TestOlder/SecondOlderThanFirst
    test-subtests-0_test.go:59: [SETUP] Hello
    test-subtests-0_test.go:102: [TEST] Hello from subtest SecondOlderThanFirst
    test-subtests-0_test.go:63: [TEARDOWN] Bye, bye!
--- PASS: TestOlder (0.00s)
    --- PASS: TestOlder/FirstOlderThanSecond (0.00s)
    --- PASS: TestOlder/SecondOlderThanFirst (0.00s)
PASS
ok      command-line-arguments  0.004s

*/
