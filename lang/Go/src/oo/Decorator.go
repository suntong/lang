////////////////////////////////////////////////////////////////////////////
// Porgram: Decorator.go
// Purpose: Go Decorator functions demo
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: As listed below
////////////////////////////////////////////////////////////////////////////

/*

Decorator functions in Go
https://stackoverflow.com/questions/45944781/decorator-functions-in-go

Q: Decorator pattern (functions) has many benefits:

- It is very useful when a method has many orthogonal concerns... I.e., None of these concerns are related, other than that we wanna do all (or some) of them whenever we call our method. This is where the decorator pattern really helps.
- By implementing the decorator pattern we subscribe to the open-closed principal. Our method is open to future extension but closed to future modification. There's a lot of groovy benefits to obeying the open-closed principle.

However, all the examples that I found are really complicated (e.g., writing HTTP servers with many middlewares). This make it difficult for me to apply the principle elsewhere. I need something that I can easily try on so as to wrap my head around it.

I need something that can illustrate this:

func Decorate(c Decorated, ds ...Decorator) Decorated {
    decorated := c
    for _, decorate := range ds {
        decorated = decorate(decorated)
    }
    return decorated
}

Can someone give me an simpler example that can best illustrate how to do Decorator pattern (functions) in Go please?

A:

A decorator can basically be a function that takes another function of a specific type as its argument and returns a function of the a same type. This essentially allows you to create a chain of functions. So in Go it would look something like this:

// this is the type of functions you want to decorate
type StringManipulator func(string) string

// this is your decorator.
func ToLower(m StringManipulator) StringManipulator {
    return func(s string) string {
        lower := strings.ToLower(s)
        return m(lower)
    }
}

here's a more complete example
https://play.golang.org/p/p3d61uLGhd

mkopriva
2017-08-29

*/

package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type StringManipulator func(string) string

type StringManipulatorDecorator func(StringManipulator) StringManipulator

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	DecoratorTest()
	OrderDemo()
}

//==========================================================================
// support functions

func ToLower(m StringManipulator) StringManipulator {
	return func(s string) string {
		lower := strings.ToLower(s)
		return m(lower)
	}
}

func ToBase64(m StringManipulator) StringManipulator {
	return func(s string) string {
		b64 := base64.StdEncoding.EncodeToString([]byte(s))
		return m(b64)
	}
}

func AppendDecorator(suffix string) StringManipulatorDecorator {
	return func(m StringManipulator) StringManipulator {
		return func(s string) string {
			return m(s + suffix)
		}
	}
}

func PrependDecorator(prefix string) StringManipulatorDecorator {
	return func(m StringManipulator) StringManipulator {
		return func(s string) string {
			return m(prefix + s)
		}
	}
}

// "identity" just return the same string
func ident(s string) string {
	return s
}

// MultiDecorator "merges" the passed in decorators and returns a singe decorator.
func MultiDecorator(ds ...StringManipulatorDecorator) StringManipulatorDecorator {
	return func(m StringManipulator) StringManipulator {
		for _, d := range ds {
			m = d(m)
		}
		return m
	}
}

/*

If you use the decorator pattern as is demonstrated in my answer you need to
pass a function to the decorators. This gets weird when you want to loop
over StringManipulators, because the "decorators" aren't "manipulators" they
just return them, so you need to pass a manipulator to the decorator to get
the decorated manipulator... if it sounds confusing, that's because it is :)

*/

func DecoratorTest() {
	s := "Hello, playground"

	var fn StringManipulator = ident
	fmt.Println(fn(s))

	fn = ToBase64(ToLower(fn))
	fmt.Println(fn(s))

	var fn2 StringManipulator = ident
	fn2 = ToLower(ToBase64(fn2))
	fmt.Println(fn2(s))

	var fn3 StringManipulator = ident
	fn3 = AppendDecorator(" -GOLANG")(ToLower(PrependDecorator("DECORATED: ")(fn3)))
	fmt.Println(fn3(s))

	// dec is now a StringManipulatorDecorator, to use it, you still need to
	// pass it the function of type StringManipulator that you want to decorate.
	dec := MultiDecorator(
		AppendDecorator(" -GOLANG"),
		ToLower,
		PrependDecorator("DECORATED: "),
	)

	fn4 := dec(ident)
	fmt.Println(fn4(s))
}

/*

Hello, playground
sgvsbg8sihbsyxlncm91bmq=
aGVsbG8sIHBsYXlncm91bmQ=
DECORATED: hello, playground -golang
decorated: hello, playground -GOLANG

*/

////////////////////////////////////////////////////////////////////////////
// Order discussion

/*

Q: For the expression

AppendDecorator(" -GOLANG")(ToLower(PrependDecorator("DECORATED: ", fn3)),

intuitive thinking is that the inner most function get applied first,
however, the test result show that it is actually the outer most function
get applied first. That's totally against my intuition. Can you explain a
bit why ToLower is applied to AppendDecorator but not the PrependDecorator,
which is the inner function, please?

A: Here's a demonstration of the order play.golang.org/p/zj-wnlXS2O hope
that helps to explain it a little.


*/

func A(s string) string {
	fmt.Println("calling A")
	return s + "A"
}

func B(s string) string {
	fmt.Println("calling B")
	return s + "B"
}

func C(s string) string {
	fmt.Println("calling C")
	return s + "C"
}

type F func(string) string

func a(f F) F {
	return func(s string) string {
		fmt.Println("calling a")
		return f(s + "a")
	}
}

func b(f F) F {
	return func(s string) string {
		fmt.Println("calling b")
		return f(s + "b")
	}
}

func c(f F) F {
	return func(s string) string {
		fmt.Println("calling c")
		return f(s + "c")
	}
}

func OrderDemo() {
	// Here the innermost gets the original string, and is called
	// and evaluated first, then its result is passed to the second
	// function and so on, resulting in a seemingly "reverse order".
	fmt.Println(A(B(C(""))))

	var id F = func(s string) string { return s }

	// here we're first constructing the abc chain and executing
	// it only at the end with ...)(""), so the first F that's called,
	// which is also the first which manipulates the passed in
	// string, is the one returned by the decorator a.
	fmt.Println(a(b(c(id)))(""))
}

/*

calling C
calling B
calling A
CBA
calling a
calling b
calling c
abc

*/
