////////////////////////////////////////////////////////////////////////////
// Porgram: lx2_search-node.go
// Purpose: gokogiri search doc/node by xml
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: tjkopena
//          https://github.com/moovweb/gokogiri/issues/15#issuecomment-13992329
////////////////////////////////////////////////////////////////////////////

// demo where the context is when doing Node.Search()

package main

import (
	"fmt"
	"github.com/moovweb/gokogiri"
)

var file = "<foo><bar></bar</foo>"

func main() {
	doc, _ := gokogiri.ParseXml([]byte(file))

	nodes, _ := doc.Search("/foo")
	for n := range nodes {
		fmt.Println(nodes[n].Name())
		subnodes, _ := nodes[n].Search("bar")
		for s := range subnodes {
			fmt.Println(subnodes[s].Name())
		}
	}

	fmt.Println("---")

	nodes, _ = doc.Search("bar")
	for n := range nodes {
		fmt.Println(nodes[n].Name())
	}
}
