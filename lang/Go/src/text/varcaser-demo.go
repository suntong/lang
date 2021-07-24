package main

import (
	"fmt"

	"github.com/danverbraganza/varcaser/varcaser"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// pre-configed varcaser transformers
// the names are self-explanatory from their definitions
var (
	Ck2lc_ = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerCamelCase}
	Ck2uc_ = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCase}

	Css2uc = varcaser.Caser{
		From: varcaser.ScreamingSnakeCase, To: varcaser.UpperCamelCaseKeepCaps}

)

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	ExampleVaribleNames()
}

func ExampleVaribleNames() {
	fmt.Printf("%v\n", Ck2lc_.String("some-init-method"))
	fmt.Printf("%v\n", Ck2lc("some-init-method"))
	fmt.Printf("%s %s %s %s %s",
		varcaser.Caser{From: varcaser.LowerCamelCase, To: varcaser.KebabCase}.
			String("someInitMethod"),
		varcaser.Caser{From: varcaser.LowerCamelCase,
			To: varcaser.ScreamingSnakeCase}.
			String("myConstantVariable"),
		Ck2lc_.String("some-init-method"),
		Ck2lc("some-init-method"),
		Ck2uc_.String("some-init-method"))

	fmt.Printf("\n%v\n", Css2uc.String("MY_CONSTANT_VARIABLE"))
	// Output:
	// someInitMethod
	// someInitMethod
	// some-init-method MY_CONSTANT_VARIABLE someInitMethod someInitMethod SomeInitMethod
	// MyConstantVariable
}

//==========================================================================
// template function
// the names are self-explanatory from the var definitions

func Ck2lc(in string) string { return Ck2lc_.String(in) }
