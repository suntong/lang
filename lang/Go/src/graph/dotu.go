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
	g, err := dot.ParseFile("dotu.dot")
	if err != nil {
		panic(err)
	}
	s := g.String()
	fmt.Println(s)
}

/*


 */
