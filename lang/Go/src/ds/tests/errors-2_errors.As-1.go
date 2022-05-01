// https://gosamples.dev/check-error-type/

package main

import (
	"errors"
	"fmt"
)

const badInput = "abc"

type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input: %s", e.input)
}

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf("validateInput: %w", &BadInputError{input: input})
	}
	return nil
}

// https://medium.com/@felipedutratine/golang-how-to-handle-errors-in-v1-13-fda7f035d027
type myError struct {
	err  string
	more string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%s: %s", e.more, e.err)
}
func x() error {
	return fmt.Errorf("adding more context: %w", &myError{"error", "more"})
}

func main() {
	{
		e := x()

		var err *myError
		if ok := errors.As(e, &err); ok {
			// handle gracefully
			fmt.Printf("Catching '%s' error for '%s' successfully: %s\n",
				err.err, err.more, err.Error())
		}
	}

	input := badInput

	err := validateInput(input)
	var badInputErr *BadInputError
	if errors.As(err, &badInputErr) {
		fmt.Printf("[Err] Bad input error occured: %s\n", badInputErr)
	}
}

/*

   the As(err error, target interface{}) bool checks if any error in the
   chain of wrapped errors matches the target. The difference is that this
   function checks whether the error has a specific type, unlike the Is(),
   which examines if it is a particular error object. Because As considers
   the whole chain of errors, it should be preferable to the type assertion
   if e, ok := err.(*BadInputError); ok.

   target argument of the As(err error, target interface{}) bool function
   should be a pointer to the error type, which in this case is
   *BadInputError

*/
