// https://go.dev/play/p/5xgyNYErHSL

package main

import (
	"fmt"
)

// Fields: structure to be printed
type Fields struct {
	Name     string
	NickName string
	Age      int
}

type Pos struct {
	Position string
	Person   Fields
}

type Pos2 struct {
	Position string
	Person   *Fields
}

func main() {

	// initializing the
	// struct with values
	var f = Fields{
		Name:     "Abc",
		NickName: "abc",
		Age:      19,
	}

	// printing the structure
	fmt.Printf("%v\n", f)
	fmt.Printf("%+v\n", f)
	fmt.Printf("%#v\n", f)

	pos := Pos{Position: "Beef-eater", Person: f}
	fmt.Printf("%#v\n", pos)
	p2 := Pos2{Position: "Beef-eater", Person: &f}
	fmt.Printf("%#v\n", p2)
	fmt.Printf("%+v\n", p2)
}

/*

{Abc abc 19}
{Name:Abc NickName:abc Age:19}
main.Fields{Name:"Abc", NickName:"abc", Age:19}
main.Pos{Position:"Beef-eater", Person:main.Fields{Name:"Abc", NickName:"abc", Age:19}}
main.Pos2{Position:"Beef-eater", Person:(*main.Fields)(0xc00008c180)}
{Position:Beef-eater Person:0xc00008c180}

*/
