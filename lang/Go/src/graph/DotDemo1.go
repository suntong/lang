////////////////////////////////////////////////////////////////////////////
// Porgram: DotDemo1.go
// Purpose: Demo the .dot (graphviz) file handling in GO
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: github.com/mewspring/dot
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/mewspring/dot"
)

func main() {
	ExampleNewGraph()
	ExampleRead()
}

func ExampleRead() {
	g, err := dot.Read([]byte(`digraph G {Hello->World}`))
	if err != nil {
		panic(err)
	}
	s := g.String()
	fmt.Println(s)
}

func ExampleNewGraph() {
	g := dot.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddNode("G", "Hello", nil)
	g.AddNode("G", "World", nil)
	g.AddEdge("Hello", "World", true, nil)
	s := g.String()
	fmt.Println(s)
}

/*

digraph G {
        World;
        Hello;
        Hello->World;

}

digraph G {
        Hello;
        World;
        Hello->World;

}


*/
