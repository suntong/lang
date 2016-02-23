////////////////////////////////////////////////////////////////////////////
// Porgram: lx2_parse-namespace.go
// Purpose: parse xml with a namespace using gokogiri (libxml2)
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Paul Crovella
//          http://stackoverflow.com/a/27475227/2125837
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

var a = `<?xml version="1.0" ?><NodeA xmlns="http://example.com/this"><NodeB>thisthat</NodeB></NodeA>`

func main() {
	doc, _ := gokogiri.ParseXml([]byte(a))
	defer doc.Free()
	xp := doc.DocXPathCtx()
	xp.RegisterNamespace("ns", "http://example.com/this")
	x := xpath.Compile("/ns:NodeA/ns:NodeB")
	groups, err := doc.Search(x)
	if err != nil {
		fmt.Println(err)
	}
	for i, group := range groups {
		fmt.Println(i, group.Content())
	}
}
