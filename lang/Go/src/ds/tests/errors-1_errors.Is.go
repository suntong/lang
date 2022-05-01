// https://gosamples.dev/check-error-type/

package main

import (
	"errors"
	"fmt"
)

const badInput = "abc"

var ErrBadInput = errors.New("bad input")

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf(`validateInput ("%s"): %w`, input, ErrBadInput)
	}
	return nil
}

func main() {
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
