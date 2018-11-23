// https://blog.golang.org/strings
// https://play.golang.org/p/w7HQAx_rsgx
// https://play.golang.org/p/XmLKa0OpCg8

package main

import "fmt"

func main() {
	BinaryString()
	UTF8StringLiteral()
}

func BinaryString() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}

/*

Println:
��=� ⌘
Byte loop:
bd b2 3d bc 20 e2 8c 98 
Printf with %x:
bdb23dbc20e28c98
Printf with % x:
bd b2 3d bc 20 e2 8c 98
Printf with %q:
"\xbd\xb2=\xbc ⌘"
Printf with %+q:
"\xbd\xb2=\xbc \u2318"

*/


func UTF8StringLiteral() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

/*

plain string: ⌘
quoted string: "\u2318"
hex bytes: e2 8c 98 

*/
