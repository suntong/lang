package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	errorsUnwrap()
	customError()
	httpError()
}

// errors.Unwrap, https://stackoverflow.com/a/58887639/2125837
func errorsUnwrap() {
	// error 1 wrapped by 2 wrapped by 3
	fmt.Println("\tWrapping ------")
	err1 := errors.New("error 1")
	err2 := fmt.Errorf("error 2: %w", err1)
	err3 := fmt.Errorf("error 3: %w", err2)

	fmt.Println(err1) // "error 1"
	fmt.Println(err2) // "error 2: error 1"
	fmt.Println(err3) // "error 3: error 2: error 1"

	// unwrap peels a layer off
	fmt.Println("\n\tUnwrapping ------")
	fmt.Println(errors.Unwrap(err3))                // "error 2: error 1"
	fmt.Println(errors.Unwrap(errors.Unwrap(err3))) // "error 1"

	// recursively unwrap
	fmt.Println("\n\tRecursively unwrapping ------")
	currentErr := err3
	for errors.Unwrap(currentErr) != nil {
		currentErr = errors.Unwrap(currentErr)
	}

	fmt.Println(currentErr) // "error 1"
}

// implement a custom error

type CustomError struct {
	LowLevel      error
	HumanReadable string
}

func (e CustomError) Error() string {
	return e.HumanReadable
}

func (e CustomError) Unwrap() error {
	return e.LowLevel
}

// helper
func NewCustomError(inner error, outer string) *CustomError {
	return &CustomError{
		LowLevel:      inner,
		HumanReadable: outer,
	}
}

func customError() {
	fmt.Println("\n\t Custom error ------")
	err0 := errors.New("db record create failure")
	err := NewCustomError(err0, "failed to create book")

	fmt.Println(errors.Unwrap(err)) // internally log low-level error
	fmt.Println(err)                // present human-readable error to user

}

// HTTP error
const (
	StatusBadRequest = 400
)

var (
	ErrBadRequest = NewError(StatusBadRequest) // 400
)

// Error represents an error that occurred while handling a request.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error makes it compatible with the `error` interface.
func (e *Error) Error() string {
	return e.Message
}

// NewError creates a new Error instance with an optional message
func NewError(code int, message ...string) *Error {
	err := &Error{
		Code:    code,
		Message: "Bad Request",
	}
	return err
}

// WrappedError wrap an error with HTTP return code
type WrappedError struct {
	Type *Error
	err  error
}

func httpError() {
	fmt.Println("\n\t HTTP error ------")
	{
		// using custom error
		err := ErrBadRequest
		fmt.Printf("%d: %+v\n\n", err.Code, err)
	}

	{
		// using custom error
		err0 := errors.New("Unmarshal body failure")
		err := fmt.Errorf("%w: Invalid payload body.", err0)
		fmt.Println(errors.Unwrap(err)) // low-level error
		fmt.Println(err)                // wrapped error

		ret := &WrappedError{Type: ErrBadRequest, err: err}
		fmt.Printf("\n%d %+v - %+v\n\n", ret.Type.Code, ret.Type, err)
		//fmt.Printf("It's a Bad Request: %+v \n", ret.Type)
	}

	{
		// using plain error & wrapping more info
		err0 := errors.New(strconv.Itoa(StatusBadRequest))
		err := fmt.Errorf("Error %w: Invalid payload body.", err0)
		fmt.Printf("(%T) %[1]s\n", errors.Unwrap(err))         // low-level error
		fmt.Printf("(%T) %[1]s\n", errors.Unwrap(err).Error()) // low-level error
		fmt.Println(err)                                       // wrapped error
	}
}

/*

The last HTTP error output:

400: Bad Request

Unmarshal body failure
Unmarshal body failure: Invalid payload body.

400 Bad Request - Unmarshal body failure: Invalid payload body.

(*errors.errorString) 400
(string) 400
Error 400: Invalid payload body.

*/
