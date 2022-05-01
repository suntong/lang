// https://gosamples.dev/check-error-type/

package main

import (
	"errors"
	"fmt"
)

const badInput = "abc"

var ErrBadInput = errors.New("BadInput")

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf(`validateInput ("%s"): %w`, input, ErrBadInput)
	}
	return nil
}

// https://medium.com/@felipedutratine/golang-how-to-handle-errors-in-v1-13-fda7f035d027
// https://go.dev/play/p/m5DmFnHsbo9
var e1 = errors.New("InternalError1")

func x() error {
	return fmt.Errorf("adding more context: %w", e1)
}

func main() {
	e := x()
	if errors.Is(e, e1) { // Magical it works
		// handle gracefully
		fmt.Printf("Catching successfully: %s\n", e)
	}

	input := badInput

	err := validateInput(input)
	if errors.Is(err, ErrBadInput) {
		fmt.Printf("[Err] Bad input error: %s\n", err)
	}
}

/*

   the function validateInput returns an error for badInput. This error is
   ErrBadInput wrapped in an error created by fmt.Errorf(). Using the
   Is(err, target error) bool function, we can detect the ErrBadInput even
   if it is wrapped since this function checks if any error in the chain of
   wrapped errors matches the target. Therefore, this form should be
   preferable to comparison if err == ErrBadInput.

*/
