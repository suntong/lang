////////////////////////////////////////////////////////////////////////////
// Porgram: et_select.go
// Purpose: etree xpath select demo
// Authors: Tong Sun (c) 2019, All rights reserved
// Credits: https://github.com/beevik/etree
//          https://github.com/beevik/etree/blob/master/example_test.go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Function main

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("et_select xml-file xml-sel-str")
		os.Exit(0)
	}
	PathQuery(os.Args[1], os.Args[2])
}

func PathQuery(xmlFile, xmlSelStr string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xmlFile); err != nil {
		panic(err)
	}

	fmt.Println("---------- Original Doc -----------")
	doc.Indent(2)
	// doc.WriteTo(os.Stdout)
	fmt.Printf("-----------------------------------\n\n")

	for _, e := range doc.FindElements(xmlSelStr) {
		fmt.Printf("%s (%s): %s\n", e.Tag, e.GetPath(), e.Text())
	}

}

/*

$ go run et_select.go et_example.xml '//book'
---------- Original Doc -----------
-----------------------------------

book (/bookstore/book):

book (/bookstore/book):

*/
