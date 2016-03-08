////////////////////////////////////////////////////////////////////////////
// Porgram: et_et_dump.go
// Purpose: etree example
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://github.com/beevik/etree
////////////////////////////////////////////////////////////////////////////

package main

import "os"

import (
	"github.com/beevik/etree"
)

func main() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("et_dump.xml"); err != nil {
		panic(err)
	}

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}
