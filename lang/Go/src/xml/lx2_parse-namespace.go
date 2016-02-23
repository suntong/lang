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

/*

https://github.com/moovweb/gokogiri/issues/77#issuecomment-66941278

Do not ue SetNamespace, because the SetNamespace call is for when you're building documents, not for querying.

Try using doc.EvalXPath(x, nil) instead of Search. I think that EvalXPath automatically registers any namespaces that occur in the document.

Alternatively you can get the XPath context directly (via doc.DocXPathCtx) and call RegisterNamespace followed by EvaluateAsNodeset on the context object ( this is what EvalXPath does in the background, see https://github.com/moovweb/gokogiri/blob/master/xml/node.go#L650 )

*/
